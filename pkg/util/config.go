package util

import (
	"encoding/json"
	"github.com/qts0312/ChaosRPC/pkg/failure"
	"github.com/qts0312/ChaosRPC/pkg/logger"
	"os"
	"strconv"
)

// GetConfig reads necessary configuration values by two ways:
// 1. from environment variables (if exists)
// 2. from /root/chaos_config.json
func GetConfig() (string, int, int, int) {
	callSite := os.Getenv("CHAOS_CALL_SITE")
	errorCode := failure.ErrorNone
	waitTime := 5
	testId := 0
	if callSite != "" {
		tmpErrorCode, err := strconv.Atoi(os.Getenv("CHAOS_ERROR_CODE"))
		if err == nil {
			errorCode = tmpErrorCode
		}
		tmpWaitTime, err := strconv.Atoi(os.Getenv("CHAOS_WAIT_TIME"))
		if err == nil {
			waitTime = tmpWaitTime
		}
		tmpTestId, err := strconv.Atoi(os.Getenv("CHAOS_TEST_ID"))
		if err == nil {
			testId = tmpTestId
		}
	} else {
		type ChaosConfig struct {
			CallSite  string `json:"call_site"`
			ErrorCode int    `json:"error_code"`
			WaitTime  int    `json:"wait_time"`
			TestId    int    `json:"test_id"`
		}
		var config ChaosConfig
		f, err := os.Open("/root/chaos_config.json")
		if err != nil {
			logger.Fatalf("open chaos_config.json error: %v", err)
			return "", errorCode, waitTime, testId
		}
		defer f.Close()
		decoder := json.NewDecoder(f)
		if err := decoder.Decode(&config); err != nil {
			logger.Fatalf("decode chaos_config.json error: %v", err)
		}
		callSite = config.CallSite
		errorCode = config.ErrorCode
		waitTime = config.WaitTime
		testId = config.TestId
	}
	return callSite, errorCode, waitTime, testId
}
