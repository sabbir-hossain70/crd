#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")

vendor/k8s.io/code-generator/generate-groups.sh all \
github.com/sabbir-hossain70/crd/pkg/client \
github.com/sabbir-hossain70/crd/pkg/apis \
crd.com:v1 \
--go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt