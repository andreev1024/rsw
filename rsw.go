/*
Package rsw implements a wrapper for rightsignature.com API for Go (Golang).
*/
package rsw

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	BaseURL    = "https://rightsignature.com/api"
	MethodGet  = "GET"
	MethodPost = "POST"
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

	Name   string `xml:"name"`
	Email  string `xml:"email"`
	Locked bool   `xml:"locked,omitempty"`

	RoleName string `xml:"role_name,attr,omitempty"`
	RoleID   string `xml:"role_id,attr,omitempty"`
}

//MergeFieldReq represents Merge Field's data for request.
type MergeFieldReq struct {
	XMLName xml.Name `xml:"merge_field"`

	Value  string `xml:"value"`
	Locked bool   `xml:"locked,omitempty"`

	RoleName string `xml:"merge_field_name,attr,omitempty"`
	RoleID   string `xml:"merge_field_id,attr,omitempty"`
}

//Role represents Role's data getting from response.
type Role struct {
	Role           string `xml:"role"`
	Name           string `xml:"name"`
	MustSign       string `xml:"must-sign"`
	DocumentRoleID string `xml:"document-role-id"`
	IsSender       string `xml:"is-sender"`
}

//Page represents Page's data getting from response.
type Page struct {
	PageNumber               int64  `xml:"page-number"`
	OriginalTemplateGUID     string `xml:"original-template-guid"`
	OriginalTemplateFilename string `xml:"original-template-filename"`
}

//MergeField represents Marge field's data getting from response.
type MergeField struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	Page int64  `xml:"page"`
}

//Error represents Error's data getting from response.
type Error struct {
	Message string `xml:"message"`
}

//NewAPI returns new API instance.
func NewAPI(apiToken string) *RightSignatureAPI {
	return &RightSignatureAPI{
		apiToken: apiToken,
	}
}

//Send http request.
func (a RightSignatureAPI) Send(method, url string, reqData []byte) (respData []byte, err error) {
	var data []byte
	if len(reqData) > 1 {
		data = []byte(xml.Header)
		data = append(data, reqData...)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Add("api-token", a.apiToken)
	req.Header.Add("Content-Type", "application/xml")

	client := http.DefaultClient

	httpResp, err := client.Do(req)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	respData, err = ioutil.ReadAll(httpResp.Body)

	if httpResp.StatusCode != http.StatusOK {
		errorResp := Error{}
		err = xml.Unmarshal(respData, &errorResp)
		if err != nil {
			return
		}
		err = errors.New(errorResp.Message)
	}

	return
}
