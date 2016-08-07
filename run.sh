#!/bin/bash

while true
do
	echo "Start"
	$(go run gograph.go)
	ls -lh database.json
	echo "Loop"
	sleep 10
done

