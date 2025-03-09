package internals

func DestroyDom(vdom Vdom) {
	vdom.RemoveNode()
}
