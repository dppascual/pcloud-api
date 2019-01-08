# PANTAHUB API

- [Introduction](#introduction)
- [Getting started](#getting-started)
  - [Testing](#testing)
  - [Production](#production)

# Introduction

The Pantahub API is an HTTP API served by Pantahub Cloud. It is the API the Pantahub client uses to communicate with the Cloud, so everything the Pantahub client can do can be done with the API.

The API specification can be found in the link below.

> URL: [Pantahub API specification](https://app.swaggerhub.com/apis-docs/dppascual/pantahub-api/1.0.0)

# Getting started

First of all, a Docker CE Engine has to be installed. It can be done by following the below steps.

> INFO: [How to install a Docker CE Engine](https://www.docker.com/products/docker-engine)

Once the Docker CE Engine has been installed, the code can be tested and deployed. Commands must be executed in the repository folder.

## Testing

Build an image from the `Dockerfile.test`:

```
docker build -t <username>/golang-test -f Dockerfile.test .
```

Provision a new container based on the previous image and run a shell on it:

```
docker run --name testing --rm -it -v ${PWD}:/api <username>/golang-test
```

Create dependencies and test the code:

```
bash-4.4# go mod init "github.com/dppascual/pantahub-api"
go: creating new go.mod: module github.com/dppascual/pantahub-api

bash-4.4# go build -v .
go: finding github.com/gorilla/mux v1.6.2
go: downloading github.com/gorilla/mux v1.6.2
golang_org/x/crypto/cryptobyte/asn1
golang_org/x/net/dns/dnsmessage
golang_org/x/crypto/curve25519
golang_org/x/crypto/cryptobyte
golang_org/x/crypto/internal/chacha20
golang_org/x/crypto/poly1305
golang_org/x/crypto/chacha20poly1305
golang_org/x/text/transform
net
golang_org/x/text/unicode/bidi
golang_org/x/text/unicode/norm
golang_org/x/net/http2/hpack
golang_org/x/text/secure/bidirule
golang_org/x/net/idna
golang_org/x/net/http/httpproxy
net/textproto
crypto/x509
golang_org/x/net/http/httpguts
crypto/tls
net/http/httptrace
net/http
github.com/gorilla/mux
github.com/dppascual/pantahub-api

bash-4.4# go test -v
=== RUN   TestGetDeviceStats
--- PASS: TestGetDeviceStats (0.00s)
=== RUN   TestGetDeviceStatsNonExistent
--- PASS: TestGetDeviceStatsNonExistent (0.00s)
PASS
ok  	github.com/dppascual/pantahub-api	0.006s
```

## Production

Build an image from the `Dockerfile`:

```
docker build -t <username>/pantahub-api .
```

Deploy the Pantahub API:

```
docker run --name pantahub-api -p 80:80 --rm -d <username>/pantahub-api
```

Automate the deployment (Docker installation and API Pantahub deployment) with the Ansible usage:

```
ansible-playbook --private-key="~/.ssh/id_rsa" -i 10.60.128.33, api_deployment.yml                                                                                                             (ansible)

PLAY [all] ***********************************************************************************************************************************************************************************************************************************

TASK [Gathering Facts] ***********************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:15 +0100 (0:00:00.040)       0:00:00.040 *******
ok: [10.60.128.33]

TASK [Add Docker GPG key] ********************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:22 +0100 (0:00:06.338)       0:00:06.378 *******
ok: [10.60.128.33]

TASK [Add Docker APT repository] *************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:23 +0100 (0:00:01.781)       0:00:08.159 *******
ok: [10.60.128.33]

TASK [Install list of packages] **************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:25 +0100 (0:00:01.329)       0:00:09.488 *******
ok: [10.60.128.33] => (item=[u'apt-transport-https', u'ca-certificates', u'curl', u'software-properties-common', u'docker-ce', u'python-pip'])

TASK [Install docker-py python package] ******************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:28 +0100 (0:00:02.772)       0:00:12.261 *******
ok: [10.60.128.33]

TASK [Create a network] **********************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:30 +0100 (0:00:02.454)       0:00:14.715 *******
ok: [10.60.128.33]

TASK [Launch api container] ******************************************************************************************************************************************************************************************************************
Tuesday 08 January 2019  00:21:31 +0100 (0:00:01.177)       0:00:15.893 *******
changed: [10.60.128.33]

PLAY RECAP ***********************************************************************************************************************************************************************************************************************************
10.60.128.33               : ok=7    changed=1    unreachable=0    failed=0

Tuesday 08 January 2019  00:21:38 +0100 (0:00:07.227)       0:00:23.121 *******
===============================================================================
Launch api container ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ 7.23s
Gathering Facts ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 6.34s
Install list of packages -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 2.77s
Install docker-py python package ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ 2.45s
Add Docker GPG key -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 1.78s
Add Docker APT repository ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 1.33s
Create a network ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- 1.18s

```

Send a request by using a HTTP client:

```
curl -sS -X GET -H "Accept: application/json" "http://10.60.128.33/devices/1/stats" | python -m json.tool
{
    "cpu": {
        "id": 0,
        "st": 0,
        "sy": 99,
        "us": 0,
        "wa": 0
    },
    "io": {
        "bi": 13,
        "bo": 69
    },
    "mem": {
        "buff": 40324,
        "cache": 1120928,
        "free": 2367340,
        "swpd": 0
    },
    "procs": {
        "b": 0,
        "r": 0
    },
    "read": "2019-01-07T23:22:43.202485791Z",
    "swap": {
        "si": 0,
        "so": 0
    },
    "system": {
        "cs": 145,
        "in": 33
    }
}
```

