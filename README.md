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

## Technical overview

Sawmill's Docker stack was built with scale in mind. The API was built with Golang and Riak KV
as database. The default docker setup consists of one load balancer (HAProxy), two API instances
and three Riak instances (with N=3 for replication and 64 for ring size). It can be easily
scaled up or down with docker-compose. This will be covered in the `Get started` section.

In addition to the HAProxy load balancing for API nodes, connection pool is used in the API
for Riak cluster. Each API instance knows about all the Riak nodes, and hosts a pool of connections
to all the Riak nodes. API automatically checks the health of each Riak node, and re-route
traffic to the healthy nodes if any Riak node goes down.

## Get started

You have two options to run sawmill locally: 1) through the Docker setup; 2) compile from source.

#### 1. With Docker

*Sawmill <3 docker*

It's a much easier approach to go with docker.

Requirements:

- Docker
- Docker-compose (with version 2 support)

```shell
git clone https://github.com/waltzofpearls/sawmill.git && cd sawmill
# Build docker images and bring up the containers
make setup
# Load demo fixture data
make fixture
```

By default, `make fixture` command loads `100` records to Riak's `urlinfo` bucket. You can pass
a custom number to the command like:

```shell
make fixture NUM=30
```

To teardown the Docker containers (shutdown and remove), run the following command:

```shell
make teardown
```

For scaling up/down API or Riak KV nodes, run the following command. Note that by default the
Docker setup uses `api=2 db_mem=2` which gives you 2 API instances and 3 Riak instances (one
of them is db_coord). The Docker Riak cluster introduced coordinator/member concept. All the
members will join the coordinator when setting up the cluster.

```shell
make dc-scale SCALE='api=5 d_mem=4'
make dc-up
```

#### 2. Compile from source

Requirements:

- Go 1.7+
- Glide
- Riak KV (single node or cluster)

Compile and run

```shell
make
# Create a new config file based on the example.
# You need to edit the config file based on your setup and needs.
cp config.example.yml config.yml
./sawmill
```

#### 3. Testing

```
# Run unit tests
make test
# Generate code coverage report
make cover
```

## API Endpoints

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
