package main

import (
	"fmt"
	"os"

	widevine "github.com/iyear/gowidevine"
)

var token string
var etpRt string
var keys []*widevine.Key
var debugLog *os.File

func init() {
	debugLog, _ = os.OpenFile("/tmp/crdl.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
}

func logf(format string, args ...any) {
	if debugLog != nil {
		fmt.Fprintf(debugLog, format, args...)
	}
}
