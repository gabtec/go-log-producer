# Go Logs Producer App

```sh
docker run --rm -p 4000:4000 IMAGE_NAME:TAG
```

## Endpoints

- GET http://localhost:4000/demo
- on each new request a random response (success or error), will be generated

## Images

```sh
docker pull ghcr.io/gabtec/go-log-producer:latest
```

## Changelog

### v0.1.1

- single endpoint "/demo", that produces a single, random, log
