Package **rsw** implements a wrapper for [RightSignature API](http://rightsignature.com/) for Go (Golang).

## Features
* Composite documents;
* Full configurable documents created from template (merged fields, tags etc.);
* API response error handling;

### API endpoints
Current version support following API endpoints:

#### Templates
* List Templates
* Prepackage Template
* Prefill Template

#### Documents
* Document Details

### Example
```
    //  Notice: error handling skipped for simplicity

  	apiToken := "123456789"
  	api := rsw.NewAPI(apiToken)

  	//  sets ListTemplate params (optional).
  	v := url.Values{}
  	v.Set("page", "1")
  	v.Set("per_page", "20")

  	listResponse, err := api.ListTemplates(v)

  	testTpl := listResponse.Templates[0]

  	//  single document PrepackageTemplate
  	guid := testTpl.GUID
  	//  composite documents PrepackageTemplate
  	//guid := testTpl.GUID + "," + testTpl.GUID

  	ptReq := rsw.PrepackageTemplateReq{CallbackLocation: "http://mytestsite.com"}

  	prepackResp, err := api.PrepackageTemplate(guid, ptReq)

    //  configure roles
  	var roles []rsw.RoleReq
  	roles = append(roles, rsw.RoleReq{
  		Name:     "Andrey",
  		Email:    "andreev1024@gmail.com",
  		Locked:   true,
  		RoleName: "Test User",
  	})

    //  configure tags
  	var tags []rsw.TagReq
  	tags = append(tags, rsw.TagReq{
  		Name: "Tag1 name",
  	})

  	tags = append(tags, rsw.TagReq{
  		Name:  "Tag2 name",
  		Value: "Tag2 value",
  	})

    //  configure merged fields
  	var mergedFields []rsw.MergeFieldReq
  	mergedFields = append(mergedFields, rsw.MergeFieldReq{
  		Value:    "custom name",
  		RoleName: "name",
  	})

  	mergedFields = append(mergedFields, rsw.MergeFieldReq{
  		Value:  "custom email",
  		RoleID: prepackResp.MergeFields[1].ID,
  		Locked: true,
  	})

  	prefillReq := rsw.PrefillTemplateReq{
  		GUID:        prepackResp.GUID,
  		Action:      "prefill",
  		Subject:     "my custom test subject",
  		Roles:       roles,
  		Tags:        tags,
  		MergeFields: mergedFields,
  	}

  	prefillResp, err := api.PrefillTemplate(prefillReq)

  	prefillReq.Action = "send"
  	prefillAndSendResp, err := api.PrefillAndSendTemplate(prefillReq)
}
```
