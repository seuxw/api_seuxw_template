package main

import (
	"context"
	"fmt"
	"net/http"
	"seuxw/embrice/rdb/wechatapp"
	"seuxw/x/logger"
	"seuxw/x/web"
)

type server struct {
	*web.Server
	db  *wechatapp.Database
	log *logger.Logger
	ctx context.Context
}

func (self *server) start() {
	self.Server = web.NewServer(self.log)
	router := self.PathPrefix("/").SubRouter()

	/* -跑操Handler- */
	// 查询跑操接口调用次数
	router.HandleFunc("/paocao/api_count", self.PaocaoAPICount).Methods("GET")

	/* -一卡通Handler- */
	// 查询一卡通绑定数量
	router.HandleFunc("/card/card_bind_count", self.CardBindCount).Methods("GET")
	// 解绑一卡通接口 出于安全考虑建议使用POST接口
	router.HandleFunc("/card/card_bind_cancel", self.CardBindCancelGET).Methods("GET")
	router.HandleFunc("/card/card_bind_cancel", self.CardBindCancelPOST).Methods("POST")

	/* -测试Handler- */
	// 测试接口 将会返回输入的日期对应的日出日落时间（需要对应时间）
	router.HandleFunc("/test", self.TestGet).Methods("GET")
	router.HandleFunc("/test", self.TestPost).Methods("POST")

	self.Serve("0.0.0.0:10000")
}

func (self *server) stop() {
	self.db.Close()
}

// Trace详情打印函数
func (self *server) PrintTrace(r *http.Request, para string, head string) {
	header := r.Header.Get("X-Trace")
	env := fmt.Sprintf("ENV:{TraceID:%v,Parameters:[%s]}", header, para)
	self.log.Trace("%s: %s", head, env)
}

func main() {
	log := logger.NewStdLogger(true, true, true, true, true)
	s := &server{
		db:  wechatapp.NewDB(log, 10, 10),
		ctx: context.TODO(),
		log: log,
	}

	s.start()

	defer s.stop()
}
