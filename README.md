Package **rsw** implements a wrapper for [RightSignature API](http://rightsignature.com/) for Go (Golang).

Example
```
  //  Notice: error handling skipped for simplicity

	apiToken := "mySmartApiToken"
	api := rsw.NewAPI(apiToken)

	//  sets ListTemplate params. If you don't wanna send any params just send empty struct (url.Values{})
	v := url.Values{}
	v.Set("page", "2")
	v.Set("per_page", "20")

	listResponse, err := api.ListTemplates(v)

	testTpl := listResponse.Templates[0]

	//  single document PrepackageTemplate
	guid := testTpl.GUID
	prepackResp, err := api.PrepackageTemplate(guid)

	guid = prepackResp.GUID

	var roles []rsw.RoleReq
	roles = append(roles, rsw.RoleReq{
		Name:     "Andrey",
		Email:    "andreev1024@gmail.com",
		Locked:   true,
		RoleName: "Test User",
	})

	var tags []rsw.TagReq
	tags = append(tags, rsw.TagReq{
		Name: "Tag1 name",
	})

	tags = append(tags, rsw.TagReq{
		Name:  "Tag2 name",
		Value: "Tag2 value",
	})

	prefillReq := rsw.PrefillTemplateReq{
		GUID:    guid,
		Action:  "prefill",
		Subject: "my custom test subject",
		Roles:   roles,
		Tags:    tags,
	}
	prefillResp, err := api.PrefillTemplate(prefillReq)

	prefillReq.Action = "send"
	prefillAndSendResp, err := api.PrefillAndSendTemplate(prefillReq)
}
```

#### Notice
Current version support following API endpoints:
##### Templates
* List Templates
* Prepackage Template
* Prefill Template
