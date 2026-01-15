package monitoring

import (
	"sync"
	"time"
)

// Metrics tracks SDK performance and usage statistics
type Metrics struct {
	mu                     sync.RWMutex
	TotalTransactions      int64
	SuccessfulTransactions int64
	FailedTransactions     int64
	TotalFeePaid           float64
	AverageLatency         time.Duration
}

var globalMetrics = &Metrics{}

// GetMetrics returns the global metrics instance
func GetMetrics() *Metrics {
	return globalMetrics
}

// RecordTransaction records a transaction attempt
func (m *Metrics) RecordTransaction(success bool, fee float64, latency time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.TotalTransactions++
	if success {
		m.SuccessfulTransactions++
		m.TotalFeePaid += fee
	} else {
		m.FailedTransactions++
	}

	// Update rolling average latency
	m.AverageLatency = (m.AverageLatency + latency) / 2
}

// GetStats returns current statistics
func (m *Metrics) GetStats() map[string]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return map[string]interface{}{
		"total_transactions":      m.TotalTransactions,
		"successful_transactions": m.SuccessfulTransactions,
		"failed_transactions":     m.FailedTransactions,
		"total_fee_paid":          m.TotalFeePaid,
		"average_latency_ms":      m.AverageLatency.Milliseconds(),
		"success_rate":            float64(m.SuccessfulTransactions) / float64(m.TotalTransactions),
	}
}

// Reset clears all metrics (useful for testing)
func (m *Metrics) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.TotalTransactions = 0
	m.SuccessfulTransactions = 0
	m.FailedTransactions = 0
	m.TotalFeePaid = 0
	m.AverageLatency = 0
}
