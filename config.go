package client

const (
	DefaultAPIURL     = "https://wd5-impl-services1.workday.com/ccx/service"
	DefaultAPIVersion = "v27.1"
	DefaultPageSize   = 999
	DefaultTimeout    = 60
)

type Config struct {
	APIURL     string
	APIVersion string
	Username   string
	Password   string
	Tenant     string
	PageSize   int
	Timeout    int
}
