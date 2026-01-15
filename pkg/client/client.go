package client

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/useeasycash/ecash-sdk-core/pkg/agent"
	"github.com/useeasycash/ecash-sdk-core/pkg/cache"
	"github.com/useeasycash/ecash-sdk-core/pkg/config"
	sdkerrors "github.com/useeasycash/ecash-sdk-core/pkg/errors"
	"github.com/useeasycash/ecash-sdk-core/pkg/monitoring"
	"github.com/useeasycash/ecash-sdk-core/pkg/types"
	"github.com/useeasycash/ecash-sdk-core/pkg/validator"
	"github.com/useeasycash/ecash-sdk-core/pkg/zk"
)

// EasyCashClient is the main entry point for the SDK
type EasyCashClient struct {
	config     *config.SDKConfig
	zk         *zk.ProofGenerator
	negotiator *agent.AgentNegotiator
	cache      *cache.Cache
	metrics    *monitoring.Metrics
}

// NewClient initializes a new EasyCash SDK client with full configuration
func NewClient(cfg *config.SDKConfig) (*EasyCashClient, error) {
	if cfg == nil {
		cfg = config.DefaultConfig()
	}

	if err := cfg.Validate(); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid configuration", err)
	}

	client := &EasyCashClient{
		config:     cfg,
		zk:         zk.NewProofGenerator("./circuits/spend.wasm"),
		negotiator: agent.NewNegotiator(cfg.Timeout),
		metrics:    monitoring.GetMetrics(),
	}

	if cfg.EnableCaching {
		client.cache = cache.NewCache(cfg.CacheTTL)
	}

	return client, nil
}

// ExecuteTransaction constructs a transfer intent and executes it with full validation
func (c *EasyCashClient) ExecuteTransaction(ctx context.Context, req *types.TransactionRequest) (*types.TransactionResponse, error) {
	startTime := time.Now()
	var success bool
	var fee float64

	defer func() {
		if c.config.EnableMetrics {
			c.metrics.RecordTransaction(success, fee, time.Since(startTime))
		}
	}()

	// 1. Validate Request
	if err := validator.ValidateTransactionRequest(req); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "validation failed", err)
	}

	// 2. Check Cache for similar recent transactions
	if c.config.EnableCaching && c.cache != nil {
		cacheKey := fmt.Sprintf("%s-%s-%s", req.Type, req.Amount, req.Asset)
		if cached, found := c.cache.Get(cacheKey); found {
			fmt.Println("[SDK] Cache hit for transaction pattern")
			if resp, ok := cached.(*types.TransactionResponse); ok {
				success = true
				return resp, nil
			}
		}
	}

	// 3. Generate ZK Proof if shielded
	if c.config.EnableZKProofs && req.IsShielded {
		proof, err := c.zk.GenerateSolvencyProof(req.Amount, "0")
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrProofGeneration, "failed to generate privacy proof", err)
		}
		fmt.Printf("[SDK] Generated ZK Proof: %s...\n", proof[:10])
	}

	// 4. Request quotes from agents
	quotes, err := c.negotiator.RequestQuotes(ctx, req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrAgentUnavailable, "failed to get agent quotes", err)
	}

	// 5. Select best route
	bestRoute, err := c.negotiator.SelectBestRoute(quotes, "balanced")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrAgentUnavailable, "no suitable route found", err)
	}

	fmt.Printf("[SDK] Selected Agent: %s (Fee: %s, Security: %.2f)\n",
		bestRoute.AgentID, bestRoute.EstimatedFee, bestRoute.SecurityScore)

	// 6. Execute via selected agent (simulated)
	select {
	case <-ctx.Done():
		return nil, sdkerrors.Wrap(sdkerrors.ErrTimeout, "transaction timeout", ctx.Err())
	case <-time.After(100 * time.Millisecond):
		// Simulated execution latency
	}

	// 7. Construct Response
	txHash := uuid.New().String()
	fee = 0.05 // Simulated fee
	success = true

	resp := &types.TransactionResponse{
		TxHash:      "0x" + txHash,
		Status:      "confirmed",
		BlockHeight: 1948201,
		FeeUsed:     bestRoute.EstimatedFee,
	}

	// 8. Cache successful result
	if c.config.EnableCaching && c.cache != nil {
		cacheKey := fmt.Sprintf("%s-%s-%s", req.Type, req.Amount, req.Asset)
		c.cache.Set(cacheKey, resp)
	}

	return resp, nil
}

// GetMetrics returns current SDK performance metrics
func (c *EasyCashClient) GetMetrics() map[string]interface{} {
	if !c.config.EnableMetrics {
		return map[string]interface{}{"metrics_disabled": true}
	}
	return c.metrics.GetStats()
}
