package fpclient

type Unknown interface{}

type Fingerprint struct {
	Plugins            Plugins              `json:"plugins"`
	GPU                GPU                  `json:"gpu"`
	DefaultCS          map[string]*any `json:"defaultCS"`
	WindowVersion      []string             `json:"windowVersion"`
	HTMLElementVersion []string             `json:"htmlElementVersion"`
	Webgl              Webgl                `json:"webgl"`
	Webgl2             Webgl2               `json:"webgl2"`
	Navigator          Navigator            `json:"navigator"`
	Window             Window               `json:"window"`
	Document           Document             `json:"document"`
	Screen             Screen               `json:"screen"`
	Body               Body                 `json:"body"`
	MediaDevices       []MediaDevice        `json:"mediaDevices"`
	Battery            Battery              `json:"battery"`
	Voices             []Voice              `json:"voices"`
	Keyboard           Keyboard             `json:"keyboard"`
	Permissions        Permissions          `json:"permissions"`
	MIMETypes          []Welcome10MIMEType  `json:"mimeTypes"`
	RTC                []RTC                `json:"rtc"`
	AllFonts           []AllFont            `json:"allFonts"`
}

type AllFont struct {
	Name   string `json:"name"`
	Exists int64  `json:"exists"`
}

type Battery struct {
	Charging        bool        `json:"charging"`
	ChargingTime    int64       `json:"chargingTime"`
	DischargingTime interface{} `json:"dischargingTime"`
	Level           int64       `json:"level"`
}

type Body struct {
	ClientWidth  int64 `json:"clientWidth"`
	ClientHeight int64 `json:"clientHeight"`
}

type Document struct {
	CharacterSet string `json:"characterSet"`
	CompatMode   string `json:"compatMode"`
	DocumentMode string `json:"documentMode"`
	Layers       string `json:"layers"`
	Images       string `json:"images"`
}

type GPU struct {
	Vendor   string `json:"vendor"`
	Renderer string `json:"renderer"`
}

type Keyboard struct {
	KeyA          string `json:"KeyA"`
	KeyB          string `json:"KeyB"`
	KeyC          string `json:"KeyC"`
	KeyD          string `json:"KeyD"`
	KeyE          string `json:"KeyE"`
	KeyF          string `json:"KeyF"`
	KeyG          string `json:"KeyG"`
	KeyH          string `json:"KeyH"`
	KeyI          string `json:"KeyI"`
	KeyJ          string `json:"KeyJ"`
	KeyK          string `json:"KeyK"`
	KeyL          string `json:"KeyL"`
	KeyM          string `json:"KeyM"`
	KeyN          string `json:"KeyN"`
	KeyO          string `json:"KeyO"`
	KeyP          string `json:"KeyP"`
	KeyQ          string `json:"KeyQ"`
	KeyR          string `json:"KeyR"`
	KeyS          string `json:"KeyS"`
	KeyT          string `json:"KeyT"`
	KeyU          string `json:"KeyU"`
	KeyV          string `json:"KeyV"`
	KeyW          string `json:"KeyW"`
	KeyX          string `json:"KeyX"`
	KeyY          string `json:"KeyY"`
	KeyZ          string `json:"KeyZ"`
	Digit1        string `json:"Digit1"`
	Digit2        string `json:"Digit2"`
	Digit3        string `json:"Digit3"`
	Digit4        string `json:"Digit4"`
	Digit5        string `json:"Digit5"`
	Digit6        string `json:"Digit6"`
	Digit7        string `json:"Digit7"`
	Digit8        string `json:"Digit8"`
	Digit9        string `json:"Digit9"`
	Digit0        string `json:"Digit0"`
	Minus         string `json:"Minus"`
	Equal         string `json:"Equal"`
	BracketLeft   string `json:"BracketLeft"`
	BracketRight  string `json:"BracketRight"`
	Backslash     string `json:"Backslash"`
	Semicolon     string `json:"Semicolon"`
	Quote         string `json:"Quote"`
	Backquote     string `json:"Backquote"`
	Comma         string `json:"Comma"`
	Period        string `json:"Period"`
	Slash         string `json:"Slash"`
	IntlBackslash string `json:"IntlBackslash"`
}

type Welcome10MIMEType struct {
	MIMEType      string     `json:"mimeType"`
	AudioPlayType *OPlayType `json:"audioPlayType,omitempty"`
	VideoPlayType *OPlayType `json:"videoPlayType,omitempty"`
	MediaSource   *bool      `json:"mediaSource,omitempty"`
	MediaRecorder *bool      `json:"mediaRecorder,omitempty"`
}

type MediaDevice struct {
	DeviceID string `json:"deviceId"`
	Kind     string `json:"kind"`
	Label    string `json:"label"`
	GroupID  string `json:"groupId"`
}

