package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

// From: https://github.com/gernest/mention
func getRuneBytes(r rune) []byte {
	rst := make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(rst, r)
	return rst
}

/* I want to understand the package 'mention' indicated in the above url.
 * It does make sense now, if you want to be able to parse
 * any variable/tag using unicode, you must use runes.
 */

func main() {
	/* Let's create a variable 'фskip_global/, has the special separator
	*  Using getRuneBytes we get the byte representation.
	*   'ф',  [209 132]
	*  For instance the Korean hangul has 3 bytes length
	*  '각', [234 176 129]
	*  However '$', [36]
	 */

	runeToByte := getRuneBytes('$') //ф
	var stringByte []byte
	data := []byte("$skip_global")
	fmt.Println("data:", data)
	fmt.Printf("runeToByte=%d\n", runeToByte)

	/* Where is located the byte representation of that rune in this []byte data slice?
	*  bytes.Index returns the position for any runeBytes representation. I want to read just after
	*  the rune. Here is why we have to use runes, because the byte representation
	*  of any rune may have a length bigger than 1 bytes.
	 */

	index := bytes.Index(data, runeToByte)
	fmt.Println("Index ф= ", index)

	// Position of the byte rune representation + length of the bytes
	begin := index + len(runeToByte)

	for i := begin; i < len(data); i++ {
		stringByte = append(stringByte, data[i])
	}

	fmt.Printf("stringByte=%s\n", string(stringByte))
}
