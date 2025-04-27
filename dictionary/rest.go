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

package dictionary

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2DictionaryRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2DictionaryRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeNamespaces(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) CreateNamespaceAsync(
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
	if request.EntryScript != nil {
		bodies["entryScript"] = request.EntryScript.ToDict()
	}
	if request.DuplicateEntryScript != nil && *request.DuplicateEntryScript != "" {
		bodies["duplicateEntryScript"] = *request.DuplicateEntryScript
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
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CreateNamespace(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetNamespaceStatus(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetNamespace(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) UpdateNamespaceAsync(
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
	if request.EntryScript != nil {
		bodies["entryScript"] = request.EntryScript.ToDict()
	}
	if request.DuplicateEntryScript != nil && *request.DuplicateEntryScript != "" {
		bodies["duplicateEntryScript"] = *request.DuplicateEntryScript
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
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) UpdateNamespace(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteNamespace(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) DumpUserDataByUserIdAsync(
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

	go dumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DumpUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) CheckDumpUserDataByUserIdAsync(
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

	go checkDumpUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CheckDumpUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) CleanUserDataByUserIdAsync(
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

	go cleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CleanUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) CheckCleanUserDataByUserIdAsync(
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

	go checkCleanUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CheckCleanUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) PrepareImportUserDataByUserIdAsync(
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

	go prepareImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) PrepareImportUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) ImportUserDataByUserIdAsync(
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

	go importUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) ImportUserDataByUserId(
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
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) CheckImportUserDataByUserIdAsync(
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

	go checkImportUserDataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CheckImportUserDataByUserId(
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

func describeEntryModelsAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEntryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEntryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEntryModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeEntryModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEntryModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEntryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeEntryModelsAsync(
	request *DescribeEntryModelsRequest,
	callback chan<- DescribeEntryModelsAsyncResult,
) {
	path := "/{namespaceName}/model"
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

	go describeEntryModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeEntryModels(
	request *DescribeEntryModelsRequest,
) (*DescribeEntryModelsResult, error) {
	callback := make(chan DescribeEntryModelsAsyncResult, 1)
	go p.DescribeEntryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryModelAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryModelResult
	if asyncResult.Err != nil {
		callback <- GetEntryModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryModelAsync(
	request *GetEntryModelRequest,
	callback chan<- GetEntryModelAsyncResult,
) {
	path := "/{namespaceName}/model/{entryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryName != nil && *request.EntryName != "" {
		path = strings.ReplaceAll(path, "{entryName}", core.ToString(*request.EntryName))
	} else {
		path = strings.ReplaceAll(path, "{entryName}", "null")
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

	go getEntryModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntryModel(
	request *GetEntryModelRequest,
) (*GetEntryModelResult, error) {
	callback := make(chan GetEntryModelAsyncResult, 1)
	go p.GetEntryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEntryModelMastersAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEntryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEntryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEntryModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeEntryModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEntryModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEntryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeEntryModelMastersAsync(
	request *DescribeEntryModelMastersRequest,
	callback chan<- DescribeEntryModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model"
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

	go describeEntryModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeEntryModelMasters(
	request *DescribeEntryModelMastersRequest,
) (*DescribeEntryModelMastersResult, error) {
	callback := make(chan DescribeEntryModelMastersAsyncResult, 1)
	go p.DescribeEntryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createEntryModelMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- CreateEntryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateEntryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateEntryModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateEntryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateEntryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateEntryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) CreateEntryModelMasterAsync(
	request *CreateEntryModelMasterRequest,
	callback chan<- CreateEntryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model"
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

	go createEntryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) CreateEntryModelMaster(
	request *CreateEntryModelMasterRequest,
) (*CreateEntryModelMasterResult, error) {
	callback := make(chan CreateEntryModelMasterAsyncResult, 1)
	go p.CreateEntryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryModelMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetEntryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryModelMasterAsync(
	request *GetEntryModelMasterRequest,
	callback chan<- GetEntryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{entryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryName != nil && *request.EntryName != "" {
		path = strings.ReplaceAll(path, "{entryName}", core.ToString(*request.EntryName))
	} else {
		path = strings.ReplaceAll(path, "{entryName}", "null")
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

	go getEntryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntryModelMaster(
	request *GetEntryModelMasterRequest,
) (*GetEntryModelMasterResult, error) {
	callback := make(chan GetEntryModelMasterAsyncResult, 1)
	go p.GetEntryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateEntryModelMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateEntryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateEntryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateEntryModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateEntryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateEntryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateEntryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) UpdateEntryModelMasterAsync(
	request *UpdateEntryModelMasterRequest,
	callback chan<- UpdateEntryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{entryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryName != nil && *request.EntryName != "" {
		path = strings.ReplaceAll(path, "{entryName}", core.ToString(*request.EntryName))
	} else {
		path = strings.ReplaceAll(path, "{entryName}", "null")
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

	go updateEntryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) UpdateEntryModelMaster(
	request *UpdateEntryModelMasterRequest,
) (*UpdateEntryModelMasterResult, error) {
	callback := make(chan UpdateEntryModelMasterAsyncResult, 1)
	go p.UpdateEntryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteEntryModelMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteEntryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteEntryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteEntryModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteEntryModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteEntryModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteEntryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteEntryModelMasterAsync(
	request *DeleteEntryModelMasterRequest,
	callback chan<- DeleteEntryModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{entryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryName != nil && *request.EntryName != "" {
		path = strings.ReplaceAll(path, "{entryName}", core.ToString(*request.EntryName))
	} else {
		path = strings.ReplaceAll(path, "{entryName}", "null")
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

	go deleteEntryModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteEntryModelMaster(
	request *DeleteEntryModelMasterRequest,
) (*DeleteEntryModelMasterResult, error) {
	callback := make(chan DeleteEntryModelMasterAsyncResult, 1)
	go p.DeleteEntryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEntriesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEntriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEntriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEntriesResult
	if asyncResult.Err != nil {
		callback <- DescribeEntriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEntriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEntriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeEntriesAsync(
	request *DescribeEntriesRequest,
	callback chan<- DescribeEntriesAsyncResult,
) {
	path := "/{namespaceName}/user/me/entry"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeEntriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeEntries(
	request *DescribeEntriesRequest,
) (*DescribeEntriesResult, error) {
	callback := make(chan DescribeEntriesAsyncResult, 1)
	go p.DescribeEntriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEntriesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEntriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEntriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEntriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeEntriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEntriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEntriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeEntriesByUserIdAsync(
	request *DescribeEntriesByUserIdRequest,
	callback chan<- DescribeEntriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry"
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

	go describeEntriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeEntriesByUserId(
	request *DescribeEntriesByUserIdRequest,
) (*DescribeEntriesByUserIdResult, error) {
	callback := make(chan DescribeEntriesByUserIdAsyncResult, 1)
	go p.DescribeEntriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addEntriesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- AddEntriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddEntriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddEntriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddEntriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddEntriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddEntriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) AddEntriesByUserIdAsync(
	request *AddEntriesByUserIdRequest,
	callback chan<- AddEntriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry"
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
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go addEntriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) AddEntriesByUserId(
	request *AddEntriesByUserIdRequest,
) (*AddEntriesByUserIdResult, error) {
	callback := make(chan AddEntriesByUserIdAsyncResult, 1)
	go p.AddEntriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryResult
	if asyncResult.Err != nil {
		callback <- GetEntryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryAsync(
	request *GetEntryRequest,
	callback chan<- GetEntryAsyncResult,
) {
	path := "/{namespaceName}/user/me/entry/{entryModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getEntryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntry(
	request *GetEntryRequest,
) (*GetEntryResult, error) {
	callback := make(chan GetEntryAsyncResult, 1)
	go p.GetEntryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetEntryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryByUserIdAsync(
	request *GetEntryByUserIdRequest,
	callback chan<- GetEntryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry/{entryModelName}"
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
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go getEntryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntryByUserId(
	request *GetEntryByUserIdRequest,
) (*GetEntryByUserIdResult, error) {
	callback := make(chan GetEntryByUserIdAsyncResult, 1)
	go p.GetEntryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryWithSignatureAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryWithSignatureResult
	if asyncResult.Err != nil {
		callback <- GetEntryWithSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryWithSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryWithSignatureAsync(
	request *GetEntryWithSignatureRequest,
	callback chan<- GetEntryWithSignatureAsyncResult,
) {
	path := "/{namespaceName}/user/me/entry/{entryModelName}/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getEntryWithSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntryWithSignature(
	request *GetEntryWithSignatureRequest,
) (*GetEntryWithSignatureResult, error) {
	callback := make(chan GetEntryWithSignatureAsyncResult, 1)
	go p.GetEntryWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEntryWithSignatureByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetEntryWithSignatureByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEntryWithSignatureByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEntryWithSignatureByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetEntryWithSignatureByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEntryWithSignatureByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEntryWithSignatureByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetEntryWithSignatureByUserIdAsync(
	request *GetEntryWithSignatureByUserIdRequest,
	callback chan<- GetEntryWithSignatureByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry/{entryModelName}/signature"
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
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
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

	go getEntryWithSignatureByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetEntryWithSignatureByUserId(
	request *GetEntryWithSignatureByUserIdRequest,
) (*GetEntryWithSignatureByUserIdResult, error) {
	callback := make(chan GetEntryWithSignatureByUserIdAsyncResult, 1)
	go p.GetEntryWithSignatureByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) ResetByUserIdAsync(
	request *ResetByUserIdRequest,
	callback chan<- ResetByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry"
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go resetByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) ResetByUserId(
	request *ResetByUserIdRequest,
) (*ResetByUserIdResult, error) {
	callback := make(chan ResetByUserIdAsyncResult, 1)
	go p.ResetByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyEntryAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyEntryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyEntryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyEntryResult
	if asyncResult.Err != nil {
		callback <- VerifyEntryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyEntryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyEntryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) VerifyEntryAsync(
	request *VerifyEntryRequest,
	callback chan<- VerifyEntryAsyncResult,
) {
	path := "/{namespaceName}/user/me/entry/{entryModelName}/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
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

	go verifyEntryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) VerifyEntry(
	request *VerifyEntryRequest,
) (*VerifyEntryResult, error) {
	callback := make(chan VerifyEntryAsyncResult, 1)
	go p.VerifyEntryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyEntryByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyEntryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyEntryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyEntryByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyEntryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyEntryByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyEntryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) VerifyEntryByUserIdAsync(
	request *VerifyEntryByUserIdRequest,
	callback chan<- VerifyEntryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry/{entryModelName}/verify/{verifyType}"
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
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
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

	go verifyEntryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) VerifyEntryByUserId(
	request *VerifyEntryByUserIdRequest,
) (*VerifyEntryByUserIdResult, error) {
	callback := make(chan VerifyEntryByUserIdAsyncResult, 1)
	go p.VerifyEntryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteEntriesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteEntriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteEntriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteEntriesResult
	if asyncResult.Err != nil {
		callback <- DeleteEntriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteEntriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteEntriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteEntriesAsync(
	request *DeleteEntriesRequest,
	callback chan<- DeleteEntriesAsyncResult,
) {
	path := "/{namespaceName}/user/me/entry/delete"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go deleteEntriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteEntries(
	request *DeleteEntriesRequest,
) (*DeleteEntriesResult, error) {
	callback := make(chan DeleteEntriesAsyncResult, 1)
	go p.DeleteEntriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteEntriesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteEntriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteEntriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteEntriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteEntriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteEntriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteEntriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteEntriesByUserIdAsync(
	request *DeleteEntriesByUserIdRequest,
	callback chan<- DeleteEntriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/entry/delete"
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
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go deleteEntriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteEntriesByUserId(
	request *DeleteEntriesByUserIdRequest,
) (*DeleteEntriesByUserIdResult, error) {
	callback := make(chan DeleteEntriesByUserIdAsyncResult, 1)
	go p.DeleteEntriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addEntriesByStampSheetAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- AddEntriesByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddEntriesByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddEntriesByStampSheetResult
	if asyncResult.Err != nil {
		callback <- AddEntriesByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddEntriesByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddEntriesByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) AddEntriesByStampSheetAsync(
	request *AddEntriesByStampSheetRequest,
	callback chan<- AddEntriesByStampSheetAsyncResult,
) {
	path := "/stamp/entry/add"

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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go addEntriesByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) AddEntriesByStampSheet(
	request *AddEntriesByStampSheetRequest,
) (*AddEntriesByStampSheetResult, error) {
	callback := make(chan AddEntriesByStampSheetAsyncResult, 1)
	go p.AddEntriesByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteEntriesByStampTaskAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteEntriesByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteEntriesByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteEntriesByStampTaskResult
	if asyncResult.Err != nil {
		callback <- DeleteEntriesByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteEntriesByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteEntriesByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteEntriesByStampTaskAsync(
	request *DeleteEntriesByStampTaskRequest,
	callback chan<- DeleteEntriesByStampTaskAsyncResult,
) {
	path := "/stamp/entry/delete"

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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteEntriesByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteEntriesByStampTask(
	request *DeleteEntriesByStampTaskRequest,
) (*DeleteEntriesByStampTaskResult, error) {
	callback := make(chan DeleteEntriesByStampTaskAsyncResult, 1)
	go p.DeleteEntriesByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyEntryByStampTaskAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyEntryByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyEntryByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyEntryByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyEntryByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyEntryByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyEntryByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) VerifyEntryByStampTaskAsync(
	request *VerifyEntryByStampTaskRequest,
	callback chan<- VerifyEntryByStampTaskAsyncResult,
) {
	path := "/stamp/entry/verify"

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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go verifyEntryByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) VerifyEntryByStampTask(
	request *VerifyEntryByStampTaskRequest,
) (*VerifyEntryByStampTaskResult, error) {
	callback := make(chan VerifyEntryByStampTaskAsyncResult, 1)
	go p.VerifyEntryByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLikesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLikesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLikesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLikesResult
	if asyncResult.Err != nil {
		callback <- DescribeLikesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLikesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLikesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeLikesAsync(
	request *DescribeLikesRequest,
	callback chan<- DescribeLikesAsyncResult,
) {
	path := "/{namespaceName}/user/me/like"
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
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLikesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeLikes(
	request *DescribeLikesRequest,
) (*DescribeLikesResult, error) {
	callback := make(chan DescribeLikesAsyncResult, 1)
	go p.DescribeLikesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeLikesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLikesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLikesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLikesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeLikesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLikesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLikesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DescribeLikesByUserIdAsync(
	request *DescribeLikesByUserIdRequest,
	callback chan<- DescribeLikesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/like"
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

	go describeLikesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DescribeLikesByUserId(
	request *DescribeLikesByUserIdRequest,
) (*DescribeLikesByUserIdResult, error) {
	callback := make(chan DescribeLikesByUserIdAsyncResult, 1)
	go p.DescribeLikesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addLikesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- AddLikesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddLikesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddLikesResult
	if asyncResult.Err != nil {
		callback <- AddLikesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddLikesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddLikesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) AddLikesAsync(
	request *AddLikesRequest,
	callback chan<- AddLikesAsyncResult,
) {
	path := "/{namespaceName}/user/me/like"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go addLikesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) AddLikes(
	request *AddLikesRequest,
) (*AddLikesResult, error) {
	callback := make(chan AddLikesAsyncResult, 1)
	go p.AddLikesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addLikesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- AddLikesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddLikesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddLikesByUserIdResult
	if asyncResult.Err != nil {
		callback <- AddLikesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddLikesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddLikesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) AddLikesByUserIdAsync(
	request *AddLikesByUserIdRequest,
	callback chan<- AddLikesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/like"
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
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go addLikesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) AddLikesByUserId(
	request *AddLikesByUserIdRequest,
) (*AddLikesByUserIdResult, error) {
	callback := make(chan AddLikesByUserIdAsyncResult, 1)
	go p.AddLikesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLikeAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLikeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLikeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLikeResult
	if asyncResult.Err != nil {
		callback <- GetLikeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLikeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLikeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetLikeAsync(
	request *GetLikeRequest,
	callback chan<- GetLikeAsyncResult,
) {
	path := "/{namespaceName}/user/me/like/{entryModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLikeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetLike(
	request *GetLikeRequest,
) (*GetLikeResult, error) {
	callback := make(chan GetLikeAsyncResult, 1)
	go p.GetLikeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLikeByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetLikeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLikeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLikeByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetLikeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLikeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLikeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetLikeByUserIdAsync(
	request *GetLikeByUserIdRequest,
	callback chan<- GetLikeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/like/{entryModelName}"
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
	if request.EntryModelName != nil && *request.EntryModelName != "" {
		path = strings.ReplaceAll(path, "{entryModelName}", core.ToString(*request.EntryModelName))
	} else {
		path = strings.ReplaceAll(path, "{entryModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go getLikeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetLikeByUserId(
	request *GetLikeByUserIdRequest,
) (*GetLikeByUserIdResult, error) {
	callback := make(chan GetLikeByUserIdAsyncResult, 1)
	go p.GetLikeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetLikesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetLikesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetLikesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetLikesResult
	if asyncResult.Err != nil {
		callback <- ResetLikesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetLikesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetLikesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) ResetLikesAsync(
	request *ResetLikesRequest,
	callback chan<- ResetLikesAsyncResult,
) {
	path := "/{namespaceName}/user/me/like"
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

	go resetLikesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) ResetLikes(
	request *ResetLikesRequest,
) (*ResetLikesResult, error) {
	callback := make(chan ResetLikesAsyncResult, 1)
	go p.ResetLikesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func resetLikesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- ResetLikesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ResetLikesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ResetLikesByUserIdResult
	if asyncResult.Err != nil {
		callback <- ResetLikesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ResetLikesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ResetLikesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) ResetLikesByUserIdAsync(
	request *ResetLikesByUserIdRequest,
	callback chan<- ResetLikesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/like"
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
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
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

	go resetLikesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) ResetLikesByUserId(
	request *ResetLikesByUserIdRequest,
) (*ResetLikesByUserIdResult, error) {
	callback := make(chan ResetLikesByUserIdAsyncResult, 1)
	go p.ResetLikesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLikesAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLikesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLikesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLikesResult
	if asyncResult.Err != nil {
		callback <- DeleteLikesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteLikesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteLikesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteLikesAsync(
	request *DeleteLikesRequest,
	callback chan<- DeleteLikesAsyncResult,
) {
	path := "/{namespaceName}/user/me/like/delete"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go deleteLikesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteLikes(
	request *DeleteLikesRequest,
) (*DeleteLikesResult, error) {
	callback := make(chan DeleteLikesAsyncResult, 1)
	go p.DeleteLikesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteLikesByUserIdAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteLikesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteLikesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteLikesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteLikesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteLikesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteLikesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) DeleteLikesByUserIdAsync(
	request *DeleteLikesByUserIdRequest,
	callback chan<- DeleteLikesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/like/delete"
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
	var bodies = core.Bodies{}
	if request.EntryModelNames != nil {
		var _entryModelNames []interface{}
		for _, item := range request.EntryModelNames {
			_entryModelNames = append(_entryModelNames, item)
		}
		bodies["entryModelNames"] = _entryModelNames
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

	go deleteLikesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) DeleteLikesByUserId(
	request *DeleteLikesByUserIdRequest,
) (*DeleteLikesByUserIdResult, error) {
	callback := make(chan DeleteLikesByUserIdAsyncResult, 1)
	go p.DeleteLikesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2DictionaryRestClient,
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

func (p Gs2DictionaryRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) ExportMaster(
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

func getCurrentEntryMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentEntryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentEntryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentEntryMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentEntryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentEntryMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentEntryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) GetCurrentEntryMasterAsync(
	request *GetCurrentEntryMasterRequest,
	callback chan<- GetCurrentEntryMasterAsyncResult,
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

	go getCurrentEntryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) GetCurrentEntryMaster(
	request *GetCurrentEntryMasterRequest,
) (*GetCurrentEntryMasterResult, error) {
	callback := make(chan GetCurrentEntryMasterAsyncResult, 1)
	go p.GetCurrentEntryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentEntryMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentEntryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentEntryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentEntryMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentEntryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentEntryMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentEntryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) PreUpdateCurrentEntryMasterAsync(
	request *PreUpdateCurrentEntryMasterRequest,
	callback chan<- PreUpdateCurrentEntryMasterAsyncResult,
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

	go preUpdateCurrentEntryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) PreUpdateCurrentEntryMaster(
	request *PreUpdateCurrentEntryMasterRequest,
) (*PreUpdateCurrentEntryMasterResult, error) {
	callback := make(chan PreUpdateCurrentEntryMasterAsyncResult, 1)
	go p.PreUpdateCurrentEntryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentEntryMasterAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentEntryMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentEntryMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentEntryMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentEntryMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentEntryMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentEntryMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) UpdateCurrentEntryMasterAsync(
	request *UpdateCurrentEntryMasterRequest,
	callback chan<- UpdateCurrentEntryMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentEntryMaster(
			&PreUpdateCurrentEntryMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentEntryMasterAsyncResult{
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
				callback <- UpdateCurrentEntryMasterAsyncResult{
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

	go updateCurrentEntryMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) UpdateCurrentEntryMaster(
	request *UpdateCurrentEntryMasterRequest,
) (*UpdateCurrentEntryMasterResult, error) {
	callback := make(chan UpdateCurrentEntryMasterAsyncResult, 1)
	go p.UpdateCurrentEntryMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentEntryMasterFromGitHubAsyncHandler(
	client Gs2DictionaryRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentEntryMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentEntryMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentEntryMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentEntryMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentEntryMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentEntryMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DictionaryRestClient) UpdateCurrentEntryMasterFromGitHubAsync(
	request *UpdateCurrentEntryMasterFromGitHubRequest,
	callback chan<- UpdateCurrentEntryMasterFromGitHubAsyncResult,
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

	go updateCurrentEntryMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("dictionary").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DictionaryRestClient) UpdateCurrentEntryMasterFromGitHub(
	request *UpdateCurrentEntryMasterFromGitHubRequest,
) (*UpdateCurrentEntryMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentEntryMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentEntryMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
