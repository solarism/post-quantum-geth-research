package pqc

/*
#cgo CFLAGS: -I/usr/local/include/liboqs
#cgo LDFLAGS: -loqs
#include <oqs/oqs.h>
#include "bridge.c" // Custom bridge for memory safety
*/
import "C"
import (
    "errors"
    "unsafe"
)

// VerifySignature wraps the OQS_SIG_verify function
func VerifySignature(data []byte) ([]byte, error) {
    // Convert Go slice to C pointer (Zero-Copy optimization)
    msgLen := C.size_t(len(data))
    msgPtr := (*C.uint8_t)(unsafe.Pointer(&data[0]))
    
    // Call liboqs (C implementation of CRYSTALS-Dilithium)
    // Using OQS_SIG_alg_dilithium_3 (NIST Level 3)
    result := C.verify_dilithium_3(msgPtr, msgLen)
    
    if result != C.OQS_SUCCESS {
        return nil, errors.New("FIPS 204 signature verification failed")
    }
    
    // Return 1 (true) aligned with EVM standard (32 bytes)
    return common.LeftPadBytes([]byte{1}, 32), nil
}