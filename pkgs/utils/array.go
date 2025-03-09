package utils

import (
	"go-fe-fwk/types"
	"reflect"
)

// func WithoutNulls(arr []internals.Vdom) []internals.Vdom {
// 	var newArr = make([]internals.Vdom, 0, len(arr))
// 	for _, value := range arr {

//			// if value != nil{
//			newArr = append(newArr, value)
//			// }
//		}
//		return newArr
//	}

func WithoutNulls() {}
func JSFuncExists(funcArray []types.JsFunc, target types.JsFunc) bool {
	for _, f := range funcArray {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return true
		}
	}
	return false
}
func FuncExists(funcArray []types.AnyFunc, target types.AnyFunc) bool {
	for _, f := range funcArray {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return true
		}
	}
	return false
}

func ReducerFuncExists(funcArray []types.ReducerFunc, target types.ReducerFunc) bool {
	for _, f := range funcArray {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return true
		}
	}
	return false
}

func IndexOfJsFunction(arr []types.JsFunc, target types.JsFunc) int {
	for i, f := range arr {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return i
		}
	}
	return -1
}

func IndexOfFunction(arr []types.AnyFunc, target types.AnyFunc) int {
	for i, f := range arr {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return i
		}
	}
	return -1
}
func IndexOfReducerFunction(arr []types.ReducerFunc, target types.ReducerFunc) int {
	for i, f := range arr {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return i
		}
	}
	return -1
}
