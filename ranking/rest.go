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

package ranking

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2RankingRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2RankingRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeNamespaces(
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
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) CreateNamespaceAsync(
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) CreateNamespace(
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
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetNamespaceStatus(
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
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetNamespace(
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
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) UpdateNamespace(
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
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DeleteNamespace(
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

func describeCategoryModelsAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCategoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCategoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCategoryModelsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCategoryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCategoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeCategoryModelsAsync(
	request *DescribeCategoryModelsRequest,
	callback chan<- DescribeCategoryModelsAsyncResult,
) {
	path := "/{namespaceName}/category"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeCategoryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeCategoryModels(
	request *DescribeCategoryModelsRequest,
) (*DescribeCategoryModelsResult, error) {
	callback := make(chan DescribeCategoryModelsAsyncResult, 1)
	go p.DescribeCategoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCategoryModelAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetCategoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCategoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCategoryModelResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCategoryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCategoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetCategoryModelAsync(
	request *GetCategoryModelRequest,
	callback chan<- GetCategoryModelAsyncResult,
) {
	path := "/{namespaceName}/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCategoryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetCategoryModel(
	request *GetCategoryModelRequest,
) (*GetCategoryModelResult, error) {
	callback := make(chan GetCategoryModelAsyncResult, 1)
	go p.GetCategoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCategoryModelMastersAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCategoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCategoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCategoryModelMastersResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCategoryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCategoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeCategoryModelMastersAsync(
	request *DescribeCategoryModelMastersRequest,
	callback chan<- DescribeCategoryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/category"
	if request.NamespaceName != nil {
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeCategoryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeCategoryModelMasters(
	request *DescribeCategoryModelMastersRequest,
) (*DescribeCategoryModelMastersResult, error) {
	callback := make(chan DescribeCategoryModelMastersAsyncResult, 1)
	go p.DescribeCategoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createCategoryModelMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- CreateCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateCategoryModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateCategoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) CreateCategoryModelMasterAsync(
	request *CreateCategoryModelMasterRequest,
	callback chan<- CreateCategoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/category"
	if request.NamespaceName != nil {
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
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.UniqueByUserId != nil {
		bodies["uniqueByUserId"] = *request.UniqueByUserId
	}
	if request.CalculateIntervalMinutes != nil {
		bodies["calculateIntervalMinutes"] = *request.CalculateIntervalMinutes
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createCategoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) CreateCategoryModelMaster(
	request *CreateCategoryModelMasterRequest,
) (*CreateCategoryModelMasterResult, error) {
	callback := make(chan CreateCategoryModelMasterAsyncResult, 1)
	go p.CreateCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCategoryModelMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCategoryModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCategoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetCategoryModelMasterAsync(
	request *GetCategoryModelMasterRequest,
	callback chan<- GetCategoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCategoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetCategoryModelMaster(
	request *GetCategoryModelMasterRequest,
) (*GetCategoryModelMasterResult, error) {
	callback := make(chan GetCategoryModelMasterAsyncResult, 1)
	go p.GetCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCategoryModelMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCategoryModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCategoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) UpdateCategoryModelMasterAsync(
	request *UpdateCategoryModelMasterRequest,
	callback chan<- UpdateCategoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.UniqueByUserId != nil {
		bodies["uniqueByUserId"] = *request.UniqueByUserId
	}
	if request.CalculateIntervalMinutes != nil {
		bodies["calculateIntervalMinutes"] = *request.CalculateIntervalMinutes
	}
	if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
		bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
	}
	if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
		bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCategoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) UpdateCategoryModelMaster(
	request *UpdateCategoryModelMasterRequest,
) (*UpdateCategoryModelMasterResult, error) {
	callback := make(chan UpdateCategoryModelMasterAsyncResult, 1)
	go p.UpdateCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCategoryModelMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCategoryModelMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCategoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DeleteCategoryModelMasterAsync(
	request *DeleteCategoryModelMasterRequest,
	callback chan<- DeleteCategoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteCategoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DeleteCategoryModelMaster(
	request *DeleteCategoryModelMasterRequest,
) (*DeleteCategoryModelMasterResult, error) {
	callback := make(chan DeleteCategoryModelMasterAsyncResult, 1)
	go p.DeleteCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesByCategoryNameAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesByCategoryNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByCategoryNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByCategoryNameResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByCategoryNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesByCategoryNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeSubscribesByCategoryNameAsync(
	request *DescribeSubscribesByCategoryNameRequest,
	callback chan<- DescribeSubscribesByCategoryNameAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeSubscribesByCategoryNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeSubscribesByCategoryName(
	request *DescribeSubscribesByCategoryNameRequest,
) (*DescribeSubscribesByCategoryNameResult, error) {
	callback := make(chan DescribeSubscribesByCategoryNameAsyncResult, 1)
	go p.DescribeSubscribesByCategoryNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesByCategoryNameAndUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesByCategoryNameAndUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByCategoryNameAndUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeSubscribesByCategoryNameAndUserIdAsync(
	request *DescribeSubscribesByCategoryNameAndUserIdRequest,
	callback chan<- DescribeSubscribesByCategoryNameAndUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/category/{categoryName}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeSubscribesByCategoryNameAndUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeSubscribesByCategoryNameAndUserId(
	request *DescribeSubscribesByCategoryNameAndUserIdRequest,
) (*DescribeSubscribesByCategoryNameAndUserIdResult, error) {
	callback := make(chan DescribeSubscribesByCategoryNameAndUserIdAsyncResult, 1)
	go p.DescribeSubscribesByCategoryNameAndUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subscribeAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- SubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) SubscribeAsync(
	request *SubscribeRequest,
	callback chan<- SubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go subscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) Subscribe(
	request *SubscribeRequest,
) (*SubscribeResult, error) {
	callback := make(chan SubscribeAsyncResult, 1)
	go p.SubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subscribeByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) SubscribeByUserIdAsync(
	request *SubscribeByUserIdRequest,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go subscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) SubscribeByUserId(
	request *SubscribeByUserIdRequest,
) (*SubscribeByUserIdResult, error) {
	callback := make(chan SubscribeByUserIdAsyncResult, 1)
	go p.SubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetSubscribeAsync(
	request *GetSubscribeRequest,
	callback chan<- GetSubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getSubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetSubscribe(
	request *GetSubscribeRequest,
) (*GetSubscribeResult, error) {
	callback := make(chan GetSubscribeAsyncResult, 1)
	go p.GetSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetSubscribeByUserIdAsync(
	request *GetSubscribeByUserIdRequest,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetSubscribeByUserId(
	request *GetSubscribeByUserIdRequest,
) (*GetSubscribeByUserIdResult, error) {
	callback := make(chan GetSubscribeByUserIdAsyncResult, 1)
	go p.GetSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unsubscribeAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- UnsubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnsubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnsubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) UnsubscribeAsync(
	request *UnsubscribeRequest,
	callback chan<- UnsubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go unsubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) Unsubscribe(
	request *UnsubscribeRequest,
) (*UnsubscribeResult, error) {
	callback := make(chan UnsubscribeAsyncResult, 1)
	go p.UnsubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func unsubscribeByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UnsubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UnsubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) UnsubscribeByUserIdAsync(
	request *UnsubscribeByUserIdRequest,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/category/{categoryName}/target/{targetUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go unsubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) UnsubscribeByUserId(
	request *UnsubscribeByUserIdRequest,
) (*UnsubscribeByUserIdResult, error) {
	callback := make(chan UnsubscribeByUserIdAsyncResult, 1)
	go p.UnsubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeScoresAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeScoresResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeScoresAsync(
	request *DescribeScoresRequest,
	callback chan<- DescribeScoresAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/scorer/{scorerUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeScoresAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeScores(
	request *DescribeScoresRequest,
) (*DescribeScoresResult, error) {
	callback := make(chan DescribeScoresAsyncResult, 1)
	go p.DescribeScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeScoresByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeScoresByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeScoresByUserIdAsync(
	request *DescribeScoresByUserIdRequest,
	callback chan<- DescribeScoresByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/category/{categoryName}/scorer/{scorerUserId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeScoresByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeScoresByUserId(
	request *DescribeScoresByUserIdRequest,
) (*DescribeScoresByUserIdResult, error) {
	callback := make(chan DescribeScoresByUserIdAsyncResult, 1)
	go p.DescribeScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getScoreAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetScoreAsync(
	request *GetScoreRequest,
	callback chan<- GetScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/scorer/{scorerUserId}/score/{uniqueId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}
	if request.UniqueId != nil {
		path = strings.ReplaceAll(path, "{uniqueId}", core.ToString(*request.UniqueId))
	} else {
		path = strings.ReplaceAll(path, "{uniqueId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetScore(
	request *GetScoreRequest,
) (*GetScoreResult, error) {
	callback := make(chan GetScoreAsyncResult, 1)
	go p.GetScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getScoreByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetScoreByUserIdAsync(
	request *GetScoreByUserIdRequest,
	callback chan<- GetScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/category/{categoryName}/scorer/{scorerUserId}/score/{uniqueId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}
	if request.UniqueId != nil {
		path = strings.ReplaceAll(path, "{uniqueId}", core.ToString(*request.UniqueId))
	} else {
		path = strings.ReplaceAll(path, "{uniqueId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetScoreByUserId(
	request *GetScoreByUserIdRequest,
) (*GetScoreByUserIdResult, error) {
	callback := make(chan GetScoreByUserIdAsyncResult, 1)
	go p.GetScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRankingsAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeRankingsAsync(
	request *DescribeRankingsRequest,
	callback chan<- DescribeRankingsAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/ranking"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.StartIndex != nil {
		queryStrings["startIndex"] = core.ToString(*request.StartIndex)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeRankingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeRankings(
	request *DescribeRankingsRequest,
) (*DescribeRankingsResult, error) {
	callback := make(chan DescribeRankingsAsyncResult, 1)
	go p.DescribeRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRankingssByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRankingssByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingssByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingssByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingssByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRankingssByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeRankingssByUserIdAsync(
	request *DescribeRankingssByUserIdRequest,
	callback chan<- DescribeRankingssByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/category/{categoryName}/ranking"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.StartIndex != nil {
		queryStrings["startIndex"] = core.ToString(*request.StartIndex)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeRankingssByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeRankingssByUserId(
	request *DescribeRankingssByUserIdRequest,
) (*DescribeRankingssByUserIdResult, error) {
	callback := make(chan DescribeRankingssByUserIdAsyncResult, 1)
	go p.DescribeRankingssByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeNearRankingsAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeNearRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeNearRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeNearRankingsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeNearRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeNearRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) DescribeNearRankingsAsync(
	request *DescribeNearRankingsRequest,
	callback chan<- DescribeNearRankingsAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/ranking/near"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Score != nil {
		queryStrings["score"] = core.ToString(*request.Score)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeNearRankingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) DescribeNearRankings(
	request *DescribeNearRankingsRequest,
) (*DescribeNearRankingsResult, error) {
	callback := make(chan DescribeNearRankingsAsyncResult, 1)
	go p.DescribeNearRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRankingAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetRankingAsync(
	request *GetRankingRequest,
	callback chan<- GetRankingAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/ranking/scorer/{scorerUserId}/score/{uniqueId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}
	if request.UniqueId != nil {
		path = strings.ReplaceAll(path, "{uniqueId}", core.ToString(*request.UniqueId))
	} else {
		path = strings.ReplaceAll(path, "{uniqueId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getRankingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetRanking(
	request *GetRankingRequest,
) (*GetRankingResult, error) {
	callback := make(chan GetRankingAsyncResult, 1)
	go p.GetRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRankingByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetRankingByUserIdAsync(
	request *GetRankingByUserIdRequest,
	callback chan<- GetRankingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/category/{categoryName}/ranking/scorer/{scorerUserId}/score/{uniqueId}"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ScorerUserId != nil {
		path = strings.ReplaceAll(path, "{scorerUserId}", core.ToString(*request.ScorerUserId))
	} else {
		path = strings.ReplaceAll(path, "{scorerUserId}", "null")
	}
	if request.UniqueId != nil {
		path = strings.ReplaceAll(path, "{uniqueId}", core.ToString(*request.UniqueId))
	} else {
		path = strings.ReplaceAll(path, "{uniqueId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getRankingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetRankingByUserId(
	request *GetRankingByUserIdRequest,
) (*GetRankingByUserIdResult, error) {
	callback := make(chan GetRankingByUserIdAsyncResult, 1)
	go p.GetRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putScoreAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- PutScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutScoreResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) PutScoreAsync(
	request *PutScoreRequest,
	callback chan<- PutScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/category/{categoryName}/ranking"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Score != nil {
		bodies["score"] = *request.Score
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go putScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) PutScore(
	request *PutScoreRequest,
) (*PutScoreResult, error) {
	callback := make(chan PutScoreAsyncResult, 1)
	go p.PutScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putScoreByUserIdAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- PutScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutScoreByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) PutScoreByUserIdAsync(
	request *PutScoreByUserIdRequest,
	callback chan<- PutScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/category/{categoryName}/ranking"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}
	if request.UserId != nil {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Score != nil {
		bodies["score"] = *request.Score
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go putScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) PutScoreByUserId(
	request *PutScoreByUserIdRequest,
) (*PutScoreByUserIdResult, error) {
	callback := make(chan PutScoreByUserIdAsyncResult, 1)
	go p.PutScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2RankingRestClient,
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

func (p Gs2RankingRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) ExportMaster(
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

func getCurrentRankingMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRankingMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentRankingMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) GetCurrentRankingMasterAsync(
	request *GetCurrentRankingMasterRequest,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCurrentRankingMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) GetCurrentRankingMaster(
	request *GetCurrentRankingMasterRequest,
) (*GetCurrentRankingMasterResult, error) {
	callback := make(chan GetCurrentRankingMasterAsyncResult, 1)
	go p.GetCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRankingMasterAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRankingMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) UpdateCurrentRankingMasterAsync(
	request *UpdateCurrentRankingMasterRequest,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Settings != nil && *request.Settings != "" {
		bodies["settings"] = *request.Settings
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentRankingMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) UpdateCurrentRankingMaster(
	request *UpdateCurrentRankingMasterRequest,
) (*UpdateCurrentRankingMasterResult, error) {
	callback := make(chan UpdateCurrentRankingMasterAsyncResult, 1)
	go p.UpdateCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRankingMasterFromGitHubAsyncHandler(
	client Gs2RankingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterFromGitHubResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingRestClient) UpdateCurrentRankingMasterFromGitHubAsync(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
	if request.NamespaceName != nil {
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentRankingMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2RankingRestClient) UpdateCurrentRankingMasterFromGitHub(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
) (*UpdateCurrentRankingMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRankingMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRankingMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
