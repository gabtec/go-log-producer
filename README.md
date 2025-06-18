# Go Logs Producer App

```sh
docker run --rm -p 4000:4000 IMAGE_NAME:TAG
```

## Endpoints

- baseURL: http://localhost:4000

| method | endpoint                   | description                                                  |
| ------ | -------------------------- | ------------------------------------------------------------ |
| GET    | http://_baseURL_/          | list all available endpoints                                 |
| GET    | http://_baseURL_/log/:type | logs a specific response type (type can be "ok", or "error") |
| GET    | http://_baseURL_/random    | logs random response, either OK or ERROR                     |
| GET    | http://_baseURL_/demo      | (deprecated) use /random                                     |

## Images

```sh
docker pull ghcr.io/gabtec/go-log-producer:latest

# container versions, match release version
v1.0.0 (latest) - uses new endpoints
v0.1.1          - single /demo endpont
```

## Changelog

### v1.0.0

- adds more endpoints
  - /
  - /logs/:type , type = "ok" or "error"
  - /random
  - /demo , will be deprecated

### v0.1.1

- single endpoint "/demo", that produces a single, random, log
