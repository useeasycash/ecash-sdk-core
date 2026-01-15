package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/useeasy/ecash-sdk-core/pkg/client"
	"github.com/useeasy/ecash-sdk-core/pkg/config"
	"github.com/useeasy/ecash-sdk-core/pkg/types"
)

func main() {
	// 1. Initialize Client with custom config
	cfg := config.DefaultConfig()
	cfg.EnableMetrics = true
	cfg.EnableCaching = true
	cfg.EnableZKProofs = true

	sdk, err := client.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize SDK: %v", err)
	}

	fmt.Println("🚀 EasyCash SDK Initialized (Advanced Mode)")

	// 2. Define a Shielded Transfer Request
	req := &types.TransactionRequest{
		ReferenceID: "ref_pay_salary_001",
		Type:        types.IntentTransfer,
		Amount:      "5000.00",
		Asset:       "USDC",
		Recipient:   "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
		SourceChain: types.ChainBase,
		IsShielded:  true, // Enable ZK Privacy
	}

	// 3. Execute
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Printf("Processing Transfer: %s %s (Shielded: %v)...\n", req.Amount, req.Asset, req.IsShielded)

	resp, err := sdk.ExecuteTransaction(ctx, req)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	fmt.Printf("✅ Success! Tx Hash: %s\n", resp.TxHash)
	fmt.Printf("   Block: %d\n", resp.BlockHeight)
	fmt.Printf("   Fee: %s\n", resp.FeeUsed)

	// 4. Display Metrics
	fmt.Println("\n📊 SDK Metrics:")
	metrics := sdk.GetMetrics()
	for key, value := range metrics {
		fmt.Printf("   %s: %v\n", key, value)
	}
}
