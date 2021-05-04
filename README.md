# IP2Country

IP2Country app exposes a REST API to get a location for an ip address.

## Prerequisites
Clone the repository

Make sure to have [Docker](https://docs.docker.com/get-docker/) installed on your machine

## Usage
```bash
cd IP2Country
docker build -t ip2country .
docker run -p 8080:8080 ip2country
```