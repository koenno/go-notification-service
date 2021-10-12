package api

type Recipients struct {
	To  []string `json:"to"`
	CC  []string `json:"cc"`
	BCC []string `json:"bcc"`
}

type EMail struct {
	Recipients Recipients `json:"recipients"`
	Subject    string     `json:"subject"`
	Message    string     `json:"message"`
}
