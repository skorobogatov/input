
package input

// #include <stdlib.h>
// #include "input.h"
//
// #cgo CFLAGS: -O3
//
import "C"

import "unsafe"

func Scanf(format string, a ...interface{}) (n int) {
	var buf []byte
	i, arg := 0, 0
	for {
		for ; i < len(format) && format[i] != '%'; i++ {
			buf = append(buf, format[i])
		}

		if i == len(format) {
			fmt := C.CString(string(buf))
			n += int(C.scanverbatim(fmt))
			C.free(unsafe.Pointer(fmt))
			break
		}

		buf = append(buf, '%')
		switch i++; format[i] {
		case '%': buf = append(buf, '%')
		case 'd':
			if arg == len(a) {
				panic("1 !!!")
			} else if p, ok := a[arg].(*int); !ok {
				panic("2 !!!")
			} else {
				buf = append(buf, 'd')
				fmt := C.CString(string(buf))
				var res C.int
				n += int(C.scanint(fmt, &res))
				C.free(unsafe.Pointer(fmt))
				*p = int(res)
				arg++
				buf = buf[:0]
			}
		}

		i++
	}

	return
}
