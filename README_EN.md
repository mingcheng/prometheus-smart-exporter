# S.M.A.R.T export for Prometheus

This project is inspire from https://grafana.com/grafana/dashboards/10664 .
That Grafana dashboard provide a simple [BASH script](https://github.com/janw/node-exporter-textfile-collector-scripts/blob/master/smartmon.sh) which
I found have some buggy and can not running on FreeBSD. So, I forked and modified this script and add some necessary features.

## Feature

* Written in #Golang
* Full tested on Linux and FreeBSD (Sorry for Windows)
* NVMe driver supported
* Simply metrics HTTP supported(with cache for performance issue)
* The BASH script is builtin, simplify to deploy
* If you don't wanna binary #Golang, just run `scripts/smartmon.sh`

## Building

Building is fairly straight, After install golang package and `make` it, that's all :-P

## Deploy and Usage

You need to install `smartmontools` package on your system first. Then run  `smart-exporter` binary with root privilege (smartctl need it). visit and check:

* `http://<machine-address>:9111/metrics` // metrics details for prometheus
* `http://<machine-address>:9111/script` // dump shell script
* `http://<machine-address>:9111/version` // version and build information

and there are simple startup scripts in `systemd` or `supervisor` directory, if you are using `systemd`(on most GNU/Linux distortion) or `supervisord`.

That's it! If your have any suggestion, contact me by mail.

`- eof -`
