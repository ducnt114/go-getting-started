package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("NumCPU: ", runtime.NumCPU())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	lastMaxProcs := runtime.GOMAXPROCS(2)
	fmt.Println("lastMaxProcs: ", lastMaxProcs)
	//fmt.Println("lastMaxProcs: ", runtime.LockOSThread)

	//fd, err := unix.Open("./test.txt", os.O_CREATE|os.O_WRONLY, 0600)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = unix.Write(fd, []byte("hello world\n"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = unix.Close(fd)
	//if err != nil {
	//	panic(err)
	//}
}
