# GraphQL based API written in Golang
[![Build Status](https://travis-ci.org/raunofreiberg/kyrene.svg?branch=master)](https://travis-ci.org/raunofreiberg/kyrene)

> Work in progress

This project serves as a GraphQL based API written in Golang, which is consumed by the front-end via the Apollo client using React.
State management is done via `apollo-link-state` and authentication is provided by JSON web tokens. Still work in progress.

# Features

- [x] Base queries & mutations for GraphQL
- [x] Actual working endpoints for queries & mutations
- [x] Database integration
- [ ] Unit & integration tests for endpoints
- [x] Front end base setup
- [ ] Client side functionality integration
- [x] JWT authentication
- [x] Docker
	- [x] Development
	- [x] Production

# Technologies
- Golang
- GraphQL
- PostgreSQL
- Docker
- React
- React-Router
- Apollo
- Webpack

# Setup

#### Docker

Development
```
docker-compose up --build
```

¯\_(ツ)_/¯

#### Non-Docker

* Start by setting env variables for Go (Note: `$` implies using the terminal)

```
$ export GOPATH=$HOME/path/to/go/projects
$ export PATH=$GOPATH/bin:$PATH
```

* Proceed with Postgres env variables

```
$ export DBHOST=localhost
$ export DBPORT=<PORT> // default 5432
$ export DBUSER=<YOUR_USERNAME>
$ export DBPASS=<YOUR_PASSWORD>
$ export DBNAME=<YOUR_DB_NAME>
```

* Install dependencies via Glide

```
glide install
````

* Development (served on localhost:8080)

```
$ npm run dev
```

* Production (served on localhost:8000)

```
$ npm run prod
```

#### Endpoint examples

* Retrieve single object

Request
```
curl -g 'http://localhost:8000/graphql?query={todo(id:<ID>){id,content,isCompleted}}'
```

Response
```
{
	"data": {
		"todo": {
			"content": "hello world",
			"id": "3",
			"isCompleted": false
		}
	}
}
```

