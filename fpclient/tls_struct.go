package fpclient

type TlsFingerprint struct {
	IP          string `json:"ip"`
	HTTPVersion string `json:"http_version"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	TLS         TLS    `json:"tls"`
	Http2       Http2  `json:"http2"`
}

type Http2 struct {
	AkamaiFingerprint     string      `json:"akamai_fingerprint"`
	AkamaiFingerprintHash string      `json:"akamai_fingerprint_hash"`
	SentFrames            []SentFrame `json:"sent_frames"`
}

type SentFrame struct {
	FrameType string    `json:"frame_type"`
	Length    int64     `json:"length"`
	Settings  []string  `json:"settings,omitempty"`
	Increment *int64    `json:"increment,omitempty"`
	StreamID  *int64    `json:"stream_id,omitempty"`
	Headers   []string  `json:"headers,omitempty"`
	Flags     []string  `json:"flags,omitempty"`
	Priority  *Priority `json:"priority,omitempty"`
}

type Priority struct {
	Weight    int64 `json:"weight"`
	DependsOn int64 `json:"depends_on"`
	Exclusive int64 `json:"exclusive"`
}

type TLS struct {
	Ciphers              []string    `json:"ciphers"`
	Extensions           []Extension `json:"extensions"`
	TLSVersionRecord     string      `json:"tls_version_record"`
	TLSVersionNegotiated string      `json:"tls_version_negotiated"`
	Ja3                  string      `json:"ja3"`
	Ja3Hash              string      `json:"ja3_hash"`
	Peetprint            string      `json:"peetprint"`
	PeetprintHash        string      `json:"peetprint_hash"`
	ClientRandom         string      `json:"client_random"`
	SessionID            string      `json:"session_id"`
}

type Extension struct {
	Name                       string         `json:"name"`
	Protocols                  []string       `json:"protocols,omitempty"`
	SharedKeys                 []SharedKey    `json:"shared_keys,omitempty"`
	MasterSecretData           *string        `json:"master_secret_data,omitempty"`
	ExtendedMasterSecretData   *string        `json:"extended_master_secret_data,omitempty"`
	EllipticCurvesPointFormats []string       `json:"elliptic_curves_point_formats,omitempty"`
	Data                       *string        `json:"data,omitempty"`
	Versions                   []string       `json:"versions,omitempty"`
	ServerName                 *string        `json:"server_name,omitempty"`
	StatusRequest              *StatusRequest `json:"status_request,omitempty"`
	Algorithms                 []string       `json:"algorithms,omitempty"`
	SupportedGroups            []string       `json:"supported_groups,omitempty"`
	SignatureAlgorithms        []string       `json:"signature_algorithms,omitempty"`
	PSKKeyExchangeMode         *string        `json:"PSK_Key_Exchange_Mode,omitempty"`
	PaddingDataLength          *int64         `json:"padding_data_length,omitempty"`
}

type SharedKey struct {
	TLSGREASE0X4A4A *string `json:"TLS_GREASE (0x4a4a),omitempty"`
	X2551929        *string `json:"X25519 (29),omitempty"`
}

type StatusRequest struct {
	CertificateStatusType   string `json:"certificate_status_type"`
	ResponderIDListLength   int64  `json:"responder_id_list_length"`
	RequestExtensionsLength int64  `json:"request_extensions_length"`
}
