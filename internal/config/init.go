package conf

import (
	go_micro_core "github.com/aka-yz/go-micro-core"
	"github.com/facebookgo/inject"
)

type provide struct{}

//Provide for inject object
func (p *provide) Provide() []*inject.Object {
	ac := &AppConf{}
	go_micro_core.LoadAppConf("app", ac)

	return []*inject.Object{
		{
			Value: ac,
		},
	}
}

//Provider for inject
var Provider = &provide{}
