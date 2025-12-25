package vm

// Opcode represents an EVM opcode
type Opcode byte

const (
	// ... existing opcodes ...
	CREATE2     Opcode = 0xf5
	
	// ----------------------------------------------------------------
	// [Research Proposal Modification]
	// New Opcode Definition for Post-Quantum Signature Verification
	// NIST FIPS 204 (ML-DSA) Support
	// ----------------------------------------------------------------
	VERIFY_DILITHIUM Opcode = 0xf6 // <-- 我們新增的指令
	// ----------------------------------------------------------------

	STATICCALL  Opcode = 0xfa
	// ...
)

func (op Opcode) IsReplica() bool {
	return op == VERIFY_DILITHIUM // Ensure it's treated as a valid operation
}