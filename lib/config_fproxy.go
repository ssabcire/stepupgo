package fproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type PathIPConfig struct {
	FilePath string `json:"filepath"`
	PxIP     string `json:"pxip"`
}

//ディレクトリを生成するための変数
var JsonDir = func() string {
		//現在のユーザーを取得
		user, err := user.Current()
		if err != nil {
			return ""
		}
		jsonPath := filepath.Join(user.HomeDir, ".sfp")
		return jsonPath
}()

var JsonPath = func() string {
	jsonPath := filepath.Join(JsonDir, "config.json")
	return jsonPath
}()

//pathとネットワークアドレスを記載するjsonファイルを作成
func createJsonFile() {
	content := []byte(`
{
  "filepath": "test",
  "pxip": "127.0.0.1"
}
	`)
	err := ioutil.WriteFile(JsonPath, content, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//jsonを読みこみ、構造体に移す
func (p *PathIPConfig) ReadJsonTransfer() error {
	//ファイルの存在確認
	_, err := os.Stat(JsonPath)
	if os.IsNotExist(err) {
		createJsonFile()
	}

	data, err := ioutil.ReadFile(JsonPath)
	if err != nil {
		fmt.Println("ファイルの読み込みに失敗", err)
		return err
	}

	err = json.Unmarshal(data, &p)
	if err != nil {
		return err
	}
	return nil
}