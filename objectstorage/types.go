package objectstorage

import "github.com/ekaputra07/warren-go/api"

// S3Bucket represents Object Storage bucket
type S3Bucket struct {
	Name             string `json:"name"`
	SizeBytes        int    `json:"size_bytes"`
	BillingAccountID int    `json:"billing_account_id"`
	NumObjects       int    `json:"num_objects"`
	CreatedAt        string `json:"created_at"`
	ModifiedAt       string `json:"modified_at"`
	IsSuspended      bool   `json:"is_suspended"`
}

// S3Credential holds information about user credentials that can be used to access S3 buckets and objects
type S3Credential struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	UserID    string `json:"userId"`
}

// S3UserInfo holds some information about user
// `caps`, `subusers` and `swiftCredentials` are ommited (documentation not clear)
type S3UserInfo struct {
	DisplayName   string
	Email         string
	MaxBuckets    int
	S3Credentials []S3Credential
	Suspended     int
	UserID        string
}

// Client is object storage client
type Client struct {
	BillingAccountID int
	API              *api.API
}
