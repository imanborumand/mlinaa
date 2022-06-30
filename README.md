# Mlinaa
## _log service_

Mlinaa is a Golang-language microservice to which you can send and save log requests from your other services.

With this service you can create a log service on your host.

You can send your system logs to Mlinaa by requesting HTTP and then retrieve them.



## Features
_Melina has three Api to manage your logs_
- set => store logs
- get/:id => get log by id
- delete/:id => delete a log from database

## Configure
The mlinna settings are in the config/config.yml file

## Installation

```sh
docker-compose build
docker-compose up -d
```

After installation, call the following address:

`http://localhost:8080`

now you can [store- get- remove] your logs




