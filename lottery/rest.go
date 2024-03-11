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

package lottery

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2LotteryRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2LotteryRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeNamespaces(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CreateNamespaceAsync(
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
	if request.LotteryTriggerScriptId != nil && *request.LotteryTriggerScriptId != "" {
		bodies["lotteryTriggerScriptId"] = *request.LotteryTriggerScriptId
	}
	if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
		bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
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

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreateNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetNamespaceStatus(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) UpdateNamespaceAsync(
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
	if request.LotteryTriggerScriptId != nil && *request.LotteryTriggerScriptId != "" {
		bodies["lotteryTriggerScriptId"] = *request.LotteryTriggerScriptId
	}
	if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
		bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
	}
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.QueueNamespaceId != nil && *request.QueueNamespaceId != "" {
		bodies["queueNamespaceId"] = *request.QueueNamespaceId
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

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeleteNamespace(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DumpUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CheckDumpUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CleanUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CheckCleanUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) PrepareImportUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ImportUserDataByUserId(
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
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CheckImportUserDataByUserId(
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

func describeLotteryModelMastersAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeLotteryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLotteryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeLotteryModelMastersAsync(
	request *DescribeLotteryModelMastersRequest,
	callback chan<- DescribeLotteryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/lottery"
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

	go describeLotteryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeLotteryModelMasters(
	request *DescribeLotteryModelMastersRequest,
) (*DescribeLotteryModelMastersResult, error) {
	callback := make(chan DescribeLotteryModelMastersAsyncResult, 1)
	go p.DescribeLotteryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateLotteryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) CreateLotteryModelMasterAsync(
	request *CreateLotteryModelMasterRequest,
	callback chan<- CreateLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery"
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
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		bodies["prizeTableName"] = *request.PrizeTableName
	}
	if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
		bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
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

	go createLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreateLotteryModelMaster(
	request *CreateLotteryModelMasterRequest,
) (*CreateLotteryModelMasterResult, error) {
	callback := make(chan CreateLotteryModelMasterAsyncResult, 1)
	go p.CreateLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetLotteryModelMasterAsync(
	request *GetLotteryModelMasterRequest,
	callback chan<- GetLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go getLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetLotteryModelMaster(
	request *GetLotteryModelMasterRequest,
) (*GetLotteryModelMasterResult, error) {
	callback := make(chan GetLotteryModelMasterAsyncResult, 1)
	go p.GetLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateLotteryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateLotteryModelMasterAsync(
	request *UpdateLotteryModelMasterRequest,
	callback chan<- UpdateLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Method != nil && *request.Method != "" {
		bodies["method"] = *request.Method
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		bodies["prizeTableName"] = *request.PrizeTableName
	}
	if request.ChoicePrizeTableScriptId != nil && *request.ChoicePrizeTableScriptId != "" {
		bodies["choicePrizeTableScriptId"] = *request.ChoicePrizeTableScriptId
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

	go updateLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateLotteryModelMaster(
	request *UpdateLotteryModelMasterRequest,
) (*UpdateLotteryModelMasterResult, error) {
	callback := make(chan UpdateLotteryModelMasterAsyncResult, 1)
	go p.UpdateLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLotteryModelMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLotteryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLotteryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLotteryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteLotteryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteLotteryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteLotteryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DeleteLotteryModelMasterAsync(
	request *DeleteLotteryModelMasterRequest,
	callback chan<- DeleteLotteryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/lottery/{lotteryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go deleteLotteryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeleteLotteryModelMaster(
	request *DeleteLotteryModelMasterRequest,
) (*DeleteLotteryModelMasterResult, error) {
	callback := make(chan DeleteLotteryModelMasterAsyncResult, 1)
	go p.DeleteLotteryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePrizeTableMastersAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePrizeTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePrizeTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePrizeTableMastersResult
	if asyncResult.Err != nil {
		callback <- DescribePrizeTableMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePrizeTableMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePrizeTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribePrizeTableMastersAsync(
	request *DescribePrizeTableMastersRequest,
	callback chan<- DescribePrizeTableMastersAsyncResult,
) {
	path := "/{namespaceName}/master/table"
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

	go describePrizeTableMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribePrizeTableMasters(
	request *DescribePrizeTableMastersRequest,
) (*DescribePrizeTableMastersResult, error) {
	callback := make(chan DescribePrizeTableMastersAsyncResult, 1)
	go p.DescribePrizeTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- CreatePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreatePrizeTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreatePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) CreatePrizeTableMasterAsync(
	request *CreatePrizeTableMasterRequest,
	callback chan<- CreatePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table"
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
	if request.Prizes != nil {
		var _prizes []interface{}
		for _, item := range request.Prizes {
			_prizes = append(_prizes, item)
		}
		bodies["prizes"] = _prizes
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

	go createPrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) CreatePrizeTableMaster(
	request *CreatePrizeTableMasterRequest,
) (*CreatePrizeTableMasterResult, error) {
	callback := make(chan CreatePrizeTableMasterAsyncResult, 1)
	go p.CreatePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetPrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- GetPrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPrizeTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetPrizeTableMasterAsync(
	request *GetPrizeTableMasterRequest,
	callback chan<- GetPrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go getPrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetPrizeTableMaster(
	request *GetPrizeTableMasterRequest,
) (*GetPrizeTableMasterResult, error) {
	callback := make(chan GetPrizeTableMasterAsyncResult, 1)
	go p.GetPrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updatePrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdatePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdatePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdatePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- UpdatePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdatePrizeTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdatePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdatePrizeTableMasterAsync(
	request *UpdatePrizeTableMasterRequest,
	callback chan<- UpdatePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Prizes != nil {
		var _prizes []interface{}
		for _, item := range request.Prizes {
			_prizes = append(_prizes, item)
		}
		bodies["prizes"] = _prizes
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

	go updatePrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdatePrizeTableMaster(
	request *UpdatePrizeTableMasterRequest,
) (*UpdatePrizeTableMasterResult, error) {
	callback := make(chan UpdatePrizeTableMasterAsyncResult, 1)
	go p.UpdatePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePrizeTableMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePrizeTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePrizeTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePrizeTableMasterResult
	if asyncResult.Err != nil {
		callback <- DeletePrizeTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePrizeTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePrizeTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DeletePrizeTableMasterAsync(
	request *DeletePrizeTableMasterRequest,
	callback chan<- DeletePrizeTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/table/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go deletePrizeTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DeletePrizeTableMaster(
	request *DeletePrizeTableMasterRequest,
) (*DeletePrizeTableMasterResult, error) {
	callback := make(chan DeletePrizeTableMasterAsyncResult, 1)
	go p.DeletePrizeTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLotteryModelsAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeLotteryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLotteryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeLotteryModelsAsync(
	request *DescribeLotteryModelsRequest,
	callback chan<- DescribeLotteryModelsAsyncResult,
) {
	path := "/{namespaceName}/lottery"
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

	go describeLotteryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeLotteryModels(
	request *DescribeLotteryModelsRequest,
) (*DescribeLotteryModelsResult, error) {
	callback := make(chan DescribeLotteryModelsAsyncResult, 1)
	go p.DescribeLotteryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLotteryModelAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryModelResult
	if asyncResult.Err != nil {
		callback <- GetLotteryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLotteryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetLotteryModelAsync(
	request *GetLotteryModelRequest,
	callback chan<- GetLotteryModelAsyncResult,
) {
	path := "/{namespaceName}/lottery/{lotteryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go getLotteryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetLotteryModel(
	request *GetLotteryModelRequest,
) (*GetLotteryModelResult, error) {
	callback := make(chan GetLotteryModelAsyncResult, 1)
	go p.GetLotteryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePrizeTablesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePrizeTablesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePrizeTablesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePrizeTablesResult
	if asyncResult.Err != nil {
		callback <- DescribePrizeTablesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePrizeTablesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePrizeTablesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribePrizeTablesAsync(
	request *DescribePrizeTablesRequest,
	callback chan<- DescribePrizeTablesAsyncResult,
) {
	path := "/{namespaceName}/table"
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

	go describePrizeTablesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribePrizeTables(
	request *DescribePrizeTablesRequest,
) (*DescribePrizeTablesResult, error) {
	callback := make(chan DescribePrizeTablesAsyncResult, 1)
	go p.DescribePrizeTablesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPrizeTableAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetPrizeTableAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPrizeTableAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPrizeTableResult
	if asyncResult.Err != nil {
		callback <- GetPrizeTableAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPrizeTableAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPrizeTableAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetPrizeTableAsync(
	request *GetPrizeTableRequest,
	callback chan<- GetPrizeTableAsyncResult,
) {
	path := "/{namespaceName}/table/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go getPrizeTableAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetPrizeTable(
	request *GetPrizeTableRequest,
) (*GetPrizeTableResult, error) {
	callback := make(chan GetPrizeTableAsyncResult, 1)
	go p.GetPrizeTableAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func drawByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DrawByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DrawByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DrawByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "box.items.empty" {
				asyncResult.Err = gs2err.SetClientError(Empty{})
			}
		}
		callback <- DrawByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DrawByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DrawByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DrawByUserIdAsync(
	request *DrawByUserIdRequest,
	callback chan<- DrawByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/draw"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Count != nil {
		bodies["count"] = *request.Count
	}
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

	go drawByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DrawByUserId(
	request *DrawByUserIdRequest,
) (*DrawByUserIdResult, error) {
	callback := make(chan DrawByUserIdAsyncResult, 1)
	go p.DrawByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func predictionAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- PredictionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PredictionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PredictionResult
	if asyncResult.Err != nil {
		callback <- PredictionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PredictionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PredictionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) PredictionAsync(
	request *PredictionRequest,
	callback chan<- PredictionAsyncResult,
) {
	path := "/{namespaceName}/user/me/lottery/{lotteryName}/prediction"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RandomSeed != nil {
		bodies["randomSeed"] = *request.RandomSeed
	}
	if request.Count != nil {
		bodies["count"] = *request.Count
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

	go predictionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) Prediction(
	request *PredictionRequest,
) (*PredictionResult, error) {
	callback := make(chan PredictionAsyncResult, 1)
	go p.PredictionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func predictionByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- PredictionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PredictionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PredictionByUserIdResult
	if asyncResult.Err != nil {
		callback <- PredictionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PredictionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PredictionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) PredictionByUserIdAsync(
	request *PredictionByUserIdRequest,
	callback chan<- PredictionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/prediction"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RandomSeed != nil {
		bodies["randomSeed"] = *request.RandomSeed
	}
	if request.Count != nil {
		bodies["count"] = *request.Count
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

	go predictionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) PredictionByUserId(
	request *PredictionByUserIdRequest,
) (*PredictionByUserIdResult, error) {
	callback := make(chan PredictionByUserIdAsyncResult, 1)
	go p.PredictionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func drawWithRandomSeedByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DrawWithRandomSeedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DrawWithRandomSeedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DrawWithRandomSeedByUserIdResult
	if asyncResult.Err != nil {
		callback <- DrawWithRandomSeedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DrawWithRandomSeedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DrawWithRandomSeedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DrawWithRandomSeedByUserIdAsync(
	request *DrawWithRandomSeedByUserIdRequest,
	callback chan<- DrawWithRandomSeedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/draw/withSeed"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RandomSeed != nil {
		bodies["randomSeed"] = *request.RandomSeed
	}
	if request.Count != nil {
		bodies["count"] = *request.Count
	}
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

	go drawWithRandomSeedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DrawWithRandomSeedByUserId(
	request *DrawWithRandomSeedByUserIdRequest,
) (*DrawWithRandomSeedByUserIdResult, error) {
	callback := make(chan DrawWithRandomSeedByUserIdAsyncResult, 1)
	go p.DrawWithRandomSeedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func drawByStampSheetAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DrawByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DrawByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DrawByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DrawByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DrawByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DrawByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DrawByStampSheetAsync(
	request *DrawByStampSheetRequest,
	callback chan<- DrawByStampSheetAsyncResult,
) {
	path := "/stamp/draw"

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

	go drawByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DrawByStampSheet(
	request *DrawByStampSheetRequest,
) (*DrawByStampSheetResult, error) {
	callback := make(chan DrawByStampSheetAsyncResult, 1)
	go p.DrawByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeProbabilitiesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProbabilitiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProbabilitiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProbabilitiesResult
	if asyncResult.Err != nil {
		callback <- DescribeProbabilitiesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeProbabilitiesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeProbabilitiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeProbabilitiesAsync(
	request *DescribeProbabilitiesRequest,
	callback chan<- DescribeProbabilitiesAsyncResult,
) {
	path := "/{namespaceName}/user/me/lottery/{lotteryName}/probability"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
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

	go describeProbabilitiesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeProbabilities(
	request *DescribeProbabilitiesRequest,
) (*DescribeProbabilitiesResult, error) {
	callback := make(chan DescribeProbabilitiesAsyncResult, 1)
	go p.DescribeProbabilitiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeProbabilitiesByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProbabilitiesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProbabilitiesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProbabilitiesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeProbabilitiesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeProbabilitiesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeProbabilitiesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeProbabilitiesByUserIdAsync(
	request *DescribeProbabilitiesByUserIdRequest,
	callback chan<- DescribeProbabilitiesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/lottery/{lotteryName}/probability"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeProbabilitiesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeProbabilitiesByUserId(
	request *DescribeProbabilitiesByUserIdRequest,
) (*DescribeProbabilitiesByUserIdResult, error) {
	callback := make(chan DescribeProbabilitiesByUserIdAsyncResult, 1)
	go p.DescribeProbabilitiesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2LotteryRestClient,
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

func (p Gs2LotteryRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ExportMaster(
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

func getCurrentLotteryMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentLotteryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentLotteryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentLotteryMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentLotteryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentLotteryMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentLotteryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetCurrentLotteryMasterAsync(
	request *GetCurrentLotteryMasterRequest,
	callback chan<- GetCurrentLotteryMasterAsyncResult,
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

	go getCurrentLotteryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetCurrentLotteryMaster(
	request *GetCurrentLotteryMasterRequest,
) (*GetCurrentLotteryMasterResult, error) {
	callback := make(chan GetCurrentLotteryMasterAsyncResult, 1)
	go p.GetCurrentLotteryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLotteryMasterAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLotteryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLotteryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLotteryMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLotteryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentLotteryMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentLotteryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterAsync(
	request *UpdateCurrentLotteryMasterRequest,
	callback chan<- UpdateCurrentLotteryMasterAsyncResult,
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

	go updateCurrentLotteryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMaster(
	request *UpdateCurrentLotteryMasterRequest,
) (*UpdateCurrentLotteryMasterResult, error) {
	callback := make(chan UpdateCurrentLotteryMasterAsyncResult, 1)
	go p.UpdateCurrentLotteryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentLotteryMasterFromGitHubAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentLotteryMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentLotteryMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentLotteryMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterFromGitHubAsync(
	request *UpdateCurrentLotteryMasterFromGitHubRequest,
	callback chan<- UpdateCurrentLotteryMasterFromGitHubAsyncResult,
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

	go updateCurrentLotteryMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) UpdateCurrentLotteryMasterFromGitHub(
	request *UpdateCurrentLotteryMasterFromGitHubRequest,
) (*UpdateCurrentLotteryMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentLotteryMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentLotteryMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePrizeLimitsAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePrizeLimitsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePrizeLimitsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePrizeLimitsResult
	if asyncResult.Err != nil {
		callback <- DescribePrizeLimitsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePrizeLimitsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePrizeLimitsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribePrizeLimitsAsync(
	request *DescribePrizeLimitsRequest,
	callback chan<- DescribePrizeLimitsAsyncResult,
) {
	path := "/{namespaceName}/prizeLimit/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go describePrizeLimitsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribePrizeLimits(
	request *DescribePrizeLimitsRequest,
) (*DescribePrizeLimitsResult, error) {
	callback := make(chan DescribePrizeLimitsAsyncResult, 1)
	go p.DescribePrizeLimitsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPrizeLimitAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetPrizeLimitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPrizeLimitAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPrizeLimitResult
	if asyncResult.Err != nil {
		callback <- GetPrizeLimitAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPrizeLimitAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPrizeLimitAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetPrizeLimitAsync(
	request *GetPrizeLimitRequest,
	callback chan<- GetPrizeLimitAsyncResult,
) {
	path := "/{namespaceName}/prizeLimit/{prizeTableName}/{prizeId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
	}
	if request.PrizeId != nil && *request.PrizeId != "" {
		path = strings.ReplaceAll(path, "{prizeId}", core.ToString(*request.PrizeId))
	} else {
		path = strings.ReplaceAll(path, "{prizeId}", "null")
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

	go getPrizeLimitAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetPrizeLimit(
	request *GetPrizeLimitRequest,
) (*GetPrizeLimitResult, error) {
	callback := make(chan GetPrizeLimitAsyncResult, 1)
	go p.GetPrizeLimitAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetPrizeLimitAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetPrizeLimitAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetPrizeLimitAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetPrizeLimitResult
	if asyncResult.Err != nil {
		callback <- ResetPrizeLimitAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetPrizeLimitAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetPrizeLimitAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetPrizeLimitAsync(
	request *ResetPrizeLimitRequest,
	callback chan<- ResetPrizeLimitAsyncResult,
) {
	path := "/{namespaceName}/prizeLimit/{prizeTableName}/{prizeId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
	}
	if request.PrizeId != nil && *request.PrizeId != "" {
		path = strings.ReplaceAll(path, "{prizeId}", core.ToString(*request.PrizeId))
	} else {
		path = strings.ReplaceAll(path, "{prizeId}", "null")
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

	go resetPrizeLimitAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetPrizeLimit(
	request *ResetPrizeLimitRequest,
) (*ResetPrizeLimitResult, error) {
	callback := make(chan ResetPrizeLimitAsyncResult, 1)
	go p.ResetPrizeLimitAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBoxesAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBoxesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBoxesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBoxesResult
	if asyncResult.Err != nil {
		callback <- DescribeBoxesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBoxesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBoxesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeBoxesAsync(
	request *DescribeBoxesRequest,
	callback chan<- DescribeBoxesAsyncResult,
) {
	path := "/{namespaceName}/user/me/box"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeBoxesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeBoxes(
	request *DescribeBoxesRequest,
) (*DescribeBoxesResult, error) {
	callback := make(chan DescribeBoxesAsyncResult, 1)
	go p.DescribeBoxesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBoxesByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBoxesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBoxesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBoxesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeBoxesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBoxesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBoxesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) DescribeBoxesByUserIdAsync(
	request *DescribeBoxesByUserIdRequest,
	callback chan<- DescribeBoxesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describeBoxesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) DescribeBoxesByUserId(
	request *DescribeBoxesByUserIdRequest,
) (*DescribeBoxesByUserIdResult, error) {
	callback := make(chan DescribeBoxesByUserIdAsyncResult, 1)
	go p.DescribeBoxesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBoxAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBoxAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBoxAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBoxResult
	if asyncResult.Err != nil {
		callback <- GetBoxAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBoxAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBoxAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetBoxAsync(
	request *GetBoxRequest,
	callback chan<- GetBoxAsyncResult,
) {
	path := "/{namespaceName}/user/me/box/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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

	go getBoxAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetBox(
	request *GetBoxRequest,
) (*GetBoxResult, error) {
	callback := make(chan GetBoxAsyncResult, 1)
	go p.GetBoxAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBoxByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- GetBoxByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBoxByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBoxByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetBoxByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBoxByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBoxByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) GetBoxByUserIdAsync(
	request *GetBoxByUserIdRequest,
	callback chan<- GetBoxByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
	}
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go getBoxByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) GetBoxByUserId(
	request *GetBoxByUserIdRequest,
) (*GetBoxByUserIdResult, error) {
	callback := make(chan GetBoxByUserIdAsyncResult, 1)
	go p.GetBoxByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetBoxAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetBoxAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetBoxAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetBoxResult
	if asyncResult.Err != nil {
		callback <- ResetBoxAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetBoxAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetBoxAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetBoxAsync(
	request *ResetBoxRequest,
	callback chan<- ResetBoxAsyncResult,
) {
	path := "/{namespaceName}/user/me/box/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go resetBoxAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetBox(
	request *ResetBoxRequest,
) (*ResetBoxResult, error) {
	callback := make(chan ResetBoxAsyncResult, 1)
	go p.ResetBoxAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetBoxByUserIdAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetBoxByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetBoxByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetBoxByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetBoxByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetBoxByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetBoxByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetBoxByUserIdAsync(
	request *ResetBoxByUserIdRequest,
	callback chan<- ResetBoxByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/box/{prizeTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PrizeTableName != nil && *request.PrizeTableName != "" {
		path = strings.ReplaceAll(path, "{prizeTableName}", core.ToString(*request.PrizeTableName))
	} else {
		path = strings.ReplaceAll(path, "{prizeTableName}", "null")
	}
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

	go resetBoxByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetBoxByUserId(
	request *ResetBoxByUserIdRequest,
) (*ResetBoxByUserIdResult, error) {
	callback := make(chan ResetBoxByUserIdAsyncResult, 1)
	go p.ResetBoxByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetByStampSheetAsyncHandler(
	client Gs2LotteryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetByStampSheetResult
	if asyncResult.Err != nil {
		callback <- ResetByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2LotteryRestClient) ResetByStampSheetAsync(
	request *ResetByStampSheetRequest,
	callback chan<- ResetByStampSheetAsyncResult,
) {
	path := "/stamp/box/reset"

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

	go resetByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("lottery").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2LotteryRestClient) ResetByStampSheet(
	request *ResetByStampSheetRequest,
) (*ResetByStampSheetResult, error) {
	callback := make(chan ResetByStampSheetAsyncResult, 1)
	go p.ResetByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
