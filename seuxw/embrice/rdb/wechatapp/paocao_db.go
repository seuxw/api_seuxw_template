package wechatapp

import (
	"fmt"
	"seuxw/embrice/entity/wechatapp"
)

func (self *Database) PaocaoAPICount(date string) (wechatapp.PaocaoAPICount, error) {
	var (
		selectSQL      string
		selectCheckSQL string
		count          int
		PaocaoAPICount wechatapp.PaocaoAPICount
		err            error
	)

	selectCheckSQL = `
	select
		count(1) as count
	from
		s_paocao_use
	where
		use_date = ?
	`

	selectSQL = `
	select
		use_date, use_count
	from
		s_paocao_use
	where
		use_date = ?
	limit 1
	`
	if err = self.Get(&count, selectCheckSQL, date); err != nil {
		err = fmt.Errorf("数据库预查询错误 err:%s", err)
		goto END
	}

	if count <= 0 {
		err = fmt.Errorf("查询日期错误")
		goto END
	}

	if err = self.Get(&PaocaoAPICount, selectSQL, date); err != nil {
		err = fmt.Errorf("数据库查询错误 err:%s", err)
		goto END
	}
	PaocaoAPICount.UseDate = date

END:
	return PaocaoAPICount, err
}
