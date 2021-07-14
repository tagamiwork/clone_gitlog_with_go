package store

import (
	"path/filepath"

	"../util"
	"github.com/shumon84/git-log/sha"
)

// 定義したobjectを読み込む処理

type Client struct {
	objectDir string
}

// コンテンツ管理を扱うクライアントの構造体を作る
func NewClient(path string) (*Client, error) {
	// path のリポジトリのルートディレクトリを探す
	rootDir, err := util.FindGitRoot(path)
	if err != nil {
		return nil, err
	}
	return &Client{
		// filepath.JoinでOSごとの区切り文字の違いを良しなに調整してくれる
		objectDir: filepath.Join(rootDir, ".git", "objects"),
	}, nil

}

//GitObject は hashで指定したオブジェクトを返す
func (c *Client) GitObject(hash sha.SHA1) (*object.Object, error) {
	hashString := hash.String()
	filepath.Join(c.objectDir, hashString[:2], hashString[2:])
}
