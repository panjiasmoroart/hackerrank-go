package main

import "fmt"

func main() {
	staircase(4)
}

func staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		// loop untuk spasi (sama seperti $j=4; $j>=$i)
		for j := n; j > i; j-- {
			// satu spasi seperti &nbsp;
			fmt.Print(" ")
		}

		// loop untuk cetak bintang
		for k := int32(1); k <= i; k++ {
			fmt.Print("#")
		}

		fmt.Println()
	}

	/*
	      #
	     ##
	    ###
	   ####
	*/
}

func segitiga(n int32) {
	for i := int32(1); i <= n; i++ {
		// loop untuk spasi (sama seperti $j=4; $j>=$i)
		for j := n; j > i; j-- {
			// satu spasi seperti &nbsp;
			fmt.Print(" ")
		}

		// loop untuk cetak bintang
		for k := int32(1); k <= i; k++ {
			fmt.Print("#")
		}

		for l := int32(1); l <= i; l++ {
			fmt.Print("#")
		}

		fmt.Println()
	}

	/*
		   ##
		  ####
		 ######
		########
	*/
}
