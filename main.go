package main

import (
	"fmt"

	"github.com/qiushi1511/hellolib/greet"
	hellolib "github.com/qiushi1511/hellolib/v2"
)

func main() {
	fmt.Println(greet.Greet("Qiushi"))
	fmt.Println(greet.WarmGreet("Qiushi"))

	hellolib.Greeting()
}
