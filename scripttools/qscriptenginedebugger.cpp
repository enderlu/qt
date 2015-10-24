#include "qscriptenginedebugger.h"
#include <QScriptEngine>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QScriptEngineDebugger>
#include "_cgo_export.h"

class MyQScriptEngineDebugger: public QScriptEngineDebugger {
public:
void Signal_EvaluationResumed(){callbackQScriptEngineDebuggerEvaluationResumed(this->objectName().toUtf8().data());};
void Signal_EvaluationSuspended(){callbackQScriptEngineDebuggerEvaluationSuspended(this->objectName().toUtf8().data());};
};

QtObjectPtr QScriptEngineDebugger_NewQScriptEngineDebugger(QtObjectPtr parent){
	return new QScriptEngineDebugger(static_cast<QObject*>(parent));
}

QtObjectPtr QScriptEngineDebugger_Action(QtObjectPtr ptr, int action){
	return static_cast<QScriptEngineDebugger*>(ptr)->action(static_cast<QScriptEngineDebugger::DebuggerAction>(action));
}

void QScriptEngineDebugger_AttachTo(QtObjectPtr ptr, QtObjectPtr engine){
	static_cast<QScriptEngineDebugger*>(ptr)->attachTo(static_cast<QScriptEngine*>(engine));
}

int QScriptEngineDebugger_AutoShowStandardWindow(QtObjectPtr ptr){
	return static_cast<QScriptEngineDebugger*>(ptr)->autoShowStandardWindow();
}

QtObjectPtr QScriptEngineDebugger_CreateStandardMenu(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardMenu(static_cast<QWidget*>(parent));
}

QtObjectPtr QScriptEngineDebugger_CreateStandardToolBar(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QScriptEngineDebugger*>(ptr)->createStandardToolBar(static_cast<QWidget*>(parent));
}

void QScriptEngineDebugger_Detach(QtObjectPtr ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->detach();
}

void QScriptEngineDebugger_ConnectEvaluationResumed(QtObjectPtr ptr){
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));;
}

void QScriptEngineDebugger_DisconnectEvaluationResumed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationResumed), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationResumed));;
}

void QScriptEngineDebugger_ConnectEvaluationSuspended(QtObjectPtr ptr){
	QObject::connect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
}

void QScriptEngineDebugger_DisconnectEvaluationSuspended(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QScriptEngineDebugger*>(ptr), static_cast<void (QScriptEngineDebugger::*)()>(&QScriptEngineDebugger::evaluationSuspended), static_cast<MyQScriptEngineDebugger*>(ptr), static_cast<void (MyQScriptEngineDebugger::*)()>(&MyQScriptEngineDebugger::Signal_EvaluationSuspended));;
}

void QScriptEngineDebugger_SetAutoShowStandardWindow(QtObjectPtr ptr, int autoShow){
	static_cast<QScriptEngineDebugger*>(ptr)->setAutoShowStandardWindow(autoShow != 0);
}

QtObjectPtr QScriptEngineDebugger_StandardWindow(QtObjectPtr ptr){
	return static_cast<QScriptEngineDebugger*>(ptr)->standardWindow();
}

QtObjectPtr QScriptEngineDebugger_Widget(QtObjectPtr ptr, int widget){
	return static_cast<QScriptEngineDebugger*>(ptr)->widget(static_cast<QScriptEngineDebugger::DebuggerWidget>(widget));
}

void QScriptEngineDebugger_DestroyQScriptEngineDebugger(QtObjectPtr ptr){
	static_cast<QScriptEngineDebugger*>(ptr)->~QScriptEngineDebugger();
}

