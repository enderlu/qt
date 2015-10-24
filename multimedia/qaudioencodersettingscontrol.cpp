#include "qaudioencodersettingscontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioEncoderSettings>
#include <QAudioEncoderSettingsControl>
#include "_cgo_export.h"

class MyQAudioEncoderSettingsControl: public QAudioEncoderSettingsControl {
public:
};

char* QAudioEncoderSettingsControl_CodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->codecDescription(QString(codec)).toUtf8().data();
}

void QAudioEncoderSettingsControl_SetAudioSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QAudioEncoderSettingsControl*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

char* QAudioEncoderSettingsControl_SupportedAudioCodecs(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettingsControl*>(ptr)->supportedAudioCodecs().join("|").toUtf8().data();
}

void QAudioEncoderSettingsControl_DestroyQAudioEncoderSettingsControl(QtObjectPtr ptr){
	static_cast<QAudioEncoderSettingsControl*>(ptr)->~QAudioEncoderSettingsControl();
}

