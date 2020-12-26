package isodd

// check int8 whether is odd
func Int8(i int8) bool {
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

// check int16 whether is odd
func Int16(i int16) bool {
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

// check int32 whether is odd
func Int32(i int32) bool {
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

// check int64 whether is odd
func Int64(i int64) bool {
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

// check int whether is odd
func Int(i int) bool {
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

// check uint8 whether is odd
func Uint8(i uint8) bool {
	return i%2 == oddCheckNum
}

// check uint16 whether is odd
func Uint16(i uint16) bool {
	return i%2 == oddCheckNum
}

// check uint32 whether is odd
func Uint32(i uint32) bool {
	return i%2 == 1
}

// check uint64 whether is odd
func Uint64(i uint64) bool {
	return i%2 == oddCheckNum
}

// check uint whether is odd
func Uint(i uint) bool {
	return i%2 == oddCheckNum
}
