// crypto/pqc/bridge.h
#ifndef PQC_BRIDGE_H
#define PQC_BRIDGE_H

#include <stdint.h>
#include <stddef.h>

int bridge_verify(uint8_t *msg, size_t msg_len, uint8_t *sig, size_t sig_len, uint8_t *pk, size_t pk_len);

#endif