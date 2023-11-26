package token

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu             sync.Mutex
	lastUpdateTime time.Time
	capacity       int // 令牌桶容量
	tokens         int // 当前令牌数量
	rate           int // 令牌生成速率 m/s
}

func NewToken(capacity, rate int) *TokenBucket {
	return &TokenBucket{
		capacity:       capacity,
		tokens:         capacity,
		rate:           rate,
		lastUpdateTime: time.Now(),
	}

}

// Allow 检查是否有可用令牌
func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	// 计算时间间隔 生成令牌
	now := time.Now()
	duration := now.Sub(t.lastUpdateTime)
	generateTokens := int(duration.Seconds()) * t.rate
	t.tokens = min(t.tokens+generateTokens, t.capacity)
	t.lastUpdateTime = now

	if t.tokens > 0 {
		t.tokens--
		return true
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
