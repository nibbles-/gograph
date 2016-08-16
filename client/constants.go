package main

import (
	"regexp"
)

var regSip = regexp.MustCompile("^.*Cisco SIP\(.*\)\\CallsInProgress$")
var regMgcpGw = regexp.MustCompile("^.*Cisco MGCP Gateways\(.*\)\\PRIChannelsActive$")
var regMgcpPri = regexp.MustCompile("^.*Cisco MGCP PRI Device\(.*\)\\CallsActive$")
