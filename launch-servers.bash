#!/bin/bash

servers=("swedish-server" "canadian-server" "gateway-server")

function status() {
	for sname in "${servers[@]}"; do
		if [ -z "$(pgrep $sname)" ]; then
			printf "%20s stopped\n" $sname
		else
			printf "%20s running\n" $sname
		fi
	done
}

function start() {
	nohup swedish-server -http.addr :5051 -grpc.addr :0 -debug.addr :0 >swedish-server.out 2>&1 &
	nohup canadian-server -http.addr :5052 -grpc.addr :0 -debug.addr :0 >canadian-server.out 2>&1 &
	nohup gateway-server >gateway-server.out 2>&1 &
}

function stop() {
	for sname in "${servers[@]}"; do
		pkill "$sname"
	done
}

function restart() {
	stop
	start
}

function rebuild() {
	go install ./...
	restart
}

case $1 in
	status)
		status
		;;
	start)
		start
		;;
	stop)
		stop
		;;
	restart)
		restart
		;;
	rebuild)
		rebuild
		;;
	*)
		echo "Usage: $0 status|start|stop|restart|rebuild"
		exit 1
esac
