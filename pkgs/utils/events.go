package utils

import (
	"go-fe-fwk/types"
	"syscall/js"
)

func AddEventListener(eventName string, handler func(this js.Value, args []js.Value) any, element js.Value) js.Func {
	element.Call("addEventListener", eventName, js.FuncOf(handler))
	return js.FuncOf(handler)
}

func AddEventListeners(listeners types.EventHandlers, element js.Value) map[string]js.Func {
	addedEventListeners := make(map[string]js.Func)

	for eventName := range listeners {
		listener := AddEventListener(eventName, listeners[eventName], element)
		addedEventListeners[eventName] = listener
	}
	return addedEventListeners
}

func RemoveEventListener(listeners map[string]js.Func, element js.Value) {
	for eventName := range listeners {
		element.Call("removeEventListener", eventName, listeners[eventName])
	}
}

// const addEventListener = (eventName,handler,el) => {
//     el.addEventListener(eventName,handler)
//     return handler
// }

// export const addEventListeners = (listeners={},el)=>{
//     const addedListeners = {}
//     Object.entries(listeners).forEach(([eventName,handler])=>{
//         const listener = addEventListener(eventName,handler,el)
//         addedListeners[eventName] = listener
//     })
//     return addedListeners
// }
// export const removeEventListeners = (listeners,el)=>{
//     Object.entries(listeners).forEach(([eventName,handler])=>{
//         el.removeEventListener(eventName,handler)
//     })
// }
