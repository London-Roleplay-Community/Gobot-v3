package http

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var RestyClient = resty.New().SetTimeout(5 * time.Second).SetRetryCount(2)
