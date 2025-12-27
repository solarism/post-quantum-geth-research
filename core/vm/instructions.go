package vm

//import (
//	"github.com/ethereum/go-ethereum/crypto/pqc" // 引用我們自己寫的套件
//	"github.com/ethereum/go-ethereum/common"
//)

import (
    "github.com/solarism/post-quantum-geth-research/crypto/pqc" // 改成本地路徑
    "github.com/ethereum/go-ethereum/common"
)

// ----------------------------------------------------------------
// [Implementation Detail]
// opVerifyDilithium executes the lattice-based verification.
//
// Stack Inputs:
// [0] Message Hash (32 bytes)
// [1] Signature Pointer (Memory Offset)
// [2] Public Key Pointer (Memory Offset)
//
// Output:
// [0] 1 if valid, 0 if invalid
// ----------------------------------------------------------------
func opVerifyDilithium(pc *uint64, interpreter *EVMInterpreter, scope *ScopeContext) ([]byte, error) {
	// 1. Pop parameters from the EVM Stack
	msgHash := scope.Stack.pop()
	sigOffset := scope.Stack.pop()
	pubKeyOffset := scope.Stack.pop()

	// 2. Read huge PQC data from Memory (Expensive operation!)
	// Dilithium signature size ~2.4KB
	signature := scope.Memory.GetPtr(sigOffset.Int64(), 2420)
	pubKey := scope.Memory.GetPtr(pubKeyOffset.Int64(), 1312)

	// 3. Call the CGO Bridge (liboqs)
	// This crosses the boundary from Go (Safe) to C (Unsafe/Fast)
	valid, err := pqc.VerifySignature(msgHash.Bytes(), signature, pubKey)
	if err != nil {
		return nil, err
	}

	// 4. Push result back to Stack
	if valid {
		scope.Stack.push(common.Big1)
	} else {
		scope.Stack.push(common.Big0)
	}

	return nil, nil
}