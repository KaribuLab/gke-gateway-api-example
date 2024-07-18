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

## Deployment in GKE

```shell
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
# Before to run this command, create a global IP in your Google Cloud project named "gateway-api-ip-test"
kubectl apply -f gateway.yaml
kubectl apply -f http-route.yaml
kubectl apply -f healthcheck-policy.yaml
```