type Navigator struct {
	Languages                   []string    `json:"languages"`
	UserAgent                   string      `json:"userAgent"`
	AppCodeName                 string      `json:"appCodeName"`
	AppMinorVersion             string      `json:"appMinorVersion"`
	AppName                     string      `json:"appName"`
	AppVersion                  string      `json:"appVersion"`
	BuildID                     string      `json:"buildID"`
	Platform                    string      `json:"platform"`
	Product                     string      `json:"product"`
	ProductSub                  string      `json:"productSub"`
	HardwareConcurrency         int64       `json:"hardwareConcurrency"`
	CPUClass                    string      `json:"cpuClass"`
	MaxTouchPoints              int64       `json:"maxTouchPoints"`
	Oscpu                       string      `json:"oscpu"`
	Vendor                      string      `json:"vendor"`
	VendorSub                   string      `json:"vendorSub"`
	DeviceMemory                int64       `json:"deviceMemory"`
	DoNotTrack                  interface{} `json:"doNotTrack"`
	MSDoNotTrack                string      `json:"msDoNotTrack"`
	Vibrate                     string      `json:"vibrate"`
	Credentials                 string      `json:"credentials"`
	Storage                     string      `json:"storage"`
	RequestMediaKeySystemAccess string      `json:"requestMediaKeySystemAccess"`
	Bluetooth                   string      `json:"bluetooth"`
	Language                    string      `json:"language"`
	SystemLanguage              string      `json:"systemLanguage"`
	UserLanguage                string      `json:"userLanguage"`
}

type Permissions struct {
	StorageAccess          AccessibilityEvents `json:"storage-access"`
	Push                   AccessibilityEvents `json:"push"`
	Speaker                AccessibilityEvents `json:"speaker"`
	DeviceInfo             AccessibilityEvents `json:"device-info"`
	Bluetooth              AccessibilityEvents `json:"bluetooth"`
	Clipboard              AccessibilityEvents `json:"clipboard"`
	AmbientLightSensor     AccessibilityEvents `json:"ambient-light-sensor"`
	AccessibilityEvents    AccessibilityEvents `json:"accessibility-events"`
	NFC                    AccessibilityEvents `json:"nfc"`
	SystemWakeLock         AccessibilityEvents `json:"system-wake-lock"`
	FontAccess             AccessibilityEvents `json:"font-access"`
	MIDI                   Accelerometer       `json:"midi"`
	BackgroundFetch        Accelerometer       `json:"background-fetch"`
	BackgroundSync         Accelerometer       `json:"background-sync"`
	Accelerometer          Accelerometer       `json:"accelerometer"`
	Gyroscope              Accelerometer       `json:"gyroscope"`
	Magnetometer           Accelerometer       `json:"magnetometer"`
	ScreenWakeLock         Accelerometer       `json:"screen-wake-lock"`
	ClipboardRead          Accelerometer       `json:"clipboard-read"`
	ClipboardWrite         Accelerometer       `json:"clipboard-write"`
	PaymentHandler         Accelerometer       `json:"payment-handler"`
	PeriodicBackgroundSync Accelerometer       `json:"periodic-background-sync"`
	Geolocation            Accelerometer       `json:"geolocation"`
	Notifications          Accelerometer       `json:"notifications"`
	Camera                 Accelerometer       `json:"camera"`
	Microphone             Accelerometer       `json:"microphone"`
	DisplayCapture         Accelerometer       `json:"display-capture"`
	PersistentStorage      Accelerometer       `json:"persistent-storage"`
	IdleDetection          Accelerometer       `json:"idle-detection"`
	WindowPlacement        Accelerometer       `json:"window-placement"`
}

type Accelerometer struct {
	State State `json:"state"`
}

type AccessibilityEvents struct {
	ExType ExType `json:"exType"`
	Msg    string `json:"msg"`
}

type Plugins struct {
	MIMETypes []PluginsMIMEType `json:"mimeTypes"`
	Plugins   []Plugin          `json:"plugins"`
}

type PluginsMIMEType struct {
	Type        TypeElement `json:"type"`
	Suffixes    string      `json:"suffixes"`
	Description string      `json:"description"`
	PluginName  string      `json:"__pluginName"`
}

type Plugin struct {
	Name        string        `json:"name"`
	Filename    string        `json:"filename"`
	Description string        `json:"description"`
	MIMETypes   []TypeElement `json:"__mimeTypes"`
}

type RTC struct {
	Candidate string      `json:"candidate"`
	Reg       interface{} `json:"reg"`
}

type Screen struct {
	AvailWidth  int64 `json:"availWidth"`
	AvailHeight int64 `json:"availHeight"`
	AvailLeft   int64 `json:"availLeft"`
	AvailTop    int64 `json:"availTop"`
	Width       int64 `json:"width"`
	Height      int64 `json:"height"`
	ColorDepth  int64 `json:"colorDepth"`
	PixelDepth  int64 `json:"pixelDepth"`
}

