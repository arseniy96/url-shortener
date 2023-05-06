package config

type Options struct {
	Host        string
	ResolveHost string
}

func SetConfig(host, resolveHost string) *Options {
	return &Options{
		Host:        host,
		ResolveHost: resolveHost,
	}
}
