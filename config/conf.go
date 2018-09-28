package conf

import (
	"fmt"
	"os"
	"strings"

	"github.com/Unknwon/goconfig"
)

var Config *goconfig.ConfigFile
var (
	Appname  string
	Httpport int
	Runmode  string
	NodeName string //节点名称 拼音简写
)

//是否是空字符串（没有字符）
func IsEmpty(s string) bool {
	if s == "" {
		return true
	} else {
		return false
	}
}

//是否是空白字符串(有字符，但是是空白字符、换行符等)
func IsWhite(s string) bool {
	if len(s) != 0 && len(strings.TrimSpace(s)) == 0 {
		return true
	} else {
		return false
	}
}

//是否是空或空白字符串
func IsNullOrEmpty(s string) bool {
	return IsEmpty(s) || IsWhite(s)
}

func init() {
	if Config == nil {
		firstcfg, configs, err := GetConfigs("../config") //不要带斜杠
		if err == nil {
			Config, err = goconfig.LoadConfigFile(firstcfg, configs...)
			if err != nil {
				panic("fail to load config file")
			}
		} else {
			panic(fmt.Errorf("fail to load config file '%s': %v", "../config", err))
		}
	}

	Appname, _ = Config.GetValue("app", "appname")
	Runmode, _ = Config.GetValue("app", "run_mode")
	Httpport = Config.MustInt(Runmode, "httpport", 8080)
	NodeName = Config.MustValue(Runmode, "nodename", "")

	if IsNullOrEmpty(NodeName) {
		fmt.Println("必须要设置nodename的值! example: nodename=shenyang")
		os.Exit(0)
	}
}
