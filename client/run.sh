#!/bin/bash

while true
do
	echo "Start"
	go run *.go
	ls -lh database.*
	echo "Loop"
	sleep 1
done

