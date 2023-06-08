package docker_importer

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"free-ent-guide-backend/models"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"gorm.io/gorm/clause"
)

// GetNetwork finds the proper network to attach to the container.
// This should exist and connect to outside the host.
func GetNetwork(cli *client.Client) (string, error) {
	ctx := context.Background()
	f := filters.NewArgs()
	// This `common` network should have been manually created.
	f.Add("name", "common")
	nets, err := cli.NetworkList(ctx, types.NetworkListOptions{Filters: f})
	if err != nil {
		return "", err
	}

	if len(nets) == 0 {
		return "", errors.New("docker network `common` not found")
	}

	// Assume only one network is found.
	log.Printf("docker network found %s", nets[0].ID)
	return nets[0].ID, nil
}

// ImportNHL runs the Docker container to fetch NHL game schedule
// for a given date, and import games and teams to the DB.
func ImportNHL(store models.Store, startDate string) error {
	image := "brianwagner/python-nhlgames:1.2"
	contName := "nhlGamesPython"
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Print(err)
		return err
	}

	// @todo is there a better way to reuse the container?
	f := filters.NewArgs()
	f.Add("name", fmt.Sprintf("/%s", contName))
	f.Add("status", "exited")
	ctrs, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: f})
	if err != nil {
		log.Print(err)
		return err
	}

	for _, v := range ctrs {
		err = cli.ContainerStop(ctx, v.ID, container.StopOptions{})
		if err != nil {
			log.Print(err)
			return err
		}
		log.Printf("removing old container %s", v.ID)
		err = cli.ContainerRemove(ctx, v.ID, types.ContainerRemoveOptions{})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		log.Print(err)
		return err
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	platform := v1.Platform{OS: "linux", Architecture: "amd64"}

	// Check for custom network which allows internet outside of host.
	netID, err := GetNetwork(cli)
	if err != nil {
		return err
	}

	cont, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			Tty:   true,
			Env:   []string{fmt.Sprintf("START_DATE=%s", startDate)},
		},
		&container.HostConfig{
			PortBindings: portBinding,
			NetworkMode:  "common",
		},
		nil,
		&platform,
		contName,
	)
	if err != nil {
		log.Print(err)
		return err
	}

	err = cli.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	contJSON, err := cli.ContainerInspect(ctx, cont.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("Container starting: %s for date %s\n", contJSON.Name, startDate)

	statusCh, errCh := cli.ContainerWait(ctx, cont.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Print(err)
			return err
		}
	case <-statusCh:
		// fmt.Println("Container is ok")
	}

	// Attach container to network to allow internet access.
	err = cli.NetworkConnect(ctx, netID, cont.ID, &network.EndpointSettings{NetworkID: netID})
	if err != nil {
		return err
	}

	out, err := cli.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
	})
	if err != nil {
		log.Print(err)
		return err
	}

	defer out.Close()

	count := 0
	db := store.DB

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		var g models.NHLGame
		err := json.Unmarshal(scanner.Bytes(), &g)
		if err != nil {
			log.Print(err)
			continue
		}

		home := g.Home
		tx := db.FirstOrCreate(&home, models.NHLTeam{TeamID: g.Home.TeamID})
		if tx.Error != nil {
			log.Print(tx.Error)
			continue
		}
		g.HomeID = home.ID

		visitor := g.Visitor
		tx = db.FirstOrCreate(&visitor, models.NHLTeam{TeamID: g.Visitor.TeamID})
		if tx.Error != nil {
			log.Print(tx.Error)
			continue
		}
		g.VisitorID = visitor.ID

		tx = db.Omit(clause.Associations).Create(&g)
		if tx.Error != nil {
			log.Print(tx.Error)
			return tx.Error
		}
		count++
	}

	log.Printf("nhl import complete, adding %d games", count)

	// fmt.Print(games)

	// Clean up.
	err = cli.ContainerStop(ctx, cont.ID, container.StopOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	err = cli.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// ImportMLB runs the Docker container to fetch MLB game schedule
// for a given date, and import games and teams to the DB.
func ImportMLB(store models.Store, startDate string) error {
	image := "brianwagner/python-mlbgames:1.2"
	contName := "mlbGamesPython"
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Print(err)
		return err
	}

	// @todo is there a better way to reuse the container?
	f := filters.NewArgs()
	f.Add("name", fmt.Sprintf("/%s", contName))
	f.Add("status", "exited")
	ctrs, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: f})
	if err != nil {
		log.Print(err)
		return err
	}

	for _, v := range ctrs {
		err = cli.ContainerStop(ctx, v.ID, container.StopOptions{})
		if err != nil {
			log.Print(err)
			return err
		}
		log.Printf("removing old container %s", v.ID)
		err = cli.ContainerRemove(ctx, v.ID, types.ContainerRemoveOptions{})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		log.Print(err)
		return err
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	platform := v1.Platform{OS: "linux", Architecture: "amd64"}

	// Check for custom network which allows internet outside of host.
	netID, err := GetNetwork(cli)
	if err != nil {
		return err
	}

	cont, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			Tty:   true,
			Env:   []string{fmt.Sprintf("START_DATE=%s", startDate)},
		},
		&container.HostConfig{
			PortBindings: portBinding,
			NetworkMode:  "common",
		},
		nil,
		&platform,
		contName,
	)
	if err != nil {
		log.Print(err)
		return err
	}

	err = cli.ContainerStart(ctx, cont.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	contJSON, err := cli.ContainerInspect(ctx, cont.ID)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("Container starting: %s for date %s\n", contJSON.Name, startDate)

	statusCh, errCh := cli.ContainerWait(ctx, cont.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Print(err)
			return err
		}
	case <-statusCh:
		// fmt.Println("Container is ok")
	}

	// Attach container to network to allow internet access.
	err = cli.NetworkConnect(ctx, netID, cont.ID, &network.EndpointSettings{NetworkID: netID})
	if err != nil {
		return err
	}

	out, err := cli.ContainerLogs(ctx, cont.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		Follow:     true,
	})
	if err != nil {
		log.Print(err)
		return err
	}

	defer out.Close()

	count := 0
	db := store.DB

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		var g models.MLBGame
		err := json.Unmarshal(scanner.Bytes(), &g)
		if err != nil {
			log.Print(err)
			continue
		}
		// games = append(games, g)

		home := g.Home
		tx := db.FirstOrCreate(&home, models.MLBTeam{TeamID: g.Home.TeamID})
		if tx.Error != nil {
			log.Print(tx.Error)
			continue
		}
		g.HomeID = home.ID

		visitor := g.Visitor
		tx = db.FirstOrCreate(&visitor, models.MLBTeam{TeamID: g.Visitor.TeamID})
		if tx.Error != nil {
			log.Print(tx.Error)
			continue
		}
		g.VisitorID = visitor.ID

		tx = db.Omit(clause.Associations).Create(&g)
		if tx.Error != nil {
			log.Print(tx.Error)
			return tx.Error
		}
		count++
	}

	log.Printf("mlb import complete, adding %d games", count)

	// fmt.Print(games)

	// Clean up.
	err = cli.ContainerStop(ctx, cont.ID, container.StopOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	err = cli.ContainerRemove(ctx, cont.ID, types.ContainerRemoveOptions{})
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
