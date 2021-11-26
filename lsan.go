package lsan

// See https://clang.llvm.org/docs/AddressSanitizer.html. You may need to run with env
// ASAN_OPTIONS=detect_leaks=1 . On MacOS you may need to specify a
// custom clang for address sanitizor support (https://stackoverflow.com/a/55778432/149482). There's
// an interface defined at
// https://chromium.googlesource.com/chromiumos/third_party/compiler-rt/+/release_34/include/sanitizer/lsan_interface.h
// . I can't work out what the proper include is.

// #cgo CPPFLAGS: -O0 -g -fsanitize=address
// #cgo LDFLAGS: -fsanitize=address -g
// #include <stdio.h>
// #include <stdlib.h>
//
// void __lsan_do_leak_check(void);
//
// void leak_a_bit(void)
// {
//     char *p = malloc(2000);
//     printf("allocated leak at %p\n", p);
// }
import "C"

func DoLeakCheck() {
	// Apparently this should only be called once per process run.
	C.__lsan_do_leak_check()
}

func LeakABit() {
	C.leak_a_bit()
}
