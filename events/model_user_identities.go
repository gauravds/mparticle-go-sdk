package events

// UserIdentities struct
type UserIdentities struct {
	Other                    string `json:"other,omitempty"`
	CustomerID               string `json:"customer_id,omitempty"`
	Facebook                 string `json:"facebook,omitempty"`
	Twitter                  string `json:"twitter,omitempty"`
	Google                   string `json:"google,omitempty"`
	Microsoft                string `json:"microsoft,omitempty"`
	Yahoo                    string `json:"yahoo,omitempty"`
	Email                    string `json:"email,omitempty"`
	Alias                    string `json:"alias,omitempty"`
	FacebookCustomAudienceID string `json:"facebook_custom_audience_id,omitempty"`
	OtherID2                 string `json:"other_id_2,omitempty"`
	OtherID3                 string `json:"other_id_3,omitempty"`
	OtherID4                 string `json:"other_id_4,omitempty"`
}
