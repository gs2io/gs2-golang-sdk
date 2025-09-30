/*
Copyright 2016 Game Server Services, Inc. or its affiliates. All Rights
Reserved.

Licensed under the Apache License, Version 2.0 (the "License").
You may not use this file except in compliance with the License.
A copy of the License is located at

 http://www.apache.org/licenses/LICENSE-2.0

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/

package megaField

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

var EndpointHost *string = nil

type Gs2MegaFieldRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2MegaFieldRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeNamespacesResult
	if asyncResult.Err != nil {
		callback <- DescribeNamespacesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeNamespacesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeNamespacesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DescribeNamespaces(
	request *DescribeNamespacesRequest,
) (*DescribeNamespacesResult, error) {
	callback := make(chan DescribeNamespacesAsyncResult, 1)
	go p.DescribeNamespacesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createNamespaceAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateNamespaceResult
	if asyncResult.Err != nil {
		callback <- CreateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) CreateNamespace(
	request *CreateNamespaceRequest,
) (*CreateNamespaceResult, error) {
	callback := make(chan CreateNamespaceAsyncResult, 1)
	go p.CreateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getNamespaceStatusAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceStatusResult
	if asyncResult.Err != nil {
		callback <- GetNamespaceStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetNamespaceStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetNamespaceStatus(
	request *GetNamespaceStatusRequest,
) (*GetNamespaceStatusResult, error) {
	callback := make(chan GetNamespaceStatusAsyncResult, 1)
	go p.GetNamespaceStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getNamespaceAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetNamespaceResult
	if asyncResult.Err != nil {
		callback <- GetNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetNamespace(
	request *GetNamespaceRequest,
) (*GetNamespaceResult, error) {
	callback := make(chan GetNamespaceAsyncResult, 1)
	go p.GetNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateNamespaceAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateNamespaceResult
	if asyncResult.Err != nil {
		callback <- UpdateNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) UpdateNamespace(
	request *UpdateNamespaceRequest,
) (*UpdateNamespaceResult, error) {
	callback := make(chan UpdateNamespaceAsyncResult, 1)
	go p.UpdateNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteNamespaceAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteNamespaceResult
	if asyncResult.Err != nil {
		callback <- DeleteNamespaceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteNamespaceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteNamespaceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DeleteNamespace(
	request *DeleteNamespaceRequest,
) (*DeleteNamespaceResult, error) {
	callback := make(chan DeleteNamespaceAsyncResult, 1)
	go p.DeleteNamespaceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getServiceVersionAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetServiceVersionResult
	if asyncResult.Err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	path := "/system/version"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getServiceVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetServiceVersion(
	request *GetServiceVersionRequest,
) (*GetServiceVersionResult, error) {
	callback := make(chan GetServiceVersionAsyncResult, 1)
	go p.GetServiceVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAreaModelsAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAreaModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAreaModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAreaModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeAreaModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAreaModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAreaModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DescribeAreaModelsAsync(
	request *DescribeAreaModelsRequest,
	callback chan<- DescribeAreaModelsAsyncResult,
) {
	path := "/{namespaceName}/area"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeAreaModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DescribeAreaModels(
	request *DescribeAreaModelsRequest,
) (*DescribeAreaModelsResult, error) {
	callback := make(chan DescribeAreaModelsAsyncResult, 1)
	go p.DescribeAreaModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAreaModelAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetAreaModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAreaModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAreaModelResult
	if asyncResult.Err != nil {
		callback <- GetAreaModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAreaModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAreaModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetAreaModelAsync(
	request *GetAreaModelRequest,
	callback chan<- GetAreaModelAsyncResult,
) {
	path := "/{namespaceName}/area/{areaModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getAreaModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetAreaModel(
	request *GetAreaModelRequest,
) (*GetAreaModelResult, error) {
	callback := make(chan GetAreaModelAsyncResult, 1)
	go p.GetAreaModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAreaModelMastersAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAreaModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAreaModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAreaModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeAreaModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAreaModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAreaModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DescribeAreaModelMastersAsync(
	request *DescribeAreaModelMastersRequest,
	callback chan<- DescribeAreaModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/area"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeAreaModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DescribeAreaModelMasters(
	request *DescribeAreaModelMastersRequest,
) (*DescribeAreaModelMastersResult, error) {
	callback := make(chan DescribeAreaModelMastersAsyncResult, 1)
	go p.DescribeAreaModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createAreaModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- CreateAreaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateAreaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateAreaModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateAreaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateAreaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateAreaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) CreateAreaModelMasterAsync(
	request *CreateAreaModelMasterRequest,
	callback chan<- CreateAreaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createAreaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) CreateAreaModelMaster(
	request *CreateAreaModelMasterRequest,
) (*CreateAreaModelMasterResult, error) {
	callback := make(chan CreateAreaModelMasterAsyncResult, 1)
	go p.CreateAreaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAreaModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetAreaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAreaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAreaModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetAreaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAreaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAreaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetAreaModelMasterAsync(
	request *GetAreaModelMasterRequest,
	callback chan<- GetAreaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getAreaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetAreaModelMaster(
	request *GetAreaModelMasterRequest,
) (*GetAreaModelMasterResult, error) {
	callback := make(chan GetAreaModelMasterAsyncResult, 1)
	go p.GetAreaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateAreaModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateAreaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateAreaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateAreaModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateAreaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateAreaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateAreaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) UpdateAreaModelMasterAsync(
	request *UpdateAreaModelMasterRequest,
	callback chan<- UpdateAreaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateAreaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) UpdateAreaModelMaster(
	request *UpdateAreaModelMasterRequest,
) (*UpdateAreaModelMasterResult, error) {
	callback := make(chan UpdateAreaModelMasterAsyncResult, 1)
	go p.UpdateAreaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAreaModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAreaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAreaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAreaModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteAreaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteAreaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteAreaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DeleteAreaModelMasterAsync(
	request *DeleteAreaModelMasterRequest,
	callback chan<- DeleteAreaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteAreaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DeleteAreaModelMaster(
	request *DeleteAreaModelMasterRequest,
) (*DeleteAreaModelMasterResult, error) {
	callback := make(chan DeleteAreaModelMasterAsyncResult, 1)
	go p.DeleteAreaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLayerModelsAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLayerModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLayerModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLayerModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeLayerModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLayerModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLayerModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DescribeLayerModelsAsync(
	request *DescribeLayerModelsRequest,
	callback chan<- DescribeLayerModelsAsyncResult,
) {
	path := "/{namespaceName}/area/{areaModelName}/layer"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLayerModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DescribeLayerModels(
	request *DescribeLayerModelsRequest,
) (*DescribeLayerModelsResult, error) {
	callback := make(chan DescribeLayerModelsAsyncResult, 1)
	go p.DescribeLayerModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLayerModelAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetLayerModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLayerModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLayerModelResult
	if asyncResult.Err != nil {
		callback <- GetLayerModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLayerModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLayerModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetLayerModelAsync(
	request *GetLayerModelRequest,
	callback chan<- GetLayerModelAsyncResult,
) {
	path := "/{namespaceName}/area/{areaModelName}/layer/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLayerModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetLayerModel(
	request *GetLayerModelRequest,
) (*GetLayerModelResult, error) {
	callback := make(chan GetLayerModelAsyncResult, 1)
	go p.GetLayerModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLayerModelMastersAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLayerModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLayerModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLayerModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeLayerModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLayerModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLayerModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DescribeLayerModelMastersAsync(
	request *DescribeLayerModelMastersRequest,
	callback chan<- DescribeLayerModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}/layer"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLayerModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DescribeLayerModelMasters(
	request *DescribeLayerModelMastersRequest,
) (*DescribeLayerModelMastersResult, error) {
	callback := make(chan DescribeLayerModelMastersAsyncResult, 1)
	go p.DescribeLayerModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createLayerModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- CreateLayerModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateLayerModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateLayerModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateLayerModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateLayerModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateLayerModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) CreateLayerModelMasterAsync(
	request *CreateLayerModelMasterRequest,
	callback chan<- CreateLayerModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}/layer"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go createLayerModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) CreateLayerModelMaster(
	request *CreateLayerModelMasterRequest,
) (*CreateLayerModelMasterResult, error) {
	callback := make(chan CreateLayerModelMasterAsyncResult, 1)
	go p.CreateLayerModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLayerModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetLayerModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLayerModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLayerModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetLayerModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLayerModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLayerModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetLayerModelMasterAsync(
	request *GetLayerModelMasterRequest,
	callback chan<- GetLayerModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}/layer/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLayerModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetLayerModelMaster(
	request *GetLayerModelMasterRequest,
) (*GetLayerModelMasterResult, error) {
	callback := make(chan GetLayerModelMasterAsyncResult, 1)
	go p.GetLayerModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateLayerModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateLayerModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateLayerModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateLayerModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateLayerModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateLayerModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateLayerModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) UpdateLayerModelMasterAsync(
	request *UpdateLayerModelMasterRequest,
	callback chan<- UpdateLayerModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}/layer/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateLayerModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) UpdateLayerModelMaster(
	request *UpdateLayerModelMasterRequest,
) (*UpdateLayerModelMasterResult, error) {
	callback := make(chan UpdateLayerModelMasterAsyncResult, 1)
	go p.UpdateLayerModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLayerModelMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLayerModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLayerModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLayerModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteLayerModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteLayerModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteLayerModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) DeleteLayerModelMasterAsync(
	request *DeleteLayerModelMasterRequest,
	callback chan<- DeleteLayerModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/area/{areaModelName}/layer/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteLayerModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) DeleteLayerModelMaster(
	request *DeleteLayerModelMasterRequest,
) (*DeleteLayerModelMasterResult, error) {
	callback := make(chan DeleteLayerModelMasterAsyncResult, 1)
	go p.DeleteLayerModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExportMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExportMasterResult
	if asyncResult.Err != nil {
		callback <- ExportMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ExportMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) ExportMaster(
	request *ExportMasterRequest,
) (*ExportMasterResult, error) {
	callback := make(chan ExportMasterAsyncResult, 1)
	go p.ExportMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCurrentFieldMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentFieldMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentFieldMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentFieldMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentFieldMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentFieldMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentFieldMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) GetCurrentFieldMasterAsync(
	request *GetCurrentFieldMasterRequest,
	callback chan<- GetCurrentFieldMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getCurrentFieldMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) GetCurrentFieldMaster(
	request *GetCurrentFieldMasterRequest,
) (*GetCurrentFieldMasterResult, error) {
	callback := make(chan GetCurrentFieldMasterAsyncResult, 1)
	go p.GetCurrentFieldMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentFieldMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentFieldMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentFieldMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentFieldMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentFieldMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentFieldMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentFieldMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) PreUpdateCurrentFieldMasterAsync(
	request *PreUpdateCurrentFieldMasterRequest,
	callback chan<- PreUpdateCurrentFieldMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go preUpdateCurrentFieldMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) PreUpdateCurrentFieldMaster(
	request *PreUpdateCurrentFieldMasterRequest,
) (*PreUpdateCurrentFieldMasterResult, error) {
	callback := make(chan PreUpdateCurrentFieldMasterAsyncResult, 1)
	go p.PreUpdateCurrentFieldMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentFieldMasterAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentFieldMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFieldMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFieldMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentFieldMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentFieldMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentFieldMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) UpdateCurrentFieldMasterAsync(
	request *UpdateCurrentFieldMasterRequest,
	callback chan<- UpdateCurrentFieldMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentFieldMaster(
			&PreUpdateCurrentFieldMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentFieldMasterAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Settings))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- UpdateCurrentFieldMasterAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Settings = nil
	}
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateCurrentFieldMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) UpdateCurrentFieldMaster(
	request *UpdateCurrentFieldMasterRequest,
) (*UpdateCurrentFieldMasterResult, error) {
	callback := make(chan UpdateCurrentFieldMasterAsyncResult, 1)
	go p.UpdateCurrentFieldMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentFieldMasterFromGitHubAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentFieldMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFieldMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFieldMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentFieldMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentFieldMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentFieldMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) UpdateCurrentFieldMasterFromGitHubAsync(
	request *UpdateCurrentFieldMasterFromGitHubRequest,
	callback chan<- UpdateCurrentFieldMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CheckoutSetting != nil {
		bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateCurrentFieldMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) UpdateCurrentFieldMasterFromGitHub(
	request *UpdateCurrentFieldMasterFromGitHubRequest,
) (*UpdateCurrentFieldMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentFieldMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentFieldMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putPositionAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- PutPositionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutPositionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutPositionResult
	if asyncResult.Err != nil {
		callback <- PutPositionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutPositionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutPositionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) PutPositionAsync(
	request *PutPositionRequest,
	callback chan<- PutPositionAsyncResult,
) {
	path := "/{namespaceName}/user/me/spatial/{areaModelName}/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Position != nil {
		bodies["position"] = request.Position.ToDict()
	}
	if request.Vector != nil {
		bodies["vector"] = request.Vector.ToDict()
	}
	if request.R != nil {
		bodies["r"] = *request.R
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go putPositionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) PutPosition(
	request *PutPositionRequest,
) (*PutPositionResult, error) {
	callback := make(chan PutPositionAsyncResult, 1)
	go p.PutPositionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putPositionByUserIdAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- PutPositionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutPositionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutPositionByUserIdResult
	if asyncResult.Err != nil {
		callback <- PutPositionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutPositionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutPositionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) PutPositionByUserIdAsync(
	request *PutPositionByUserIdRequest,
	callback chan<- PutPositionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/spatial/{areaModelName}/{layerModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Position != nil {
		bodies["position"] = request.Position.ToDict()
	}
	if request.Vector != nil {
		bodies["vector"] = request.Vector.ToDict()
	}
	if request.R != nil {
		bodies["r"] = *request.R
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go putPositionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) PutPositionByUserId(
	request *PutPositionByUserIdRequest,
) (*PutPositionByUserIdResult, error) {
	callback := make(chan PutPositionByUserIdAsyncResult, 1)
	go p.PutPositionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func fetchPositionAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- FetchPositionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FetchPositionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FetchPositionResult
	if asyncResult.Err != nil {
		callback <- FetchPositionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FetchPositionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FetchPositionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) FetchPositionAsync(
	request *FetchPositionRequest,
	callback chan<- FetchPositionAsyncResult,
) {
	path := "/{namespaceName}/area/{areaModelName}/layer/{layerModelName}/spatial/fetch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserIds != nil {
		var _userIds []interface{}
		for _, item := range request.UserIds {
			_userIds = append(_userIds, item)
		}
		bodies["userIds"] = _userIds
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go fetchPositionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) FetchPosition(
	request *FetchPositionRequest,
) (*FetchPositionResult, error) {
	callback := make(chan FetchPositionAsyncResult, 1)
	go p.FetchPositionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func fetchPositionFromSystemAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- FetchPositionFromSystemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FetchPositionFromSystemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FetchPositionFromSystemResult
	if asyncResult.Err != nil {
		callback <- FetchPositionFromSystemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FetchPositionFromSystemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FetchPositionFromSystemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) FetchPositionFromSystemAsync(
	request *FetchPositionFromSystemRequest,
	callback chan<- FetchPositionFromSystemAsyncResult,
) {
	path := "/{namespaceName}/system/area/{areaModelName}/layer/{layerModelName}/spatial/fetch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserIds != nil {
		var _userIds []interface{}
		for _, item := range request.UserIds {
			_userIds = append(_userIds, item)
		}
		bodies["userIds"] = _userIds
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go fetchPositionFromSystemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) FetchPositionFromSystem(
	request *FetchPositionFromSystemRequest,
) (*FetchPositionFromSystemResult, error) {
	callback := make(chan FetchPositionFromSystemAsyncResult, 1)
	go p.FetchPositionFromSystemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func nearUserIdsAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- NearUserIdsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- NearUserIdsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result NearUserIdsResult
	if asyncResult.Err != nil {
		callback <- NearUserIdsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- NearUserIdsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- NearUserIdsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) NearUserIdsAsync(
	request *NearUserIdsRequest,
	callback chan<- NearUserIdsAsyncResult,
) {
	path := "/{namespaceName}/area/{areaModelName}/layer/{layerModelName}/spatial/near"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Point != nil {
		bodies["point"] = request.Point.ToDict()
	}
	if request.R != nil {
		bodies["r"] = *request.R
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go nearUserIdsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) NearUserIds(
	request *NearUserIdsRequest,
) (*NearUserIdsResult, error) {
	callback := make(chan NearUserIdsAsyncResult, 1)
	go p.NearUserIdsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func nearUserIdsFromSystemAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- NearUserIdsFromSystemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- NearUserIdsFromSystemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result NearUserIdsFromSystemResult
	if asyncResult.Err != nil {
		callback <- NearUserIdsFromSystemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- NearUserIdsFromSystemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- NearUserIdsFromSystemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) NearUserIdsFromSystemAsync(
	request *NearUserIdsFromSystemRequest,
	callback chan<- NearUserIdsFromSystemAsyncResult,
) {
	path := "/{namespaceName}/system/area/{areaModelName}/layer/{layerModelName}/spatial/near"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Point != nil {
		bodies["point"] = request.Point.ToDict()
	}
	if request.R != nil {
		bodies["r"] = *request.R
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go nearUserIdsFromSystemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) NearUserIdsFromSystem(
	request *NearUserIdsFromSystemRequest,
) (*NearUserIdsFromSystemResult, error) {
	callback := make(chan NearUserIdsFromSystemAsyncResult, 1)
	go p.NearUserIdsFromSystemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func actionAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- ActionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ActionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ActionResult
	if asyncResult.Err != nil {
		callback <- ActionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ActionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ActionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) ActionAsync(
	request *ActionRequest,
	callback chan<- ActionAsyncResult,
) {
	path := "/{namespaceName}/user/me/spatial/{areaModelName}/{layerModelName}/action"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Position != nil {
		bodies["position"] = request.Position.ToDict()
	}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go actionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) Action(
	request *ActionRequest,
) (*ActionResult, error) {
	callback := make(chan ActionAsyncResult, 1)
	go p.ActionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func actionByUserIdAsyncHandler(
	client Gs2MegaFieldRestClient,
	job *core.NetworkJob,
	callback chan<- ActionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ActionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ActionByUserIdResult
	if asyncResult.Err != nil {
		callback <- ActionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ActionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ActionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MegaFieldRestClient) ActionByUserIdAsync(
	request *ActionByUserIdRequest,
	callback chan<- ActionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/spatial/{areaModelName}/{layerModelName}/action"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.AreaModelName != nil && *request.AreaModelName != "" {
		path = strings.ReplaceAll(path, "{areaModelName}", core.ToString(*request.AreaModelName))
	} else {
		path = strings.ReplaceAll(path, "{areaModelName}", "null")
	}
	if request.LayerModelName != nil && *request.LayerModelName != "" {
		path = strings.ReplaceAll(path, "{layerModelName}", core.ToString(*request.LayerModelName))
	} else {
		path = strings.ReplaceAll(path, "{layerModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Position != nil {
		bodies["position"] = request.Position.ToDict()
	}
	if request.Scopes != nil {
		var _scopes []interface{}
		for _, item := range request.Scopes {
			_scopes = append(_scopes, item)
		}
		bodies["scopes"] = _scopes
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go actionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("mega-field", EndpointHost).AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2MegaFieldRestClient) ActionByUserId(
	request *ActionByUserIdRequest,
) (*ActionByUserIdResult, error) {
	callback := make(chan ActionByUserIdAsyncResult, 1)
	go p.ActionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
