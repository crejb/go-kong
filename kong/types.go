package kong

import "encoding/json"

// Service represents a Service in Kong.
// Read https://getkong.org/docs/0.13.x/admin-api/#Service-object
// +k8s:deepcopy-gen=true
type Service struct {
	ClientCertificate *Certificate `json:"client_certificate,omitempty" yaml:"client_certificate,omitempty"`
	ConnectTimeout    *int         `json:"connect_timeout,omitempty" yaml:"connect_timeout,omitempty"`
	CreatedAt         *int         `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Host              *string      `json:"host,omitempty" yaml:"host,omitempty"`
	ID                *string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name              *string      `json:"name,omitempty" yaml:"name,omitempty"`
	Path              *string      `json:"path,omitempty" yaml:"path,omitempty"`
	Port              *int         `json:"port,omitempty" yaml:"port,omitempty"`
	Protocol          *string      `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	ReadTimeout       *int         `json:"read_timeout,omitempty" yaml:"read_timeout,omitempty"`
	Retries           *int         `json:"retries,omitempty" yaml:"retries,omitempty"`
	UpdatedAt         *int         `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	WriteTimeout      *int         `json:"write_timeout,omitempty" yaml:"write_timeout,omitempty"`
	Tags              []*string    `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// CIDRPort represents a set of CIDR and a port.
// +k8s:deepcopy-gen=true
type CIDRPort struct {
	IP   *string `json:"ip,omitempty" yaml:"ip,omitempty"`
	Port *int    `json:"port,omitempty" yaml:"port,omitempty"`
}

// Route represents a Route in Kong.
// Read https://getkong.org/docs/0.13.x/admin-api/#Route-object
// +k8s:deepcopy-gen=true
type Route struct {
	CreatedAt     *int                `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Hosts         []*string           `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Headers       map[string][]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	ID            *string             `json:"id,omitempty" yaml:"id,omitempty"`
	Name          *string             `json:"name,omitempty" yaml:"name,omitempty"`
	Methods       []*string           `json:"methods,omitempty" yaml:"methods,omitempty"`
	Paths         []*string           `json:"paths,omitempty" yaml:"paths,omitempty"`
	PreserveHost  *bool               `json:"preserve_host,omitempty" yaml:"preserve_host,omitempty"`
	Protocols     []*string           `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	RegexPriority *int                `json:"regex_priority,omitempty" yaml:"regex_priority,omitempty"`
	Service       *Service            `json:"service,omitempty" yaml:"service,omitempty"`
	StripPath     *bool               `json:"strip_path,omitempty" yaml:"strip_path,omitempty"`
	UpdatedAt     *int                `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	SNIs          []*string           `json:"snis,omitempty" yaml:"snis,omitempty"`
	Sources       []*CIDRPort         `json:"sources,omitempty" yaml:"sources,omitempty"`
	Destinations  []*CIDRPort         `json:"destinations,omitempty" yaml:"destinations,omitempty"`
	Tags          []*string           `json:"tags,omitempty" yaml:"tags,omitempty"`

	HTTPSRedirectStatusCode *int `json:"https_redirect_status_code,omitempty" yaml:"https_redirect_status_code,omitempty"`
}

// Consumer represents a Consumer in Kong.
// Read https://getkong.org/docs/0.13.x/admin-api/#consumer-object
// +k8s:deepcopy-gen=true
type Consumer struct {
	ID        *string   `json:"id,omitempty" yaml:"id,omitempty"`
	CustomID  *string   `json:"custom_id,omitempty" yaml:"custom_id,omitempty"`
	Username  *string   `json:"username,omitempty" yaml:"username,omitempty"`
	CreatedAt *int64    `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Tags      []*string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Certificate represents a Certificate in Kong.
// Read https://getkong.org/docs/0.14.x/admin-api/#certificate-object
// +k8s:deepcopy-gen=true
type Certificate struct {
	ID        *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Cert      *string   `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key       *string   `json:"key,omitempty" yaml:"key,omitempty"`
	CreatedAt *int64    `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	SNIs      []*string `json:"snis,omitempty" yaml:"snis,omitempty"`
	Tags      []*string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// SNI represents a SNI in Kong.
// Read https://getkong.org/docs/0.14.x/admin-api/#sni-object
// +k8s:deepcopy-gen=true
type SNI struct {
	ID          *string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name        *string      `json:"name,omitempty" yaml:"name,omitempty"`
	CreatedAt   *int64       `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	Certificate *Certificate `json:"certificate,omitempty" yaml:"certificate,omitempty"`
	Tags        []*string    `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Healthy configures thresholds and HTTP status codes
// to mark targets healthy for an upstream.
// +k8s:deepcopy-gen=true
type Healthy struct {
	HTTPStatuses []int `json:"http_statuses,omitempty" yaml:"http_statuses,omitempty"`
	Interval     *int  `json:"interval,omitempty" yaml:"interval,omitempty"`
	Successes    *int  `json:"successes,omitempty" yaml:"successes,omitempty"`
}

// Unhealthy configures thresholds and HTTP status codes
// to mark targets unhealthy.
// +k8s:deepcopy-gen=true
type Unhealthy struct {
	HTTPFailures *int  `json:"http_failures,omitempty" yaml:"http_failures,omitempty"`
	HTTPStatuses []int `json:"http_statuses,omitempty" yaml:"http_statuses,omitempty"`
	TCPFailures  *int  `json:"tcp_failures,omitempty" yaml:"tcp_failures,omitempty"`
	Timeouts     *int  `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
	Interval     *int  `json:"interval,omitempty" yaml:"interval,omitempty"`
}

// ActiveHealthcheck configures active health check probing.
// +k8s:deepcopy-gen=true
type ActiveHealthcheck struct {
	Concurrency            *int       `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	Healthy                *Healthy   `json:"healthy,omitempty" yaml:"healthy,omitempty"`
	HTTPPath               *string    `json:"http_path,omitempty" yaml:"http_path,omitempty"`
	HTTPSSni               *string    `json:"https_sni,omitempty" yaml:"https_sni,omitempty"`
	HTTPSVerifyCertificate *bool      `json:"https_verify_certificate,omitempty" yaml:"https_verify_certificate,omitempty"`
	Type                   *string    `json:"type,omitempty" yaml:"type,omitempty"`
	Timeout                *int       `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Unhealthy              *Unhealthy `json:"unhealthy,omitempty" yaml:"unhealthy,omitempty"`
}

// PassiveHealthcheck configures passive checks around
// passive health checks.
// +k8s:deepcopy-gen=true
type PassiveHealthcheck struct {
	Healthy   *Healthy   `json:"healthy,omitempty" yaml:"healthy,omitempty"`
	Unhealthy *Unhealthy `json:"unhealthy,omitempty" yaml:"unhealthy,omitempty"`
}

// Healthcheck represents a health-check config of an upstream
// in Kong.
// +k8s:deepcopy-gen=true
type Healthcheck struct {
	Active  *ActiveHealthcheck  `json:"active,omitempty" yaml:"active,omitempty"`
	Passive *PassiveHealthcheck `json:"passive,omitempty" yaml:"passive,omitempty"`
}

// Upstream represents a Consumer in Kong.
// +k8s:deepcopy-gen=true
type Upstream struct {
	ID                 *string      `json:"id,omitempty" yaml:"id,omitempty"`
	Name               *string      `json:"name,omitempty" yaml:"name,omitempty"`
	HostHeader         *string      `json:"host_header,omitempty" yaml:"host_header,omitempty"`
	Algorithm          *string      `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Slots              *int         `json:"slots,omitempty" yaml:"slots,omitempty"`
	Healthchecks       *Healthcheck `json:"healthchecks,omitempty" yaml:"healthchecks,omitempty"`
	CreatedAt          *int64       `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	HashOn             *string      `json:"hash_on,omitempty" yaml:"hash_on,omitempty"`
	HashFallback       *string      `json:"hash_fallback,omitempty" yaml:"hash_fallback,omitempty"`
	HashOnHeader       *string      `json:"hash_on_header,omitempty" yaml:"hash_on_header,omitempty"`
	HashFallbackHeader *string      `json:"hash_fallback_header,omitempty" yaml:"hash_fallback_header,omitempty"`
	HashOnCookie       *string      `json:"hash_on_cookie,omitempty" yaml:"hash_on_cookie,omitempty"`
	HashOnCookiePath   *string      `json:"hash_on_cookie_path,omitempty" yaml:"hash_on_cookie_path,omitempty"`
	Tags               []*string    `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Target represents a Target in Kong.
// +k8s:deepcopy-gen=true
type Target struct {
	CreatedAt *float64  `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	ID        *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Target    *string   `json:"target,omitempty" yaml:"target,omitempty"`
	Upstream  *Upstream `json:"upstream,omitempty" yaml:"upstream,omitempty"`
	Weight    *int      `json:"weight,omitempty" yaml:"weight,omitempty"`
	Tags      []*string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// Configuration represents a config of a plugin in Kong.
type Configuration map[string]interface{}

// DeepCopyInto copies the receiver, writing into out. in must be non-nil.
func (in Configuration) DeepCopyInto(out *Configuration) {
	// Resorting to JSON since interface{} cannot be DeepCopied easily.
	// This could be replaced using reflection-fu.
	// XXX Ignoring errors
	b, _ := json.Marshal(&in)
	_ = json.Unmarshal(b, out)
}

// DeepCopy copies the receiver, creating a new Configuration.
func (in Configuration) DeepCopy() Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return *out
}

// Plugin represents a Plugin in Kong.
// Read https://getkong.org/docs/0.13.x/admin-api/#Plugin-object
// +k8s:deepcopy-gen=true
type Plugin struct {
	CreatedAt *int          `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	ID        *string       `json:"id,omitempty" yaml:"id,omitempty"`
	Name      *string       `json:"name,omitempty" yaml:"name,omitempty"`
	Route     *Route        `json:"route,omitempty" yaml:"route,omitempty"`
	Service   *Service      `json:"service,omitempty" yaml:"service,omitempty"`
	Consumer  *Consumer     `json:"consumer,omitempty" yaml:"consumer,omitempty"`
	Config    Configuration `json:"config,omitempty" yaml:"config,omitempty"`
	Enabled   *bool         `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	RunOn     *string       `json:"run_on,omitempty" yaml:"run_on,omitempty"`
	Protocols []*string     `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	Tags      []*string     `json:"tags,omitempty" yaml:"tags,omitempty"`
}
