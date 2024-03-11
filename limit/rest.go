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

package limit

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2LimitRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2LimitRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) DescribeNamespacesAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeNamespacesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DescribeNamespaces(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) CreateNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CreateNamespace(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) GetNamespaceStatusAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetNamespaceStatus(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) GetNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetNamespace(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) UpdateNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) UpdateNamespace(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) DeleteNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DeleteNamespace(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) DumpUserDataByUserIdAsync(
	request *DumpUserDataByUserIdRequest,
	callback chan<- DumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DumpUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) CheckDumpUserDataByUserIdAsync(
	request *CheckDumpUserDataByUserIdRequest,
	callback chan<- CheckDumpUserDataByUserIdAsyncResult,
) {
	path := "/system/dump/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CheckDumpUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) CleanUserDataByUserIdAsync(
	request *CleanUserDataByUserIdRequest,
	callback chan<- CleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CleanUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) CheckCleanUserDataByUserIdAsync(
	request *CheckCleanUserDataByUserIdRequest,
	callback chan<- CheckCleanUserDataByUserIdAsyncResult,
) {
	path := "/system/clean/user/{userId}"
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CheckCleanUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) PrepareImportUserDataByUserIdAsync(
	request *PrepareImportUserDataByUserIdRequest,
	callback chan<- PrepareImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/prepare"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) PrepareImportUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) ImportUserDataByUserIdAsync(
	request *ImportUserDataByUserIdRequest,
	callback chan<- ImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) ImportUserDataByUserId(
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
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) CheckImportUserDataByUserIdAsync(
	request *CheckImportUserDataByUserIdRequest,
	callback chan<- CheckImportUserDataByUserIdAsyncResult,
) {
	path := "/system/import/user/{userId}/{uploadToken}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CheckImportUserDataByUserId(
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

func describeCountersAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCountersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCountersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCountersResult
	if asyncResult.Err != nil {
		callback <- DescribeCountersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCountersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DescribeCountersAsync(
	request *DescribeCountersRequest,
	callback chan<- DescribeCountersAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.LimitName != nil {
		queryStrings["limitName"] = core.ToString(*request.LimitName)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeCountersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DescribeCounters(
	request *DescribeCountersRequest,
) (*DescribeCountersResult, error) {
	callback := make(chan DescribeCountersAsyncResult, 1)
	go p.DescribeCountersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCountersByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCountersByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCountersByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeCountersByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCountersByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCountersByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DescribeCountersByUserIdAsync(
	request *DescribeCountersByUserIdRequest,
	callback chan<- DescribeCountersByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter"
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
	if request.LimitName != nil {
		queryStrings["limitName"] = core.ToString(*request.LimitName)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeCountersByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DescribeCountersByUserId(
	request *DescribeCountersByUserIdRequest,
) (*DescribeCountersByUserIdResult, error) {
	callback := make(chan DescribeCountersByUserIdAsyncResult, 1)
	go p.DescribeCountersByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCounterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterResult
	if asyncResult.Err != nil {
		callback <- GetCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) GetCounterAsync(
	request *GetCounterRequest,
	callback chan<- GetCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{limitName}/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go getCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetCounter(
	request *GetCounterRequest,
) (*GetCounterResult, error) {
	callback := make(chan GetCounterAsyncResult, 1)
	go p.GetCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCounterByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCounterByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) GetCounterByUserIdAsync(
	request *GetCounterByUserIdRequest,
	callback chan<- GetCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{limitName}/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go getCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetCounterByUserId(
	request *GetCounterByUserIdRequest,
) (*GetCounterByUserIdResult, error) {
	callback := make(chan GetCounterByUserIdAsyncResult, 1)
	go p.GetCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func countUpAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CountUpAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountUpAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountUpResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "limit.counter.overflow" {
				asyncResult.Err = gs2err.SetClientError(Overflow{})
			}
		}
		callback <- CountUpAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountUpAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountUpAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CountUpAsync(
	request *CountUpRequest,
	callback chan<- CountUpAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{limitName}/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CountUpValue != nil {
		bodies["countUpValue"] = *request.CountUpValue
	}
	if request.MaxValue != nil {
		bodies["maxValue"] = *request.MaxValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go countUpAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CountUp(
	request *CountUpRequest,
) (*CountUpResult, error) {
	callback := make(chan CountUpAsyncResult, 1)
	go p.CountUpAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func countUpByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CountUpByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountUpByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountUpByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "limit.counter.overflow" {
				asyncResult.Err = gs2err.SetClientError(Overflow{})
			}
		}
		callback <- CountUpByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountUpByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountUpByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CountUpByUserIdAsync(
	request *CountUpByUserIdRequest,
	callback chan<- CountUpByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{limitName}/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CountUpValue != nil {
		bodies["countUpValue"] = *request.CountUpValue
	}
	if request.MaxValue != nil {
		bodies["maxValue"] = *request.MaxValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go countUpByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CountUpByUserId(
	request *CountUpByUserIdRequest,
) (*CountUpByUserIdResult, error) {
	callback := make(chan CountUpByUserIdAsyncResult, 1)
	go p.CountUpByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func countDownByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CountDownByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountDownByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountDownByUserIdResult
	if asyncResult.Err != nil {
		callback <- CountDownByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountDownByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountDownByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CountDownByUserIdAsync(
	request *CountDownByUserIdRequest,
	callback chan<- CountDownByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{limitName}/{counterName}/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.CountDownValue != nil {
		bodies["countDownValue"] = *request.CountDownValue
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go countDownByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CountDownByUserId(
	request *CountDownByUserIdRequest,
) (*CountDownByUserIdResult, error) {
	callback := make(chan CountDownByUserIdAsyncResult, 1)
	go p.CountDownByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCounterByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCounterByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DeleteCounterByUserIdAsync(
	request *DeleteCounterByUserIdRequest,
	callback chan<- DeleteCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{limitName}/{counterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go deleteCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DeleteCounterByUserId(
	request *DeleteCounterByUserIdRequest,
) (*DeleteCounterByUserIdResult, error) {
	callback := make(chan DeleteCounterByUserIdAsyncResult, 1)
	go p.DeleteCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCounterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) VerifyCounterAsync(
	request *VerifyCounterRequest,
	callback chan<- VerifyCounterAsyncResult,
) {
	path := "/{namespaceName}/user/me/counter/{limitName}/{counterName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go verifyCounterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) VerifyCounter(
	request *VerifyCounterRequest,
) (*VerifyCounterResult, error) {
	callback := make(chan VerifyCounterAsyncResult, 1)
	go p.VerifyCounterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCounterByUserIdAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) VerifyCounterByUserIdAsync(
	request *VerifyCounterByUserIdRequest,
	callback chan<- VerifyCounterByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/counter/{limitName}/{counterName}/verify/{verifyType}"
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
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}
	if request.CounterName != nil && *request.CounterName != "" {
		path = strings.ReplaceAll(path, "{counterName}", core.ToString(*request.CounterName))
	} else {
		path = strings.ReplaceAll(path, "{counterName}", "null")
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
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go verifyCounterByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) VerifyCounterByUserId(
	request *VerifyCounterByUserIdRequest,
) (*VerifyCounterByUserIdResult, error) {
	callback := make(chan VerifyCounterByUserIdAsyncResult, 1)
	go p.VerifyCounterByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func countUpByStampTaskAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CountUpByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountUpByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountUpByStampTaskResult
	if asyncResult.Err != nil {
		callback <- CountUpByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountUpByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountUpByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CountUpByStampTaskAsync(
	request *CountUpByStampTaskRequest,
	callback chan<- CountUpByStampTaskAsyncResult,
) {
	path := "/stamp/counter/increase"

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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go countUpByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CountUpByStampTask(
	request *CountUpByStampTaskRequest,
) (*CountUpByStampTaskResult, error) {
	callback := make(chan CountUpByStampTaskAsyncResult, 1)
	go p.CountUpByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func countDownByStampSheetAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CountDownByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CountDownByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CountDownByStampSheetResult
	if asyncResult.Err != nil {
		callback <- CountDownByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CountDownByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CountDownByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CountDownByStampSheetAsync(
	request *CountDownByStampSheetRequest,
	callback chan<- CountDownByStampSheetAsyncResult,
) {
	path := "/stamp/counter/decrease"

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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go countDownByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CountDownByStampSheet(
	request *CountDownByStampSheetRequest,
) (*CountDownByStampSheetResult, error) {
	callback := make(chan CountDownByStampSheetAsyncResult, 1)
	go p.CountDownByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteByStampSheetAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DeleteByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DeleteByStampSheetAsync(
	request *DeleteByStampSheetRequest,
	callback chan<- DeleteByStampSheetAsyncResult,
) {
	path := "/stamp/counter/delete"

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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DeleteByStampSheet(
	request *DeleteByStampSheetRequest,
) (*DeleteByStampSheetResult, error) {
	callback := make(chan DeleteByStampSheetAsyncResult, 1)
	go p.DeleteByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCounterByStampTaskAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCounterByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCounterByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCounterByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyCounterByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCounterByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCounterByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) VerifyCounterByStampTaskAsync(
	request *VerifyCounterByStampTaskRequest,
	callback chan<- VerifyCounterByStampTaskAsyncResult,
) {
	path := "/stamp/counter/verify"

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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go verifyCounterByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) VerifyCounterByStampTask(
	request *VerifyCounterByStampTaskRequest,
) (*VerifyCounterByStampTaskResult, error) {
	callback := make(chan VerifyCounterByStampTaskAsyncResult, 1)
	go p.VerifyCounterByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLimitModelMastersAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLimitModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLimitModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLimitModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeLimitModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLimitModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DescribeLimitModelMastersAsync(
	request *DescribeLimitModelMastersRequest,
	callback chan<- DescribeLimitModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/limit"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeLimitModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DescribeLimitModelMasters(
	request *DescribeLimitModelMastersRequest,
) (*DescribeLimitModelMastersResult, error) {
	callback := make(chan DescribeLimitModelMastersAsyncResult, 1)
	go p.DescribeLimitModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createLimitModelMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- CreateLimitModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateLimitModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateLimitModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateLimitModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateLimitModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateLimitModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) CreateLimitModelMasterAsync(
	request *CreateLimitModelMasterRequest,
	callback chan<- CreateLimitModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/limit"
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
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ResetDayOfMonth != nil {
		bodies["resetDayOfMonth"] = *request.ResetDayOfMonth
	}
	if request.ResetDayOfWeek != nil && *request.ResetDayOfWeek != "" {
		bodies["resetDayOfWeek"] = *request.ResetDayOfWeek
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createLimitModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) CreateLimitModelMaster(
	request *CreateLimitModelMasterRequest,
) (*CreateLimitModelMasterResult, error) {
	callback := make(chan CreateLimitModelMasterAsyncResult, 1)
	go p.CreateLimitModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLimitModelMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- GetLimitModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLimitModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLimitModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetLimitModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLimitModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) GetLimitModelMasterAsync(
	request *GetLimitModelMasterRequest,
	callback chan<- GetLimitModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/limit/{limitName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getLimitModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetLimitModelMaster(
	request *GetLimitModelMasterRequest,
) (*GetLimitModelMasterResult, error) {
	callback := make(chan GetLimitModelMasterAsyncResult, 1)
	go p.GetLimitModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateLimitModelMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateLimitModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateLimitModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateLimitModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateLimitModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateLimitModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateLimitModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) UpdateLimitModelMasterAsync(
	request *UpdateLimitModelMasterRequest,
	callback chan<- UpdateLimitModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/limit/{limitName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ResetType != nil && *request.ResetType != "" {
		bodies["resetType"] = *request.ResetType
	}
	if request.ResetDayOfMonth != nil {
		bodies["resetDayOfMonth"] = *request.ResetDayOfMonth
	}
	if request.ResetDayOfWeek != nil && *request.ResetDayOfWeek != "" {
		bodies["resetDayOfWeek"] = *request.ResetDayOfWeek
	}
	if request.ResetHour != nil {
		bodies["resetHour"] = *request.ResetHour
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateLimitModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) UpdateLimitModelMaster(
	request *UpdateLimitModelMasterRequest,
) (*UpdateLimitModelMasterResult, error) {
	callback := make(chan UpdateLimitModelMasterAsyncResult, 1)
	go p.UpdateLimitModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLimitModelMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLimitModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLimitModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLimitModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteLimitModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteLimitModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteLimitModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DeleteLimitModelMasterAsync(
	request *DeleteLimitModelMasterRequest,
	callback chan<- DeleteLimitModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/limit/{limitName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteLimitModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DeleteLimitModelMaster(
	request *DeleteLimitModelMasterRequest,
) (*DeleteLimitModelMasterResult, error) {
	callback := make(chan DeleteLimitModelMasterAsyncResult, 1)
	go p.DeleteLimitModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2LimitRestClient,
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

func (p Gs2LimitRestClient) ExportMasterAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) ExportMaster(
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

func getCurrentLimitMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentLimitMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentLimitMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentLimitMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentLimitMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentLimitMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentLimitMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) GetCurrentLimitMasterAsync(
	request *GetCurrentLimitMasterRequest,
	callback chan<- GetCurrentLimitMasterAsyncResult,
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCurrentLimitMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetCurrentLimitMaster(
	request *GetCurrentLimitMasterRequest,
) (*GetCurrentLimitMasterResult, error) {
	callback := make(chan GetCurrentLimitMasterAsyncResult, 1)
	go p.GetCurrentLimitMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLimitMasterAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLimitMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLimitMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLimitMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLimitMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentLimitMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentLimitMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) UpdateCurrentLimitMasterAsync(
	request *UpdateCurrentLimitMasterRequest,
	callback chan<- UpdateCurrentLimitMasterAsyncResult,
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentLimitMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) UpdateCurrentLimitMaster(
	request *UpdateCurrentLimitMasterRequest,
) (*UpdateCurrentLimitMasterResult, error) {
	callback := make(chan UpdateCurrentLimitMasterAsyncResult, 1)
	go p.UpdateCurrentLimitMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLimitMasterFromGitHubAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLimitMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLimitMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLimitMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLimitMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentLimitMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentLimitMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) UpdateCurrentLimitMasterFromGitHubAsync(
	request *UpdateCurrentLimitMasterFromGitHubRequest,
	callback chan<- UpdateCurrentLimitMasterFromGitHubAsyncResult,
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateCurrentLimitMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) UpdateCurrentLimitMasterFromGitHub(
	request *UpdateCurrentLimitMasterFromGitHubRequest,
) (*UpdateCurrentLimitMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentLimitMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentLimitMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLimitModelsAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLimitModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLimitModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLimitModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeLimitModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLimitModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) DescribeLimitModelsAsync(
	request *DescribeLimitModelsRequest,
	callback chan<- DescribeLimitModelsAsyncResult,
) {
	path := "/{namespaceName}/limit"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeLimitModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) DescribeLimitModels(
	request *DescribeLimitModelsRequest,
) (*DescribeLimitModelsResult, error) {
	callback := make(chan DescribeLimitModelsAsyncResult, 1)
	go p.DescribeLimitModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLimitModelAsyncHandler(
	client Gs2LimitRestClient,
	job *core.NetworkJob,
	callback chan<- GetLimitModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLimitModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLimitModelResult
	if asyncResult.Err != nil {
		callback <- GetLimitModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLimitModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LimitRestClient) GetLimitModelAsync(
	request *GetLimitModelRequest,
	callback chan<- GetLimitModelAsyncResult,
) {
	path := "/{namespaceName}/limit/{limitName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getLimitModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("limit").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LimitRestClient) GetLimitModel(
	request *GetLimitModelRequest,
) (*GetLimitModelResult, error) {
	callback := make(chan GetLimitModelAsyncResult, 1)
	go p.GetLimitModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
