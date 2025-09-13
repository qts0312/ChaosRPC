package grpc

import (
	"context"
	"github.com/qts0312/ChaosRPC/pkg/call_site"
	"github.com/qts0312/ChaosRPC/pkg/failure"
	"github.com/qts0312/ChaosRPC/pkg/logger"
	"github.com/qts0312/ChaosRPC/pkg/util"
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
		logger.Debugf(fullCallSite)
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
		waitTime, err := strconv.Atoi(os.Getenv("CHAOS_WAIT_SEC"))
		if err != nil {
			waitTime = 5
		}
		_ = invoker(ctx, method, req, reply, cc, opts...)
		reply = nil
		time.Sleep(time.Duration(waitTime) * time.Second)
		return status.Errorf(codes.DeadlineExceeded, "Timeout by ChaosRPC")
	default:
		logger.Fatalf("Unknown error code %d", errorCode)
	}
	return nil
}

func StreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
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
		return streamer(ctx, desc, cc, method, opts...)
	case failure.ErrorOutboundUnavailable:
		logger.Infof("Outbound unavailable on %s", fullCallSite)
		return nil, status.Errorf(codes.Unavailable, "Outbound unavailable by ChaosRPC")
	case failure.ErrorInboundUnavailable:
		_, _ = streamer(ctx, desc, cc, method, opts...)
		logger.Infof("Inbound unavailable on %s", fullCallSite)
		return nil, status.Errorf(codes.Unavailable, "Inbound unavailable by ChaosRPC")
	case failure.ErrorInboundTimeout:
		waitTime, err := strconv.Atoi(os.Getenv("CHAOS_WAIT_SEC"))
		if err != nil {
			waitTime = 5
		}
		_, _ = streamer(ctx, desc, cc, method, opts...)
		time.Sleep(time.Duration(waitTime) * time.Second)
		return nil, status.Errorf(codes.DeadlineExceeded, "Timeout by ChaosRPC")
	default:
		logger.Fatalf("Unknown error code %d", errorCode)
	}
	return nil, nil
}

func StreamSendInterceptor(m any, f func(any) error) error {
	callSite := call_site.GetCallSite()
	fullCallSite := strings.Join(callSite, ";")
	fullCallSite = fullCallSite + ";STREAM_SEND"

	targetCallSite := os.Getenv("CHAOS_CALL_SITE")
	errorCode, err := strconv.Atoi(os.Getenv("CHAOS_ERROR_CODE"))
	if err != nil || fullCallSite != targetCallSite {
		errorCode = failure.ErrorNone
	}

	switch errorCode {
	case failure.ErrorNone:
		return f(m)
	case failure.ErrorOutboundUnavailable:
		logger.Infof("Outbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Outbound unavailable by ChaosRPC")
	case failure.ErrorInboundUnavailable:
		_ = f(m)
		logger.Infof("Inbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Inbound unavailable by ChaosRPC")
	case failure.ErrorInboundTimeout:
		logger.Infof("Timeout is not supported in stream")
		return f(m)
	default:
		logger.Fatalf("Unknown error code %d", errorCode)
	}
	return nil
}

func StreamRecvInterceptor(m any, f func(any) error) error {
	callSite := call_site.GetCallSite()
	fullCallSite := strings.Join(callSite, ";")
	fullCallSite = fullCallSite + ";STREAM_RECV"

	targetCallSite := os.Getenv("CHAOS_CALL_SITE")
	errorCode, err := strconv.Atoi(os.Getenv("CHAOS_ERROR_CODE"))
	if err != nil || fullCallSite != targetCallSite {
		errorCode = failure.ErrorNone
	}

	switch errorCode {
	case failure.ErrorNone:
		return f(m)
	case failure.ErrorOutboundUnavailable:
		logger.Infof("Outbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Outbound unavailable by ChaosRPC")
	case failure.ErrorInboundUnavailable:
		_ = f(m)
		logger.Infof("Inbound unavailable on %s", fullCallSite)
		return status.Errorf(codes.Unavailable, "Inbound unavailable by ChaosRPC")
	case failure.ErrorInboundTimeout:
		logger.Infof("Timeout is not supported in stream")
		return f(m)
	default:
		logger.Fatalf("Unknown error code %d", errorCode)
	}
	return nil
}

func Init() {
	util.Assert(grpc.ChaosUnaryClientInterceptor == nil && grpc.ChaosStreamClientInterceptor == nil, "Client interceptors already set\n")
	grpc.ChaosUnaryClientInterceptor = UnaryClientInterceptor
	grpc.ChaosStreamClientInterceptor = StreamClientInterceptor
	util.Assert(grpc.ChaosStreamSendInterceptor == nil && grpc.ChaosStreamRecvInterceptor == nil, "Stream interceptors already set\n")
	grpc.ChaosStreamSendInterceptor = StreamSendInterceptor
	grpc.ChaosStreamRecvInterceptor = StreamRecvInterceptor
	logger.Infof("Initializing ChaosRPC interceptors")
}
