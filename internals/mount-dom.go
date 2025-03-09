package internals

import (
	"go-fe-fwk/pkgs/utils"
	"go-fe-fwk/types"
	"syscall/js"
)

type Vdom interface {
	CreateNode(parentEl js.Value)
	RemoveNode()
}

type ElementNode struct {
	Tag       string             `json:"tag"`
	Props     types.Props        `json:"props"`
	Children  []Vdom             `json:"children"`
	Type      string             `json:"type"`
	Listeners map[string]js.Func `json:"listeners"`
	El        js.Value           `json:"el"`
}

func (vdom *ElementNode) CreateNode(parentEl js.Value) {
	document := js.Global().Get("document")
	tag := vdom.Tag
	props := vdom.Props
	children := vdom.Children
	element := document.Call("createElement", tag)
	addProps(element, props, vdom)
	vdom.El = element
	for _, child := range children {
		MountDom(child, element)
	}
	parentEl.Call("appendChild", element)
}

func (vdom *ElementNode) RemoveNode() {
	// fmt.Println("aaaaaaaaaaaa destroying el", vdom)
	element := vdom.El
	children := vdom.Children
	listeners := vdom.Listeners
	element.Set("className", "awaaww")
	element.Call("remove")
	for _, child := range children {
		DestroyDom(child)
	}
	// fmt.Println("Removing this shit")
	utils.RemoveEventListener(listeners, element)
	vdom.Listeners = nil
}

type StringDom struct {
	Type  string   `json:"type"`
	Value string   `json:"value"`
	El    js.Value `json:"el"`
}

func (vdom *StringDom) CreateNode(parentEl js.Value) {
	document := js.Global().Get("document")
	value := vdom.Value
	textNode := document.Call("createTextNode", value)
	vdom.El = textNode
	parentEl.Call("appendChild", textNode)
}

func (vdom *StringDom) RemoveNode() {
	// fmt.Println("aaaaaaaaaaaaaaaaa destroying", vdom)
	element := vdom.El
	element.Call("remove")
}

func MountDom(vdom Vdom, parentEl js.Value) {
	// fmt.Println("printing the pointer", vdom)
	vdom.CreateNode(parentEl)
}

func addProps(element js.Value, props types.Props, vdom *ElementNode) {
	attributes := props.Attributes
	eventHandlers := props.On
	vdom.Listeners = utils.AddEventListeners(eventHandlers, element)
	utils.SetAttributes(element, attributes)
}
