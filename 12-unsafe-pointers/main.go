package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func state() {
	var onload = createEvents("onload")
	var receive = createEvents("receive")
	var success = createEvents("success")
	mapEvents := make(map[string]interface{})
	mapEvents["messageOnload"] = unsafe.Pointer(onload)
	mapEvents["messageReceive"] = unsafe.Pointer(receive)
	mapEvents["messageSuccess"] = uintptr(unsafe.Pointer(success))
	//This line is safe - retains original value
	fmt.Println(*(*string)(mapEvents["messageReceive"].(unsafe.Pointer)))
	//This line is unsafe - original value could be garbage collected
	fmt.Println(*(*string)(unsafe.Pointer(mapEvents["messageSuccess"].(uintptr))))

	runtime.KeepAlive(success)
}
func createEvents(s string) *string {
	return &s
}

func main() {
	state()
}
