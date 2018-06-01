package g

/*
#include "./greet.h"
*/
import "C"

func Greet1() string {
	return C.GoString(C.greet1())
}
