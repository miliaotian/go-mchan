package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	// MchanBech32Prefix defines the Bech32 prefix of an account's address
	MchanBech32Prefix = "mchan"

	// MchanCoinType Mchan coin in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	MchanCoinType = 118

	// MchanFullFundraiserPath BIP44Prefix is the parts of the BIP44 HD path that are fixed by what we used during the fundraiser.
	// use the registered cosmos stake token ATOM 118 as coin_type
	// m / purpose' / coin_type' / account' / change / address_index
	MchanFullFundraiserPath = "44'/118'/0'/0/0"

	// MchanBech32PrefixAccAddr defines the Bech32 prefix of an account's address
	MchanBech32PrefixAccAddr = MchanBech32Prefix
	// MchanBech32PrefixAccPub defines the Bech32 prefix of an account's public key
	MchanBech32PrefixAccPub = MchanBech32Prefix + sdk.PrefixPublic
	// MchanBech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	MchanBech32PrefixValAddr = MchanBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// MchanBech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	MchanBech32PrefixValPub = MchanBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// MchanBech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	MchanBech32PrefixConsAddr = MchanBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// MchanBech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	MchanBech32PrefixConsPub = MchanBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)
