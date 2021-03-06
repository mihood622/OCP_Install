package alibabacloud

// Metadata holds additional metadata for InstallConfig resources that
// does not need to be user-supplied (e.g. because it can be retrieved
// from external APIs).
type Metadata struct {
	client *Client
	Region string `json:"region,omitempty"`
}

// NewMetadata initializes a new Metadata object.
func NewMetadata(region string) *Metadata {
	return &Metadata{Region: region}
}

// Client returns a client used for making API calls to Alibaba Cloud services.
func (m *Metadata) Client() (*Client, error) {
	if m.client == nil {
		client, err := NewClient(m.Region)
		if err != nil {
			return nil, err
		}
		m.client = client
	}
	return m.client, nil
}
