package redis

import (
	"context"
	"io"
	"net"
	"strings"

	"github.com/2637309949/dolphin/cli/packages/redis/internal/proto"
)

func isRetryableError(err error, retryTimeout bool) bool {
	switch err {
	case nil, context.Canceled, context.DeadlineExceeded:
		return false
	case io.EOF:
		return true
	}
	if netErr, ok := err.(net.Error); ok {
		if netErr.Timeout() {
			return retryTimeout
		}
		return true
	}

	s := err.Error()
	if s == "ERR max number of clients reached" {
		return true
	}
	if strings.HasPrefix(s, "LOADING ") {
		return true
	}
	if strings.HasPrefix(s, "READONLY ") {
		return true
	}
	if strings.HasPrefix(s, "CLUSTERDOWN ") {
		return true
	}
	return false
}

func isRedisError(err error) bool {
	_, ok := err.(proto.RedisError)
	return ok
}

func isBadConn(err error, allowTimeout bool) bool {
	if err == nil {
		return false
	}
	if isRedisError(err) {
		// Close connections in read only state in case domain addr is used
		// and domain resolves to a different Redis Server. See #790.
		return isReadOnlyError(err)
	}
	if allowTimeout {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			return false
		}
	}
	return true
}

func isMovedError(err error) (moved bool, ask bool, addr string) {
	if !isRedisError(err) {
		return
	}

	s := err.Error()
	switch {
	case strings.HasPrefix(s, "MOVED "):
		moved = true
	case strings.HasPrefix(s, "ASK "):
		ask = true
	default:
		return
	}

	ind := strings.LastIndex(s, " ")
	if ind == -1 {
		return false, false, ""
	}
	addr = s[ind+1:]
	return
}

func isLoadingError(err error) bool {
	return strings.HasPrefix(err.Error(), "LOADING ")
}

func isReadOnlyError(err error) bool {
	return strings.HasPrefix(err.Error(), "READONLY ")
}