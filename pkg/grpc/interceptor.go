package grpc

import (
	"context"
	"github.com/qts0312/ChaosRPC/pkg/call_site"
	"github.com/qts0312/ChaosRPC/pkg/failure"
	"github.com/qts0312/ChaosRPC/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"strconv"
	"strings"
	"time"
)

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	callSite := call_site.GetCallSite()
	fullCallSite := strings.Join(callSite, ";")
	fullCallSite = fullCallSite + ";" + method

	targetCallSite := os.Getenv("CHAOS_CALL_SITE")
	errorCode, err := strconv.Atoi(os.Getenv("CHAOS_ERROR_CODE"))
	if err != nil || fullCallSite != targetCallSite {
		errorCode = failure.ErrorNone
	}

	switch errorCode {
	case failure.ErrorNone:
		return invoker(ctx, method, req, reply, cc, opts...)
	case failure.ErrorOutboundUnavailable:
		reply = nil
		logger.Infof("Outbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Outbound unavailable by ChaosRPC")
	case failure.ErrorInboundUnavailable:
		_ = invoker(ctx, method, req, reply, cc, opts...)
		reply = nil
		logger.Infof("Inbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Inbound unavailable by ChaosRPC")
	case failure.ErrorInboundTimeout:
		waitTime, _ := strconv.Atoi(os.Getenv("CHAOS_WAIT_SEC"))
		_ = invoker(ctx, method, req, reply, cc, opts...)
		reply = nil
		time.Sleep(time.Duration(waitTime) * time.Second)
		return status.Errorf(codes.DeadlineExceeded, "Timeout by ChaosRPC")
	default:
		logger.Fatalf("Unknown error code %d", errorCode)
	}
	return nil
}
