package types

import "errors"

// ChainID represents supported blockchain networks
type ChainID string

const (
	ChainEthereum ChainID = "ethereum"
	ChainBase     ChainID = "base"
	ChainSolana   ChainID = "solana"
)

// IntentType defines the classification of the operation
type IntentType string

const (
	IntentTransfer IntentType = "transfer"
	IntentSwap     IntentType = "swap"
	IntentShield   IntentType = "shield"
)

// TransactionRequest is the standard payload for initiating an operation
type TransactionRequest struct {
	ReferenceID string     `json:"reference_id"`
	Type        IntentType `json:"type"`
	Amount      string     `json:"amount"` // String to support big integers
	Asset       string     `json:"asset"`  // e.g., "USDC"
	Recipient   string     `json:"recipient,omitempty"`
	SourceChain ChainID    `json:"source_chain"`
	TargetChain ChainID    `json:"target_chain,omitempty"`
	// Privacy options
	IsShielded bool `json:"is_shielded"`
}

// Validation methods
func (r *TransactionRequest) Validate() error {
	if r.Amount == "" {
		return errors.New("amount is required")
	}
	if r.Asset == "" {
		return errors.New("asset is required")
	}
	return nil
}

// TransactionResponse is the result of an intent execution
type TransactionResponse struct {
	TxHash      string `json:"tx_hash"`
	Status      string `json:"status"`
	BlockHeight uint64 `json:"block_height"`
	FeeUsed     string `json:"fee_used"`
}
