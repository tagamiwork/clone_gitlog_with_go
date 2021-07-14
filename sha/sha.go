package sha

import "encoding/hex"

type SHA1 []byte // len(SHA1) == 20

func (sha1 SHA1) String() string {
	// デバッグしやすいようにbyte列を文字列へ変換
	return hex.EncodeToString(sha1)
}

/*

SHA1の組み込みクラスが用意されているが
オブジェクトが20バイトの配列なので、扱いづらいため
今回は自分で定義した

*/
