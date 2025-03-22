package model

type Config struct {
	ClientId     string   `env:"CLIENT_ID,notEmpty,required"`
	TenantId     string   `env:"TENANT_ID,notEmpty,required"`
	ClientSecret string   `env:"CLIENT_SECRET,notEmpty,required"`
	BaseUrl      string   `env:"BASE_URL,notEmpty,required"`
	AdminToken   string   `env:"ADMIN_TOKEN,notEmpty,required"`
	Groups       []string `env:"GROUPS,notEmpty,required"`
	Whitelist    []string `env:"WHITELIST"`
	Cron         string   `env:"CRON" envDefault:"*/15 * * * *"`
}
