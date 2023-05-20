package net

type Payload struct {
	Type string `json:"type"`
}

type RemoteExecPayload struct {
	Payload
	Command string `json:"command"`
}

const (
	RemoteExecPayloadType = "remote"
)

const KnownAddressesPath = "./known_addresses.txt"