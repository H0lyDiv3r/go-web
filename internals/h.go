package internals

import (
	"fmt"
	"go-fe-fwk/types"
	"reflect"
)

func H(tag string, props types.Props, children []any) *ElementNode {
	newChildren := []Vdom{}
	for i, child := range children {
		if reflect.TypeOf(children[i]) == reflect.TypeOf("") {
			textNode := Hstring(child.(string))
			newChildren = append(newChildren, &textNode)
		} else {
			fmt.Println("showing children", child)
			newChildren = append(newChildren, child.(*ElementNode))
		}
	}
	return &ElementNode{
		Tag:      tag,
		Props:    props,
		Children: newChildren,
		Type:     types.DOM_TYPES["ELEMENT"],
	}
}

func Hstring(value string) StringDom {
	return StringDom{
		Type:  types.DOM_TYPES["TEXT"],
		Value: value,
	}
}
