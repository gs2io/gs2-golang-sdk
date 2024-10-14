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

package ranking2

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2Ranking2RestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2Ranking2RestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) DescribeNamespacesAsync(
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeNamespaces(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) CreateNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateNamespace(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) GetNamespaceStatusAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetNamespaceStatus(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) GetNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetNamespace(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) UpdateNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateNamespace(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) DeleteNamespaceAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteNamespace(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) DumpUserDataByUserIdAsync(
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DumpUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) CheckDumpUserDataByUserIdAsync(
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CheckDumpUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) CleanUserDataByUserIdAsync(
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CleanUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) CheckCleanUserDataByUserIdAsync(
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CheckCleanUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) PrepareImportUserDataByUserIdAsync(
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PrepareImportUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) ImportUserDataByUserIdAsync(
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ImportUserDataByUserId(
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
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) CheckImportUserDataByUserIdAsync(
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CheckImportUserDataByUserId(
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

func describeGlobalRankingModelsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingModelsAsync(
	request *DescribeGlobalRankingModelsRequest,
	callback chan<- DescribeGlobalRankingModelsAsyncResult,
) {
	path := "/{namespaceName}/model/global"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeGlobalRankingModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingModels(
	request *DescribeGlobalRankingModelsRequest,
) (*DescribeGlobalRankingModelsResult, error) {
	callback := make(chan DescribeGlobalRankingModelsAsyncResult, 1)
	go p.DescribeGlobalRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingModelAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingModelResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingModelAsync(
	request *GetGlobalRankingModelRequest,
	callback chan<- GetGlobalRankingModelAsyncResult,
) {
	path := "/{namespaceName}/model/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getGlobalRankingModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingModel(
	request *GetGlobalRankingModelRequest,
) (*GetGlobalRankingModelResult, error) {
	callback := make(chan GetGlobalRankingModelAsyncResult, 1)
	go p.GetGlobalRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingModelMastersAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingModelMastersAsync(
	request *DescribeGlobalRankingModelMastersRequest,
	callback chan<- DescribeGlobalRankingModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/global"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeGlobalRankingModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingModelMasters(
	request *DescribeGlobalRankingModelMastersRequest,
) (*DescribeGlobalRankingModelMastersResult, error) {
	callback := make(chan DescribeGlobalRankingModelMastersAsyncResult, 1)
	go p.DescribeGlobalRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGlobalRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateGlobalRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateGlobalRankingModelMasterAsync(
	request *CreateGlobalRankingModelMasterRequest,
	callback chan<- CreateGlobalRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/global"
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
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createGlobalRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateGlobalRankingModelMaster(
	request *CreateGlobalRankingModelMasterRequest,
) (*CreateGlobalRankingModelMasterResult, error) {
	callback := make(chan CreateGlobalRankingModelMasterAsyncResult, 1)
	go p.CreateGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingModelMasterAsync(
	request *GetGlobalRankingModelMasterRequest,
	callback chan<- GetGlobalRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getGlobalRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingModelMaster(
	request *GetGlobalRankingModelMasterRequest,
) (*GetGlobalRankingModelMasterResult, error) {
	callback := make(chan GetGlobalRankingModelMasterAsyncResult, 1)
	go p.GetGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGlobalRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGlobalRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateGlobalRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) UpdateGlobalRankingModelMasterAsync(
	request *UpdateGlobalRankingModelMasterRequest,
	callback chan<- UpdateGlobalRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
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
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateGlobalRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateGlobalRankingModelMaster(
	request *UpdateGlobalRankingModelMasterRequest,
) (*UpdateGlobalRankingModelMasterResult, error) {
	callback := make(chan UpdateGlobalRankingModelMasterAsyncResult, 1)
	go p.UpdateGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGlobalRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGlobalRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteGlobalRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGlobalRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingModelMasterAsync(
	request *DeleteGlobalRankingModelMasterRequest,
	callback chan<- DeleteGlobalRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteGlobalRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingModelMaster(
	request *DeleteGlobalRankingModelMasterRequest,
) (*DeleteGlobalRankingModelMasterResult, error) {
	callback := make(chan DeleteGlobalRankingModelMasterAsyncResult, 1)
	go p.DeleteGlobalRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingScoresAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingScoresResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingScoresAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingScoresAsync(
	request *DescribeGlobalRankingScoresRequest,
	callback chan<- DescribeGlobalRankingScoresAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/global"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeGlobalRankingScoresAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingScores(
	request *DescribeGlobalRankingScoresRequest,
) (*DescribeGlobalRankingScoresResult, error) {
	callback := make(chan DescribeGlobalRankingScoresAsyncResult, 1)
	go p.DescribeGlobalRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingScoresByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingScoresByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingScoresByUserIdAsync(
	request *DescribeGlobalRankingScoresByUserIdRequest,
	callback chan<- DescribeGlobalRankingScoresByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeGlobalRankingScoresByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingScoresByUserId(
	request *DescribeGlobalRankingScoresByUserIdRequest,
) (*DescribeGlobalRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putGlobalRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutGlobalRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutGlobalRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutGlobalRankingScoreResult
	if asyncResult.Err != nil {
		callback <- PutGlobalRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutGlobalRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutGlobalRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutGlobalRankingScoreAsync(
	request *PutGlobalRankingScoreRequest,
	callback chan<- PutGlobalRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
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

	go putGlobalRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutGlobalRankingScore(
	request *PutGlobalRankingScoreRequest,
) (*PutGlobalRankingScoreResult, error) {
	callback := make(chan PutGlobalRankingScoreAsyncResult, 1)
	go p.PutGlobalRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putGlobalRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutGlobalRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- PutGlobalRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutGlobalRankingScoreByUserIdAsync(
	request *PutGlobalRankingScoreByUserIdRequest,
	callback chan<- PutGlobalRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
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

	go putGlobalRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutGlobalRankingScoreByUserId(
	request *PutGlobalRankingScoreByUserIdRequest,
) (*PutGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan PutGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.PutGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingScoreResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingScoreAsync(
	request *GetGlobalRankingScoreRequest,
	callback chan<- GetGlobalRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingScore(
	request *GetGlobalRankingScoreRequest,
) (*GetGlobalRankingScoreResult, error) {
	callback := make(chan GetGlobalRankingScoreAsyncResult, 1)
	go p.GetGlobalRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingScoreByUserIdAsync(
	request *GetGlobalRankingScoreByUserIdRequest,
	callback chan<- GetGlobalRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingScoreByUserId(
	request *GetGlobalRankingScoreByUserIdRequest,
) (*GetGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan GetGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.GetGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGlobalRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingScoreByUserIdAsync(
	request *DeleteGlobalRankingScoreByUserIdRequest,
	callback chan<- DeleteGlobalRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global/{rankingName}/{season}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteGlobalRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingScoreByUserId(
	request *DeleteGlobalRankingScoreByUserIdRequest,
) (*DeleteGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyGlobalRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyGlobalRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGlobalRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGlobalRankingScoreResult
	if asyncResult.Err != nil {
		callback <- VerifyGlobalRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGlobalRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyGlobalRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScoreAsync(
	request *VerifyGlobalRankingScoreRequest,
	callback chan<- VerifyGlobalRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global/{rankingName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifyGlobalRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScore(
	request *VerifyGlobalRankingScoreRequest,
) (*VerifyGlobalRankingScoreResult, error) {
	callback := make(chan VerifyGlobalRankingScoreAsyncResult, 1)
	go p.VerifyGlobalRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyGlobalRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyGlobalRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGlobalRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGlobalRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyGlobalRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGlobalRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyGlobalRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScoreByUserIdAsync(
	request *VerifyGlobalRankingScoreByUserIdRequest,
	callback chan<- VerifyGlobalRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/global/{rankingName}/verify/{verifyType}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifyGlobalRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScoreByUserId(
	request *VerifyGlobalRankingScoreByUserIdRequest,
) (*VerifyGlobalRankingScoreByUserIdResult, error) {
	callback := make(chan VerifyGlobalRankingScoreByUserIdAsyncResult, 1)
	go p.VerifyGlobalRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyGlobalRankingScoreByStampTaskAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyGlobalRankingScoreByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyGlobalRankingScoreByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyGlobalRankingScoreByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyGlobalRankingScoreByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyGlobalRankingScoreByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyGlobalRankingScoreByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScoreByStampTaskAsync(
	request *VerifyGlobalRankingScoreByStampTaskRequest,
	callback chan<- VerifyGlobalRankingScoreByStampTaskAsyncResult,
) {
	path := "/stamp/global/score/verify"

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

	go verifyGlobalRankingScoreByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyGlobalRankingScoreByStampTask(
	request *VerifyGlobalRankingScoreByStampTaskRequest,
) (*VerifyGlobalRankingScoreByStampTaskResult, error) {
	callback := make(chan VerifyGlobalRankingScoreByStampTaskAsyncResult, 1)
	go p.VerifyGlobalRankingScoreByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingReceivedRewardsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingReceivedRewardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingReceivedRewardsResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingReceivedRewardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingReceivedRewardsAsync(
	request *DescribeGlobalRankingReceivedRewardsRequest,
	callback chan<- DescribeGlobalRankingReceivedRewardsAsyncResult,
) {
	path := "/{namespaceName}/user/me/global/reward/received"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeGlobalRankingReceivedRewardsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingReceivedRewards(
	request *DescribeGlobalRankingReceivedRewardsRequest,
) (*DescribeGlobalRankingReceivedRewardsResult, error) {
	callback := make(chan DescribeGlobalRankingReceivedRewardsAsyncResult, 1)
	go p.DescribeGlobalRankingReceivedRewardsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingReceivedRewardsByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingReceivedRewardsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingReceivedRewardsByUserIdAsync(
	request *DescribeGlobalRankingReceivedRewardsByUserIdRequest,
	callback chan<- DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/global/reward/received"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeGlobalRankingReceivedRewardsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingReceivedRewardsByUserId(
	request *DescribeGlobalRankingReceivedRewardsByUserIdRequest,
) (*DescribeGlobalRankingReceivedRewardsByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingReceivedRewardsByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingReceivedRewardsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGlobalRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- CreateGlobalRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedRewardAsync(
	request *CreateGlobalRankingReceivedRewardRequest,
	callback chan<- CreateGlobalRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/global/reward/received/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go createGlobalRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedReward(
	request *CreateGlobalRankingReceivedRewardRequest,
) (*CreateGlobalRankingReceivedRewardResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGlobalRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedRewardByUserIdAsync(
	request *CreateGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- CreateGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/global/reward/received/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go createGlobalRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedRewardByUserId(
	request *CreateGlobalRankingReceivedRewardByUserIdRequest,
) (*CreateGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveGlobalRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) ReceiveGlobalRankingReceivedRewardAsync(
	request *ReceiveGlobalRankingReceivedRewardRequest,
	callback chan<- ReceiveGlobalRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/global/reward/received/{rankingName}/{season}/reward/receive"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.Season != nil {
		path = strings.ReplaceAll(path, "{season}", core.ToString(*request.Season))
	} else {
		path = strings.ReplaceAll(path, "{season}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go receiveGlobalRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ReceiveGlobalRankingReceivedReward(
	request *ReceiveGlobalRankingReceivedRewardRequest,
) (*ReceiveGlobalRankingReceivedRewardResult, error) {
	callback := make(chan ReceiveGlobalRankingReceivedRewardAsyncResult, 1)
	go p.ReceiveGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveGlobalRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) ReceiveGlobalRankingReceivedRewardByUserIdAsync(
	request *ReceiveGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/global/reward/received/{rankingName}/{season}/reward/receive"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.Season != nil {
		path = strings.ReplaceAll(path, "{season}", core.ToString(*request.Season))
	} else {
		path = strings.ReplaceAll(path, "{season}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go receiveGlobalRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ReceiveGlobalRankingReceivedRewardByUserId(
	request *ReceiveGlobalRankingReceivedRewardByUserIdRequest,
) (*ReceiveGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan ReceiveGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.ReceiveGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingReceivedRewardAsync(
	request *GetGlobalRankingReceivedRewardRequest,
	callback chan<- GetGlobalRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/global/reward/received/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingReceivedReward(
	request *GetGlobalRankingReceivedRewardRequest,
) (*GetGlobalRankingReceivedRewardResult, error) {
	callback := make(chan GetGlobalRankingReceivedRewardAsyncResult, 1)
	go p.GetGlobalRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingReceivedRewardByUserIdAsync(
	request *GetGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- GetGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/global/reward/received/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingReceivedRewardByUserId(
	request *GetGlobalRankingReceivedRewardByUserIdRequest,
) (*GetGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan GetGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.GetGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGlobalRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingReceivedRewardByUserIdAsync(
	request *DeleteGlobalRankingReceivedRewardByUserIdRequest,
	callback chan<- DeleteGlobalRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/global/reward/received/{rankingName}/{season}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteGlobalRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteGlobalRankingReceivedRewardByUserId(
	request *DeleteGlobalRankingReceivedRewardByUserIdRequest,
) (*DeleteGlobalRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan DeleteGlobalRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.DeleteGlobalRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGlobalRankingReceivedRewardByStampTaskAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalRankingReceivedRewardByStampTaskResult
	if asyncResult.Err != nil {
		callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedRewardByStampTaskAsync(
	request *CreateGlobalRankingReceivedRewardByStampTaskRequest,
	callback chan<- CreateGlobalRankingReceivedRewardByStampTaskAsyncResult,
) {
	path := "/stamp/ranking/global/reward/receive"

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

	go createGlobalRankingReceivedRewardByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateGlobalRankingReceivedRewardByStampTask(
	request *CreateGlobalRankingReceivedRewardByStampTaskRequest,
) (*CreateGlobalRankingReceivedRewardByStampTaskResult, error) {
	callback := make(chan CreateGlobalRankingReceivedRewardByStampTaskAsyncResult, 1)
	go p.CreateGlobalRankingReceivedRewardByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingsResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingsAsync(
	request *DescribeGlobalRankingsRequest,
	callback chan<- DescribeGlobalRankingsAsyncResult,
) {
	path := "/{namespaceName}/ranking/global/{rankingName}/user/me"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeGlobalRankingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankings(
	request *DescribeGlobalRankingsRequest,
) (*DescribeGlobalRankingsResult, error) {
	callback := make(chan DescribeGlobalRankingsAsyncResult, 1)
	go p.DescribeGlobalRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalRankingsByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalRankingsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalRankingsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingsByUserIdAsync(
	request *DescribeGlobalRankingsByUserIdRequest,
	callback chan<- DescribeGlobalRankingsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/global/{rankingName}/user/{userId}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeGlobalRankingsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeGlobalRankingsByUserId(
	request *DescribeGlobalRankingsByUserIdRequest,
) (*DescribeGlobalRankingsByUserIdResult, error) {
	callback := make(chan DescribeGlobalRankingsByUserIdAsyncResult, 1)
	go p.DescribeGlobalRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingAsync(
	request *GetGlobalRankingRequest,
	callback chan<- GetGlobalRankingAsyncResult,
) {
	path := "/{namespaceName}/ranking/global/{rankingName}/user/me/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRanking(
	request *GetGlobalRankingRequest,
) (*GetGlobalRankingResult, error) {
	callback := make(chan GetGlobalRankingAsyncResult, 1)
	go p.GetGlobalRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalRankingByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalRankingByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetGlobalRankingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetGlobalRankingByUserIdAsync(
	request *GetGlobalRankingByUserIdRequest,
	callback chan<- GetGlobalRankingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/global/{rankingName}/user/{userId}/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getGlobalRankingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetGlobalRankingByUserId(
	request *GetGlobalRankingByUserIdRequest,
) (*GetGlobalRankingByUserIdResult, error) {
	callback := make(chan GetGlobalRankingByUserIdAsyncResult, 1)
	go p.GetGlobalRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingModelsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingModelsAsync(
	request *DescribeClusterRankingModelsRequest,
	callback chan<- DescribeClusterRankingModelsAsyncResult,
) {
	path := "/{namespaceName}/model/cluster"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeClusterRankingModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingModels(
	request *DescribeClusterRankingModelsRequest,
) (*DescribeClusterRankingModelsResult, error) {
	callback := make(chan DescribeClusterRankingModelsAsyncResult, 1)
	go p.DescribeClusterRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingModelAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingModelResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingModelAsync(
	request *GetClusterRankingModelRequest,
	callback chan<- GetClusterRankingModelAsyncResult,
) {
	path := "/{namespaceName}/model/cluster/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getClusterRankingModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingModel(
	request *GetClusterRankingModelRequest,
) (*GetClusterRankingModelResult, error) {
	callback := make(chan GetClusterRankingModelAsyncResult, 1)
	go p.GetClusterRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingModelMastersAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingModelMastersAsync(
	request *DescribeClusterRankingModelMastersRequest,
	callback chan<- DescribeClusterRankingModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/cluster"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeClusterRankingModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingModelMasters(
	request *DescribeClusterRankingModelMastersRequest,
) (*DescribeClusterRankingModelMastersResult, error) {
	callback := make(chan DescribeClusterRankingModelMastersAsyncResult, 1)
	go p.DescribeClusterRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createClusterRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateClusterRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateClusterRankingModelMasterAsync(
	request *CreateClusterRankingModelMasterRequest,
	callback chan<- CreateClusterRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/cluster"
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
	if request.ClusterType != nil && *request.ClusterType != "" {
		bodies["clusterType"] = *request.ClusterType
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createClusterRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateClusterRankingModelMaster(
	request *CreateClusterRankingModelMasterRequest,
) (*CreateClusterRankingModelMasterResult, error) {
	callback := make(chan CreateClusterRankingModelMasterAsyncResult, 1)
	go p.CreateClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingModelMasterAsync(
	request *GetClusterRankingModelMasterRequest,
	callback chan<- GetClusterRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/cluster/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getClusterRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingModelMaster(
	request *GetClusterRankingModelMasterRequest,
) (*GetClusterRankingModelMasterResult, error) {
	callback := make(chan GetClusterRankingModelMasterAsyncResult, 1)
	go p.GetClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateClusterRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- UpdateClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateClusterRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateClusterRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) UpdateClusterRankingModelMasterAsync(
	request *UpdateClusterRankingModelMasterRequest,
	callback chan<- UpdateClusterRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/cluster/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ClusterType != nil && *request.ClusterType != "" {
		bodies["clusterType"] = *request.ClusterType
	}
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
	}
	if request.RankingRewards != nil {
		var _rankingRewards []interface{}
		for _, item := range request.RankingRewards {
			_rankingRewards = append(_rankingRewards, item)
		}
		bodies["rankingRewards"] = _rankingRewards
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateClusterRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateClusterRankingModelMaster(
	request *UpdateClusterRankingModelMasterRequest,
) (*UpdateClusterRankingModelMasterResult, error) {
	callback := make(chan UpdateClusterRankingModelMasterAsyncResult, 1)
	go p.UpdateClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteClusterRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteClusterRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteClusterRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteClusterRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteClusterRankingModelMasterAsync(
	request *DeleteClusterRankingModelMasterRequest,
	callback chan<- DeleteClusterRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/cluster/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteClusterRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteClusterRankingModelMaster(
	request *DeleteClusterRankingModelMasterRequest,
) (*DeleteClusterRankingModelMasterResult, error) {
	callback := make(chan DeleteClusterRankingModelMasterAsyncResult, 1)
	go p.DeleteClusterRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingScoresAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingScoresResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingScoresAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingScoresAsync(
	request *DescribeClusterRankingScoresRequest,
	callback chan<- DescribeClusterRankingScoresAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/cluster"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.ClusterName != nil {
		queryStrings["clusterName"] = core.ToString(*request.ClusterName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeClusterRankingScoresAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingScores(
	request *DescribeClusterRankingScoresRequest,
) (*DescribeClusterRankingScoresResult, error) {
	callback := make(chan DescribeClusterRankingScoresAsyncResult, 1)
	go p.DescribeClusterRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingScoresByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingScoresByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingScoresByUserIdAsync(
	request *DescribeClusterRankingScoresByUserIdRequest,
	callback chan<- DescribeClusterRankingScoresByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.ClusterName != nil {
		queryStrings["clusterName"] = core.ToString(*request.ClusterName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeClusterRankingScoresByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingScoresByUserId(
	request *DescribeClusterRankingScoresByUserIdRequest,
) (*DescribeClusterRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putClusterRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutClusterRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutClusterRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutClusterRankingScoreResult
	if asyncResult.Err != nil {
		callback <- PutClusterRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutClusterRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutClusterRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutClusterRankingScoreAsync(
	request *PutClusterRankingScoreRequest,
	callback chan<- PutClusterRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/cluster/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
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

	go putClusterRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutClusterRankingScore(
	request *PutClusterRankingScoreRequest,
) (*PutClusterRankingScoreResult, error) {
	callback := make(chan PutClusterRankingScoreAsyncResult, 1)
	go p.PutClusterRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putClusterRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutClusterRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- PutClusterRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutClusterRankingScoreByUserIdAsync(
	request *PutClusterRankingScoreByUserIdRequest,
	callback chan<- PutClusterRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
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

	go putClusterRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutClusterRankingScoreByUserId(
	request *PutClusterRankingScoreByUserIdRequest,
) (*PutClusterRankingScoreByUserIdResult, error) {
	callback := make(chan PutClusterRankingScoreByUserIdAsyncResult, 1)
	go p.PutClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingScoreResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingScoreAsync(
	request *GetClusterRankingScoreRequest,
	callback chan<- GetClusterRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/cluster/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingScore(
	request *GetClusterRankingScoreRequest,
) (*GetClusterRankingScoreResult, error) {
	callback := make(chan GetClusterRankingScoreAsyncResult, 1)
	go p.GetClusterRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingScoreByUserIdAsync(
	request *GetClusterRankingScoreByUserIdRequest,
	callback chan<- GetClusterRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingScoreByUserId(
	request *GetClusterRankingScoreByUserIdRequest,
) (*GetClusterRankingScoreByUserIdResult, error) {
	callback := make(chan GetClusterRankingScoreByUserIdAsyncResult, 1)
	go p.GetClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteClusterRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteClusterRankingScoreByUserIdAsync(
	request *DeleteClusterRankingScoreByUserIdRequest,
	callback chan<- DeleteClusterRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster/{rankingName}/{clusterName}/{season}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteClusterRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteClusterRankingScoreByUserId(
	request *DeleteClusterRankingScoreByUserIdRequest,
) (*DeleteClusterRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteClusterRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyClusterRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyClusterRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyClusterRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyClusterRankingScoreResult
	if asyncResult.Err != nil {
		callback <- VerifyClusterRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyClusterRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyClusterRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScoreAsync(
	request *VerifyClusterRankingScoreRequest,
	callback chan<- VerifyClusterRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster/{rankingName}/{clusterName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifyClusterRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScore(
	request *VerifyClusterRankingScoreRequest,
) (*VerifyClusterRankingScoreResult, error) {
	callback := make(chan VerifyClusterRankingScoreAsyncResult, 1)
	go p.VerifyClusterRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyClusterRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyClusterRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyClusterRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyClusterRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyClusterRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyClusterRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyClusterRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScoreByUserIdAsync(
	request *VerifyClusterRankingScoreByUserIdRequest,
	callback chan<- VerifyClusterRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/cluster/{rankingName}/{clusterName}/verify/{verifyType}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifyClusterRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScoreByUserId(
	request *VerifyClusterRankingScoreByUserIdRequest,
) (*VerifyClusterRankingScoreByUserIdResult, error) {
	callback := make(chan VerifyClusterRankingScoreByUserIdAsyncResult, 1)
	go p.VerifyClusterRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyClusterRankingScoreByStampTaskAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyClusterRankingScoreByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyClusterRankingScoreByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyClusterRankingScoreByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyClusterRankingScoreByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyClusterRankingScoreByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyClusterRankingScoreByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScoreByStampTaskAsync(
	request *VerifyClusterRankingScoreByStampTaskRequest,
	callback chan<- VerifyClusterRankingScoreByStampTaskAsyncResult,
) {
	path := "/stamp/cluster/score/verify"

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

	go verifyClusterRankingScoreByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifyClusterRankingScoreByStampTask(
	request *VerifyClusterRankingScoreByStampTaskRequest,
) (*VerifyClusterRankingScoreByStampTaskResult, error) {
	callback := make(chan VerifyClusterRankingScoreByStampTaskAsyncResult, 1)
	go p.VerifyClusterRankingScoreByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingReceivedRewardsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingReceivedRewardsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingReceivedRewardsResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingReceivedRewardsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingReceivedRewardsAsync(
	request *DescribeClusterRankingReceivedRewardsRequest,
	callback chan<- DescribeClusterRankingReceivedRewardsAsyncResult,
) {
	path := "/{namespaceName}/user/me/cluster/reward/received"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.ClusterName != nil {
		queryStrings["clusterName"] = core.ToString(*request.ClusterName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeClusterRankingReceivedRewardsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingReceivedRewards(
	request *DescribeClusterRankingReceivedRewardsRequest,
) (*DescribeClusterRankingReceivedRewardsResult, error) {
	callback := make(chan DescribeClusterRankingReceivedRewardsAsyncResult, 1)
	go p.DescribeClusterRankingReceivedRewardsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingReceivedRewardsByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingReceivedRewardsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingReceivedRewardsByUserIdAsync(
	request *DescribeClusterRankingReceivedRewardsByUserIdRequest,
	callback chan<- DescribeClusterRankingReceivedRewardsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/cluster/reward/received"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
	}
	if request.ClusterName != nil {
		queryStrings["clusterName"] = core.ToString(*request.ClusterName)
	}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeClusterRankingReceivedRewardsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingReceivedRewardsByUserId(
	request *DescribeClusterRankingReceivedRewardsByUserIdRequest,
) (*DescribeClusterRankingReceivedRewardsByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingReceivedRewardsByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingReceivedRewardsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createClusterRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- CreateClusterRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedRewardAsync(
	request *CreateClusterRankingReceivedRewardRequest,
	callback chan<- CreateClusterRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/cluster/reward/received/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go createClusterRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedReward(
	request *CreateClusterRankingReceivedRewardRequest,
) (*CreateClusterRankingReceivedRewardResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createClusterRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedRewardByUserIdAsync(
	request *CreateClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- CreateClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/cluster/reward/received/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
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

	go createClusterRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedRewardByUserId(
	request *CreateClusterRankingReceivedRewardByUserIdRequest,
) (*CreateClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveClusterRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveClusterRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) ReceiveClusterRankingReceivedRewardAsync(
	request *ReceiveClusterRankingReceivedRewardRequest,
	callback chan<- ReceiveClusterRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/cluster/reward/received/{rankingName}/{clusterName}/{season}/reward/receive"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.Season != nil {
		path = strings.ReplaceAll(path, "{season}", core.ToString(*request.Season))
	} else {
		path = strings.ReplaceAll(path, "{season}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go receiveClusterRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ReceiveClusterRankingReceivedReward(
	request *ReceiveClusterRankingReceivedRewardRequest,
) (*ReceiveClusterRankingReceivedRewardResult, error) {
	callback := make(chan ReceiveClusterRankingReceivedRewardAsyncResult, 1)
	go p.ReceiveClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveClusterRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) ReceiveClusterRankingReceivedRewardByUserIdAsync(
	request *ReceiveClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- ReceiveClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/cluster/reward/received/{rankingName}/{clusterName}/{season}/reward/receive"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.Season != nil {
		path = strings.ReplaceAll(path, "{season}", core.ToString(*request.Season))
	} else {
		path = strings.ReplaceAll(path, "{season}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Config != nil {
		var _config []interface{}
		for _, item := range request.Config {
			_config = append(_config, item)
		}
		bodies["config"] = _config
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

	go receiveClusterRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ReceiveClusterRankingReceivedRewardByUserId(
	request *ReceiveClusterRankingReceivedRewardByUserIdRequest,
) (*ReceiveClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan ReceiveClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.ReceiveClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingReceivedRewardAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingReceivedRewardAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingReceivedRewardAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingReceivedRewardResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingReceivedRewardAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingReceivedRewardAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingReceivedRewardAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingReceivedRewardAsync(
	request *GetClusterRankingReceivedRewardRequest,
	callback chan<- GetClusterRankingReceivedRewardAsyncResult,
) {
	path := "/{namespaceName}/user/me/cluster/reward/received/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingReceivedRewardAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingReceivedReward(
	request *GetClusterRankingReceivedRewardRequest,
) (*GetClusterRankingReceivedRewardResult, error) {
	callback := make(chan GetClusterRankingReceivedRewardAsyncResult, 1)
	go p.GetClusterRankingReceivedRewardAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingReceivedRewardByUserIdAsync(
	request *GetClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- GetClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/cluster/reward/received/{rankingName}/{clusterName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingReceivedRewardByUserId(
	request *GetClusterRankingReceivedRewardByUserIdRequest,
) (*GetClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan GetClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.GetClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteClusterRankingReceivedRewardByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteClusterRankingReceivedRewardByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteClusterRankingReceivedRewardByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteClusterRankingReceivedRewardByUserIdAsync(
	request *DeleteClusterRankingReceivedRewardByUserIdRequest,
	callback chan<- DeleteClusterRankingReceivedRewardByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/cluster/reward/received/{rankingName}/{clusterName}/{season}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteClusterRankingReceivedRewardByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteClusterRankingReceivedRewardByUserId(
	request *DeleteClusterRankingReceivedRewardByUserIdRequest,
) (*DeleteClusterRankingReceivedRewardByUserIdResult, error) {
	callback := make(chan DeleteClusterRankingReceivedRewardByUserIdAsyncResult, 1)
	go p.DeleteClusterRankingReceivedRewardByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createClusterRankingReceivedRewardByStampTaskAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateClusterRankingReceivedRewardByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateClusterRankingReceivedRewardByStampTaskResult
	if asyncResult.Err != nil {
		callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateClusterRankingReceivedRewardByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedRewardByStampTaskAsync(
	request *CreateClusterRankingReceivedRewardByStampTaskRequest,
	callback chan<- CreateClusterRankingReceivedRewardByStampTaskAsyncResult,
) {
	path := "/stamp/ranking/cluster/reward/receive"

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

	go createClusterRankingReceivedRewardByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateClusterRankingReceivedRewardByStampTask(
	request *CreateClusterRankingReceivedRewardByStampTaskRequest,
) (*CreateClusterRankingReceivedRewardByStampTaskResult, error) {
	callback := make(chan CreateClusterRankingReceivedRewardByStampTaskAsyncResult, 1)
	go p.CreateClusterRankingReceivedRewardByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingsResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingsAsync(
	request *DescribeClusterRankingsRequest,
	callback chan<- DescribeClusterRankingsAsyncResult,
) {
	path := "/{namespaceName}/ranking/cluster/{rankingName}/{clusterName}/user/me"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeClusterRankingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankings(
	request *DescribeClusterRankingsRequest,
) (*DescribeClusterRankingsResult, error) {
	callback := make(chan DescribeClusterRankingsAsyncResult, 1)
	go p.DescribeClusterRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeClusterRankingsByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeClusterRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeClusterRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeClusterRankingsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeClusterRankingsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeClusterRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeClusterRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeClusterRankingsByUserIdAsync(
	request *DescribeClusterRankingsByUserIdRequest,
	callback chan<- DescribeClusterRankingsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/cluster/{rankingName}/{clusterName}/user/{userId}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeClusterRankingsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeClusterRankingsByUserId(
	request *DescribeClusterRankingsByUserIdRequest,
) (*DescribeClusterRankingsByUserIdResult, error) {
	callback := make(chan DescribeClusterRankingsByUserIdAsyncResult, 1)
	go p.DescribeClusterRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingAsync(
	request *GetClusterRankingRequest,
	callback chan<- GetClusterRankingAsyncResult,
) {
	path := "/{namespaceName}/ranking/cluster/{rankingName}/{clusterName}/user/me/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRanking(
	request *GetClusterRankingRequest,
) (*GetClusterRankingResult, error) {
	callback := make(chan GetClusterRankingAsyncResult, 1)
	go p.GetClusterRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getClusterRankingByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetClusterRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetClusterRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetClusterRankingByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetClusterRankingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetClusterRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetClusterRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetClusterRankingByUserIdAsync(
	request *GetClusterRankingByUserIdRequest,
	callback chan<- GetClusterRankingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/cluster/{rankingName}/{clusterName}/user/{userId}/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.ClusterName != nil && *request.ClusterName != "" {
		path = strings.ReplaceAll(path, "{clusterName}", core.ToString(*request.ClusterName))
	} else {
		path = strings.ReplaceAll(path, "{clusterName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getClusterRankingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetClusterRankingByUserId(
	request *GetClusterRankingByUserIdRequest,
) (*GetClusterRankingByUserIdResult, error) {
	callback := make(chan GetClusterRankingByUserIdAsyncResult, 1)
	go p.GetClusterRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingModelsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingModelsAsync(
	request *DescribeSubscribeRankingModelsRequest,
	callback chan<- DescribeSubscribeRankingModelsAsyncResult,
) {
	path := "/{namespaceName}/model/subscribe"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeSubscribeRankingModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingModels(
	request *DescribeSubscribeRankingModelsRequest,
) (*DescribeSubscribeRankingModelsResult, error) {
	callback := make(chan DescribeSubscribeRankingModelsAsyncResult, 1)
	go p.DescribeSubscribeRankingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingModelAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingModelResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingModelAsync(
	request *GetSubscribeRankingModelRequest,
	callback chan<- GetSubscribeRankingModelAsyncResult,
) {
	path := "/{namespaceName}/model/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSubscribeRankingModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRankingModel(
	request *GetSubscribeRankingModelRequest,
) (*GetSubscribeRankingModelResult, error) {
	callback := make(chan GetSubscribeRankingModelAsyncResult, 1)
	go p.GetSubscribeRankingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingModelMastersAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingModelMastersAsync(
	request *DescribeSubscribeRankingModelMastersRequest,
	callback chan<- DescribeSubscribeRankingModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/subscribe"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeSubscribeRankingModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingModelMasters(
	request *DescribeSubscribeRankingModelMastersRequest,
) (*DescribeSubscribeRankingModelMastersResult, error) {
	callback := make(chan DescribeSubscribeRankingModelMastersAsyncResult, 1)
	go p.DescribeSubscribeRankingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSubscribeRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSubscribeRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateSubscribeRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) CreateSubscribeRankingModelMasterAsync(
	request *CreateSubscribeRankingModelMasterRequest,
	callback chan<- CreateSubscribeRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/subscribe"
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
	if request.MinimumValue != nil {
		bodies["minimumValue"] = *request.MinimumValue
	}
	if request.MaximumValue != nil {
		bodies["maximumValue"] = *request.MaximumValue
	}
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createSubscribeRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) CreateSubscribeRankingModelMaster(
	request *CreateSubscribeRankingModelMasterRequest,
) (*CreateSubscribeRankingModelMasterResult, error) {
	callback := make(chan CreateSubscribeRankingModelMasterAsyncResult, 1)
	go p.CreateSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingModelMasterAsync(
	request *GetSubscribeRankingModelMasterRequest,
	callback chan<- GetSubscribeRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSubscribeRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRankingModelMaster(
	request *GetSubscribeRankingModelMasterRequest,
) (*GetSubscribeRankingModelMasterResult, error) {
	callback := make(chan GetSubscribeRankingModelMasterAsyncResult, 1)
	go p.GetSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSubscribeRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSubscribeRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateSubscribeRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) UpdateSubscribeRankingModelMasterAsync(
	request *UpdateSubscribeRankingModelMasterRequest,
	callback chan<- UpdateSubscribeRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
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
	if request.Sum != nil {
		bodies["sum"] = *request.Sum
	}
	if request.ScoreTtlDays != nil {
		bodies["scoreTtlDays"] = *request.ScoreTtlDays
	}
	if request.OrderDirection != nil && *request.OrderDirection != "" {
		bodies["orderDirection"] = *request.OrderDirection
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateSubscribeRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateSubscribeRankingModelMaster(
	request *UpdateSubscribeRankingModelMasterRequest,
) (*UpdateSubscribeRankingModelMasterResult, error) {
	callback := make(chan UpdateSubscribeRankingModelMasterAsyncResult, 1)
	go p.UpdateSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSubscribeRankingModelMasterAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSubscribeRankingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeRankingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeRankingModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteSubscribeRankingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeRankingModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSubscribeRankingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteSubscribeRankingModelMasterAsync(
	request *DeleteSubscribeRankingModelMasterRequest,
	callback chan<- DeleteSubscribeRankingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteSubscribeRankingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteSubscribeRankingModelMaster(
	request *DeleteSubscribeRankingModelMasterRequest,
) (*DeleteSubscribeRankingModelMasterResult, error) {
	callback := make(chan DeleteSubscribeRankingModelMasterAsyncResult, 1)
	go p.DeleteSubscribeRankingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribesAsync(
	request *DescribeSubscribesRequest,
	callback chan<- DescribeSubscribesAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/score"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeSubscribesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribes(
	request *DescribeSubscribesRequest,
) (*DescribeSubscribesResult, error) {
	callback := make(chan DescribeSubscribesAsyncResult, 1)
	go p.DescribeSubscribesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribesByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribesByUserIdAsync(
	request *DescribeSubscribesByUserIdRequest,
	callback chan<- DescribeSubscribesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/score"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeSubscribesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribesByUserId(
	request *DescribeSubscribesByUserIdRequest,
) (*DescribeSubscribesByUserIdResult, error) {
	callback := make(chan DescribeSubscribesByUserIdAsyncResult, 1)
	go p.DescribeSubscribesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addSubscribeAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- AddSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddSubscribeResult
	if asyncResult.Err != nil {
		callback <- AddSubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) AddSubscribeAsync(
	request *AddSubscribeRequest,
	callback chan<- AddSubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
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

	go addSubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) AddSubscribe(
	request *AddSubscribeRequest,
) (*AddSubscribeResult, error) {
	callback := make(chan AddSubscribeAsyncResult, 1)
	go p.AddSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addSubscribeByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- AddSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddSubscribeByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddSubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) AddSubscribeByUserIdAsync(
	request *AddSubscribeByUserIdRequest,
	callback chan<- AddSubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
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

	go addSubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) AddSubscribeByUserId(
	request *AddSubscribeByUserIdRequest,
) (*AddSubscribeByUserIdResult, error) {
	callback := make(chan AddSubscribeByUserIdAsyncResult, 1)
	go p.AddSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingScoresAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingScoresResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingScoresAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingScoresAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingScoresAsync(
	request *DescribeSubscribeRankingScoresRequest,
	callback chan<- DescribeSubscribeRankingScoresAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/subscribe"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeSubscribeRankingScoresAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingScores(
	request *DescribeSubscribeRankingScoresRequest,
) (*DescribeSubscribeRankingScoresResult, error) {
	callback := make(chan DescribeSubscribeRankingScoresAsyncResult, 1)
	go p.DescribeSubscribeRankingScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingScoresByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingScoresByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingScoresByUserIdAsync(
	request *DescribeSubscribeRankingScoresByUserIdRequest,
	callback chan<- DescribeSubscribeRankingScoresByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe"
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
	if request.RankingName != nil {
		queryStrings["rankingName"] = core.ToString(*request.RankingName)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeSubscribeRankingScoresByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingScoresByUserId(
	request *DescribeSubscribeRankingScoresByUserIdRequest,
) (*DescribeSubscribeRankingScoresByUserIdResult, error) {
	callback := make(chan DescribeSubscribeRankingScoresByUserIdAsyncResult, 1)
	go p.DescribeSubscribeRankingScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putSubscribeRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutSubscribeRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutSubscribeRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutSubscribeRankingScoreResult
	if asyncResult.Err != nil {
		callback <- PutSubscribeRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutSubscribeRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutSubscribeRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutSubscribeRankingScoreAsync(
	request *PutSubscribeRankingScoreRequest,
	callback chan<- PutSubscribeRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
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

	go putSubscribeRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutSubscribeRankingScore(
	request *PutSubscribeRankingScoreRequest,
) (*PutSubscribeRankingScoreResult, error) {
	callback := make(chan PutSubscribeRankingScoreAsyncResult, 1)
	go p.PutSubscribeRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putSubscribeRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- PutSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutSubscribeRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PutSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) PutSubscribeRankingScoreByUserIdAsync(
	request *PutSubscribeRankingScoreByUserIdRequest,
	callback chan<- PutSubscribeRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
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

	go putSubscribeRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) PutSubscribeRankingScoreByUserId(
	request *PutSubscribeRankingScoreByUserIdRequest,
) (*PutSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan PutSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.PutSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingScoreResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingScoreAsync(
	request *GetSubscribeRankingScoreRequest,
	callback chan<- GetSubscribeRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/me/score/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRankingScore(
	request *GetSubscribeRankingScoreRequest,
) (*GetSubscribeRankingScoreResult, error) {
	callback := make(chan GetSubscribeRankingScoreAsyncResult, 1)
	go p.GetSubscribeRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingScoreByUserIdAsync(
	request *GetSubscribeRankingScoreByUserIdRequest,
	callback chan<- GetSubscribeRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe/{rankingName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRankingScoreByUserId(
	request *GetSubscribeRankingScoreByUserIdRequest,
) (*GetSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan GetSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.GetSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSubscribeRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteSubscribeRankingScoreByUserIdAsync(
	request *DeleteSubscribeRankingScoreByUserIdRequest,
	callback chan<- DeleteSubscribeRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe/{rankingName}/{season}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteSubscribeRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteSubscribeRankingScoreByUserId(
	request *DeleteSubscribeRankingScoreByUserIdRequest,
) (*DeleteSubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan DeleteSubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.DeleteSubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySubscribeRankingScoreAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifySubscribeRankingScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySubscribeRankingScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySubscribeRankingScoreResult
	if asyncResult.Err != nil {
		callback <- VerifySubscribeRankingScoreAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySubscribeRankingScoreAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySubscribeRankingScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScoreAsync(
	request *VerifySubscribeRankingScoreRequest,
	callback chan<- VerifySubscribeRankingScoreAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe/{rankingName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifySubscribeRankingScoreAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScore(
	request *VerifySubscribeRankingScoreRequest,
) (*VerifySubscribeRankingScoreResult, error) {
	callback := make(chan VerifySubscribeRankingScoreAsyncResult, 1)
	go p.VerifySubscribeRankingScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySubscribeRankingScoreByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifySubscribeRankingScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySubscribeRankingScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySubscribeRankingScoreByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifySubscribeRankingScoreByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySubscribeRankingScoreByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySubscribeRankingScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScoreByUserIdAsync(
	request *VerifySubscribeRankingScoreByUserIdRequest,
	callback chan<- VerifySubscribeRankingScoreByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/score/subscribe/{rankingName}/verify/{verifyType}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Season != nil {
		bodies["season"] = *request.Season
	}
	if request.Score != nil {
		bodies["score"] = *request.Score
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

	go verifySubscribeRankingScoreByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScoreByUserId(
	request *VerifySubscribeRankingScoreByUserIdRequest,
) (*VerifySubscribeRankingScoreByUserIdResult, error) {
	callback := make(chan VerifySubscribeRankingScoreByUserIdAsyncResult, 1)
	go p.VerifySubscribeRankingScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifySubscribeRankingScoreByStampTaskAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifySubscribeRankingScoreByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifySubscribeRankingScoreByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifySubscribeRankingScoreByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifySubscribeRankingScoreByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifySubscribeRankingScoreByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifySubscribeRankingScoreByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScoreByStampTaskAsync(
	request *VerifySubscribeRankingScoreByStampTaskRequest,
	callback chan<- VerifySubscribeRankingScoreByStampTaskAsyncResult,
) {
	path := "/stamp/subscribe/score/verify"

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

	go verifySubscribeRankingScoreByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) VerifySubscribeRankingScoreByStampTask(
	request *VerifySubscribeRankingScoreByStampTaskRequest,
) (*VerifySubscribeRankingScoreByStampTaskResult, error) {
	callback := make(chan VerifySubscribeRankingScoreByStampTaskAsyncResult, 1)
	go p.VerifySubscribeRankingScoreByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingsAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingsResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingsAsync(
	request *DescribeSubscribeRankingsRequest,
	callback chan<- DescribeSubscribeRankingsAsyncResult,
) {
	path := "/{namespaceName}/ranking/subscribe/{rankingName}/user/me"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeSubscribeRankingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankings(
	request *DescribeSubscribeRankingsRequest,
) (*DescribeSubscribeRankingsResult, error) {
	callback := make(chan DescribeSubscribeRankingsAsyncResult, 1)
	go p.DescribeSubscribeRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscribeRankingsByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscribeRankingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribeRankingsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscribeRankingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingsByUserIdAsync(
	request *DescribeSubscribeRankingsByUserIdRequest,
	callback chan<- DescribeSubscribeRankingsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/subscribe/{rankingName}/user/{userId}"
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
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeSubscribeRankingsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DescribeSubscribeRankingsByUserId(
	request *DescribeSubscribeRankingsByUserIdRequest,
) (*DescribeSubscribeRankingsByUserIdResult, error) {
	callback := make(chan DescribeSubscribeRankingsByUserIdAsyncResult, 1)
	go p.DescribeSubscribeRankingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingAsync(
	request *GetSubscribeRankingRequest,
	callback chan<- GetSubscribeRankingAsyncResult,
) {
	path := "/{namespaceName}/ranking/subscribe/{rankingName}/user/me/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ScorerUserId != nil {
		queryStrings["scorerUserId"] = core.ToString(*request.ScorerUserId)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeRankingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRanking(
	request *GetSubscribeRankingRequest,
) (*GetSubscribeRankingResult, error) {
	callback := make(chan GetSubscribeRankingAsyncResult, 1)
	go p.GetSubscribeRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscribeRankingByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscribeRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeRankingByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSubscribeRankingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscribeRankingByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscribeRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) GetSubscribeRankingByUserIdAsync(
	request *GetSubscribeRankingByUserIdRequest,
	callback chan<- GetSubscribeRankingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/ranking/subscribe/{rankingName}/user/{userId}/rank"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Season != nil {
		queryStrings["season"] = core.ToString(*request.Season)
	}
	if request.ScorerUserId != nil {
		queryStrings["scorerUserId"] = core.ToString(*request.ScorerUserId)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeRankingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeRankingByUserId(
	request *GetSubscribeRankingByUserIdRequest,
) (*GetSubscribeRankingByUserIdResult, error) {
	callback := make(chan GetSubscribeRankingByUserIdAsyncResult, 1)
	go p.GetSubscribeRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2Ranking2RestClient,
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

func (p Gs2Ranking2RestClient) ExportMasterAsync(
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) ExportMaster(
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
	client Gs2Ranking2RestClient,
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
	if asyncResult.Err != nil {
		callback <- GetCurrentRankingMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2Ranking2RestClient) GetCurrentRankingMasterAsync(
	request *GetCurrentRankingMasterRequest,
	callback chan<- GetCurrentRankingMasterAsyncResult,
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCurrentRankingMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetCurrentRankingMaster(
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
	client Gs2Ranking2RestClient,
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
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRankingMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2Ranking2RestClient) UpdateCurrentRankingMasterAsync(
	request *UpdateCurrentRankingMasterRequest,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
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

	go updateCurrentRankingMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateCurrentRankingMaster(
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
	client Gs2Ranking2RestClient,
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
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2Ranking2RestClient) UpdateCurrentRankingMasterFromGitHubAsync(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
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

	go updateCurrentRankingMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) UpdateCurrentRankingMasterFromGitHub(
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

func getSubscribeAsyncHandler(
	client Gs2Ranking2RestClient,
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
	if asyncResult.Err != nil {
		callback <- GetSubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2Ranking2RestClient) GetSubscribeAsync(
	request *GetSubscribeRequest,
	callback chan<- GetSubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribe(
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
	client Gs2Ranking2RestClient,
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
	if asyncResult.Err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2Ranking2RestClient) GetSubscribeByUserIdAsync(
	request *GetSubscribeByUserIdRequest,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go getSubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) GetSubscribeByUserId(
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

func deleteSubscribeAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeResult
	if asyncResult.Err != nil {
		callback <- DeleteSubscribeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteSubscribeAsync(
	request *DeleteSubscribeRequest,
	callback chan<- DeleteSubscribeAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteSubscribeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteSubscribe(
	request *DeleteSubscribeRequest,
) (*DeleteSubscribeResult, error) {
	callback := make(chan DeleteSubscribeAsyncResult, 1)
	go p.DeleteSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSubscribeByUserIdAsyncHandler(
	client Gs2Ranking2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSubscribeByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteSubscribeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSubscribeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Ranking2RestClient) DeleteSubscribeByUserIdAsync(
	request *DeleteSubscribeByUserIdRequest,
	callback chan<- DeleteSubscribeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscribe/{rankingName}/target/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RankingName != nil && *request.RankingName != "" {
		path = strings.ReplaceAll(path, "{rankingName}", core.ToString(*request.RankingName))
	} else {
		path = strings.ReplaceAll(path, "{rankingName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go deleteSubscribeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("ranking2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Ranking2RestClient) DeleteSubscribeByUserId(
	request *DeleteSubscribeByUserIdRequest,
) (*DeleteSubscribeByUserIdResult, error) {
	callback := make(chan DeleteSubscribeByUserIdAsyncResult, 1)
	go p.DeleteSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
