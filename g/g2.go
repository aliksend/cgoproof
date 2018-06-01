package g

/*
#include "./c-api/greet.h"
*/
import "C"

func Greet2() string {
	return C.GoString(C.greet2())
}
