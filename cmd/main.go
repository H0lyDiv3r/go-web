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

type task struct {
	idx   int
	title string
	done  bool
}

func mountWithApp(this js.Value, args []js.Value) interface{} {
	var app internals.App
	app.State = map[string]any{
		"tmp":     "",
		"editing": -1,
		"tasks":   []task{},
	}
	app.Reducers = internals.Reducers{
		"add": func(state map[string]any, payload map[string]any) map[string]any {
			state["tasks"] = append(state["tasks"].([]task), task{title: payload["title"].(string), done: payload["done"].(bool), idx: payload["idx"].(int)})
			fmt.Println("emmitting from adddddddd babat", state, payload)
			return state
		},
		"setString": func(state, payload map[string]any) map[string]any {
			fmt.Println("tell me why", payload["value"], state["tmp"])
			state["tmp"] = payload["value"].(string)
			return state
		},
		"edit": func(state, payload map[string]any) map[string]any {
			idx := payload["idx"].(int)
			state["tasks"].([]task)[idx].title = payload["newVal"].(string)
			return state
		},
		"delete": func(state, payload map[string]any) map[string]any {
			idx := payload["idx"].(int)
			state["tasks"] = append(state["tasks"].([]task)[0:idx], state["tasks"].([]task)[idx+1:]...)
			return state
		},
		"setEdit": func(state, payload map[string]any) map[string]any {
			state["editing"] = payload["value"].(int)
			return state
		},
	}
	app.View = func(state map[string]any, emit func(eventName string, payload map[string]any)) internals.Vdom {
		buttonText := "Add Task"
		editButtonText := "edit"
		tmpVal := ""
		if state["editing"] != -1 {
			tmpVal = state["tasks"].([]task)[state["editing"].(int)].title
			buttonText = "update value"
			editButtonText = "cancel"
		}
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
		//
		tasks := []any{}

		for idx, tsk := range state["tasks"].([]task) {
			var task *internals.ElementNode

			task = internals.H("div", types.Props{
				Attributes: types.Attributes[string]{
					"class": "task",
				},
			}, []any{

				internals.H("div", types.Props{}, []any{
					internals.H("p", types.Props{}, []any{tsk.title}),
				}),
				internals.H("div", types.Props{}, []any{

					internals.H("button", types.Props{
						On: types.EventHandlers{
							"click": func(this js.Value, args []js.Value) interface{} {
								if state["editing"].(int) == -1 {
									emit("setEdit", map[string]any{"value": tsk.idx})
								} else {
									emit("setEdit", map[string]any{"value": -1})
								}
								fmt.Println("editing", tsk.idx)
								return nil
							},
						},
					}, []any{editButtonText}),
					internals.H("button", types.Props{
						Attributes: types.Attributes[string]{},
						On: types.EventHandlers{
							"click": func(this js.Value, args []js.Value) interface{} {
								fmt.Println("deleting", tsk.idx)
								if state["editing"] != -1 {
									fmt.Println("cant delete in edit mode")
									return nil
								}
								emit("delete", map[string]any{
									"idx": idx,
								})
								return nil
							},
						},
					}, []any{"delete"}),
				}),
			})

			tasks = append(tasks, internals.H("div", types.Props{}, []any{
				task,
			}))
		}

		// if state["editing"].(int) == -1 {
		// 	state["tmp"] = state["tmp"].(string)
		// } else {
		// 	state["tmp"] = state["tasks"].([]task)[state["editing"].(int)].title
		// }

		return internals.H("div", types.Props{
			Attributes: types.Attributes[string]{
				"class": "container",
			},
		}, []any{
			internals.H("div", types.Props{}, tasks),
			//input
			internals.H("input", types.Props{
				Attributes: types.Attributes[string]{
					"value":       tmpVal,
					"placeholder": "title",
					"class":       "input",
				},
				On: types.EventHandlers{
					"input": func(this js.Value, args []js.Value) interface{} {
						tmpVal = args[0].Get("target").Get("value").String()
						// emit("setString", map[string]any{
						// 	"value": args[0].Get("target").Get("value").String(),
						// })
						return nil
					},
				},
			}, []any{}),

			//button
			internals.H("button", types.Props{
				Attributes: types.Attributes[string]{
					"class": "colored",
				},
				On: types.EventHandlers{
					"click": func(this js.Value, args []js.Value) interface{} {
						if state["editing"].(int) != -1 {
							emit("edit", map[string]any{"newVal": tmpVal, "idx": state["editing"].(int)})
							emit("setEdit", map[string]any{"value": -1})
							emit("setString", map[string]any{
								"value": "",
							})
							return nil
						}
						emit("add", map[string]any{"title": tmpVal, "done": false, "idx": len(state["tasks"].([]task))})
						emit("setString", map[string]any{
							"value": "",
						})
						return nil
					},
				},
			}, []any{buttonText}),
			// internals.H("img", types.Props{
			// 	Attributes: types.Attributes[string]{
			// 		"src": "Screenshot 2025-03-09 at 14-18-58 Instagram.png",
			// 	},
			// }, []any{}),
		})
		// return internals.H("button", types.Props{
		// 	Attributes: types.Attributes[string]{
		// 		"class": "colored",
		// 	},
		// 	On: types.EventHandlers{
		// 		"click": func(this js.Value, args []js.Value) interface{} {
		// 			fmt.Println("button clicked")
		// 			emit("add", map[string]any{"count": 5})
		// 			return nil
		// 		},
		// 	},
		// }, []any{stringState})
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
