// Code generated by dol build. Only Generate by tools if not existed.
// source: UserSrv.go

package rpc

import (
	"github.com/2637309949/dolphin/platform/rpc/proto"

	"golang.org/x/net/context"
)

// UserSrv defined
type UserSrv struct{}

// List defined
func (srv *UserSrv) List(ctx context.Context, in *proto.UserCond) (*proto.UserReply, error) {
	return &proto.UserReply{}, nil
}
