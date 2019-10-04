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
