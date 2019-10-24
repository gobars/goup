package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tkrajina/go-reflector/reflector"
	"os"
	"strings"
)

// FileExists 检查文件是否存在，并且不是目录
func FileExists(filename string) error {
	if fi, err := os.Stat(filename); err != nil {
		return err
	} else if fi.IsDir() {
		return fmt.Errorf("file %s is a directory", filename)
	}

	return nil
}

// ViperToStruct read viper value to struct
func ViperToStruct(structVar interface{}) {
	for _, f := range reflector.New(structVar).Fields() {
		if !f.IsExported() {
			continue
		}

		switch t, _ := f.Get(); t.(type) {
		case []string:
			value := strings.TrimSpace(viper.GetString(f.Name()))
			valueSlice := make([]string, 0)

			for _, v := range strings.Split(value, ",") {
				v = strings.TrimSpace(v)
				if v != "" {
					valueSlice = append(valueSlice, v)
				}
			}

			if len(valueSlice) > 0 {
				if err := f.Set(valueSlice); err != nil {
					logrus.Warnf("Fail to set %s to value %v, error %v", f.Name(), value, err)
				}
			}
		case string:
			if value := strings.TrimSpace(viper.GetString(f.Name())); value != "" {
				if err := f.Set(value); err != nil {
					logrus.Warnf("Fail to set %s to value %v, error %v", f.Name(), value, err)
				}
			}
		case int:
			if value := viper.GetInt(f.Name()); value != 0 {
				if err := f.Set(value); err != nil {
					logrus.Warnf("Fail to set %s to value %v, error %v", f.Name(), value, err)
				}
			}
		case bool:
			if value := viper.GetBool(f.Name()); value {
				if err := f.Set(value); err != nil {
					logrus.Warnf("Fail to set %s to value %v, error %v", f.Name(), value, err)
				}
			}
		}
	}
}

func Contains(arr []string, str string) int {
	for i, a := range arr {
		if a == str {
			return i
		}
	}
	return -1
}
