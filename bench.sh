#!/bin/bash

autocannon -m POST -H 'Content-Type: application/json' -i body-node.json -c 100 -d 60 http://localhost:8080/users

autocannon -m POST -H 'Content-Type: application/json' -i body-go.json -c 100 -d 60 http://localhost:3000/users
