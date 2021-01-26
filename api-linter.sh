#!/bin/sh

bazel build //api/...

api-linter --config api-linter.yaml \
  -I bazel-"$(basename "$(pwd)")"/external/com_google_googleapis \
  api/*.proto
