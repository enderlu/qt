#include "qsensorgesturemanager.h"
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorGestureRecognizer>
#include <QSensorGesture>
#include <QObject>
#include <QSensorGestureManager>
#include "_cgo_export.h"

class MyQSensorGestureManager: public QSensorGestureManager {
public:
void Signal_NewSensorGestureAvailable(){callbackQSensorGestureManagerNewSensorGestureAvailable(this->objectName().toUtf8().data());};
};

QtObjectPtr QSensorGestureManager_NewQSensorGestureManager(QtObjectPtr parent){
	return new QSensorGestureManager(static_cast<QObject*>(parent));
}

char* QSensorGestureManager_GestureIds(QtObjectPtr ptr){
	return static_cast<QSensorGestureManager*>(ptr)->gestureIds().join("|").toUtf8().data();
}

void QSensorGestureManager_ConnectNewSensorGestureAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

void QSensorGestureManager_DisconnectNewSensorGestureAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

char* QSensorGestureManager_RecognizerSignals(QtObjectPtr ptr, char* gestureId){
	return static_cast<QSensorGestureManager*>(ptr)->recognizerSignals(QString(gestureId)).join("|").toUtf8().data();
}

int QSensorGestureManager_RegisterSensorGestureRecognizer(QtObjectPtr ptr, QtObjectPtr recognizer){
	return static_cast<QSensorGestureManager*>(ptr)->registerSensorGestureRecognizer(static_cast<QSensorGestureRecognizer*>(recognizer));
}

QtObjectPtr QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id){
	return QSensorGestureManager::sensorGestureRecognizer(QString(id));
}

void QSensorGestureManager_DestroyQSensorGestureManager(QtObjectPtr ptr){
	static_cast<QSensorGestureManager*>(ptr)->~QSensorGestureManager();
}

