package utils

import (
	"go-fe-fwk/types"
	"strings"
	"syscall/js"
)

func SetAttributes(element js.Value, attributes types.Attributes[string]) {
	otherAttrs := make(map[string]interface{})
	for key, value := range attributes {
		if key != "class" && key != "style" {
			otherAttrs[key] = value
		}
	}
	if className, ok := attributes["class"]; ok {
		setClass(element, className)
	}
	if styleName, ok := attributes["style"]; ok {
		setStyle(element, styleName, "")
	}
	for key, value := range otherAttrs {
		setAttribute(element, key, value.(string))
	}
}
func setClass[T string | []string](element js.Value, className T) {
	switch any(className).(type) {
	case string:
		element.Set("className", "")
		element.Set("className", className)
	case []string:
		for _, class := range any(className).([]string) {
			element.Get("classList").Set("length", 0)
			element.Get("classList").Call("add", class)
		}
	}
}
func setStyle(element js.Value, name, value string) {
	element.Get("style").Set(name, value)
}

func removeStyle(element js.Value, name string) {
	element.Get("style").Set(name, js.Null())
}

func setAttribute(element js.Value, name string, value string) {
	// if value == "" {
	// 	removeAttribute(element, name)
	// }
	if strings.HasPrefix(name, "data-") {
		element.Call("setAttribute", name, value)
	} else {
		element.Set(name, value)
	}

}

func removeAttribute(element js.Value, name string) {
	element.Set(name, js.Null)
	element.Call("removeAttribute", name)
}

// const setAttribute = (el,name,value)=>{
//     if(value==null){
//         removeAttribute(el,name)
//     }else if(name.startsWith('data-')){
//         el.setAttribute(name,value)
//     }else{
//         el[name]=value
//     }
// }

// const setClass = (el,className) => {
//     el.className = ''
//     if(typeof className === 'string') {
//         el.className = className
//     }
//     if(Array.isArray(className)){
//         el.classList.add(...className)
//     }
// }

// export const setAttributes = (el,attrs) =>{
//     const { class:className,style,...otherAttrs} = attrs
//     if(className){
//         setClass(el,className)
//     }
//     if(style){
//         Object.entries(style).forEach(([prop,value])=>{
//             setStyle(el,prop,value)
//         })
//     }
//     for(const [name,value] of Object.entries(otherAttrs)){
//         setAttribute(el,name,value)
//     }
// }

// const setStyle = (el,name,value)=>{
//     el.style[name] = value
// }
// const removeStyle = (el,name)=>{
//     el.style[name]=null
// }
// const setAttribute = (el,name,value)=>{
//     if(value==null){
//         removeAttribute(el,name)
//     }else if(name.startsWith('data-')){
//         el.setAttribute(name,value)
//     }else{
//         el[name]=value
//     }
// }

// const removeAttribute = (el,name) =>{
//     el[name]=null
//     el.removeAttribute(name)
// }
