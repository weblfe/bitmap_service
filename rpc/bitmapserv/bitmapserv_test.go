package bitmapserv

import (
	"context"
	"fmt"
	"github.com/tal-tech/go-zero/zrpc"
	"testing"
	"time"
)

func TestNewBitMapServ(t *testing.T) {
	var (
		cnf = zrpc.NewEtcdClientConf([]string{`127.0.0.1:2379`},`bitmaprpc.rpc`,``,``)

		client, err = zrpc.NewClient(cnf)
	)
	fmt.Println(fmt.Sprintf("%v",[]byte(`bitmaprpc`)))
	fmt.Println(fmt.Sprintf("%T,%T",[]string{`test`,`123`},[3]string{}))
	if err != nil {
		t.Error(err)
		return
	}
	var serv = NewBitMapServ(client)
	res, err := serv.Ping(context.TODO(), &PingRequest{
		Ack:       1,
		Timestamp: time.Now().Unix(),
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(res.String())
}
