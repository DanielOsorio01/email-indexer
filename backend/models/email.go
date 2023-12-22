package models

type Email struct {
	MessageID               string
	Date                    string
	From                    string
	To                      []string
	Subject                 string
	MimeVersion             string
	ContentType             string
	ContentTransferEncoding string
	XFrom                   string
	XTo                     []string
	XCC                     []string
	XBCC                    []string
	XFolder                 string
	XOrigin                 string
	XFileName               string
	Body                    string
}
