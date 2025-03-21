package main

type ValidatorRequest struct {
	RequestID     string   `json:"request_id"`
	NumValidators int      `json:"num_validators"`
	FeeRecipient  string   `json:"fee_recipient"`
	Status        string   `json:"status"`
	Keys          []string `json:"keys,omitempty"` // Only included if status = successful
}
