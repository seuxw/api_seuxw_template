package main

import (
	"fmt"
	"net/http"
	"seuxw/embrice/entity/wechatapp"
	"seuxw/embrice/extension"
)

func (self *server) PaocaoAPICount(w http.ResponseWriter, r *http.Request) {
	var (
		getPara        string
		PaocaoAPICount wechatapp.PaocaoAPICount
		err            error
	)

	// 获取date参数
	getPara = r.FormValue("date")
	if len(getPara) == 0 {
		getPara = extension.CurrentDateInStr()
	}

	self.PrintTrace(r, fmt.Sprintf("date:%s", getPara), "GET: 开始调用查询跑操接口调用次数接口")

	if PaocaoAPICount, err = self.db.PaocaoAPICount(getPara); err != nil {
		err = fmt.Errorf("[PaocaoAPICount]数据库调用错误！ %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, PaocaoAPICount)
}
