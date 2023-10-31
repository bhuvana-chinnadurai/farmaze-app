#!/bin/bash

go get -tags 'postgres' -u github.com/golang-migrate/migrate/v4/cmd/migrate
migrate -path=/migrations/ -database=postgres://farmaze:farmaze@db:5432/farmaze?sslmode=disable up

