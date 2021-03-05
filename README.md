Data Sources

TMS
  * get tv movie listings (/v1/tv-movies?date={YYYY-MM-DD})
  * get tv sports listings (/v1/tv-sports?date={YYYY-MM-DD})
  * get cinema listings by zip (/v1/movies?zip={ZIP})

TV Maze
  * tv show search (/v1/tv-search?title={TITLE})

TheMovieDB.org
  * discover movies (/v1/discover?date={YYYY-MM-DD})
  * TODO: tv show popular: https://developers.themoviedb.org/3/tv/get-popular-tv-shows

TODO:
  - set a username and pass on Redis; externalize to creds.yaml file
  - add caching to movie by zip lookup, GetTMSReq
  - fix episode links showing on search results
  - fetch theaters in zip code

Creds:

Application looks for a "creds.yaml" file in the root of the project. This file is required to build, as it is read by Viper to set up the application.

  - tms: App ID for TMS.
  - moviedb: Access key for Movie DB
  - port: Port number
  - env: Environment flag, e.g. "dev" or "prod"
  - redis_port: URL for Redis with hostname and port
  - redis_password: Redis password
  - redis_db: Redis database number
  - use_cache: Boolean flag to enable cache/storage
  - timezone: Default timezone for server, used in querying results