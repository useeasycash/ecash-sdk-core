package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/useeasycash/ecash-sdk-core/pkg/types"
)

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		wantErr bool
	}{
		{"valid address", "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb", false},
		{"invalid prefix", "742d35Cc6634C0532925a3b844Bc9e7595f0bEb", true},
		{"invalid length", "0x742d35Cc", true},
		{"invalid chars", "0xZZZZ35Cc6634C0532925a3b844Bc9e7595f0bEb", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.address)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  string
		wantErr bool
	}{
		{"valid integer", "100", false},
		{"valid decimal", "100.50", false},
		{"zero", "0", true},
		{"negative", "-100", true},
		{"invalid format", "abc", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAmount(tt.amount)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateChain(t *testing.T) {
	assert.NoError(t, ValidateChain(types.ChainBase))
	assert.NoError(t, ValidateChain(types.ChainEthereum))
	assert.Error(t, ValidateChain(types.ChainID("invalid")))
}
