# SMART exporter for Prometheus

Tips: [English version is here.](README_EN.md)

这个项目源自  https://grafana.com/grafana/dashboards/10664   的 [BASH 脚本](https://github.com/janw/node-exporter-textfile-collector-scripts/blob/master/smartmon.sh) 。但其有部分的 Bug 没有处理，同时写与 2012 年所以并不支持如 NVMe 等较新的硬件，因此改造增强了这个脚本并使用 Golang 封装以支持 HTTP 输出方便部署。

## 特性

* 完全使用 Goalng 编写
* 在 macOS、Linux 以及 FreeBSD 下经过完全的测试
* 支持 NVMe 磁盘
* 简单的 HTTP 服务器（并增加了缓存功能）
* 将 BASH 脚本内置增加安全性和方便部署
* 当然，如果不想使用二进制文件，直接执行 `scripts/smartmon.sh`

## 编译

编译相对简单，安装完 Golang 环境以后，直接使用 `make` 即可。

## 部署


首先在你的系统上安装  `smartmontools`  这个软件包，然后执行 `smart-exporter` 即可，无需任何的配置。访问如下地址：

* `http://<machine-address>:9111/metrics` // metrics details for prometheus
* `http://<machine-address>:9111/script` // dump shell script
* `http://<machine-address>:9111/version` // version and build information

对应的 `systemd` 和 `supervisor` 有相应的启动脚本供参考。如有任何的意见和建议，欢迎邮件联系，非常感谢。

`- eof-`
