// crypto/pqc/bridge.c
#include <oqs/oqs.h>
#include <stdlib.h>
#include <string.h>

// 這是 Helper 函數，用來幫助 Go 呼叫 C
int bridge_verify(uint8_t *msg, size_t msg_len, uint8_t *sig, size_t sig_len, uint8_t *pk, size_t pk_len) {
    // -------------------------------------------------------------
    // [Update] 使用 NIST FIPS 204 正式名稱: ML-DSA-65
    // 舊名稱 OQS_SIG_alg_dilithium_3 在新版 liboqs 已被移除
    // -------------------------------------------------------------
    OQS_SIG *sig_alg = OQS_SIG_new(OQS_SIG_alg_ml_dsa_65);
    
    if (sig_alg == NULL) {
        // 如果初始化失敗 (例如系統不支援 AVX2 指令集導致無法載入)，回傳失敗
        return OQS_ERROR;
    }

    // 執行驗證
    OQS_STATUS status = OQS_SIG_verify(sig_alg, msg, msg_len, sig, sig_len, pk);
    
    // 釋放記憶體
    OQS_SIG_free(sig_alg);
    
    return (status == OQS_SUCCESS) ? 1 : 0;
}