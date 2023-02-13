package main

import "os"

func main() {
	check()
	os.Remove("list.txt")

}
