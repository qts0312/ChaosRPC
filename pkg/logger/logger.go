package logger

import "log"

func Infof(format string, v ...any) {
	log.Printf("[ChaosRPC] info: "+format, v...)
}

func Errorf(format string, v ...any) {
	log.Printf("[ChaosRPC] error: "+format, v...)
}

func Fatalf(format string, v ...any) {
	log.Fatalf("[ChaosRPC] fatal: "+format, v...)
}
