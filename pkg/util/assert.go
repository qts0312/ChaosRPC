package util

import "github.com/qts0312/ChaosRPC/pkg/logger"

func Assert(condition bool, message string) {
	if !condition {
		logger.Fatalf(message)
	}
}
