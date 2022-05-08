package dart_go_com

// #include "stdint.h"
// #include "string.h"
// #include "include/dart_api_dl.c"
//
// // Go does not allow calling C function pointers directly. So we are
// // forced to provide a trampoline.
// bool GoDart_PostCObject(Dart_Port_DL port, Dart_CObject* obj) {
//   return Dart_PostCObject_DL(port, obj);
// }
//
// typedef struct ComObj{
//      char *data;
// }ComObj;
//
// int64_t GetCom(void **ppCom, char* data) {
//      ComObj *pCom= (ComObj *)malloc(sizeof(ComObj));
//      pCom->data=data;
//      *ppCom = pCom;
//      int64_t ptr = (int64_t)pCom;
//      return ptr;
// }

import "C"
import "unsafe"

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		panic("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, data string) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64
	var pcom unsafe.Pointer
	ptrAddr := C.GetCom(&pcom, C.CString(data))
	*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrAddr
	C.GoDart_PostCObject(C.longlong(port), &obj)
	defer C.free(pcom)
}
