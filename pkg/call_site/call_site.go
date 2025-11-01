package call_site

import (
	"runtime"
	"strings"
)

type CallSite []string

func GetCallSite() CallSite {
	funcNames := make([]string, 0)
	for i := 0; ; i++ {
		pc, _, _, ok := runtime.Caller(i)
		if !ok {
			break
		}
		funcNames = append(funcNames, runtime.FuncForPC(pc).Name())
	}

	callSite := make([]string, 0)
	for i := len(funcNames) - 1; i >= 0; i-- {
		if strings.HasSuffix(funcNames[i], "Invoke") || strings.HasSuffix(funcNames[i], "NewStream") {
			// Skip common parts of gRPC call stack
			break
		}
		callSite = append(callSite, funcNames[i])
	}

	return callSite
}
