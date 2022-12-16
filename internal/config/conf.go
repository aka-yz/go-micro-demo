package conf

//AppConf struct
type AppConf struct {
	ValidNameUrl string
	ChainId      string
	BlockNums    string
	Token        string `yaml:"TOKEN"`
}

//Init Appconf
func (c *AppConf) Init() {
}
