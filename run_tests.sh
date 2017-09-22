#!/bin/bash

if [ "${TRAVIS_PULL_REQUEST}" = "false" ]; then
    make container
    docker login -u $docker_user -p $docker_pass

    docker tag multidex reaxun/multidex:$TRAVIS_BUILD_NUMBER
    docker push reaxun/multidex:$TRAVIS_BUILD_NUMBER

    docker tag multidex reaxun/multidex:latest
    docker push reaxun/multidex:latest
else
    make test
    make coverage
fi
