package handlers

import (
    "sync"
    "github.com/gofiber/fiber/v2"
)

// EndpointStats represents statistics for a specific HTTP request endpoint.
type EndpointStats struct {
    Count           int            `json:"count"`              // Number of requests made to this endpoint.
    UniqueUserAgent map[string]int `json:"unique_user_agent"`  // Count of unique User-Agents that accessed this endpoint.
}

// Stats represents a collection of statistics for various HTTP request endpoints.
type Stats struct {
    mu        sync.Mutex           // Mutex for protecting concurrent access to the 'endpoints' map.
    endpoints map[string]*EndpointStats // A map of endpoint paths to their statistics.
}

var statsMapMu sync.Mutex

// NewStats creates and initializes a new 'Stats' instance.
func NewStats() *Stats {
    return &Stats{
        endpoints: make(map[string]*EndpointStats),
    }
}

// RecordRequest records statistics for an HTTP request to a specific endpoint with a given User-Agent.
func (s *Stats) RecordRequest(endpoint, userAgent string) {
    s.mu.Lock()   // Lock to ensure thread safety while updating 'endpoints' map.
    defer s.mu.Unlock()

    // If the endpoint doesn't exist in the 'endpoints' map, create it.
    if _, ok := s.endpoints[endpoint]; !ok {
        s.endpoints[endpoint] = &EndpointStats{
            UniqueUserAgent: make(map[string]int),
        }
    }

    // Increment the request count and User-Agent count for the specified endpoint.
    s.endpoints[endpoint].Count++
    s.endpoints[endpoint].UniqueUserAgent[userAgent]++
}

// StatsMiddleware is a Fiber middleware function that records statistics for incoming HTTP requests.
func StatsMiddleware(statsMap map[string]*EndpointStats) func(*fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        endpoint := c.Method() + " " + c.Path()   // Generate a unique endpoint identifier.
        userAgent := c.Get("User-Agent")          // Get the User-Agent header from the request.

        statsMapMu.Lock()   // Lock to ensure thread safety while updating 'statsMap'.
        defer statsMapMu.Unlock()

        // If the endpoint doesn't exist in the 'statsMap' map, create it.
        if _, ok := statsMap[endpoint]; !ok {
            statsMap[endpoint] = &EndpointStats{
                UniqueUserAgent: make(map[string]int),
            }
        }

        // Increment the request count and User-Agent count for the specified endpoint in 'statsMap'.
        statsMap[endpoint].Count++
        statsMap[endpoint].UniqueUserAgent[userAgent]++

        return c.Next()   // Continue processing the request.
    }
}
