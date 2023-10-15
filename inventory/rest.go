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

package inventory

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2InventoryRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2InventoryRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeNamespaces(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) CreateNamespaceAsync(
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
	if request.AcquireScript != nil {
		bodies["acquireScript"] = request.AcquireScript.ToDict()
	}
	if request.OverflowScript != nil {
		bodies["overflowScript"] = request.OverflowScript.ToDict()
	}
	if request.ConsumeScript != nil {
		bodies["consumeScript"] = request.ConsumeScript.ToDict()
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
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetNamespaceStatus(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) UpdateNamespaceAsync(
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
	if request.AcquireScript != nil {
		bodies["acquireScript"] = request.AcquireScript.ToDict()
	}
	if request.OverflowScript != nil {
		bodies["overflowScript"] = request.OverflowScript.ToDict()
	}
	if request.ConsumeScript != nil {
		bodies["consumeScript"] = request.ConsumeScript.ToDict()
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
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteNamespace(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DumpUserDataByUserId(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CheckDumpUserDataByUserId(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CleanUserDataByUserId(
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
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CheckCleanUserDataByUserId(
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

func prepareImportUserDataByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/import/prepare"
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

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) PrepareImportUserDataByUserId(
	request *PrepareImportUserDataByUserIdRequest,
) (*PrepareImportUserDataByUserIdResult, error) {
	callback := make(chan PrepareImportUserDataByUserIdAsyncResult, 1)
	go p.PrepareImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func importUserDataByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- ImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/import"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ImportUserDataByUserId(
	request *ImportUserDataByUserIdRequest,
) (*ImportUserDataByUserIdResult, error) {
	callback := make(chan ImportUserDataByUserIdAsyncResult, 1)
	go p.ImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkImportUserDataByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckImportUserDataByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckImportUserDataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckImportUserDataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckImportUserDataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	path := "/system/user/{userId}/import/{uploadToken}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		path = strings.ReplaceAll(path, "{uploadToken}", core.ToString(*request.UploadToken))
	} else {
		path = strings.ReplaceAll(path, "{uploadToken}", "null")
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

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CheckImportUserDataByUserId(
	request *CheckImportUserDataByUserIdRequest,
) (*CheckImportUserDataByUserIdResult, error) {
	callback := make(chan CheckImportUserDataByUserIdAsyncResult, 1)
	go p.CheckImportUserDataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoryModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoryModelMastersAsync(
	request *DescribeInventoryModelMastersRequest,
	callback chan<- DescribeInventoryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/inventory"
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

	go describeInventoryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoryModelMasters(
	request *DescribeInventoryModelMastersRequest,
) (*DescribeInventoryModelMastersResult, error) {
	callback := make(chan DescribeInventoryModelMastersAsyncResult, 1)
	go p.DescribeInventoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateInventoryModelMasterAsync(
	request *CreateInventoryModelMasterRequest,
	callback chan<- CreateInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory"
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
	if request.InitialCapacity != nil {
		bodies["initialCapacity"] = *request.InitialCapacity
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
	}
	if request.ProtectReferencedItem != nil {
		bodies["protectReferencedItem"] = *request.ProtectReferencedItem
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateInventoryModelMaster(
	request *CreateInventoryModelMasterRequest,
) (*CreateInventoryModelMasterResult, error) {
	callback := make(chan CreateInventoryModelMasterAsyncResult, 1)
	go p.CreateInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryModelMasterAsync(
	request *GetInventoryModelMasterRequest,
	callback chan<- GetInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryModelMaster(
	request *GetInventoryModelMasterRequest,
) (*GetInventoryModelMasterResult, error) {
	callback := make(chan GetInventoryModelMasterAsyncResult, 1)
	go p.GetInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateInventoryModelMasterAsync(
	request *UpdateInventoryModelMasterRequest,
	callback chan<- UpdateInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.InitialCapacity != nil {
		bodies["initialCapacity"] = *request.InitialCapacity
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
	}
	if request.ProtectReferencedItem != nil {
		bodies["protectReferencedItem"] = *request.ProtectReferencedItem
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateInventoryModelMaster(
	request *UpdateInventoryModelMasterRequest,
) (*UpdateInventoryModelMasterResult, error) {
	callback := make(chan UpdateInventoryModelMasterAsyncResult, 1)
	go p.UpdateInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteInventoryModelMasterAsync(
	request *DeleteInventoryModelMasterRequest,
	callback chan<- DeleteInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteInventoryModelMaster(
	request *DeleteInventoryModelMasterRequest,
) (*DeleteInventoryModelMasterResult, error) {
	callback := make(chan DeleteInventoryModelMasterAsyncResult, 1)
	go p.DeleteInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoryModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoryModelsAsync(
	request *DescribeInventoryModelsRequest,
	callback chan<- DescribeInventoryModelsAsyncResult,
) {
	path := "/{namespaceName}/inventory"
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

	go describeInventoryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoryModels(
	request *DescribeInventoryModelsRequest,
) (*DescribeInventoryModelsResult, error) {
	callback := make(chan DescribeInventoryModelsAsyncResult, 1)
	go p.DescribeInventoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryModelResult
	if asyncResult.Err != nil {
		callback <- GetInventoryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInventoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryModelAsync(
	request *GetInventoryModelRequest,
	callback chan<- GetInventoryModelAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getInventoryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryModel(
	request *GetInventoryModelRequest,
) (*GetInventoryModelResult, error) {
	callback := make(chan GetInventoryModelAsyncResult, 1)
	go p.GetInventoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeItemModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeItemModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeItemModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemModelMastersAsync(
	request *DescribeItemModelMastersRequest,
	callback chan<- DescribeItemModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeItemModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemModelMasters(
	request *DescribeItemModelMastersRequest,
) (*DescribeItemModelMastersResult, error) {
	callback := make(chan DescribeItemModelMastersAsyncResult, 1)
	go p.DescribeItemModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateItemModelMasterAsync(
	request *CreateItemModelMasterRequest,
	callback chan<- CreateItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
	if request.StackingLimit != nil {
		bodies["stackingLimit"] = *request.StackingLimit
	}
	if request.AllowMultipleStacks != nil {
		bodies["allowMultipleStacks"] = *request.AllowMultipleStacks
	}
	if request.SortValue != nil {
		bodies["sortValue"] = *request.SortValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateItemModelMaster(
	request *CreateItemModelMasterRequest,
) (*CreateItemModelMasterResult, error) {
	callback := make(chan CreateItemModelMasterAsyncResult, 1)
	go p.CreateItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemModelMasterAsync(
	request *GetItemModelMasterRequest,
	callback chan<- GetItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemModelMaster(
	request *GetItemModelMasterRequest,
) (*GetItemModelMasterResult, error) {
	callback := make(chan GetItemModelMasterAsyncResult, 1)
	go p.GetItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateItemModelMasterAsync(
	request *UpdateItemModelMasterRequest,
	callback chan<- UpdateItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.StackingLimit != nil {
		bodies["stackingLimit"] = *request.StackingLimit
	}
	if request.AllowMultipleStacks != nil {
		bodies["allowMultipleStacks"] = *request.AllowMultipleStacks
	}
	if request.SortValue != nil {
		bodies["sortValue"] = *request.SortValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateItemModelMaster(
	request *UpdateItemModelMasterRequest,
) (*UpdateItemModelMasterResult, error) {
	callback := make(chan UpdateItemModelMasterAsyncResult, 1)
	go p.UpdateItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteItemModelMasterAsync(
	request *DeleteItemModelMasterRequest,
	callback chan<- DeleteItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteItemModelMaster(
	request *DeleteItemModelMasterRequest,
) (*DeleteItemModelMasterResult, error) {
	callback := make(chan DeleteItemModelMasterAsyncResult, 1)
	go p.DeleteItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeItemModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeItemModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeItemModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemModelsAsync(
	request *DescribeItemModelsRequest,
	callback chan<- DescribeItemModelsAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeItemModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemModels(
	request *DescribeItemModelsRequest,
) (*DescribeItemModelsResult, error) {
	callback := make(chan DescribeItemModelsAsyncResult, 1)
	go p.DescribeItemModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemModelResult
	if asyncResult.Err != nil {
		callback <- GetItemModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemModelAsync(
	request *GetItemModelRequest,
	callback chan<- GetItemModelAsyncResult,
) {
	path := "/{namespaceName}/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getItemModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemModel(
	request *GetItemModelRequest,
) (*GetItemModelResult, error) {
	callback := make(chan GetItemModelAsyncResult, 1)
	go p.GetItemModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleInventoryModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleInventoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleInventoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleInventoryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleInventoryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleInventoryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleInventoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleInventoryModelMastersAsync(
	request *DescribeSimpleInventoryModelMastersRequest,
	callback chan<- DescribeSimpleInventoryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory"
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

	go describeSimpleInventoryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleInventoryModelMasters(
	request *DescribeSimpleInventoryModelMastersRequest,
) (*DescribeSimpleInventoryModelMastersResult, error) {
	callback := make(chan DescribeSimpleInventoryModelMastersAsyncResult, 1)
	go p.DescribeSimpleInventoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSimpleInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateSimpleInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSimpleInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSimpleInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateSimpleInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateSimpleInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateSimpleInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateSimpleInventoryModelMasterAsync(
	request *CreateSimpleInventoryModelMasterRequest,
	callback chan<- CreateSimpleInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createSimpleInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateSimpleInventoryModelMaster(
	request *CreateSimpleInventoryModelMasterRequest,
) (*CreateSimpleInventoryModelMasterResult, error) {
	callback := make(chan CreateSimpleInventoryModelMasterAsyncResult, 1)
	go p.CreateSimpleInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetSimpleInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleInventoryModelMasterAsync(
	request *GetSimpleInventoryModelMasterRequest,
	callback chan<- GetSimpleInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleInventoryModelMaster(
	request *GetSimpleInventoryModelMasterRequest,
) (*GetSimpleInventoryModelMasterResult, error) {
	callback := make(chan GetSimpleInventoryModelMasterAsyncResult, 1)
	go p.GetSimpleInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSimpleInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSimpleInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSimpleInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSimpleInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateSimpleInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateSimpleInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateSimpleInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateSimpleInventoryModelMasterAsync(
	request *UpdateSimpleInventoryModelMasterRequest,
	callback chan<- UpdateSimpleInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateSimpleInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateSimpleInventoryModelMaster(
	request *UpdateSimpleInventoryModelMasterRequest,
) (*UpdateSimpleInventoryModelMasterResult, error) {
	callback := make(chan UpdateSimpleInventoryModelMasterAsyncResult, 1)
	go p.UpdateSimpleInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSimpleInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSimpleInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSimpleInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSimpleInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteSimpleInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSimpleInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSimpleInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteSimpleInventoryModelMasterAsync(
	request *DeleteSimpleInventoryModelMasterRequest,
	callback chan<- DeleteSimpleInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteSimpleInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteSimpleInventoryModelMaster(
	request *DeleteSimpleInventoryModelMasterRequest,
) (*DeleteSimpleInventoryModelMasterResult, error) {
	callback := make(chan DeleteSimpleInventoryModelMasterAsyncResult, 1)
	go p.DeleteSimpleInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleInventoryModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleInventoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleInventoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleInventoryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleInventoryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleInventoryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleInventoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleInventoryModelsAsync(
	request *DescribeSimpleInventoryModelsRequest,
	callback chan<- DescribeSimpleInventoryModelsAsyncResult,
) {
	path := "/{namespaceName}/simple/inventory"
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

	go describeSimpleInventoryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleInventoryModels(
	request *DescribeSimpleInventoryModelsRequest,
) (*DescribeSimpleInventoryModelsResult, error) {
	callback := make(chan DescribeSimpleInventoryModelsAsyncResult, 1)
	go p.DescribeSimpleInventoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleInventoryModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleInventoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleInventoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleInventoryModelResult
	if asyncResult.Err != nil {
		callback <- GetSimpleInventoryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleInventoryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleInventoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleInventoryModelAsync(
	request *GetSimpleInventoryModelRequest,
	callback chan<- GetSimpleInventoryModelAsyncResult,
) {
	path := "/{namespaceName}/simple/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleInventoryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleInventoryModel(
	request *GetSimpleInventoryModelRequest,
) (*GetSimpleInventoryModelResult, error) {
	callback := make(chan GetSimpleInventoryModelAsyncResult, 1)
	go p.GetSimpleInventoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleItemModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleItemModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleItemModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleItemModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleItemModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleItemModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleItemModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleItemModelMastersAsync(
	request *DescribeSimpleItemModelMastersRequest,
	callback chan<- DescribeSimpleItemModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeSimpleItemModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleItemModelMasters(
	request *DescribeSimpleItemModelMastersRequest,
) (*DescribeSimpleItemModelMastersResult, error) {
	callback := make(chan DescribeSimpleItemModelMastersAsyncResult, 1)
	go p.DescribeSimpleItemModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSimpleItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateSimpleItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSimpleItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSimpleItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateSimpleItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateSimpleItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateSimpleItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateSimpleItemModelMasterAsync(
	request *CreateSimpleItemModelMasterRequest,
	callback chan<- CreateSimpleItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createSimpleItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateSimpleItemModelMaster(
	request *CreateSimpleItemModelMasterRequest,
) (*CreateSimpleItemModelMasterResult, error) {
	callback := make(chan CreateSimpleItemModelMasterAsyncResult, 1)
	go p.CreateSimpleItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemModelMasterAsync(
	request *GetSimpleItemModelMasterRequest,
	callback chan<- GetSimpleItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItemModelMaster(
	request *GetSimpleItemModelMasterRequest,
) (*GetSimpleItemModelMasterResult, error) {
	callback := make(chan GetSimpleItemModelMasterAsyncResult, 1)
	go p.GetSimpleItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSimpleItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSimpleItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSimpleItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSimpleItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateSimpleItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateSimpleItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateSimpleItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateSimpleItemModelMasterAsync(
	request *UpdateSimpleItemModelMasterRequest,
	callback chan<- UpdateSimpleItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateSimpleItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateSimpleItemModelMaster(
	request *UpdateSimpleItemModelMasterRequest,
) (*UpdateSimpleItemModelMasterResult, error) {
	callback := make(chan UpdateSimpleItemModelMasterAsyncResult, 1)
	go p.UpdateSimpleItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSimpleItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSimpleItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSimpleItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSimpleItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteSimpleItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSimpleItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSimpleItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteSimpleItemModelMasterAsync(
	request *DeleteSimpleItemModelMasterRequest,
	callback chan<- DeleteSimpleItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteSimpleItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteSimpleItemModelMaster(
	request *DeleteSimpleItemModelMasterRequest,
) (*DeleteSimpleItemModelMasterResult, error) {
	callback := make(chan DeleteSimpleItemModelMasterAsyncResult, 1)
	go p.DeleteSimpleItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleItemModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleItemModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleItemModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleItemModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleItemModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleItemModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleItemModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleItemModelsAsync(
	request *DescribeSimpleItemModelsRequest,
	callback chan<- DescribeSimpleItemModelsAsyncResult,
) {
	path := "/{namespaceName}/simple/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeSimpleItemModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleItemModels(
	request *DescribeSimpleItemModelsRequest,
) (*DescribeSimpleItemModelsResult, error) {
	callback := make(chan DescribeSimpleItemModelsAsyncResult, 1)
	go p.DescribeSimpleItemModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemModelResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemModelAsync(
	request *GetSimpleItemModelRequest,
	callback chan<- GetSimpleItemModelAsyncResult,
) {
	path := "/{namespaceName}/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleItemModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItemModel(
	request *GetSimpleItemModelRequest,
) (*GetSimpleItemModelResult, error) {
	callback := make(chan GetSimpleItemModelAsyncResult, 1)
	go p.GetSimpleItemModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigInventoryModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigInventoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigInventoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigInventoryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeBigInventoryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigInventoryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigInventoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigInventoryModelMastersAsync(
	request *DescribeBigInventoryModelMastersRequest,
	callback chan<- DescribeBigInventoryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory"
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

	go describeBigInventoryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigInventoryModelMasters(
	request *DescribeBigInventoryModelMastersRequest,
) (*DescribeBigInventoryModelMastersResult, error) {
	callback := make(chan DescribeBigInventoryModelMastersAsyncResult, 1)
	go p.DescribeBigInventoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createBigInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateBigInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBigInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBigInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateBigInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateBigInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateBigInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateBigInventoryModelMasterAsync(
	request *CreateBigInventoryModelMasterRequest,
	callback chan<- CreateBigInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createBigInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateBigInventoryModelMaster(
	request *CreateBigInventoryModelMasterRequest,
) (*CreateBigInventoryModelMasterResult, error) {
	callback := make(chan CreateBigInventoryModelMasterAsyncResult, 1)
	go p.CreateBigInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetBigInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigInventoryModelMasterAsync(
	request *GetBigInventoryModelMasterRequest,
	callback chan<- GetBigInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBigInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigInventoryModelMaster(
	request *GetBigInventoryModelMasterRequest,
) (*GetBigInventoryModelMasterResult, error) {
	callback := make(chan GetBigInventoryModelMasterAsyncResult, 1)
	go p.GetBigInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBigInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBigInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBigInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBigInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateBigInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBigInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBigInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateBigInventoryModelMasterAsync(
	request *UpdateBigInventoryModelMasterRequest,
	callback chan<- UpdateBigInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateBigInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateBigInventoryModelMaster(
	request *UpdateBigInventoryModelMasterRequest,
) (*UpdateBigInventoryModelMasterResult, error) {
	callback := make(chan UpdateBigInventoryModelMasterAsyncResult, 1)
	go p.UpdateBigInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBigInventoryModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBigInventoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBigInventoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBigInventoryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteBigInventoryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBigInventoryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBigInventoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteBigInventoryModelMasterAsync(
	request *DeleteBigInventoryModelMasterRequest,
	callback chan<- DeleteBigInventoryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteBigInventoryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteBigInventoryModelMaster(
	request *DeleteBigInventoryModelMasterRequest,
) (*DeleteBigInventoryModelMasterResult, error) {
	callback := make(chan DeleteBigInventoryModelMasterAsyncResult, 1)
	go p.DeleteBigInventoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigInventoryModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigInventoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigInventoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigInventoryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeBigInventoryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigInventoryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigInventoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigInventoryModelsAsync(
	request *DescribeBigInventoryModelsRequest,
	callback chan<- DescribeBigInventoryModelsAsyncResult,
) {
	path := "/{namespaceName}/big/inventory"
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

	go describeBigInventoryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigInventoryModels(
	request *DescribeBigInventoryModelsRequest,
) (*DescribeBigInventoryModelsResult, error) {
	callback := make(chan DescribeBigInventoryModelsAsyncResult, 1)
	go p.DescribeBigInventoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigInventoryModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigInventoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigInventoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigInventoryModelResult
	if asyncResult.Err != nil {
		callback <- GetBigInventoryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigInventoryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigInventoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigInventoryModelAsync(
	request *GetBigInventoryModelRequest,
	callback chan<- GetBigInventoryModelAsyncResult,
) {
	path := "/{namespaceName}/big/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBigInventoryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigInventoryModel(
	request *GetBigInventoryModelRequest,
) (*GetBigInventoryModelResult, error) {
	callback := make(chan GetBigInventoryModelAsyncResult, 1)
	go p.GetBigInventoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigItemModelMastersAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigItemModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigItemModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigItemModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeBigItemModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigItemModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigItemModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigItemModelMastersAsync(
	request *DescribeBigItemModelMastersRequest,
	callback chan<- DescribeBigItemModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeBigItemModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigItemModelMasters(
	request *DescribeBigItemModelMastersRequest,
) (*DescribeBigItemModelMastersResult, error) {
	callback := make(chan DescribeBigItemModelMastersAsyncResult, 1)
	go p.DescribeBigItemModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createBigItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateBigItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBigItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBigItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateBigItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateBigItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateBigItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) CreateBigItemModelMasterAsync(
	request *CreateBigItemModelMasterRequest,
	callback chan<- CreateBigItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createBigItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) CreateBigItemModelMaster(
	request *CreateBigItemModelMasterRequest,
) (*CreateBigItemModelMasterResult, error) {
	callback := make(chan CreateBigItemModelMasterAsyncResult, 1)
	go p.CreateBigItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetBigItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigItemModelMasterAsync(
	request *GetBigItemModelMasterRequest,
	callback chan<- GetBigItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBigItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigItemModelMaster(
	request *GetBigItemModelMasterRequest,
) (*GetBigItemModelMasterResult, error) {
	callback := make(chan GetBigItemModelMasterAsyncResult, 1)
	go p.GetBigItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBigItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBigItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBigItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBigItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateBigItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBigItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBigItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateBigItemModelMasterAsync(
	request *UpdateBigItemModelMasterRequest,
	callback chan<- UpdateBigItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateBigItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateBigItemModelMaster(
	request *UpdateBigItemModelMasterRequest,
) (*UpdateBigItemModelMasterResult, error) {
	callback := make(chan UpdateBigItemModelMasterAsyncResult, 1)
	go p.UpdateBigItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBigItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBigItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBigItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBigItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteBigItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBigItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBigItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteBigItemModelMasterAsync(
	request *DeleteBigItemModelMasterRequest,
	callback chan<- DeleteBigItemModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteBigItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteBigItemModelMaster(
	request *DeleteBigItemModelMasterRequest,
) (*DeleteBigItemModelMasterResult, error) {
	callback := make(chan DeleteBigItemModelMasterAsyncResult, 1)
	go p.DeleteBigItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigItemModelsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigItemModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigItemModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigItemModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeBigItemModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigItemModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigItemModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigItemModelsAsync(
	request *DescribeBigItemModelsRequest,
	callback chan<- DescribeBigItemModelsAsyncResult,
) {
	path := "/{namespaceName}/big/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeBigItemModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigItemModels(
	request *DescribeBigItemModelsRequest,
) (*DescribeBigItemModelsResult, error) {
	callback := make(chan DescribeBigItemModelsAsyncResult, 1)
	go p.DescribeBigItemModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigItemModelAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigItemModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigItemModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigItemModelResult
	if asyncResult.Err != nil {
		callback <- GetBigItemModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigItemModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigItemModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigItemModelAsync(
	request *GetBigItemModelRequest,
	callback chan<- GetBigItemModelAsyncResult,
) {
	path := "/{namespaceName}/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBigItemModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigItemModel(
	request *GetBigItemModelRequest,
) (*GetBigItemModelResult, error) {
	callback := make(chan GetBigItemModelAsyncResult, 1)
	go p.GetBigItemModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2InventoryRestClient,
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

func (p Gs2InventoryRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ExportMaster(
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

func getCurrentItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetCurrentItemModelMasterAsync(
	request *GetCurrentItemModelMasterRequest,
	callback chan<- GetCurrentItemModelMasterAsyncResult,
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

	go getCurrentItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetCurrentItemModelMaster(
	request *GetCurrentItemModelMasterRequest,
) (*GetCurrentItemModelMasterResult, error) {
	callback := make(chan GetCurrentItemModelMasterAsyncResult, 1)
	go p.GetCurrentItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentItemModelMasterAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentItemModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentItemModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentItemModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentItemModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentItemModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterAsync(
	request *UpdateCurrentItemModelMasterRequest,
	callback chan<- UpdateCurrentItemModelMasterAsyncResult,
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

	go updateCurrentItemModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMaster(
	request *UpdateCurrentItemModelMasterRequest,
) (*UpdateCurrentItemModelMasterResult, error) {
	callback := make(chan UpdateCurrentItemModelMasterAsyncResult, 1)
	go p.UpdateCurrentItemModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentItemModelMasterFromGitHubAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentItemModelMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentItemModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterFromGitHubAsync(
	request *UpdateCurrentItemModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentItemModelMasterFromGitHubAsyncResult,
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

	go updateCurrentItemModelMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) UpdateCurrentItemModelMasterFromGitHub(
	request *UpdateCurrentItemModelMasterFromGitHubRequest,
) (*UpdateCurrentItemModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentItemModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentItemModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoriesAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoriesResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoriesAsync(
	request *DescribeInventoriesRequest,
	callback chan<- DescribeInventoriesAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeInventoriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventories(
	request *DescribeInventoriesRequest,
) (*DescribeInventoriesResult, error) {
	callback := make(chan DescribeInventoriesAsyncResult, 1)
	go p.DescribeInventoriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeInventoriesByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeInventoriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeInventoriesByUserIdAsync(
	request *DescribeInventoriesByUserIdRequest,
	callback chan<- DescribeInventoriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory"
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

	go describeInventoriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeInventoriesByUserId(
	request *DescribeInventoriesByUserIdRequest,
) (*DescribeInventoriesByUserIdResult, error) {
	callback := make(chan DescribeInventoriesByUserIdAsyncResult, 1)
	go p.DescribeInventoriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryResult
	if asyncResult.Err != nil {
		callback <- GetInventoryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInventoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryAsync(
	request *GetInventoryRequest,
	callback chan<- GetInventoryAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go getInventoryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventory(
	request *GetInventoryRequest,
) (*GetInventoryResult, error) {
	callback := make(chan GetInventoryAsyncResult, 1)
	go p.GetInventoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getInventoryByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetInventoryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetInventoryByUserIdAsync(
	request *GetInventoryByUserIdRequest,
	callback chan<- GetInventoryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
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

	go getInventoryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetInventoryByUserId(
	request *GetInventoryByUserIdRequest,
) (*GetInventoryByUserIdResult, error) {
	callback := make(chan GetInventoryByUserIdAsyncResult, 1)
	go p.GetInventoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addCapacityByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddCapacityByUserIdAsync(
	request *AddCapacityByUserIdRequest,
	callback chan<- AddCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/capacity"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AddCapacityValue != nil {
		bodies["addCapacityValue"] = *request.AddCapacityValue
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

	go addCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddCapacityByUserId(
	request *AddCapacityByUserIdRequest,
) (*AddCapacityByUserIdResult, error) {
	callback := make(chan AddCapacityByUserIdAsyncResult, 1)
	go p.AddCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setCapacityByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) SetCapacityByUserIdAsync(
	request *SetCapacityByUserIdRequest,
	callback chan<- SetCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/capacity"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.NewCapacityValue != nil {
		bodies["newCapacityValue"] = *request.NewCapacityValue
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

	go setCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) SetCapacityByUserId(
	request *SetCapacityByUserIdRequest,
) (*SetCapacityByUserIdResult, error) {
	callback := make(chan SetCapacityByUserIdAsyncResult, 1)
	go p.SetCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteInventoryByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteInventoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteInventoryByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteInventoryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteInventoryByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteInventoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteInventoryByUserIdAsync(
	request *DeleteInventoryByUserIdRequest,
	callback chan<- DeleteInventoryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
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

	go deleteInventoryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteInventoryByUserId(
	request *DeleteInventoryByUserIdRequest,
) (*DeleteInventoryByUserIdResult, error) {
	callback := make(chan DeleteInventoryByUserIdAsyncResult, 1)
	go p.DeleteInventoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyInventoryCurrentMaxCapacityAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyInventoryCurrentMaxCapacityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyInventoryCurrentMaxCapacityResult
	if asyncResult.Err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyInventoryCurrentMaxCapacityAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyInventoryCurrentMaxCapacityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacityAsync(
	request *VerifyInventoryCurrentMaxCapacityRequest,
	callback chan<- VerifyInventoryCurrentMaxCapacityAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CurrentInventoryMaxCapacity != nil {
		bodies["currentInventoryMaxCapacity"] = *request.CurrentInventoryMaxCapacity
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

	go verifyInventoryCurrentMaxCapacityAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacity(
	request *VerifyInventoryCurrentMaxCapacityRequest,
) (*VerifyInventoryCurrentMaxCapacityResult, error) {
	callback := make(chan VerifyInventoryCurrentMaxCapacityAsyncResult, 1)
	go p.VerifyInventoryCurrentMaxCapacityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyInventoryCurrentMaxCapacityByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyInventoryCurrentMaxCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacityByUserIdAsync(
	request *VerifyInventoryCurrentMaxCapacityByUserIdRequest,
	callback chan<- VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/verify/{verifyType}"
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
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CurrentInventoryMaxCapacity != nil {
		bodies["currentInventoryMaxCapacity"] = *request.CurrentInventoryMaxCapacity
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

	go verifyInventoryCurrentMaxCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacityByUserId(
	request *VerifyInventoryCurrentMaxCapacityByUserIdRequest,
) (*VerifyInventoryCurrentMaxCapacityByUserIdResult, error) {
	callback := make(chan VerifyInventoryCurrentMaxCapacityByUserIdAsyncResult, 1)
	go p.VerifyInventoryCurrentMaxCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyInventoryCurrentMaxCapacityByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyInventoryCurrentMaxCapacityByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacityByStampTaskAsync(
	request *VerifyInventoryCurrentMaxCapacityByStampTaskRequest,
	callback chan<- VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult,
) {
	path := "/stamp/inventory/verify"

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

	go verifyInventoryCurrentMaxCapacityByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyInventoryCurrentMaxCapacityByStampTask(
	request *VerifyInventoryCurrentMaxCapacityByStampTaskRequest,
) (*VerifyInventoryCurrentMaxCapacityByStampTaskResult, error) {
	callback := make(chan VerifyInventoryCurrentMaxCapacityByStampTaskAsyncResult, 1)
	go p.VerifyInventoryCurrentMaxCapacityByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addCapacityByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddCapacityByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddCapacityByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddCapacityByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddCapacityByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddCapacityByStampSheetAsync(
	request *AddCapacityByStampSheetRequest,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/inventory/capacity/add"

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

	go addCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddCapacityByStampSheet(
	request *AddCapacityByStampSheetRequest,
) (*AddCapacityByStampSheetResult, error) {
	callback := make(chan AddCapacityByStampSheetAsyncResult, 1)
	go p.AddCapacityByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setCapacityByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetCapacityByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetCapacityByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetCapacityByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetCapacityByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetCapacityByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) SetCapacityByStampSheetAsync(
	request *SetCapacityByStampSheetRequest,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/inventory/capacity/set"

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

	go setCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) SetCapacityByStampSheet(
	request *SetCapacityByStampSheetRequest,
) (*SetCapacityByStampSheetResult, error) {
	callback := make(chan SetCapacityByStampSheetAsyncResult, 1)
	go p.SetCapacityByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemSetsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemSetsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemSetsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemSetsResult
	if asyncResult.Err != nil {
		callback <- DescribeItemSetsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeItemSetsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeItemSetsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemSetsAsync(
	request *DescribeItemSetsRequest,
	callback chan<- DescribeItemSetsAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeItemSetsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemSets(
	request *DescribeItemSetsRequest,
) (*DescribeItemSetsResult, error) {
	callback := make(chan DescribeItemSetsAsyncResult, 1)
	go p.DescribeItemSetsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeItemSetsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeItemSetsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeItemSetsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeItemSetsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeItemSetsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeItemSetsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeItemSetsByUserIdAsync(
	request *DescribeItemSetsByUserIdRequest,
	callback chan<- DescribeItemSetsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go describeItemSetsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeItemSetsByUserId(
	request *DescribeItemSetsByUserIdRequest,
) (*DescribeItemSetsByUserIdResult, error) {
	callback := make(chan DescribeItemSetsByUserIdAsyncResult, 1)
	go p.DescribeItemSetsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemSetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemSetResult
	if asyncResult.Err != nil {
		callback <- GetItemSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemSetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemSetAsync(
	request *GetItemSetRequest,
	callback chan<- GetItemSetAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getItemSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemSet(
	request *GetItemSetRequest,
) (*GetItemSetResult, error) {
	callback := make(chan GetItemSetAsyncResult, 1)
	go p.GetItemSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemSetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemSetByUserIdAsync(
	request *GetItemSetByUserIdRequest,
	callback chan<- GetItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemSetByUserId(
	request *GetItemSetByUserIdRequest,
) (*GetItemSetByUserIdResult, error) {
	callback := make(chan GetItemSetByUserIdAsyncResult, 1)
	go p.GetItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemWithSignatureAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetItemWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemWithSignatureAsync(
	request *GetItemWithSignatureRequest,
	callback chan<- GetItemWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getItemWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemWithSignature(
	request *GetItemWithSignatureRequest,
) (*GetItemWithSignatureResult, error) {
	callback := make(chan GetItemWithSignatureAsyncResult, 1)
	go p.GetItemWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getItemWithSignatureByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetItemWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetItemWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetItemWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetItemWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetItemWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetItemWithSignatureByUserIdAsync(
	request *GetItemWithSignatureByUserIdRequest,
	callback chan<- GetItemWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getItemWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetItemWithSignatureByUserId(
	request *GetItemWithSignatureByUserIdRequest,
) (*GetItemWithSignatureByUserIdResult, error) {
	callback := make(chan GetItemWithSignatureByUserIdAsyncResult, 1)
	go p.GetItemWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireItemSetByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- AcquireItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireItemSetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireItemSetByUserIdAsync(
	request *AcquireItemSetByUserIdRequest,
	callback chan<- AcquireItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/acquire"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireCount != nil {
		bodies["acquireCount"] = *request.AcquireCount
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.CreateNewItemSet != nil {
		bodies["createNewItemSet"] = *request.CreateNewItemSet
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		bodies["itemSetName"] = *request.ItemSetName
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

	go acquireItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireItemSetByUserId(
	request *AcquireItemSetByUserIdRequest,
) (*AcquireItemSetByUserIdResult, error) {
	callback := make(chan AcquireItemSetByUserIdAsyncResult, 1)
	go p.AcquireItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeItemSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeItemSetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetAsync(
	request *ConsumeItemSetRequest,
	callback chan<- ConsumeItemSetAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCount != nil {
		bodies["consumeCount"] = *request.ConsumeCount
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		bodies["itemSetName"] = *request.ItemSetName
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

	go consumeItemSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSet(
	request *ConsumeItemSetRequest,
) (*ConsumeItemSetResult, error) {
	callback := make(chan ConsumeItemSetAsyncResult, 1)
	go p.ConsumeItemSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeItemSetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetByUserIdAsync(
	request *ConsumeItemSetByUserIdRequest,
	callback chan<- ConsumeItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCount != nil {
		bodies["consumeCount"] = *request.ConsumeCount
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		bodies["itemSetName"] = *request.ItemSetName
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

	go consumeItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSetByUserId(
	request *ConsumeItemSetByUserIdRequest,
) (*ConsumeItemSetByUserIdResult, error) {
	callback := make(chan ConsumeItemSetByUserIdAsyncResult, 1)
	go p.ConsumeItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteItemSetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteItemSetByUserIdAsync(
	request *DeleteItemSetByUserIdRequest,
	callback chan<- DeleteItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ItemSetName != nil {
		queryStrings["itemSetName"] = core.ToString(*request.ItemSetName)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteItemSetByUserId(
	request *DeleteItemSetByUserIdRequest,
) (*DeleteItemSetByUserIdResult, error) {
	callback := make(chan DeleteItemSetByUserIdAsyncResult, 1)
	go p.DeleteItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyItemSetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyItemSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyItemSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyItemSetResult
	if asyncResult.Err != nil {
		callback <- VerifyItemSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyItemSetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyItemSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyItemSetAsync(
	request *VerifyItemSetRequest,
	callback chan<- VerifyItemSetAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		bodies["itemSetName"] = *request.ItemSetName
	}
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyItemSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyItemSet(
	request *VerifyItemSetRequest,
) (*VerifyItemSetResult, error) {
	callback := make(chan VerifyItemSetAsyncResult, 1)
	go p.VerifyItemSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyItemSetByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyItemSetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyItemSetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyItemSetByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyItemSetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyItemSetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyItemSetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyItemSetByUserIdAsync(
	request *VerifyItemSetByUserIdRequest,
	callback chan<- VerifyItemSetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
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
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		bodies["itemSetName"] = *request.ItemSetName
	}
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

	go verifyItemSetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyItemSetByUserId(
	request *VerifyItemSetByUserIdRequest,
) (*VerifyItemSetByUserIdResult, error) {
	callback := make(chan VerifyItemSetByUserIdAsyncResult, 1)
	go p.VerifyItemSetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireItemSetByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireItemSetByStampSheetAsync(
	request *AcquireItemSetByStampSheetRequest,
	callback chan<- AcquireItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/acquire"

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

	go acquireItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireItemSetByStampSheet(
	request *AcquireItemSetByStampSheetRequest,
) (*AcquireItemSetByStampSheetResult, error) {
	callback := make(chan AcquireItemSetByStampSheetAsyncResult, 1)
	go p.AcquireItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeItemSetByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeItemSetByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeItemSetByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ConsumeItemSetByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeItemSetByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeItemSetByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeItemSetByStampTaskAsync(
	request *ConsumeItemSetByStampTaskRequest,
	callback chan<- ConsumeItemSetByStampTaskAsyncResult,
) {
	path := "/stamp/item/consume"

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

	go consumeItemSetByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeItemSetByStampTask(
	request *ConsumeItemSetByStampTaskRequest,
) (*ConsumeItemSetByStampTaskResult, error) {
	callback := make(chan ConsumeItemSetByStampTaskAsyncResult, 1)
	go p.ConsumeItemSetByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyItemSetByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyItemSetByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyItemSetByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyItemSetByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyItemSetByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyItemSetByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyItemSetByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyItemSetByStampTaskAsync(
	request *VerifyItemSetByStampTaskRequest,
	callback chan<- VerifyItemSetByStampTaskAsyncResult,
) {
	path := "/stamp/item/verify"

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

	go verifyItemSetByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyItemSetByStampTask(
	request *VerifyItemSetByStampTaskRequest,
) (*VerifyItemSetByStampTaskResult, error) {
	callback := make(chan VerifyItemSetByStampTaskAsyncResult, 1)
	go p.VerifyItemSetByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReferenceOfResult
	if asyncResult.Err != nil {
		callback <- DescribeReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReferenceOfAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeReferenceOfAsync(
	request *DescribeReferenceOfRequest,
	callback chan<- DescribeReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
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

	go describeReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeReferenceOf(
	request *DescribeReferenceOfRequest,
) (*DescribeReferenceOfResult, error) {
	callback := make(chan DescribeReferenceOfAsyncResult, 1)
	go p.DescribeReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReferenceOfByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeReferenceOfByUserIdAsync(
	request *DescribeReferenceOfByUserIdRequest,
	callback chan<- DescribeReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeReferenceOfByUserId(
	request *DescribeReferenceOfByUserIdRequest,
) (*DescribeReferenceOfByUserIdResult, error) {
	callback := make(chan DescribeReferenceOfByUserIdAsyncResult, 1)
	go p.DescribeReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReferenceOfResult
	if asyncResult.Err != nil {
		callback <- GetReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReferenceOfAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetReferenceOfAsync(
	request *GetReferenceOfRequest,
	callback chan<- GetReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
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

	go getReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetReferenceOf(
	request *GetReferenceOfRequest,
) (*GetReferenceOfResult, error) {
	callback := make(chan GetReferenceOfAsyncResult, 1)
	go p.GetReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReferenceOfByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetReferenceOfByUserIdAsync(
	request *GetReferenceOfByUserIdRequest,
	callback chan<- GetReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetReferenceOfByUserId(
	request *GetReferenceOfByUserIdRequest,
) (*GetReferenceOfByUserIdResult, error) {
	callback := make(chan GetReferenceOfByUserIdAsyncResult, 1)
	go p.GetReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReferenceOfAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfAsync(
	request *VerifyReferenceOfRequest,
	callback chan<- VerifyReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOf(
	request *VerifyReferenceOfRequest,
) (*VerifyReferenceOfResult, error) {
	callback := make(chan VerifyReferenceOfAsyncResult, 1)
	go p.VerifyReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReferenceOfByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfByUserIdAsync(
	request *VerifyReferenceOfByUserIdRequest,
	callback chan<- VerifyReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
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

	go verifyReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOfByUserId(
	request *VerifyReferenceOfByUserIdRequest,
) (*VerifyReferenceOfByUserIdResult, error) {
	callback := make(chan VerifyReferenceOfByUserIdAsyncResult, 1)
	go p.VerifyReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddReferenceOfAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfAsync(
	request *AddReferenceOfRequest,
	callback chan<- AddReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		bodies["referenceOf"] = *request.ReferenceOf
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

	go addReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOf(
	request *AddReferenceOfRequest,
) (*AddReferenceOfResult, error) {
	callback := make(chan AddReferenceOfAsyncResult, 1)
	go p.AddReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddReferenceOfByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfByUserIdAsync(
	request *AddReferenceOfByUserIdRequest,
	callback chan<- AddReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		bodies["referenceOf"] = *request.ReferenceOf
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

	go addReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOfByUserId(
	request *AddReferenceOfByUserIdRequest,
) (*AddReferenceOfByUserIdResult, error) {
	callback := make(chan AddReferenceOfByUserIdAsyncResult, 1)
	go p.AddReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReferenceOfAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReferenceOfAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfAsync(
	request *DeleteReferenceOfRequest,
	callback chan<- DeleteReferenceOfAsyncResult,
) {
	path := "/{namespaceName}/user/me/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go deleteReferenceOfAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOf(
	request *DeleteReferenceOfRequest,
) (*DeleteReferenceOfResult, error) {
	callback := make(chan DeleteReferenceOfAsyncResult, 1)
	go p.DeleteReferenceOfAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReferenceOfByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReferenceOfByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfByUserIdAsync(
	request *DeleteReferenceOfByUserIdRequest,
	callback chan<- DeleteReferenceOfByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/inventory/{inventoryName}/item/{itemName}/{itemSetName}/reference/{referenceOf}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.ItemSetName != nil && *request.ItemSetName != "" {
		path = strings.ReplaceAll(path, "{itemSetName}", core.ToString(*request.ItemSetName))
	} else {
		path = strings.ReplaceAll(path, "{itemSetName}", "null")
	}
	if request.ReferenceOf != nil && *request.ReferenceOf != "" {
		path = strings.ReplaceAll(path, "{referenceOf}", core.ToString(*request.ReferenceOf))
	} else {
		path = strings.ReplaceAll(path, "{referenceOf}", "null")
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

	go deleteReferenceOfByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOfByUserId(
	request *DeleteReferenceOfByUserIdRequest,
) (*DeleteReferenceOfByUserIdResult, error) {
	callback := make(chan DeleteReferenceOfByUserIdAsyncResult, 1)
	go p.DeleteReferenceOfByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addReferenceOfItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddReferenceOfItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AddReferenceOfItemSetByStampSheetAsync(
	request *AddReferenceOfItemSetByStampSheetRequest,
	callback chan<- AddReferenceOfItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/reference/add"

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

	go addReferenceOfItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AddReferenceOfItemSetByStampSheet(
	request *AddReferenceOfItemSetByStampSheetRequest,
) (*AddReferenceOfItemSetByStampSheetResult, error) {
	callback := make(chan AddReferenceOfItemSetByStampSheetAsyncResult, 1)
	go p.AddReferenceOfItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReferenceOfItemSetByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReferenceOfItemSetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReferenceOfItemSetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteReferenceOfItemSetByStampSheetAsync(
	request *DeleteReferenceOfItemSetByStampSheetRequest,
	callback chan<- DeleteReferenceOfItemSetByStampSheetAsyncResult,
) {
	path := "/stamp/item/reference/delete"

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

	go deleteReferenceOfItemSetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteReferenceOfItemSetByStampSheet(
	request *DeleteReferenceOfItemSetByStampSheetRequest,
) (*DeleteReferenceOfItemSetByStampSheetResult, error) {
	callback := make(chan DeleteReferenceOfItemSetByStampSheetAsyncResult, 1)
	go p.DeleteReferenceOfItemSetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReferenceOfByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReferenceOfByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReferenceOfByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyReferenceOfByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReferenceOfByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReferenceOfByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyReferenceOfByStampTaskAsync(
	request *VerifyReferenceOfByStampTaskRequest,
	callback chan<- VerifyReferenceOfByStampTaskAsyncResult,
) {
	path := "/stamp/item/verify"

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

	go verifyReferenceOfByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyReferenceOfByStampTask(
	request *VerifyReferenceOfByStampTaskRequest,
) (*VerifyReferenceOfByStampTaskResult, error) {
	callback := make(chan VerifyReferenceOfByStampTaskAsyncResult, 1)
	go p.VerifyReferenceOfByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleItemsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleItemsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleItemsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleItemsResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleItemsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleItemsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleItemsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleItemsAsync(
	request *DescribeSimpleItemsRequest,
	callback chan<- DescribeSimpleItemsAsyncResult,
) {
	path := "/{namespaceName}/user/me/simple/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeSimpleItemsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleItems(
	request *DescribeSimpleItemsRequest,
) (*DescribeSimpleItemsResult, error) {
	callback := make(chan DescribeSimpleItemsAsyncResult, 1)
	go p.DescribeSimpleItemsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSimpleItemsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSimpleItemsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSimpleItemsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSimpleItemsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSimpleItemsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSimpleItemsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSimpleItemsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeSimpleItemsByUserIdAsync(
	request *DescribeSimpleItemsByUserIdRequest,
	callback chan<- DescribeSimpleItemsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go describeSimpleItemsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeSimpleItemsByUserId(
	request *DescribeSimpleItemsByUserIdRequest,
) (*DescribeSimpleItemsByUserIdResult, error) {
	callback := make(chan DescribeSimpleItemsByUserIdAsyncResult, 1)
	go p.DescribeSimpleItemsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemAsync(
	request *GetSimpleItemRequest,
	callback chan<- GetSimpleItemAsyncResult,
) {
	path := "/{namespaceName}/user/me/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
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

	go getSimpleItemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItem(
	request *GetSimpleItemRequest,
) (*GetSimpleItemResult, error) {
	callback := make(chan GetSimpleItemAsyncResult, 1)
	go p.GetSimpleItemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemByUserIdAsync(
	request *GetSimpleItemByUserIdRequest,
	callback chan<- GetSimpleItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItemByUserId(
	request *GetSimpleItemByUserIdRequest,
) (*GetSimpleItemByUserIdResult, error) {
	callback := make(chan GetSimpleItemByUserIdAsyncResult, 1)
	go p.GetSimpleItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemWithSignatureAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemWithSignatureAsync(
	request *GetSimpleItemWithSignatureRequest,
	callback chan<- GetSimpleItemWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/simple/inventory/{inventoryName}/item/{itemName}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getSimpleItemWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItemWithSignature(
	request *GetSimpleItemWithSignatureRequest,
) (*GetSimpleItemWithSignatureResult, error) {
	callback := make(chan GetSimpleItemWithSignatureAsyncResult, 1)
	go p.GetSimpleItemWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSimpleItemWithSignatureByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetSimpleItemWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSimpleItemWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSimpleItemWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSimpleItemWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSimpleItemWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSimpleItemWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetSimpleItemWithSignatureByUserIdAsync(
	request *GetSimpleItemWithSignatureByUserIdRequest,
	callback chan<- GetSimpleItemWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/item/{itemName}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSimpleItemWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetSimpleItemWithSignatureByUserId(
	request *GetSimpleItemWithSignatureByUserIdRequest,
) (*GetSimpleItemWithSignatureByUserIdResult, error) {
	callback := make(chan GetSimpleItemWithSignatureByUserIdAsyncResult, 1)
	go p.GetSimpleItemWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireSimpleItemsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireSimpleItemsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireSimpleItemsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireSimpleItemsByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- AcquireSimpleItemsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireSimpleItemsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireSimpleItemsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireSimpleItemsByUserIdAsync(
	request *AcquireSimpleItemsByUserIdRequest,
	callback chan<- AcquireSimpleItemsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/acquire"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireCounts != nil {
		var _acquireCounts []interface{}
		for _, item := range request.AcquireCounts {
			_acquireCounts = append(_acquireCounts, item)
		}
		bodies["acquireCounts"] = _acquireCounts
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

	go acquireSimpleItemsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireSimpleItemsByUserId(
	request *AcquireSimpleItemsByUserIdRequest,
) (*AcquireSimpleItemsByUserIdResult, error) {
	callback := make(chan AcquireSimpleItemsByUserIdAsyncResult, 1)
	go p.AcquireSimpleItemsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeSimpleItemsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeSimpleItemsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeSimpleItemsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeSimpleItemsResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeSimpleItemsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeSimpleItemsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeSimpleItemsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeSimpleItemsAsync(
	request *ConsumeSimpleItemsRequest,
	callback chan<- ConsumeSimpleItemsAsyncResult,
) {
	path := "/{namespaceName}/user/me/simple/inventory/{inventoryName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCounts != nil {
		var _consumeCounts []interface{}
		for _, item := range request.ConsumeCounts {
			_consumeCounts = append(_consumeCounts, item)
		}
		bodies["consumeCounts"] = _consumeCounts
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

	go consumeSimpleItemsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeSimpleItems(
	request *ConsumeSimpleItemsRequest,
) (*ConsumeSimpleItemsResult, error) {
	callback := make(chan ConsumeSimpleItemsAsyncResult, 1)
	go p.ConsumeSimpleItemsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeSimpleItemsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeSimpleItemsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeSimpleItemsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeSimpleItemsByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeSimpleItemsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeSimpleItemsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeSimpleItemsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeSimpleItemsByUserIdAsync(
	request *ConsumeSimpleItemsByUserIdRequest,
	callback chan<- ConsumeSimpleItemsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCounts != nil {
		var _consumeCounts []interface{}
		for _, item := range request.ConsumeCounts {
			_consumeCounts = append(_consumeCounts, item)
		}
		bodies["consumeCounts"] = _consumeCounts
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

	go consumeSimpleItemsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeSimpleItemsByUserId(
	request *ConsumeSimpleItemsByUserIdRequest,
) (*ConsumeSimpleItemsByUserIdResult, error) {
	callback := make(chan ConsumeSimpleItemsByUserIdAsyncResult, 1)
	go p.ConsumeSimpleItemsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSimpleItemsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSimpleItemsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSimpleItemsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSimpleItemsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteSimpleItemsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSimpleItemsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSimpleItemsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteSimpleItemsByUserIdAsync(
	request *DeleteSimpleItemsByUserIdRequest,
	callback chan<- DeleteSimpleItemsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
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

	go deleteSimpleItemsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteSimpleItemsByUserId(
	request *DeleteSimpleItemsByUserIdRequest,
) (*DeleteSimpleItemsByUserIdResult, error) {
	callback := make(chan DeleteSimpleItemsByUserIdAsyncResult, 1)
	go p.DeleteSimpleItemsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySimpleItemAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifySimpleItemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySimpleItemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySimpleItemResult
	if asyncResult.Err != nil {
		callback <- VerifySimpleItemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySimpleItemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySimpleItemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifySimpleItemAsync(
	request *VerifySimpleItemRequest,
	callback chan<- VerifySimpleItemAsyncResult,
) {
	path := "/{namespaceName}/user/me/simple/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifySimpleItemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifySimpleItem(
	request *VerifySimpleItemRequest,
) (*VerifySimpleItemResult, error) {
	callback := make(chan VerifySimpleItemAsyncResult, 1)
	go p.VerifySimpleItemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySimpleItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifySimpleItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySimpleItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySimpleItemByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifySimpleItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySimpleItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySimpleItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifySimpleItemByUserIdAsync(
	request *VerifySimpleItemByUserIdRequest,
	callback chan<- VerifySimpleItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/simple/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
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
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
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

	go verifySimpleItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifySimpleItemByUserId(
	request *VerifySimpleItemByUserIdRequest,
) (*VerifySimpleItemByUserIdResult, error) {
	callback := make(chan VerifySimpleItemByUserIdAsyncResult, 1)
	go p.VerifySimpleItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireSimpleItemsByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireSimpleItemsByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireSimpleItemsByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireSimpleItemsByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireSimpleItemsByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireSimpleItemsByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireSimpleItemsByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireSimpleItemsByStampSheetAsync(
	request *AcquireSimpleItemsByStampSheetRequest,
	callback chan<- AcquireSimpleItemsByStampSheetAsyncResult,
) {
	path := "/stamp/simple/item/acquire"

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

	go acquireSimpleItemsByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireSimpleItemsByStampSheet(
	request *AcquireSimpleItemsByStampSheetRequest,
) (*AcquireSimpleItemsByStampSheetResult, error) {
	callback := make(chan AcquireSimpleItemsByStampSheetAsyncResult, 1)
	go p.AcquireSimpleItemsByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeSimpleItemsByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeSimpleItemsByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeSimpleItemsByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeSimpleItemsByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ConsumeSimpleItemsByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeSimpleItemsByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeSimpleItemsByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeSimpleItemsByStampTaskAsync(
	request *ConsumeSimpleItemsByStampTaskRequest,
	callback chan<- ConsumeSimpleItemsByStampTaskAsyncResult,
) {
	path := "/stamp/simple/item/consume"

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

	go consumeSimpleItemsByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeSimpleItemsByStampTask(
	request *ConsumeSimpleItemsByStampTaskRequest,
) (*ConsumeSimpleItemsByStampTaskResult, error) {
	callback := make(chan ConsumeSimpleItemsByStampTaskAsyncResult, 1)
	go p.ConsumeSimpleItemsByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySimpleItemByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifySimpleItemByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySimpleItemByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySimpleItemByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifySimpleItemByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySimpleItemByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySimpleItemByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifySimpleItemByStampTaskAsync(
	request *VerifySimpleItemByStampTaskRequest,
	callback chan<- VerifySimpleItemByStampTaskAsyncResult,
) {
	path := "/stamp/simple/item/verify"

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

	go verifySimpleItemByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifySimpleItemByStampTask(
	request *VerifySimpleItemByStampTaskRequest,
) (*VerifySimpleItemByStampTaskResult, error) {
	callback := make(chan VerifySimpleItemByStampTaskAsyncResult, 1)
	go p.VerifySimpleItemByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigItemsAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigItemsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigItemsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigItemsResult
	if asyncResult.Err != nil {
		callback <- DescribeBigItemsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigItemsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigItemsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigItemsAsync(
	request *DescribeBigItemsRequest,
	callback chan<- DescribeBigItemsAsyncResult,
) {
	path := "/{namespaceName}/user/me/big/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
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

	go describeBigItemsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigItems(
	request *DescribeBigItemsRequest,
) (*DescribeBigItemsResult, error) {
	callback := make(chan DescribeBigItemsAsyncResult, 1)
	go p.DescribeBigItemsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBigItemsByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBigItemsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBigItemsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBigItemsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeBigItemsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBigItemsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBigItemsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DescribeBigItemsByUserIdAsync(
	request *DescribeBigItemsByUserIdRequest,
	callback chan<- DescribeBigItemsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
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

	go describeBigItemsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DescribeBigItemsByUserId(
	request *DescribeBigItemsByUserIdRequest,
) (*DescribeBigItemsByUserIdResult, error) {
	callback := make(chan DescribeBigItemsByUserIdAsyncResult, 1)
	go p.DescribeBigItemsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigItemAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigItemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigItemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigItemResult
	if asyncResult.Err != nil {
		callback <- GetBigItemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigItemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigItemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigItemAsync(
	request *GetBigItemRequest,
	callback chan<- GetBigItemAsyncResult,
) {
	path := "/{namespaceName}/user/me/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
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

	go getBigItemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigItem(
	request *GetBigItemRequest,
) (*GetBigItemResult, error) {
	callback := make(chan GetBigItemAsyncResult, 1)
	go p.GetBigItemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBigItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBigItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBigItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBigItemByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetBigItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBigItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBigItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) GetBigItemByUserIdAsync(
	request *GetBigItemByUserIdRequest,
	callback chan<- GetBigItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getBigItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) GetBigItemByUserId(
	request *GetBigItemByUserIdRequest,
) (*GetBigItemByUserIdResult, error) {
	callback := make(chan GetBigItemByUserIdAsyncResult, 1)
	go p.GetBigItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireBigItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireBigItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireBigItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireBigItemByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- AcquireBigItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireBigItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireBigItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireBigItemByUserIdAsync(
	request *AcquireBigItemByUserIdRequest,
	callback chan<- AcquireBigItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item/{itemName}/acquire"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireCount != nil && *request.AcquireCount != "" {
		bodies["acquireCount"] = *request.AcquireCount
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

	go acquireBigItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireBigItemByUserId(
	request *AcquireBigItemByUserIdRequest,
) (*AcquireBigItemByUserIdResult, error) {
	callback := make(chan AcquireBigItemByUserIdAsyncResult, 1)
	go p.AcquireBigItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeBigItemAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeBigItemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeBigItemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeBigItemResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeBigItemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeBigItemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeBigItemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeBigItemAsync(
	request *ConsumeBigItemRequest,
	callback chan<- ConsumeBigItemAsyncResult,
) {
	path := "/{namespaceName}/user/me/big/inventory/{inventoryName}/item/{itemName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCount != nil && *request.ConsumeCount != "" {
		bodies["consumeCount"] = *request.ConsumeCount
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

	go consumeBigItemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeBigItem(
	request *ConsumeBigItemRequest,
) (*ConsumeBigItemResult, error) {
	callback := make(chan ConsumeBigItemAsyncResult, 1)
	go p.ConsumeBigItemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeBigItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeBigItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeBigItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeBigItemByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "itemSet.count.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeBigItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeBigItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeBigItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeBigItemByUserIdAsync(
	request *ConsumeBigItemByUserIdRequest,
	callback chan<- ConsumeBigItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item/{itemName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeCount != nil && *request.ConsumeCount != "" {
		bodies["consumeCount"] = *request.ConsumeCount
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

	go consumeBigItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeBigItemByUserId(
	request *ConsumeBigItemByUserIdRequest,
) (*ConsumeBigItemByUserIdResult, error) {
	callback := make(chan ConsumeBigItemByUserIdAsyncResult, 1)
	go p.ConsumeBigItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBigItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBigItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBigItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBigItemByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteBigItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBigItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBigItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) DeleteBigItemByUserIdAsync(
	request *DeleteBigItemByUserIdRequest,
	callback chan<- DeleteBigItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item/{itemName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
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

	go deleteBigItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) DeleteBigItemByUserId(
	request *DeleteBigItemByUserIdRequest,
) (*DeleteBigItemByUserIdResult, error) {
	callback := make(chan DeleteBigItemByUserIdAsyncResult, 1)
	go p.DeleteBigItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyBigItemAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyBigItemAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyBigItemAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyBigItemResult
	if asyncResult.Err != nil {
		callback <- VerifyBigItemAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyBigItemAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyBigItemAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyBigItemAsync(
	request *VerifyBigItemRequest,
	callback chan<- VerifyBigItemAsyncResult,
) {
	path := "/{namespaceName}/user/me/big/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Count != nil && *request.Count != "" {
		bodies["count"] = *request.Count
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

	go verifyBigItemAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyBigItem(
	request *VerifyBigItemRequest,
) (*VerifyBigItemResult, error) {
	callback := make(chan VerifyBigItemAsyncResult, 1)
	go p.VerifyBigItemAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyBigItemByUserIdAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyBigItemByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyBigItemByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyBigItemByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyBigItemByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyBigItemByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyBigItemByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyBigItemByUserIdAsync(
	request *VerifyBigItemByUserIdRequest,
	callback chan<- VerifyBigItemByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/big/inventory/{inventoryName}/item/{itemName}/verify/{verifyType}"
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
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}
	if request.ItemName != nil && *request.ItemName != "" {
		path = strings.ReplaceAll(path, "{itemName}", core.ToString(*request.ItemName))
	} else {
		path = strings.ReplaceAll(path, "{itemName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Count != nil && *request.Count != "" {
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

	go verifyBigItemByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyBigItemByUserId(
	request *VerifyBigItemByUserIdRequest,
) (*VerifyBigItemByUserIdResult, error) {
	callback := make(chan VerifyBigItemByUserIdAsyncResult, 1)
	go p.VerifyBigItemByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireBigItemByStampSheetAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireBigItemByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireBigItemByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireBigItemByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireBigItemByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireBigItemByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireBigItemByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) AcquireBigItemByStampSheetAsync(
	request *AcquireBigItemByStampSheetRequest,
	callback chan<- AcquireBigItemByStampSheetAsyncResult,
) {
	path := "/stamp/big/item/acquire"

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

	go acquireBigItemByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) AcquireBigItemByStampSheet(
	request *AcquireBigItemByStampSheetRequest,
) (*AcquireBigItemByStampSheetResult, error) {
	callback := make(chan AcquireBigItemByStampSheetAsyncResult, 1)
	go p.AcquireBigItemByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeBigItemByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeBigItemByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeBigItemByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeBigItemByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ConsumeBigItemByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeBigItemByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeBigItemByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) ConsumeBigItemByStampTaskAsync(
	request *ConsumeBigItemByStampTaskRequest,
	callback chan<- ConsumeBigItemByStampTaskAsyncResult,
) {
	path := "/stamp/big/item/consume"

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

	go consumeBigItemByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) ConsumeBigItemByStampTask(
	request *ConsumeBigItemByStampTaskRequest,
) (*ConsumeBigItemByStampTaskResult, error) {
	callback := make(chan ConsumeBigItemByStampTaskAsyncResult, 1)
	go p.ConsumeBigItemByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyBigItemByStampTaskAsyncHandler(
	client Gs2InventoryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyBigItemByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyBigItemByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyBigItemByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyBigItemByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyBigItemByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyBigItemByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InventoryRestClient) VerifyBigItemByStampTaskAsync(
	request *VerifyBigItemByStampTaskRequest,
	callback chan<- VerifyBigItemByStampTaskAsyncResult,
) {
	path := "/stamp/big/item/verify"

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

	go verifyBigItemByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inventory").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InventoryRestClient) VerifyBigItemByStampTask(
	request *VerifyBigItemByStampTaskRequest,
) (*VerifyBigItemByStampTaskResult, error) {
	callback := make(chan VerifyBigItemByStampTaskAsyncResult, 1)
	go p.VerifyBigItemByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
