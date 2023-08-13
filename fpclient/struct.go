package fpclient

type LoadingConfig struct {
	FilePath       string
	String         string
	Isbase64String bool
}

type LoadingTlsConfig struct {
	FilePath       string
	String         string
	Isbase64String bool
}

type DumpTlsConfig struct {
	OutputPath string
	Headless   bool
}

type DumpConfig struct {
	OutputPath  string
	Headless    bool
	DumpPageSrc string
}
