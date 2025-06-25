package main

import (
	"fmt"

	"github.com/qiushi1511/hellolib/greet"
)

func main() {
	fmt.Println(greet.Greet("Qiushi"))
	fmt.Println(greet.WarmGreet("Qiushi"))
}
