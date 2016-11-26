package main

// cgo_1 START OMIT

/*
#cgo LDFLAGS: -lglfw
#include <stdlib.h>
#include <GLFW/glfw3.h>

extern void myErrorCallback(int, char*);
static inline void myErrorCallback_cgo(int code, const char* desc) {
	myErrorCallback(code, (char*)desc);
}
static void setCallback() {
	glfwSetErrorCallback(myErrorCallback_cgo);
}
*/
import "C"
import (
	"log"
	"time"
	"unsafe"
)

//export myErrorCallback
func myErrorCallback(code C.int, desc *C.char) {
	log.Fatalln("got glfw error:", code, C.GoString(desc))
}

// cgo_1 END OMIT

// cgo_2 START OMIT

func cgo_1() {
	C.setCallback()

	log.Println("initializing")
	if ret := C.glfwInit(); ret != 1 {
		log.Fatalln("error:", ret)
	}
	defer C.glfwTerminate()

	name := C.CString("My window")
	defer C.free(unsafe.Pointer(name))
	log.Println("creating window")
	window := C.glfwCreateWindow(640, 480, name, nil, nil)
	if window == nil {
		log.Fatalln("error creating window")

	}
	time.Sleep(5 * time.Second)
}

// cgo_2 END OMIT
func main() {
	cgo_1()
}
