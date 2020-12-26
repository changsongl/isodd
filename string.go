package isodd

import (
	"github.com/pkg/errors"
	"strconv"
)

func String(s string) (bool, error) {
	i, e := strconv.ParseInt(s, 10, 64)
	if e != nil {
		return false, wrapStringErr(e)
	}

	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg, nil
}

func wrapStringErr(e error) error {
	return errors.Wrap(e, "isodd.String")
}
