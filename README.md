## Description

Simple implementation loosely following this [tutorial](https://go.dev/doc/tutorial/web-service-gin)

## How to run

From `src/` run:
`go run. data/albums.json`

## How to test

GET all albums:
`curl localhost:8080/albums`

GET a specific album by its id:
`curl localhost:8080/albums/:id`

POST a new album:

`curl localhost:8080/albums --request POST --data '{"id":"4","title": "twilight 7", "artist": "Clint Eastwood", "price": 1000.00}'`