# S.M.A.R.T export for Prometheus

This project is inspire from https://grafana.com/grafana/dashboards/10664 .
That Grafana dashboard provide a simple [BASH script](https://github.com/janw/node-exporter-textfile-collector-scripts/blob/master/smartmon.sh) which
I found have some buggy and can not running on FreeBSD. So, I forked and modified this script and add some necessary features.

## Feature

* Written in #Golang
* Full tested on Linux and FreeBSD
* Simply metrics HTTP supported(with cache for performance issue)
* The BASH script is builtin, simplify to deploy
* If you don't wanna binary #Golang, just run `scripts/smartmon.sh`

## Building

Building is fairly straight, Just `make` it, that's all :-P

## Deploy and Usage

@TODO

`- eof -`
