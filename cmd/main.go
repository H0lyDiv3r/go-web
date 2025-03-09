package main

import (
	"fmt"
	"go-fe-fwk/internals"
	"go-fe-fwk/types"
	"syscall/js"
)

var props = types.Props{
	Attributes: types.Attributes[string]{
		"class": "colored",
	},
	On: types.EventHandlers{
		"click": func(this js.Value, args []js.Value) interface{} {
			fmt.Println("aaaaaaaaaaaaaaaaaaaaaaauuuuuuuu")
			return nil
		},
	},
}
var vdom = internals.H("div", props, []any{
	"aaaaaaaaa",
	internals.H("p", types.Props{}, []any{"awawaw"}),
})

func mountDom(this js.Value, args []js.Value) interface{} {
	vdm := internals.H("button", props, []any{"aaa"})
	// l := []any{
	// 	internals.H("ul", types.Props{}, []any{
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 		internals.H("li", types.Props{}, []any{"one"}),
	// 	}),
	// }

	// vdom.CreateNode(args[0])
	internals.MountDom(vdm, args[0])
	fmt.Println(vdom)

	// time.Sleep(2 * time.Second)
	// // fmt.Println(args[0], vdom)
	// internals.DestroyDom(vdom)
	return nil
}

func remDom(this js.Value, args []js.Value) interface{} {
	vdom.RemoveNode()
	return nil
}

func mountWithApp(this js.Value, args []js.Value) interface{} {
	var app internals.App
	app.State = map[string]any{
		"count": 0,
	}
	app.Reducers = internals.Reducers{
		"add": func(state map[string]any, payload map[string]any) map[string]any {
			// state := payload[0]
			// val := payload[1]
			state["count"] = state["count"].(int) + payload["count"].(int)
			fmt.Println("emmitting from adddddddd babat", state, payload)
			return state
		},
	}
	app.View = func(state map[string]any, emit func(eventName string, payload map[string]any)) internals.Vdom {

		stringState := fmt.Sprint(state["count"])
		fmt.Println("i am coming from here with", string(stringState))
		// input := internals.H("input", types.Props{
		// 	Attributes: types.Attributes[string]{
		// 		"placeholder": "this is a placeholder",
		// 	},
		// 	On: types.EventHandlers{
		// 		"change": func(this js.Value, args []js.Value) interface{} {
		// 			fmt.Println("changed", this.Get("value"), args)
		// 			return nil
		// 		},
		// 	},
		// }, []any{})
		// return input
		return internals.H("button", types.Props{
			Attributes: types.Attributes[string]{
				"class": "colored",
			},
			On: types.EventHandlers{
				"click": func(this js.Value, args []js.Value) interface{} {
					fmt.Println("button clicked")
					emit("add", map[string]any{"count": 5})
					return nil
				},
			},
		}, []any{stringState})
	}
	// emit := func(eventName string, payload any) {
	// 	fmt.Printf("Event emitted: %s with payload: %v\n", eventName, payload)
	// }
	// vdm := app.View(app.State, emit)
	// internals.MountDom(vdm, args[0])
	p := app.CreateApp()
	p["mount"](args[0])
	return nil
}

func main() {

	fmt.Println("wasm connected")
	// js.Global().Set("mountDom", js.FuncOf(mountDom))
	// js.Global().Set("removeDom", js.FuncOf(remDom))

	js.Global().Set("mountDom", js.FuncOf(mountWithApp))
	<-make(chan struct{})
	// props := types.Props{
	// 	Attributes: types.Attributes[string]{
	// 		"aa": "aa",
	// 	},
	// 	On: types.EventHandlers{},
	// }

	// vdom := []any{
	// 	internals.H("p", props, []any{"paragraph"}),
	// }
	// jsonBytes, err := json.Marshal(internals.H("div", props, vdom))

	// if err != nil {
	// 	fmt.Println("there has been an error mate")
	// }

	// fmt.Println(string(jsonBytes))

}
