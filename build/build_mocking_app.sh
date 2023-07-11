#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
# shellcheck source=../.env
source "$DIR/../.env" || {
  echo "error while sourcing env file"
  exit
}

# shellcheck source=../.env.pw
source "$DIR/../.env.pw" || {
  echo "error while sourcing env.pw file"
  exit
}

cd "$DIR/.." || exit
FULL_MOCKING_IMAGE_NAME="${DOCKER_REPOSITORY}:${MOCKING_TAG}"
DOCKERFILE_NAME="Dockerfile_Mocking_App"
DOCKERFILE_PATH="${DIR}/${DOCKERFILE_NAME}"

if ! docker build --build-arg VERSION="${OWN_VERSION}" --build-arg DEPLOY_USER="${REYMOM_DEPLOY_USER}" --build-arg DEPLOY_PW="${REYMOM_DEPLOY_PW}" -t "${FULL_MOCKING_IMAGE_NAME}" --target executor -f "${DOCKERFILE_PATH}" .; then
  echo "Could not build image ${FULL_MOCKING_IMAGE_NAME}"
  exit 1
fi


PS3="Push image ${FULL_MOCKING_IMAGE_NAME}?: "
select opt in yes no; do

  case $opt in
  yes)
    docker push "${FULL_MOCKING_IMAGE_NAME}"
    break
    ;;
  *)
    echo "NO PUSH"
    break
    ;;
  esac
done