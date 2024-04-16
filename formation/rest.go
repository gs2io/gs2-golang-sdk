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

package formation

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2FormationRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2FormationRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeNamespaces(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) CreateNamespaceAsync(
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
	if request.UpdateMoldScript != nil {
		bodies["updateMoldScript"] = request.UpdateMoldScript.ToDict()
	}
	if request.UpdateFormScript != nil {
		bodies["updateFormScript"] = request.UpdateFormScript.ToDict()
	}
	if request.UpdatePropertyFormScript != nil {
		bodies["updatePropertyFormScript"] = request.UpdatePropertyFormScript.ToDict()
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CreateNamespace(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetNamespaceStatus(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetNamespace(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) UpdateNamespaceAsync(
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
	if request.UpdateMoldScript != nil {
		bodies["updateMoldScript"] = request.UpdateMoldScript.ToDict()
	}
	if request.UpdateFormScript != nil {
		bodies["updateFormScript"] = request.UpdateFormScript.ToDict()
	}
	if request.UpdatePropertyFormScript != nil {
		bodies["updatePropertyFormScript"] = request.UpdatePropertyFormScript.ToDict()
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdateNamespace(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteNamespace(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DumpUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) CheckDumpUserDataByUserIdAsync(
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CheckDumpUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CleanUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) CheckCleanUserDataByUserIdAsync(
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CheckCleanUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) PrepareImportUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) ImportUserDataByUserId(
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
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) CheckImportUserDataByUserIdAsync(
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
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CheckImportUserDataByUserId(
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

func getFormModelAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormModelResult
	if asyncResult.Err != nil {
		callback <- GetFormModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormModelAsync(
	request *GetFormModelRequest,
	callback chan<- GetFormModelAsyncResult,
) {
	path := "/{namespaceName}/model/{moldModelName}/form"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go getFormModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetFormModel(
	request *GetFormModelRequest,
) (*GetFormModelResult, error) {
	callback := make(chan GetFormModelAsyncResult, 1)
	go p.GetFormModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeFormModelMastersAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeFormModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeFormModelMastersAsync(
	request *DescribeFormModelMastersRequest,
	callback chan<- DescribeFormModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/form"
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

	go describeFormModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeFormModelMasters(
	request *DescribeFormModelMastersRequest,
) (*DescribeFormModelMastersResult, error) {
	callback := make(chan DescribeFormModelMastersAsyncResult, 1)
	go p.DescribeFormModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- CreateFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) CreateFormModelMasterAsync(
	request *CreateFormModelMasterRequest,
	callback chan<- CreateFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/form"
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
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go createFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CreateFormModelMaster(
	request *CreateFormModelMasterRequest,
) (*CreateFormModelMasterResult, error) {
	callback := make(chan CreateFormModelMasterAsyncResult, 1)
	go p.CreateFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormModelMasterAsync(
	request *GetFormModelMasterRequest,
	callback chan<- GetFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/form/{formModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.FormModelName != nil && *request.FormModelName != "" {
		path = strings.ReplaceAll(path, "{formModelName}", core.ToString(*request.FormModelName))
	} else {
		path = strings.ReplaceAll(path, "{formModelName}", "null")
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

	go getFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetFormModelMaster(
	request *GetFormModelMasterRequest,
) (*GetFormModelMasterResult, error) {
	callback := make(chan GetFormModelMasterAsyncResult, 1)
	go p.GetFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) UpdateFormModelMasterAsync(
	request *UpdateFormModelMasterRequest,
	callback chan<- UpdateFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/form/{formModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.FormModelName != nil && *request.FormModelName != "" {
		path = strings.ReplaceAll(path, "{formModelName}", core.ToString(*request.FormModelName))
	} else {
		path = strings.ReplaceAll(path, "{formModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go updateFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdateFormModelMaster(
	request *UpdateFormModelMasterRequest,
) (*UpdateFormModelMasterResult, error) {
	callback := make(chan UpdateFormModelMasterAsyncResult, 1)
	go p.UpdateFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteFormModelMasterAsync(
	request *DeleteFormModelMasterRequest,
	callback chan<- DeleteFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/form/{formModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.FormModelName != nil && *request.FormModelName != "" {
		path = strings.ReplaceAll(path, "{formModelName}", core.ToString(*request.FormModelName))
	} else {
		path = strings.ReplaceAll(path, "{formModelName}", "null")
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

	go deleteFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteFormModelMaster(
	request *DeleteFormModelMasterRequest,
) (*DeleteFormModelMasterResult, error) {
	callback := make(chan DeleteFormModelMasterAsyncResult, 1)
	go p.DeleteFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMoldModelsAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoldModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeMoldModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoldModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoldModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeMoldModelsAsync(
	request *DescribeMoldModelsRequest,
	callback chan<- DescribeMoldModelsAsyncResult,
) {
	path := "/{namespaceName}/model/mold"
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

	go describeMoldModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeMoldModels(
	request *DescribeMoldModelsRequest,
) (*DescribeMoldModelsResult, error) {
	callback := make(chan DescribeMoldModelsAsyncResult, 1)
	go p.DescribeMoldModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMoldModelAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetMoldModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldModelResult
	if asyncResult.Err != nil {
		callback <- GetMoldModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoldModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMoldModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetMoldModelAsync(
	request *GetMoldModelRequest,
	callback chan<- GetMoldModelAsyncResult,
) {
	path := "/{namespaceName}/model/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go getMoldModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetMoldModel(
	request *GetMoldModelRequest,
) (*GetMoldModelResult, error) {
	callback := make(chan GetMoldModelAsyncResult, 1)
	go p.GetMoldModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMoldModelMastersAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoldModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeMoldModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoldModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoldModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeMoldModelMastersAsync(
	request *DescribeMoldModelMastersRequest,
	callback chan<- DescribeMoldModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/mold"
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

	go describeMoldModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeMoldModelMasters(
	request *DescribeMoldModelMastersRequest,
) (*DescribeMoldModelMastersResult, error) {
	callback := make(chan DescribeMoldModelMastersAsyncResult, 1)
	go p.DescribeMoldModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createMoldModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- CreateMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMoldModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateMoldModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMoldModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) CreateMoldModelMasterAsync(
	request *CreateMoldModelMasterRequest,
	callback chan<- CreateMoldModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/mold"
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
	if request.FormModelName != nil && *request.FormModelName != "" {
		bodies["formModelName"] = *request.FormModelName
	}
	if request.InitialMaxCapacity != nil {
		bodies["initialMaxCapacity"] = *request.InitialMaxCapacity
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
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

	go createMoldModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CreateMoldModelMaster(
	request *CreateMoldModelMasterRequest,
) (*CreateMoldModelMasterResult, error) {
	callback := make(chan CreateMoldModelMasterAsyncResult, 1)
	go p.CreateMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMoldModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetMoldModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoldModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetMoldModelMasterAsync(
	request *GetMoldModelMasterRequest,
	callback chan<- GetMoldModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go getMoldModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetMoldModelMaster(
	request *GetMoldModelMasterRequest,
) (*GetMoldModelMasterResult, error) {
	callback := make(chan GetMoldModelMasterAsyncResult, 1)
	go p.GetMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMoldModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMoldModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateMoldModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMoldModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) UpdateMoldModelMasterAsync(
	request *UpdateMoldModelMasterRequest,
	callback chan<- UpdateMoldModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.FormModelName != nil && *request.FormModelName != "" {
		bodies["formModelName"] = *request.FormModelName
	}
	if request.InitialMaxCapacity != nil {
		bodies["initialMaxCapacity"] = *request.InitialMaxCapacity
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
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

	go updateMoldModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdateMoldModelMaster(
	request *UpdateMoldModelMasterRequest,
) (*UpdateMoldModelMasterResult, error) {
	callback := make(chan UpdateMoldModelMasterAsyncResult, 1)
	go p.UpdateMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMoldModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMoldModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteMoldModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMoldModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMoldModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteMoldModelMasterAsync(
	request *DeleteMoldModelMasterRequest,
	callback chan<- DeleteMoldModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go deleteMoldModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteMoldModelMaster(
	request *DeleteMoldModelMasterRequest,
) (*DeleteMoldModelMasterResult, error) {
	callback := make(chan DeleteMoldModelMasterAsyncResult, 1)
	go p.DeleteMoldModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePropertyFormModelsAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePropertyFormModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePropertyFormModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePropertyFormModelsResult
	if asyncResult.Err != nil {
		callback <- DescribePropertyFormModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePropertyFormModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePropertyFormModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribePropertyFormModelsAsync(
	request *DescribePropertyFormModelsRequest,
	callback chan<- DescribePropertyFormModelsAsyncResult,
) {
	path := "/{namespaceName}/model/propertyForm"
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

	go describePropertyFormModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribePropertyFormModels(
	request *DescribePropertyFormModelsRequest,
) (*DescribePropertyFormModelsResult, error) {
	callback := make(chan DescribePropertyFormModelsAsyncResult, 1)
	go p.DescribePropertyFormModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormModelAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormModelResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormModelAsync(
	request *GetPropertyFormModelRequest,
	callback chan<- GetPropertyFormModelAsyncResult,
) {
	path := "/{namespaceName}/model/propertyForm/{propertyFormModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
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

	go getPropertyFormModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyFormModel(
	request *GetPropertyFormModelRequest,
) (*GetPropertyFormModelResult, error) {
	callback := make(chan GetPropertyFormModelAsyncResult, 1)
	go p.GetPropertyFormModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePropertyFormModelMastersAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePropertyFormModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePropertyFormModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePropertyFormModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribePropertyFormModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePropertyFormModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePropertyFormModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribePropertyFormModelMastersAsync(
	request *DescribePropertyFormModelMastersRequest,
	callback chan<- DescribePropertyFormModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/propertyForm"
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

	go describePropertyFormModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribePropertyFormModelMasters(
	request *DescribePropertyFormModelMastersRequest,
) (*DescribePropertyFormModelMastersResult, error) {
	callback := make(chan DescribePropertyFormModelMastersAsyncResult, 1)
	go p.DescribePropertyFormModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPropertyFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePropertyFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePropertyFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePropertyFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreatePropertyFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreatePropertyFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreatePropertyFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) CreatePropertyFormModelMasterAsync(
	request *CreatePropertyFormModelMasterRequest,
	callback chan<- CreatePropertyFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/propertyForm"
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
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go createPropertyFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) CreatePropertyFormModelMaster(
	request *CreatePropertyFormModelMasterRequest,
) (*CreatePropertyFormModelMasterResult, error) {
	callback := make(chan CreatePropertyFormModelMasterAsyncResult, 1)
	go p.CreatePropertyFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormModelMasterAsync(
	request *GetPropertyFormModelMasterRequest,
	callback chan<- GetPropertyFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/propertyForm/{propertyFormModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
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

	go getPropertyFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyFormModelMaster(
	request *GetPropertyFormModelMasterRequest,
) (*GetPropertyFormModelMasterResult, error) {
	callback := make(chan GetPropertyFormModelMasterAsyncResult, 1)
	go p.GetPropertyFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updatePropertyFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- UpdatePropertyFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdatePropertyFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdatePropertyFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdatePropertyFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdatePropertyFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdatePropertyFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) UpdatePropertyFormModelMasterAsync(
	request *UpdatePropertyFormModelMasterRequest,
	callback chan<- UpdatePropertyFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/propertyForm/{propertyFormModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go updatePropertyFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdatePropertyFormModelMaster(
	request *UpdatePropertyFormModelMasterRequest,
) (*UpdatePropertyFormModelMasterResult, error) {
	callback := make(chan UpdatePropertyFormModelMasterAsyncResult, 1)
	go p.UpdatePropertyFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePropertyFormModelMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePropertyFormModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePropertyFormModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePropertyFormModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeletePropertyFormModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePropertyFormModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePropertyFormModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeletePropertyFormModelMasterAsync(
	request *DeletePropertyFormModelMasterRequest,
	callback chan<- DeletePropertyFormModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/propertyForm/{propertyFormModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
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

	go deletePropertyFormModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeletePropertyFormModelMaster(
	request *DeletePropertyFormModelMasterRequest,
) (*DeletePropertyFormModelMasterResult, error) {
	callback := make(chan DeletePropertyFormModelMasterAsyncResult, 1)
	go p.DeletePropertyFormModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) ExportMaster(
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

func getCurrentFormMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentFormMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentFormMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentFormMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentFormMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentFormMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentFormMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetCurrentFormMasterAsync(
	request *GetCurrentFormMasterRequest,
	callback chan<- GetCurrentFormMasterAsyncResult,
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

	go getCurrentFormMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetCurrentFormMaster(
	request *GetCurrentFormMasterRequest,
) (*GetCurrentFormMasterResult, error) {
	callback := make(chan GetCurrentFormMasterAsyncResult, 1)
	go p.GetCurrentFormMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentFormMasterAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentFormMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFormMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFormMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentFormMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentFormMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentFormMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) UpdateCurrentFormMasterAsync(
	request *UpdateCurrentFormMasterRequest,
	callback chan<- UpdateCurrentFormMasterAsyncResult,
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

	go updateCurrentFormMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdateCurrentFormMaster(
	request *UpdateCurrentFormMasterRequest,
) (*UpdateCurrentFormMasterResult, error) {
	callback := make(chan UpdateCurrentFormMasterAsyncResult, 1)
	go p.UpdateCurrentFormMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentFormMasterFromGitHubAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentFormMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentFormMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentFormMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) UpdateCurrentFormMasterFromGitHubAsync(
	request *UpdateCurrentFormMasterFromGitHubRequest,
	callback chan<- UpdateCurrentFormMasterFromGitHubAsyncResult,
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

	go updateCurrentFormMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) UpdateCurrentFormMasterFromGitHub(
	request *UpdateCurrentFormMasterFromGitHubRequest,
) (*UpdateCurrentFormMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentFormMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentFormMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMoldsAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoldsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldsResult
	if asyncResult.Err != nil {
		callback <- DescribeMoldsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoldsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoldsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeMoldsAsync(
	request *DescribeMoldsRequest,
	callback chan<- DescribeMoldsAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeMoldsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeMolds(
	request *DescribeMoldsRequest,
) (*DescribeMoldsResult, error) {
	callback := make(chan DescribeMoldsAsyncResult, 1)
	go p.DescribeMoldsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMoldsByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoldsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoldsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoldsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeMoldsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoldsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoldsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeMoldsByUserIdAsync(
	request *DescribeMoldsByUserIdRequest,
	callback chan<- DescribeMoldsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold"
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

	go describeMoldsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeMoldsByUserId(
	request *DescribeMoldsByUserIdRequest,
) (*DescribeMoldsByUserIdResult, error) {
	callback := make(chan DescribeMoldsByUserIdAsyncResult, 1)
	go p.DescribeMoldsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMoldAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetMoldAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldResult
	if asyncResult.Err != nil {
		callback <- GetMoldAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoldAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMoldAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetMoldAsync(
	request *GetMoldRequest,
	callback chan<- GetMoldAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go getMoldAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetMold(
	request *GetMoldRequest,
) (*GetMoldResult, error) {
	callback := make(chan GetMoldAsyncResult, 1)
	go p.GetMoldAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMoldByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetMoldByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoldByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoldByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetMoldByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoldByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMoldByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetMoldByUserIdAsync(
	request *GetMoldByUserIdRequest,
	callback chan<- GetMoldByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go getMoldByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetMoldByUserId(
	request *GetMoldByUserIdRequest,
) (*GetMoldByUserIdResult, error) {
	callback := make(chan GetMoldByUserIdAsyncResult, 1)
	go p.GetMoldByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMoldCapacityByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetMoldCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMoldCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMoldCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetMoldCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMoldCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMoldCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetMoldCapacityByUserIdAsync(
	request *SetMoldCapacityByUserIdRequest,
	callback chan<- SetMoldCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Capacity != nil {
		bodies["capacity"] = *request.Capacity
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

	go setMoldCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetMoldCapacityByUserId(
	request *SetMoldCapacityByUserIdRequest,
) (*SetMoldCapacityByUserIdResult, error) {
	callback := make(chan SetMoldCapacityByUserIdAsyncResult, 1)
	go p.SetMoldCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addMoldCapacityByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- AddMoldCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddMoldCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddMoldCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddMoldCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddMoldCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddMoldCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) AddMoldCapacityByUserIdAsync(
	request *AddMoldCapacityByUserIdRequest,
	callback chan<- AddMoldCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Capacity != nil {
		bodies["capacity"] = *request.Capacity
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

	go addMoldCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AddMoldCapacityByUserId(
	request *AddMoldCapacityByUserIdRequest,
) (*AddMoldCapacityByUserIdResult, error) {
	callback := make(chan AddMoldCapacityByUserIdAsyncResult, 1)
	go p.AddMoldCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func subMoldCapacityByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SubMoldCapacityByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubMoldCapacityByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubMoldCapacityByUserIdResult
	if asyncResult.Err != nil {
		callback <- SubMoldCapacityByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubMoldCapacityByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubMoldCapacityByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SubMoldCapacityByUserIdAsync(
	request *SubMoldCapacityByUserIdRequest,
	callback chan<- SubMoldCapacityByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/sub"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Capacity != nil {
		bodies["capacity"] = *request.Capacity
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

	go subMoldCapacityByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SubMoldCapacityByUserId(
	request *SubMoldCapacityByUserIdRequest,
) (*SubMoldCapacityByUserIdResult, error) {
	callback := make(chan SubMoldCapacityByUserIdAsyncResult, 1)
	go p.SubMoldCapacityByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMoldAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMoldAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldResult
	if asyncResult.Err != nil {
		callback <- DeleteMoldAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMoldAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMoldAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteMoldAsync(
	request *DeleteMoldRequest,
	callback chan<- DeleteMoldAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go deleteMoldAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteMold(
	request *DeleteMoldRequest,
) (*DeleteMoldResult, error) {
	callback := make(chan DeleteMoldAsyncResult, 1)
	go p.DeleteMoldAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMoldByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMoldByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMoldByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMoldByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteMoldByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMoldByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMoldByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteMoldByUserIdAsync(
	request *DeleteMoldByUserIdRequest,
	callback chan<- DeleteMoldByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go deleteMoldByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteMoldByUserId(
	request *DeleteMoldByUserIdRequest,
) (*DeleteMoldByUserIdResult, error) {
	callback := make(chan DeleteMoldByUserIdAsyncResult, 1)
	go p.DeleteMoldByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addCapacityByStampSheetAsyncHandler(
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) AddCapacityByStampSheetAsync(
	request *AddCapacityByStampSheetRequest,
	callback chan<- AddCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/mold/capacity/add"

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

	go addCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AddCapacityByStampSheet(
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

func subCapacityByStampTaskAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SubCapacityByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubCapacityByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubCapacityByStampTaskResult
	if asyncResult.Err != nil {
		callback <- SubCapacityByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SubCapacityByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SubCapacityByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SubCapacityByStampTaskAsync(
	request *SubCapacityByStampTaskRequest,
	callback chan<- SubCapacityByStampTaskAsyncResult,
) {
	path := "/stamp/mold/capacity/sub"

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

	go subCapacityByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SubCapacityByStampTask(
	request *SubCapacityByStampTaskRequest,
) (*SubCapacityByStampTaskResult, error) {
	callback := make(chan SubCapacityByStampTaskAsyncResult, 1)
	go p.SubCapacityByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setCapacityByStampSheetAsyncHandler(
	client Gs2FormationRestClient,
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

func (p Gs2FormationRestClient) SetCapacityByStampSheetAsync(
	request *SetCapacityByStampSheetRequest,
	callback chan<- SetCapacityByStampSheetAsyncResult,
) {
	path := "/stamp/mold/capacity/set"

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

	go setCapacityByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetCapacityByStampSheet(
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

func describeFormsAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormsResult
	if asyncResult.Err != nil {
		callback <- DescribeFormsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeFormsAsync(
	request *DescribeFormsRequest,
	callback chan<- DescribeFormsAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}/form"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describeFormsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeForms(
	request *DescribeFormsRequest,
) (*DescribeFormsResult, error) {
	callback := make(chan DescribeFormsAsyncResult, 1)
	go p.DescribeFormsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeFormsByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeFormsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribeFormsByUserIdAsync(
	request *DescribeFormsByUserIdRequest,
	callback chan<- DescribeFormsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
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

	go describeFormsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribeFormsByUserId(
	request *DescribeFormsByUserIdRequest,
) (*DescribeFormsByUserIdResult, error) {
	callback := make(chan DescribeFormsByUserIdAsyncResult, 1)
	go p.DescribeFormsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFormAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormResult
	if asyncResult.Err != nil {
		callback <- GetFormAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormAsync(
	request *GetFormRequest,
	callback chan<- GetFormAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}/form/{index}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
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

	go getFormAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetForm(
	request *GetFormRequest,
) (*GetFormResult, error) {
	callback := make(chan GetFormAsyncResult, 1)
	go p.GetFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormByUserIdAsync(
	request *GetFormByUserIdRequest,
	callback chan<- GetFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form/{index}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
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

	go getFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetFormByUserId(
	request *GetFormByUserIdRequest,
) (*GetFormByUserIdResult, error) {
	callback := make(chan GetFormByUserIdAsyncResult, 1)
	go p.GetFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFormWithSignatureAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetFormWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormWithSignatureAsync(
	request *GetFormWithSignatureRequest,
	callback chan<- GetFormWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}/form/{index}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
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

	go getFormWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetFormWithSignature(
	request *GetFormWithSignatureRequest,
) (*GetFormWithSignatureResult, error) {
	callback := make(chan GetFormWithSignatureAsyncResult, 1)
	go p.GetFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFormWithSignatureByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetFormWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetFormWithSignatureByUserIdAsync(
	request *GetFormWithSignatureByUserIdRequest,
	callback chan<- GetFormWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form/{index}/signature"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
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

	go getFormWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetFormWithSignatureByUserId(
	request *GetFormWithSignatureByUserIdRequest,
) (*GetFormWithSignatureByUserIdResult, error) {
	callback := make(chan GetFormWithSignatureByUserIdAsyncResult, 1)
	go p.GetFormWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetFormByUserIdAsync(
	request *SetFormByUserIdRequest,
	callback chan<- SetFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form/{index}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go setFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetFormByUserId(
	request *SetFormByUserIdRequest,
) (*SetFormByUserIdResult, error) {
	callback := make(chan SetFormByUserIdAsyncResult, 1)
	go p.SetFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setFormWithSignatureAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFormWithSignatureResult
	if asyncResult.Err != nil {
		callback <- SetFormWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetFormWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetFormWithSignatureAsync(
	request *SetFormWithSignatureRequest,
	callback chan<- SetFormWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}/form/{index}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setFormWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetFormWithSignature(
	request *SetFormWithSignatureRequest,
) (*SetFormWithSignatureResult, error) {
	callback := make(chan SetFormWithSignatureAsyncResult, 1)
	go p.SetFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireActionsToFormPropertiesAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireActionsToFormPropertiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionsToFormPropertiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionsToFormPropertiesResult
	if asyncResult.Err != nil {
		callback <- AcquireActionsToFormPropertiesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireActionsToFormPropertiesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireActionsToFormPropertiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) AcquireActionsToFormPropertiesAsync(
	request *AcquireActionsToFormPropertiesRequest,
	callback chan<- AcquireActionsToFormPropertiesAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form/{index}/stamp/delegate"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireAction != nil {
		bodies["acquireAction"] = request.AcquireAction.ToDict()
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

	go acquireActionsToFormPropertiesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AcquireActionsToFormProperties(
	request *AcquireActionsToFormPropertiesRequest,
) (*AcquireActionsToFormPropertiesResult, error) {
	callback := make(chan AcquireActionsToFormPropertiesAsyncResult, 1)
	go p.AcquireActionsToFormPropertiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteFormAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormResult
	if asyncResult.Err != nil {
		callback <- DeleteFormAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFormAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteFormAsync(
	request *DeleteFormRequest,
	callback chan<- DeleteFormAsyncResult,
) {
	path := "/{namespaceName}/user/me/mold/{moldModelName}/form/{index}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
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

	go deleteFormAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteForm(
	request *DeleteFormRequest,
) (*DeleteFormResult, error) {
	callback := make(chan DeleteFormAsyncResult, 1)
	go p.DeleteFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeleteFormByUserIdAsync(
	request *DeleteFormByUserIdRequest,
	callback chan<- DeleteFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/mold/{moldModelName}/form/{index}"
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
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}
	if request.Index != nil {
		path = strings.ReplaceAll(path, "{index}", core.ToString(*request.Index))
	} else {
		path = strings.ReplaceAll(path, "{index}", "null")
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

	go deleteFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeleteFormByUserId(
	request *DeleteFormByUserIdRequest,
) (*DeleteFormByUserIdResult, error) {
	callback := make(chan DeleteFormByUserIdAsyncResult, 1)
	go p.DeleteFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireActionToFormPropertiesByStampSheetAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireActionToFormPropertiesByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionToFormPropertiesByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireActionToFormPropertiesByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) AcquireActionToFormPropertiesByStampSheetAsync(
	request *AcquireActionToFormPropertiesByStampSheetRequest,
	callback chan<- AcquireActionToFormPropertiesByStampSheetAsyncResult,
) {
	path := "/stamp/form/acquire"

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

	go acquireActionToFormPropertiesByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AcquireActionToFormPropertiesByStampSheet(
	request *AcquireActionToFormPropertiesByStampSheetRequest,
) (*AcquireActionToFormPropertiesByStampSheetResult, error) {
	callback := make(chan AcquireActionToFormPropertiesByStampSheetAsyncResult, 1)
	go p.AcquireActionToFormPropertiesByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setFormByStampSheetAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetFormByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFormByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFormByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetFormByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetFormByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetFormByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetFormByStampSheetAsync(
	request *SetFormByStampSheetRequest,
	callback chan<- SetFormByStampSheetAsyncResult,
) {
	path := "/stamp/form/set"

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

	go setFormByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetFormByStampSheet(
	request *SetFormByStampSheetRequest,
) (*SetFormByStampSheetResult, error) {
	callback := make(chan SetFormByStampSheetAsyncResult, 1)
	go p.SetFormByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePropertyFormsAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePropertyFormsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePropertyFormsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePropertyFormsResult
	if asyncResult.Err != nil {
		callback <- DescribePropertyFormsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePropertyFormsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePropertyFormsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribePropertyFormsAsync(
	request *DescribePropertyFormsRequest,
	callback chan<- DescribePropertyFormsAsyncResult,
) {
	path := "/{namespaceName}/user/me/property/{propertyFormModelName}/form"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}

	go describePropertyFormsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribePropertyForms(
	request *DescribePropertyFormsRequest,
) (*DescribePropertyFormsResult, error) {
	callback := make(chan DescribePropertyFormsAsyncResult, 1)
	go p.DescribePropertyFormsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePropertyFormsByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePropertyFormsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePropertyFormsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePropertyFormsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribePropertyFormsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePropertyFormsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePropertyFormsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DescribePropertyFormsByUserIdAsync(
	request *DescribePropertyFormsByUserIdRequest,
	callback chan<- DescribePropertyFormsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
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
	if request.TimeOffsetToken != nil {
		headers["X-GS2-TIME-OFFSET-TOKEN"] = string(*request.TimeOffsetToken)
	}

	go describePropertyFormsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DescribePropertyFormsByUserId(
	request *DescribePropertyFormsByUserIdRequest,
) (*DescribePropertyFormsByUserIdResult, error) {
	callback := make(chan DescribePropertyFormsByUserIdAsyncResult, 1)
	go p.DescribePropertyFormsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormAsync(
	request *GetPropertyFormRequest,
	callback chan<- GetPropertyFormAsyncResult,
) {
	path := "/{namespaceName}/user/me/property/{propertyFormModelName}/form/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go getPropertyFormAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyForm(
	request *GetPropertyFormRequest,
) (*GetPropertyFormResult, error) {
	callback := make(chan GetPropertyFormAsyncResult, 1)
	go p.GetPropertyFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormByUserIdAsync(
	request *GetPropertyFormByUserIdRequest,
	callback chan<- GetPropertyFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form/{propertyId}"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go getPropertyFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyFormByUserId(
	request *GetPropertyFormByUserIdRequest,
) (*GetPropertyFormByUserIdResult, error) {
	callback := make(chan GetPropertyFormByUserIdAsyncResult, 1)
	go p.GetPropertyFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormWithSignatureAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormWithSignatureAsync(
	request *GetPropertyFormWithSignatureRequest,
	callback chan<- GetPropertyFormWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/property/{propertyFormModelName}/form/{propertyId}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
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

	go getPropertyFormWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyFormWithSignature(
	request *GetPropertyFormWithSignatureRequest,
) (*GetPropertyFormWithSignatureResult, error) {
	callback := make(chan GetPropertyFormWithSignatureAsyncResult, 1)
	go p.GetPropertyFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPropertyFormWithSignatureByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- GetPropertyFormWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPropertyFormWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPropertyFormWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetPropertyFormWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPropertyFormWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPropertyFormWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) GetPropertyFormWithSignatureByUserIdAsync(
	request *GetPropertyFormWithSignatureByUserIdRequest,
	callback chan<- GetPropertyFormWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form/{propertyId}/signature"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.KeyId != nil {
		queryStrings["keyId"] = core.ToString(*request.KeyId)
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

	go getPropertyFormWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) GetPropertyFormWithSignatureByUserId(
	request *GetPropertyFormWithSignatureByUserIdRequest,
) (*GetPropertyFormWithSignatureByUserIdResult, error) {
	callback := make(chan GetPropertyFormWithSignatureByUserIdAsyncResult, 1)
	go p.GetPropertyFormWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setPropertyFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetPropertyFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetPropertyFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetPropertyFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetPropertyFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetPropertyFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetPropertyFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetPropertyFormByUserIdAsync(
	request *SetPropertyFormByUserIdRequest,
	callback chan<- SetPropertyFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form/{propertyId}"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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

	go setPropertyFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetPropertyFormByUserId(
	request *SetPropertyFormByUserIdRequest,
) (*SetPropertyFormByUserIdResult, error) {
	callback := make(chan SetPropertyFormByUserIdAsyncResult, 1)
	go p.SetPropertyFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setPropertyFormWithSignatureAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- SetPropertyFormWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetPropertyFormWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetPropertyFormWithSignatureResult
	if asyncResult.Err != nil {
		callback <- SetPropertyFormWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetPropertyFormWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetPropertyFormWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) SetPropertyFormWithSignatureAsync(
	request *SetPropertyFormWithSignatureRequest,
	callback chan<- SetPropertyFormWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/property/{propertyFormModelName}/form/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Slots != nil {
		var _slots []interface{}
		for _, item := range request.Slots {
			_slots = append(_slots, item)
		}
		bodies["slots"] = _slots
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
	}

	go setPropertyFormWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) SetPropertyFormWithSignature(
	request *SetPropertyFormWithSignatureRequest,
) (*SetPropertyFormWithSignatureResult, error) {
	callback := make(chan SetPropertyFormWithSignatureAsyncResult, 1)
	go p.SetPropertyFormWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireActionsToPropertyFormPropertiesAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireActionsToPropertyFormPropertiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionsToPropertyFormPropertiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionsToPropertyFormPropertiesResult
	if asyncResult.Err != nil {
		callback <- AcquireActionsToPropertyFormPropertiesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireActionsToPropertyFormPropertiesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireActionsToPropertyFormPropertiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) AcquireActionsToPropertyFormPropertiesAsync(
	request *AcquireActionsToPropertyFormPropertiesRequest,
	callback chan<- AcquireActionsToPropertyFormPropertiesAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form/{propertyId}/stamp/delegate"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AcquireAction != nil {
		bodies["acquireAction"] = request.AcquireAction.ToDict()
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

	go acquireActionsToPropertyFormPropertiesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AcquireActionsToPropertyFormProperties(
	request *AcquireActionsToPropertyFormPropertiesRequest,
) (*AcquireActionsToPropertyFormPropertiesResult, error) {
	callback := make(chan AcquireActionsToPropertyFormPropertiesAsyncResult, 1)
	go p.AcquireActionsToPropertyFormPropertiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePropertyFormAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePropertyFormAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePropertyFormAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePropertyFormResult
	if asyncResult.Err != nil {
		callback <- DeletePropertyFormAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePropertyFormAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePropertyFormAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeletePropertyFormAsync(
	request *DeletePropertyFormRequest,
	callback chan<- DeletePropertyFormAsyncResult,
) {
	path := "/{namespaceName}/user/me/property/{propertyFormModelName}/form/{propertyId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go deletePropertyFormAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeletePropertyForm(
	request *DeletePropertyFormRequest,
) (*DeletePropertyFormResult, error) {
	callback := make(chan DeletePropertyFormAsyncResult, 1)
	go p.DeletePropertyFormAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePropertyFormByUserIdAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePropertyFormByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePropertyFormByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePropertyFormByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeletePropertyFormByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePropertyFormByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePropertyFormByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) DeletePropertyFormByUserIdAsync(
	request *DeletePropertyFormByUserIdRequest,
	callback chan<- DeletePropertyFormByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/property/{propertyFormModelName}/form/{propertyId}"
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
	if request.PropertyFormModelName != nil && *request.PropertyFormModelName != "" {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", core.ToString(*request.PropertyFormModelName))
	} else {
		path = strings.ReplaceAll(path, "{propertyFormModelName}", "null")
	}
	if request.PropertyId != nil && *request.PropertyId != "" {
		path = strings.ReplaceAll(path, "{propertyId}", core.ToString(*request.PropertyId))
	} else {
		path = strings.ReplaceAll(path, "{propertyId}", "null")
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

	go deletePropertyFormByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) DeletePropertyFormByUserId(
	request *DeletePropertyFormByUserIdRequest,
) (*DeletePropertyFormByUserIdResult, error) {
	callback := make(chan DeletePropertyFormByUserIdAsyncResult, 1)
	go p.DeletePropertyFormByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acquireActionToPropertyFormPropertiesByStampSheetAsyncHandler(
	client Gs2FormationRestClient,
	job *core.NetworkJob,
	callback chan<- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcquireActionToPropertyFormPropertiesByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2FormationRestClient) AcquireActionToPropertyFormPropertiesByStampSheetAsync(
	request *AcquireActionToPropertyFormPropertiesByStampSheetRequest,
	callback chan<- AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult,
) {
	path := "/stamp/property/form/acquire"

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

	go acquireActionToPropertyFormPropertiesByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("formation").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2FormationRestClient) AcquireActionToPropertyFormPropertiesByStampSheet(
	request *AcquireActionToPropertyFormPropertiesByStampSheetRequest,
) (*AcquireActionToPropertyFormPropertiesByStampSheetResult, error) {
	callback := make(chan AcquireActionToPropertyFormPropertiesByStampSheetAsyncResult, 1)
	go p.AcquireActionToPropertyFormPropertiesByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
