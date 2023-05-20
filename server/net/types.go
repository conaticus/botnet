package net

type Payload struct {
	Type string `json:"type"`
}

const KnownAddressesPath = "./known_addresses.txt"