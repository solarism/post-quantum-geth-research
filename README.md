# Post-Quantum Ethereum (Geth) Research Prototype, 威震＠NTUB 

[![Go Report Card](https://goreportcard.com/badge/github.com/ethereum/go-ethereum)](https://goreportcard.com/report/github.com/ethereum/go-ethereum)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/ethereum/go-ethereum)](https://github.com/ethereum/go-ethereum)
[![NIST Standard](https://img.shields.io/badge/NIST-FIPS%20204-blue)](https://csrc.nist.gov/pubs/fips/204/final)

## 📖 Project Overview

This repository hosts the experimental implementation of the **Post-Quantum DeFi Resilience Framework**, a research initiative aimed at integrating **NIST FIPS 204 (ML-DSA)** standards directly into the Ethereum Layer 1 core and Layer 2 scaling solutions.

The goal is to address the "Harvest Now, Decrypt Later" threat by introducing lattice-based cryptography support at the EVM level while solving the associated data availability and gas cost bottlenecks via Zero-Knowledge Proofs (ZK-Rollup).

**Principal Investigator:** Prof. Wei-Chen Wu  
**Grant Proposal:** Engineering Division, NSTC (Taiwan)

---

## 🏗 System Architecture & File Tree

This prototype modifies the core `go-ethereum` codebase to support post-quantum primitives. Below is the directory structure highlighting our modifications:

```text
post-quantum-geth-research/
├── benchmarks/                        <-- [Perf] Performance Benchmarking Module
│   └── opcode_pricing/
│       └── Cargo.toml                 <-- Rust/Criterion.rs config for SGX benchmarks
│
├── circuits/                          <-- [L2] ZK-Rollup Scaling Layer
│   ├── dilithium_verifier.circom      <-- Main Circuit: Aggregated PQC signature verification
│   └── lib/
│       └── bigint_mod_q.circom        <-- Math Core: Custom gadget for Dilithium field arithmetic
│
├── core/                              <-- [L1] EVM Core Modifications
│   └── vm/
│       ├── contracts.go               <-- Mod: Precompiled contract registration (Address 0x10)
│       ├── instructions.go            <-- Mod: Implementation of opVerifyDilithium logic
│       ├── jump_table.go              <-- Mod: Opcode 0xF6 routing & gas binding
│       ├── opcodes.go                 <-- New: Definition of VERIFY_DILITHIUM (0xf6)
│       └── pqc_gas_pricing.go         <-- New: Dynamic Gas Pricing Formula (Eq. 3.1)
│
├── crypto/                            <-- [Impl] Post-Quantum Cryptography Implementation
│   └── pqc/
│       ├── bridge.go                  <-- Interface: CGO linking logic
│       ├── dilithium.go               <-- Logic: Wrapper for liboqs (NIST FIPS 204)
│       └── dilithium_fuzz_test.go     <-- Test: Automated Fuzzing for security boundaries
│
└── README.md
