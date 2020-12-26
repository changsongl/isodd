package isodd

import "errors"

const (
	oddCheckNum    = 1
	oddCheckNumNeg = -1
)

var (
	ErrorInterfaceNoMatch = errors.New("isodd.Interface: interface not match")
)
