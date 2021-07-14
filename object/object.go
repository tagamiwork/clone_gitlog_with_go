package object

import (
	"fmt"

	"github.com/shumon84/git-log/sha"
)

// objectの定義
type Object struct {
	Hash sha.SHA1
	Type Type // object_type.goに定義
	Size int
	Data []byte
}

func (o *Object) Header() []byte {
	return []byte(fmt.Sprintf("%s %d\x00", o.Type, o.Size))
}
