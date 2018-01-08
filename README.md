# GraphQL based API written in Golang

- [x] Base queries & mutations for GraphQL
- [x] Working endpoint
- [x] Database integration
- [ ] Unit & integration tests
- [ ] Client side API integration
- [ ] Docker

#### Setup

* Start by setting env variables for Go (Note: `$` implies using the terminal)

```
$ export GOPATH=$HOME/path/to/proj
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

* Go crazy

```
$ go run *.go
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

