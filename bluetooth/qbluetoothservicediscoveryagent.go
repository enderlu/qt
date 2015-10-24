package bluetooth

//#include "qbluetoothservicediscoveryagent.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgentITF interface {
	core.QObjectITF
	QBluetoothServiceDiscoveryAgentPTR() *QBluetoothServiceDiscoveryAgent
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgentPTR().Pointer()
	}
	return nil
}

func QBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QBluetoothServiceDiscoveryAgent_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgentPTR() *QBluetoothServiceDiscoveryAgent {
	return ptr
}

//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int

var (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int

var (
	QBluetoothServiceDiscoveryAgent__NoError                      = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentCanceled
func callbackQBluetoothServiceDiscoveryAgentCanceled(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "canceled").(func())()
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

//export callbackQBluetoothServiceDiscoveryAgentFinished
func callbackQBluetoothServiceDiscoveryAgentFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "finished").(func())()
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObjectITF) *QBluetoothServiceDiscoveryAgent {
	return QBluetoothServiceDiscoveryAgentFromPointer(unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddressITF, parent core.QObjectITF) *QBluetoothServiceDiscoveryAgent {
	return QBluetoothServiceDiscoveryAgentFromPointer(unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(C.QtObjectPtr(PointerFromQBluetoothAddress(deviceAdapter)), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddressITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothAddress(address))) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuidITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothUuid(uuid)))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
