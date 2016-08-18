#!/bin/bash

while true
do
	echo "Start"
	go run *.go
	ls -lh database.json
	echo "Loop"
	sleep 5
done

