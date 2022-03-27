package desoapi

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	desoCore "github.com/deso-protocol/core/lib"
)

func SignTransactionWithDerivedKey(transactionHex string, derivedKeySeedHex string) (string, error) {
	txnBytes, err := hex.DecodeString(transactionHex)
	if err != nil {
		return "", fmt.Errorf("problem decoding transaction hex: %w", err)
	}

	derivedKeyBytes, err := hex.DecodeString(derivedKeySeedHex)
	if err != nil {
		return "", fmt.Errorf("problem decoding derived key seed hex: %w", err)
	}

	privateKeyBytes, _ := btcec.PrivKeyFromBytes(btcec.S256(), derivedKeyBytes)
	newTxnBytes, txnSignatureBytes, err := desoCore.SignTransactionWithDerivedKey(txnBytes, privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("problem signing transaction: %w", err)
	}

	var signedTxnHex []byte
	signedTxnHex = newTxnBytes[0 : len(newTxnBytes)-1]
	signedTxnHex = append(signedTxnHex, desoCore.UintToBuf(uint64(len(txnSignatureBytes)))...)
	signedTxnHex = append(signedTxnHex, txnSignatureBytes...)

	return hex.EncodeToString(signedTxnHex), nil
}
