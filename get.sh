#!/bin/bash
# set -x
set -eo pipefail

GREEN='\033[0;32m'
NC='\033[0m'
RED='\033[0;31m'
INSTALL_PATH="/usr/local/bin"

if ! [ -x "$(command -v tar)" ]; then
  echo 'Error: curl is not installed.' >&2
  exit 1
fi

# Check for root permission
touch ${INSTALL_PATH}/.hashnode &> /dev/null || (echo -e "${RED}Root access is required to install to ${GREEEN}${INSTALL_PATH}${NC}" && exit 1)
rm ${INSTALL_PATH}/.hashnode

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

