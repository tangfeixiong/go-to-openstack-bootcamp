package osctl

import (
	"os"
)

type CounterConfig struct {
	IdentityHost string
	DomainId     string
	DomainName   string
	ProjectId    string
	ProjectName  string
	Userid       string
	Username     string
	Password     string
	Role         string
	Token        string
	RegionName   string
}

func NewCounterConfig() *CounterConfig {
	cc := &CounterConfig{
		IdentityHost: "http://10.100.151.247:5000/v2.0/", // "http://127.0.0.1:5000/v2.0
		DomainName:   "",                                 // "default",
		ProjectName:  "admin",                            // "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
		Username:     "admin",                            // "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
		Password:     "7uj8ik",                           // "00000000-0000-0000-0000-000000000000",
		Role:         "",                                 // "_member_",
		RegionName:   "openstack",                        // "regionOne"
	}

	cc.ProjectId = os.Getenv("COUNTER_OSP_ID")
	cc.Userid = os.Getenv("COUNTER_OSU_ID")

	return cc
}
