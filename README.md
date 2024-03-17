# Description 
This is a simple test project to practice with Go, validate some microservices concepts, use some that that I have in a daily job but don't have much space to play around with or don't work with everyday like configuring things from scratch or dealing with the ops myself.

This project uses the following stack
- __Go__: the main programming language, I want to practice more with it, specially with generics and go routines and channels (async programming)
- __Kafka__: to communicate between services, the idea is to create very small services with the purpose of practicing, I want to configure things over and over to have them clear in my head, specially the problems that I will find
- __Kubernetes__: I know the concepts and how things works, but use too little of it, it will be nice to know how to do at least small tweaks in the app pods

# Starting the application (WIP)
So far I'm using [Rancher Desktop](https://rancherdesktop.io/) for containers and kubernetes, the project has a docker-compose file where I plan to add every part of the application to make the startup process as easy as possible.

### To start the containers
So far I'm only using Kafka
``` bash
docker-compose up -d
```

## Start the application
Use the launch.json for debug or run
``` go
go run ./src/main
```

## Tests
Run tests and coverage
``` bash
./scritps/test_cover.sh
```
## API calls
I'm using curl for that since it's simple and is the most agnostic way to do it
``` bash
curl -X POST -d '{"title": "test", "pages": 100}' localhost/books