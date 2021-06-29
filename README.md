<div>
  <h1 align="center">gosk</h1>
  <h4 align="center">
    Golang Starter-Kit, a dockerized project with a sample HTTP Server.
    Useful for prototyping and writing tests with Golang
  </h4>
</div>

## Usage

> Make sure you hace Docker, Docker-Compose and Make available in your system

1. Clone this repository or click on the **Use template** button to create a
repository from this template

2. Run `make up` to build and execute the project

```log
api_1  | Running build command!
api_1  | Build ok.
api_1  | Restarting the given command.
api_1  | 2021/05/30 02:58:54 Listening on http://0.0.0.0:4200
```

3. Make a GET HTTP request to `http://0.0.0.0:4200`

```log
$ curl -v http://0.0.0.0:4200
*   Trying 0.0.0.0...
* TCP_NODELAY set
* Connected to 0.0.0.0 (127.0.0.1) port 4200 (#0)
> GET / HTTP/1.1
> Host: 0.0.0.0:4200
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Sun, 30 May 2021 03:03:45 GMT
< Content-Length: 18
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host 0.0.0.0 left intact
Hello from Golang!
* Closing connection 0
```

### Stop the Server and Clean up

Use `Ctrl+C` to stop the Docker process and then run `make clean`

```log
$ make down
docker-compose down
Removing gosk_1 ... done
Removing network gosk_default
Cleaning up process
rm -f out
docker system prune -f
Deleted build cache objects:
o0hk8vfugh3s39gbt18bhmdbr
9frn6250oe8vhx3vermtg42fa
jgn1flx1o9tdo6fzt5wshgxaf
kbk1w9lo5bworx8ixr1waj0qj
4t7egfao289sp8um4gkqfp0dd
6cv0qskklp34c56vcse3aigkb

Total reclaimed space: 12.42MB
docker volume prune -f
Total reclaimed space: 0B
```

## Caveats

* To avoid the error:

```
  Error while building:
  cmd/main.go:5:2: no required module provides package <package url>; to add it:
      go get <package url>
```

The `go.mod` file is removed and the environment variable `GO111MODULE` is set
to `off`.
