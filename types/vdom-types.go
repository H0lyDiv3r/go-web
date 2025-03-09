package types

import "syscall/js"

type Attributes[T string | []string] map[string]T
type JsFunc func(this js.Value, args []js.Value) any
type AnyFunc func(payload ...any) any
type ReducerFunc func(state map[string]any, payload map[string]any) map[string]any
type EventHandlers map[string]func(this js.Value, args []js.Value) interface{}
type Props struct {
	Attributes Attributes[string] `json:"attribute"`
	On         EventHandlers      `json:"on"`
}

var DOM_TYPES = map[string]string{
	"TEXT":     "text",
	"ELEMENT":  "element",
	"FRAGMENT": "fragment",
}
