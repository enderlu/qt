#include "qbluetoothdeviceinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QBluetoothUuid>
#include <QBluetoothDeviceInfo>
#include "_cgo_export.h"

class MyQBluetoothDeviceInfo: public QBluetoothDeviceInfo {
public:
};

QtObjectPtr QBluetoothDeviceInfo_NewQBluetoothDeviceInfo(){
	return new QBluetoothDeviceInfo();
}

QtObjectPtr QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(QtObjectPtr other){
	return new QBluetoothDeviceInfo(*static_cast<QBluetoothDeviceInfo*>(other));
}

int QBluetoothDeviceInfo_CoreConfigurations(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->coreConfigurations();
}

int QBluetoothDeviceInfo_IsCached(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isCached();
}

int QBluetoothDeviceInfo_IsValid(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->isValid();
}

int QBluetoothDeviceInfo_MajorDeviceClass(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->majorDeviceClass();
}

char* QBluetoothDeviceInfo_Name(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->name().toUtf8().data();
}

int QBluetoothDeviceInfo_ServiceClasses(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceClasses();
}

int QBluetoothDeviceInfo_ServiceUuidsCompleteness(QtObjectPtr ptr){
	return static_cast<QBluetoothDeviceInfo*>(ptr)->serviceUuidsCompleteness();
}

void QBluetoothDeviceInfo_SetCached(QtObjectPtr ptr, int cached){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCached(cached != 0);
}

void QBluetoothDeviceInfo_SetCoreConfigurations(QtObjectPtr ptr, int coreConfigs){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setCoreConfigurations(static_cast<QBluetoothDeviceInfo::CoreConfiguration>(coreConfigs));
}

void QBluetoothDeviceInfo_SetDeviceUuid(QtObjectPtr ptr, QtObjectPtr uuid){
	static_cast<QBluetoothDeviceInfo*>(ptr)->setDeviceUuid(*static_cast<QBluetoothUuid*>(uuid));
}

void QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(QtObjectPtr ptr){
	static_cast<QBluetoothDeviceInfo*>(ptr)->~QBluetoothDeviceInfo();
}

