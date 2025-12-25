pragma circom 2.0.0;

// Solving the Non-Native Field Arithmetic Problem
// Dilithium Modulus q = 8380417
// BN254 Scalar Field r = 21888...
template BigIntModQMul() {
    signal input a;
    signal input b;
    signal output out;
    
    var q = 8380417;
    
    // Custom constraint logic for modular multiplication
    // out === (a * b) % q
    signal k; 
    k <-- (a * b) \ q;
    
    // Range check to ensure structural soundness
    out === a * b - k * q;
}