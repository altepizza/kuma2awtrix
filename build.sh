#!/bin/bash

# Build and push the image to GitHub Container Registry

docker buildx build --platform=linux/amd64 -t ghcr.io/altepizza/kuma2awtrix:latest .
docker push ghcr.io/altepizza/kuma2awtrix:latest
