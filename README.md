# Kuma2Awtrix

This project is a bridge between [Uptime Kuma](https://github.com/louislam/uptime-kuma)
and [AWTRIX 3](https://github.com/Blueforcer/awtrix3), written in Go.

## Health Check

The application has a /health endpoint for monitoring. A GET request to `/health`
will return a `200` status code and a message indicating that the application is
healthy.
