#!/usr/bin/env bash

set -ex

# Exit code always 0
golangci-lint run --out-format=line-number \
| reviewdog -f=golangci-lint -diff="git diff ${DEFAULT_BRANCH}..HEAD" -reporter=github-pr-review