package ctx

import (
	"flag"
	"uc/pkg/tool"

	"github.com/BurntSushi/toml"
)

//Options 配置信息供外部使用
var Options *options

// Options 配置选项
type options struct {
	LogPath string `toml:"log_path"`
}

func init() {
	var confFile string
	f := flag.String("c", "", "配置文件路径")
	flag.Parse()
	if len(*f) > 0 {
		confFile = *f
	} else {
		confFile, _ = tool.GetAppDirectory()
		confFile = confFile + "/conf.toml"
	}
	newOptions(confFile)
}

//NewOptions 初始化配置
func newOptions(confFile string) {
	var err error
	Options, err = configFileparse(confFile)
	if err != nil {
		Log.Info("未能读取配置文件，将执行默认配置...")
	}
}

//configFileparse 日志文件解析
func configFileparse(filePath string) (*options, error) {
	o := new(options)
	_, err := toml.DecodeFile(filePath, &o)
	return o, err
}
