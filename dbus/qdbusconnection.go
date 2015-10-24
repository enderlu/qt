package dbus

//#include "qdbusconnection.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QDBusConnection struct {
	ptr unsafe.Pointer
}

type QDBusConnectionITF interface {
	QDBusConnectionPTR() *QDBusConnection
}

func (p *QDBusConnection) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDBusConnection) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDBusConnection(ptr QDBusConnectionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusConnectionPTR().Pointer()
	}
	return nil
}

func QDBusConnectionFromPointer(ptr unsafe.Pointer) *QDBusConnection {
	var n = new(QDBusConnection)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDBusConnection) QDBusConnectionPTR() *QDBusConnection {
	return ptr
}

//QDBusConnection::BusType
type QDBusConnection__BusType int

var (
	QDBusConnection__SessionBus    = QDBusConnection__BusType(0)
	QDBusConnection__SystemBus     = QDBusConnection__BusType(1)
	QDBusConnection__ActivationBus = QDBusConnection__BusType(2)
)

//QDBusConnection::ConnectionCapability
type QDBusConnection__ConnectionCapability int

var (
	QDBusConnection__UnixFileDescriptorPassing = QDBusConnection__ConnectionCapability(0x0001)
)

//QDBusConnection::RegisterOption
type QDBusConnection__RegisterOption int

var (
	QDBusConnection__ExportAdaptors                = QDBusConnection__RegisterOption(0x01)
	QDBusConnection__ExportScriptableSlots         = QDBusConnection__RegisterOption(0x10)
	QDBusConnection__ExportScriptableSignals       = QDBusConnection__RegisterOption(0x20)
	QDBusConnection__ExportScriptableProperties    = QDBusConnection__RegisterOption(0x40)
	QDBusConnection__ExportScriptableInvokables    = QDBusConnection__RegisterOption(0x80)
	QDBusConnection__ExportScriptableContents      = QDBusConnection__RegisterOption(0xf0)
	QDBusConnection__ExportNonScriptableSlots      = QDBusConnection__RegisterOption(0x100)
	QDBusConnection__ExportNonScriptableSignals    = QDBusConnection__RegisterOption(0x200)
	QDBusConnection__ExportNonScriptableProperties = QDBusConnection__RegisterOption(0x400)
	QDBusConnection__ExportNonScriptableInvokables = QDBusConnection__RegisterOption(0x800)
	QDBusConnection__ExportNonScriptableContents   = QDBusConnection__RegisterOption(0xf00)
	QDBusConnection__ExportAllSlots                = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSlots | QDBusConnection__ExportNonScriptableSlots)
	QDBusConnection__ExportAllSignals              = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableSignals | QDBusConnection__ExportNonScriptableSignals)
	QDBusConnection__ExportAllProperties           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableProperties | QDBusConnection__ExportNonScriptableProperties)
	QDBusConnection__ExportAllInvokables           = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableInvokables | QDBusConnection__ExportNonScriptableInvokables)
	QDBusConnection__ExportAllContents             = QDBusConnection__RegisterOption(QDBusConnection__ExportScriptableContents | QDBusConnection__ExportNonScriptableContents)
	QDBusConnection__ExportChildObjects            = QDBusConnection__RegisterOption(0x1000)
)

//QDBusConnection::UnregisterMode
type QDBusConnection__UnregisterMode int

var (
	QDBusConnection__UnregisterNode = QDBusConnection__UnregisterMode(0)
	QDBusConnection__UnregisterTree = QDBusConnection__UnregisterMode(1)
)

func NewQDBusConnection2(other QDBusConnectionITF) *QDBusConnection {
	return QDBusConnectionFromPointer(unsafe.Pointer(C.QDBusConnection_NewQDBusConnection2(C.QtObjectPtr(PointerFromQDBusConnection(other)))))
}

func NewQDBusConnection(name string) *QDBusConnection {
	return QDBusConnectionFromPointer(unsafe.Pointer(C.QDBusConnection_NewQDBusConnection(C.CString(name))))
}

