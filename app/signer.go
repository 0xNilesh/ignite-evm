package app

import (
	"cosmossdk.io/x/tx/signing"
	evmtypes "github.com/evmos/os/x/evm/types"
)

func ProvideCustomGetSigners() []signing.CustomGetSigner {
	return []signing.CustomGetSigner{
		{
			MsgType: "os.evm.v1.MsgEthereumTx",
			Fn:      evmtypes.MsgEthereumTxCustomGetSigner.Fn,
		},
	}
}
