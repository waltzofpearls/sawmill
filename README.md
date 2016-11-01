# Sawmill

[![Build Status](https://travis-ci.org/waltzofpearls/sawmill.svg)](https://travis-ci.org/waltzofpearls/sawmill)
[![Go Report Card](https://goreportcard.com/badge/github.com/waltzofpearls/sawmill)](https://goreportcard.com/report/github.com/waltzofpearls/sawmill)

```
███████╗ █████╗ ██╗    ██╗███╗   ███╗██╗██╗     ██╗
██╔════╝██╔══██╗██║    ██║████╗ ████║██║██║     ██║
███████╗███████║██║ █╗ ██║██╔████╔██║██║██║     ██║
╚════██║██╔══██║██║███╗██║██║╚██╔╝██║██║██║     ██║
███████║██║  ██║╚███╔███╔╝██║ ╚═╝ ██║██║███████╗███████╗
╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝ ╚═╝     ╚═╝╚═╝╚══════╝╚══════╝
```

Sawmill is a simple web service with a tiny REST API. It helps check if a given URL/website
could contain phishing, malware, viruses, unwanted software or reported suspicious contents.

For more details about technical overview and architecture design, go to `Technical details`
section for more information.

## Get started

You have two options to run sawmill locally: 1) through the Docker setup; 2) compile from source.

### 1. With Docker

*Sawmill <3 docker*

TL;DR: It's a much easier approach to go with Docker, since it uses Compose and has all the
necessary services configured within the `docker-compose.yml` file.

*Requirements*

- Docker
- Docker-compose (with version 2 support)

*Build and run*

```shell
git clone https://github.com/waltzofpearls/sawmill.git && cd sawmill
# Build docker images and bring up the containers
make setup
# Load demo fixture data
make fixture
```

*Data fixture*

By default, `make fixture` loads `100` records to `urlinfo` bucket. You can pass a different
number to the command like `make fixture NUM=30`.

If you see errors from the fixture loading command, try the following steps, and it may fix
your problem:

```
make teardown
make setup
make fixture
```

*Teardown*

To shutdown and remove all the sawmill related Docker containers, run `make teardown`.

*Scaling*

For scaling up/down API or Riak KV nodes, take the following two steps:

(Note that by default `SCALE='api=2 db_mem=2'` which gives you two (2) API instances and
three (3) Riak nodes (one `db_coord` and two `db_mem`). The Docker Riak cluster uses a
coordinator-and-member setup. All the members joins the coordinator when scaling or starting
up)

First, edit `docker/haproxy.cfg` and `docker/config.docker.yml` config files to add or remove
API instances and Riak nodes, and then run:

```shell
make dc-scale SCALE='api=5 d_mem=4'
make dc-up
```

### 2. Compile from source

*Requirements*

- Go 1.7+
- Glide
- Riak KV (single node or cluster)
- HAProxy (optional)

*Build and run*

```shell
make
# Create a new config file based on the example.
# Need to edit the config file based on your setup.
cp config.example.yml config.yml
./sawmill
```

### 3. Testing

For unit test, run `make test`, and to generate code coverage report, run `make cover`.

## API Endpoints

### Check a given URL for malware

*Request*

```
GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
```

*Response*

Code: `200 OK`

Headers: `Content-Type: application/json`

Body:

```json
{
  "url": "reillyjones.org/iliana.raynor?8hzo2=r7x8k242",
  "description": "Est maiores fuga ipsum omnis eveniet quia. Voluptatem excepturi pariatur ab debitis. Quae omnis quia alias.",
  "has_malware":true,
  "created":"2016-11-01T04:43:51.704359757Z",
  "updated":"2016-11-01T04:43:51.704360064Z"
}
```

*Example*

```
curl http://localhost:8000/bartonberge.com/edd?8om1g=j7wdv6zj
```

## Technical details

### Overview

Sawmill was built with scale in mind. The API is written Go, and Riak KV is used as database.
The default Docker setup consists of one (1) load balancer (HAProxy), two (2) API instances
and three (3) Riak nodes (with N=3 for replication and 64 for ring size). It can be easily
scaled up or down with `make scale SCALE='api=10 db_mem=15'` command.

For load balancing and high availability, API instances are balanced by HAProxy with round-robin.
Riak nodes are managed by a connection pool in each API instance. Every API knows about all the
Riak nodes in a cluster. This is done through configuration `config.yml`. The connection pool
automatically checks the health for each Riak node, and re-route traffic to the healthy nodes
if any node goes down.

### Folder structure

```
├── LICENSE ------------------------- MIT License
├── Makefile ------------------------ go development related make targets
├── README.md ----------------------- this file itself
├── app ----------------------------- package for the CLI app
│   ├── api ------------------------- package for the API server
│   │   ├── api.go ------------------ api.New() package entry point
│   │   ├── api_test.go ------------- test file for api
│   │   ├── json_error.go ----------- JsonError struct with status and message attributes for serialization
│   │   ├── middleware.go ----------- HTTP middlewares handles 404 and logging
│   │   ├── middleware_test.go ------ test file middleware
│   │   ├── subroute.go ------------- middleware that handles sub-route with base JSON response handlers
│   │   ├── subroute_test.go -------- test file for subroute
│   │   ├── version1.go ------------- /urlinfo/1 is defined in api, and it points version1 to handle sub-routes
│   │   └── version1_test.go -------- test file for version1
│   ├── app.go ---------------------- app.New() package entry point
│   ├── app_test.go ----------------- test file for app
│   ├── cmd.go ---------------------- cli methods wrapper for easily mocking and testing app
│   ├── config ---------------------- package for YAML configuration
│   │   ├── config.go --------------- config.New() package entry point
│   │   └── config_test.go ---------- test file for config
│   ├── database -------------------- package for Riak database back-end
│   │   ├── database.go ------------- database.New() package entry point, deals with Riak connection
│   │   ├── database_test.go -------- test for database
│   │   └── riak.go ----------------- Riak methods wrapper for easily mocking and testing database
│   ├── logger ---------------------- package for logger (uses Uber zap log)
│   │   ├── logger.go --------------- logger.New()
│   │   └── logger_test.go ---------- test file for logger
│   ├── manager --------------------- package for mediator layer that works with repo and model
│   │   ├── urlinfo.go -------------- manager.NewUrlInfo()
│   │   └── urlinfo_test.go --------- test file for UrlInfo manager
│   ├── model ----------------------- package for ORM models
│   │   ├── model.go ---------------- base model
│   │   ├── urlinfo.go -------------- model.NewUrlInfo()
│   │   └── urlinfo_test.go --------- test file for UrlInfo model
│   └── repository ------------------ package for ORM repositories
│       ├── repository.go ----------- base repository
│       ├── repository_test.go ------ test file for base repository
│       └── urlinfo.go -------------- repository.NewUrlInfo()
├── config.example.yml -------------- an example of config.yml
├── docker -------------------------- everything docker!! (Just kidding)
│   ├── Dockerfile ------------------ the mighty Dockerfile...
│   ├── config.docker.yml ----------- a copy of sawmill config.yml customized the Docker setup
│   ├── docker-compose.yml ---------- and the mighty docker-compose.yml file...
│   ├── docker.mk ------------------- Docker related make targets
│   └── haproxy.cfg ----------------- HAProxy config file
├── fixture ------------------------- data fixture generation program
│   └── gen.go ---------------------- run "make fixture" (Docker) or "go run fixture/gen.go" (no Docker)
├── glide.lock ---------------------- glide package manager lock file
├── glide.yaml ---------------------- glide package manager config file
├── logs ---------------------------- this folder can be used to store all the log files
├── main.go ------------------------- go main.main. Create CLI app and run it
```

### A bit more about Riak KV

The default Docker setup was created for development and demo purposes, and it uses riak's
default Bitcask backend, N=3 replication and ring size 64. In a production environment setup,
backend, replication and ring size need to be planned carefully with actual use case.

*Backend*

Bitcask vs LevelDB

Bitcask loads the keyspace into memory. It's faster with low latency, but if the keys are
larger than what the memory can contain, it will severely impact performance with a lot of
swapping.

*Replication*

N value can be set differently for each bucket or object. Have a look at Riak's official
doc for replication and N value.

http://docs.basho.com/riak/kv/2.1.4/learn/concepts/replication/

*Ring size*

By default, Riak uses 64 as ring size. It's works well for a 3-node development cluster.
For actual production, see the following table for recommended ring size.

http://docs.basho.com/riak/kv/2.1.4/setup/planning/cluster-capacity/#ring-size-number-of-partitions

```
-----------------------------------------------
| Number of nodes | Number of data partitions |
-----------------------------------------------
| 3, 4, 5         | 64, 128                   |
-----------------------------------------------
| 5               | 64, 128                   |
-----------------------------------------------
| 6               | 64, 128, 256              |
-----------------------------------------------
| 7, 8, 9, 10     | 128, 256                  |
-----------------------------------------------
| 11, 12          | 128, 256, 512             |
-----------------------------------------------
```
