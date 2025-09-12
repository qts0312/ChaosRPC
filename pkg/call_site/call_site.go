package call_site

import (
	"runtime"
	"strconv"
	"strings"
)

type CallSite []string

func GetCallSite() CallSite {
	funcNames := make([]string, 0)
	linenos := make([]int, 0)
	for i := 0; ; i++ {
		pc, _, lineno, ok := runtime.Caller(i)
		if !ok {
			break
		}
		funcNames = append(funcNames, runtime.FuncForPC(pc).Name())
		linenos = append(linenos, lineno)
	}

	callSite := make([]string, 0)
	for i := len(funcNames) - 1; i >= 0; i-- {
		if strings.HasSuffix(funcNames[i], "Invoke") || strings.HasSuffix(funcNames[i], "NewStream") {
			// Skip common parts of gRPC call stack
			break
		}
		callSite = append(callSite, funcNames[i]+":"+strconv.Itoa(linenos[i]))
	}

	return callSite
}
