package widgets

//#include "qfileiconprovider.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFileIconProvider struct {
	ptr unsafe.Pointer
}

type QFileIconProviderITF interface {
	QFileIconProviderPTR() *QFileIconProvider
}

func (p *QFileIconProvider) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QFileIconProvider) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQFileIconProvider(ptr QFileIconProviderITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileIconProviderPTR().Pointer()
	}
	return nil
}

func QFileIconProviderFromPointer(ptr unsafe.Pointer) *QFileIconProvider {
	var n = new(QFileIconProvider)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFileIconProvider) QFileIconProviderPTR() *QFileIconProvider {
	return ptr
}

//QFileIconProvider::IconType
type QFileIconProvider__IconType int

var (
	QFileIconProvider__Computer = QFileIconProvider__IconType(0)
	QFileIconProvider__Desktop  = QFileIconProvider__IconType(1)
	QFileIconProvider__Trashcan = QFileIconProvider__IconType(2)
	QFileIconProvider__Network  = QFileIconProvider__IconType(3)
	QFileIconProvider__Drive    = QFileIconProvider__IconType(4)
	QFileIconProvider__Folder   = QFileIconProvider__IconType(5)
	QFileIconProvider__File     = QFileIconProvider__IconType(6)
)

//QFileIconProvider::Option
type QFileIconProvider__Option int

var (
	QFileIconProvider__DontUseCustomDirectoryIcons = QFileIconProvider__Option(0x00000001)
)

func NewQFileIconProvider() *QFileIconProvider {
	return QFileIconProviderFromPointer(unsafe.Pointer(C.QFileIconProvider_NewQFileIconProvider()))
}

func (ptr *QFileIconProvider) Options() QFileIconProvider__Option {
	if ptr.Pointer() != nil {
		return QFileIconProvider__Option(C.QFileIconProvider_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileIconProvider) SetOptions(options QFileIconProvider__Option) {
	if ptr.Pointer() != nil {
		C.QFileIconProvider_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func (ptr *QFileIconProvider) Type(info core.QFileInfoITF) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileIconProvider_Type(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQFileInfo(info))))
	}
	return ""
}

func (ptr *QFileIconProvider) DestroyQFileIconProvider() {
	if ptr.Pointer() != nil {
		C.QFileIconProvider_DestroyQFileIconProvider(C.QtObjectPtr(ptr.Pointer()))
	}
}
