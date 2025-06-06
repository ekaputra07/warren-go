package lb

import "github.com/google/uuid"

type RuleSettings struct {
	ConnectionLimit    int    `json:"connection_limit"`
	SessionPersistence string `json:"session_persistence"`
}

type Rule struct {
	UUID       uuid.UUID    `json:"uuid"`
	Protocol   string       `json:"protocol"`
	CreatedAt  string       `json:"created_at"`
	SourcePort int          `json:"source_port"`
	TargetPort int          `json:"target_port"`
	Settings   RuleSettings `json:"settings"`
}

type Target struct {
	UUID      uuid.UUID `json:"target_uuid"`
	Type      string    `json:"target_type"`
	IPAddress string    `json:"target_ip_address"`
}

type LoadBalancer struct {
	UUID             uuid.UUID `json:"uuid"`
	NetworkUUID      uuid.UUID `json:"network_uuid"`
	UserID           int       `json:"user_id"`
	BillingAccountID int       `json:"billing_account_id"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
	PrivateAddress   string    `json:"private_address"`
	IsDeleted        bool      `json:"is_deleted"`
	ForwardingRules  []Rule    `json:"forwarding_rules"`
	Targets          []Target  `json:"targets"`
}
