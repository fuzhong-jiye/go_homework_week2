package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

/**
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答: 需要Wrap, 包含当前出错时的参数(固定案发现场)
 */

func service(params uint)(res interface{}, err error) {
	err = getByID(params)
	if err != nil {
		 return nil, errors.Wrapf(err, "根据主键%d查询出错", params)
	}

	return res, nil
}


// 模拟数据库操作出错
func getByID(id uint) error {
	return sql.ErrNoRows
}