#!/bin/bash
set -euo pipefail

if ! command -v asdf >/dev/null
then
  echo >&2 'error: you need to install asdf-vm. Follow steps 1-3 here:

    https://asdf-vm.com/guide/getting-started.html#getting-started
'
  exit 1
fi

required_asdf_plugins=(
  golang
)

for p in "${required_asdf_plugins[@]}"
do
  set +e
  asdf plugin add "$p"
  case $? in
    0 | 2) break;;
    *) exit 1;;
  esac
  set -e
done

asdf install

npm install @openapitools/openapi-generator-cli
