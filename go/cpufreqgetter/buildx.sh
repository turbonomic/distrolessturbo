#!/usr/bin/env bash
set -Eeuo pipefail
set -x
REPO_NAME=icr.io/cpopen/turbonomic
IMAGE_NAME=cpufreqgetter
PLATFORM_OS_ARCH_LIST="linux/amd64,linux/arm64,linux/ppc64le,linux/s390x"
docker buildx build --platform $PLATFORM_OS_ARCH_LIST --label "git-version=latest" --push -t $REPO_NAME/$IMAGE_NAME:latest .
