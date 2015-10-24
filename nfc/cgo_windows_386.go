package nfc

/*
#cgo CPPFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -Wextra -fexceptions -mthreads
#cgo CPPFLAGS: -DUNICODE -DQT_NO_DEBUG -DQT_CORE_LIB -DQT_NFC_LIB
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include -IC:/Qt/Qt5.5.1/5.5/mingw492_32/mkspecs/win32-g++
#cgo CPPFLAGS: -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtCore -IC:/Qt/Qt5.5.1/5.5/mingw492_32/include/QtNfc

#cgo LDFLAGS: -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS: -LC:/Qt/Qt5.5.1/5.5/mingw492_32/bin -lQt5Core -lQt5Nfc
*/
import "C"
