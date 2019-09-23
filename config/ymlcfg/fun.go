/**
 * @version: 1.0.0
 * @author: zgd: general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: fun.go
 * @time: 2019-08-03 22:09
 */

package ymlcfg

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadYaml(path string, tar interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	bts, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(bts, tar)
}
