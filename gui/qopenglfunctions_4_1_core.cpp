#include "qopenglfunctions_4_1_core.h"
#include <QOpenGLFunctions>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOpenGLFunctions_4_1_Core>
#include "_cgo_export.h"

class MyQOpenGLFunctions_4_1_Core: public QOpenGLFunctions_4_1_Core {
public:
};

QtObjectPtr QOpenGLFunctions_4_1_Core_NewQOpenGLFunctions_4_1_Core(){
	return new QOpenGLFunctions_4_1_Core();
}

int QOpenGLFunctions_4_1_Core_InitializeOpenGLFunctions(QtObjectPtr ptr){
	return static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->initializeOpenGLFunctions();
}

void QOpenGLFunctions_4_1_Core_DestroyQOpenGLFunctions_4_1_Core(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->~QOpenGLFunctions_4_1_Core();
}

void QOpenGLFunctions_4_1_Core_GlEndConditionalRender(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glEndConditionalRender();
}

void QOpenGLFunctions_4_1_Core_GlEndTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glEndTransformFeedback();
}

void QOpenGLFunctions_4_1_Core_GlFinish(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glFinish();
}

void QOpenGLFunctions_4_1_Core_GlFlush(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glFlush();
}

void QOpenGLFunctions_4_1_Core_GlPauseTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glPauseTransformFeedback();
}

void QOpenGLFunctions_4_1_Core_GlReleaseShaderCompiler(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glReleaseShaderCompiler();
}

void QOpenGLFunctions_4_1_Core_GlResumeTransformFeedback(QtObjectPtr ptr){
	static_cast<QOpenGLFunctions_4_1_Core*>(ptr)->glResumeTransformFeedback();
}

