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

package enchant

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2EnchantRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2EnchantRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeNamespaces(
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
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) CreateNamespaceAsync(
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
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
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
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CreateNamespace(
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
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) GetNamespaceStatusAsync(
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetNamespaceStatus(
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
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) GetNamespaceAsync(
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetNamespace(
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
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) UpdateNamespaceAsync(
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
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
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
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) UpdateNamespace(
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
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) DeleteNamespaceAsync(
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DeleteNamespace(
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

func dumpUserDataByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DumpUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- DumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/dump"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DumpUserDataByUserId(
	request *DumpUserDataByUserIdRequest,
) (*DumpUserDataByUserIdResult, error) {
	callback := make(chan DumpUserDataByUserIdAsyncResult, 1)
	go p.DumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkDumpUserDataByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckDumpUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckDumpUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckDumpUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckDumpUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/dump"
	if request.UserId != nil && *request.UserId != "" {
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CheckDumpUserDataByUserId(
	request *CheckDumpUserDataByUserIdRequest,
) (*CheckDumpUserDataByUserIdResult, error) {
	callback := make(chan CheckDumpUserDataByUserIdAsyncResult, 1)
	go p.CheckDumpUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func cleanUserDataByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CleanUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/clean"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CleanUserDataByUserId(
	request *CleanUserDataByUserIdRequest,
) (*CleanUserDataByUserIdResult, error) {
	callback := make(chan CleanUserDataByUserIdAsyncResult, 1)
	go p.CleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkCleanUserDataByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckCleanUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckCleanUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckCleanUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckCleanUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/clean"
	if request.UserId != nil && *request.UserId != "" {
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CheckCleanUserDataByUserId(
	request *CheckCleanUserDataByUserIdRequest,
) (*CheckCleanUserDataByUserIdResult, error) {
	callback := make(chan CheckCleanUserDataByUserIdAsyncResult, 1)
	go p.CheckCleanUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBalanceParameterModelsAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBalanceParameterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeBalanceParameterModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBalanceParameterModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBalanceParameterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeBalanceParameterModelsAsync(
	request *DescribeBalanceParameterModelsRequest,
	callback chan<- DescribeBalanceParameterModelsAsyncResult,
) {
	path := "/{namespaceName}/model/balance"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
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

	go describeBalanceParameterModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeBalanceParameterModels(
	request *DescribeBalanceParameterModelsRequest,
) (*DescribeBalanceParameterModelsResult, error) {
	callback := make(chan DescribeBalanceParameterModelsAsyncResult, 1)
	go p.DescribeBalanceParameterModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBalanceParameterModelAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetBalanceParameterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterModelResult
	if asyncResult.Err != nil {
		callback <- GetBalanceParameterModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBalanceParameterModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBalanceParameterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetBalanceParameterModelAsync(
	request *GetBalanceParameterModelRequest,
	callback chan<- GetBalanceParameterModelAsyncResult,
) {
	path := "/{namespaceName}/model/balance/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBalanceParameterModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetBalanceParameterModel(
	request *GetBalanceParameterModelRequest,
) (*GetBalanceParameterModelResult, error) {
	callback := make(chan GetBalanceParameterModelAsyncResult, 1)
	go p.GetBalanceParameterModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBalanceParameterModelMastersAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBalanceParameterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeBalanceParameterModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBalanceParameterModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBalanceParameterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeBalanceParameterModelMastersAsync(
	request *DescribeBalanceParameterModelMastersRequest,
	callback chan<- DescribeBalanceParameterModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/balance"
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeBalanceParameterModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeBalanceParameterModelMasters(
	request *DescribeBalanceParameterModelMastersRequest,
) (*DescribeBalanceParameterModelMastersResult, error) {
	callback := make(chan DescribeBalanceParameterModelMastersAsyncResult, 1)
	go p.DescribeBalanceParameterModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createBalanceParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- CreateBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBalanceParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateBalanceParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateBalanceParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) CreateBalanceParameterModelMasterAsync(
	request *CreateBalanceParameterModelMasterRequest,
	callback chan<- CreateBalanceParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/balance"
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
	if request.TotalValue != nil {
		bodies["totalValue"] = *request.TotalValue
	}
	if request.InitialValueStrategy != nil && *request.InitialValueStrategy != "" {
		bodies["initialValueStrategy"] = *request.InitialValueStrategy
	}
	if request.Parameters != nil {
		var _parameters []interface{}
		for _, item := range request.Parameters {
			_parameters = append(_parameters, item)
		}
		bodies["parameters"] = _parameters
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createBalanceParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CreateBalanceParameterModelMaster(
	request *CreateBalanceParameterModelMasterRequest,
) (*CreateBalanceParameterModelMasterResult, error) {
	callback := make(chan CreateBalanceParameterModelMasterAsyncResult, 1)
	go p.CreateBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBalanceParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetBalanceParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBalanceParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetBalanceParameterModelMasterAsync(
	request *GetBalanceParameterModelMasterRequest,
	callback chan<- GetBalanceParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/balance/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBalanceParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetBalanceParameterModelMaster(
	request *GetBalanceParameterModelMasterRequest,
) (*GetBalanceParameterModelMasterResult, error) {
	callback := make(chan GetBalanceParameterModelMasterAsyncResult, 1)
	go p.GetBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBalanceParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBalanceParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateBalanceParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBalanceParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) UpdateBalanceParameterModelMasterAsync(
	request *UpdateBalanceParameterModelMasterRequest,
	callback chan<- UpdateBalanceParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/balance/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.TotalValue != nil {
		bodies["totalValue"] = *request.TotalValue
	}
	if request.InitialValueStrategy != nil && *request.InitialValueStrategy != "" {
		bodies["initialValueStrategy"] = *request.InitialValueStrategy
	}
	if request.Parameters != nil {
		var _parameters []interface{}
		for _, item := range request.Parameters {
			_parameters = append(_parameters, item)
		}
		bodies["parameters"] = _parameters
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateBalanceParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) UpdateBalanceParameterModelMaster(
	request *UpdateBalanceParameterModelMasterRequest,
) (*UpdateBalanceParameterModelMasterResult, error) {
	callback := make(chan UpdateBalanceParameterModelMasterAsyncResult, 1)
	go p.UpdateBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBalanceParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBalanceParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBalanceParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBalanceParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteBalanceParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBalanceParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBalanceParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DeleteBalanceParameterModelMasterAsync(
	request *DeleteBalanceParameterModelMasterRequest,
	callback chan<- DeleteBalanceParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/balance/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteBalanceParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DeleteBalanceParameterModelMaster(
	request *DeleteBalanceParameterModelMasterRequest,
) (*DeleteBalanceParameterModelMasterResult, error) {
	callback := make(chan DeleteBalanceParameterModelMasterAsyncResult, 1)
	go p.DeleteBalanceParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRarityParameterModelsAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRarityParameterModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeRarityParameterModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRarityParameterModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRarityParameterModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeRarityParameterModelsAsync(
	request *DescribeRarityParameterModelsRequest,
	callback chan<- DescribeRarityParameterModelsAsyncResult,
) {
	path := "/{namespaceName}/model/rarity"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
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

	go describeRarityParameterModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeRarityParameterModels(
	request *DescribeRarityParameterModelsRequest,
) (*DescribeRarityParameterModelsResult, error) {
	callback := make(chan DescribeRarityParameterModelsAsyncResult, 1)
	go p.DescribeRarityParameterModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRarityParameterModelAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetRarityParameterModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterModelResult
	if asyncResult.Err != nil {
		callback <- GetRarityParameterModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRarityParameterModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRarityParameterModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetRarityParameterModelAsync(
	request *GetRarityParameterModelRequest,
	callback chan<- GetRarityParameterModelAsyncResult,
) {
	path := "/{namespaceName}/model/rarity/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getRarityParameterModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetRarityParameterModel(
	request *GetRarityParameterModelRequest,
) (*GetRarityParameterModelResult, error) {
	callback := make(chan GetRarityParameterModelAsyncResult, 1)
	go p.GetRarityParameterModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRarityParameterModelMastersAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRarityParameterModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeRarityParameterModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRarityParameterModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRarityParameterModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeRarityParameterModelMastersAsync(
	request *DescribeRarityParameterModelMastersRequest,
	callback chan<- DescribeRarityParameterModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/rarity"
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeRarityParameterModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeRarityParameterModelMasters(
	request *DescribeRarityParameterModelMastersRequest,
) (*DescribeRarityParameterModelMastersResult, error) {
	callback := make(chan DescribeRarityParameterModelMastersAsyncResult, 1)
	go p.DescribeRarityParameterModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRarityParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRarityParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateRarityParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRarityParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) CreateRarityParameterModelMasterAsync(
	request *CreateRarityParameterModelMasterRequest,
	callback chan<- CreateRarityParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/rarity"
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
	if request.MaximumParameterCount != nil {
		bodies["maximumParameterCount"] = *request.MaximumParameterCount
	}
	if request.ParameterCounts != nil {
		var _parameterCounts []interface{}
		for _, item := range request.ParameterCounts {
			_parameterCounts = append(_parameterCounts, item)
		}
		bodies["parameterCounts"] = _parameterCounts
	}
	if request.Parameters != nil {
		var _parameters []interface{}
		for _, item := range request.Parameters {
			_parameters = append(_parameters, item)
		}
		bodies["parameters"] = _parameters
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createRarityParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) CreateRarityParameterModelMaster(
	request *CreateRarityParameterModelMasterRequest,
) (*CreateRarityParameterModelMasterResult, error) {
	callback := make(chan CreateRarityParameterModelMasterAsyncResult, 1)
	go p.CreateRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRarityParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetRarityParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRarityParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetRarityParameterModelMasterAsync(
	request *GetRarityParameterModelMasterRequest,
	callback chan<- GetRarityParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/rarity/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getRarityParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetRarityParameterModelMaster(
	request *GetRarityParameterModelMasterRequest,
) (*GetRarityParameterModelMasterResult, error) {
	callback := make(chan GetRarityParameterModelMasterAsyncResult, 1)
	go p.GetRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRarityParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRarityParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateRarityParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRarityParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) UpdateRarityParameterModelMasterAsync(
	request *UpdateRarityParameterModelMasterRequest,
	callback chan<- UpdateRarityParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/rarity/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MaximumParameterCount != nil {
		bodies["maximumParameterCount"] = *request.MaximumParameterCount
	}
	if request.ParameterCounts != nil {
		var _parameterCounts []interface{}
		for _, item := range request.ParameterCounts {
			_parameterCounts = append(_parameterCounts, item)
		}
		bodies["parameterCounts"] = _parameterCounts
	}
	if request.Parameters != nil {
		var _parameters []interface{}
		for _, item := range request.Parameters {
			_parameters = append(_parameters, item)
		}
		bodies["parameters"] = _parameters
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateRarityParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) UpdateRarityParameterModelMaster(
	request *UpdateRarityParameterModelMasterRequest,
) (*UpdateRarityParameterModelMasterResult, error) {
	callback := make(chan UpdateRarityParameterModelMasterAsyncResult, 1)
	go p.UpdateRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRarityParameterModelMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRarityParameterModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRarityParameterModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRarityParameterModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteRarityParameterModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRarityParameterModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRarityParameterModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DeleteRarityParameterModelMasterAsync(
	request *DeleteRarityParameterModelMasterRequest,
	callback chan<- DeleteRarityParameterModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/rarity/{parameterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteRarityParameterModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DeleteRarityParameterModelMaster(
	request *DeleteRarityParameterModelMasterRequest,
) (*DeleteRarityParameterModelMasterResult, error) {
	callback := make(chan DeleteRarityParameterModelMasterAsyncResult, 1)
	go p.DeleteRarityParameterModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2EnchantRestClient,
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

func (p Gs2EnchantRestClient) ExportMasterAsync(
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) ExportMaster(
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

func getCurrentParameterMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentParameterMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentParameterMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentParameterMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentParameterMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentParameterMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentParameterMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetCurrentParameterMasterAsync(
	request *GetCurrentParameterMasterRequest,
	callback chan<- GetCurrentParameterMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
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

	go getCurrentParameterMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetCurrentParameterMaster(
	request *GetCurrentParameterMasterRequest,
) (*GetCurrentParameterMasterResult, error) {
	callback := make(chan GetCurrentParameterMasterAsyncResult, 1)
	go p.GetCurrentParameterMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentParameterMasterAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentParameterMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentParameterMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentParameterMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentParameterMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentParameterMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentParameterMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) UpdateCurrentParameterMasterAsync(
	request *UpdateCurrentParameterMasterRequest,
	callback chan<- UpdateCurrentParameterMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
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

	go updateCurrentParameterMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) UpdateCurrentParameterMaster(
	request *UpdateCurrentParameterMasterRequest,
) (*UpdateCurrentParameterMasterResult, error) {
	callback := make(chan UpdateCurrentParameterMasterAsyncResult, 1)
	go p.UpdateCurrentParameterMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentParameterMasterFromGitHubAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentParameterMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentParameterMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentParameterMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentParameterMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentParameterMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentParameterMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) UpdateCurrentParameterMasterFromGitHubAsync(
	request *UpdateCurrentParameterMasterFromGitHubRequest,
	callback chan<- UpdateCurrentParameterMasterFromGitHubAsyncResult,
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentParameterMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) UpdateCurrentParameterMasterFromGitHub(
	request *UpdateCurrentParameterMasterFromGitHubRequest,
) (*UpdateCurrentParameterMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentParameterMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentParameterMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBalanceParameterStatusesAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBalanceParameterStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterStatusesResult
	if asyncResult.Err != nil {
		callback <- DescribeBalanceParameterStatusesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBalanceParameterStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBalanceParameterStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeBalanceParameterStatusesAsync(
	request *DescribeBalanceParameterStatusesRequest,
	callback chan<- DescribeBalanceParameterStatusesAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/balance"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ParameterName != nil {
		queryStrings["parameterName"] = core.ToString(*request.ParameterName)
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

	go describeBalanceParameterStatusesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeBalanceParameterStatuses(
	request *DescribeBalanceParameterStatusesRequest,
) (*DescribeBalanceParameterStatusesResult, error) {
	callback := make(chan DescribeBalanceParameterStatusesAsyncResult, 1)
	go p.DescribeBalanceParameterStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBalanceParameterStatusesByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBalanceParameterStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBalanceParameterStatusesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBalanceParameterStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeBalanceParameterStatusesByUserIdAsync(
	request *DescribeBalanceParameterStatusesByUserIdRequest,
	callback chan<- DescribeBalanceParameterStatusesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/balance"
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

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ParameterName != nil {
		queryStrings["parameterName"] = core.ToString(*request.ParameterName)
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

	go describeBalanceParameterStatusesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeBalanceParameterStatusesByUserId(
	request *DescribeBalanceParameterStatusesByUserIdRequest,
) (*DescribeBalanceParameterStatusesByUserIdResult, error) {
	callback := make(chan DescribeBalanceParameterStatusesByUserIdAsyncResult, 1)
	go p.DescribeBalanceParameterStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBalanceParameterStatusAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetBalanceParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterStatusResult
	if asyncResult.Err != nil {
		callback <- GetBalanceParameterStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBalanceParameterStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBalanceParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetBalanceParameterStatusAsync(
	request *GetBalanceParameterStatusRequest,
	callback chan<- GetBalanceParameterStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/balance/{parameterName}/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go getBalanceParameterStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetBalanceParameterStatus(
	request *GetBalanceParameterStatusRequest,
) (*GetBalanceParameterStatusResult, error) {
	callback := make(chan GetBalanceParameterStatusAsyncResult, 1)
	go p.GetBalanceParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBalanceParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBalanceParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetBalanceParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBalanceParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetBalanceParameterStatusByUserIdAsync(
	request *GetBalanceParameterStatusByUserIdRequest,
	callback chan<- GetBalanceParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/balance/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBalanceParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetBalanceParameterStatusByUserId(
	request *GetBalanceParameterStatusByUserIdRequest,
) (*GetBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan GetBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.GetBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBalanceParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBalanceParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DeleteBalanceParameterStatusByUserIdAsync(
	request *DeleteBalanceParameterStatusByUserIdRequest,
	callback chan<- DeleteBalanceParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/balance/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteBalanceParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DeleteBalanceParameterStatusByUserId(
	request *DeleteBalanceParameterStatusByUserIdRequest,
) (*DeleteBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan DeleteBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.DeleteBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func reDrawBalanceParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- ReDrawBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawBalanceParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReDrawBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) ReDrawBalanceParameterStatusByUserIdAsync(
	request *ReDrawBalanceParameterStatusByUserIdRequest,
	callback chan<- ReDrawBalanceParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/balance/{parameterName}/{propertyId}/redraw"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.FixedParameterNames != nil {
		var _fixedParameterNames []interface{}
		for _, item := range request.FixedParameterNames {
			_fixedParameterNames = append(_fixedParameterNames, item)
		}
		bodies["fixedParameterNames"] = _fixedParameterNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go reDrawBalanceParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) ReDrawBalanceParameterStatusByUserId(
	request *ReDrawBalanceParameterStatusByUserIdRequest,
) (*ReDrawBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan ReDrawBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.ReDrawBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func reDrawBalanceParameterStatusByStampSheetAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- ReDrawBalanceParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawBalanceParameterStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReDrawBalanceParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) ReDrawBalanceParameterStatusByStampSheetAsync(
	request *ReDrawBalanceParameterStatusByStampSheetRequest,
	callback chan<- ReDrawBalanceParameterStatusByStampSheetAsyncResult,
) {
	path := "/stamp/balance/redraw"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go reDrawBalanceParameterStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) ReDrawBalanceParameterStatusByStampSheet(
	request *ReDrawBalanceParameterStatusByStampSheetRequest,
) (*ReDrawBalanceParameterStatusByStampSheetResult, error) {
	callback := make(chan ReDrawBalanceParameterStatusByStampSheetAsyncResult, 1)
	go p.ReDrawBalanceParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setBalanceParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- SetBalanceParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetBalanceParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetBalanceParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetBalanceParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetBalanceParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetBalanceParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) SetBalanceParameterStatusByUserIdAsync(
	request *SetBalanceParameterStatusByUserIdRequest,
	callback chan<- SetBalanceParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/balance/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ParameterValues != nil {
		var _parameterValues []interface{}
		for _, item := range request.ParameterValues {
			_parameterValues = append(_parameterValues, item)
		}
		bodies["parameterValues"] = _parameterValues
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setBalanceParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) SetBalanceParameterStatusByUserId(
	request *SetBalanceParameterStatusByUserIdRequest,
) (*SetBalanceParameterStatusByUserIdResult, error) {
	callback := make(chan SetBalanceParameterStatusByUserIdAsyncResult, 1)
	go p.SetBalanceParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setBalanceParameterStatusByStampSheetAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- SetBalanceParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetBalanceParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetBalanceParameterStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetBalanceParameterStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetBalanceParameterStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetBalanceParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) SetBalanceParameterStatusByStampSheetAsync(
	request *SetBalanceParameterStatusByStampSheetRequest,
	callback chan<- SetBalanceParameterStatusByStampSheetAsyncResult,
) {
	path := "/stamp/balance/set"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go setBalanceParameterStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) SetBalanceParameterStatusByStampSheet(
	request *SetBalanceParameterStatusByStampSheetRequest,
) (*SetBalanceParameterStatusByStampSheetResult, error) {
	callback := make(chan SetBalanceParameterStatusByStampSheetAsyncResult, 1)
	go p.SetBalanceParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRarityParameterStatusesAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRarityParameterStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterStatusesResult
	if asyncResult.Err != nil {
		callback <- DescribeRarityParameterStatusesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRarityParameterStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRarityParameterStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeRarityParameterStatusesAsync(
	request *DescribeRarityParameterStatusesRequest,
	callback chan<- DescribeRarityParameterStatusesAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/rarity"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ParameterName != nil {
		queryStrings["parameterName"] = core.ToString(*request.ParameterName)
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

	go describeRarityParameterStatusesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeRarityParameterStatuses(
	request *DescribeRarityParameterStatusesRequest,
) (*DescribeRarityParameterStatusesResult, error) {
	callback := make(chan DescribeRarityParameterStatusesAsyncResult, 1)
	go p.DescribeRarityParameterStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRarityParameterStatusesByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRarityParameterStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRarityParameterStatusesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRarityParameterStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DescribeRarityParameterStatusesByUserIdAsync(
	request *DescribeRarityParameterStatusesByUserIdRequest,
	callback chan<- DescribeRarityParameterStatusesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity"
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

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ParameterName != nil {
		queryStrings["parameterName"] = core.ToString(*request.ParameterName)
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

	go describeRarityParameterStatusesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DescribeRarityParameterStatusesByUserId(
	request *DescribeRarityParameterStatusesByUserIdRequest,
) (*DescribeRarityParameterStatusesByUserIdResult, error) {
	callback := make(chan DescribeRarityParameterStatusesByUserIdAsyncResult, 1)
	go p.DescribeRarityParameterStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRarityParameterStatusAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetRarityParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterStatusResult
	if asyncResult.Err != nil {
		callback <- GetRarityParameterStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRarityParameterStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRarityParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetRarityParameterStatusAsync(
	request *GetRarityParameterStatusRequest,
	callback chan<- GetRarityParameterStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/rarity/{parameterName}/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go getRarityParameterStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetRarityParameterStatus(
	request *GetRarityParameterStatusRequest,
) (*GetRarityParameterStatusResult, error) {
	callback := make(chan GetRarityParameterStatusAsyncResult, 1)
	go p.GetRarityParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- GetRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) GetRarityParameterStatusByUserIdAsync(
	request *GetRarityParameterStatusByUserIdRequest,
	callback chan<- GetRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) GetRarityParameterStatusByUserId(
	request *GetRarityParameterStatusByUserIdRequest,
) (*GetRarityParameterStatusByUserIdResult, error) {
	callback := make(chan GetRarityParameterStatusByUserIdAsyncResult, 1)
	go p.GetRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) DeleteRarityParameterStatusByUserIdAsync(
	request *DeleteRarityParameterStatusByUserIdRequest,
	callback chan<- DeleteRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) DeleteRarityParameterStatusByUserId(
	request *DeleteRarityParameterStatusByUserIdRequest,
) (*DeleteRarityParameterStatusByUserIdResult, error) {
	callback := make(chan DeleteRarityParameterStatusByUserIdAsyncResult, 1)
	go p.DeleteRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func reDrawRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- ReDrawRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReDrawRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) ReDrawRarityParameterStatusByUserIdAsync(
	request *ReDrawRarityParameterStatusByUserIdRequest,
	callback chan<- ReDrawRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}/redraw"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.FixedParameterNames != nil {
		var _fixedParameterNames []interface{}
		for _, item := range request.FixedParameterNames {
			_fixedParameterNames = append(_fixedParameterNames, item)
		}
		bodies["fixedParameterNames"] = _fixedParameterNames
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go reDrawRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) ReDrawRarityParameterStatusByUserId(
	request *ReDrawRarityParameterStatusByUserIdRequest,
) (*ReDrawRarityParameterStatusByUserIdResult, error) {
	callback := make(chan ReDrawRarityParameterStatusByUserIdAsyncResult, 1)
	go p.ReDrawRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func reDrawRarityParameterStatusByStampSheetAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- ReDrawRarityParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReDrawRarityParameterStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReDrawRarityParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) ReDrawRarityParameterStatusByStampSheetAsync(
	request *ReDrawRarityParameterStatusByStampSheetRequest,
	callback chan<- ReDrawRarityParameterStatusByStampSheetAsyncResult,
) {
	path := "/stamp/rarity/parameter/redraw"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go reDrawRarityParameterStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) ReDrawRarityParameterStatusByStampSheet(
	request *ReDrawRarityParameterStatusByStampSheetRequest,
) (*ReDrawRarityParameterStatusByStampSheetResult, error) {
	callback := make(chan ReDrawRarityParameterStatusByStampSheetAsyncResult, 1)
	go p.ReDrawRarityParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- AddRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) AddRarityParameterStatusByUserIdAsync(
	request *AddRarityParameterStatusByUserIdRequest,
	callback chan<- AddRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}/add"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Count != nil {
		bodies["count"] = *request.Count
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go addRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) AddRarityParameterStatusByUserId(
	request *AddRarityParameterStatusByUserIdRequest,
) (*AddRarityParameterStatusByUserIdResult, error) {
	callback := make(chan AddRarityParameterStatusByUserIdAsyncResult, 1)
	go p.AddRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addRarityParameterStatusByStampSheetAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- AddRarityParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRarityParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRarityParameterStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddRarityParameterStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddRarityParameterStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddRarityParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) AddRarityParameterStatusByStampSheetAsync(
	request *AddRarityParameterStatusByStampSheetRequest,
	callback chan<- AddRarityParameterStatusByStampSheetAsyncResult,
) {
	path := "/stamp/rarity/parameter/add"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go addRarityParameterStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) AddRarityParameterStatusByStampSheet(
	request *AddRarityParameterStatusByStampSheetRequest,
) (*AddRarityParameterStatusByStampSheetResult, error) {
	callback := make(chan AddRarityParameterStatusByStampSheetAsyncResult, 1)
	go p.AddRarityParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRarityParameterStatusAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRarityParameterStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusResult
	if asyncResult.Err != nil {
		callback <- VerifyRarityParameterStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRarityParameterStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRarityParameterStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatusAsync(
	request *VerifyRarityParameterStatusRequest,
	callback chan<- VerifyRarityParameterStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/status/rarity/{parameterName}/{propertyId}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ParameterValueName != nil && *request.ParameterValueName != "" {
		bodies["parameterValueName"] = *request.ParameterValueName
	}
	if request.ParameterCount != nil {
		bodies["parameterCount"] = *request.ParameterCount
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRarityParameterStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatus(
	request *VerifyRarityParameterStatusRequest,
) (*VerifyRarityParameterStatusResult, error) {
	callback := make(chan VerifyRarityParameterStatusAsyncResult, 1)
	go p.VerifyRarityParameterStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatusByUserIdAsync(
	request *VerifyRarityParameterStatusByUserIdRequest,
	callback chan<- VerifyRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ParameterValueName != nil && *request.ParameterValueName != "" {
		bodies["parameterValueName"] = *request.ParameterValueName
	}
	if request.ParameterCount != nil {
		bodies["parameterCount"] = *request.ParameterCount
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatusByUserId(
	request *VerifyRarityParameterStatusByUserIdRequest,
) (*VerifyRarityParameterStatusByUserIdResult, error) {
	callback := make(chan VerifyRarityParameterStatusByUserIdAsyncResult, 1)
	go p.VerifyRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyRarityParameterStatusByStampTaskAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyRarityParameterStatusByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyRarityParameterStatusByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyRarityParameterStatusByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatusByStampTaskAsync(
	request *VerifyRarityParameterStatusByStampTaskRequest,
	callback chan<- VerifyRarityParameterStatusByStampTaskAsyncResult,
) {
	path := "/stamp/rarity/parameter/verify"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampTask != nil && *request.StampTask != "" {
		bodies["stampTask"] = *request.StampTask
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go verifyRarityParameterStatusByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) VerifyRarityParameterStatusByStampTask(
	request *VerifyRarityParameterStatusByStampTaskRequest,
) (*VerifyRarityParameterStatusByStampTaskResult, error) {
	callback := make(chan VerifyRarityParameterStatusByStampTaskAsyncResult, 1)
	go p.VerifyRarityParameterStatusByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRarityParameterStatusByUserIdAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- SetRarityParameterStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRarityParameterStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRarityParameterStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetRarityParameterStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRarityParameterStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRarityParameterStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) SetRarityParameterStatusByUserIdAsync(
	request *SetRarityParameterStatusByUserIdRequest,
	callback chan<- SetRarityParameterStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/status/rarity/{parameterName}/{propertyId}"
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
	if request.ParameterName != nil && *request.ParameterName != "" {
		path = strings.ReplaceAll(path, "{parameterName}", core.ToString(*request.ParameterName))
	} else {
		path = strings.ReplaceAll(path, "{parameterName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ParameterValues != nil {
		var _parameterValues []interface{}
		for _, item := range request.ParameterValues {
			_parameterValues = append(_parameterValues, item)
		}
		bodies["parameterValues"] = _parameterValues
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setRarityParameterStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) SetRarityParameterStatusByUserId(
	request *SetRarityParameterStatusByUserIdRequest,
) (*SetRarityParameterStatusByUserIdResult, error) {
	callback := make(chan SetRarityParameterStatusByUserIdAsyncResult, 1)
	go p.SetRarityParameterStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRarityParameterStatusByStampSheetAsyncHandler(
	client Gs2EnchantRestClient,
	job *core.NetworkJob,
	callback chan<- SetRarityParameterStatusByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRarityParameterStatusByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRarityParameterStatusByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetRarityParameterStatusByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRarityParameterStatusByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRarityParameterStatusByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2EnchantRestClient) SetRarityParameterStatusByStampSheetAsync(
	request *SetRarityParameterStatusByStampSheetRequest,
	callback chan<- SetRarityParameterStatusByStampSheetAsyncResult,
) {
	path := "/stamp/rarity/parameter/set"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.StampSheet != nil && *request.StampSheet != "" {
		bodies["stampSheet"] = *request.StampSheet
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go setRarityParameterStatusByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("enchant").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2EnchantRestClient) SetRarityParameterStatusByStampSheet(
	request *SetRarityParameterStatusByStampSheetRequest,
) (*SetRarityParameterStatusByStampSheetResult, error) {
	callback := make(chan SetRarityParameterStatusByStampSheetAsyncResult, 1)
	go p.SetRarityParameterStatusByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
