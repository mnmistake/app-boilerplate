# GraphQL based API written in Golang
[![Build Status](https://travis-ci.org/raunofreiberg/kyrene.svg?branch=master)](https://travis-ci.org/raunofreiberg/kyrene)

- [x] Base queries & mutations for GraphQL
- [x] Actual working endpoints for queries & mutations
- [x] Database integration
- [ ] Unit & integration tests for endpoints
- [x] Front end base setup
- [ ] Client side functionality integration
- [x] Docker
	- [x] Development
	- [ ] Production

#### Setup

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

