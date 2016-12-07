#!/bin/bash

while true
do
	echo "Start"
	go run *.go
	ls -lh *.db
	echo "Loop"
	sleep 1
done

