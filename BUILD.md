## Building the Application
### Logging
- We are using the [lumberjack](https://pkg.go.dev/gopkg.in/lumberjack.v2) library to write logs and handle log rotation.
- [logrus](https://pkg.go.dev/github.com/sirupsen/logrus) for appending logs
- For local deployment, it is available on STDOUT using the format configured in [logger.go](../src/config/logger.go)

## Setup Guide
### Requirements

For building and running the application you need:

- [GO](https://go.dev/)

### **Here is our quickstart guide.**
* Clone the repo
```shell  
git clone git@github.com:AlexKimani/mpesa-daraja-api-go.git  
```  
* When you attempt to clone the repository, you receive the error message. [Fix – git@github.com : permission denied](https://dev.classmethod.jp/articles/fix-gitgithub-com-permission-denied-publickey-fatal-could-not-read-from-remote-repository/)
* [Install docker](https://docs.docker.com/get-docker/). Ensure that docker is always running.

## Running and testing
* via IDE for local debugging (recommended)
* Run all the required Services on Docker `docker-compose up`. 