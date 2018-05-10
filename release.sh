#!/usr/bin/env bash

version=0.1
rm hotspots-${version}.tar.gz

# frontend build
# cd frontend && npm run build && cd ..
# sed -i "s#height:88px#min-height:88px; height:auto!important; height:88px;#g" frontend/dist/static/css/*.css

# go build
export GIN_MODE=release
go build -o hotspots main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hotspots-linux-amd64 main.go

# tar
sleep 2
tar zcvf hotspots-${version}.tar.gz hotspots* frontend/dist data