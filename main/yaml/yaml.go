package yaml

import (
	"io"

	"github.com/aiempire79/mr-xray/common"
	"github.com/aiempire79/mr-xray/common/cmdarg"
	"github.com/aiempire79/mr-xray/core"
	"github.com/aiempire79/mr-xray/infra/conf"
	"github.com/aiempire79/mr-xray/infra/conf/serial"
	"github.com/aiempire79/mr-xray/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "YAML",
		Extension: []string{"yaml", "yml"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				cf := &conf.Config{}
				for i, arg := range v {
					newError("Reading config: ", arg).AtInfo().WriteToLog()
					r, err := confloader.LoadConfig(arg)
					if err != nil {
						return nil, newError("failed to read config: ", arg).Base(err)
					}
					c, err := serial.DecodeYAMLConfig(r)
					if err != nil {
						return nil, newError("failed to decode config: ", arg).Base(err)
					}
					if i == 0 {
						// This ensure even if the muti-json parser do not support a setting,
						// It is still respected automatically for the first configure file
						*cf = *c
						continue
					}
					cf.Override(c, arg)
				}
				return cf.Build()
			case io.Reader:
				return serial.LoadYAMLConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
