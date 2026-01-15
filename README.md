# ecash-sdk-core (Go)

[![Go Report Card](https://goreportcard.com/badge/github.com/useeasycash/ecash-sdk-core)](https://goreportcard.com/report/github.com/useeasycash/ecash-sdk-core)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**The Official Server-Side SDK for the EasyCash Protocol.**

`ecash-sdk-core` provides a robust, type-safe Golang interface for interacting with the EasyCash Agent Network. It is designed for high-throughput, secure backend integrations such as:

*   **Exchanges & Custodians**: Offering private withdrawals for users.
*   **Payment Gateways**: Processing merchant settlements via stablecoins.
*   **Payroll Providers**: Batching private salary streams.

## 🚀 Key Features

*   **🔒 ZK-Proof Generation**: Built-in logic to generate solvency proofs locally before broadcasting.
*   **🤖 Agentic Routing**: Automatically selects the optimal execution path with multi-factor optimization (cost, speed, security).
*   **📊 Built-in Observability**: Metrics tracking for transaction success rates, latency, and fee analysis.
*   **⚡ Performance Optimized**: In-memory caching with TTL for repeated transaction patterns.
*   **🛡 Type Safety**: Strict typing for Assets, Chains, and Intent structures to prevent financial errors.
*   **✅ Comprehensive Validation**: Input validation for addresses, amounts, and chain compatibility.

## 📦 Installation

```bash
go get github.com/useeasycash/ecash-sdk-core
```

## 🛠 Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "log"

    "github.com/useeasycash/ecash-sdk-core/pkg/client"
    "github.com/useeasycash/ecash-sdk-core/pkg/config"
    "github.com/useeasycash/ecash-sdk-core/pkg/types"
)

func main() {
    // Initialize with default config
    cfg := config.DefaultConfig()
    cfg.APIKey = "your_api_key_here"
    
    sdk, err := client.NewClient(cfg)
    if err != nil {
        log.Fatal(err)
    }

    // Execute a private transfer
    req := &types.TransactionRequest{
        Type:        types.IntentTransfer,
        Amount:      "1000.00",
        Asset:       "USDC",
        Recipient:   "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
        SourceChain: types.ChainBase,
        IsShielded:  true, // Enable ZK Privacy
    }

    resp, err := sdk.ExecuteTransaction(context.Background(), req)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Transaction confirmed: %s", resp.TxHash)
}
```

### Advanced Configuration

```go
cfg := &config.SDKConfig{
    APIKey:         os.Getenv("ECASH_API_KEY"),
    Environment:    "mainnet",
    Timeout:        30 * time.Second,
    EnableZKProofs: true,
    EnableMetrics:  true,
    EnableCaching:  true,
    CacheTTL:       5 * time.Minute,
}

sdk, _ := client.NewClient(cfg)
```

### Monitoring & Metrics

```go
// Get performance metrics
metrics := sdk.GetMetrics()
fmt.Printf("Success Rate: %.2f%%\n", metrics["success_rate"].(float64) * 100)
fmt.Printf("Average Latency: %dms\n", metrics["average_latency_ms"])
```

## 🏗 Architecture

```
pkg/
├── agent/          # Route negotiation & quote selection
├── cache/          # In-memory caching with TTL
├── client/         # Main SDK client interface
├── config/         # Configuration management
├── crypto/         # Cryptographic signing utilities
├── errors/         # Structured error handling
├── monitoring/     # Metrics & observability
├── types/          # Domain models & types
├── validator/      # Input validation
└── zk/             # Zero-Knowledge proof generation
```

### Package Overview

*   **`client`**: High-level facade for API interaction with automatic route optimization.
*   **`agent`**: Negotiates with the decentralized agent network to find optimal execution paths.
*   **`zk`**: Zero-Knowledge cryptographic primitives and proof generation logic.
*   **`monitoring`**: Real-time metrics collection for performance analysis.
*   **`cache`**: Performance optimization through intelligent caching.
*   **`validator`**: Comprehensive input validation to prevent errors.
*   **`config`**: Environment-aware configuration with sensible defaults.
*   **`errors`**: Structured error codes for better error handling.

## 🧪 Testing

Run the test suite:

```bash
make test
```

Run a specific package test:

```bash
go test ./pkg/validator -v
```

## 📖 Examples

Check out the `examples/` directory for complete working examples:

*   **`simple_transfer/`**: Basic private transfer with ZK proofs
*   More examples coming soon...

## 🔧 Development

```bash
# Install dependencies
make deps

# Run linter
make lint

# Build
make build

# Run example
make example
```

## 🌐 Environment Variables

```bash
ECASH_API_KEY=your_api_key_here
ECASH_API_ENDPOINT=https://api.useeasy.cash
ECASH_ENV=mainnet  # or testnet, devnet
```

## 🤝 Contributing

We welcome Pull Requests! Please ensure you run `make test` and `make lint` before submitting.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.

## 🔗 Links

*   [EasyCash Website](https://useeasy.cash)
*   [Documentation](https://useeasycash.gitbook.io/ecash-docs/)
*   [Twitter](https://x.com/useeasycash)
*   [GitHub Organization](https://github.com/useeasycash)
