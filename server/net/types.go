package net

const (
	CodeKeepAlive = "keepalive"
)

type Payload struct {
	Type string `json:"type"`
}