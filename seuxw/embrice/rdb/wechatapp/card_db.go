package wechatapp

import (
	"database/sql"
	"fmt"
	"seuxw/embrice/entity/wechatapp"
)

// 通过QQ、学号查询一卡通
// Type:0 - 一卡通 1 - QQ	Type:2 - 学号
func (self *Database) SelectCard(InputNO string, Type int) (wechatapp.CardBindCancel, error) {
	var (
		selectSQL      string
		whereSQL       string
		err            error
		CardBindCancel wechatapp.CardBindCancel
	)

	selectSQL = `
	select
		cardNo as card_no, name
	from 
		student_info
	where
		%s
	`
	if Type == 1 {
		whereSQL = `QQ = ?`
	} else if Type == 2 {
		whereSQL = `stuNo = ?`
	} else if Type == 0 {
		whereSQL = `cardNo = ?`
	} else {
		err = fmt.Errorf("参数 类型传递错误，请检查！")
		goto END
	}

	selectSQL = fmt.Sprintf(selectSQL, whereSQL)

	if err = self.Get(&CardBindCancel, selectSQL, InputNO); err != nil {
		err = fmt.Errorf("数据库查询错误 err：%s", err)
		goto END
	}
END:
	return CardBindCancel, err
}

// CardBindCount 一卡通绑定人数统计
func (self *Database) CardBindCount(grade string, major string, Bind int64) (wechatapp.CardBindCount, error) {
	var (
		selectSQL     string
		selectAllSQL  string
		whereStr      string
		CardBindCount wechatapp.CardBindCount
		count_all     int64
		err           error
	)

	selectAllSQL = `
	select
		count(1) as count_all
	from
		student_info
	where
	 	1=1 %s
	`

	selectSQL = `
	select
		count(1) as count_bind, (count(1)/%d) as count_rate
	from
		student_info
	where
	 	1=1 %s
	`

	// 拼接查询条件

	if grade != "" {
		whereStr = fmt.Sprintf(`%s and grade = "%s" `, whereStr, grade)
	}

	if major != "" {
		whereStr = fmt.Sprintf(`%s and SUBSTR(stuNo,1,2) = "%s" `, whereStr, major)
	}

	// 查询总体人数
	if err = self.Get(&count_all, fmt.Sprintf(selectAllSQL, whereStr)); err != nil {
		err = fmt.Errorf("数据库查询错误 err:%s", err)
		goto END
	}

	if count_all == 0 {
		err = fmt.Errorf("该查询条件下的总人数为0！")
		goto END
	}

	// 查询绑定人数
	if Bind == 0 {
		whereStr = fmt.Sprintf(`%s and QQ is null `, whereStr)
	} else {
		whereStr = fmt.Sprintf(`%s and QQ is not null `, whereStr)
	}

	if err = self.Get(&CardBindCount, fmt.Sprintf(selectSQL, count_all, whereStr)); err != nil {
		err = fmt.Errorf("数据库查询错误 err:%s", err)
		goto END
	}

	CardBindCount.CountAll = count_all
END:
	return CardBindCount, err
}

// CardBindCancel
func (self *Database) CardBindCancel(QQ string, CardNO string, stuNO string) (wechatapp.CardBindCancel, error) {
	var (
		CardBindCancel    wechatapp.CardBindCancel
		CardBindCancelSQL string
		Result            sql.Result
		err               error
	)

	// 为了避免代码的复杂度，均使用CardNO来更新字段
	CardBindCancelSQL = `
	update
		student_info
	set
		QQ = null
	where
		cardNo = ?
	`
	if len(QQ) != 0 {
		CardBindCancel, err = self.SelectCard(QQ, 1)
	} else if len(stuNO) != 0 {
		CardBindCancel, err = self.SelectCard(stuNO, 2)
	} else if len(CardNO) != 0 {
		CardBindCancel, err = self.SelectCard(CardNO, 0)
	}

	if Result, err = self.Exec(CardBindCancelSQL, CardBindCancel.CardNO); err != nil {
		err = fmt.Errorf("数据库执行更新错误！ err:%s", err)
		goto END
	}

	if err = self.JudgeAffect(Result); err != nil {
		goto END
	}
END:
	return CardBindCancel, err
}