func (ptr *QDBusConnection) BaseService() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_BaseService(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusConnection) CallWithCallback(message QDBusMessageITF, receiver core.QObjectITF, returnMethod string, errorMethod string, timeout int) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_CallWithCallback(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusMessage(message)), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(returnMethod), C.CString(errorMethod), C.int(timeout)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect(service string, path string, interfa string, name string, receiver core.QObjectITF, slot string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect(C.QtObjectPtr(ptr.Pointer()), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect2(service string, path string, interfa string, name string, signature string, receiver core.QObjectITF, slot string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect2(C.QtObjectPtr(ptr.Pointer()), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(signature), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Connect3(service string, path string, interfa string, name string, argumentMatch []string, signature string, receiver core.QObjectITF, slot string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_Connect3(C.QtObjectPtr(ptr.Pointer()), C.CString(service), C.CString(path), C.CString(interfa), C.CString(name), C.CString(strings.Join(argumentMatch, "|")), C.CString(signature), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(slot)) != 0
	}
	return false
}

func (ptr *QDBusConnection) ConnectionCapabilities() QDBusConnection__ConnectionCapability {
	if ptr.Pointer() != nil {
		return QDBusConnection__ConnectionCapability(C.QDBusConnection_ConnectionCapabilities(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func QDBusConnection_DisconnectFromBus(name string) {
	C.QDBusConnection_QDBusConnection_DisconnectFromBus(C.CString(name))
}

func QDBusConnection_DisconnectFromPeer(name string) {
	C.QDBusConnection_QDBusConnection_DisconnectFromPeer(C.CString(name))
}

func (ptr *QDBusConnection) Interface() *QDBusConnectionInterface {
	if ptr.Pointer() != nil {
		return QDBusConnectionInterfaceFromPointer(unsafe.Pointer(C.QDBusConnection_Interface(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QDBusConnection) IsConnected() bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_IsConnected(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QDBusConnection) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusConnection_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QDBusConnection) ObjectRegisteredAt(path string) *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QDBusConnection_ObjectRegisteredAt(C.QtObjectPtr(ptr.Pointer()), C.CString(path))))
	}
	return nil
}

func (ptr *QDBusConnection) RegisterObject(path string, object core.QObjectITF, options QDBusConnection__RegisterOption) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject(C.QtObjectPtr(ptr.Pointer()), C.CString(path), C.QtObjectPtr(core.PointerFromQObject(object)), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterObject2(path string, interfa string, object core.QObjectITF, options QDBusConnection__RegisterOption) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterObject2(C.QtObjectPtr(ptr.Pointer()), C.CString(path), C.CString(interfa), C.QtObjectPtr(core.PointerFromQObject(object)), C.int(options)) != 0
	}
	return false
}

func (ptr *QDBusConnection) RegisterService(serviceName string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_RegisterService(C.QtObjectPtr(ptr.Pointer()), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) Send(message QDBusMessageITF) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_Send(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQDBusMessage(message))) != 0
	}
	return false
}

func (ptr *QDBusConnection) UnregisterObject(path string, mode QDBusConnection__UnregisterMode) {
	if ptr.Pointer() != nil {
		C.QDBusConnection_UnregisterObject(C.QtObjectPtr(ptr.Pointer()), C.CString(path), C.int(mode))
	}
}

func (ptr *QDBusConnection) UnregisterService(serviceName string) bool {
	if ptr.Pointer() != nil {
		return C.QDBusConnection_UnregisterService(C.QtObjectPtr(ptr.Pointer()), C.CString(serviceName)) != 0
	}
	return false
}

func (ptr *QDBusConnection) DestroyQDBusConnection() {
	if ptr.Pointer() != nil {
		C.QDBusConnection_DestroyQDBusConnection(C.QtObjectPtr(ptr.Pointer()))
	}
}
