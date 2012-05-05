
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
			if len(buf) > 0 {
				fmt := C.CString(string(buf))
				n += int(C.scanverbatim(fmt))
				C.free(unsafe.Pointer(fmt))
			}
			break
		}

		buf = append(buf, '%')
		if i++; i == len(format) {
			panic("input.Scanf: format string ends with '%'")
		} else if format[i] == '%' {
			buf = append(buf, '%')
		} else {
			spec := format[i]
			if spec != 'd' && spec != 'c' {
				panic("input.Scanf: only '%d' and '%c' format specifiers allowed")
			}

			if arg == len(a) {
				panic("input.Scanf: the number of format specifiers exeeds the number of arguments")
			}

			buf = append(buf, spec)
			fmt := C.CString(string(buf))
			var res C.int
			n += int(C.scanint(fmt, &res))
			C.free(unsafe.Pointer(fmt))
			buf = buf[:0]

			switch p := a[arg].(type) {
			case nil:
				panic("input.Scanf: nil pointer passed as argument")
			case *uint: *p = uint(res)
			case *uint8: *p = uint8(res)
			case *uint16: *p = uint16(res)
			case *uint32: *p = uint32(res)
			case *uint64: *p = uint64(res)
			case *int: *p = int(res)
			case *int8: *p = int8(res)
			case *int16: *p = int16(res)
			case *int32: *p = int32(res)
			case *int64: *p = int64(res)
			default:
				panic("input.Scanf: argument must be pointer to some integer type")
			}
			arg++
		}
		i++
	}

	return
}
