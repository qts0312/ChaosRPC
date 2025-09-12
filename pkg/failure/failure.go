package failure

const (
	ErrorNone = iota
	ErrorOutboundUnavailable
	ErrorInboundUnavailable
	ErrorInboundTimeout
)
