package dart_go_com

// #include "stdlib.h"
// #include "stdint.h"
// #include "stdio.h"
// #include "include/dart_api_dl.c"
//
// // Go does not allow calling C function pointers directly. So we are
// // forced to provide a trampoline.
// bool GoDart_PostCObject(Dart_Port_DL port, int64_t ptrAddr) {
//   return Dart_PostCObject_DL(port, obj);
//   Dart_CObject dartObj;
//   dartObj.type = Dart_CObject_kInt64;
//	 dartObj.value.as_int64 = ptrAddr;
//	 return Dart_PostCObject_DL(port, &dartObj);
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
//
// void clearComStructMemory(ComObj pCom) {
//      free(&pCom.data);
// }

import "C"
import (
	"fmt"
	"unsafe"
)

func Init(api unsafe.Pointer) {
	if C.Dart_InitializeApiDL(api) != 0 {
		fmt.Errorf("failed to initialize Dart DL C API: version mismatch. " +
			"must update include/ to match Dart SDK version")
	}
}

func SendToPort(port int64, data string) {
	var obj C.Dart_CObject
	obj._type = C.Dart_CObject_kInt64

	var pcom unsafe.Pointer
	ptrAddr := C.GetCom(&pcom, C.CString(data))

	//*(*C.int64_t)(unsafe.Pointer(&obj.value[0])) = ptrAddr

	C.GoDart_PostCObject(C.int64_t(port), C.int64_t(ptrAddr))

}

func FreeComStructMemory(pointer *int64) {
	ptr := (*C.struct_ComObj)(unsafe.Pointer(pointer))
	C.clearComStructMemory(*ptr)
}
