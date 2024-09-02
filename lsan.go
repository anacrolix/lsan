//go:build lsan

// See https://clang.llvm.org/docs/AddressSanitizer.html. You may need to run with env
// ASAN_OPTIONS=detect_leaks=1,halt_on_error=0. Since Go 1.18 you might also need
// CGO_CPPFLAGS_ALLOW='.*' and CGO_LDFLAGS_ALLOW='.*' for the -fsanitize-recover. On MacOS you may
// need to specify a custom clang for address sanitizor support
// (https://stackoverflow.com/a/55778432/149482). There's an interface defined at
// https://chromium.googlesource.com/chromiumos/third_party/compiler-rt/+/release_34/include/sanitizer/lsan_interface.h
// . I can't work out what the proper include is.
package lsan

// #cgo CPPFLAGS: -O0 -g -fsanitize=address -fsanitize-recover=address
// #cgo LDFLAGS: -fsanitize=address -g -fsanitize-recover=address
import "C"

func DoLeakCheckIfLsanEnabled() {
	DoLeakCheck()
}

func DoLeakCheck() {
	// Apparently this should only be called once per process run.
	C.__lsan_do_leak_check()
}
