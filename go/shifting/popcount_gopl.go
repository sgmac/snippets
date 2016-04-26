package main

import (
	"flag"
	"fmt"
)

/* So the idea is to count how many bits are set to 1
* in an unsigned 64 bit number.

A 8 bit (byte) number has 256 posible values and a 64bit  number is 8 x 8bit number.
Compute a table with all the 256 posible values, counting the bits set to 1.

Then return the sum of all the bits enabled for the 8 groups of 8 bits.
For instance the number 4096:

7byte	    6byt      5byte       4byte     3byte      2byte      1byte	     0byte
00000000 +00000000 + 00000000 + 00000000 + 00000000 + 00000000 + 00001000 + 00000000

pc[byte(4096>>(0*8)]= 0 there is not right shifting, so 0 positions are moved.
pc[byte(4096>>(1*8)]=  Now we have to move 8 times to the right

00001000 00000000 (Shift 1 pos)
00000100 00000000 (Shift 2 pos)
00000001 00000000 (Shift 3 pos)
00000000 10000000 (Shift 4 pos)
00000000 01000000 (Shift 5 pos)
00000000 00100000 (Shift 6 pos)
00000000 00010000 (Shift 7 pos)
00000000 00001000 (Shift 8 pos)

So the number is 16, this is the index used in the "pc" table to lookup the number of bits set to 1.
Now if the number was 4099 (4096 +3) that should be pc[16] + pc[3] lookup

The idea is to right shifting the N-byte from the unsigned 64bit so you can see the equivalent in the lookup table.

*/

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		// fmt.Printf("pc[%d]=%d\n", i, pc[i])
	}
}

var pc [256]byte

func PopCount(x uint64) int {
	fmt.Printf("pc[byte(x>>(0*8))]=%d\n", pc[byte(x>>(0*8))])
	fmt.Printf("pc[byte(x>>(1*8))]=%d\n", pc[byte(x>>(1*8))])
	fmt.Printf("pc[byte(x>>(2*8))]=%d\n", pc[byte(x>>(2*8))])
	fmt.Printf("pc[byte(x>>(3*8))]=%d\n", pc[byte(x>>(3*8))])
	fmt.Printf("pc[byte(x>>(4*8))]=%d\n", pc[byte(x>>(4*8))])
	fmt.Printf("pc[byte(x>>(5*8))]=%d\n", pc[byte(x>>(5*8))])
	fmt.Printf("pc[byte(x>>(6*8))]=%d\n", pc[byte(x>>(6*8))])
	fmt.Printf("pc[byte(x>>(7*8))]=%d\n", pc[byte(x>>(7*8))])
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
func main() {
	number := flag.Uint64("n", 0, "Number to count bits.")
	flag.Parse()

	fmt.Printf("Number=%d %b\n", *number, *number)
	fmt.Printf("PopCount=%d\n", PopCount(*number))
	fmt.Printf("%d>>0*8=%d\n", *number, byte(*number>>(0*8)))
	fmt.Printf("%d>>1*8=%d\n", *number, byte(*number>>(1*8)))
	fmt.Printf("%d>>2*8=%d\n", *number, byte(*number>>(2*8)))
	fmt.Printf("%d>>2*8=%d\n", *number, byte(*number>>(3*8)))
	fmt.Printf("%d>>3*8=%d\n", *number, byte(*number>>(4*8)))
	fmt.Printf("%d>>4*8=%d\n", *number, byte(*number>>(5*8)))
	fmt.Printf("%d>>5*8=%d\n", *number, byte(*number>>(6*8)))
	fmt.Printf("%d>>6*8=%d\n", *number, byte(*number>>(7*8)))
}
