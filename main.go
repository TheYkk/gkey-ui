package main

import (
	"fmt"
	"github.com/theykk/gkey"
	"syscall/js"
)

func getElementByID(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func genPP(this js.Value, inputs []js.Value) interface{} {
	realm := getElementByID("realm")
	password := getElementByID("password")
	secure := getElementByID("secure")

	pass,err :=gkey.GenPass(realm.Get("value").String(),password.Get("value").String(),16)
	if err !=nil {
		fmt.Print(err)
	}
	secure.Set("value",pass)

	return nil
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("genPP", js.FuncOf(genPP))
	<-make(chan bool)
}

