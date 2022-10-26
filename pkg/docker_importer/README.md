Steps to make this work

1. install docker
1. above should create a usergroup `docker`. Now add the app user to this group. https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket
1. create a custom network. The default bridge network won't give access to the host internet. https://superuser.com/questions/1130898/no-internet-connection-inside-docker-containers/1582710#1582710