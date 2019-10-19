package client

import "errors"

var (
	errInvalidClient = errors.New("invalid client: nil client provided")
)

type Measure struct {
	Metric string `json:"metric"`
}
