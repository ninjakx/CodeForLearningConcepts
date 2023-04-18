import (
    "github.com/go-redis/redis"
    "time"
)

type RateLimiter struct {
    client *redis.Client
    limit  int
    interval time.Duration
}

func NewRateLimiter(client *redis.Client, limit int, interval time.Duration) *RateLimiter {
    return &RateLimiter{
        client:   client,
        limit:    limit,
        interval: interval,
    }
}

func (rl *RateLimiter) AllowVisit(ip string) (bool, error) {
    key := "ratelimit:" + ip
    count, err := rl.client.Incr(key).Result()
    if err != nil {
        return false, err
    }
    if count == 1 {
        err = rl.client.Expire(key, rl.interval).Err()
        if err != nil {
            return false, err
        }
    }
    if count > rl.limit {
        return false, nil
    }
    return true, nil
}

import (
    "github.com/go-redis/redis"
)

func main() {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    defer client.Close()

    // ...
}
