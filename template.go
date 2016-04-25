package rsw

import (
	"encoding/xml"
	"errors"
	"net/url"
)

//Template represents Template entity.
//Notice: Tags implementation work only like a mock.
type Template struct {
	Type            string `xml:"type"`
	GUID            string `xml:"guid"`
	CreatedAt       string `xml:"created-at"`
	Filename        string `xml:"filename"`
	Size            int64  `xml:"size"`
	ContentType     string `xml:"content-type"`
	PageCount       string `xml:"page-count"`
	Subject         string `xml:"subject"`
	Message         string `xml:"message"`
	Tags            string `xml:"tags"`
	ProcessingState string `xml:"processing-state"`
	ThumbnailURL    string `xml:"thumbnail-url"`

	Roles         []Role       `xml:"roles>role"`
	Pages         []Page       `xml:"pages>page"`
	MergeFields   []MergeField `xml:"merge-fields>merge-field"`
	RedirectToken string       `xml:"redirect-token"`
}

//PrefillTemplateReq represents arguments for Prefill Template API endpoint.
type PrefillTemplateReq struct {
	XMLName xml.Name `xml:"template"`

	GUID             string          `xml:"guid"`
	Action           string          `xml:"action"`
	Subject          string          `xml:"subject,omitempty"`
	Roles            []RoleReq       `xml:"roles>role,omitempty"`
	Description      string          `xml:"description,omitempty"`
	ExpiresIn        string          `xml:"expires_in,omitempty"`
	Tags             []TagReq        `xml:"tags>tag,omitempty"`
	MergeFields      []MergeFieldReq `xml:"merge_fields>merge_field,omitempty"`
	CallbackLocation string          `xml:"callback_location,omitempty"`
}

//PrepackageTemplateReq represents arguments for Prepackage Template API endpoint.
type PrepackageTemplateReq struct {
	XMLName xml.Name `xml:"callback_location"`

	CallbackLocation string `xml:",chardata"`
}

//PrefillAndSendTemplateResp represents Prefill Template (send) API endpoint response.
type PrefillAndSendTemplateResp struct {
	Status string `xml:"status"`
	GUID   string `xml:"guid"`
}

//ListTemplatesResp represents List Templates API endpoint response.
type ListTemplatesResp struct {
	Templates      []Template `xml:"templates>template"`
	TotalTemplates int64      `xml:"total-templates"`
	TotalPages     int64      `xml:"total-pages"`
	PerPage        int64      `xml:"per-page"`
	CurrentPage    int64      `xml:"current-page"`
}

//ListTemplates perform request to the same-name API endpoint.
//It has some optional arguments (see documentation).
//p - optional and can be missed.
func (a RightSignatureAPI) ListTemplates(p ...url.Values) (resp ListTemplatesResp, err error) {
	params := url.Values{}
	if len(p) > 0 {
		params = p[0]
	}

	encodeParams := params.Encode()
	if len(encodeParams) > 0 {
		encodeParams = "?" + encodeParams
	}

	url := BaseURL + "/templates.xml" + encodeParams
	method := MethodGet

	body, err := a.Send(method, url, nil)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//PrepackageTemplate perform request to the same-name API endpoint.
//p - optional and can be missed.
func (a RightSignatureAPI) PrepackageTemplate(guid string, p ...PrepackageTemplateReq) (resp Template, err error) {
	var data []byte
	if len(p) > 0 {
		data, err = xml.Marshal(p[0])
		if err != nil {
			return
		}
	}

	url := BaseURL + "/templates/" + guid + "/prepackage.xml"
	method := MethodPost

	body, err := a.Send(method, url, data)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//PrefillTemplate perform 'prefill' action request to the Prefill Template API endpoint.
func (a RightSignatureAPI) PrefillTemplate(p PrefillTemplateReq) (resp Template, err error) {
	err = a.prefillValidation(p)
	if err != nil {
		return
	}

	if p.Action != "prefill" {
		err = errors.New("For this method action must be 'prefill'")
		return
	}

	url := BaseURL + "/templates.xml"
	method := MethodPost

	xmlData, err := xml.Marshal(p)
	if err != nil {
		return
	}

	body, err := a.Send(method, url, xmlData)
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

	url := BaseURL + "/templates.xml"
	method := MethodPost

	xmlData, err := xml.Marshal(p)
	if err != nil {
		return
	}

	body, err := a.Send(method, url, xmlData)
	if err != nil {
		return
	}

	err = xml.Unmarshal(body, &resp)
	return
}

//prefillValidation represents validation for PrefillTemplateReq.
func (a RightSignatureAPI) prefillValidation(p PrefillTemplateReq) error {
	if len(p.GUID) < 1 || len(p.Action) < 1 {
		return errors.New("Requred arguments missed (check guid, action).")
	}
	return nil
}
