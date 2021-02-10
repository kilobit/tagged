/* Copyright 2020 Kilobit Labs Inc. */

// Example of embedding an XML Object
//
package main

import "fmt"
import _ "errors"
import "encoding/xml"

import x "kilobit.ca/go/tagged"

var zepxml string = `
 <Band>
  <Vocals>Robert Plant</Vocals>
  <Bass>John Paul Jones</Bass>
  <Guitar>Jimmy Page</Guitar>
  <Drums>John Bonham</Drums>
 </Band>
`

type Band struct {
	x.XMLElement
}

func(b *Band) Vocals() string {

	result := ""

	el := b.GetByTagName("Vocals")
	if el != nil {
		result = el.GetContent()
	}

	return result
}

func(b *Band) Guitar() string {

	result := ""

	el := b.GetByTagName("Guitar")
	if el != nil {
		result = el.GetContent()
	}

	return result
}

func(b *Band) Bass() string {

	result := ""

	el := b.GetByTagName("Bass")
	if el != nil {
		result = el.GetContent()
	}

	return result
}

func(b *Band) Drums() string {

	result := ""

	el := b.GetByTagName("Drums")
	if el != nil {
		result = el.GetContent()
	}

	return result
}

func main() {

	zep := Band{}
	err := xml.Unmarshal([]byte(zepxml), &zep)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"Vocals:\t%s\nGuitar:\t%s\nBass:\t%s\nDrums:\t%s\n",
		zep.Vocals(),
		zep.Guitar(),
		zep.Bass(),
		zep.Drums(),
	)
}
