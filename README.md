# go-server-sample
⚡ REST API sample written in Go using Gin, Gorm and JWT Authentication

## Songs-Share
Songs-Share used to be the concept of the project as a feed to share songs.
This project is created for educational purposes.
Any contribution or improvement is welcome.

## What's inside?
This API is capable to store Users with authentication capabilities.

## Getting Started
Theres two ways of running this project, either locally (requires *Go* and *dep* installed in your system), or with Docker (requires *Docker* installed in your system).

### Locally
You will have to clone the project in your `GOPATH/src/github.com/estebanborai/*` directory.
In order to achieve that, run the following command:

```bash
# create github.com/estebanborai directory in your GOPATH/src directory
mkdir -p $GOPATH/src/github.com/estebanborai
```

Then you will be able to clone the project repository as follows:
```bash
# step into the main file directory
cd $GOPATH/src/github.com/estebanborai/

# clone the project in your system using Git
git clone https://github.com/estebanborai/songs-share-server.git
```

Songs-Share server is a GIN powered REST API that connects to MySQL using an ORM (Object Relational Mapping) library called GORM,
you will need to setup MySQL or build and run the docker image that comes with the repository which is ready for this purpose.

`docker-compose -f docker-compose.db.yml up --build`

Now we are able to run our database, first open a new terminal, as the following terminal will be attached to the *Docker/MySQL* container process.

```bash
# install dependencies using dep
dep ensure

# make sure you have MySQL server running in your system
# and run main.go as follows
go run main.go
```

### Docker

```bash
# run docker-compose
docker-compose up --build

# Expect the following log
Starting songs-share-db ... done
Recreating songs-share-server ... done
```

Now you are able to SSH into the Docker Container for the server and the database instances of SongsShare.

#### Running Server (SSH into Docker Container)

In order to run the server, SSH into `songs-share-server` container.
- Make sure you are in the following path:

`root@<your container id>:/go/src/github.com/estebanborai/songs-share-server/server/src/#`

- Install dependencies using `dep`:
`dep ensure`

- Then run `go run main.go` command and expect the following output:

`[GIN-debug] Listening and serving HTTP on :8080`

You are able to make HTTP requests to this server using Postman or any other tool.

## References
[API Documentation](https://github.com/estebanborai/songs-share-server/server/src/blob/master/docs/Api.md)
