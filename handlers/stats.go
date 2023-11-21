package handlers

import (
    "sync"
    "github.com/gofiber/fiber/v2"
)

type EndpointStats struct {
    Count           int            `json:"count"`
    UniqueUserAgent map[string]int `json:"unique_user_agent"`
}

type Stats struct {
    mu        sync.Mutex
    endpoints map[string]*EndpointStats
}

var statsMapMu sync.Mutex

func NewStats() *Stats {
    return &Stats{
        endpoints: make(map[string]*EndpointStats),
    }
}

func (s *Stats) RecordRequest(endpoint, userAgent string) {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, ok := s.endpoints[endpoint]; !ok {
        s.endpoints[endpoint] = &EndpointStats{
            UniqueUserAgent: make(map[string]int),
        }
    }

    s.endpoints[endpoint].Count++
    s.endpoints[endpoint].UniqueUserAgent[userAgent]++
}

func StatsMiddleware(statsMap map[string]*EndpointStats) func(*fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        endpoint := c.Method() + " " + c.Path()
        userAgent := c.Get("User-Agent")

        statsMapMu.Lock()
        defer statsMapMu.Unlock()

        if _, ok := statsMap[endpoint]; !ok {
            statsMap[endpoint] = &EndpointStats{
                UniqueUserAgent: make(map[string]int),
            }
        }

        statsMap[endpoint].Count++
        statsMap[endpoint].UniqueUserAgent[userAgent]++

        return c.Next()
    }
}
