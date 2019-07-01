# go-server-sample
⚡ REST API sample written in Go using Gin, Gorm and JWT Authentication

<div align="center">
	<img src="" alt="Gopher Brown" />
</div>

> Hobbyist project that aims to develop a REST API written in Go, using Gin as Web API Framework, Gorm as ORM for MySQL and authentication through JWT.

## Getting Started
This project is written in Go but you can run in your machine using Docker,
with *docker-compose* you will be able to build and run a couple containers, one for
the server and another for the database.

## Requirements
-	[Docker (optional)](https://www.docker.com/get-started)
- [Go (optional/recommended)](https://www.docker.com/get-started)

## Setup
Clone the project. If you are using Docker you can clone it anywhere in your system, if you are using Go
you should clone it inside your `$GOPATH`'s `src` directory.

```bash
# if you are using Go installed locally, create a github.com/username folder for
# this project
mkdir -p $GOPATH/src/github.com/estebanborai

# then clone the project, either anywhere or $GOPATH/src/github.com/estebanborai/
git clone https://github.com/estebanborai/go-server-sample.git
```

The following are instructions for Go users, if you are running the project using *Dcoker* installed in your machine, go to *Docker*'s section.

### Go
1. Install dependencies using *Golang's* [dep](https://golang.github.io/dep/):
```bash
# step into project directory
cd $GOPATH/src/github.com/estebanborai

# gather dependencies
dep ensure
```

2. Running the database using Docker
Docker is required to run the `MySQL` database.
If you are not using Docker, you can also run the project with your local `MySQL` instance.
[src/data/Connection.go]() Manages the connection to the database, you can setup your connection string
from there in order to run this project with your local setup.

> Note: If you are setting up MySQL for this project, please follow Gorm instructions for compatibility with MySQL databases. [Gorm/MySQL Setup](http://gorm.io/docs/connecting_to_the_database.html#MySQL)

Theres two *docker-compose* files you can use in this project, one for a complete environment with Docker (*docker-compose.yml*) and the other for a database only usage (*docker-compose.db.yml*).

For this case you should use the second one.
```bash
# $GOPATH/src/github.com/estebanborai/go-server-sample
docker-compose -f docker-compose.db.yml up --build
```

3. Set your environment variable
Make sure the `.env` file is using `SERVER_ENV` as `LOCAL`
You can find the `.env` file at: `$GOPATH/src/github.com/estebanborai/go-server-sample/server/src/`
```
SERVER_ENV=LOCAL

```

4. Running the server using go
Finally you are able to run the server using *Go*.
```bash
# $GOPATH/src/github.com/estebanborai/go-server-sample/server/src/
go run main.go
```

### Docker
In order to run the project using docker, you should build the docker images for *Go* and *MySQL*, then run them using `docker-compose` command.

1. Step into project root directory
```bash
cd $GOPATH/src/github.com/estebanborai/go-server-sample
```

2. Make sure the environment variable is set to `DOCKER`
```bash
# move to server/src
cd $GOPATH/src/github.com/estebanborai/go-server-sample/server/src/

# print the contents of .env file
cat .env

# expect the following output
SERVER_ENV=DOCKER

# if this argument is set to another value, please set it to 'DOCKER'
```

3. Build and run Docker containers
```bash
# $GOPATH/src/github.com/estebanborai/go-server-sample/
docker-compose up --build
```

The following URL should be runing the project.
`http://localhost:8080/`

You should be able to make requests to this API using Postman or another HTTP requests client.

## Contributions
Contributions and enhancements are welcome for this project, there's a lot of improvements to be made.
