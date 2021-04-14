package events

// Batch struct for Batch
type Batch struct {
	SourceRequestID string        `json:"source_request_id,omitempty"`
	BatchContext    *BatchContext `json:"context,omitempty"`
	// Provide a list of event objects - such as CustomEvent, ScreenViewEvent, or CommerceEvent
	Events                []Event                      `json:"events,omitempty"`
	DeviceInfo            *DeviceInformation           `json:"device_info,omitempty"`
	ApplicationInfo       *ApplicationInformation      `json:"application_info,omitempty"`
	UserAttributes        map[string]interface{}       `json:"user_attributes,omitempty"`
	DeletedUserAttributes []string                     `json:"deleted_user_attributes,omitempty"`
	UserIdentities        *UserIdentities              `json:"user_identities,omitempty"`
	Environment           Environment                  `json:"environment"`
	APIKey                string                       `json:"api_key,omitempty"`
	APIKeys               []string                     `json:"api_keys,omitempty"`
	IP                    string                       `json:"ip,omitempty"`
	IntegrationAttributes map[string]map[string]string `json:"integration_attributes,omitempty"`
	PartnerIdentities     map[string]string            `json:"partner_identities,omitempty"`
	SourceInfo            *SourceInformation           `json:"source_info,omitempty"`
	MpDeviceid            string                       `json:"mp_deviceid,omitempty"`
	AttributionInfo       *AttributionInfo             `json:"attribution_info,omitempty"`
	TimestampUnixtimeMS   int64                        `json:"timestamp_unixtime_ms,omitempty"`
	BatchID               int64                        `json:"batch_id,omitempty"`
	MPID                  int64                        `json:"mpid,omitempty"`
	SdkVersion            string                       `json:"sdk_version,omitempty"`
	ConsentState          *ConsentState                `json:"consent_state,omitempty"`
	JobID                 string                       `json:"job_id,omitempty"`
}

// Environment model
type Environment string

// List of CustomEventType
const (
	DevelopmentEnvironment Environment = "development"
	ProductionEnvironment  Environment = "production"
)
