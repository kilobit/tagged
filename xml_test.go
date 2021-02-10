/* Copyright 2020 Kilobit Labs Inc. */

package tagged_test

import "strings"
import "encoding/xml"

import "testing"
import "kilobit.ca/go/tested/assert"

import "kilobit.ca/go/tagged"

func TestXMLTest(t *testing.T) {
	assert.Expect(t, true, true)
}

type Band struct {
	Vocals string
	Bass   string
	Guitar string
	Drums  string
}

func TestAttr(t *testing.T) {

	xstr := `<rss xmlns="rss.org/rss.dtd" xmlns:example="http://foo.com/" x="why"><example:foo bing="boom"/></rss>`
	xe := tagged.XMLElement{}
	err := xml.Unmarshal([]byte(xstr), &xe)
	assert.Ok(t, err)

	val, ok := xe.Attr("x")
	assert.Expect(t, true, ok)

	assert.Expect(t, "why", val)
}

func TestGetChildByName(t *testing.T) {

	zep := Band{"Robert Plant", "John Paul Jones", "Jimmy Page", "John Bonham"}

	bs, err := xml.MarshalIndent(zep, "", " ")
	assert.Ok(t, err)

	xe := tagged.XMLElement{}
	err = xml.Unmarshal(bs, &xe)
	assert.Ok(t, err)

	xgtr := xe.GetChildByName("guitar")
	assert.Expect(t, true, xgtr != nil)

	cd := xgtr.GetCharData()
	assert.Expect(t, 1, len(cd))
	assert.Expect(t, "Jimmy Page", cd[0])
}

func TestGetChildrenByName(t *testing.T) {

	xstr := `<rss><item d="one" /><item d="two" /><item d="three" /></rss>`
	xe := tagged.XMLElement{}
	err := xml.Unmarshal([]byte(xstr), &xe)
	assert.Ok(t, err)

	// Non-existent
	items := xe.GetChildrenByName("foo")
	assert.Expect(t, 0, len(items))

	// Existent
	items = xe.GetChildrenByName("item")
	assert.Expect(t, 3, len(items))

	val, ok := items[2].Attr("d")
	assert.Expect(t, true, ok)
	assert.Expect(t, "three", val)

}

func TestGetTagByName(t *testing.T) {

	zep := Band{"Robert Plant", "John Paul Jones", "Jimmy Page", "John Bonham"}

	bs, err := xml.MarshalIndent(zep, "", " ")
	assert.Ok(t, err)

	obj := tagged.XMLElement{}
	err = xml.Unmarshal(bs, &obj)
	assert.Ok(t, err)

	act := obj.GetByTagName("Guitar")
	assert.Expect(t, true, act != nil)

	strs := act.GetCharData()
	assert.Expect(t, "Jimmy Page", strs[0])
}

func TestNSTags(t *testing.T) {

	xstr := `<rss xmlns="rss.org/rss.dtd" xmlns:example="http://foo.com/" x="why"><example:foo bing="boom"/></rss>`
	obj := tagged.XMLElement{}
	err := xml.Unmarshal([]byte(xstr), &obj)
	assert.Ok(t, err)

	// t.Logf("%#v\n", obj)
	// t.Logf("%#v\n", obj.NthChild(0))
	// t.Log(obj.String())

}

func TestXMLElementString(t *testing.T) {

	zep := Band{"Robert Plant", "John Paul Jones", "Jimmy Page", "John Bonham"}

	bs, err := xml.MarshalIndent(zep, "", " ")
	assert.Ok(t, err)

	obj := tagged.XMLElement{}
	err = xml.Unmarshal(bs, &obj)
	assert.Ok(t, err)

	//t.Logf("%s", obj)

	assert.Expect(t, string(bs), obj.String())
}

func TestUnmarshalMarshalXML(t *testing.T) {

	zep := Band{"Robert Plant", "John Paul Jones", "Jimmy Page", "John Bonham"}

	bs, err := xml.MarshalIndent(zep, "", " ")
	assert.Ok(t, err)

	obj := tagged.XMLElement{}
	err = xml.Unmarshal(bs, &obj)
	assert.Ok(t, err)

	sb := &strings.Builder{}
	e := xml.NewEncoder(sb)
	e.Indent("", " ")
	err = obj.MarshalXML(e, xml.StartElement{})
	assert.Ok(t, err)

	assert.Expect(t, string(bs), sb.String())
}
