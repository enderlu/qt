package gui

//#include "qopenglfunctions_4_4_core.h"
import "C"
import (
	"unsafe"
)

type QOpenGLFunctions_4_4_Core struct {
	QAbstractOpenGLFunctions
}

type QOpenGLFunctions_4_4_CoreITF interface {
	QAbstractOpenGLFunctionsITF
	QOpenGLFunctions_4_4_CorePTR() *QOpenGLFunctions_4_4_Core
}

func PointerFromQOpenGLFunctions_4_4_Core(ptr QOpenGLFunctions_4_4_CoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLFunctions_4_4_CorePTR().Pointer()
	}
	return nil
}

func QOpenGLFunctions_4_4_CoreFromPointer(ptr unsafe.Pointer) *QOpenGLFunctions_4_4_Core {
	var n = new(QOpenGLFunctions_4_4_Core)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLFunctions_4_4_Core) QOpenGLFunctions_4_4_CorePTR() *QOpenGLFunctions_4_4_Core {
	return ptr
}

func NewQOpenGLFunctions_4_4_Core() *QOpenGLFunctions_4_4_Core {
	return QOpenGLFunctions_4_4_CoreFromPointer(unsafe.Pointer(C.QOpenGLFunctions_4_4_Core_NewQOpenGLFunctions_4_4_Core()))
}

func (ptr *QOpenGLFunctions_4_4_Core) InitializeOpenGLFunctions() bool {
	if ptr.Pointer() != nil {
		return C.QOpenGLFunctions_4_4_Core_InitializeOpenGLFunctions(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QOpenGLFunctions_4_4_Core) DestroyQOpenGLFunctions_4_4_Core() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_DestroyQOpenGLFunctions_4_4_Core(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlEndConditionalRender() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlEndConditionalRender(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlEndTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlEndTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlFinish() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlFinish(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlFlush() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlFlush(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlPauseTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlPauseTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlPopDebugGroup() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlPopDebugGroup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlReleaseShaderCompiler() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlReleaseShaderCompiler(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QOpenGLFunctions_4_4_Core) GlResumeTransformFeedback() {
	if ptr.Pointer() != nil {
		C.QOpenGLFunctions_4_4_Core_GlResumeTransformFeedback(C.QtObjectPtr(ptr.Pointer()))
	}
}
