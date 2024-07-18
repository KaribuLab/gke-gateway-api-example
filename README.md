# GKE Gateway API Example

## Application

This simple application use *Echo Framework* and *Docker* to deploy.

Run locally this application:

```shell
go mod download
go run main.go
```

Test the application with `curl`:

```shell
curl http://localhost:1323
```

This return always 401 HTTP status code to simulate a private resource. To get a 200 HTTP status code use the following resource.

```shell
curl http://localhost:1323/health
```