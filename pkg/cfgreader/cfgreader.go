package cfgreader

//test only

import (
	"fmt"
	"github.com/van-pelt/pools/pkg/config"
)

type CfgReader struct {
	Cfg *config.Config
}

func ReadCfg(cfg *config.Config) *CfgReader {
	return &CfgReader{Cfg: cfg}
}

func (C *CfgReader) R() {
	fmt.Println(C.Cfg.App, " ", C.Cfg.Database.Host)
}
