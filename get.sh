#!/usr/bin/env bash
set -eo pipefail

release=$(curl --silent "https://api.github.com/repos/hashnode/hashnode-cli/releases/latest" | sed -n 's/.*"tag_name": *"\([^"]*\)".*/\1/p')
