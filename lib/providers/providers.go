package providers

type Provider interface {
	UpdateRecord(ip string) error
	UpdateRecord6(ip string) error
}

type ProviderInit func() (Provider, error)

// Store init function for each provider
var Providers = map[string]ProviderInit{}

func RegisterProvider(name string, init ProviderInit) {
	Providers[name] = init
}
