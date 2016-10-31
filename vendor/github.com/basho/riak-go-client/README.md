Riak Go Client
==================

The **Riak Go Client** is a client which makes it easy to communicate with
[Riak](http://basho.com/riak/), an open source, distributed database that
focuses on high availability, horizontal scalability, and *predictable* latency.
Both Riak and this code is maintained by [Basho](http://www.basho.com/).

1. [Installation](#installation)
2. [Documentation](#documentation)
3. [Contributing](#contributing)
	* [An honest disclaimer](#an-honest-disclaimer)
4. [Roadmap](#roadmap)
5. [License and Authors](#license-and-authors)

## Build Status

[![Build Status](https://travis-ci.org/basho/riak-go-client.svg?branch=master)](https://travis-ci.org/basho/riak-go-client)

## Installation

`go get github.com/basho/riak-go-client`

## Documentation

* [API documentation on Godoc](https://godoc.org/github.com/basho/riak-go-client)
* [Wiki](https://github.com/basho/riak-go-client/wiki)
* [Release Notes](https://github.com/basho/riak-go-client/blob/master/RELNOTES.md). 

## Contributing

*Note:* Please clone this repository in such a manner that submodules are also cloned:

```
git clone --recursive https://github.com/basho/riak-go-client
```

OR:

```
git clone https://github.com/basho/riak-go-client
git submodule init --update
```

This repository's maintainers are engineers at Basho and we welcome your contribution to the project! Review the details in [CONTRIBUTING.md](CONTRIBUTING.md) in order to give back to this project.

### An honest disclaimer

Due to our obsession with stability and our rich ecosystem of users, community updates on this repo may take a little longer to review. 

The most helpful way to contribute is by reporting your experience through issues. Issues may not be updated while we review internally, but they're still incredibly appreciated.

Thank you for being part of the community! We love you for it. 

## Roadmap

* 1.0.0 - Full Riak 2 support with command queuing and retries.

## License

The **Riak Go Client** is Open Source software released under the Apache 2.0
License. Please see the [LICENSE](LICENSE) file for full license details.

These excellent community projects inspired this client and parts of their code
are in `riak-go-client` as well:

* [`goriakpbc`](https://github.com/tpjg/goriakpbc)
* [`riaken-core`](https://github.com/riaken/riaken-core)
* [`backoff`](https://github.com/jpillora/backoff)

## Authors

* [Luke Bakken](https://github.com/lukebakken)
* [Christopher Mancini](https://github.com/christophermancini)

## Contributors

Thank you to all of our contributors!

* [Ian Lozinski](https://github.com/i)
* [Sergio C. Arteaga](https://github.com/tegioz)
* [Andrew Zeneski](https://github.com/andrewzeneski)
* [Кирилл Александрович Журавлев](https://github.com/kazhuravlev)
* [Paul Guelpa](https://github.com/pguelpa)
* [Xabier Larrakoetxea Gallego](https://github.com/slok)
