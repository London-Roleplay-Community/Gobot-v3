package http

import (
	"Dbot/console"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	OpenCloud *resty.Client
	UsersV1   *resty.Client
	GroupsV2  *resty.Client
	GroupsV1  *resty.Client
	AuthV3    *resty.Client
)

func BuildClients() { // this should fix headers not being sent?
	OpenCloud = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"x-api-key":    os.Getenv("RBLX_APIKEY"),
		"Cookie":       os.Getenv("RANKING_COOKIE"),
	}).SetBaseURL("https://apis.roblox.com/cloud/v2").OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		xsrfMutex.RLock()
		token := xsrfToken
		xsrfMutex.RUnlock()
		if token != "" {
			r.SetHeader("x-csrf-token", token)
		}
		return nil
	}).AddRetryCondition(func(r *resty.Response, err error) bool {

		if err != nil {
			return false
		}
		if r.StatusCode() == 403 {
			if t := r.Header().Get("x-csrf-token"); t != "" {
				console.Info("403 returned, will try again with xsrf token")

				xsrfMutex.Lock()
				xsrfToken = t
				xsrfMutex.Unlock()
				return true
			}
		}
		return r.StatusCode() >= 500
	})

	UsersV1 = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"x-api-key":    os.Getenv("RBLX_APIKEY"),
		"Cookie":       os.Getenv("RANKING_COOKIE"),
	}).SetBaseURL("https://users.roblox.com/v1").OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		xsrfMutex.RLock()
		token := xsrfToken
		xsrfMutex.RUnlock()
		if token != "" {
			r.SetHeader("x-csrf-token", token)
		}
		return nil
	}).AddRetryCondition(func(r *resty.Response, err error) bool {

		if err != nil {
			return false
		}
		if r.StatusCode() == 403 {
			if t := r.Header().Get("x-csrf-token"); t != "" {
				console.Info("403 returned, will try again with xsrf token")

				xsrfMutex.Lock()
				xsrfToken = t
				xsrfMutex.Unlock()
				return true
			}
		}
		return r.StatusCode() >= 500
	})

	GroupsV2 = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"x-api-key":    os.Getenv("RBLX_APIKEY"),
		"Cookie":       os.Getenv("RANKING_COOKIE"),
	}).SetBaseURL("https://groups.roblox.com/v2").OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		xsrfMutex.RLock()
		token := xsrfToken
		xsrfMutex.RUnlock()
		if token != "" {
			r.SetHeader("x-csrf-token", token)
		}
		return nil
	}).AddRetryCondition(func(r *resty.Response, err error) bool {
		if err != nil {
			return false
		}
		if r.StatusCode() == 403 {
			if t := r.Header().Get("x-csrf-token"); t != "" {
				console.Info("403 returned, will try again with xsrf token")

				xsrfMutex.Lock()
				xsrfToken = t
				xsrfMutex.Unlock()
				return true
			}
		}
		return r.StatusCode() >= 500
	})

	GroupsV1 = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2).SetHeaders(map[string]string{ // DELETE https://groups.roblox.com/v1/groups/{group_id}/users/{userId}
		"Content-Type": "application/json",
		"x-api-key":    os.Getenv("RBLX_APIKEY"),
		"Cookie":       os.Getenv("RANKING_COOKIE"),
	}).SetBaseURL("https://groups.roblox.com/v1").OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		xsrfMutex.RLock()
		token := xsrfToken
		xsrfMutex.RUnlock()
		if token != "" {
			r.SetHeader("x-csrf-token", token)
		}
		return nil
	}).AddRetryCondition(func(r *resty.Response, err error) bool {
		if err != nil {
			return false
		}
		if r.StatusCode() == 403 {
			if t := r.Header().Get("x-csrf-token"); t != "" {
				console.Info("403 returned, will try again with xsrf token")
				xsrfMutex.Lock()
				xsrfToken = t
				xsrfMutex.Unlock()
				return true
			}
		}
		return r.StatusCode() >= 500
	})

	AuthV3 = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2).SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"x-api-key":    os.Getenv("RBLX_APIKEY"),
		"Cookie":       os.Getenv("RANKING_COOKIE"),
	}).SetBaseURL("https://auth.roblox.com/v3").OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		xsrfMutex.RLock()
		token := xsrfToken
		xsrfMutex.RUnlock()
		if token != "" {
			r.SetHeader("x-csrf-token", token)
		}
		return nil
	}).AddRetryCondition(func(r *resty.Response, err error) bool {
		if err != nil {
			return false
		}
		if r.StatusCode() == 403 {
			if t := r.Header().Get("x-csrf-token"); t != "" {
				console.Info("403 returned, will try again with xsrf token")
				xsrfMutex.Lock()
				xsrfToken = t
				xsrfMutex.Unlock()
				return true
			}
		}
		return r.StatusCode() >= 500
	})

	XSRFRefresh(10 * time.Minute)
}
