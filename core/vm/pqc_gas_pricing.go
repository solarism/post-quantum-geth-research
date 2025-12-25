package vm

import (
	"github.com/ethereum/go-ethereum/params"
)

// ----------------------------------------------------------------
// [Performance Optimization]
// Dynamic Gas Pricing Model based on Intel SGX Benchmarks.
//
// Formula Reference (Proposal Eq. 3.1):
// Cost = G_base + G_mem * S_sig + G_ops * N_matrix
// ----------------------------------------------------------------

const (
	// G_base: Fixed cost for calling the PQC precompile
	// Adjusted from 21000 to 50000 based on initial overhead analysis
	GasDilithiumBase = 50000

	// G_ops: Cost per lattice matrix operation unit (AVX2 optimized)
	// Derived from Criterion.rs benchmarks on SGX environment
	GasDilithiumOpUnit = 120 
)

// CalculateDilithiumGas computes the exact gas cost for a transaction
// ensuring validators are compensated for CPU cycles.
func CalculateDilithiumGas(sigLen int, matrixDim int) uint64 {
	// 1. Base Cost
	gas := uint64(GasDilithiumBase)

	// 2. Memory Expansion Cost (Linear)
	// EIP-2028: 16 gas per non-zero byte
	memCost := uint64(sigLen) * params.TxDataNonZeroGas
	gas += memCost

	// 3. Computational Cost (Polynomial)
	// Complexity grows with matrix dimension squared in ML-DSA
	computeCost := uint64(matrixDim * matrixDim) * GasDilithiumOpUnit
	gas += computeCost

	return gas
}