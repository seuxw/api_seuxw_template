package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seuxw/embrice/entity/wechatapp"
	"seuxw/embrice/extension"
)

func (self *server) CardBindCount(w http.ResponseWriter, r *http.Request) {
	var (
		grade         string
		major         string
		bind          string
		Bind          int64
		CardBindCount wechatapp.CardBindCount
		err           error
	)

	// 获取参数
	grade = r.FormValue("grade")
	major = r.FormValue("major")
	bind = r.FormValue("bind")

	// 参数检查
	if bind == "1" {
		Bind = 1
	} else if bind == "" || bind == "0" {
		Bind = 0
	} else {
		err = fmt.Errorf("bind参数只能是0或者1！")
		goto END
	}

	self.PrintTrace(r, fmt.Sprintf("grade:%s major:%s Bind:%d", grade, major, Bind), "GET: 开始调用查询跑操接口调用次数接口")

	if CardBindCount, err = self.db.CardBindCount(grade, major, Bind); err != nil {
		err = fmt.Errorf("[CardBindCount]数据库调用错误！ %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, CardBindCount)
}

func (self *server) CardBindCancelGET(w http.ResponseWriter, r *http.Request) {
	var (
		QQ             string
		CardNO         string
		StuNO          string
		err            error
		CardBindCancel wechatapp.CardBindCancel
	)

	// 获取参数
	QQ = r.FormValue("qq")
	CardNO = r.FormValue("card_no")
	StuNO = r.FormValue("stu_no")

	// 参数判断 - 不能全为空
	if len(QQ) == 0 && len(CardNO) == 0 && len(StuNO) == 0 {
		err = fmt.Errorf("请至少输入一个参数[qq, card_no, stu_no]")
		goto END
	}

	if CardBindCancel, err = self.db.CardBindCancel(QQ, CardNO, StuNO); err != nil {
		err = fmt.Errorf("[CardBindCancel]数据库操作错误 %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, CardBindCancel)
}

type StudentInfo struct {
	QQ     string `json:"qq"`
	CardNO string `json:"card_no"`
	StuNO  string `json:"stu_no"`
}

func (self *server) CardBindCancelPOST(w http.ResponseWriter, r *http.Request) {
	var (
		err            error
		StudentInfo    StudentInfo
		CardBindCancel wechatapp.CardBindCancel
	)

	// 获取参数
	defer r.Body.Close()
	if err = json.NewDecoder(r.Body).Decode(&StudentInfo); err != nil {
		err = fmt.Errorf("[StudentInfo] Json decode err: %v", err)
		goto END
	}

	// 参数判断 - 不能全为空
	if len(StudentInfo.QQ) == 0 && len(StudentInfo.CardNO) == 0 && len(StudentInfo.StuNO) == 0 {
		err = fmt.Errorf("请至少输入一个参数[qq, card_no, stu_no]")
		goto END
	}

	if CardBindCancel, err = self.db.CardBindCancel(StudentInfo.QQ, StudentInfo.CardNO, StudentInfo.StuNO); err != nil {
		err = fmt.Errorf("[CardBindCancel]数据库操作错误 %s", err)
		goto END
	}
END:
	extension.EndLabel(self.log, err, w, nil, CardBindCancel)
}
