package rsw

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/url"
)

//PrefillTemplateReq represents PrefillTemplat method response.
//	@todo MergeFields
type PrefillTemplateReq struct {
	XMLName xml.Name `xml:"template"`

	GUID             string    `xml:"guid"`
	Action           string    `xml:"action"`
	Subject          string    `xml:"subject,omitempty"`
	Roles            []RoleReq `xml:"roles>role,omitempty"`
	Description      string    `xml:"description,omitempty"`
	ExpiresIn        string    `xml:"expires_in,omitempty"`
	Tags             []TagReq  `xml:"tags>tag,omitempty"`
	CallbackLocation string    `xml:"callback_location,omitempty"`
}

//ListTemplatesResp represents ListTemplates method request parameters.
type ListTemplatesResp struct {
	Templates      []TemplateResp `xml:"templates>template"`
	TotalTemplates int64          `xml:"total-templates"`
	TotalPages     int64          `xml:"total-pages"`
	PerPage        int64          `xml:"per-page"`
	CurrentPage    int64          `xml:"current-page"`
}

//TemplateResp represents Template entity.
//@todo MergeFields
//			Tags
type TemplateResp struct {
	Type            string `xml:"type"`
	GUID            string `xml:"guid"`
	CreatedAt       string `xml:"created-at"`
	Filename        string `xml:"filename"`
	Size            int64  `xml:"size"`
	ContentType     string `xml:"content-type"`
	PageCount       int64  `xml:"page-count"`
	Subject         string `xml:"subject"`
	Message         string `xml:"message"`
	Tags            string `xml:"tags"`
	ProcessingState string `xml:"processing-state"`
	ThumbnailURL    string `xml:"thumbnail-url"`

	Roles         []RoleResp `xml:"roles>role"`
	Pages         []PageResp `xml:"pages>page"`
	RedirectToken string     `xml:"redirect-token"`
}

//PrefillAndSendTemplateResp represents PrefillAndSendTemplate response.
type PrefillAndSendTemplateResp struct {
	Status string `xml:"status"`
	GUID   string `xml:"guid"`
}

//ListTemplates perform request to the same-name API endpoint.
//List Template has some arguments (see documentation).
//If you don't wanna send any params just send empty struct url.Values{}
func (a RightSignatureAPI) ListTemplates(params url.Values) (resp ListTemplatesResp, err error) {

	encodeParams := params.Encode()
	if len(encodeParams) > 0 {
		encodeParams = "?" + encodeParams
	}

	url := BaseURL + "/templates.xml" + encodeParams
	method := "GET"

	req, err := a.newRequest(method, url, nil)
	if err != nil {
		return
	}

	httpResp, err := a.do(req)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//	@todo skiped argument callback_location
//PrepackageTemplate perform request to the same-name API endpoint.
func (a RightSignatureAPI) PrepackageTemplate(guid string) (resp TemplateResp, err error) {
	url := BaseURL + "/templates/" + guid + "/prepackage.xml"
	method := "POST"

	req, err := a.newRequest(method, url, nil)
	if err != nil {
		return
	}

	httpResp, err := a.do(req)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//PrefillTemplate perform 'prefill' action request to the Prefill Template API endpoint.
func (a RightSignatureAPI) PrefillTemplate(p PrefillTemplateReq) (resp TemplateResp, err error) {
	err = a.prefillValidation(p)
	if err != nil {
		return
	}

	if p.Action != "prefill" {
		err = errors.New("For this method action must be 'prefill'")
		return
	}

	body, err := a.prefillTemplateBody(p)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//PrefillAndSendTemplate perform 'send' action request to the Prefill Template API endpoint.
func (a RightSignatureAPI) PrefillAndSendTemplate(p PrefillTemplateReq) (resp PrefillAndSendTemplateResp, err error) {
	err = a.prefillValidation(p)
	if err != nil {
		return
	}

	if p.Action != "send" {
		err = errors.New("For this method action must be 'send'")
		return
	}

	body, err := a.prefillTemplateBody(p)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//prefillTemplateBody perform request to Prefill Template API endpoint and returns the request body.
func (a RightSignatureAPI) prefillTemplateBody(p PrefillTemplateReq) (body []byte, err error) {
	url := BaseURL + "/templates.xml"
	method := "POST"

	xmlData, err := xml.Marshal(p)
	if err != nil {
		return
	}
	xmlData = append([]byte(xml.Header), xmlData...)

	req, err := a.newRequest(method, url, bytes.NewBuffer(xmlData))
	if err != nil {
		return
	}

	httpResp, err := a.do(req)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()

	body, err = ioutil.ReadAll(httpResp.Body)
	return
}

//prefillValidation represents validation for PrefillTemplateReq.
func (a RightSignatureAPI) prefillValidation(p PrefillTemplateReq) error {
	if len(p.GUID) < 1 || len(p.Action) < 1 {
		return errors.New("Requred arguments missed (check guid, action).")
	}
	return nil
}
