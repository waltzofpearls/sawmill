# Sawmill

[![Build Status](https://travis-ci.org/waltzofpearls/sawmill.svg)](https://travis-ci.org/waltzofpearls/sawmill)

Sawmill is a simple web service with a REST API that helps check if a given URL/website could contain
phishing, malware, viruses, unwanted software or reported suspicious.

## Get started

```
git clone https://github.com/waltzofpearls/sawmill.git
cd sawmill
make docker
```

## API Endpoints

```
GET /urlinfo/1/{hostname_and_port}/{original_path_and_query_string}
```
