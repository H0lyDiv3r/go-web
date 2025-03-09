package internals

import (
	"fmt"
	"go-fe-fwk/types"
	"syscall/js"
)

type Reducers map[string]types.ReducerFunc

type App struct {
	State    map[string]any                                                                       `json:"state"`
	View     func(state map[string]any, emit func(eventName string, payload map[string]any)) Vdom `json:"view"`
	Reducers Reducers                                                                             `json:"reducers"`
}

func (app *App) CreateApp() map[string]func(parentEl js.Value) {
	fmt.Println("CREATING THIS APP")
	var parentEl js.Value
	var vdom Vdom = nil

	var dispatcher Dispatcher = Dispatcher{Subs: map[string][]types.ReducerFunc{}, AfterHandlers: []types.AnyFunc{}}
	emit := func(eventName string, payload map[string]any) {
		dispatcher.Dispatch(eventName, payload)
	}
	renderApp := func(payload ...any) any {
		if vdom != nil {
			DestroyDom(vdom)
		}
		fmt.Println("RE RENDERING THE APP")
		vdom = app.View(app.State, emit)
		MountDom(vdom, parentEl)
		return nil
	}
	afterHandler := dispatcher.AfterEveryCommand(renderApp)
	subscriptions := []func(){afterHandler}
	for actionName, reducer := range app.Reducers {
		subs := dispatcher.Subscribe(actionName, func(state map[string]any, payload map[string]any) map[string]any {
			app.State = reducer(app.State, payload)
			return nil
		})
		subscriptions = append(subscriptions, subs)
		fmt.Println("subscribing the functions.", dispatcher.Subs)
	}

	return map[string]func(payload js.Value){
		"mount": func(_parentEl js.Value) {
			parentEl = _parentEl
			renderApp()
			fmt.Println("mouting")
		},
		"unmount": func(_parentEl js.Value) {
			fmt.Println("dismounting")
		},
	}

}

//     return {
//         mount(_parentEl){
//             parentEl=_parentEl
//             console.log("aaaaaaaaaaaaaaaaaaaaaaaa",_parentEl)
//             renderApp()
//         },
//     unmount(){
//         destroyDom(vdom),
//         vdom=null,
//         subscriptions.forEach((unsubscribe)=>unsubscribe())
//     }
//     }

// export const createApp = ({state,view,reducers={}})=>{
//     let parentEl = null
//     let vdom = null
//     const emit = (eventName,payload)=>{
//         dispatcher.dispatch(eventName, payload)
//     }
//     const renderApp = () => {
//             if (vdom) {
//                 destroyDom(vdom)
//             }
//             vdom = view(state,emit)
//         console.log("vdom", vdom);
//             mountDom(vdom, parentEl)
//         };
//     const dispatcher = new Dispatcher()
//     const subscriptions = [dispatcher.afterEveryCommand(renderApp)]

//     for(const actionName in reducers){
//         const reducer = reducers[actionName]
//         const subs=dispatcher.subscribe(actionName, (payload)=>{
//             state=reducer(state,payload)
//         })
//         subscriptions.push(subs)
//     }

//     return {
//         mount(_parentEl){
//             parentEl=_parentEl
//             console.log("aaaaaaaaaaaaaaaaaaaaaaaa",_parentEl)
//             renderApp()
//         },
//     unmount(){
//         destroyDom(vdom),
//         vdom=null,
//         subscriptions.forEach((unsubscribe)=>unsubscribe())
//     }
//     }
// }
