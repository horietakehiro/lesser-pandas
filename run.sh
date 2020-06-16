#!/bin/bash

docker run -d \
        --hostname lesser-pandas \
        --name lesser-pandas \
        -v ${PWD}/codes:/codes \
        -v ${PWD}/root:/root \
        lesser-pandas:latest
