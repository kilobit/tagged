TaggEd
======

Work simply with XML backed data structures in Golang.

Status: In Development

```
import x "kilobit.ca/go/tagged"

type Band struct {
	Vocals string
	Bass   string
	Guitar string
	Drums  string
}

func main() {

	zep := Band{"Robert Plant", "John Paul Jones", "Jimmy Page", "John Bonham"}

	bs, _ := xml.MarshalIndent(zep, "", " ")

	xe := tagged.XMLElement{}
	xml.Unmarshal(bs, &xe)

	xgtr := xe.GetChildByName("guitar")

	cstr := xgtr.GetContent()

	fmt.Println(cstr)
}
```


```
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

func main() {

	zep := Band{}
	err := xml.Unmarshal([]byte(zepxml), &zep)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Vocals:\t%s\n",	zep.Vocals())
}
```

Features
--------

- Unmarshal arbitrary XML.
- Work with XMLElement directly or embed it in your type.
- Walk trees of tag data.
- Marshal back to XML.

Installation
------------

```
go get kilobit.ca/go/tagged
go test -v
```

Building
--------

```
go get kilobit.ca/go/tagged
go build

```

Contribute
----------

Please help!  Submit pull requests through
[Github](https://github.com/kilobit/tagged).

Support
-------

Please submit issues through
[Github](https://github.com/kilobit/tagged).

License
-------

See LICENSE.

--  
Created: Feb 10, 2021  
By: Christian Saunders <cps@kilobit.ca>  
Copyright 2021 Kilobit Labs Inc.  
