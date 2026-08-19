# https://just.systems

default:
    echo 'Hello, world!'

build:
    go build .

up: build
    sudo ./own-vpn up

down: build
    sudo ./own-vpn down

run neighbour: build
    sudo ./own-vpn run --neighbours {{neighbour}}