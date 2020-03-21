#!/usr/bin/env bash
###
# File: test.sh
# Author: Ming Cheng<mingcheng@outlook.com>
#
# Created Date: Saturday, March 21st 2020, 6:45:10 pm
# Last Modified: Saturday, March 21st 2020, 6:52:44 pm
#
# http://www.opensource.org/licenses/MIT
###

if [ ! -x ./smart-exporter ]; then
  echo "smart-exporter binary file is not exists or execuable."
  exit -1
fi

if [ "$(./smart-exporter --script | md5sum)" = "$(cat ./scripts/smartmon.sh | md5sum)" ]; then
  echo "smart-exporter buildin script is equal to ./scripts/smartmon.sh"
  exit 0
fi

exit -1