type Voice struct {
	Default      bool   `json:"default"`
	Lang         string `json:"lang"`
	LocalService bool   `json:"localService"`
	Name         string `json:"name"`
	VoiceURI     string `json:"voiceURI"`
}

type Webgl struct {
	SupportedExtensions    []string                `json:"supportedExtensions"`
	ContextAttributes      ContextAttributes       `json:"contextAttributes"`
	MaxAnisotropy          int64                   `json:"maxAnisotropy"`
	Params                 map[string]any        `json:"params"`
	ShaderPrecisionFormats []ShaderPrecisionFormat `json:"shaderPrecisionFormats"`
}

type ContextAttributes struct {
	Alpha                        bool   `json:"alpha"`
	Antialias                    bool   `json:"antialias"`
	Depth                        bool   `json:"depth"`
	Desynchronized               bool   `json:"desynchronized"`
	FailIfMajorPerformanceCaveat bool   `json:"failIfMajorPerformanceCaveat"`
	PowerPreference              string `json:"powerPreference"`
	PremultipliedAlpha           bool   `json:"premultipliedAlpha"`
	PreserveDrawingBuffer        bool   `json:"preserveDrawingBuffer"`
	Stencil                      bool   `json:"stencil"`
	XrCompatible                 bool   `json:"xrCompatible"`
}

type Param struct {
	Type  ParamType `json:"type"`
	Value *Value    `json:"value"`
}

type ShaderPrecisionFormat struct {
	ShaderType    int64 `json:"shaderType"`
	PrecisionType int64 `json:"precisionType"`
	R             R     `json:"r"`
}

type R struct {
	RangeMin  int64 `json:"rangeMin"`
	RangeMax  int64 `json:"rangeMax"`
	Precision int64 `json:"precision"`
}

type Webgl2 struct {
	SupportedExtensions    []string                `json:"supportedExtensions"`
	ContextAttributes      ContextAttributes       `json:"contextAttributes"`
	MaxAnisotropy          int64                   `json:"maxAnisotropy"`
	Params                 map[string]any        `json:"params"`
	ShaderPrecisionFormats []ShaderPrecisionFormat `json:"shaderPrecisionFormats"`
}

type Window struct {
	InnerWidth                  int64  `json:"innerWidth"`
	InnerHeight                 int64  `json:"innerHeight"`
	OuterWidth                  int64  `json:"outerWidth"`
	OuterHeight                 int64  `json:"outerHeight"`
	ScreenX                     int64  `json:"screenX"`
	ScreenY                     int64  `json:"screenY"`
	PageXOffset                 int64  `json:"pageXOffset"`
	PageYOffset                 int64  `json:"pageYOffset"`
	Image                       string `json:"Image"`
	IsSecureContext             bool   `json:"isSecureContext"`
	DevicePixelRatio            int64  `json:"devicePixelRatio"`
	Toolbar                     string `json:"toolbar"`
	Locationbar                 string `json:"locationbar"`
	ActiveXObject               string `json:"ActiveXObject"`
	External                    string `json:"external"`
	MozRTCPeerConnection        string `json:"mozRTCPeerConnection"`
	PostMessage                 string `json:"postMessage"`
	WebkitRequestAnimationFrame string `json:"webkitRequestAnimationFrame"`
	BluetoothUUID               string `json:"BluetoothUUID"`
	Netscape                    string `json:"netscape"`
	LocalStorage                string `json:"localStorage"`
	SessionStorage              string `json:"sessionStorage"`
	IndexDB                     string `json:"indexDB"`
	BarcodeDetector             string `json:"BarcodeDetector"`
}

type OPlayType string

const (
	Maybe    OPlayType = "maybe"
	Probably OPlayType = "probably"
)

type State string

const (
	Denied  State = "denied"
	Granted State = "granted"
	Prompt  State = "prompt"
)

type ExType string

const (
	DOMException ExType = "DOMException"
	TypeError    ExType = "TypeError"
)

type TypeElement string

const (
	ApplicationPDF TypeElement = "application/pdf"
	TextPDF        TypeElement = "text/pdf"
)

type ParamType string

const (
	Array        ParamType = "Array"
	Boolean      ParamType = "Boolean"
	Empty        ParamType = ""
	Float32Array ParamType = "Float32Array"
	Int32Array   ParamType = "Int32Array"
	Number       ParamType = "Number"
	String       ParamType = "String"
	Uint32Array  ParamType = "Uint32Array"
)

type DefaultC struct {
	Integer *int64
	String  *string
	Other   *Unknown
}

type Value struct {
	Bool       *bool
	BoolArray  []bool
	Integer    *int64
	IntegerMap map[string]int64
	String     *string
	
}
