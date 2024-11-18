#! /bin/bash

# STEP 1: Determinate the required values

PACKAGE="go-getting-started"
VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')

# STEP 2: Build the ldflags

LDFLAGS=(
  "-X '${PACKAGE}/goredis/conf.Version=${VERSION}'"
  "-X '${PACKAGE}/goredis/conf.CommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}/goredis/conf.BuildTimestamp=${BUILD_TIMESTAMP}'"
)

# STEP 3: Actual Go build process

go build -ldflags="${LDFLAGS[*]}" ./goredis/main.go