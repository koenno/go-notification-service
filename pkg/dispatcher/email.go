package dispatcher

type Recipients struct {
	To  []string
	CC  []string
	BCC []string
}

type EMail struct {
	Recipients Recipients
	Subject    string
	Message    string
}
