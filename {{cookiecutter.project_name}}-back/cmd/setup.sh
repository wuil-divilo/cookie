#!/bin/bash

if [ ! -d .make-out/ ]; then
  mkdir .make-out
fi

GOREPORTCARD=$(which goreportcard-cli)
if [ -z "${GOREPORTCARD}" -o ! -x "${GOREPORTCARD}" ]; then
  echo "Installing goreportcard-cli ..."
  git clone https://github.com/gojp/goreportcard.git /tmp/goreportcard
  cd /tmp/goreportcard
  make install
  go install ./cmd/goreportcard-cli
  cd - || return
  rm -rf /tmp/goreportcard
  echo "Done: Installing goreportcard-cli"
fi

OAPICODEGEN=$(which oapi-codegen)
if [ -z "${OAPICODEGEN}" -o ! -x "${OAPICODEGEN}" ]; then
  echo "Installing oapi-codegen ..."
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.1
  echo "Done: Installing goreportcard-cli"
fi

if [ -d .git/hooks/ ] && ([ ! -f .git/hooks/pre-commit ] || [ $(diff -q .githooks/pre-commit .git/hooks/pre-commit | wc -l) -eq 1 ]); then
  echo "Installing pre-commit hook ..."
  cp .githooks/pre-commit .git/hooks/
  echo "Done: Installing pre-commit hook"
fi
