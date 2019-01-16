#!/bin/bash
# set -x
set -eo pipefail
#!/bin/bash

GREEN='\033[0;32m'
NC='\033[0m'
RED='\033[0;31m'

if [[ $EUID -ne 0 ]]; then
  echo -e "${RED}This script must be run as root, please try again with sudo${NC}" 
   exit 1
fi

if ! [ -x "$(command -v tar)" ]; then
  echo 'Error: curl is not installed.' >&2
  exit 1
fi

BINARY_NAME="hashnode"
HOST_OS=${HOST_OS:-$(uname | tr '[:upper:]' '[:lower:]')}

if [[ $(uname -m) == "x86_64" ]]; then
  HOST_ARCH="amd64"
else
  HOST_ARCH=${HOST_ARCH:-$(uname -m)}
fi

ARTIFACT_NAME=${BINARY_NAME}-${HOST_OS}-${HOST_ARCH}.tar.gz

LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/hashnode/hashnode-cli/releases/latest)
LATEST_VERSION=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
ARTIFACT_URL="https://github.com/hashnode/hashnode-cli/releases/download/$LATEST_VERSION/$ARTIFACT_NAME"

curl -L $ARTIFACT_URL | tar xvz
mv hashnode /usr/local/bin/hashnode


echo -e "${GREEN}Installed Successfully${NC}"

