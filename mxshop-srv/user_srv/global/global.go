package global

import (
	"user_srv/config"
	"user_srv/proto"
)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	CouponsClient proto.CouponsClient
)
