package http

import (
	"Dbot/console"
	"os"
	"sync"
	"time"
)

var (
	xsrfMutex sync.RWMutex
	xsrfToken string
)

func XSRFRefresh(interval time.Duration) {
	refresh := func() {
		console.Info("Refreshing xsrf token")
		response, err := RestyClient.R().SetBody(map[string]any{
			"logoutReason": "xsrf",
			"url":          "https://roblox.com/",
			"userId":       "10603273948",
		}).SetHeader("Cookie", os.Getenv("RANKING_COOKIE")).Post("https://auth.roblox.com/v3/logout")
		if err != nil {
			console.Warn("Failed to refresh XSRF token, this may cause some issues; %v", err)
			return
		}
		t := response.Header().Get("x-csrf-token")
		if t == "" {
			console.Warn("xsrf token returned no token... status (%v)", response.Status())
			return
		}
		xsrfMutex.Lock() // only one goroutine can fetch or write this
		xsrfToken = t
		console.Info("xsrf token is now %v", xsrfToken)
		xsrfMutex.Unlock() // k now everyone can

	}

	refresh()
	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		for range t.C {
			refresh()
		}
	}()
}
