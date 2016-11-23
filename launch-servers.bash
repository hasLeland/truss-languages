#!/bin/bash


function start() {
	nohup swedish-server -http.addr :5051 -grpc.addr :0 -debug.addr :0 >swedish-server.out 2>&1 &
	canadian-server -http.addr :5052 -grpc.addr :0 -debug.addr :0 >canadian-server.out 2>&1 &
	gateway-server >gateway-server.out 2>&1 &
}

function stop() {
	pkill swedish-server
	pkill canadian-server
	pkill gateway-server
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
	start)
		start
		;;
	stop)
		stop
		;;
	restart)
		restart
		;;
	*)
		echo "Usage: $0 start|stop|restart|rebuild"
		exit 1
esac
