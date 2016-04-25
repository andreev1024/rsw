package rsw

import "encoding/xml"

//Document represents Document entity.
type Document struct {
	GUID              string       `xml:"guid"`
	CreatedAt         string       `xml:"created-at"`
	CompletedAt       string       `xml:"completed-at"`
	LastActivityAt    string       `xml:"last-activity-at"`
	ExpiresOn         string       `xml:"expires-on"`
	IsTrashed         string       `xml:"is-trashed"`
	Size              string       `xml:"size"`
	ContentType       string       `xml:"content-type"`
	OriginalFilename  string       `xml:"original-filename"`
	SignedPdfChecksum string       `xml:"signed-pdf-checksum"`
	Subject           string       `xml:"subject"`
	Message           string       `xml:"message"`
	ProcessingState   string       `xml:"processing-state"`
	MergeState        string       `xml:"merge-state"`
	State             string       `xml:"state"`
	CallbackLocation  string       `xml:"callback-location"`
	Tags              string       `xml:"tags"`
	Recipients        []Recipient  `xml:"recipients>recipient"`
	AuditTrails       []AuditTrail `xml:"audit-trails>audit-trail"`
	Pages             []Page       `xml:"pages>page"`
	FormFields        []FormField  `xml:"form-fields>form-field"`
	OriginalURL       string       `xml:"original-url"`
	PdfURL            string       `xml:"pdf-url"`
	ThumbnailURL      string       `xml:"thumbnail-url"`
	LargeURL          string       `xml:"large-url"`
	SignedPdfURL      string       `xml:"signed-pdf-url"`
}

//Recipient represents Recipient entity.
type Recipient struct {
	Name           string `xml:"name"`
	Email          string `xml:"email"`
	MustSign       string `xml:"must-sign"`
	DocumentRoleID string `xml:"document-role-id"`
	RoleID         string `xml:"role-id"`
	State          string `xml:"state"`
	IsSender       string `xml:"is-sender"`
	ViewedAt       string `xml:"viewed-at"`
	CompletedAt    string `xml:"completed-at"`
}

//AuditTrail represents AuditTrail entity.
type AuditTrail struct {
	Timestamp string `xml:"timestamp"`
	Message   string `xml:"message"`
}

//FormField represents FormField entity.
type FormField struct {
	ID     string `xml:"id"`
	Name   string `xml:"name"`
	RoleID string `xml:"role-id"`
	Value  string `xml:"value"`
	Page   string `xml:"page"`
}

//DocumentDetails perform request to the same-name API endpoint.
func (a RightSignatureAPI) DocumentDetails(guid string) (resp Document, err error) {
	url := BaseURL + "/documents/" + guid + ".xml"
	method := MethodGet

	body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}
