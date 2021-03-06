package wechatapp

import (
	"database/sql"
	"fmt"
	"seuxw/embrice/constant/config"
	"seuxw/x/logger"
	_ "seuxw/x/mysql"
	"seuxw/x/sqlx"
)

type Database struct {
	*sqlx.DB
	log *logger.Logger
}

func NewDB(log *logger.Logger, maxConns, maxIdles int) *Database {
	db := &Database{
		log: log,
		DB: sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			config.DATABASE_USERNAME, config.DATABASE_PASSWORD, config.DATABASE_HOST, config.
				DATABASE_PORT, config.DATABASE_NAME)), //连接数据库
	}
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdles)
	return db
}

// 判断Update执行语句后影响的函数并输出错误
func (self *Database) JudgeAffect(Result sql.Result) error {
	var (
		err       error
		AffectRow int64
	)
	if AffectRow, err = Result.RowsAffected(); err != nil {
		err = fmt.Errorf("获取更新数据库执行状态失败！ err:%s", err)
		goto END
	}

	if AffectRow == 0 {
		err = fmt.Errorf("没有数据被修改！")
		goto END
	}
END:
	return err
}
