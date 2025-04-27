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

package inbox

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2InboxRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2InboxRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DescribeNamespaces(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) CreateNamespaceAsync(
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
	if request.IsAutomaticDeletingEnabled != nil {
		bodies["isAutomaticDeletingEnabled"] = *request.IsAutomaticDeletingEnabled
	}
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReceiveMessageScript != nil {
		bodies["receiveMessageScript"] = request.ReceiveMessageScript.ToDict()
	}
	if request.ReadMessageScript != nil {
		bodies["readMessageScript"] = request.ReadMessageScript.ToDict()
	}
	if request.DeleteMessageScript != nil {
		bodies["deleteMessageScript"] = request.DeleteMessageScript.ToDict()
	}
	if request.ReceiveNotification != nil {
		bodies["receiveNotification"] = request.ReceiveNotification.ToDict()
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CreateNamespace(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetNamespaceStatus(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetNamespace(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) UpdateNamespaceAsync(
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
	if request.IsAutomaticDeletingEnabled != nil {
		bodies["isAutomaticDeletingEnabled"] = *request.IsAutomaticDeletingEnabled
	}
	if request.TransactionSetting != nil {
		bodies["transactionSetting"] = request.TransactionSetting.ToDict()
	}
	if request.ReceiveMessageScript != nil {
		bodies["receiveMessageScript"] = request.ReceiveMessageScript.ToDict()
	}
	if request.ReadMessageScript != nil {
		bodies["readMessageScript"] = request.ReadMessageScript.ToDict()
	}
	if request.DeleteMessageScript != nil {
		bodies["deleteMessageScript"] = request.DeleteMessageScript.ToDict()
	}
	if request.ReceiveNotification != nil {
		bodies["receiveNotification"] = request.ReceiveNotification.ToDict()
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) UpdateNamespace(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteNamespace(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DumpUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CheckDumpUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CleanUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CheckCleanUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) PrepareImportUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ImportUserDataByUserId(
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
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CheckImportUserDataByUserId(
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

func describeMessagesAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesResult
	if asyncResult.Err != nil {
		callback <- DescribeMessagesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMessagesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DescribeMessagesAsync(
	request *DescribeMessagesRequest,
	callback chan<- DescribeMessagesAsyncResult,
) {
	path := "/{namespaceName}/user/me/message"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.IsRead != nil {
		queryStrings["isRead"] = core.ToString(*request.IsRead)
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

	go describeMessagesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DescribeMessages(
	request *DescribeMessagesRequest,
) (*DescribeMessagesResult, error) {
	callback := make(chan DescribeMessagesAsyncResult, 1)
	go p.DescribeMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMessagesByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMessagesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMessagesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeMessagesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMessagesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMessagesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DescribeMessagesByUserIdAsync(
	request *DescribeMessagesByUserIdRequest,
	callback chan<- DescribeMessagesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/message"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserId != nil {
		queryStrings["userId"] = core.ToString(*request.UserId)
	}
	if request.IsRead != nil {
		queryStrings["isRead"] = core.ToString(*request.IsRead)
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

	go describeMessagesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DescribeMessagesByUserId(
	request *DescribeMessagesByUserIdRequest,
) (*DescribeMessagesByUserIdResult, error) {
	callback := make(chan DescribeMessagesByUserIdAsyncResult, 1)
	go p.DescribeMessagesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- SendMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendMessageByUserIdResult
	if asyncResult.Err != nil {
		callback <- SendMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) SendMessageByUserIdAsync(
	request *SendMessageByUserIdRequest,
	callback chan<- SendMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/message"
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
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReadAcquireActions != nil {
		var _readAcquireActions []interface{}
		for _, item := range request.ReadAcquireActions {
			_readAcquireActions = append(_readAcquireActions, item)
		}
		bodies["readAcquireActions"] = _readAcquireActions
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.ExpiresTimeSpan != nil {
		bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
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

	go sendMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) SendMessageByUserId(
	request *SendMessageByUserIdRequest,
) (*SendMessageByUserIdResult, error) {
	callback := make(chan SendMessageByUserIdAsyncResult, 1)
	go p.SendMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageResult
	if asyncResult.Err != nil {
		callback <- GetMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetMessageAsync(
	request *GetMessageRequest,
	callback chan<- GetMessageAsyncResult,
) {
	path := "/{namespaceName}/user/me/{messageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go getMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetMessage(
	request *GetMessageRequest,
) (*GetMessageResult, error) {
	callback := make(chan GetMessageAsyncResult, 1)
	go p.GetMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMessageByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetMessageByUserIdAsync(
	request *GetMessageByUserIdRequest,
	callback chan<- GetMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{messageName}"
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
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go getMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetMessageByUserId(
	request *GetMessageByUserIdRequest,
) (*GetMessageByUserIdResult, error) {
	callback := make(chan GetMessageByUserIdAsyncResult, 1)
	go p.GetMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveGlobalMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveGlobalMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalMessageResult
	if asyncResult.Err != nil {
		callback <- ReceiveGlobalMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveGlobalMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) ReceiveGlobalMessageAsync(
	request *ReceiveGlobalMessageRequest,
	callback chan<- ReceiveGlobalMessageAsyncResult,
) {
	path := "/{namespaceName}/user/me/message/globalMessage/receive"
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

	go receiveGlobalMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ReceiveGlobalMessage(
	request *ReceiveGlobalMessageRequest,
) (*ReceiveGlobalMessageResult, error) {
	callback := make(chan ReceiveGlobalMessageAsyncResult, 1)
	go p.ReceiveGlobalMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func receiveGlobalMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- ReceiveGlobalMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReceiveGlobalMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReceiveGlobalMessageByUserIdResult
	if asyncResult.Err != nil {
		callback <- ReceiveGlobalMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReceiveGlobalMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReceiveGlobalMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) ReceiveGlobalMessageByUserIdAsync(
	request *ReceiveGlobalMessageByUserIdRequest,
	callback chan<- ReceiveGlobalMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/message/globalMessage/receive"
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

	go receiveGlobalMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ReceiveGlobalMessageByUserId(
	request *ReceiveGlobalMessageByUserIdRequest,
) (*ReceiveGlobalMessageByUserIdResult, error) {
	callback := make(chan ReceiveGlobalMessageByUserIdAsyncResult, 1)
	go p.ReceiveGlobalMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func openMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- OpenMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenMessageResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- OpenMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- OpenMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- OpenMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) OpenMessageAsync(
	request *OpenMessageRequest,
	callback chan<- OpenMessageAsyncResult,
) {
	path := "/{namespaceName}/user/me/{messageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go openMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) OpenMessage(
	request *OpenMessageRequest,
) (*OpenMessageResult, error) {
	callback := make(chan OpenMessageAsyncResult, 1)
	go p.OpenMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func openMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- OpenMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenMessageByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- OpenMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- OpenMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- OpenMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) OpenMessageByUserIdAsync(
	request *OpenMessageByUserIdRequest,
	callback chan<- OpenMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{messageName}"
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
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go openMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) OpenMessageByUserId(
	request *OpenMessageByUserIdRequest,
) (*OpenMessageByUserIdResult, error) {
	callback := make(chan OpenMessageByUserIdAsyncResult, 1)
	go p.OpenMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func closeMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- CloseMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CloseMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CloseMessageByUserIdResult
	if asyncResult.Err != nil {
		callback <- CloseMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CloseMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CloseMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) CloseMessageByUserIdAsync(
	request *CloseMessageByUserIdRequest,
	callback chan<- CloseMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{messageName}/close"
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
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go closeMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CloseMessageByUserId(
	request *CloseMessageByUserIdRequest,
) (*CloseMessageByUserIdResult, error) {
	callback := make(chan CloseMessageByUserIdAsyncResult, 1)
	go p.CloseMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func readMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- ReadMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReadMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReadMessageResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- ReadMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReadMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReadMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) ReadMessageAsync(
	request *ReadMessageRequest,
	callback chan<- ReadMessageAsyncResult,
) {
	path := "/{namespaceName}/user/me/{messageName}/read"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go readMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ReadMessage(
	request *ReadMessageRequest,
) (*ReadMessageResult, error) {
	callback := make(chan ReadMessageAsyncResult, 1)
	go p.ReadMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func readMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- ReadMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ReadMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ReadMessageByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- ReadMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ReadMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ReadMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) ReadMessageByUserIdAsync(
	request *ReadMessageByUserIdRequest,
	callback chan<- ReadMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{messageName}/read"
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
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go readMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ReadMessageByUserId(
	request *ReadMessageByUserIdRequest,
) (*ReadMessageByUserIdResult, error) {
	callback := make(chan ReadMessageByUserIdAsyncResult, 1)
	go p.ReadMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchReadMessagesAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- BatchReadMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchReadMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchReadMessagesResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- BatchReadMessagesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchReadMessagesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchReadMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) BatchReadMessagesAsync(
	request *BatchReadMessagesRequest,
	callback chan<- BatchReadMessagesAsyncResult,
) {
	path := "/{namespaceName}/user/me/messages/read/batch"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.MessageNames != nil {
		var _messageNames []interface{}
		for _, item := range request.MessageNames {
			_messageNames = append(_messageNames, item)
		}
		bodies["messageNames"] = _messageNames
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

	go batchReadMessagesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) BatchReadMessages(
	request *BatchReadMessagesRequest,
) (*BatchReadMessagesResult, error) {
	callback := make(chan BatchReadMessagesAsyncResult, 1)
	go p.BatchReadMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchReadMessagesByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- BatchReadMessagesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchReadMessagesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchReadMessagesByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "inbox.message.expired" {
				asyncResult.Err = gs2err.SetClientError(MessageExpired{})
			}
		}
		callback <- BatchReadMessagesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchReadMessagesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchReadMessagesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) BatchReadMessagesByUserIdAsync(
	request *BatchReadMessagesByUserIdRequest,
	callback chan<- BatchReadMessagesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/messages/read/batch"
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
	if request.MessageNames != nil {
		var _messageNames []interface{}
		for _, item := range request.MessageNames {
			_messageNames = append(_messageNames, item)
		}
		bodies["messageNames"] = _messageNames
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

	go batchReadMessagesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) BatchReadMessagesByUserId(
	request *BatchReadMessagesByUserIdRequest,
) (*BatchReadMessagesByUserIdResult, error) {
	callback := make(chan BatchReadMessagesByUserIdAsyncResult, 1)
	go p.BatchReadMessagesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageResult
	if asyncResult.Err != nil {
		callback <- DeleteMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DeleteMessageAsync(
	request *DeleteMessageRequest,
	callback chan<- DeleteMessageAsyncResult,
) {
	path := "/{namespaceName}/user/me/{messageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go deleteMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteMessage(
	request *DeleteMessageRequest,
) (*DeleteMessageResult, error) {
	callback := make(chan DeleteMessageAsyncResult, 1)
	go p.DeleteMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMessageByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMessageByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteMessageByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMessageByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMessageByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DeleteMessageByUserIdAsync(
	request *DeleteMessageByUserIdRequest,
	callback chan<- DeleteMessageByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{messageName}"
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
	if request.MessageName != nil && *request.MessageName != "" {
		path = strings.ReplaceAll(path, "{messageName}", core.ToString(*request.MessageName))
	} else {
		path = strings.ReplaceAll(path, "{messageName}", "null")
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

	go deleteMessageByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteMessageByUserId(
	request *DeleteMessageByUserIdRequest,
) (*DeleteMessageByUserIdResult, error) {
	callback := make(chan DeleteMessageByUserIdAsyncResult, 1)
	go p.DeleteMessageByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendByStampSheetAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- SendByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SendByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) SendByStampSheetAsync(
	request *SendByStampSheetRequest,
	callback chan<- SendByStampSheetAsyncResult,
) {
	path := "/stamp/send"

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

	go sendByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) SendByStampSheet(
	request *SendByStampSheetRequest,
) (*SendByStampSheetResult, error) {
	callback := make(chan SendByStampSheetAsyncResult, 1)
	go p.SendByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func openByStampTaskAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- OpenByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- OpenByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result OpenByStampTaskResult
	if asyncResult.Err != nil {
		callback <- OpenByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- OpenByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- OpenByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) OpenByStampTaskAsync(
	request *OpenByStampTaskRequest,
	callback chan<- OpenByStampTaskAsyncResult,
) {
	path := "/stamp/open"

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

	go openByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) OpenByStampTask(
	request *OpenByStampTaskRequest,
) (*OpenByStampTaskResult, error) {
	callback := make(chan OpenByStampTaskAsyncResult, 1)
	go p.OpenByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMessageByStampTaskAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMessageByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMessageByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMessageByStampTaskResult
	if asyncResult.Err != nil {
		callback <- DeleteMessageByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMessageByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMessageByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DeleteMessageByStampTaskAsync(
	request *DeleteMessageByStampTaskRequest,
	callback chan<- DeleteMessageByStampTaskAsyncResult,
) {
	path := "/stamp/delete"

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

	go deleteMessageByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteMessageByStampTask(
	request *DeleteMessageByStampTaskRequest,
) (*DeleteMessageByStampTaskResult, error) {
	callback := make(chan DeleteMessageByStampTaskAsyncResult, 1)
	go p.DeleteMessageByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2InboxRestClient,
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

func (p Gs2InboxRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) ExportMaster(
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

func getCurrentMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentMessageMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetCurrentMessageMasterAsync(
	request *GetCurrentMessageMasterRequest,
	callback chan<- GetCurrentMessageMasterAsyncResult,
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

	go getCurrentMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetCurrentMessageMaster(
	request *GetCurrentMessageMasterRequest,
) (*GetCurrentMessageMasterResult, error) {
	callback := make(chan GetCurrentMessageMasterAsyncResult, 1)
	go p.GetCurrentMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentMessageMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) PreUpdateCurrentMessageMasterAsync(
	request *PreUpdateCurrentMessageMasterRequest,
	callback chan<- PreUpdateCurrentMessageMasterAsyncResult,
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

	go preUpdateCurrentMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) PreUpdateCurrentMessageMaster(
	request *PreUpdateCurrentMessageMasterRequest,
) (*PreUpdateCurrentMessageMasterResult, error) {
	callback := make(chan PreUpdateCurrentMessageMasterAsyncResult, 1)
	go p.PreUpdateCurrentMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMessageMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) UpdateCurrentMessageMasterAsync(
	request *UpdateCurrentMessageMasterRequest,
	callback chan<- UpdateCurrentMessageMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentMessageMaster(
			&PreUpdateCurrentMessageMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentMessageMasterAsyncResult{
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
				callback <- UpdateCurrentMessageMasterAsyncResult{
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

	go updateCurrentMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) UpdateCurrentMessageMaster(
	request *UpdateCurrentMessageMasterRequest,
) (*UpdateCurrentMessageMasterResult, error) {
	callback := make(chan UpdateCurrentMessageMasterAsyncResult, 1)
	go p.UpdateCurrentMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentMessageMasterFromGitHubAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentMessageMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentMessageMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentMessageMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) UpdateCurrentMessageMasterFromGitHubAsync(
	request *UpdateCurrentMessageMasterFromGitHubRequest,
	callback chan<- UpdateCurrentMessageMasterFromGitHubAsyncResult,
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

	go updateCurrentMessageMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) UpdateCurrentMessageMasterFromGitHub(
	request *UpdateCurrentMessageMasterFromGitHubRequest,
) (*UpdateCurrentMessageMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentMessageMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentMessageMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalMessageMastersAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalMessageMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalMessageMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalMessageMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalMessageMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalMessageMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalMessageMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DescribeGlobalMessageMastersAsync(
	request *DescribeGlobalMessageMastersRequest,
	callback chan<- DescribeGlobalMessageMastersAsyncResult,
) {
	path := "/{namespaceName}/master/globalMessage"
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

	go describeGlobalMessageMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DescribeGlobalMessageMasters(
	request *DescribeGlobalMessageMastersRequest,
) (*DescribeGlobalMessageMastersResult, error) {
	callback := make(chan DescribeGlobalMessageMastersAsyncResult, 1)
	go p.DescribeGlobalMessageMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGlobalMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGlobalMessageMasterResult
	if asyncResult.Err != nil {
		callback <- CreateGlobalMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGlobalMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) CreateGlobalMessageMasterAsync(
	request *CreateGlobalMessageMasterRequest,
	callback chan<- CreateGlobalMessageMasterAsyncResult,
) {
	path := "/{namespaceName}/master/globalMessage"
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
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReadAcquireActions != nil {
		var _readAcquireActions []interface{}
		for _, item := range request.ReadAcquireActions {
			_readAcquireActions = append(_readAcquireActions, item)
		}
		bodies["readAcquireActions"] = _readAcquireActions
	}
	if request.ExpiresTimeSpan != nil {
		bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.MessageReceptionPeriodEventId != nil && *request.MessageReceptionPeriodEventId != "" {
		bodies["messageReceptionPeriodEventId"] = *request.MessageReceptionPeriodEventId
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

	go createGlobalMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) CreateGlobalMessageMaster(
	request *CreateGlobalMessageMasterRequest,
) (*CreateGlobalMessageMasterResult, error) {
	callback := make(chan CreateGlobalMessageMasterAsyncResult, 1)
	go p.CreateGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalMessageMasterResult
	if asyncResult.Err != nil {
		callback <- GetGlobalMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetGlobalMessageMasterAsync(
	request *GetGlobalMessageMasterRequest,
	callback chan<- GetGlobalMessageMasterAsyncResult,
) {
	path := "/{namespaceName}/master/globalMessage/{globalMessageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
		path = strings.ReplaceAll(path, "{globalMessageName}", core.ToString(*request.GlobalMessageName))
	} else {
		path = strings.ReplaceAll(path, "{globalMessageName}", "null")
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

	go getGlobalMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetGlobalMessageMaster(
	request *GetGlobalMessageMasterRequest,
) (*GetGlobalMessageMasterResult, error) {
	callback := make(chan GetGlobalMessageMasterAsyncResult, 1)
	go p.GetGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGlobalMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGlobalMessageMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateGlobalMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGlobalMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) UpdateGlobalMessageMasterAsync(
	request *UpdateGlobalMessageMasterRequest,
	callback chan<- UpdateGlobalMessageMasterAsyncResult,
) {
	path := "/{namespaceName}/master/globalMessage/{globalMessageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
		path = strings.ReplaceAll(path, "{globalMessageName}", core.ToString(*request.GlobalMessageName))
	} else {
		path = strings.ReplaceAll(path, "{globalMessageName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ReadAcquireActions != nil {
		var _readAcquireActions []interface{}
		for _, item := range request.ReadAcquireActions {
			_readAcquireActions = append(_readAcquireActions, item)
		}
		bodies["readAcquireActions"] = _readAcquireActions
	}
	if request.ExpiresTimeSpan != nil {
		bodies["expiresTimeSpan"] = request.ExpiresTimeSpan.ToDict()
	}
	if request.ExpiresAt != nil {
		bodies["expiresAt"] = *request.ExpiresAt
	}
	if request.MessageReceptionPeriodEventId != nil && *request.MessageReceptionPeriodEventId != "" {
		bodies["messageReceptionPeriodEventId"] = *request.MessageReceptionPeriodEventId
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

	go updateGlobalMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) UpdateGlobalMessageMaster(
	request *UpdateGlobalMessageMasterRequest,
) (*UpdateGlobalMessageMasterResult, error) {
	callback := make(chan UpdateGlobalMessageMasterAsyncResult, 1)
	go p.UpdateGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGlobalMessageMasterAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGlobalMessageMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGlobalMessageMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGlobalMessageMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteGlobalMessageMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGlobalMessageMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGlobalMessageMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DeleteGlobalMessageMasterAsync(
	request *DeleteGlobalMessageMasterRequest,
	callback chan<- DeleteGlobalMessageMasterAsyncResult,
) {
	path := "/{namespaceName}/master/globalMessage/{globalMessageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
		path = strings.ReplaceAll(path, "{globalMessageName}", core.ToString(*request.GlobalMessageName))
	} else {
		path = strings.ReplaceAll(path, "{globalMessageName}", "null")
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

	go deleteGlobalMessageMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteGlobalMessageMaster(
	request *DeleteGlobalMessageMasterRequest,
) (*DeleteGlobalMessageMasterResult, error) {
	callback := make(chan DeleteGlobalMessageMasterAsyncResult, 1)
	go p.DeleteGlobalMessageMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGlobalMessagesAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGlobalMessagesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGlobalMessagesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGlobalMessagesResult
	if asyncResult.Err != nil {
		callback <- DescribeGlobalMessagesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGlobalMessagesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGlobalMessagesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DescribeGlobalMessagesAsync(
	request *DescribeGlobalMessagesRequest,
	callback chan<- DescribeGlobalMessagesAsyncResult,
) {
	path := "/{namespaceName}/globalMessage"
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

	go describeGlobalMessagesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DescribeGlobalMessages(
	request *DescribeGlobalMessagesRequest,
) (*DescribeGlobalMessagesResult, error) {
	callback := make(chan DescribeGlobalMessagesAsyncResult, 1)
	go p.DescribeGlobalMessagesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGlobalMessageAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetGlobalMessageAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGlobalMessageAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGlobalMessageResult
	if asyncResult.Err != nil {
		callback <- GetGlobalMessageAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGlobalMessageAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGlobalMessageAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetGlobalMessageAsync(
	request *GetGlobalMessageRequest,
	callback chan<- GetGlobalMessageAsyncResult,
) {
	path := "/{namespaceName}/globalMessage/{globalMessageName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GlobalMessageName != nil && *request.GlobalMessageName != "" {
		path = strings.ReplaceAll(path, "{globalMessageName}", core.ToString(*request.GlobalMessageName))
	} else {
		path = strings.ReplaceAll(path, "{globalMessageName}", "null")
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

	go getGlobalMessageAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetGlobalMessage(
	request *GetGlobalMessageRequest,
) (*GetGlobalMessageResult, error) {
	callback := make(chan GetGlobalMessageAsyncResult, 1)
	go p.GetGlobalMessageAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReceivedByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- GetReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceivedByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetReceivedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceivedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) GetReceivedByUserIdAsync(
	request *GetReceivedByUserIdRequest,
	callback chan<- GetReceivedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/received"
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

	go getReceivedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) GetReceivedByUserId(
	request *GetReceivedByUserIdRequest,
) (*GetReceivedByUserIdResult, error) {
	callback := make(chan GetReceivedByUserIdAsyncResult, 1)
	go p.GetReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateReceivedByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateReceivedByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateReceivedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateReceivedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) UpdateReceivedByUserIdAsync(
	request *UpdateReceivedByUserIdRequest,
	callback chan<- UpdateReceivedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/received"
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
	if request.ReceivedGlobalMessageNames != nil {
		var _receivedGlobalMessageNames []interface{}
		for _, item := range request.ReceivedGlobalMessageNames {
			_receivedGlobalMessageNames = append(_receivedGlobalMessageNames, item)
		}
		bodies["receivedGlobalMessageNames"] = _receivedGlobalMessageNames
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

	go updateReceivedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) UpdateReceivedByUserId(
	request *UpdateReceivedByUserIdRequest,
) (*UpdateReceivedByUserIdResult, error) {
	callback := make(chan UpdateReceivedByUserIdAsyncResult, 1)
	go p.UpdateReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteReceivedByUserIdAsyncHandler(
	client Gs2InboxRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteReceivedByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteReceivedByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteReceivedByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteReceivedByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteReceivedByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteReceivedByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2InboxRestClient) DeleteReceivedByUserIdAsync(
	request *DeleteReceivedByUserIdRequest,
	callback chan<- DeleteReceivedByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/received"
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

	go deleteReceivedByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("inbox").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2InboxRestClient) DeleteReceivedByUserId(
	request *DeleteReceivedByUserIdRequest,
) (*DeleteReceivedByUserIdResult, error) {
	callback := make(chan DeleteReceivedByUserIdAsyncResult, 1)
	go p.DeleteReceivedByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
