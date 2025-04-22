# internal-health-check-api

A simple and very light weight health check api with golang

## Get started

Create a env, then run this app via docker or directly `go run src/main.go`

## ENV Pattern

this app required these env to response:

 - `{SERVICE_NAME}_HEALTH_HOST`
 - `{SERVICE_NAME}_HEALTH_PORT`
 - `{SERVICE_NAME}_HEALTH_TYPE` // TCP or UDP

## Example

```env
WEB_HEALTH_HOST="192.168.1.1"
WEB_HEALTH_PORT="80"
WEB_HEALTH_TYPE="TCP"
```