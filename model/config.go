package model

type Config struct {
	ClientId     string   `env:"CLIENT_ID,notempty,required"`
	TenantId     string   `env:"TENANT_ID,notempty,required"`
	ClientSecret string   `env:"CLIENT_SECRET,notempty,required"`
	BaseUrl      string   `env:"BASE_URL,notempty,required"`
	AdminToken   string   `env:"ADMIN_TOKEN,notempty,required"`
	Groups       []string `env:"GROUPS,notempty,required"`
	Whitelist    []string `env:"WHITELIST,notempty,required"`
	Cron         string   `env:"CRON" envDefault:"*/15 * * * *"`
}
