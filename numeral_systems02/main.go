package main

import "fmt"

func main() {
	fmt.Println("the number 42 in various numbering systems :")
	fmt.Printf("%d decimal\n", 42)
	fmt.Printf("%b binary\n", 42)
	fmt.Printf("%o octal\n", 42)
	fmt.Printf("%#o octal leading 0 \n", 42)
	fmt.Printf("%x hexadecimal (lower case)\n", 42)
	fmt.Printf("%X hexadecimal (upper case)\n", 42)
	fmt.Printf("%#x hexadecimal (lower case) leading 0x\n", 42)
	fmt.Printf("%#X hexadecimal (upper case) leading 0x\n", 42)

	fmt.Println("ASCII upper case alphabets : ")
	for i := 65; i <= 90; i++ {
		fmt.Printf("%c ", i)

	}
	fmt.Println()

	fmt.Println("ASCII lower case alphabets : ")
	for i := 97; i <= 122; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()

	fmt.Println("First 100 digits in dec, bin, hex and utf-8")
	fmt.Println("\ndec\thex\toct\tbin\tutf-8")
	for i := 0; i <= 50; i++ {
		fmt.Printf("%d\t%#X\t%#o\t%b\t%q\n", i, i, i, i, i)
	}
	fmt.Println()

}
