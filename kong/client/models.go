package client

type KongService struct {
	Id             string `json:"id,omitempty"`
	CreatedAt      int64  `json:"created_at,omitempty"`
	UpdatedAt      int64  `json:"updated_at,omitempty"`
	ConnectTimeout int    `json:"connect_timeout,omitempty"`
	Name           string `json:"name,omitempty"`
	Retries        int    `json:"retries,omitempty"`
	ReadTimeout    int    `json:"read_timeout,omitempty"`
	WriteTimeout   int    `json:"write_timeout,omitempty"`

	// Kong's api treats `url` as a write-only property.
	// This is useful for creating or updating a service (simply supply the url instead of four other fields),
	// However, in the interest of a consistent model, this package only exposes a url field, for both reading and writing.
	// The other fields (protocol, host, port, path) are in the struct only so that the url field can be populated.
	Url      string `json:"url,omitempty"`
	protocol string
	host     string
	port     int
	path     string
}

type KongServiceReference struct {
	Id string `json:"id"`
}

type KongRoute struct {
	Id            string               `json:"id,omitempty"`
	CreatedAt     int64                `json:"created_at,omitempty"`
	UpdatedAt     int64                `json:"updated_at,omitempty"`
	Protocols     []string             `json:"protocols,omitempty"`
	Methods       []string             `json:"methods,omitempty"`
	Hosts         []string             `json:"hosts,omitempty"`
	Paths         []string             `json:"paths,omitempty"`
	RegexPriority int                  `json:"regex_priority"`
	StripPath     bool                 `json:"strip_path"`
	PreserveHost  bool                 `json:"preserve_host"`
	Service       KongServiceReference `json:"service"`
}

type KongPlugin struct {
	Id         string
	ServiceId  string
	ConsumerId string
	Name       string
	Config     map[string]interface{}
	Enabled    bool
	CreatedAt  int64
}