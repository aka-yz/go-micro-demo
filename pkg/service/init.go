package service

import "github.com/aka-yz/go-micro-core"

var Provider = go_micro_core.NewProvider(
	&UserInfoServiceImpl{},
)
