package lang

import (
	"GoVM/native"
	"GoVM/chapter4-rtdt"
	"GoVM/chapter6-obj/heap"
)

const jlSystem = "java/lang/System"

func init() {
	native.Register(jlSystem, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

func arraycopy(frame *chapter4_rtdt.Frame) {
	vars := frame.LocalVars()

	//source
	src := vars.GetRef(0)
	if src == nil {
		panic("java.lang.NullPointerException")
	}
	srcPos := vars.GetInt(1)

	//destination
	dest := vars.GetRef(2)
	if dest == nil {
		panic("java.lang.NullPointerException")
	}
	destPos := vars.GetInt(3)

	//data length
	length := vars.GetInt(4)

	//源数组和目标数组必须兼容，否则不能拷贝
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || destPos < 0 || length < 0 || srcPos + length > src.ArrayLength() || destPos + length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}