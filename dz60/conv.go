package main

import "fmt"
import "os"
import "log"
import "encoding/json"

// KeyCode holds layers by Key
type KeyCode struct {
	ID string `json:"id"`
}

// Key holds key data
type Key struct {
	ID       int `json:"id"`
	Row      int `json:"row"`
	Col      int `json:"col"`
	KeyCodes []KeyCode
}

// Keyboard Contains Array of Keys
type Keyboard struct {
	Keys    []Key `json:"keys"`
	Rows    int   `json:"rows"`
	Columns int   `json:"cols"`
}

// KeyboardConfig contains a Keyboard config
type KeyboardConfig struct {
	Version  int      `json:"version"`
	Keyboard Keyboard `json:"keyboard"`
}

func main() {
	fh, err := os.Open("dz60-model42.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(fh)
	doc := KeyboardConfig{}
	if err := dec.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	fmt.Println(`#include "dz60.h"`)
	fmt.Println(`#include "action_layer.h"`)
	fmt.Printf("const uint16_t PROGMEM keymaps[][%d][%d] = {\n", doc.Keyboard.Rows, doc.Keyboard.Columns)

	layers := [4]string{}
	for _, v := range doc.Keyboard.Keys {
		for j := 0; j < 4; j++ {
			layers[j] = fmt.Sprintf("%s %s,", layers[j], v.KeyCodes[j].ID)
		}
	}

	for i, layer := range layers {
		fmt.Printf("// layer %d\n", i)
		fmt.Print("KEYMAP(")
		fmt.Printf("%s", layer)
		fmt.Print("),\n\n")
	}
	fmt.Println("}")

}
