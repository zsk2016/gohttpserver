#!/bin/bash
echo "please input build option: 0-[winlocal],1-[linux],2-[linux-a]"
read host_idx

if [ ${host_idx} = 0 ]; then
  export CGO_ENABLED=0
	export GOARCH=386
	export GOOS=windows

  echo "install main"
  go install -a gohttpserver/main
fi
if [ ${host_idx} = 1 ]; then
  export CGO_ENABLED=0
	export GOARCH=386
	export GOOS=linux

  echo "install main"
  go install gohttpserver/main
fi