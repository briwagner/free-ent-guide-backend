package models

import (
	"encoding/json"
	"free-ent-guide-backend/pkg/cred"
	"testing"
)

func init() {
	c := cred.Cred{DB: "gorm:password@tcp(127.0.0.1:3306)/demo_gorm_test?charset=utf8mb4&parseTime=True&loc=Local"}
	AppStore = Setup(c)
}

func TestUpdateGameNHL(t *testing.T) {
	data := `{
		"copyright" : "NHL and the NHL Shield are registered trademarks of the National Hockey League. NHL and NHL team marks are the property of the NHL and its teams. © NHL 2022. All Rights Reserved.",
		"gamePk" : 2022020148,
		"link" : "/api/v1/game/2022020148/feed/live",
		"metaData" : {
			"wait" : 10,
			"timeStamp" : "20221102_005519"
		},
		"gameData" : {
			"game" : {
				"pk" : 2022020148,
				"season" : "20222023",
				"type" : "R"
			},
			"datetime" : {
				"dateTime" : "2022-11-01T23:00:00Z"
			},
			"status" : {
				"abstractGameState" : "Live",
				"codedGameState" : "3",
				"detailedState" : "In Progress",
				"statusCode" : "3",
				"startTimeTBD" : false
			},
			"teams" : {
				"away" : {
					"id" : 9,
					"name" : "Ottawa Senators",
					"link" : "/api/v1/teams/9",
					"venue" : {
						"id" : 5031,
						"name" : "Canadian Tire Centre",
						"link" : "/api/v1/venues/5031",
						"city" : "Ottawa",
						"timeZone" : {
							"id" : "America/New_York",
							"offset" : -4,
							"tz" : "EDT"
						}
					},
					"abbreviation" : "OTT",
					"triCode" : "OTT",
					"teamName" : "Senators",
					"locationName" : "Ottawa",
					"firstYearOfPlay" : "1992",
					"division" : {
						"id" : 17,
						"name" : "Atlantic",
						"nameShort" : "ATL",
						"link" : "/api/v1/divisions/17",
						"abbreviation" : "A"
					},
					"conference" : {
						"id" : 6,
						"name" : "Eastern",
						"link" : "/api/v1/conferences/6"
					},
					"franchise" : {
						"franchiseId" : 30,
						"teamName" : "Senators",
						"link" : "/api/v1/franchises/30"
					},
					"shortName" : "Ottawa",
					"officialSiteUrl" : "http://www.ottawasenators.com/",
					"franchiseId" : 30,
					"active" : true
				},
				"home" : {
					"id" : 14,
					"name" : "Tampa Bay Lightning",
					"link" : "/api/v1/teams/14",
					"venue" : {
						"name" : "AMALIE Arena",
						"link" : "/api/v1/venues/null",
						"city" : "Tampa",
						"timeZone" : {
							"id" : "America/New_York",
							"offset" : -4,
							"tz" : "EDT"
						}
					},
					"abbreviation" : "TBL",
					"triCode" : "TBL",
					"teamName" : "Lightning",
					"locationName" : "Tampa Bay",
					"firstYearOfPlay" : "1992",
					"division" : {
						"id" : 17,
						"name" : "Atlantic",
						"nameShort" : "ATL",
						"link" : "/api/v1/divisions/17",
						"abbreviation" : "A"
					},
					"conference" : {
						"id" : 6,
						"name" : "Eastern",
						"link" : "/api/v1/conferences/6"
					},
					"franchise" : {
						"franchiseId" : 31,
						"teamName" : "Lightning",
						"link" : "/api/v1/franchises/31"
					},
					"shortName" : "Tampa Bay",
					"officialSiteUrl" : "http://www.tampabaylightning.com/",
					"franchiseId" : 31,
					"active" : true
				}
			},
			"players" : {
				"ID8480073" : {
					"id" : 8480073,
					"fullName" : "Erik Brannstrom",
					"link" : "/api/v1/people/8480073",
					"firstName" : "Erik",
					"lastName" : "Brannstrom",
					"primaryNumber" : "26",
					"birthDate" : "1999-09-02",
					"currentAge" : 23,
					"birthCity" : "Eksjo",
					"birthCountry" : "SWE",
					"nationality" : "SWE",
					"height" : "5' 10\"",
					"weight" : 185,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8481043" : {
					"id" : 8481043,
					"fullName" : "Cole Koepke",
					"link" : "/api/v1/people/8481043",
					"firstName" : "Cole",
					"lastName" : "Koepke",
					"primaryNumber" : "45",
					"birthDate" : "1998-05-17",
					"currentAge" : 24,
					"birthCity" : "Two Harbors",
					"birthStateProvince" : "MN",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 1\"",
					"weight" : 196,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : true,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8473512" : {
					"id" : 8473512,
					"fullName" : "Claude Giroux",
					"link" : "/api/v1/people/8473512",
					"firstName" : "Claude",
					"lastName" : "Giroux",
					"primaryNumber" : "28",
					"birthDate" : "1988-01-12",
					"currentAge" : 34,
					"birthCity" : "Hearst",
					"birthStateProvince" : "ON",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "5' 11\"",
					"weight" : 190,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8477938" : {
					"id" : 8477938,
					"fullName" : "Haydn Fleury",
					"link" : "/api/v1/people/8477938",
					"firstName" : "Haydn",
					"lastName" : "Fleury",
					"primaryNumber" : "7",
					"birthDate" : "1996-07-08",
					"currentAge" : 26,
					"birthCity" : "Carlyle",
					"birthStateProvince" : "SK",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 4\"",
					"weight" : 208,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8470880" : {
					"id" : 8470880,
					"fullName" : "Brian Elliott",
					"link" : "/api/v1/people/8470880",
					"firstName" : "Brian",
					"lastName" : "Elliott",
					"primaryNumber" : "1",
					"birthDate" : "1985-04-09",
					"currentAge" : 37,
					"birthCity" : "Newmarket",
					"birthStateProvince" : "ON",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 2\"",
					"weight" : 221,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "G",
						"name" : "Goalie",
						"type" : "Goalie",
						"abbreviation" : "G"
					}
				},
				"ID8474207" : {
					"id" : 8474207,
					"fullName" : "Nick Holden",
					"link" : "/api/v1/people/8474207",
					"firstName" : "Nick",
					"lastName" : "Holden",
					"primaryNumber" : "5",
					"birthDate" : "1987-05-15",
					"currentAge" : 35,
					"birthCity" : "St. Albert",
					"birthStateProvince" : "AB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 4\"",
					"weight" : 210,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8470621" : {
					"id" : 8470621,
					"fullName" : "Corey Perry",
					"link" : "/api/v1/people/8470621",
					"firstName" : "Corey",
					"lastName" : "Perry",
					"primaryNumber" : "10",
					"birthDate" : "1985-05-16",
					"currentAge" : 37,
					"birthCity" : "Peterborough",
					"birthStateProvince" : "ON",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 3\"",
					"weight" : 208,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8476453" : {
					"id" : 8476453,
					"fullName" : "Nikita Kucherov",
					"link" : "/api/v1/people/8476453",
					"firstName" : "Nikita",
					"lastName" : "Kucherov",
					"primaryNumber" : "86",
					"birthDate" : "1993-06-17",
					"currentAge" : 29,
					"birthCity" : "Maykop",
					"birthCountry" : "RUS",
					"nationality" : "RUS",
					"height" : "5' 11\"",
					"weight" : 181,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8480448" : {
					"id" : 8480448,
					"fullName" : "Parker Kelly",
					"link" : "/api/v1/people/8480448",
					"firstName" : "Parker",
					"lastName" : "Kelly",
					"primaryNumber" : "45",
					"birthDate" : "1999-05-14",
					"currentAge" : 23,
					"birthCity" : "Camrose",
					"birthStateProvince" : "AB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 0\"",
					"weight" : 190,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8479525" : {
					"id" : 8479525,
					"fullName" : "Ross Colton",
					"link" : "/api/v1/people/8479525",
					"firstName" : "Ross",
					"lastName" : "Colton",
					"primaryNumber" : "79",
					"birthDate" : "1996-09-11",
					"currentAge" : 26,
					"birthCity" : "Robbinsville",
					"birthStateProvince" : "NJ",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 0\"",
					"weight" : 194,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8475167" : {
					"id" : 8475167,
					"fullName" : "Victor Hedman",
					"link" : "/api/v1/people/8475167",
					"firstName" : "Victor",
					"lastName" : "Hedman",
					"primaryNumber" : "77",
					"birthDate" : "1990-12-18",
					"currentAge" : 31,
					"birthCity" : "Ornskoldsvik",
					"birthCountry" : "SWE",
					"nationality" : "SWE",
					"height" : "6' 7\"",
					"weight" : 244,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8480801" : {
					"id" : 8480801,
					"fullName" : "Brady Tkachuk",
					"link" : "/api/v1/people/8480801",
					"firstName" : "Brady",
					"lastName" : "Tkachuk",
					"primaryNumber" : "7",
					"birthDate" : "1999-09-16",
					"currentAge" : 23,
					"birthCity" : "Scottsdale",
					"birthStateProvince" : "AZ",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 4\"",
					"weight" : 221,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8476433" : {
					"id" : 8476433,
					"fullName" : "Magnus Hellberg",
					"link" : "/api/v1/people/8476433",
					"firstName" : "Magnus",
					"lastName" : "Hellberg",
					"primaryNumber" : "39",
					"birthDate" : "1991-04-04",
					"currentAge" : 31,
					"birthCity" : "Uppsala",
					"birthCountry" : "SWE",
					"nationality" : "SWE",
					"height" : "6' 6\"",
					"weight" : 220,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "G",
						"name" : "Goalie",
						"type" : "Goalie",
						"abbreviation" : "G"
					}
				},
				"ID8479984" : {
					"id" : 8479984,
					"fullName" : "Cal Foote",
					"link" : "/api/v1/people/8479984",
					"firstName" : "Cal",
					"lastName" : "Foote",
					"primaryNumber" : "52",
					"birthDate" : "1998-12-13",
					"currentAge" : 23,
					"birthCity" : "Denver",
					"birthStateProvince" : "CO",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 5\"",
					"weight" : 224,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8480208" : {
					"id" : 8480208,
					"fullName" : "Drake Batherson",
					"link" : "/api/v1/people/8480208",
					"firstName" : "Drake",
					"lastName" : "Batherson",
					"primaryNumber" : "19",
					"birthDate" : "1998-04-27",
					"currentAge" : 24,
					"birthCity" : "Fort Wayne",
					"birthStateProvince" : "IN",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 3\"",
					"weight" : 200,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8479542" : {
					"id" : 8479542,
					"fullName" : "Brandon Hagel",
					"link" : "/api/v1/people/8479542",
					"firstName" : "Brandon",
					"lastName" : "Hagel",
					"primaryNumber" : "38",
					"birthDate" : "1998-08-27",
					"currentAge" : 24,
					"birthCity" : "Saskatoon",
					"birthStateProvince" : "SK",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 2\"",
					"weight" : 179,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8474013" : {
					"id" : 8474013,
					"fullName" : "Ian Cole",
					"link" : "/api/v1/people/8474013",
					"firstName" : "Ian",
					"lastName" : "Cole",
					"primaryNumber" : "28",
					"birthDate" : "1989-02-21",
					"currentAge" : 33,
					"birthCity" : "Ann Arbor",
					"birthStateProvince" : "MI",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 1\"",
					"weight" : 225,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8479026" : {
					"id" : 8479026,
					"fullName" : "Philippe Myers",
					"link" : "/api/v1/people/8479026",
					"firstName" : "Philippe",
					"lastName" : "Myers",
					"primaryNumber" : "5",
					"birthDate" : "1997-01-25",
					"currentAge" : 25,
					"birthCity" : "Moncton",
					"birthStateProvince" : "NB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 6\"",
					"weight" : 219,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8474034" : {
					"id" : 8474034,
					"fullName" : "Pat Maroon",
					"link" : "/api/v1/people/8474034",
					"firstName" : "Pat",
					"lastName" : "Maroon",
					"primaryNumber" : "14",
					"birthDate" : "1988-04-23",
					"currentAge" : 34,
					"birthCity" : "St. Louis",
					"birthStateProvince" : "MO",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 3\"",
					"weight" : 234,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8478472" : {
					"id" : 8478472,
					"fullName" : "Mathieu Joseph",
					"link" : "/api/v1/people/8478472",
					"firstName" : "Mathieu",
					"lastName" : "Joseph",
					"primaryNumber" : "21",
					"birthDate" : "1997-02-09",
					"currentAge" : 25,
					"birthCity" : "Laval",
					"birthStateProvince" : "QC",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 1\"",
					"weight" : 187,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8479580" : {
					"id" : 8479580,
					"fullName" : "Dylan Gambrell",
					"link" : "/api/v1/people/8479580",
					"firstName" : "Dylan",
					"lastName" : "Gambrell",
					"primaryNumber" : "27",
					"birthDate" : "1996-08-26",
					"currentAge" : 26,
					"birthCity" : "Bonney Lake",
					"birthStateProvince" : "WA",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "5' 11\"",
					"weight" : 185,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8481596" : {
					"id" : 8481596,
					"fullName" : "Shane Pinto",
					"link" : "/api/v1/people/8481596",
					"firstName" : "Shane",
					"lastName" : "Pinto",
					"primaryNumber" : "57",
					"birthDate" : "2000-11-12",
					"currentAge" : 21,
					"birthCity" : "Franklin Square",
					"birthStateProvince" : "NY",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 3\"",
					"weight" : 201,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : true,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8478010" : {
					"id" : 8478010,
					"fullName" : "Brayden Point",
					"link" : "/api/v1/people/8478010",
					"firstName" : "Brayden",
					"lastName" : "Point",
					"primaryNumber" : "21",
					"birthDate" : "1996-03-13",
					"currentAge" : 26,
					"birthCity" : "Calgary",
					"birthStateProvince" : "AB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "5' 11\"",
					"weight" : 180,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8482105" : {
					"id" : 8482105,
					"fullName" : "Jake Sanderson",
					"link" : "/api/v1/people/8482105",
					"firstName" : "Jake",
					"lastName" : "Sanderson",
					"primaryNumber" : "85",
					"birthDate" : "2002-07-08",
					"currentAge" : 20,
					"birthCity" : "Whitefish",
					"birthStateProvince" : "MT",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 3\"",
					"weight" : 195,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : true,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8480246" : {
					"id" : 8480246,
					"fullName" : "Nick Perbix",
					"link" : "/api/v1/people/8480246",
					"firstName" : "Nick",
					"lastName" : "Perbix",
					"primaryNumber" : "48",
					"birthDate" : "1998-06-15",
					"currentAge" : 24,
					"birthCity" : "Minneapolis",
					"birthStateProvince" : "MN",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 2\"",
					"weight" : 191,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : true,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8473544" : {
					"id" : 8473544,
					"fullName" : "Derick Brassard",
					"link" : "/api/v1/people/8473544",
					"firstName" : "Derick",
					"lastName" : "Brassard",
					"primaryNumber" : "61",
					"birthDate" : "1987-09-22",
					"currentAge" : 35,
					"birthCity" : "Hull",
					"birthStateProvince" : "QC",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 1\"",
					"weight" : 200,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8477426" : {
					"id" : 8477426,
					"fullName" : "Nicholas Paul",
					"link" : "/api/v1/people/8477426",
					"firstName" : "Nicholas",
					"lastName" : "Paul",
					"primaryNumber" : "20",
					"birthDate" : "1995-03-20",
					"currentAge" : 27,
					"birthCity" : "Mississauga",
					"birthStateProvince" : "ON",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 3\"",
					"weight" : 223,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8478416" : {
					"id" : 8478416,
					"fullName" : "Erik Cernak",
					"link" : "/api/v1/people/8478416",
					"firstName" : "Erik",
					"lastName" : "Cernak",
					"primaryNumber" : "81",
					"birthDate" : "1997-05-28",
					"currentAge" : 25,
					"birthCity" : "Kosice",
					"birthCountry" : "SVK",
					"nationality" : "SVK",
					"height" : "6' 4\"",
					"weight" : 224,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8475766" : {
					"id" : 8475766,
					"fullName" : "Austin Watson",
					"link" : "/api/v1/people/8475766",
					"firstName" : "Austin",
					"lastName" : "Watson",
					"primaryNumber" : "16",
					"birthDate" : "1992-01-13",
					"currentAge" : 30,
					"birthCity" : "Ann Arbor",
					"birthStateProvince" : "MI",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 4\"",
					"weight" : 204,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8474612" : {
					"id" : 8474612,
					"fullName" : "Travis Hamonic",
					"link" : "/api/v1/people/8474612",
					"firstName" : "Travis",
					"lastName" : "Hamonic",
					"primaryNumber" : "23",
					"birthDate" : "1990-08-16",
					"currentAge" : 32,
					"birthCity" : "St. Malo",
					"birthStateProvince" : "MB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 1\"",
					"weight" : 200,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8473986" : {
					"id" : 8473986,
					"fullName" : "Alex Killorn",
					"link" : "/api/v1/people/8473986",
					"firstName" : "Alex",
					"lastName" : "Killorn",
					"primaryNumber" : "17",
					"birthDate" : "1989-09-14",
					"currentAge" : 33,
					"birthCity" : "Halifax",
					"birthStateProvince" : "NS",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 1\"",
					"weight" : 194,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8479337" : {
					"id" : 8479337,
					"fullName" : "Alex DeBrincat",
					"link" : "/api/v1/people/8479337",
					"firstName" : "Alex",
					"lastName" : "DeBrincat",
					"primaryNumber" : "12",
					"birthDate" : "1997-12-18",
					"currentAge" : 24,
					"birthCity" : "Farmington Hills",
					"birthStateProvince" : "MI",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "5' 8\"",
					"weight" : 178,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "R",
						"name" : "Right Wing",
						"type" : "Forward",
						"abbreviation" : "RW"
					}
				},
				"ID8479458" : {
					"id" : 8479458,
					"fullName" : "Nikita Zaitsev",
					"link" : "/api/v1/people/8479458",
					"firstName" : "Nikita",
					"lastName" : "Zaitsev",
					"primaryNumber" : "22",
					"birthDate" : "1991-10-29",
					"currentAge" : 31,
					"birthCity" : "Moscow",
					"birthCountry" : "RUS",
					"nationality" : "RUS",
					"height" : "6' 2\"",
					"weight" : 192,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8476883" : {
					"id" : 8476883,
					"fullName" : "Andrei Vasilevskiy",
					"link" : "/api/v1/people/8476883",
					"firstName" : "Andrei",
					"lastName" : "Vasilevskiy",
					"primaryNumber" : "88",
					"birthDate" : "1994-07-25",
					"currentAge" : 28,
					"birthCity" : "Tyumen",
					"birthCountry" : "RUS",
					"nationality" : "RUS",
					"height" : "6' 4\"",
					"weight" : 220,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "G",
						"name" : "Goalie",
						"type" : "Goalie",
						"abbreviation" : "G"
					}
				},
				"ID8478469" : {
					"id" : 8478469,
					"fullName" : "Thomas Chabot",
					"link" : "/api/v1/people/8478469",
					"firstName" : "Thomas",
					"lastName" : "Chabot",
					"primaryNumber" : "72",
					"birthDate" : "1997-01-30",
					"currentAge" : 25,
					"birthCity" : "Sainte-Marie",
					"birthStateProvince" : "QC",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 2\"",
					"weight" : 195,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8477930" : {
					"id" : 8477930,
					"fullName" : "Pierre-Edouard Bellemare",
					"link" : "/api/v1/people/8477930",
					"firstName" : "Pierre-Edouard",
					"lastName" : "Bellemare",
					"primaryNumber" : "41",
					"birthDate" : "1985-03-06",
					"currentAge" : 37,
					"birthCity" : "Le Blanc-Mesnil",
					"birthCountry" : "FRA",
					"nationality" : "FRA",
					"height" : "5' 11\"",
					"weight" : 198,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8474564" : {
					"id" : 8474564,
					"fullName" : "Steven Stamkos",
					"link" : "/api/v1/people/8474564",
					"firstName" : "Steven",
					"lastName" : "Stamkos",
					"primaryNumber" : "91",
					"birthDate" : "1990-02-07",
					"currentAge" : 32,
					"birthCity" : "Markham",
					"birthStateProvince" : "ON",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 0\"",
					"weight" : 182,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : true,
					"rookie" : false,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8479410" : {
					"id" : 8479410,
					"fullName" : "Mikhail Sergachev",
					"link" : "/api/v1/people/8479410",
					"firstName" : "Mikhail",
					"lastName" : "Sergachev",
					"primaryNumber" : "98",
					"birthDate" : "1998-06-25",
					"currentAge" : 24,
					"birthCity" : "Nizhnekamsk",
					"birthCountry" : "RUS",
					"nationality" : "RUS",
					"height" : "6' 1\"",
					"weight" : 217,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8476341" : {
					"id" : 8476341,
					"fullName" : "Anton Forsberg",
					"link" : "/api/v1/people/8476341",
					"firstName" : "Anton",
					"lastName" : "Forsberg",
					"primaryNumber" : "31",
					"birthDate" : "1992-11-27",
					"currentAge" : 29,
					"birthCity" : "Härnösand",
					"birthCountry" : "SWE",
					"nationality" : "SWE",
					"height" : "6' 3\"",
					"weight" : 193,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "G",
						"name" : "Goalie",
						"type" : "Goalie",
						"abbreviation" : "G"
					}
				},
				"ID8477353" : {
					"id" : 8477353,
					"fullName" : "Tyler Motte",
					"link" : "/api/v1/people/8477353",
					"firstName" : "Tyler",
					"lastName" : "Motte",
					"primaryNumber" : "14",
					"birthDate" : "1995-03-10",
					"currentAge" : 27,
					"birthCity" : "St. Clair",
					"birthStateProvince" : "MI",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "5' 10\"",
					"weight" : 192,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8477471" : {
					"id" : 8477471,
					"fullName" : "Dillon Heatherington",
					"link" : "/api/v1/people/8477471",
					"firstName" : "Dillon",
					"lastName" : "Heatherington",
					"primaryNumber" : "29",
					"birthDate" : "1995-05-09",
					"currentAge" : 27,
					"birthCity" : "Calgary",
					"birthStateProvince" : "AB",
					"birthCountry" : "CAN",
					"nationality" : "CAN",
					"height" : "6' 4\"",
					"weight" : 215,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "D",
						"name" : "Defenseman",
						"type" : "Defenseman",
						"abbreviation" : "D"
					}
				},
				"ID8476480" : {
					"id" : 8476480,
					"fullName" : "Vladislav Namestnikov",
					"link" : "/api/v1/people/8476480",
					"firstName" : "Vladislav",
					"lastName" : "Namestnikov",
					"primaryNumber" : "90",
					"birthDate" : "1992-11-22",
					"currentAge" : 29,
					"birthCity" : "Zhukovskiy",
					"birthCountry" : "RUS",
					"nationality" : "RUS",
					"height" : "6' 0\"",
					"weight" : 179,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				},
				"ID8482116" : {
					"id" : 8482116,
					"fullName" : "Tim Stützle",
					"link" : "/api/v1/people/8482116",
					"firstName" : "Tim",
					"lastName" : "Stützle",
					"primaryNumber" : "18",
					"birthDate" : "2002-01-15",
					"currentAge" : 20,
					"birthCity" : "Viersen",
					"birthCountry" : "DEU",
					"nationality" : "DEU",
					"height" : "6' 0\"",
					"weight" : 193,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : false,
					"shootsCatches" : "L",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "L",
						"name" : "Left Wing",
						"type" : "Forward",
						"abbreviation" : "LW"
					}
				},
				"ID8480355" : {
					"id" : 8480355,
					"fullName" : "Mark Kastelic",
					"link" : "/api/v1/people/8480355",
					"firstName" : "Mark",
					"lastName" : "Kastelic",
					"primaryNumber" : "47",
					"birthDate" : "1999-03-11",
					"currentAge" : 23,
					"birthCity" : "Phoenix",
					"birthStateProvince" : "AZ",
					"birthCountry" : "USA",
					"nationality" : "USA",
					"height" : "6' 3\"",
					"weight" : 233,
					"active" : true,
					"alternateCaptain" : false,
					"captain" : false,
					"rookie" : true,
					"shootsCatches" : "R",
					"rosterStatus" : "Y",
					"currentTeam" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					},
					"primaryPosition" : {
						"code" : "C",
						"name" : "Center",
						"type" : "Forward",
						"abbreviation" : "C"
					}
				}
			},
			"venue" : {
				"id" : 5017,
				"name" : "Amalie Arena",
				"link" : "/api/v1/venues/5017"
			}
		},
		"liveData" : {
			"plays" : {
				"allPlays" : [ {
					"result" : {
						"event" : "Game Scheduled",
						"eventCode" : "TBL1",
						"eventTypeId" : "GAME_SCHEDULED",
						"description" : "Game Scheduled"
					},
					"about" : {
						"eventIdx" : 0,
						"eventId" : 1,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-01T22:13:59Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"result" : {
						"event" : "Period Ready",
						"eventCode" : "TBL5",
						"eventTypeId" : "PERIOD_READY",
						"description" : "Period Ready"
					},
					"about" : {
						"eventIdx" : 1,
						"eventId" : 5,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-01T23:04:34Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"result" : {
						"event" : "Period Start",
						"eventCode" : "TBL51",
						"eventTypeId" : "PERIOD_START",
						"description" : "Period Start"
					},
					"about" : {
						"eventIdx" : 2,
						"eventId" : 51,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-01T23:10:00Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL52",
						"eventTypeId" : "FACEOFF",
						"description" : "Mark Kastelic faceoff won against Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 3,
						"eventId" : 52,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-01T23:10:00Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 0.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL8",
						"eventTypeId" : "HIT",
						"description" : "Ross Colton hit Travis Hamonic"
					},
					"about" : {
						"eventIdx" : 4,
						"eventId" : 8,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:08",
						"periodTimeRemaining" : "19:52",
						"dateTime" : "2022-11-01T23:10:08Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -6.0,
						"y" : 40.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL151",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Jake Sanderson shot blocked shot by Erik Cernak"
					},
					"about" : {
						"eventIdx" : 5,
						"eventId" : 151,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:16",
						"periodTimeRemaining" : "19:44",
						"dateTime" : "2022-11-01T23:10:16Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -75.0,
						"y" : -4.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL10",
						"eventTypeId" : "HIT",
						"description" : "Mark Kastelic hit Erik Cernak"
					},
					"about" : {
						"eventIdx" : 6,
						"eventId" : 10,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:18",
						"periodTimeRemaining" : "19:42",
						"dateTime" : "2022-11-01T23:10:18Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -96.0,
						"y" : 1.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL11",
						"eventTypeId" : "HIT",
						"description" : "Ross Colton hit Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 7,
						"eventId" : 11,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:24",
						"periodTimeRemaining" : "19:36",
						"dateTime" : "2022-11-01T23:10:24Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -83.0,
						"y" : -34.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL356",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Tyler Motte shot blocked shot by Erik Cernak"
					},
					"about" : {
						"eventIdx" : 8,
						"eventId" : 356,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:26",
						"periodTimeRemaining" : "19:34",
						"dateTime" : "2022-11-01T23:10:26Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -71.0,
						"y" : 9.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL12",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Cole Koepke Wide of Net Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 9,
						"eventId" : 12,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:37",
						"periodTimeRemaining" : "19:23",
						"dateTime" : "2022-11-01T23:10:37Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 52.0,
						"y" : 28.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL352",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Pierre-Edouard Bellemare shot blocked shot by Travis Hamonic"
					},
					"about" : {
						"eventIdx" : 10,
						"eventId" : 352,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "00:41",
						"periodTimeRemaining" : "19:19",
						"dateTime" : "2022-11-01T23:10:41Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 56.0,
						"y" : 24.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL153",
						"eventTypeId" : "SHOT",
						"description" : "Nicholas Paul Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 11,
						"eventId" : 153,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:06",
						"periodTimeRemaining" : "18:54",
						"dateTime" : "2022-11-01T23:11:06Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 9.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL13",
						"eventTypeId" : "HIT",
						"description" : "Thomas Chabot hit Nicholas Paul"
					},
					"about" : {
						"eventIdx" : 12,
						"eventId" : 13,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:24",
						"periodTimeRemaining" : "18:36",
						"dateTime" : "2022-11-01T23:11:24Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 66.0,
						"y" : -37.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL14",
						"eventTypeId" : "STOP",
						"description" : "Puck Frozen"
					},
					"about" : {
						"eventIdx" : 13,
						"eventId" : 14,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:28",
						"periodTimeRemaining" : "18:32",
						"dateTime" : "2022-11-01T23:11:28Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL53",
						"eventTypeId" : "FACEOFF",
						"description" : "Dylan Gambrell faceoff won against Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 14,
						"eventId" : 53,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:28",
						"periodTimeRemaining" : "18:32",
						"dateTime" : "2022-11-01T23:12:08Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL15",
						"eventTypeId" : "HIT",
						"description" : "Parker Kelly hit Cal Foote"
					},
					"about" : {
						"eventIdx" : 15,
						"eventId" : 15,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:40",
						"periodTimeRemaining" : "18:20",
						"dateTime" : "2022-11-01T23:12:20Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -98.0,
						"y" : 3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL16",
						"eventTypeId" : "HIT",
						"description" : "Pierre-Edouard Bellemare hit Thomas Chabot"
					},
					"about" : {
						"eventIdx" : 16,
						"eventId" : 16,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:49",
						"periodTimeRemaining" : "18:11",
						"dateTime" : "2022-11-01T23:12:29Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 89.0,
						"y" : 35.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474034,
							"fullName" : "Pat Maroon",
							"link" : "/api/v1/people/8474034"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8475766,
							"fullName" : "Austin Watson",
							"link" : "/api/v1/people/8475766"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL17",
						"eventTypeId" : "PENALTY",
						"description" : "Pat Maroon Fighting against Austin Watson",
						"secondaryType" : "Fighting",
						"penaltySeverity" : "Major",
						"penaltyMinutes" : 5
					},
					"about" : {
						"eventIdx" : 17,
						"eventId" : 17,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:57",
						"periodTimeRemaining" : "18:03",
						"dateTime" : "2022-11-01T23:12:37Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 48.0,
						"y" : 40.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8475766,
							"fullName" : "Austin Watson",
							"link" : "/api/v1/people/8475766"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8474034,
							"fullName" : "Pat Maroon",
							"link" : "/api/v1/people/8474034"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL20",
						"eventTypeId" : "PENALTY",
						"description" : "Austin Watson Fighting against Pat Maroon",
						"secondaryType" : "Fighting",
						"penaltySeverity" : "Major",
						"penaltyMinutes" : 5
					},
					"about" : {
						"eventIdx" : 18,
						"eventId" : 20,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:57",
						"periodTimeRemaining" : "18:03",
						"dateTime" : "2022-11-01T23:12:37Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 46.0,
						"y" : 39.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8482116,
							"fullName" : "Tim Stützle",
							"link" : "/api/v1/people/8482116"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL54",
						"eventTypeId" : "FACEOFF",
						"description" : "Brandon Hagel faceoff won against Tim Stützle"
					},
					"about" : {
						"eventIdx" : 19,
						"eventId" : 54,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:57",
						"periodTimeRemaining" : "18:03",
						"dateTime" : "2022-11-01T23:14:04Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL154",
						"eventTypeId" : "SHOT",
						"description" : "Nikita Kucherov Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 20,
						"eventId" : 154,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "01:58",
						"periodTimeRemaining" : "18:02",
						"dateTime" : "2022-11-01T23:14:05Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 50.0,
						"y" : -6.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL23",
						"eventTypeId" : "STOP",
						"description" : "Offside"
					},
					"about" : {
						"eventIdx" : 21,
						"eventId" : 23,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:08",
						"periodTimeRemaining" : "17:52",
						"dateTime" : "2022-11-01T23:14:14Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL55",
						"eventTypeId" : "FACEOFF",
						"description" : "Brayden Point faceoff won against Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 22,
						"eventId" : 55,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:08",
						"periodTimeRemaining" : "17:52",
						"dateTime" : "2022-11-01T23:14:42Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL155",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Nikita Kucherov shot blocked shot by Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 23,
						"eventId" : 155,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:15",
						"periodTimeRemaining" : "17:45",
						"dateTime" : "2022-11-01T23:14:49Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 76.0,
						"y" : 1.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL24",
						"eventTypeId" : "HIT",
						"description" : "Erik Cernak hit Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 24,
						"eventId" : 24,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:27",
						"periodTimeRemaining" : "17:33",
						"dateTime" : "2022-11-01T23:15:01Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 19.0,
						"y" : -38.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL156",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Travis Hamonic shot blocked shot by Ian Cole"
					},
					"about" : {
						"eventIdx" : 25,
						"eventId" : 156,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:31",
						"periodTimeRemaining" : "17:29",
						"dateTime" : "2022-11-01T23:15:05Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -68.0,
						"y" : 7.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL25",
						"eventTypeId" : "STOP",
						"description" : "Icing"
					},
					"about" : {
						"eventIdx" : 26,
						"eventId" : 25,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:32",
						"periodTimeRemaining" : "17:28",
						"dateTime" : "2022-11-01T23:15:06Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8473512,
							"fullName" : "Claude Giroux",
							"link" : "/api/v1/people/8473512"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL56",
						"eventTypeId" : "FACEOFF",
						"description" : "Claude Giroux faceoff won against Brayden Point"
					},
					"about" : {
						"eventIdx" : 27,
						"eventId" : 56,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:32",
						"periodTimeRemaining" : "17:28",
						"dateTime" : "2022-11-01T23:15:38Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL26",
						"eventTypeId" : "STOP",
						"description" : "Puck in Netting"
					},
					"about" : {
						"eventIdx" : 28,
						"eventId" : 26,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:36",
						"periodTimeRemaining" : "17:24",
						"dateTime" : "2022-11-01T23:15:42Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8473512,
							"fullName" : "Claude Giroux",
							"link" : "/api/v1/people/8473512"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL57",
						"eventTypeId" : "FACEOFF",
						"description" : "Nicholas Paul faceoff won against Claude Giroux"
					},
					"about" : {
						"eventIdx" : 29,
						"eventId" : 57,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:36",
						"periodTimeRemaining" : "17:24",
						"dateTime" : "2022-11-01T23:16:08Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL27",
						"eventTypeId" : "HIT",
						"description" : "Travis Hamonic hit Steven Stamkos"
					},
					"about" : {
						"eventIdx" : 30,
						"eventId" : 27,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:55",
						"periodTimeRemaining" : "17:05",
						"dateTime" : "2022-11-01T23:16:27Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 95.0,
						"y" : 17.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480246,
							"fullName" : "Nick Perbix",
							"link" : "/api/v1/people/8480246"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL157",
						"eventTypeId" : "SHOT",
						"description" : "Nick Perbix Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 31,
						"eventId" : 157,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "02:57",
						"periodTimeRemaining" : "17:03",
						"dateTime" : "2022-11-01T23:16:29Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 39.0,
						"y" : -7.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL159",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Nicholas Paul shot blocked shot by Travis Hamonic"
					},
					"about" : {
						"eventIdx" : 32,
						"eventId" : 159,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "03:06",
						"periodTimeRemaining" : "16:54",
						"dateTime" : "2022-11-01T23:16:38Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 71.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL158",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Alex Killorn shot blocked shot by Alex DeBrincat"
					},
					"about" : {
						"eventIdx" : 33,
						"eventId" : 158,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "03:09",
						"periodTimeRemaining" : "16:51",
						"dateTime" : "2022-11-01T23:16:41Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 77.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL28",
						"eventTypeId" : "HIT",
						"description" : "Alex DeBrincat hit Mikhail Sergachev"
					},
					"about" : {
						"eventIdx" : 34,
						"eventId" : 28,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "03:18",
						"periodTimeRemaining" : "16:42",
						"dateTime" : "2022-11-01T23:16:50Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -97.0,
						"y" : -4.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL160",
						"eventTypeId" : "SHOT",
						"description" : "Nicholas Paul Backhand saved by Anton Forsberg",
						"secondaryType" : "Backhand"
					},
					"about" : {
						"eventIdx" : 35,
						"eventId" : 160,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "03:23",
						"periodTimeRemaining" : "16:37",
						"dateTime" : "2022-11-01T23:16:55Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 67.0,
						"y" : 2.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL29",
						"eventTypeId" : "HIT",
						"description" : "Brady Tkachuk hit Nicholas Paul"
					},
					"about" : {
						"eventIdx" : 36,
						"eventId" : 29,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "03:44",
						"periodTimeRemaining" : "16:16",
						"dateTime" : "2022-11-01T23:17:16Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 49.0,
						"y" : -41.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479458,
							"fullName" : "Nikita Zaitsev",
							"link" : "/api/v1/people/8479458"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL30",
						"eventTypeId" : "HIT",
						"description" : "Nikita Zaitsev hit Brayden Point"
					},
					"about" : {
						"eventIdx" : 37,
						"eventId" : 30,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "04:19",
						"periodTimeRemaining" : "15:41",
						"dateTime" : "2022-11-01T23:17:51Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 76.0,
						"y" : -40.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL351",
						"eventTypeId" : "SHOT",
						"description" : "Brayden Point Tip-In saved by Anton Forsberg",
						"secondaryType" : "Tip-In"
					},
					"about" : {
						"eventIdx" : 38,
						"eventId" : 351,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "04:30",
						"periodTimeRemaining" : "15:30",
						"dateTime" : "2022-11-01T23:18:02Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 23.0,
						"y" : -31.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL31",
						"eventTypeId" : "STOP",
						"description" : "Puck Frozen"
					},
					"about" : {
						"eventIdx" : 39,
						"eventId" : 31,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "04:31",
						"periodTimeRemaining" : "15:29",
						"dateTime" : "2022-11-01T23:18:02Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL58",
						"eventTypeId" : "FACEOFF",
						"description" : "Mark Kastelic faceoff won against Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 40,
						"eventId" : 58,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "04:31",
						"periodTimeRemaining" : "15:29",
						"dateTime" : "2022-11-01T23:18:50Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL32",
						"eventTypeId" : "HIT",
						"description" : "Cole Koepke hit Nick Holden"
					},
					"about" : {
						"eventIdx" : 41,
						"eventId" : 32,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "04:48",
						"periodTimeRemaining" : "15:12",
						"dateTime" : "2022-11-01T23:19:07Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 98.0,
						"y" : -10.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL161",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Haydn Fleury shot blocked shot by Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 42,
						"eventId" : 161,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "05:11",
						"periodTimeRemaining" : "14:49",
						"dateTime" : "2022-11-01T23:19:30Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 50.0,
						"y" : 17.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL59",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 43,
						"eventId" : 59,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "05:42",
						"periodTimeRemaining" : "14:18",
						"dateTime" : "2022-11-01T23:20:01Z",
						"goals" : {
							"away" : 0,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -68.0,
						"y" : 3.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "Scorer",
						"seasonTotal" : 1
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Goal",
						"eventCode" : "TBL162",
						"eventTypeId" : "GOAL",
						"description" : "Dylan Gambrell (1) Wrist Shot, assists: none",
						"secondaryType" : "Wrist Shot",
						"strength" : {
							"code" : "EVEN",
							"name" : "Even"
						},
						"emptyNet" : false
					},
					"about" : {
						"eventIdx" : 44,
						"eventId" : 162,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "05:44",
						"periodTimeRemaining" : "14:16",
						"dateTime" : "2022-11-01T23:20:03Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -80.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL60",
						"eventTypeId" : "FACEOFF",
						"description" : "Nicholas Paul faceoff won against Brady Tkachuk"
					},
					"about" : {
						"eventIdx" : 45,
						"eventId" : 60,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "05:44",
						"periodTimeRemaining" : "14:16",
						"dateTime" : "2022-11-01T23:20:49Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 0.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL163",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Travis Hamonic Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 46,
						"eventId" : 163,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:05",
						"periodTimeRemaining" : "13:55",
						"dateTime" : "2022-11-01T23:21:10Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -52.0,
						"y" : 31.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL33",
						"eventTypeId" : "HIT",
						"description" : "Brady Tkachuk hit Erik Cernak"
					},
					"about" : {
						"eventIdx" : 47,
						"eventId" : 33,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:08",
						"periodTimeRemaining" : "13:52",
						"dateTime" : "2022-11-01T23:21:13Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -93.0,
						"y" : -20.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL212",
						"eventTypeId" : "STOP",
						"description" : "TV timeout"
					},
					"about" : {
						"eventIdx" : 49,
						"eventId" : 212,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:12",
						"periodTimeRemaining" : "13:48",
						"dateTime" : "2022-11-01T23:21:17Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482116,
							"fullName" : "Tim Stützle",
							"link" : "/api/v1/people/8482116"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL301",
						"eventTypeId" : "PENALTY",
						"description" : "Tim Stützle Charging against Ian Cole",
						"secondaryType" : "Charging",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 48,
						"eventId" : 301,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:12",
						"periodTimeRemaining" : "13:48",
						"dateTime" : "2022-11-01T23:21:17Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -55.0,
						"y" : 40.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL62",
						"eventTypeId" : "FACEOFF",
						"description" : "Tyler Motte faceoff won against Brayden Point"
					},
					"about" : {
						"eventIdx" : 50,
						"eventId" : 62,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:12",
						"periodTimeRemaining" : "13:48",
						"dateTime" : "2022-11-01T23:23:42Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL34",
						"eventTypeId" : "STOP",
						"description" : "Offside"
					},
					"about" : {
						"eventIdx" : 51,
						"eventId" : 34,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:31",
						"periodTimeRemaining" : "13:29",
						"dateTime" : "2022-11-01T23:24:00Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL63",
						"eventTypeId" : "FACEOFF",
						"description" : "Tyler Motte faceoff won against Brayden Point"
					},
					"about" : {
						"eventIdx" : 52,
						"eventId" : 63,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "06:31",
						"periodTimeRemaining" : "13:29",
						"dateTime" : "2022-11-01T23:24:25Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 20.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL164",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 53,
						"eventId" : 164,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "07:00",
						"periodTimeRemaining" : "13:00",
						"dateTime" : "2022-11-01T23:24:54Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 48.0,
						"y" : 26.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL35",
						"eventTypeId" : "HIT",
						"description" : "Brayden Point hit Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 54,
						"eventId" : 35,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "07:22",
						"periodTimeRemaining" : "12:38",
						"dateTime" : "2022-11-01T23:25:16Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 92.0,
						"y" : -30.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8470621,
							"fullName" : "Corey Perry",
							"link" : "/api/v1/people/8470621"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL36",
						"eventTypeId" : "HIT",
						"description" : "Corey Perry hit Jake Sanderson"
					},
					"about" : {
						"eventIdx" : 55,
						"eventId" : 36,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "07:32",
						"periodTimeRemaining" : "12:28",
						"dateTime" : "2022-11-01T23:25:26Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 90.0,
						"y" : -30.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL64",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Parker Kelly"
					},
					"about" : {
						"eventIdx" : 56,
						"eventId" : 64,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "07:59",
						"periodTimeRemaining" : "12:01",
						"dateTime" : "2022-11-01T23:25:53Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 52.0,
						"y" : 12.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL165",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Ross Colton Wide of Net Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 57,
						"eventId" : 165,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "08:05",
						"periodTimeRemaining" : "11:55",
						"dateTime" : "2022-11-01T23:25:59Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 74.0,
						"y" : -15.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL37",
						"eventTypeId" : "HIT",
						"description" : "Nicholas Paul hit Nick Holden"
					},
					"about" : {
						"eventIdx" : 58,
						"eventId" : 37,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "08:31",
						"periodTimeRemaining" : "11:29",
						"dateTime" : "2022-11-01T23:26:25Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 96.0,
						"y" : 23.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL166",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Drake Batherson Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 59,
						"eventId" : 166,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "09:07",
						"periodTimeRemaining" : "10:53",
						"dateTime" : "2022-11-01T23:27:01Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -76.0,
						"y" : 35.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL358",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Brandon Hagel shot blocked shot by Erik Brannstrom"
					},
					"about" : {
						"eventIdx" : 60,
						"eventId" : 358,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "09:41",
						"periodTimeRemaining" : "10:19",
						"dateTime" : "2022-11-01T23:27:35Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 77.0,
						"y" : 5.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL65",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Nikita Kucherov"
					},
					"about" : {
						"eventIdx" : 61,
						"eventId" : 65,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "09:42",
						"periodTimeRemaining" : "10:18",
						"dateTime" : "2022-11-01T23:27:36Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 52.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL168",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Jake Sanderson shot blocked shot by Nikita Kucherov"
					},
					"about" : {
						"eventIdx" : 62,
						"eventId" : 168,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "09:55",
						"periodTimeRemaining" : "10:05",
						"dateTime" : "2022-11-01T23:27:49Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -80.0,
						"y" : 5.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL38",
						"eventTypeId" : "HIT",
						"description" : "Shane Pinto hit Erik Cernak"
					},
					"about" : {
						"eventIdx" : 63,
						"eventId" : 38,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "10:18",
						"periodTimeRemaining" : "09:42",
						"dateTime" : "2022-11-01T23:28:12Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -94.0,
						"y" : 14.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL169",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Cole Koepke shot blocked shot by Nick Holden"
					},
					"about" : {
						"eventIdx" : 64,
						"eventId" : 169,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "10:27",
						"periodTimeRemaining" : "09:33",
						"dateTime" : "2022-11-01T23:28:21Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 78.0,
						"y" : 2.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL360",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Tyler Motte shot blocked shot by Ross Colton"
					},
					"about" : {
						"eventIdx" : 65,
						"eventId" : 360,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "10:42",
						"periodTimeRemaining" : "09:18",
						"dateTime" : "2022-11-01T23:28:36Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -78.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL170",
						"eventTypeId" : "SHOT",
						"description" : "Dylan Gambrell Wrist Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 66,
						"eventId" : 170,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "11:42",
						"periodTimeRemaining" : "08:18",
						"dateTime" : "2022-11-01T23:29:36Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 35.0,
						"y" : 6.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL171",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Cal Foote shot blocked shot by Jake Sanderson"
					},
					"about" : {
						"eventIdx" : 67,
						"eventId" : 171,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:03",
						"periodTimeRemaining" : "07:57",
						"dateTime" : "2022-11-01T23:29:57Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 61.0,
						"y" : -10.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL172",
						"eventTypeId" : "SHOT",
						"description" : "Mikhail Sergachev Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 68,
						"eventId" : 172,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:13",
						"periodTimeRemaining" : "07:47",
						"dateTime" : "2022-11-01T23:30:07Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 41.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL40",
						"eventTypeId" : "STOP",
						"description" : "Goalie Stopped"
					},
					"about" : {
						"eventIdx" : 70,
						"eventId" : 40,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:42",
						"periodTimeRemaining" : "07:18",
						"dateTime" : "2022-11-01T23:30:35Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL173",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 69,
						"eventId" : 173,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:42",
						"periodTimeRemaining" : "07:18",
						"dateTime" : "2022-11-01T23:30:35Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 77.0,
						"y" : 18.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL67",
						"eventTypeId" : "FACEOFF",
						"description" : "Vladislav Namestnikov faceoff won against Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 71,
						"eventId" : 67,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:42",
						"periodTimeRemaining" : "07:18",
						"dateTime" : "2022-11-01T23:33:08Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL41",
						"eventTypeId" : "STOP",
						"description" : "Goalie Stopped"
					},
					"about" : {
						"eventIdx" : 73,
						"eventId" : 41,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:46",
						"periodTimeRemaining" : "07:14",
						"dateTime" : "2022-11-01T23:33:12Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL174",
						"eventTypeId" : "SHOT",
						"description" : "Haydn Fleury Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 72,
						"eventId" : 174,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:46",
						"periodTimeRemaining" : "07:14",
						"dateTime" : "2022-11-01T23:33:12Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 37.0,
						"y" : 18.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL68",
						"eventTypeId" : "FACEOFF",
						"description" : "Vladislav Namestnikov faceoff won against Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 74,
						"eventId" : 68,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:46",
						"periodTimeRemaining" : "07:14",
						"dateTime" : "2022-11-01T23:33:27Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL175",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Haydn Fleury shot blocked shot by Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 75,
						"eventId" : 175,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "12:48",
						"periodTimeRemaining" : "07:12",
						"dateTime" : "2022-11-01T23:33:29Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 47.0,
						"y" : 16.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL42",
						"eventTypeId" : "HIT",
						"description" : "Nick Holden hit Cole Koepke"
					},
					"about" : {
						"eventIdx" : 76,
						"eventId" : 42,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "13:08",
						"periodTimeRemaining" : "06:52",
						"dateTime" : "2022-11-01T23:33:49Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 98.0,
						"y" : 4.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL44",
						"eventTypeId" : "HIT",
						"description" : "Haydn Fleury hit Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 77,
						"eventId" : 44,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "13:18",
						"periodTimeRemaining" : "06:42",
						"dateTime" : "2022-11-01T23:33:59Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 95.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL176",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Cal Foote shot blocked shot by Tyler Motte"
					},
					"about" : {
						"eventIdx" : 78,
						"eventId" : 176,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "13:19",
						"periodTimeRemaining" : "06:41",
						"dateTime" : "2022-11-01T23:34:00Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 50.0,
						"y" : 21.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL69",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Shane Pinto"
					},
					"about" : {
						"eventIdx" : 79,
						"eventId" : 69,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "13:45",
						"periodTimeRemaining" : "06:15",
						"dateTime" : "2022-11-01T23:34:26Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -11.0,
						"y" : 8.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL70",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Shane Pinto"
					},
					"about" : {
						"eventIdx" : 80,
						"eventId" : 70,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "13:53",
						"periodTimeRemaining" : "06:07",
						"dateTime" : "2022-11-01T23:34:34Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -32.0,
						"y" : -27.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL71",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Erik Brannstrom"
					},
					"about" : {
						"eventIdx" : 81,
						"eventId" : 71,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "14:07",
						"periodTimeRemaining" : "05:53",
						"dateTime" : "2022-11-01T23:34:48Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 98.0,
						"y" : 1.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8470621,
							"fullName" : "Corey Perry",
							"link" : "/api/v1/people/8470621"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL47",
						"eventTypeId" : "HIT",
						"description" : "Parker Kelly hit Corey Perry"
					},
					"about" : {
						"eventIdx" : 82,
						"eventId" : 47,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "14:37",
						"periodTimeRemaining" : "05:23",
						"dateTime" : "2022-11-01T23:35:18Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -93.0,
						"y" : 29.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL354",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Jake Sanderson shot blocked shot by Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 83,
						"eventId" : 354,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "14:56",
						"periodTimeRemaining" : "05:04",
						"dateTime" : "2022-11-01T23:35:37Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -78.0,
						"y" : -12.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL48",
						"eventTypeId" : "STOP",
						"description" : "Offside"
					},
					"about" : {
						"eventIdx" : 84,
						"eventId" : 48,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "15:40",
						"periodTimeRemaining" : "04:20",
						"dateTime" : "2022-11-01T23:36:21Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL75",
						"eventTypeId" : "FACEOFF",
						"description" : "Mark Kastelic faceoff won against Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 85,
						"eventId" : 75,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "15:40",
						"periodTimeRemaining" : "04:20",
						"dateTime" : "2022-11-01T23:38:33Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 20.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL363",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Vladislav Namestnikov shot blocked shot by Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 86,
						"eventId" : 363,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:09",
						"periodTimeRemaining" : "03:51",
						"dateTime" : "2022-11-01T23:39:02Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 77.0,
						"y" : -5.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL177",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Haydn Fleury shot blocked shot by Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 87,
						"eventId" : 177,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:13",
						"periodTimeRemaining" : "03:47",
						"dateTime" : "2022-11-01T23:39:06Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 76.0,
						"y" : 6.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL49",
						"eventTypeId" : "HIT",
						"description" : "Thomas Chabot hit Alex Killorn"
					},
					"about" : {
						"eventIdx" : 88,
						"eventId" : 49,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:51",
						"periodTimeRemaining" : "03:09",
						"dateTime" : "2022-11-01T23:39:44Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 76.0,
						"y" : -36.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL50",
						"eventTypeId" : "STOP",
						"description" : "Referee or Linesman"
					},
					"about" : {
						"eventIdx" : 89,
						"eventId" : 50,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:59",
						"periodTimeRemaining" : "03:01",
						"dateTime" : "2022-11-01T23:39:53Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL76",
						"eventTypeId" : "FACEOFF",
						"description" : "Brayden Point faceoff won against Shane Pinto"
					},
					"about" : {
						"eventIdx" : 91,
						"eventId" : 76,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:59",
						"periodTimeRemaining" : "03:01",
						"dateTime" : "2022-11-01T23:41:02Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 20.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL364",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Travis Hamonic shot blocked shot by Brandon Hagel"
					},
					"about" : {
						"eventIdx" : 90,
						"eventId" : 364,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "16:59",
						"periodTimeRemaining" : "03:01",
						"dateTime" : "2022-11-01T23:41:02Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -68.0,
						"y" : 20.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL201",
						"eventTypeId" : "HIT",
						"description" : "Brandon Hagel hit Erik Brannstrom"
					},
					"about" : {
						"eventIdx" : 92,
						"eventId" : 201,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "17:09",
						"periodTimeRemaining" : "02:51",
						"dateTime" : "2022-11-01T23:41:12Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 96.0,
						"y" : 16.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL202",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Nikita Kucherov"
					},
					"about" : {
						"eventIdx" : 93,
						"eventId" : 202,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "17:18",
						"periodTimeRemaining" : "02:42",
						"dateTime" : "2022-11-01T23:41:21Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 90.0,
						"y" : -23.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482116,
							"fullName" : "Tim Stützle",
							"link" : "/api/v1/people/8482116"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL178",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Tim Stützle Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 94,
						"eventId" : 178,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "17:49",
						"periodTimeRemaining" : "02:11",
						"dateTime" : "2022-11-01T23:41:52Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -71.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL39",
						"eventTypeId" : "HIT",
						"description" : "Mikhail Sergachev hit Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 96,
						"eventId" : 39,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "17:54",
						"periodTimeRemaining" : "02:06",
						"dateTime" : "2022-11-01T23:41:57Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -97.0,
						"y" : -5.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL365",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Travis Hamonic shot blocked shot by Brandon Hagel"
					},
					"about" : {
						"eventIdx" : 95,
						"eventId" : 365,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "17:54",
						"periodTimeRemaining" : "02:06",
						"dateTime" : "2022-11-01T23:41:57Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -64.0,
						"y" : 25.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL203",
						"eventTypeId" : "SHOT",
						"description" : "Drake Batherson Snap Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Snap Shot"
					},
					"about" : {
						"eventIdx" : 97,
						"eventId" : 203,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "18:03",
						"periodTimeRemaining" : "01:57",
						"dateTime" : "2022-11-01T23:42:06Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -54.0,
						"y" : 3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL179",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Drake Batherson Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 98,
						"eventId" : 179,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "18:06",
						"periodTimeRemaining" : "01:54",
						"dateTime" : "2022-11-01T23:42:09Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -76.0,
						"y" : 5.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL366",
						"eventTypeId" : "SHOT",
						"description" : "Ian Cole Snap Shot saved by Anton Forsberg",
						"secondaryType" : "Snap Shot"
					},
					"about" : {
						"eventIdx" : 99,
						"eventId" : 366,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "18:38",
						"periodTimeRemaining" : "01:22",
						"dateTime" : "2022-11-01T23:42:41Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -31.0,
						"y" : 40.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL204",
						"eventTypeId" : "HIT",
						"description" : "Nicholas Paul hit Jake Sanderson"
					},
					"about" : {
						"eventIdx" : 100,
						"eventId" : 204,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "18:47",
						"periodTimeRemaining" : "01:13",
						"dateTime" : "2022-11-01T23:42:50Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 87.0,
						"y" : -33.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL205",
						"eventTypeId" : "HIT",
						"description" : "Cal Foote hit Parker Kelly"
					},
					"about" : {
						"eventIdx" : 101,
						"eventId" : 205,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "18:58",
						"periodTimeRemaining" : "01:02",
						"dateTime" : "2022-11-01T23:43:01Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -65.0,
						"y" : -7.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL575",
						"eventTypeId" : "PENALTY",
						"description" : "Parker Kelly Interference against Cal Foote",
						"secondaryType" : "Interference",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 102,
						"eventId" : 575,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "19:27",
						"periodTimeRemaining" : "00:33",
						"dateTime" : "2022-11-01T23:43:30Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -80.0,
						"y" : -8.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL80",
						"eventTypeId" : "FACEOFF",
						"description" : "Steven Stamkos faceoff won against Dylan Gambrell"
					},
					"about" : {
						"eventIdx" : 103,
						"eventId" : 80,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "19:27",
						"periodTimeRemaining" : "00:33",
						"dateTime" : "2022-11-01T23:44:16Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL180",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 104,
						"eventId" : 180,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "19:39",
						"periodTimeRemaining" : "00:21",
						"dateTime" : "2022-11-01T23:44:28Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 71.0,
						"y" : 20.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL181",
						"eventTypeId" : "SHOT",
						"description" : "Vladislav Namestnikov Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 105,
						"eventId" : 181,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "19:42",
						"periodTimeRemaining" : "00:18",
						"dateTime" : "2022-11-01T23:44:31Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 80.0,
						"y" : 2.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Period End",
						"eventCode" : "TBL209",
						"eventTypeId" : "PERIOD_END",
						"description" : "End of 1st Period"
					},
					"about" : {
						"eventIdx" : 106,
						"eventId" : 209,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "20:00",
						"periodTimeRemaining" : "00:00",
						"dateTime" : "2022-11-01T23:44:49Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"result" : {
						"event" : "Period Official",
						"eventCode" : "TBL213",
						"eventTypeId" : "PERIOD_OFFICIAL",
						"description" : "Period Official"
					},
					"about" : {
						"eventIdx" : 107,
						"eventId" : 213,
						"period" : 1,
						"periodType" : "REGULAR",
						"ordinalNum" : "1st",
						"periodTime" : "20:00",
						"periodTimeRemaining" : "00:00",
						"dateTime" : "2022-11-02T00:00:43Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"result" : {
						"event" : "Period Ready",
						"eventCode" : "TBL214",
						"eventTypeId" : "PERIOD_READY",
						"description" : "Period Ready"
					},
					"about" : {
						"eventIdx" : 108,
						"eventId" : 214,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-02T00:00:47Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"result" : {
						"event" : "Period Start",
						"eventCode" : "TBL81",
						"eventTypeId" : "PERIOD_START",
						"description" : "Period Start"
					},
					"about" : {
						"eventIdx" : 109,
						"eventId" : 81,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-02T00:03:47Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL82",
						"eventTypeId" : "FACEOFF",
						"description" : "Steven Stamkos faceoff won against Tyler Motte"
					},
					"about" : {
						"eventIdx" : 110,
						"eventId" : 82,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:00",
						"periodTimeRemaining" : "20:00",
						"dateTime" : "2022-11-02T00:03:47Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 0.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL182",
						"eventTypeId" : "SHOT",
						"description" : "Nikita Kucherov Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 111,
						"eventId" : 182,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:22",
						"periodTimeRemaining" : "19:38",
						"dateTime" : "2022-11-02T00:04:09Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -54.0,
						"y" : 24.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479458,
							"fullName" : "Nikita Zaitsev",
							"link" : "/api/v1/people/8479458"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL183",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Vladislav Namestnikov shot blocked shot by Nikita Zaitsev"
					},
					"about" : {
						"eventIdx" : 112,
						"eventId" : 183,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:46",
						"periodTimeRemaining" : "19:14",
						"dateTime" : "2022-11-02T00:04:33Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -78.0,
						"y" : -3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479458,
							"fullName" : "Nikita Zaitsev",
							"link" : "/api/v1/people/8479458"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL184",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Mikhail Sergachev shot blocked shot by Nikita Zaitsev"
					},
					"about" : {
						"eventIdx" : 113,
						"eventId" : 184,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "00:47",
						"periodTimeRemaining" : "19:13",
						"dateTime" : "2022-11-02T00:04:34Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -76.0,
						"y" : 3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL185",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 114,
						"eventId" : 185,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:04",
						"periodTimeRemaining" : "18:56",
						"dateTime" : "2022-11-02T00:04:50Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -63.0,
						"y" : -17.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL215",
						"eventTypeId" : "STOP",
						"description" : "Goalie Stopped"
					},
					"about" : {
						"eventIdx" : 115,
						"eventId" : 215,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:04",
						"periodTimeRemaining" : "18:56",
						"dateTime" : "2022-11-02T00:04:50Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL83",
						"eventTypeId" : "FACEOFF",
						"description" : "Mark Kastelic faceoff won against Ross Colton"
					},
					"about" : {
						"eventIdx" : 116,
						"eventId" : 83,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:04",
						"periodTimeRemaining" : "18:56",
						"dateTime" : "2022-11-02T00:05:20Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8470621,
							"fullName" : "Corey Perry",
							"link" : "/api/v1/people/8470621"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL216",
						"eventTypeId" : "HIT",
						"description" : "Corey Perry hit Jake Sanderson"
					},
					"about" : {
						"eventIdx" : 117,
						"eventId" : 216,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:09",
						"periodTimeRemaining" : "18:51",
						"dateTime" : "2022-11-02T00:05:25Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -96.0,
						"y" : 9.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL84",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 118,
						"eventId" : 84,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:15",
						"periodTimeRemaining" : "18:45",
						"dateTime" : "2022-11-02T00:05:31Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -31.0,
						"y" : 39.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL217",
						"eventTypeId" : "HIT",
						"description" : "Ross Colton hit Travis Hamonic"
					},
					"about" : {
						"eventIdx" : 119,
						"eventId" : 217,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:35",
						"periodTimeRemaining" : "18:25",
						"dateTime" : "2022-11-02T00:05:51Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -99.0,
						"y" : -4.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL590",
						"eventTypeId" : "PENALTY",
						"description" : "Jake Sanderson High-sticking against Ross Colton",
						"secondaryType" : "High-sticking",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 120,
						"eventId" : 590,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:42",
						"periodTimeRemaining" : "18:18",
						"dateTime" : "2022-11-02T00:05:59Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -94.0,
						"y" : 21.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL85",
						"eventTypeId" : "FACEOFF",
						"description" : "Tyler Motte faceoff won against Steven Stamkos"
					},
					"about" : {
						"eventIdx" : 121,
						"eventId" : 85,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "01:42",
						"periodTimeRemaining" : "18:18",
						"dateTime" : "2022-11-02T00:06:55Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL86",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Tyler Motte"
					},
					"about" : {
						"eventIdx" : 122,
						"eventId" : 86,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:08",
						"periodTimeRemaining" : "17:52",
						"dateTime" : "2022-11-02T00:07:21Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -34.0,
						"y" : 24.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL186",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Nikita Kucherov shot blocked shot by Nick Holden"
					},
					"about" : {
						"eventIdx" : 123,
						"eventId" : 186,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:30",
						"periodTimeRemaining" : "17:30",
						"dateTime" : "2022-11-02T00:07:43Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 6.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL218",
						"eventTypeId" : "HIT",
						"description" : "Nikita Kucherov hit Nick Holden"
					},
					"about" : {
						"eventIdx" : 124,
						"eventId" : 218,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:31",
						"periodTimeRemaining" : "17:29",
						"dateTime" : "2022-11-02T00:07:44Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -73.0,
						"y" : 38.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL87",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Nick Holden"
					},
					"about" : {
						"eventIdx" : 125,
						"eventId" : 87,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:37",
						"periodTimeRemaining" : "17:23",
						"dateTime" : "2022-11-02T00:07:50Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -40.0,
						"y" : -18.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL187",
						"eventTypeId" : "SHOT",
						"description" : "Nick Holden Wrist Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 126,
						"eventId" : 187,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:44",
						"periodTimeRemaining" : "17:16",
						"dateTime" : "2022-11-02T00:07:57Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : 53.0,
						"y" : 10.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8470621,
							"fullName" : "Corey Perry",
							"link" : "/api/v1/people/8470621"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL188",
						"eventTypeId" : "SHOT",
						"description" : "Corey Perry Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 127,
						"eventId" : 188,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "02:54",
						"periodTimeRemaining" : "17:06",
						"dateTime" : "2022-11-02T00:08:07Z",
						"goals" : {
							"away" : 1,
							"home" : 0
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : -11.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Scorer",
						"seasonTotal" : 2
					}, {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Assist",
						"seasonTotal" : 3
					}, {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Assist",
						"seasonTotal" : 10
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Goal",
						"eventCode" : "TBL189",
						"eventTypeId" : "GOAL",
						"description" : "Mikhail Sergachev (2) Slap Shot, assists: Steven Stamkos (3), Nikita Kucherov (10)",
						"secondaryType" : "Slap Shot",
						"strength" : {
							"code" : "PPG",
							"name" : "Power Play"
						},
						"emptyNet" : false
					},
					"about" : {
						"eventIdx" : 128,
						"eventId" : 189,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "03:17",
						"periodTimeRemaining" : "16:43",
						"dateTime" : "2022-11-02T00:08:30Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -52.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL88",
						"eventTypeId" : "FACEOFF",
						"description" : "Shane Pinto faceoff won against Nicholas Paul"
					},
					"about" : {
						"eventIdx" : 129,
						"eventId" : 88,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "03:17",
						"periodTimeRemaining" : "16:43",
						"dateTime" : "2022-11-02T00:09:10Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 0.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477426,
							"fullName" : "Nicholas Paul",
							"link" : "/api/v1/people/8477426"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL190",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Shane Pinto shot blocked shot by Nicholas Paul"
					},
					"about" : {
						"eventIdx" : 130,
						"eventId" : 190,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "03:29",
						"periodTimeRemaining" : "16:31",
						"dateTime" : "2022-11-02T00:09:22Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : -6.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL191",
						"eventTypeId" : "SHOT",
						"description" : "Erik Cernak Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 131,
						"eventId" : 191,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "03:41",
						"periodTimeRemaining" : "16:19",
						"dateTime" : "2022-11-02T00:09:34Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -41.0,
						"y" : 29.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL219",
						"eventTypeId" : "HIT",
						"description" : "Alex DeBrincat hit Erik Cernak"
					},
					"about" : {
						"eventIdx" : 132,
						"eventId" : 219,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "03:49",
						"periodTimeRemaining" : "16:11",
						"dateTime" : "2022-11-02T00:09:42Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -72.0,
						"y" : 38.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL220",
						"eventTypeId" : "HIT",
						"description" : "Alex DeBrincat hit Ian Cole"
					},
					"about" : {
						"eventIdx" : 133,
						"eventId" : 220,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "04:03",
						"periodTimeRemaining" : "15:57",
						"dateTime" : "2022-11-02T00:09:56Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 96.0,
						"y" : 19.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL89",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Cole Koepke"
					},
					"about" : {
						"eventIdx" : 134,
						"eventId" : 89,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "04:19",
						"periodTimeRemaining" : "15:41",
						"dateTime" : "2022-11-02T00:10:12Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 33.0,
						"y" : -32.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL654",
						"eventTypeId" : "PENALTY",
						"description" : "Mikhail Sergachev Cross-checking against Thomas Chabot",
						"secondaryType" : "Cross-checking",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 135,
						"eventId" : 654,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "04:25",
						"periodTimeRemaining" : "15:35",
						"dateTime" : "2022-11-02T00:10:19Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -67.0,
						"y" : 23.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL90",
						"eventTypeId" : "FACEOFF",
						"description" : "Brady Tkachuk faceoff won against Steven Stamkos"
					},
					"about" : {
						"eventIdx" : 136,
						"eventId" : 90,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "04:25",
						"periodTimeRemaining" : "15:35",
						"dateTime" : "2022-11-02T00:11:08Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL192",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Alex DeBrincat shot blocked shot by Erik Cernak"
					},
					"about" : {
						"eventIdx" : 137,
						"eventId" : 192,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "04:36",
						"periodTimeRemaining" : "15:24",
						"dateTime" : "2022-11-02T00:11:19Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 75.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL91",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Alex DeBrincat"
					},
					"about" : {
						"eventIdx" : 138,
						"eventId" : 91,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "05:02",
						"periodTimeRemaining" : "14:58",
						"dateTime" : "2022-11-02T00:11:45Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 18.0,
						"y" : 30.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL193",
						"eventTypeId" : "SHOT",
						"description" : "Drake Batherson Wrist Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 139,
						"eventId" : 193,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "05:16",
						"periodTimeRemaining" : "14:44",
						"dateTime" : "2022-11-02T00:11:59Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 59.0,
						"y" : -3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL194",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Jake Sanderson Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 140,
						"eventId" : 194,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "05:55",
						"periodTimeRemaining" : "14:05",
						"dateTime" : "2022-11-02T00:12:38Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 36.0,
						"y" : 32.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8481596,
							"fullName" : "Shane Pinto",
							"link" : "/api/v1/people/8481596"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL195",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Shane Pinto shot blocked shot by Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 141,
						"eventId" : 195,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "06:11",
						"periodTimeRemaining" : "13:49",
						"dateTime" : "2022-11-02T00:12:54Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 71.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL92",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 142,
						"eventId" : 92,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "06:23",
						"periodTimeRemaining" : "13:37",
						"dateTime" : "2022-11-02T00:13:06Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -97.0,
						"y" : 6.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL196",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Mathieu Joseph Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 144,
						"eventId" : 196,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "06:56",
						"periodTimeRemaining" : "13:04",
						"dateTime" : "2022-11-02T00:13:39Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 68.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL221",
						"eventTypeId" : "HIT",
						"description" : "Cole Koepke hit Mathieu Joseph"
					},
					"about" : {
						"eventIdx" : 143,
						"eventId" : 221,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "06:56",
						"periodTimeRemaining" : "13:04",
						"dateTime" : "2022-11-02T00:13:39Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 88.0,
						"y" : -33.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL367",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Tyler Motte shot blocked shot by Ian Cole"
					},
					"about" : {
						"eventIdx" : 145,
						"eventId" : 367,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "06:58",
						"periodTimeRemaining" : "13:02",
						"dateTime" : "2022-11-02T00:13:41Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 67.0,
						"y" : 12.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL222",
						"eventTypeId" : "HIT",
						"description" : "Tyler Motte hit Ian Cole"
					},
					"about" : {
						"eventIdx" : 146,
						"eventId" : 222,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:08",
						"periodTimeRemaining" : "12:52",
						"dateTime" : "2022-11-02T00:13:51Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 98.0,
						"y" : -19.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL197",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Cole Koepke Wide of Net Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 147,
						"eventId" : 197,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:25",
						"periodTimeRemaining" : "12:35",
						"dateTime" : "2022-11-02T00:14:08Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -84.0,
						"y" : 15.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL234",
						"eventTypeId" : "PENALTY",
						"description" : "Dylan Gambrell Minor against Erik Cernak",
						"secondaryType" : "Minor",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 149,
						"eventId" : 234,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:27",
						"periodTimeRemaining" : "12:33",
						"dateTime" : "2022-11-02T00:14:10Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 61.0,
						"y" : 27.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL237",
						"eventTypeId" : "STOP",
						"description" : "Referee or Linesman"
					},
					"about" : {
						"eventIdx" : 150,
						"eventId" : 237,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:27",
						"periodTimeRemaining" : "12:33",
						"dateTime" : "2022-11-02T00:14:10Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479580,
							"fullName" : "Dylan Gambrell",
							"link" : "/api/v1/people/8479580"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "DrewBy"
					}, {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "ServedBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL680",
						"eventTypeId" : "PENALTY",
						"description" : "Dylan Gambrell Attempt to injure against Erik Cernak served by Parker Kelly",
						"secondaryType" : "Missing key [PD_154]",
						"penaltySeverity" : "Major",
						"penaltyMinutes" : 5
					},
					"about" : {
						"eventIdx" : 148,
						"eventId" : 680,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:27",
						"periodTimeRemaining" : "12:33",
						"dateTime" : "2022-11-02T00:14:10Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 64.0,
						"y" : 36.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL93",
						"eventTypeId" : "FACEOFF",
						"description" : "Mark Kastelic faceoff won against Steven Stamkos"
					},
					"about" : {
						"eventIdx" : 151,
						"eventId" : 93,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "07:27",
						"periodTimeRemaining" : "12:33",
						"dateTime" : "2022-11-02T00:19:31Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL198",
						"eventTypeId" : "SHOT",
						"description" : "Mikhail Sergachev Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 152,
						"eventId" : 198,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "08:07",
						"periodTimeRemaining" : "11:53",
						"dateTime" : "2022-11-02T00:20:11Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -33.0,
						"y" : -12.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL199",
						"eventTypeId" : "SHOT",
						"description" : "Vladislav Namestnikov Tip-In saved by Anton Forsberg",
						"secondaryType" : "Tip-In"
					},
					"about" : {
						"eventIdx" : 153,
						"eventId" : 199,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "08:09",
						"periodTimeRemaining" : "11:51",
						"dateTime" : "2022-11-02T00:20:13Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -78.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL94",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Steven Stamkos"
					},
					"about" : {
						"eventIdx" : 154,
						"eventId" : 94,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "08:30",
						"periodTimeRemaining" : "11:30",
						"dateTime" : "2022-11-02T00:20:34Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : -31.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8475766,
							"fullName" : "Austin Watson",
							"link" : "/api/v1/people/8475766"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL200",
						"eventTypeId" : "HIT",
						"description" : "Austin Watson hit Alex Killorn"
					},
					"about" : {
						"eventIdx" : 155,
						"eventId" : 200,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "09:01",
						"periodTimeRemaining" : "10:59",
						"dateTime" : "2022-11-02T00:21:05Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -38.0,
						"y" : 37.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL451",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Ross Colton Wide of Net Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 156,
						"eventId" : 451,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "09:28",
						"periodTimeRemaining" : "10:32",
						"dateTime" : "2022-11-02T00:21:32Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -81.0,
						"y" : -3.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL452",
						"eventTypeId" : "SHOT",
						"description" : "Brayden Point Tip-In saved by Anton Forsberg",
						"secondaryType" : "Tip-In"
					},
					"about" : {
						"eventIdx" : 157,
						"eventId" : 452,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:31",
						"periodTimeRemaining" : "09:29",
						"dateTime" : "2022-11-02T00:22:35Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -75.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL453",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 158,
						"eventId" : 453,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:40",
						"periodTimeRemaining" : "09:20",
						"dateTime" : "2022-11-02T00:22:44Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -73.0,
						"y" : -16.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL454",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Nikita Kucherov Wide of Net Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 159,
						"eventId" : 454,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:51",
						"periodTimeRemaining" : "09:09",
						"dateTime" : "2022-11-02T00:22:55Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -73.0,
						"y" : 8.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "PenaltyOn"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL712",
						"eventTypeId" : "PENALTY",
						"description" : "Too many men/ice served by Brady Tkachuk",
						"secondaryType" : "Too many men on the ice",
						"penaltySeverity" : "Bench Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 160,
						"eventId" : 712,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:53",
						"periodTimeRemaining" : "09:07",
						"dateTime" : "2022-11-02T00:22:57Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 13.0,
						"y" : 41.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL95",
						"eventTypeId" : "FACEOFF",
						"description" : "Brayden Point faceoff won against Tyler Motte"
					},
					"about" : {
						"eventIdx" : 161,
						"eventId" : 95,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:53",
						"periodTimeRemaining" : "09:07",
						"dateTime" : "2022-11-02T00:23:52Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL455",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 162,
						"eventId" : 455,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "10:57",
						"periodTimeRemaining" : "09:03",
						"dateTime" : "2022-11-02T00:23:56Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -60.0,
						"y" : -16.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474207,
							"fullName" : "Nick Holden",
							"link" : "/api/v1/people/8474207"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8476453,
							"fullName" : "Nikita Kucherov",
							"link" : "/api/v1/people/8476453"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL224",
						"eventTypeId" : "HIT",
						"description" : "Nick Holden hit Nikita Kucherov"
					},
					"about" : {
						"eventIdx" : 163,
						"eventId" : 224,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "11:22",
						"periodTimeRemaining" : "08:38",
						"dateTime" : "2022-11-02T00:24:21Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -87.0,
						"y" : 37.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479458,
							"fullName" : "Nikita Zaitsev",
							"link" : "/api/v1/people/8479458"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL456",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Alex Killorn shot blocked shot by Nikita Zaitsev"
					},
					"about" : {
						"eventIdx" : 164,
						"eventId" : 456,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "11:44",
						"periodTimeRemaining" : "08:16",
						"dateTime" : "2022-11-02T00:24:43Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -79.0,
						"y" : 2.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL457",
						"eventTypeId" : "SHOT",
						"description" : "Brayden Point Wrist Shot saved by Anton Forsberg",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 165,
						"eventId" : 457,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "11:51",
						"periodTimeRemaining" : "08:09",
						"dateTime" : "2022-11-02T00:24:50Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -76.0,
						"y" : -16.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL458",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Steven Stamkos Goalpost Anton Forsberg"
					},
					"about" : {
						"eventIdx" : 166,
						"eventId" : 458,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "11:51",
						"periodTimeRemaining" : "08:09",
						"dateTime" : "2022-11-02T00:24:50Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -55.0,
						"y" : -18.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477353,
							"fullName" : "Tyler Motte",
							"link" : "/api/v1/people/8477353"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL97",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Tyler Motte"
					},
					"about" : {
						"eventIdx" : 167,
						"eventId" : 97,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "11:55",
						"periodTimeRemaining" : "08:05",
						"dateTime" : "2022-11-02T00:24:54Z",
						"goals" : {
							"away" : 1,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -33.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Scorer",
						"seasonTotal" : 1
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Goal",
						"eventCode" : "TBL459",
						"eventTypeId" : "GOAL",
						"description" : "Mathieu Joseph (1) Wrist Shot, assists: none",
						"secondaryType" : "Wrist Shot",
						"strength" : {
							"code" : "SHG",
							"name" : "Short Handed"
						},
						"emptyNet" : false
					},
					"about" : {
						"eventIdx" : 168,
						"eventId" : 459,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "12:21",
						"periodTimeRemaining" : "07:39",
						"dateTime" : "2022-11-02T00:25:20Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 73.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8482116,
							"fullName" : "Tim Stützle",
							"link" : "/api/v1/people/8482116"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL98",
						"eventTypeId" : "FACEOFF",
						"description" : "Tim Stützle faceoff won against Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 169,
						"eventId" : 98,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "12:21",
						"periodTimeRemaining" : "07:39",
						"dateTime" : "2022-11-02T00:26:02Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 0.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480448,
							"fullName" : "Parker Kelly",
							"link" : "/api/v1/people/8480448"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL460",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Mikhail Sergachev shot blocked shot by Parker Kelly"
					},
					"about" : {
						"eventIdx" : 170,
						"eventId" : 460,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "12:45",
						"periodTimeRemaining" : "07:15",
						"dateTime" : "2022-11-02T00:26:26Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -49.0,
						"y" : 3.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL99",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 171,
						"eventId" : 99,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "13:08",
						"periodTimeRemaining" : "06:52",
						"dateTime" : "2022-11-02T00:26:49Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 39.0,
						"y" : -21.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL225",
						"eventTypeId" : "HIT",
						"description" : "Brady Tkachuk hit Erik Cernak"
					},
					"about" : {
						"eventIdx" : 172,
						"eventId" : 225,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "13:18",
						"periodTimeRemaining" : "06:42",
						"dateTime" : "2022-11-02T00:26:59Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 13.0,
						"y" : 41.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL100",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 173,
						"eventId" : 100,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "13:43",
						"periodTimeRemaining" : "06:17",
						"dateTime" : "2022-11-02T00:27:24Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 5.0,
						"y" : 11.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8479337,
							"fullName" : "Alex DeBrincat",
							"link" : "/api/v1/people/8479337"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL226",
						"eventTypeId" : "HIT",
						"description" : "Erik Cernak hit Alex DeBrincat"
					},
					"about" : {
						"eventIdx" : 174,
						"eventId" : 226,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "13:49",
						"periodTimeRemaining" : "06:11",
						"dateTime" : "2022-11-02T00:27:30Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 77.0,
						"y" : 37.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL461",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Thomas Chabot shot blocked shot by Cole Koepke"
					},
					"about" : {
						"eventIdx" : 175,
						"eventId" : 461,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "13:56",
						"periodTimeRemaining" : "06:04",
						"dateTime" : "2022-11-02T00:27:37Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 52.0,
						"y" : 15.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8480246,
							"fullName" : "Nick Perbix",
							"link" : "/api/v1/people/8480246"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL368",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Nick Perbix shot blocked shot by Thomas Chabot"
					},
					"about" : {
						"eventIdx" : 176,
						"eventId" : 368,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "14:15",
						"periodTimeRemaining" : "05:45",
						"dateTime" : "2022-11-02T00:27:56Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -79.0,
						"y" : 9.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Stoppage",
						"eventCode" : "TBL227",
						"eventTypeId" : "STOP",
						"description" : "Goalie Stopped"
					},
					"about" : {
						"eventIdx" : 178,
						"eventId" : 227,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "14:51",
						"periodTimeRemaining" : "05:09",
						"dateTime" : "2022-11-02T00:28:32Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : { }
				}, {
					"players" : [ {
						"player" : {
							"id" : 8470621,
							"fullName" : "Corey Perry",
							"link" : "/api/v1/people/8470621"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL462",
						"eventTypeId" : "SHOT",
						"description" : "Corey Perry Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 177,
						"eventId" : 462,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "14:51",
						"periodTimeRemaining" : "05:09",
						"dateTime" : "2022-11-02T00:28:32Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -37.0,
						"y" : -1.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8482116,
							"fullName" : "Tim Stützle",
							"link" : "/api/v1/people/8482116"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL601",
						"eventTypeId" : "FACEOFF",
						"description" : "Vladislav Namestnikov faceoff won against Tim Stützle"
					},
					"about" : {
						"eventIdx" : 179,
						"eventId" : 601,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "14:51",
						"periodTimeRemaining" : "05:09",
						"dateTime" : "2022-11-02T00:31:31Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -20.0,
						"y" : -22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8474013,
							"fullName" : "Ian Cole",
							"link" : "/api/v1/people/8474013"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL463",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Ian Cole shot blocked shot by Erik Brannstrom"
					},
					"about" : {
						"eventIdx" : 180,
						"eventId" : 463,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:21",
						"periodTimeRemaining" : "04:39",
						"dateTime" : "2022-11-02T00:32:01Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : -2.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478416,
							"fullName" : "Erik Cernak",
							"link" : "/api/v1/people/8478416"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL464",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Brady Tkachuk shot blocked shot by Erik Cernak"
					},
					"about" : {
						"eventIdx" : 181,
						"eventId" : 464,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:24",
						"periodTimeRemaining" : "04:36",
						"dateTime" : "2022-11-02T00:32:04Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 75.0,
						"y" : 1.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8481043,
							"fullName" : "Cole Koepke",
							"link" : "/api/v1/people/8481043"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL369",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Brady Tkachuk shot blocked shot by Cole Koepke"
					},
					"about" : {
						"eventIdx" : 182,
						"eventId" : 369,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:25",
						"periodTimeRemaining" : "04:35",
						"dateTime" : "2022-11-02T00:32:05Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 68.0,
						"y" : 17.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL602",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Ross Colton"
					},
					"about" : {
						"eventIdx" : 183,
						"eventId" : 602,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:37",
						"periodTimeRemaining" : "04:23",
						"dateTime" : "2022-11-02T00:32:17Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 2.0,
						"y" : 25.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474612,
							"fullName" : "Travis Hamonic",
							"link" : "/api/v1/people/8474612"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL743",
						"eventTypeId" : "PENALTY",
						"description" : "Travis Hamonic Hooking against Vladislav Namestnikov",
						"secondaryType" : "Hooking",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 184,
						"eventId" : 743,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:44",
						"periodTimeRemaining" : "04:16",
						"dateTime" : "2022-11-02T00:32:24Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -74.0,
						"y" : -5.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478010,
							"fullName" : "Brayden Point",
							"link" : "/api/v1/people/8478010"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8480355,
							"fullName" : "Mark Kastelic",
							"link" : "/api/v1/people/8480355"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL603",
						"eventTypeId" : "FACEOFF",
						"description" : "Brayden Point faceoff won against Mark Kastelic"
					},
					"about" : {
						"eventIdx" : 185,
						"eventId" : 603,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:44",
						"periodTimeRemaining" : "04:16",
						"dateTime" : "2022-11-02T00:33:17Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8474564,
							"fullName" : "Steven Stamkos",
							"link" : "/api/v1/people/8474564"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL465",
						"eventTypeId" : "SHOT",
						"description" : "Steven Stamkos Slap Shot saved by Anton Forsberg",
						"secondaryType" : "Slap Shot"
					},
					"about" : {
						"eventIdx" : 186,
						"eventId" : 465,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "15:49",
						"periodTimeRemaining" : "04:11",
						"dateTime" : "2022-11-02T00:33:22Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -49.0,
						"y" : -12.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479410,
							"fullName" : "Mikhail Sergachev",
							"link" : "/api/v1/people/8479410"
						},
						"playerType" : "PenaltyOn"
					}, {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "DrewBy"
					} ],
					"result" : {
						"event" : "Penalty",
						"eventCode" : "TBL751",
						"eventTypeId" : "PENALTY",
						"description" : "Mikhail Sergachev Hooking against Mathieu Joseph",
						"secondaryType" : "Hooking",
						"penaltySeverity" : "Minor",
						"penaltyMinutes" : 2
					},
					"about" : {
						"eventIdx" : 187,
						"eventId" : 751,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "16:50",
						"periodTimeRemaining" : "03:10",
						"dateTime" : "2022-11-02T00:34:22Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -25.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8473512,
							"fullName" : "Claude Giroux",
							"link" : "/api/v1/people/8473512"
						},
						"playerType" : "Winner"
					}, {
						"player" : {
							"id" : 8479542,
							"fullName" : "Brandon Hagel",
							"link" : "/api/v1/people/8479542"
						},
						"playerType" : "Loser"
					} ],
					"result" : {
						"event" : "Faceoff",
						"eventCode" : "TBL604",
						"eventTypeId" : "FACEOFF",
						"description" : "Claude Giroux faceoff won against Brandon Hagel"
					},
					"about" : {
						"eventIdx" : 188,
						"eventId" : 604,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "16:50",
						"periodTimeRemaining" : "03:10",
						"dateTime" : "2022-11-02T00:35:16Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 69.0,
						"y" : 22.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL466",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Thomas Chabot Wide of Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 189,
						"eventId" : 466,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "17:00",
						"periodTimeRemaining" : "03:00",
						"dateTime" : "2022-11-02T00:35:26Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 67.0,
						"y" : -25.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477938,
							"fullName" : "Haydn Fleury",
							"link" : "/api/v1/people/8477938"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8473512,
							"fullName" : "Claude Giroux",
							"link" : "/api/v1/people/8473512"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL467",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Claude Giroux shot blocked shot by Haydn Fleury"
					},
					"about" : {
						"eventIdx" : 190,
						"eventId" : 467,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "17:35",
						"periodTimeRemaining" : "02:25",
						"dateTime" : "2022-11-02T00:36:01Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 63.0,
						"y" : 0.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL468",
						"eventTypeId" : "SHOT",
						"description" : "Erik Brannstrom Wrist Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 191,
						"eventId" : 468,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "17:40",
						"periodTimeRemaining" : "02:20",
						"dateTime" : "2022-11-02T00:36:06Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 67.0,
						"y" : 14.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8477930,
							"fullName" : "Pierre-Edouard Bellemare",
							"link" : "/api/v1/people/8477930"
						},
						"playerType" : "Blocker"
					}, {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					} ],
					"result" : {
						"event" : "Blocked Shot",
						"eventCode" : "TBL469",
						"eventTypeId" : "BLOCKED_SHOT",
						"description" : "Drake Batherson shot blocked shot by Pierre-Edouard Bellemare"
					},
					"about" : {
						"eventIdx" : 192,
						"eventId" : 469,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "18:25",
						"periodTimeRemaining" : "01:35",
						"dateTime" : "2022-11-02T00:36:51Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 60.0,
						"y" : 2.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL605",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Alex Killorn"
					},
					"about" : {
						"eventIdx" : 193,
						"eventId" : 605,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "18:36",
						"periodTimeRemaining" : "01:24",
						"dateTime" : "2022-11-02T00:37:02Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 61.0,
						"y" : 39.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8473986,
							"fullName" : "Alex Killorn",
							"link" : "/api/v1/people/8473986"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL606",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Alex Killorn"
					},
					"about" : {
						"eventIdx" : 194,
						"eventId" : 606,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "18:41",
						"periodTimeRemaining" : "01:19",
						"dateTime" : "2022-11-02T00:37:07Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -35.0,
						"y" : 39.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480208,
							"fullName" : "Drake Batherson",
							"link" : "/api/v1/people/8480208"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL470",
						"eventTypeId" : "SHOT",
						"description" : "Drake Batherson Wrist Shot saved by Andrei Vasilevskiy",
						"secondaryType" : "Wrist Shot"
					},
					"about" : {
						"eventIdx" : 195,
						"eventId" : 470,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "18:47",
						"periodTimeRemaining" : "01:13",
						"dateTime" : "2022-11-02T00:37:13Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 82.0,
						"y" : -2.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480801,
							"fullName" : "Brady Tkachuk",
							"link" : "/api/v1/people/8480801"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Takeaway",
						"eventCode" : "TBL228",
						"eventTypeId" : "TAKEAWAY",
						"description" : "Takeaway by Brady Tkachuk"
					},
					"about" : {
						"eventIdx" : 196,
						"eventId" : 228,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "18:50",
						"periodTimeRemaining" : "01:10",
						"dateTime" : "2022-11-02T00:37:16Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 89.0,
						"y" : -31.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479984,
							"fullName" : "Cal Foote",
							"link" : "/api/v1/people/8479984"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8482105,
							"fullName" : "Jake Sanderson",
							"link" : "/api/v1/people/8482105"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL229",
						"eventTypeId" : "HIT",
						"description" : "Cal Foote hit Jake Sanderson"
					},
					"about" : {
						"eventIdx" : 197,
						"eventId" : 229,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "19:16",
						"periodTimeRemaining" : "00:44",
						"dateTime" : "2022-11-02T00:37:42Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -86.0,
						"y" : -35.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8479525,
							"fullName" : "Ross Colton",
							"link" : "/api/v1/people/8479525"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476341,
							"fullName" : "Anton Forsberg",
							"link" : "/api/v1/people/8476341"
						},
						"playerType" : "Goalie"
					} ],
					"result" : {
						"event" : "Shot",
						"eventCode" : "TBL471",
						"eventTypeId" : "SHOT",
						"description" : "Ross Colton Snap Shot saved by Anton Forsberg",
						"secondaryType" : "Snap Shot"
					},
					"about" : {
						"eventIdx" : 198,
						"eventId" : 471,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "19:24",
						"periodTimeRemaining" : "00:36",
						"dateTime" : "2022-11-02T00:37:50Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -81.0,
						"y" : 23.0
					},
					"team" : {
						"id" : 14,
						"name" : "Tampa Bay Lightning",
						"link" : "/api/v1/teams/14",
						"triCode" : "TBL"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478469,
							"fullName" : "Thomas Chabot",
							"link" : "/api/v1/people/8478469"
						},
						"playerType" : "Hitter"
					}, {
						"player" : {
							"id" : 8476480,
							"fullName" : "Vladislav Namestnikov",
							"link" : "/api/v1/people/8476480"
						},
						"playerType" : "Hittee"
					} ],
					"result" : {
						"event" : "Hit",
						"eventCode" : "TBL230",
						"eventTypeId" : "HIT",
						"description" : "Thomas Chabot hit Vladislav Namestnikov"
					},
					"about" : {
						"eventIdx" : 199,
						"eventId" : 230,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "19:28",
						"periodTimeRemaining" : "00:32",
						"dateTime" : "2022-11-02T00:37:54Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : -92.0,
						"y" : 28.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8478472,
							"fullName" : "Mathieu Joseph",
							"link" : "/api/v1/people/8478472"
						},
						"playerType" : "Shooter"
					}, {
						"player" : {
							"id" : 8476883,
							"fullName" : "Andrei Vasilevskiy",
							"link" : "/api/v1/people/8476883"
						},
						"playerType" : "Unknown"
					} ],
					"result" : {
						"event" : "Missed Shot",
						"eventCode" : "TBL472",
						"eventTypeId" : "MISSED_SHOT",
						"description" : "Mathieu Joseph Over Net Andrei Vasilevskiy"
					},
					"about" : {
						"eventIdx" : 200,
						"eventId" : 472,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "19:54",
						"periodTimeRemaining" : "00:06",
						"dateTime" : "2022-11-02T00:38:20Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 44.0,
						"y" : -18.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"players" : [ {
						"player" : {
							"id" : 8480073,
							"fullName" : "Erik Brannstrom",
							"link" : "/api/v1/people/8480073"
						},
						"playerType" : "PlayerID"
					} ],
					"result" : {
						"event" : "Giveaway",
						"eventCode" : "TBL607",
						"eventTypeId" : "GIVEAWAY",
						"description" : "Giveaway by Erik Brannstrom"
					},
					"about" : {
						"eventIdx" : 201,
						"eventId" : 607,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "19:55",
						"periodTimeRemaining" : "00:05",
						"dateTime" : "2022-11-02T00:38:21Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : {
						"x" : 42.0,
						"y" : -30.0
					},
					"team" : {
						"id" : 9,
						"name" : "Ottawa Senators",
						"link" : "/api/v1/teams/9",
						"triCode" : "OTT"
					}
				}, {
					"result" : {
						"event" : "Period End",
						"eventCode" : "TBL231",
						"eventTypeId" : "PERIOD_END",
						"description" : "End of 2nd Period"
					},
					"about" : {
						"eventIdx" : 202,
						"eventId" : 231,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "20:00",
						"periodTimeRemaining" : "00:00",
						"dateTime" : "2022-11-02T00:38:27Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : { }
				} ],
				"scoringPlays" : [ 44, 128, 168 ],
				"penaltyPlays" : [ 17, 18, 49, 102, 120, 135, 148, 150, 160, 184, 187 ],
				"playsByPeriod" : [ {
					"startIndex" : 0,
					"plays" : [ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107 ],
					"endIndex" : 106
				}, {
					"startIndex" : 108,
					"plays" : [ 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 127, 128, 129, 130, 131, 132, 133, 134, 135, 136, 137, 138, 139, 140, 141, 142, 143, 144, 145, 146, 147, 148, 149, 150, 151, 152, 153, 154, 155, 156, 157, 158, 159, 160, 161, 162, 163, 164, 165, 166, 167, 168, 169, 170, 171, 172, 173, 174, 175, 176, 177, 178, 179, 180, 181, 182, 183, 184, 185, 186, 187, 188, 189, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199, 200, 201, 202 ],
					"endIndex" : 202
				} ],
				"currentPlay" : {
					"result" : {
						"event" : "Period End",
						"eventCode" : "TBL231",
						"eventTypeId" : "PERIOD_END",
						"description" : "End of 2nd Period"
					},
					"about" : {
						"eventIdx" : 202,
						"eventId" : 231,
						"period" : 2,
						"periodType" : "REGULAR",
						"ordinalNum" : "2nd",
						"periodTime" : "20:00",
						"periodTimeRemaining" : "00:00",
						"dateTime" : "2022-11-02T00:38:27Z",
						"goals" : {
							"away" : 2,
							"home" : 1
						}
					},
					"coordinates" : { }
				}
			},
			"linescore" : {
				"currentPeriod" : 2,
				"currentPeriodOrdinal" : "2nd",
				"currentPeriodTimeRemaining" : "END",
				"periods" : [ {
					"periodType" : "REGULAR",
					"startTime" : "2022-11-01T23:10:00Z",
					"endTime" : "2022-11-01T23:44:49Z",
					"num" : 1,
					"ordinalNum" : "1st",
					"home" : {
						"goals" : 0,
						"shotsOnGoal" : 12,
						"rinkSide" : "left"
					},
					"away" : {
						"goals" : 1,
						"shotsOnGoal" : 3,
						"rinkSide" : "right"
					}
				}, {
					"periodType" : "REGULAR",
					"startTime" : "2022-11-02T00:03:47Z",
					"endTime" : "2022-11-02T00:38:27Z",
					"num" : 2,
					"ordinalNum" : "2nd",
					"home" : {
						"goals" : 1,
						"shotsOnGoal" : 14,
						"rinkSide" : "right"
					},
					"away" : {
						"goals" : 1,
						"shotsOnGoal" : 5,
						"rinkSide" : "left"
					}
				} ],
				"shootoutInfo" : {
					"away" : {
						"scores" : 0,
						"attempts" : 0
					},
					"home" : {
						"scores" : 0,
						"attempts" : 0
					}
				},
				"teams" : {
					"home" : {
						"team" : {
							"id" : 14,
							"name" : "Tampa Bay Lightning",
							"link" : "/api/v1/teams/14",
							"abbreviation" : "TBL",
							"triCode" : "TBL"
						},
						"goals" : 1,
						"shotsOnGoal" : 26,
						"goaliePulled" : false,
						"numSkaters" : 5,
						"powerPlay" : false
					},
					"away" : {
						"team" : {
							"id" : 9,
							"name" : "Ottawa Senators",
							"link" : "/api/v1/teams/9",
							"abbreviation" : "OTT",
							"triCode" : "OTT"
						},
						"goals" : 2,
						"shotsOnGoal" : 8,
						"goaliePulled" : false,
						"numSkaters" : 5,
						"powerPlay" : false
					}
				},
				"powerPlayStrength" : "Even",
				"hasShootout" : false,
				"intermissionInfo" : {
					"intermissionTimeRemaining" : 68,
					"intermissionTimeElapsed" : 1012,
					"inIntermission" : true
				},
				"powerPlayInfo" : {
					"situationTimeRemaining" : 0,
					"situationTimeElapsed" : 70,
					"inSituation" : false
				}
			},
			"boxscore" : {
				"teams" : {
					"away" : {
						"team" : {
							"id" : 9,
							"name" : "Ottawa Senators",
							"link" : "/api/v1/teams/9",
							"abbreviation" : "OTT",
							"triCode" : "OTT"
						},
						"teamStats" : {
							"teamSkaterStats" : {
								"goals" : 2,
								"pim" : 22,
								"shots" : 8,
								"powerPlayPercentage" : "0.0",
								"powerPlayGoals" : 0.0,
								"powerPlayOpportunities" : 2.0,
								"faceOffWinPercentage" : "53.8",
								"blocked" : 19,
								"takeaways" : 10,
								"giveaways" : 2,
								"hits" : 19
							}
						},
						"players" : {
							"ID8480073" : {
								"person" : {
									"id" : 8480073,
									"fullName" : "Erik Brannstrom",
									"link" : "/api/v1/people/8480073",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "26",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "11:00",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 1,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 1,
										"evenTimeOnIce" : "10:58",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "0:02"
									}
								}
							},
							"ID8473512" : {
								"person" : {
									"id" : 8473512,
									"fullName" : "Claude Giroux",
									"link" : "/api/v1/people/8473512",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "28",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:08",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 67.0,
										"faceOffWins" : 2,
										"faceoffTaken" : 3,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:58",
										"powerPlayTimeOnIce" : "1:10",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8473544" : {
								"person" : {
									"id" : 8473544,
									"fullName" : "Derick Brassard",
									"link" : "/api/v1/people/8473544",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "61",
								"position" : {
									"code" : "N/A",
									"name" : "Unknown",
									"type" : "Unknown",
									"abbreviation" : "N/A"
								},
								"stats" : { }
							},
							"ID8474207" : {
								"person" : {
									"id" : 8474207,
									"fullName" : "Nick Holden",
									"link" : "/api/v1/people/8474207",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "5",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "14:29",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 0,
										"evenTimeOnIce" : "6:05",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "8:24"
									}
								}
							},
							"ID8475766" : {
								"person" : {
									"id" : 8475766,
									"fullName" : "Austin Watson",
									"link" : "/api/v1/people/8475766",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "16",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "5:06",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 5,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "2:09",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "2:57"
									}
								}
							},
							"ID8474612" : {
								"person" : {
									"id" : 8474612,
									"fullName" : "Travis Hamonic",
									"link" : "/api/v1/people/8474612",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "23",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "12:25",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 2,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 1,
										"evenTimeOnIce" : "7:09",
										"powerPlayTimeOnIce" : "0:03",
										"shortHandedTimeOnIce" : "5:13"
									}
								}
							},
							"ID8479337" : {
								"person" : {
									"id" : 8479337,
									"fullName" : "Alex DeBrincat",
									"link" : "/api/v1/people/8479337",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "12",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "11:09",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 3,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 1,
										"evenTimeOnIce" : "8:17",
										"powerPlayTimeOnIce" : "2:52",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8479458" : {
								"person" : {
									"id" : 8479458,
									"fullName" : "Nikita Zaitsev",
									"link" : "/api/v1/people/8479458",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "22",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "13:23",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 3,
										"plusMinus" : 0,
										"evenTimeOnIce" : "5:15",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "8:08"
									}
								}
							},
							"ID8480448" : {
								"person" : {
									"id" : 8480448,
									"fullName" : "Parker Kelly",
									"link" : "/api/v1/people/8480448",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "45",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "5:10",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 2,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 1,
										"evenTimeOnIce" : "3:40",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "1:30"
									}
								}
							},
							"ID8480801" : {
								"person" : {
									"id" : 8480801,
									"fullName" : "Brady Tkachuk",
									"link" : "/api/v1/people/8480801",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "7",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:02",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 3,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 50.0,
										"faceOffWins" : 1,
										"faceoffTaken" : 2,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "5:53",
										"powerPlayTimeOnIce" : "2:59",
										"shortHandedTimeOnIce" : "0:10"
									}
								}
							},
							"ID8478469" : {
								"person" : {
									"id" : 8478469,
									"fullName" : "Thomas Chabot",
									"link" : "/api/v1/people/8478469",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "72",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "15:43",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 3,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 1,
										"evenTimeOnIce" : "13:17",
										"powerPlayTimeOnIce" : "2:26",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8476433" : {
								"person" : {
									"id" : 8476433,
									"fullName" : "Magnus Hellberg",
									"link" : "/api/v1/people/8476433",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "39",
								"position" : {
									"code" : "G",
									"name" : "Goalie",
									"type" : "Goalie",
									"abbreviation" : "G"
								},
								"stats" : {
									"goalieStats" : {
										"timeOnIce" : "0:00",
										"assists" : 0,
										"goals" : 0,
										"pim" : 0,
										"shots" : 0,
										"saves" : 0,
										"powerPlaySaves" : 0,
										"shortHandedSaves" : 0,
										"evenSaves" : 0,
										"shortHandedShotsAgainst" : 0,
										"evenShotsAgainst" : 0,
										"powerPlayShotsAgainst" : 0
									}
								}
							},
							"ID8480208" : {
								"person" : {
									"id" : 8480208,
									"fullName" : "Drake Batherson",
									"link" : "/api/v1/people/8480208",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "19",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:42",
										"assists" : 0,
										"goals" : 0,
										"shots" : 3,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:43",
										"powerPlayTimeOnIce" : "1:59",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8476341" : {
								"person" : {
									"id" : 8476341,
									"fullName" : "Anton Forsberg",
									"link" : "/api/v1/people/8476341",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "31",
								"position" : {
									"code" : "G",
									"name" : "Goalie",
									"type" : "Goalie",
									"abbreviation" : "G"
								},
								"stats" : {
									"goalieStats" : {
										"timeOnIce" : "40:00",
										"assists" : 0,
										"goals" : 0,
										"pim" : 0,
										"shots" : 26,
										"saves" : 25,
										"powerPlaySaves" : 13,
										"shortHandedSaves" : 0,
										"evenSaves" : 12,
										"shortHandedShotsAgainst" : 0,
										"evenShotsAgainst" : 12,
										"powerPlayShotsAgainst" : 14,
										"savePercentage" : 96.15384615384616,
										"powerPlaySavePercentage" : 92.85714285714286,
										"evenStrengthSavePercentage" : 100.0
									}
								}
							},
							"ID8477353" : {
								"person" : {
									"id" : 8477353,
									"fullName" : "Tyler Motte",
									"link" : "/api/v1/people/8477353",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "14",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "12:53",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 60.0,
										"faceOffWins" : 3,
										"faceoffTaken" : 5,
										"takeaways" : 2,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 0,
										"evenTimeOnIce" : "6:00",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "6:53"
									}
								}
							},
							"ID8478472" : {
								"person" : {
									"id" : 8478472,
									"fullName" : "Mathieu Joseph",
									"link" : "/api/v1/people/8478472",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "21",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "11:27",
										"assists" : 0,
										"goals" : 1,
										"shots" : 1,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 1,
										"shortHandedAssists" : 0,
										"blocked" : 3,
										"plusMinus" : 1,
										"evenTimeOnIce" : "6:43",
										"powerPlayTimeOnIce" : "0:04",
										"shortHandedTimeOnIce" : "4:40"
									}
								}
							},
							"ID8479580" : {
								"person" : {
									"id" : 8479580,
									"fullName" : "Dylan Gambrell",
									"link" : "/api/v1/people/8479580",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "27",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "7:47",
										"assists" : 0,
										"goals" : 1,
										"shots" : 2,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 5,
										"faceOffPct" : 50.0,
										"faceOffWins" : 1,
										"faceoffTaken" : 2,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 1,
										"evenTimeOnIce" : "4:17",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "3:30"
									}
								}
							},
							"ID8481596" : {
								"person" : {
									"id" : 8481596,
									"fullName" : "Shane Pinto",
									"link" : "/api/v1/people/8481596",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "57",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "10:28",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 50.0,
										"faceOffWins" : 1,
										"faceoffTaken" : 2,
										"takeaways" : 1,
										"giveaways" : 1,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "8:39",
										"powerPlayTimeOnIce" : "1:10",
										"shortHandedTimeOnIce" : "0:39"
									}
								}
							},
							"ID8477471" : {
								"person" : {
									"id" : 8477471,
									"fullName" : "Dillon Heatherington",
									"link" : "/api/v1/people/8477471",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "29",
								"position" : {
									"code" : "N/A",
									"name" : "Unknown",
									"type" : "Unknown",
									"abbreviation" : "N/A"
								},
								"stats" : { }
							},
							"ID8482116" : {
								"person" : {
									"id" : 8482116,
									"fullName" : "Tim Stützle",
									"link" : "/api/v1/people/8482116",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "18",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:58",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 2,
										"faceOffPct" : 33.0,
										"faceOffWins" : 1,
										"faceoffTaken" : 3,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "6:15",
										"powerPlayTimeOnIce" : "1:56",
										"shortHandedTimeOnIce" : "1:47"
									}
								}
							},
							"ID8480355" : {
								"person" : {
									"id" : 8480355,
									"fullName" : "Mark Kastelic",
									"link" : "/api/v1/people/8480355",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "47",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "6:28",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 56.0,
										"faceOffWins" : 5,
										"faceoffTaken" : 9,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 0,
										"evenTimeOnIce" : "5:57",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "0:31"
									}
								}
							},
							"ID8482105" : {
								"person" : {
									"id" : 8482105,
									"fullName" : "Jake Sanderson",
									"link" : "/api/v1/people/8482105",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "85",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "10:07",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 2,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 1,
										"evenTimeOnIce" : "6:46",
										"powerPlayTimeOnIce" : "0:51",
										"shortHandedTimeOnIce" : "2:30"
									}
								}
							}
						},
						"goalies" : [ 8476341, 8476433 ],
						"skaters" : [ 8474207, 8480801, 8479337, 8477353, 8475766, 8482116, 8480208, 8478472, 8479458, 8474612, 8480073, 8479580, 8473512, 8480448, 8480355, 8481596, 8478469, 8482105, 8477471, 8473544 ],
						"onIce" : [ 8474612, 8475766, 8476341, 8477353, 8480073, 8480355 ],
						"onIcePlus" : [ {
							"playerId" : 8474612,
							"shiftDuration" : 26,
							"stamina" : 100
						}, {
							"playerId" : 8475766,
							"shiftDuration" : 4,
							"stamina" : 100
						}, {
							"playerId" : 8476341,
							"shiftDuration" : 256,
							"stamina" : 33
						}, {
							"playerId" : 8477353,
							"shiftDuration" : 5,
							"stamina" : 100
						}, {
							"playerId" : 8480073,
							"shiftDuration" : 24,
							"stamina" : 100
						}, {
							"playerId" : 8480355,
							"shiftDuration" : 20,
							"stamina" : 100
						} ],
						"scratches" : [ 8477471, 8473544 ],
						"penaltyBox" : [ ],
						"coaches" : [ {
							"person" : {
								"fullName" : "D.J. Smith",
								"link" : "/api/v1/people/null"
							},
							"position" : {
								"code" : "HC",
								"name" : "Head Coach",
								"type" : "Head Coach",
								"abbreviation" : "Head Coach"
							}
						} ]
					},
					"home" : {
						"team" : {
							"id" : 14,
							"name" : "Tampa Bay Lightning",
							"link" : "/api/v1/teams/14",
							"abbreviation" : "TBL",
							"triCode" : "TBL"
						},
						"teamStats" : {
							"teamSkaterStats" : {
								"goals" : 1,
								"pim" : 9,
								"shots" : 26,
								"powerPlayPercentage" : "16.7",
								"powerPlayGoals" : 1.0,
								"powerPlayOpportunities" : 6.0,
								"faceOffWinPercentage" : "46.2",
								"blocked" : 17,
								"takeaways" : 5,
								"giveaways" : 5,
								"hits" : 19
							}
						},
						"players" : {
							"ID8481043" : {
								"person" : {
									"id" : 8481043,
									"fullName" : "Cole Koepke",
									"link" : "/api/v1/people/8481043",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "45",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "7:45",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:43",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "0:02"
									}
								}
							},
							"ID8477938" : {
								"person" : {
									"id" : 8477938,
									"fullName" : "Haydn Fleury",
									"link" : "/api/v1/people/8477938",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "7",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:27",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 0,
										"evenTimeOnIce" : "8:24",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "1:03"
									}
								}
							},
							"ID8470880" : {
								"person" : {
									"id" : 8470880,
									"fullName" : "Brian Elliott",
									"link" : "/api/v1/people/8470880",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "1",
								"position" : {
									"code" : "G",
									"name" : "Goalie",
									"type" : "Goalie",
									"abbreviation" : "G"
								},
								"stats" : {
									"goalieStats" : {
										"timeOnIce" : "0:00",
										"assists" : 0,
										"goals" : 0,
										"pim" : 0,
										"shots" : 0,
										"saves" : 0,
										"powerPlaySaves" : 0,
										"shortHandedSaves" : 0,
										"evenSaves" : 0,
										"shortHandedShotsAgainst" : 0,
										"evenShotsAgainst" : 0,
										"powerPlayShotsAgainst" : 0
									}
								}
							},
							"ID8477426" : {
								"person" : {
									"id" : 8477426,
									"fullName" : "Nicholas Paul",
									"link" : "/api/v1/people/8477426",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "20",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "10:42",
										"assists" : 0,
										"goals" : 0,
										"shots" : 2,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 67.0,
										"faceOffWins" : 2,
										"faceoffTaken" : 3,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 0,
										"evenTimeOnIce" : "9:44",
										"powerPlayTimeOnIce" : "0:01",
										"shortHandedTimeOnIce" : "0:57"
									}
								}
							},
							"ID8478416" : {
								"person" : {
									"id" : 8478416,
									"fullName" : "Erik Cernak",
									"link" : "/api/v1/people/8478416",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "81",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "11:06",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 4,
										"plusMinus" : 0,
										"evenTimeOnIce" : "9:01",
										"powerPlayTimeOnIce" : "0:01",
										"shortHandedTimeOnIce" : "2:04"
									}
								}
							},
							"ID8470621" : {
								"person" : {
									"id" : 8470621,
									"fullName" : "Corey Perry",
									"link" : "/api/v1/people/8470621",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "10",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:15",
										"assists" : 0,
										"goals" : 0,
										"shots" : 2,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -2,
										"evenTimeOnIce" : "3:54",
										"powerPlayTimeOnIce" : "5:21",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8473986" : {
								"person" : {
									"id" : 8473986,
									"fullName" : "Alex Killorn",
									"link" : "/api/v1/people/8473986",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "17",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "14:35",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 1,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -1,
										"evenTimeOnIce" : "7:41",
										"powerPlayTimeOnIce" : "5:06",
										"shortHandedTimeOnIce" : "1:48"
									}
								}
							},
							"ID8476453" : {
								"person" : {
									"id" : 8476453,
									"fullName" : "Nikita Kucherov",
									"link" : "/api/v1/people/8476453",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "86",
								"position" : {
									"code" : "R",
									"name" : "Right Wing",
									"type" : "Forward",
									"abbreviation" : "RW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "15:08",
										"assists" : 1,
										"goals" : 0,
										"shots" : 2,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 1,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 1,
										"giveaways" : 1,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : -2,
										"evenTimeOnIce" : "5:38",
										"powerPlayTimeOnIce" : "9:30",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8479525" : {
								"person" : {
									"id" : 8479525,
									"fullName" : "Ross Colton",
									"link" : "/api/v1/people/8479525",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "79",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "9:44",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 3,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 0.0,
										"faceOffWins" : 0,
										"faceoffTaken" : 1,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 1,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:08",
										"powerPlayTimeOnIce" : "2:36",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8476883" : {
								"person" : {
									"id" : 8476883,
									"fullName" : "Andrei Vasilevskiy",
									"link" : "/api/v1/people/8476883",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "88",
								"position" : {
									"code" : "G",
									"name" : "Goalie",
									"type" : "Goalie",
									"abbreviation" : "G"
								},
								"stats" : {
									"goalieStats" : {
										"timeOnIce" : "39:38",
										"assists" : 0,
										"goals" : 0,
										"pim" : 0,
										"shots" : 8,
										"saves" : 6,
										"powerPlaySaves" : 2,
										"shortHandedSaves" : 1,
										"evenSaves" : 3,
										"shortHandedShotsAgainst" : 2,
										"evenShotsAgainst" : 4,
										"powerPlayShotsAgainst" : 2,
										"savePercentage" : 75.0,
										"powerPlaySavePercentage" : 100.0,
										"shortHandedSavePercentage" : 50.0,
										"evenStrengthSavePercentage" : 75.0
									}
								}
							},
							"ID8475167" : {
								"person" : {
									"id" : 8475167,
									"fullName" : "Victor Hedman",
									"link" : "/api/v1/people/8475167",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "77",
								"position" : {
									"code" : "N/A",
									"name" : "Unknown",
									"type" : "Unknown",
									"abbreviation" : "N/A"
								},
								"stats" : { }
							},
							"ID8477930" : {
								"person" : {
									"id" : 8477930,
									"fullName" : "Pierre-Edouard Bellemare",
									"link" : "/api/v1/people/8477930",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "41",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "6:38",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 0.0,
										"faceOffWins" : 0,
										"faceoffTaken" : 1,
										"takeaways" : 0,
										"giveaways" : 2,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 3,
										"plusMinus" : -1,
										"evenTimeOnIce" : "4:35",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "2:03"
									}
								}
							},
							"ID8479984" : {
								"person" : {
									"id" : 8479984,
									"fullName" : "Cal Foote",
									"link" : "/api/v1/people/8479984",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "52",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "10:54",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 2,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:35",
										"powerPlayTimeOnIce" : "2:20",
										"shortHandedTimeOnIce" : "0:59"
									}
								}
							},
							"ID8474564" : {
								"person" : {
									"id" : 8474564,
									"fullName" : "Steven Stamkos",
									"link" : "/api/v1/people/8474564",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "91",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "16:56",
										"assists" : 1,
										"goals" : 0,
										"shots" : 7,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 1,
										"penaltyMinutes" : 0,
										"faceOffPct" : 40.0,
										"faceOffWins" : 2,
										"faceoffTaken" : 5,
										"takeaways" : 0,
										"giveaways" : 1,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -1,
										"evenTimeOnIce" : "7:29",
										"powerPlayTimeOnIce" : "9:17",
										"shortHandedTimeOnIce" : "0:10"
									}
								}
							},
							"ID8479542" : {
								"person" : {
									"id" : 8479542,
									"fullName" : "Brandon Hagel",
									"link" : "/api/v1/people/8479542",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "38",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "10:39",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 50.0,
										"faceOffWins" : 1,
										"faceoffTaken" : 2,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 0,
										"evenTimeOnIce" : "6:48",
										"powerPlayTimeOnIce" : "2:47",
										"shortHandedTimeOnIce" : "1:04"
									}
								}
							},
							"ID8479410" : {
								"person" : {
									"id" : 8479410,
									"fullName" : "Mikhail Sergachev",
									"link" : "/api/v1/people/8479410",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "98",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "16:18",
										"assists" : 0,
										"goals" : 1,
										"shots" : 3,
										"hits" : 1,
										"powerPlayGoals" : 1,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 4,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -1,
										"evenTimeOnIce" : "7:59",
										"powerPlayTimeOnIce" : "8:19",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8474013" : {
								"person" : {
									"id" : 8474013,
									"fullName" : "Ian Cole",
									"link" : "/api/v1/people/8474013",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "28",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "11:23",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 2,
										"plusMinus" : 0,
										"evenTimeOnIce" : "9:10",
										"powerPlayTimeOnIce" : "0:08",
										"shortHandedTimeOnIce" : "2:05"
									}
								}
							},
							"ID8479026" : {
								"person" : {
									"id" : 8479026,
									"fullName" : "Philippe Myers",
									"link" : "/api/v1/people/8479026",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "5",
								"position" : {
									"code" : "N/A",
									"name" : "Unknown",
									"type" : "Unknown",
									"abbreviation" : "N/A"
								},
								"stats" : { }
							},
							"ID8474034" : {
								"person" : {
									"id" : 8474034,
									"fullName" : "Pat Maroon",
									"link" : "/api/v1/people/8474034",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "14",
								"position" : {
									"code" : "L",
									"name" : "Left Wing",
									"type" : "Forward",
									"abbreviation" : "LW"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "1:46",
										"assists" : 0,
										"goals" : 0,
										"shots" : 0,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 5,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "1:15",
										"powerPlayTimeOnIce" : "0:31",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8478010" : {
								"person" : {
									"id" : 8478010,
									"fullName" : "Brayden Point",
									"link" : "/api/v1/people/8478010",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "21",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "13:20",
										"assists" : 0,
										"goals" : 0,
										"shots" : 3,
										"hits" : 1,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 57.0,
										"faceOffWins" : 4,
										"faceoffTaken" : 7,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -1,
										"evenTimeOnIce" : "4:17",
										"powerPlayTimeOnIce" : "9:03",
										"shortHandedTimeOnIce" : "0:00"
									}
								}
							},
							"ID8476480" : {
								"person" : {
									"id" : 8476480,
									"fullName" : "Vladislav Namestnikov",
									"link" : "/api/v1/people/8476480",
									"shootsCatches" : "L",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "90",
								"position" : {
									"code" : "C",
									"name" : "Center",
									"type" : "Forward",
									"abbreviation" : "C"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "12:55",
										"assists" : 0,
										"goals" : 0,
										"shots" : 2,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffPct" : 43.0,
										"faceOffWins" : 3,
										"faceoffTaken" : 7,
										"takeaways" : 1,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : 0,
										"evenTimeOnIce" : "7:12",
										"powerPlayTimeOnIce" : "5:35",
										"shortHandedTimeOnIce" : "0:08"
									}
								}
							},
							"ID8480246" : {
								"person" : {
									"id" : 8480246,
									"fullName" : "Nick Perbix",
									"link" : "/api/v1/people/8480246",
									"shootsCatches" : "R",
									"rosterStatus" : "Y"
								},
								"jerseyNumber" : "48",
								"position" : {
									"code" : "D",
									"name" : "Defenseman",
									"type" : "Defenseman",
									"abbreviation" : "D"
								},
								"stats" : {
									"skaterStats" : {
										"timeOnIce" : "7:51",
										"assists" : 0,
										"goals" : 0,
										"shots" : 1,
										"hits" : 0,
										"powerPlayGoals" : 0,
										"powerPlayAssists" : 0,
										"penaltyMinutes" : 0,
										"faceOffWins" : 0,
										"faceoffTaken" : 0,
										"takeaways" : 0,
										"giveaways" : 0,
										"shortHandedGoals" : 0,
										"shortHandedAssists" : 0,
										"blocked" : 0,
										"plusMinus" : -1,
										"evenTimeOnIce" : "7:50",
										"powerPlayTimeOnIce" : "0:00",
										"shortHandedTimeOnIce" : "0:01"
									}
								}
							}
						},
						"goalies" : [ 8470880, 8476883 ],
						"skaters" : [ 8477938, 8470621, 8474034, 8473986, 8477426, 8478010, 8474013, 8479542, 8477930, 8481043, 8480246, 8479984, 8479525, 8478416, 8476453, 8476480, 8474564, 8479410, 8479026, 8475167 ],
						"onIce" : [ 8476480, 8476883, 8477938, 8479525, 8479984, 8481043 ],
						"onIcePlus" : [ {
							"playerId" : 8476480,
							"shiftDuration" : 58,
							"stamina" : 66
						}, {
							"playerId" : 8476883,
							"shiftDuration" : 400,
							"stamina" : 33
						}, {
							"playerId" : 8477938,
							"shiftDuration" : 58,
							"stamina" : 66
						}, {
							"playerId" : 8479525,
							"shiftDuration" : 55,
							"stamina" : 66
						}, {
							"playerId" : 8479984,
							"shiftDuration" : 59,
							"stamina" : 66
						}, {
							"playerId" : 8481043,
							"shiftDuration" : 56,
							"stamina" : 66
						} ],
						"scratches" : [ 8479026, 8475167 ],
						"penaltyBox" : [ ],
						"coaches" : [ {
							"person" : {
								"fullName" : "Jon Cooper",
								"link" : "/api/v1/people/null"
							},
							"position" : {
								"code" : "HC",
								"name" : "Head Coach",
								"type" : "Head Coach",
								"abbreviation" : "Head Coach"
							}
						} ]
					}
				},
				"officials" : [ {
					"official" : {
						"id" : 6992,
						"fullName" : "Furman South",
						"link" : "/api/v1/people/6992"
					},
					"officialType" : "Referee"
				}, {
					"official" : {
						"id" : 7404,
						"fullName" : "Michael Markovic",
						"link" : "/api/v1/people/7404"
					},
					"officialType" : "Referee"
				}, {
					"official" : {
						"id" : 4694,
						"fullName" : "Shandor Alphonso",
						"link" : "/api/v1/people/4694"
					},
					"officialType" : "Linesman"
				}, {
					"official" : {
						"id" : 7664,
						"fullName" : "Kyle Flemington",
						"link" : "/api/v1/people/7664"
					},
					"officialType" : "Linesman"
				} ]
			},
			"decisions" : { }
		}
	}`

	gu := NHLGameUpdate{}
	err := json.Unmarshal([]byte(data), &gu)
	if err != nil {
		panic(err)
	}

	if gu.ID != 2022020148 {
		t.Errorf("Game ID. Got %d, want %d", gu.ID, 2022020148)
	}
	if gu.Status != "In Progress" {
		t.Errorf("Game status. Got %s, want %s", gu.Status, "Live")
	}
	if gu.Period != 2 {
		t.Errorf("Game period. Got %d, want %d", gu.Period, 2)
	}
	if gu.HomeScore != 1 {
		t.Errorf("Home score. Got %d, want %d", gu.HomeScore, 1)
	}
	if gu.VisitorScore != 2 {
		t.Errorf("Visitor score. Got %d, want %d", gu.VisitorScore, 2)
	}
}
