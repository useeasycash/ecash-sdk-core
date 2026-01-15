package agent

import (
	"context"
	"fmt"
	"time"

	"github.com/useeasy/ecash-sdk-core/pkg/types"
)

// AgentNegotiator handles fee negotiation and route selection with the Agent Network
type AgentNegotiator struct {
	timeout time.Duration
}

func NewNegotiator(timeout time.Duration) *AgentNegotiator {
	return &AgentNegotiator{
		timeout: timeout,
	}
}

// RouteQuote represents a quote from an agent for executing a transaction
type RouteQuote struct {
	AgentID       string
	EstimatedFee  string
	EstimatedTime time.Duration
	Route         []string // Chain hops
	SecurityScore float64  // 0.0 - 1.0
}

// RequestQuotes fetches multiple route quotes from available agents
func (n *AgentNegotiator) RequestQuotes(ctx context.Context, req *types.TransactionRequest) ([]RouteQuote, error) {
	// Simulate network call to Agent Discovery Service
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(50 * time.Millisecond):
		// Mock latency
	}

	// Return simulated quotes
	quotes := []RouteQuote{
		{
			AgentID:       "agent-001",
			EstimatedFee:  "0.05 USDC",
			EstimatedTime: 15 * time.Second,
			Route:         []string{string(req.SourceChain), string(req.TargetChain)},
			SecurityScore: 0.98,
		},
		{
			AgentID:       "agent-002",
			EstimatedFee:  "0.03 USDC",
			EstimatedTime: 30 * time.Second,
			Route:         []string{string(req.SourceChain), "polygon", string(req.TargetChain)},
			SecurityScore: 0.85,
		},
	}

	return quotes, nil
}

// SelectBestRoute applies multi-factor optimization to choose the best agent
func (n *AgentNegotiator) SelectBestRoute(quotes []RouteQuote, preference string) (*RouteQuote, error) {
	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes available")
	}

	// Simple heuristic: prefer high security score
	best := &quotes[0]
	for i := range quotes {
		if quotes[i].SecurityScore > best.SecurityScore {
			best = &quotes[i]
		}
	}

	return best, nil
}
