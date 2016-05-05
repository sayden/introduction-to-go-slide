package main

import (
	"fmt"
	"reflect"
)

func main() {
	me := "Jebediah Kerman"
	fmt.Println(reflect.TypeOf(me))
}
