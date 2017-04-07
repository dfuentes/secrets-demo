#!/usr/bin/env bash

export ENVIRONMENT=development
export SERVICE=test-service

$(./bootstrap)

env | grep _SECRET
