package pqc

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -loqs -Wl,-rpath,/usr/local/lib
#include "bridge.h"
*/
import "C"
import (
	"unsafe"
)

// VerifySignature exposes the C logic to Go
func VerifySignature(msg []byte, sig []byte, pubKey []byte) (bool, error) {
	// 1. 基本檢查：如果簽章或公鑰是空的，直接回傳失敗，避免後續指針錯誤
	if len(sig) == 0 || len(pubKey) == 0 {
		return false, nil
	}

	// 2. 安全地取得 C 指標 (Safe Pointer Conversion)
	// 如果 slice 長度 > 0，才取第 0 個元素的位址；否則給 nil
	var cMsg *C.uint8_t
	if len(msg) > 0 {
		cMsg = (*C.uint8_t)(unsafe.Pointer(&msg[0]))
	}

	// 因為上面已經檢查過 sig 和 pubKey 不為空，這裡可以直接取址
	// 但為了養成好習慣，我們還是用安全的方式寫
	var cSig *C.uint8_t
	if len(sig) > 0 {
		cSig = (*C.uint8_t)(unsafe.Pointer(&sig[0]))
	}

	var cPk *C.uint8_t
	if len(pubKey) > 0 {
		cPk = (*C.uint8_t)(unsafe.Pointer(&pubKey[0]))
	}

	// 呼叫 C 語言定義的 bridge_verify 函數
	res := C.bridge_verify(
		cMsg, C.size_t(len(msg)),
		cSig, C.size_t(len(sig)),
		cPk, C.size_t(len(pubKey)),
	)

	if res == 1 {
		return true, nil
	}
	return false, nil
}