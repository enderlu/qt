package gui

//#include "qaccessibleactioninterface.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QAccessibleActionInterface struct {
	ptr unsafe.Pointer
}

type QAccessibleActionInterfaceITF interface {
	QAccessibleActionInterfacePTR() *QAccessibleActionInterface
}

func (p *QAccessibleActionInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessibleActionInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessibleActionInterface(ptr QAccessibleActionInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleActionInterfacePTR().Pointer()
	}
	return nil
}

func QAccessibleActionInterfaceFromPointer(ptr unsafe.Pointer) *QAccessibleActionInterface {
	var n = new(QAccessibleActionInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleActionInterface) QAccessibleActionInterfacePTR() *QAccessibleActionInterface {
	return ptr
}

func (ptr *QAccessibleActionInterface) LocalizedActionDescription(actionName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleActionInterface_LocalizedActionDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName)))
	}
	return ""
}

func (ptr *QAccessibleActionInterface) LocalizedActionName(actionName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleActionInterface_LocalizedActionName(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName)))
	}
	return ""
}

func (ptr *QAccessibleActionInterface) ActionNames() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleActionInterface_ActionNames(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func QAccessibleActionInterface_DecreaseAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction())
}

func (ptr *QAccessibleActionInterface) DoAction(actionName string) {
	if ptr.Pointer() != nil {
		C.QAccessibleActionInterface_DoAction(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName))
	}
}

func QAccessibleActionInterface_IncreaseAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction())
}

func (ptr *QAccessibleActionInterface) KeyBindingsForAction(actionName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QAccessibleActionInterface_KeyBindingsForAction(C.QtObjectPtr(ptr.Pointer()), C.CString(actionName))), "|")
	}
	return make([]string, 0)
}

func QAccessibleActionInterface_NextPageAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_NextPageAction())
}

func QAccessibleActionInterface_PressAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_PressAction())
}

func QAccessibleActionInterface_PreviousPageAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_PreviousPageAction())
}

func QAccessibleActionInterface_ScrollDownAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollDownAction())
}

func QAccessibleActionInterface_ScrollLeftAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollLeftAction())
}

func QAccessibleActionInterface_ScrollRightAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollRightAction())
}

func QAccessibleActionInterface_ScrollUpAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ScrollUpAction())
}

func QAccessibleActionInterface_SetFocusAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_SetFocusAction())
}

func QAccessibleActionInterface_ShowMenuAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ShowMenuAction())
}

func QAccessibleActionInterface_ToggleAction() string {
	return C.GoString(C.QAccessibleActionInterface_QAccessibleActionInterface_ToggleAction())
}

func (ptr *QAccessibleActionInterface) DestroyQAccessibleActionInterface() {
	if ptr.Pointer() != nil {
		C.QAccessibleActionInterface_DestroyQAccessibleActionInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
