package lsan

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

func LeakABit() {
	C.leak_a_bit()
}
