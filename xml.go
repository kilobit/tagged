/* Copyright 2020 Kilobit Labs Inc. */

// Work simply with XML backed data structures in Golang.
//
package tagged

import "fmt"
import _ "errors"
import "io"
import "strings"
import "encoding/xml"

// https://play.golang.org/p/iD6pnVu9Sin

// Base type of XML encoded structures.  Use directly or embedd
// XMLElement in your own types.
//
// Methods approximate the interface provided by a DOM.
//
type XMLElement struct {
	name     string
	ns       string
	attrs    map[string]string
	children []interface{}
}

func (el XMLElement) Attr(name string) (string, bool) {

	val, ok := el.attrs[name]
	return val, ok
}

func (el XMLElement) NChildren() int {
	return len(el.children)
}

func (el XMLElement) NthChild(i int) interface{} {

	if len(el.children) < i-1 {
		return nil
	}

	return el.children[i]
}

func (el *XMLElement) GetByTagName(tag string) *XMLElement {

	for _, child := range el.children {

		xe, ok := child.(*XMLElement)
		if !ok {
			continue
		}

		if xe.name == tag {
			return xe
		}

		result := xe.GetByTagName(tag)
		if result != nil {
			return result
		}
	}

	return nil
}

func (el *XMLElement) GetChildByName(name string) *XMLElement {

	name = strings.ToLower(strings.TrimSpace(name))

	for _, child := range el.children {

		xe, ok := child.(*XMLElement)
		if !ok {
			continue
		}

		if strings.ToLower(strings.TrimSpace(xe.name)) == name {
			return xe
		}
	}

	return nil
}

func (el *XMLElement) GetChildrenByName(name string) []*XMLElement {

	name = strings.ToLower(strings.TrimSpace(name))
	results := []*XMLElement{}

	for _, child := range el.children {

		xe, ok := child.(*XMLElement)
		if !ok {
			continue
		}

		if strings.ToLower(strings.TrimSpace(xe.name)) == name {
			results = append(results, xe)
		}
	}

	return results
}

func (el *XMLElement) GetCharData() []string {

	results := []string{}

	for _, child := range el.children {

		cd, ok := child.(xml.CharData)
		if ok {
			results = append(results, string(cd))
		}
	}

	return results
}

// Returns all of the CharData as a single concatenated string.
//
func (el *XMLElement) GetContent() string {

	results := el.GetCharData()

	return strings.Join(results, "")
}

func (el XMLElement) String() string {

	sb := &strings.Builder{}
	e := xml.NewEncoder(sb)
	e.Indent("", " ")

	err := el.MarshalXML(e, xml.StartElement{})
	if err != nil {
		return fmt.Sprintf("Failed to marshal xml: %s", err.Error())
	}

	return sb.String()
}

func (el *XMLElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	s := xml.StartElement{}

	s.Name.Local = el.name
	s.Name.Space = el.ns

	for name, value := range el.attrs {
		s.Attr = append(s.Attr, xml.Attr{xml.Name{"", name}, value})
	}

	//fmt.Printf("%#v\n", s)

	err := e.EncodeToken(s)
	if err != nil {
		return err
	}

	for _, child := range el.children {

		switch val := child.(type) {

		case xml.CharData, xml.Comment:
			err = e.EncodeToken(val)
			if err != nil {
				return err
			}

		default:
			err = e.Encode(child)
			if err != nil {
				return err
			}
		}
	}

	err = e.EncodeToken(xml.EndElement{xml.Name{Local: el.name, Space: el.ns}})
	if err != nil {
		return err
	}

	return e.Flush()
}

//
// Note: This parser ignores comments, processing instructions and directives.
//
// Todo: Store namespaces and pass them to child tags, then replace
// xmlns attributes with the names.  e.g. podcast:transcript
//
func (el *XMLElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	el.name = start.Name.Local
	el.ns = start.Name.Space
	el.attrs = map[string]string{}

	for _, attr := range start.Attr {
		//fmt.Printf("%#v\n", attr)
		el.attrs[attr.Name.Local] = attr.Value
	}

	done := false
	for {

		if done {
			break
		}

		tkn, err := d.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch val := tkn.(type) {

		case xml.StartElement:
			child := &XMLElement{}
			err = child.UnmarshalXML(d, val)
			if err != io.EOF && err != nil {
				return err
			}

			el.children = append(el.children, child)

		case xml.CharData:
			if strings.TrimSpace(string(val)) != "" {
				el.children = append(el.children, val.Copy())
			}

		case xml.Comment:
			el.children = append(el.children, string(val.Copy()))

		case xml.EndElement:
			done = true
			break

		default:
			continue
		}
	}

	return nil
}
