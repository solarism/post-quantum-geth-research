package pqc

import (
	"testing"
)

// ----------------------------------------------------------------
// [Security Testing]
// Fuzz testing for Dilithium implementation.
// This automated test generates random byte sequences to ensure
// the CGO bridge handles invalid inputs gracefully without crashing (Panic).
// ----------------------------------------------------------------
func FuzzVerifySignature(f *testing.F) {
	// 1. Add Seed Corpus (Valid-looking data samples to start with)
	f.Add([]byte("test_message"), []byte("mock_sig"), []byte("mock_pk"))

	// 2. Fuzz Target
	f.Fuzz(func(t *testing.T, msg []byte, sig []byte, pubKey []byte) {
		// Limit input size to prevent OOM during fuzzing
		if len(sig) > 5000 || len(pubKey) > 2000 {
			return
		}

		// Call the target function
		// We expect errors for random data, but we MUST NOT see crashes/panics
		valid, err := VerifySignature(msg, sig, pubKey)

		// Assertion: If it's valid, it must not return error
		if valid && err != nil {
			t.Errorf("Valid signature returned error: %v", err)
		}
	})
}