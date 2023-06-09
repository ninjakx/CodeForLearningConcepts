import (
	"sync"
	"time"
)

type RateLimiter struct {
	visits     map[string]int
	lastVisits map[string]time.Time
	mutex      sync.Mutex
	interval   time.Duration
	limit      int
}

func NewRateLimiter(interval time.Duration, limit int) *RateLimiter {
	return &RateLimiter{
		visits:     make(map[string]int),
		lastVisits: make(map[string]time.Time),
		mutex:      sync.Mutex{},
		interval:   interval,
		limit:      limit,
	}
}

func (r *RateLimiter) AllowVisit(ip string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	now := time.Now()
	if _, ok := r.lastVisits[ip]; !ok {
		r.lastVisits[ip] = now
		r.visits[ip] = 1
		return true
	}

	if now.Sub(r.lastVisits[ip]) > r.interval {
		r.lastVisits[ip] = now
		r.visits[ip] = 1
		return true
	}

	if r.visits[ip] < r.limit {
		r.visits[ip]++
		return true
	}

	return false
}

func main() {
	limiter := NewRateLimiter(1*time.Minute, 5)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // get the IP address from the request
		if !limiter.AllowVisit(ip) {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		
		// handle the request

	})
	
	http.ListenAndServe(":8080", nil)
}
