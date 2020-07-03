// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.cli.go

package rpc

import (
	"demo/rpc/proto"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"google.golang.org/grpc"
)

// MessageSrvClient defined
var MessageSrvClient proto.MessageSrvClient

func init() {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(viper.GetString("rpc.message_srv"), opt)
	if err != nil {
		logrus.Error("grpc dial failed: %v", err)
	}
	MessageSrvClient = proto.NewMessageSrvClient(conn)
}
