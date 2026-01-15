package validator

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/useeasy/ecash-sdk-core/pkg/types"
)

var (
	addressRegex = regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	amountRegex  = regexp.MustCompile(`^\d+(\.\d+)?$`)
)

// ValidateAddress checks if an address is valid
func ValidateAddress(address string) error {
	if !addressRegex.MatchString(address) {
		return fmt.Errorf("invalid address format: %s", address)
	}
	return nil
}

// ValidateAmount checks if an amount string is valid
func ValidateAmount(amount string) error {
	if !amountRegex.MatchString(amount) {
		return fmt.Errorf("invalid amount format: %s", amount)
	}

	// Check if amount is positive
	val, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("failed to parse amount: %w", err)
	}

	if val <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	return nil
}

// ValidateChain checks if a chain ID is supported
func ValidateChain(chain types.ChainID) error {
	validChains := map[types.ChainID]bool{
		types.ChainEthereum: true,
		types.ChainBase:     true,
		types.ChainSolana:   true,
	}

	if !validChains[chain] {
		return fmt.Errorf("unsupported chain: %s", chain)
	}

	return nil
}

// ValidateTransactionRequest performs comprehensive validation
func ValidateTransactionRequest(req *types.TransactionRequest) error {
	if err := ValidateAmount(req.Amount); err != nil {
		return fmt.Errorf("amount validation failed: %w", err)
	}

	if err := ValidateChain(req.SourceChain); err != nil {
		return fmt.Errorf("source chain validation failed: %w", err)
	}

	if req.TargetChain != "" {
		if err := ValidateChain(req.TargetChain); err != nil {
			return fmt.Errorf("target chain validation failed: %w", err)
		}
	}

	if req.Recipient != "" {
		if err := ValidateAddress(req.Recipient); err != nil {
			return fmt.Errorf("recipient validation failed: %w", err)
		}
	}

	return nil
}
