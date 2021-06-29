<div>
  <h1 align="center">go-post-api</h1>
  <h4 align="center">
    A minimalist REST API with CRUD capabilities for _Blog Posts_ to learn Go
  </h4>
</div>

## Motivation

This is a "Drive Test" on Golang to learn a bit about the language and
have fun programming.

The solution is pretty simple, a REST API with a CRUD logic on Post.
There's 2 rules for the Post entity:

1. Every post title must have at least 5 characters
2. Every post body must have at least 10 characters

The previous version of this project: [v0.1.0](https://github.com/EstebanBorai/hkroom/releases/tag/v0.1.0).

Was one of the first Golang projects I've ever made, I decided to
update the project to follow better practices and to keep the
repository clean

## Run Locally

```bash
go run cmd/main.go
```

This project is created by cloning the template from [gosk](https://github.com/EstebanBorai/gosk).

## cURL Request Examples

### List Posts

```bash
curl http://localhost:8080
```

### Find Post

```bash
curl http://localhost:8080/:id
```

### Create a Post

> Posts must have at least 5 characters for the title and 10 characters for the body

```bash
curl -d '{"title": "Hello World!", "body": "This is a test on the Post Creation"}' \
  -H "Content-Type: application/json" \
  -X POST http://localhost:8080
```

### Update a Post

> Posts must have at least 5 characters for the title and 10 characters for the body

```bash
curl -d '{"title": "This is an updated post!", "body": "Lorem Ipsum for the updated post"}' \
  -H "Content-Type: application/json" \
  -X PUT http://localhost:8080/:id
```

### Delete a Post

```bash
curl -X DELETE http://localhost:8080/:id
```
