#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

# shellcheck source=../.env
source "$DIR/../.env" || {
  echo "error while sourcing env file"
  exit
}

cd "$DIR/.." || exit

FULL_PSQL_IMAGE_NAME="${DOCKER_REPOSITORY}:${PSQL_DEV_IMAGE_TAG}"

if ! docker build -t "${FULL_PSQL_IMAGE_NAME}" -f "$DIR/Dockerfile_Dev_PSQL" .; then
  echo "Could not build image ${FULL_PSQL_IMAGE_NAME}"
  exit 1
fi

PS3="Push image ${FULL_PSQL_IMAGE_NAME}: "

select opt in yes no; do

  case $opt in
  yes)
    docker push "${FULL_PSQL_IMAGE_NAME}"
    break
    ;;
  *)
    echo "NO PUSH"
    break
    ;;
  esac
done
