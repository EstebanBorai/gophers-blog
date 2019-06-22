# songs-share-server
🎶 Back-End Development for "Songs Share"

### Getting Started
This project is made to run it with Docker, you will need to install Docker in your system in order to run this project.

```bash
# run docker-compose
docker-compose up --build

# Expect the following log
Starting songs-share-db ... done
Recreating songs-share-server ... done
```

Now you are able to SSH into the Docker Container for the server and the database instances of SongsShare.

### Running Server
In order to run the server, SSH into `songs-share-server` container.
- Make sure you are in the following path:

`root@<your container id>:/go/src/github.com/estebanborai/songs-share-server/server/src/#`

- Install dependencies using `dep`:
`dep ensure`

- Then run `go run main.go` command and expect the following output:

`[GIN-debug] Listening and serving HTTP on :8080`

You are able to make HTTP requests to this server using Postman or any other tool.

### References
[API Documentation](https://github.com/estebanborai/songs-share-server/server/src/blob/master/docs/Api.md)
