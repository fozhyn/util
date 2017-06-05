package convert

import "unsafe"

//convert string to []byte without mem alloc and copy
//confirm the life cycle of the background array by yourself
//appropriate use
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	y := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&y))
}

//convert []byte to string without mem alloc and copy
//confirm the life cycle of the background array by yourself
//appropriate use
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
