# Personal Information
Personal Information is sample boilerplate usecase for workshop tokopedia academy. 

## Configuration & Dependency
- [Docker Compose](https://docs.docker.com/compose/install/) : Container manager.
- [MySQL](https://hub.docker.com/_/mysql): Used docker mysql image.
- [Adminer](https://hub.docker.com/_/adminer): Used for manage mysql on docker.

## How to run the project
### Initial config docker env configuration
```
# Copy the config env
piilot$ cp .env.sample .env

# Change password default mysql root on .env
MYSQL_ROOT_PASSWORD=sample

# Run docker for installation
docker-compose up
```
### Initial config docker env configuration
```
# Copy the config env
piilot/files/etc/stat-hybrid/$ cp piilot.development.yaml.sample piilot.development.yaml

# Change the config piilot.development.yaml
version: "0.0.1"

server: 
  name: "piilot"
  http: 
    address: ":8080"
    write_timeout: 1
    read_timeout: 2
    max_header_bytes: 500000
    enable: true


database:
  testing: true
  dsn: "root:sample@tcp(127.0.0.1:3306)/test_pii_lot?charset=utf8mb4&parseTime=True&loc=Local"
  max_conns: 15
  max_idle_conns: 5
  max_retry: 5

new_relic:
  app_name: "Personal Information"
  secret: "this_secret_generated" #generated from one.newrelic.com
```

### Running golang apps
```
piilot$ make app
```
## The endpoint list
### Get PII List

- GET [/pii/list](http://localhost:8080/pii/list)
```
{
    "data": [
        {
            "plat_no": "DA123DEF",
            "slot_number": 1,
            "registration_date": "2022-03-12T01:49:41+08:00"
        }
    ]
}
```