# Kuma2Awtrix

This project is a bridge between Kuma and Awtrix, written in Go.

## Health Check

The application has a /health endpoint for monitoring. A GET request to `/health`
will return a `200` status code and a message indicating that the application is
healthy.
