#!/bin/bash

cd pkg/graphql || return
go get github.com/99designs/gqlgen@v0.17.43
go run github.com/99designs/gqlgen
cd ../..