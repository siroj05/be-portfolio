package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/siroj05/portfolio/internal/response"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type rateLimiterStore struct {
	mu      sync.Mutex
	clients map[string]*client
	rate    rate.Limit
	burst   int
}

func NewRateLimiterStore(interval time.Duration, b int) *rateLimiterStore {
	rls := &rateLimiterStore{
		clients: make(map[string]*client),
		rate:    rate.Every(interval),
		burst:   b,
	}

	// cleanup background
	go func() {
		for {
			time.Sleep(time.Minute)
			rls.mu.Lock()
			for ip, c := range rls.clients {
				if time.Since(c.lastSeen) > 3*time.Minute {
					delete(rls.clients, ip)
				}
			}
			rls.mu.Unlock()
		}
	}()

	return rls
}

func (rls *rateLimiterStore) getClient(ip string) *rate.Limiter {
	rls.mu.Lock()
	defer rls.mu.Unlock()

	c, exists := rls.clients[ip]
	if !exists {
		// rate.NewLimiter(rate.Every(200ms), 5)
		// artinya max 5 request burst, lalu 1 request tiap 200ms (~5 req/detik)
		limiter := rate.NewLimiter(rate.Every(200*time.Millisecond), 20)
		rls.clients[ip] = &client{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	c.lastSeen = time.Now()
	return c.limiter
}

func (rls *rateLimiterStore) RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Unable to parse IP", err.Error())
			return
		}

		limiter := rls.getClient(ip)
		if limiter == nil {
			response.Error(w, http.StatusInternalServerError, "Limiter not initialized", "")
			return
		}

		if !limiter.Allow() {
			response.Error(w, http.StatusTooManyRequests, "Too many requests", "")
			return
		}

		next.ServeHTTP(w, r)
	})
}
