package isodd

func Int8(i int8) bool{
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

func Int16(i int16) bool{
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

func Int32(i int32) bool{
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

func Int64(i int64) bool{
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

func Int(i int) bool{
	return i%2 == oddCheckNum || i%2 == oddCheckNumNeg
}

func Uint8(i uint8) bool{
	return i%2 == oddCheckNum
}

func Uint16(i uint16) bool{
	return i%2 == oddCheckNum
}

func Uint32(i uint32) bool{
	return i%2 == 1
}

func Uint64(i uint64) bool{
	return i%2 == oddCheckNum
}

func Uint(i uint) bool{
	return i%2 == oddCheckNum
}


