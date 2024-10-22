package client

import (
	"encoding/json"
	"fmt"
	"github.com/ben-haas/climate-compare/backend/internal/config"
	"net/http"
	"time"
)

// MonthlyData represents the response structure from Meteostat API
type MonthlyData struct {
	Date string  `json:"date"`
	Tavg float64 `json:"tavg"`
	Tmin float64 `json:"tmin"`
	Tmax float64 `json:"tmax"`
	Prcp float64 `json:"prcp"`
	Snow int     `json:"snow"`
	Wdir int     `json:"wdir"`
	Wspd float64 `json:"wspd"`
	Wpgt float64 `json:"wpgt"`
	Pres float64 `json:"pres"`
	Tsun int     `json:"tsun"`
}

// Client handles API requests to Meteostat
type Client struct {
	config     *config.Config
	httpClient *http.Client
}

// NewClient creates a new Meteostat API client using the provided config
func NewClient(cfg *config.Config) *Client {
	return &Client{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetMonthlyData fetches monthly weather data for a specific location and time period
func (c *Client) GetMonthlyData(lat, lon float64, start, end time.Time) ([]MonthlyData, error) {
	url := fmt.Sprintf("%s/point/monthly?lat=%f&lon=%f&start=%s&end=%s",
		c.config.MeteoBaseUrl,
		lat,
		lon,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// Set required headers using API key from config
	req.Header.Add("x-rapidapi-host", "meteostat.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", c.config.MeteoApiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var data []MonthlyData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return data, nil
}
