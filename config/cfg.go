package conf

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
)

func FetchFilePathNames(dirPath, regExpr string) ([]string, error) {
	reg := regexp.MustCompile(regExpr)
	filePathNames := make([]string, 0, 10)

	err := filepath.Walk(dirPath, func(filePathName string, fi os.FileInfo, err error) error { //遍历目录
		if err == nil {
			if fi.IsDir() { // 忽略目录
				return nil
			}
			if reg.MatchString(fi.Name()) {
				filePathNames = append(filePathNames, filePathName)
			}
			return nil
		}
		return err
	})
	return filePathNames, err
}

func GetConfigs(dir string) (string, []string, error) {
	fileCfgs, _ := FetchFilePathNames(dir, `.*(app.conf)$`)    //以app开头，以conf结尾
	fileCuss, _ := FetchFilePathNames(dir, `.*(custom.conf)$`) //以app开头，以conf结尾
	fileCfgs = append(fileCfgs, fileCuss...)
	l := len(fileCfgs)
	if l > 0 {
		return fileCfgs[0], fileCfgs[1:], nil
	} else {
		return "", []string{}, errors.New("no config files!")
	}
}

//框架已经定时更新Token了
//func WxTimerService() {
//	t := time.NewTicker(90 * time.Minute)
//	for {
//		select {
//		case <-t.C:
//			AppLog.Infof("refetch wx token and create new client!")
//			AppWxTokenServ = mp.NewDefaultAccessTokenServer(WxAppId, WxAppSecret, nil)
//			AppWxClient = template.NewClient(AppWxTokenServ, nil)
//		}
//	}
//}
