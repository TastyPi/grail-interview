#!/bin/sh

bazel build //api/proto/...

api-linter --config api-linter.yaml \
  -I bazel-"$(basename "$(pwd)")"/external/com_google_googleapis \
  api/proto/*.proto
