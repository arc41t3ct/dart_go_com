package dart_go_com

// #include "stdint.h"
// #include "include/dart_api_dl.c"
//
// // Go does not allow calling C function pointers directly. So we are
// // forced to provide a trampoline.
// bool GoDart_PostCObject(Dart_Port_DL port, Dart_CObject* obj) {
//   return Dart_PostCObject_DL(port, obj);
// }
import "C"
import "unsafe"

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, msg int64) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64
	// cgo does not support unions so we are forced to do this
	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = C.int64_t(msg)
	// func _Cfunc_GoDart_PostCObject(p0 _Ctype_Dart_Port_DL, p1 *_Ctype_struct__Dart_CObject) (r1 _Ctype__Bool)
	C.GoDart_PostCObject(C.longlong(port), &obj)
}
