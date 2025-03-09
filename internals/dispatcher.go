package internals

import (
	"fmt"
	"go-fe-fwk/pkgs/utils"
	"go-fe-fwk/types"
)

type Dispatcher struct {
	Subs          map[string][]types.ReducerFunc `json:"subs"`
	AfterHandlers []types.AnyFunc                `json:"afterHandlers"`
}

func (dispatcher *Dispatcher) Subscribe(commandName string, handler types.ReducerFunc) func() {
	if _, ok := dispatcher.Subs[commandName]; !ok {
		dispatcher.Subs[commandName] = []types.ReducerFunc{}
	}
	handlers := dispatcher.Subs[commandName]
	if utils.ReducerFuncExists(handlers, handler) {
		return func() {}
	}
	handlers = append(handlers, handler)
	dispatcher.Subs[commandName] = handlers
	return func() {
		idx := utils.IndexOfReducerFunction(handlers, handler)
		handlers = append(handlers[:idx], handlers[idx+1:]...)
	}
}

func (dispatcher *Dispatcher) AfterEveryCommand(handler types.AnyFunc) func() {
	dispatcher.AfterHandlers = append(dispatcher.AfterHandlers, handler)
	return func() {
		idx := utils.IndexOfFunction(dispatcher.AfterHandlers, handler)
		if idx != -1 {
			dispatcher.AfterHandlers = append(dispatcher.AfterHandlers[:idx], dispatcher.AfterHandlers[idx+1:]...)
		}
	}
}

func (dispatcher *Dispatcher) Dispatch(commandName string, payload map[string]any) {
	if _, ok := dispatcher.Subs[commandName]; ok {
		fmt.Println("dispatching", commandName, dispatcher.Subs[commandName])

		for _, handler := range dispatcher.Subs[commandName] {
			fmt.Println("dispatching inside loop", commandName, payload)
			handler(nil, payload)
		}
	} else {
		fmt.Println("there is no command by that name")
	}
	fmt.Println("printing after every command man", dispatcher.AfterHandlers)
	for _, handler := range dispatcher.AfterHandlers {
		handler()
	}
}

//     dispatch(commandName, payload){
//         if(this.#subs.has(commandName)){
//             this.#subs.get(commandName).forEach((handler)=>{
//                 handler(payload)
//             })
//         }else{
//             console.log(`there is no command by the name ${commandName}`)
//         }

//         this.#afterHandlers.forEach((handler)=>{handler()})
//     }
// export class Dispatcher {
//     #subs = new Map()
//     #afterHandlers =[]
//     subscribe(commandName, handler){
//         if(!this.#subs.has(commandName)){
//             this.#subs.set(commandName, [])
//         }
//         const handlers = this.#subs.get(commandName)
//         if(handlers.includes(handler)){
//             return ()=>{}
//         }
//         handlers.push(handler)
//         return ()=>{
//             const idx = handlers.indexOf(handler)
//             handlers.splice(idx,1)
//         }
//     }
//     afterEveryCommand(handler) {
//         this.#afterHandlers.push(handler)
//         return () =>{
//             const idx= this.#afterHandlers.indexOf(handler)
//             this.#afterHandlers.splice(idx,1)
//         }
//     }

//     dispatch(commandName, payload){
//         if(this.#subs.has(commandName)){
//             this.#subs.get(commandName).forEach((handler)=>{
//                 handler(payload)
//             })
//         }else{
//             console.log(`there is no command by the name ${commandName}`)
//         }

//         this.#afterHandlers.forEach((handler)=>{handler()})
//     }

// }
