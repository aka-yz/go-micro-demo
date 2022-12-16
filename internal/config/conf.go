package conf

//AppConf struct
type AppConf struct {
	ConfigKey   string
	SellerAdmin SellerAdmin
	Token       string `yaml:"TOKEN"`
	APIHost     string `yaml:"API_HOST"`
}

//Init Appconf
func (c *AppConf) Init() {
}

//SellerAdmin struct
type SellerAdmin struct {
}
