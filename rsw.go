/*
Package rsw implements a wrapper for rightsignature.com API for Go (Golang).
*/
package rsw

import (
	"encoding/xml"
	"io"
	"net/http"
)

const (
	//BaseURL for all API requests.
	BaseURL = "https://rightsignature.com/api"
)

//RightSignatureAPI struct represents API for communicate with RightSignature.
type RightSignatureAPI struct {
	apiToken string
}

//TagReq represents Tag's data for request.
type TagReq struct {
	XMLName xml.Name `xml:"tag"`

	Name  string `xml:"name"`
	Value string `xml:"value,omitempty"`
}

//RoleReq represents Role's data for request.
type RoleReq struct {
	XMLName xml.Name `xml:"role"`

	Name     string `xml:"name"`
	Email    string `xml:"email"`
	Locked   bool   `xml:"locked,omitempty"`
	RoleName string `xml:"role_name,attr,omitempty"`
	RoleID   string `xml:"role_id,attr,omitempty"`
}

//RoleResp represents Role's data for response.
type RoleResp struct {
	Role           string `xml:"role"`
	Name           string `xml:"name"`
	MustSign       string `xml:"must-sign"`
	DocumentRoleID string `xml:"document-role-id"`
	IsSender       string `xml:"is-sender"`
}

//PageResp represents Page's data for response.
type PageResp struct {
	PageNumber               int64  `xml:"page-number"`
	OriginalTemplateGUID     string `xml:"original-template-guid"`
	OriginalTemplateFilename string `xml:"original-template-filename"`
}

//NewAPI returns new API instance.
func NewAPI(apiToken string) *RightSignatureAPI {
	return &RightSignatureAPI{
		apiToken: apiToken,
	}
}

//newRequest returns new Request instance.
func (a RightSignatureAPI) newRequest(method, urlStr string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, urlStr, body)
	if err != nil {
		return
	}
	req.Header.Add("api-token", a.apiToken)
	req.Header.Add("Content-Type", "application/xml")
	return
}

//do perform http Request.
func (a RightSignatureAPI) do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}
