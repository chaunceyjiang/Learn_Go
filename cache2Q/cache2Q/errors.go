package cache2Q

import "github.com/pkg/errors"

var (
	ErrKeyNotFound = errors.New("该Key没有找到！！！")
	ErrKeyNotFoundAndLoadable = errors.New("该Key没有找到,并且添加失败!")
)
