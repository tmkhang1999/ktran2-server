package structs

type LogglyMessage struct {
	StatusCode  uint   `json:"status_code"`
	MethodType  string `json:"method_type"`
	SourceIp    string `json:"source_ip"`
	RequestPath string `json:"request_path"`
}
