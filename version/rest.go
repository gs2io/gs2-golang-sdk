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

package version

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2VersionRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2VersionRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DescribeNamespaces(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) CreateNamespaceAsync(
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
	if request.AssumeUserId != nil && *request.AssumeUserId != "" {
		bodies["assumeUserId"] = *request.AssumeUserId
	}
	if request.AcceptVersionScript != nil {
		bodies["acceptVersionScript"] = request.AcceptVersionScript.ToDict()
	}
	if request.CheckVersionTriggerScriptId != nil && *request.CheckVersionTriggerScriptId != "" {
		bodies["checkVersionTriggerScriptId"] = *request.CheckVersionTriggerScriptId
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CreateNamespace(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetNamespaceStatus(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetNamespace(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) UpdateNamespaceAsync(
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
	if request.AssumeUserId != nil && *request.AssumeUserId != "" {
		bodies["assumeUserId"] = *request.AssumeUserId
	}
	if request.AcceptVersionScript != nil {
		bodies["acceptVersionScript"] = request.AcceptVersionScript.ToDict()
	}
	if request.CheckVersionTriggerScriptId != nil && *request.CheckVersionTriggerScriptId != "" {
		bodies["checkVersionTriggerScriptId"] = *request.CheckVersionTriggerScriptId
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) UpdateNamespace(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DeleteNamespace(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DumpUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CheckDumpUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CleanUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CheckCleanUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) PrepareImportUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) ImportUserDataByUserId(
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
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CheckImportUserDataByUserId(
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

func describeVersionModelMastersAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeVersionModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeVersionModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeVersionModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeVersionModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeVersionModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeVersionModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DescribeVersionModelMastersAsync(
	request *DescribeVersionModelMastersRequest,
	callback chan<- DescribeVersionModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/version"
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

	go describeVersionModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DescribeVersionModelMasters(
	request *DescribeVersionModelMastersRequest,
) (*DescribeVersionModelMastersResult, error) {
	callback := make(chan DescribeVersionModelMastersAsyncResult, 1)
	go p.DescribeVersionModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createVersionModelMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- CreateVersionModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateVersionModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateVersionModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateVersionModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateVersionModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateVersionModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) CreateVersionModelMasterAsync(
	request *CreateVersionModelMasterRequest,
	callback chan<- CreateVersionModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/version"
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
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.CurrentVersion != nil {
		bodies["currentVersion"] = request.CurrentVersion.ToDict()
	}
	if request.WarningVersion != nil {
		bodies["warningVersion"] = request.WarningVersion.ToDict()
	}
	if request.ErrorVersion != nil {
		bodies["errorVersion"] = request.ErrorVersion.ToDict()
	}
	if request.ScheduleVersions != nil {
		var _scheduleVersions []interface{}
		for _, item := range request.ScheduleVersions {
			_scheduleVersions = append(_scheduleVersions, item)
		}
		bodies["scheduleVersions"] = _scheduleVersions
	}
	if request.NeedSignature != nil {
		bodies["needSignature"] = *request.NeedSignature
	}
	if request.SignatureKeyId != nil && *request.SignatureKeyId != "" {
		bodies["signatureKeyId"] = *request.SignatureKeyId
	}
	if request.ApproveRequirement != nil && *request.ApproveRequirement != "" {
		bodies["approveRequirement"] = *request.ApproveRequirement
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

	go createVersionModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CreateVersionModelMaster(
	request *CreateVersionModelMasterRequest,
) (*CreateVersionModelMasterResult, error) {
	callback := make(chan CreateVersionModelMasterAsyncResult, 1)
	go p.CreateVersionModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getVersionModelMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- GetVersionModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetVersionModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetVersionModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetVersionModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetVersionModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetVersionModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) GetVersionModelMasterAsync(
	request *GetVersionModelMasterRequest,
	callback chan<- GetVersionModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/version/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go getVersionModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetVersionModelMaster(
	request *GetVersionModelMasterRequest,
) (*GetVersionModelMasterResult, error) {
	callback := make(chan GetVersionModelMasterAsyncResult, 1)
	go p.GetVersionModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateVersionModelMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateVersionModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateVersionModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateVersionModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateVersionModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateVersionModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateVersionModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) UpdateVersionModelMasterAsync(
	request *UpdateVersionModelMasterRequest,
	callback chan<- UpdateVersionModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/version/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.Scope != nil && *request.Scope != "" {
		bodies["scope"] = *request.Scope
	}
	if request.Type != nil && *request.Type != "" {
		bodies["type"] = *request.Type
	}
	if request.CurrentVersion != nil {
		bodies["currentVersion"] = request.CurrentVersion.ToDict()
	}
	if request.WarningVersion != nil {
		bodies["warningVersion"] = request.WarningVersion.ToDict()
	}
	if request.ErrorVersion != nil {
		bodies["errorVersion"] = request.ErrorVersion.ToDict()
	}
	if request.ScheduleVersions != nil {
		var _scheduleVersions []interface{}
		for _, item := range request.ScheduleVersions {
			_scheduleVersions = append(_scheduleVersions, item)
		}
		bodies["scheduleVersions"] = _scheduleVersions
	}
	if request.NeedSignature != nil {
		bodies["needSignature"] = *request.NeedSignature
	}
	if request.SignatureKeyId != nil && *request.SignatureKeyId != "" {
		bodies["signatureKeyId"] = *request.SignatureKeyId
	}
	if request.ApproveRequirement != nil && *request.ApproveRequirement != "" {
		bodies["approveRequirement"] = *request.ApproveRequirement
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

	go updateVersionModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) UpdateVersionModelMaster(
	request *UpdateVersionModelMasterRequest,
) (*UpdateVersionModelMasterResult, error) {
	callback := make(chan UpdateVersionModelMasterAsyncResult, 1)
	go p.UpdateVersionModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteVersionModelMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteVersionModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteVersionModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteVersionModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteVersionModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteVersionModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteVersionModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DeleteVersionModelMasterAsync(
	request *DeleteVersionModelMasterRequest,
	callback chan<- DeleteVersionModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/version/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go deleteVersionModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DeleteVersionModelMaster(
	request *DeleteVersionModelMasterRequest,
) (*DeleteVersionModelMasterResult, error) {
	callback := make(chan DeleteVersionModelMasterAsyncResult, 1)
	go p.DeleteVersionModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeVersionModelsAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeVersionModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeVersionModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeVersionModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeVersionModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeVersionModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeVersionModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DescribeVersionModelsAsync(
	request *DescribeVersionModelsRequest,
	callback chan<- DescribeVersionModelsAsyncResult,
) {
	path := "/{namespaceName}/version"
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

	go describeVersionModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DescribeVersionModels(
	request *DescribeVersionModelsRequest,
) (*DescribeVersionModelsResult, error) {
	callback := make(chan DescribeVersionModelsAsyncResult, 1)
	go p.DescribeVersionModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getVersionModelAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- GetVersionModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetVersionModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetVersionModelResult
	if asyncResult.Err != nil {
		callback <- GetVersionModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetVersionModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetVersionModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) GetVersionModelAsync(
	request *GetVersionModelRequest,
	callback chan<- GetVersionModelAsyncResult,
) {
	path := "/{namespaceName}/version/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go getVersionModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetVersionModel(
	request *GetVersionModelRequest,
) (*GetVersionModelResult, error) {
	callback := make(chan GetVersionModelAsyncResult, 1)
	go p.GetVersionModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAcceptVersionsAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAcceptVersionsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAcceptVersionsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAcceptVersionsResult
	if asyncResult.Err != nil {
		callback <- DescribeAcceptVersionsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAcceptVersionsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAcceptVersionsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DescribeAcceptVersionsAsync(
	request *DescribeAcceptVersionsRequest,
	callback chan<- DescribeAcceptVersionsAsyncResult,
) {
	path := "/{namespaceName}/user/me/acceptVersion"
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

	go describeAcceptVersionsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DescribeAcceptVersions(
	request *DescribeAcceptVersionsRequest,
) (*DescribeAcceptVersionsResult, error) {
	callback := make(chan DescribeAcceptVersionsAsyncResult, 1)
	go p.DescribeAcceptVersionsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeAcceptVersionsByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAcceptVersionsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAcceptVersionsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAcceptVersionsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeAcceptVersionsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAcceptVersionsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAcceptVersionsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DescribeAcceptVersionsByUserIdAsync(
	request *DescribeAcceptVersionsByUserIdRequest,
	callback chan<- DescribeAcceptVersionsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/acceptVersion"
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

	go describeAcceptVersionsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DescribeAcceptVersionsByUserId(
	request *DescribeAcceptVersionsByUserIdRequest,
) (*DescribeAcceptVersionsByUserIdResult, error) {
	callback := make(chan DescribeAcceptVersionsByUserIdAsyncResult, 1)
	go p.DescribeAcceptVersionsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acceptAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- AcceptAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "version.accept.version.invalid" {
				asyncResult.Err = gs2err.SetClientError(AcceptVersionInvalid{})
			}
		}
		callback <- AcceptAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcceptAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) AcceptAsync(
	request *AcceptRequest,
	callback chan<- AcceptAsyncResult,
) {
	path := "/{namespaceName}/user/me/acceptVersion"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VersionName != nil && *request.VersionName != "" {
		bodies["versionName"] = *request.VersionName
	}
	if request.Version != nil {
		bodies["version"] = request.Version.ToDict()
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

	go acceptAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) Accept(
	request *AcceptRequest,
) (*AcceptResult, error) {
	callback := make(chan AcceptAsyncResult, 1)
	go p.AcceptAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acceptByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- AcceptByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "version.accept.version.invalid" {
				asyncResult.Err = gs2err.SetClientError(AcceptVersionInvalid{})
			}
		}
		callback <- AcceptByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcceptByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) AcceptByUserIdAsync(
	request *AcceptByUserIdRequest,
	callback chan<- AcceptByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/acceptVersion"
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
	if request.VersionName != nil && *request.VersionName != "" {
		bodies["versionName"] = *request.VersionName
	}
	if request.Version != nil {
		bodies["version"] = request.Version.ToDict()
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

	go acceptByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) AcceptByUserId(
	request *AcceptByUserIdRequest,
) (*AcceptByUserIdResult, error) {
	callback := make(chan AcceptByUserIdAsyncResult, 1)
	go p.AcceptByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func rejectAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- RejectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "version.accept.version.invalid" {
				asyncResult.Err = gs2err.SetClientError(AcceptVersionInvalid{})
			}
		}
		callback <- RejectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RejectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) RejectAsync(
	request *RejectRequest,
	callback chan<- RejectAsyncResult,
) {
	path := "/{namespaceName}/user/me/acceptVersion/reject"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VersionName != nil && *request.VersionName != "" {
		bodies["versionName"] = *request.VersionName
	}
	if request.Version != nil {
		bodies["version"] = request.Version.ToDict()
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

	go rejectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) Reject(
	request *RejectRequest,
) (*RejectResult, error) {
	callback := make(chan RejectAsyncResult, 1)
	go p.RejectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func rejectByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- RejectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "version.accept.version.invalid" {
				asyncResult.Err = gs2err.SetClientError(AcceptVersionInvalid{})
			}
		}
		callback <- RejectByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RejectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) RejectByUserIdAsync(
	request *RejectByUserIdRequest,
	callback chan<- RejectByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/acceptVersion/reject"
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
	if request.VersionName != nil && *request.VersionName != "" {
		bodies["versionName"] = *request.VersionName
	}
	if request.Version != nil {
		bodies["version"] = request.Version.ToDict()
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

	go rejectByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) RejectByUserId(
	request *RejectByUserIdRequest,
) (*RejectByUserIdResult, error) {
	callback := make(chan RejectByUserIdAsyncResult, 1)
	go p.RejectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAcceptVersionAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- GetAcceptVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAcceptVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAcceptVersionResult
	if asyncResult.Err != nil {
		callback <- GetAcceptVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAcceptVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAcceptVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) GetAcceptVersionAsync(
	request *GetAcceptVersionRequest,
	callback chan<- GetAcceptVersionAsyncResult,
) {
	path := "/{namespaceName}/user/me/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go getAcceptVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetAcceptVersion(
	request *GetAcceptVersionRequest,
) (*GetAcceptVersionResult, error) {
	callback := make(chan GetAcceptVersionAsyncResult, 1)
	go p.GetAcceptVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAcceptVersionByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- GetAcceptVersionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAcceptVersionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAcceptVersionByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetAcceptVersionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAcceptVersionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAcceptVersionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) GetAcceptVersionByUserIdAsync(
	request *GetAcceptVersionByUserIdRequest,
	callback chan<- GetAcceptVersionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{versionName}"
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
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go getAcceptVersionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetAcceptVersionByUserId(
	request *GetAcceptVersionByUserIdRequest,
) (*GetAcceptVersionByUserIdResult, error) {
	callback := make(chan GetAcceptVersionByUserIdAsyncResult, 1)
	go p.GetAcceptVersionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAcceptVersionAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAcceptVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAcceptVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAcceptVersionResult
	if asyncResult.Err != nil {
		callback <- DeleteAcceptVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteAcceptVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteAcceptVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DeleteAcceptVersionAsync(
	request *DeleteAcceptVersionRequest,
	callback chan<- DeleteAcceptVersionAsyncResult,
) {
	path := "/{namespaceName}/user/me/{versionName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go deleteAcceptVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DeleteAcceptVersion(
	request *DeleteAcceptVersionRequest,
) (*DeleteAcceptVersionResult, error) {
	callback := make(chan DeleteAcceptVersionAsyncResult, 1)
	go p.DeleteAcceptVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAcceptVersionByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAcceptVersionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAcceptVersionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAcceptVersionByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteAcceptVersionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteAcceptVersionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteAcceptVersionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) DeleteAcceptVersionByUserIdAsync(
	request *DeleteAcceptVersionByUserIdRequest,
	callback chan<- DeleteAcceptVersionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/{versionName}"
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
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
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

	go deleteAcceptVersionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) DeleteAcceptVersionByUserId(
	request *DeleteAcceptVersionByUserIdRequest,
) (*DeleteAcceptVersionByUserIdResult, error) {
	callback := make(chan DeleteAcceptVersionByUserIdAsyncResult, 1)
	go p.DeleteAcceptVersionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkVersionAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- CheckVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckVersionResult
	if asyncResult.Err != nil {
		callback <- CheckVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) CheckVersionAsync(
	request *CheckVersionRequest,
	callback chan<- CheckVersionAsyncResult,
) {
	path := "/{namespaceName}/user/me/check"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.TargetVersions != nil {
		var _targetVersions []interface{}
		for _, item := range request.TargetVersions {
			_targetVersions = append(_targetVersions, item)
		}
		bodies["targetVersions"] = _targetVersions
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

	go checkVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CheckVersion(
	request *CheckVersionRequest,
) (*CheckVersionResult, error) {
	callback := make(chan CheckVersionAsyncResult, 1)
	go p.CheckVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func checkVersionByUserIdAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- CheckVersionByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CheckVersionByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CheckVersionByUserIdResult
	if asyncResult.Err != nil {
		callback <- CheckVersionByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CheckVersionByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CheckVersionByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) CheckVersionByUserIdAsync(
	request *CheckVersionByUserIdRequest,
	callback chan<- CheckVersionByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/check"
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
	if request.TargetVersions != nil {
		var _targetVersions []interface{}
		for _, item := range request.TargetVersions {
			_targetVersions = append(_targetVersions, item)
		}
		bodies["targetVersions"] = _targetVersions
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

	go checkVersionByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CheckVersionByUserId(
	request *CheckVersionByUserIdRequest,
) (*CheckVersionByUserIdResult, error) {
	callback := make(chan CheckVersionByUserIdAsyncResult, 1)
	go p.CheckVersionByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func calculateSignatureAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- CalculateSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CalculateSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CalculateSignatureResult
	if asyncResult.Err != nil {
		callback <- CalculateSignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CalculateSignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CalculateSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) CalculateSignatureAsync(
	request *CalculateSignatureRequest,
	callback chan<- CalculateSignatureAsyncResult,
) {
	path := "/{namespaceName}/version/{versionName}/calculate/signature"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.VersionName != nil && *request.VersionName != "" {
		path = strings.ReplaceAll(path, "{versionName}", core.ToString(*request.VersionName))
	} else {
		path = strings.ReplaceAll(path, "{versionName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Version != nil {
		bodies["version"] = request.Version.ToDict()
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

	go calculateSignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) CalculateSignature(
	request *CalculateSignatureRequest,
) (*CalculateSignatureResult, error) {
	callback := make(chan CalculateSignatureAsyncResult, 1)
	go p.CalculateSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2VersionRestClient,
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

func (p Gs2VersionRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) ExportMaster(
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

func getCurrentVersionMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentVersionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentVersionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentVersionMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentVersionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentVersionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentVersionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) GetCurrentVersionMasterAsync(
	request *GetCurrentVersionMasterRequest,
	callback chan<- GetCurrentVersionMasterAsyncResult,
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

	go getCurrentVersionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) GetCurrentVersionMaster(
	request *GetCurrentVersionMasterRequest,
) (*GetCurrentVersionMasterResult, error) {
	callback := make(chan GetCurrentVersionMasterAsyncResult, 1)
	go p.GetCurrentVersionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentVersionMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentVersionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentVersionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentVersionMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentVersionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentVersionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentVersionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) PreUpdateCurrentVersionMasterAsync(
	request *PreUpdateCurrentVersionMasterRequest,
	callback chan<- PreUpdateCurrentVersionMasterAsyncResult,
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

	go preUpdateCurrentVersionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) PreUpdateCurrentVersionMaster(
	request *PreUpdateCurrentVersionMasterRequest,
) (*PreUpdateCurrentVersionMasterResult, error) {
	callback := make(chan PreUpdateCurrentVersionMasterAsyncResult, 1)
	go p.PreUpdateCurrentVersionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentVersionMasterAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentVersionMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentVersionMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentVersionMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentVersionMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentVersionMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentVersionMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) UpdateCurrentVersionMasterAsync(
	request *UpdateCurrentVersionMasterRequest,
	callback chan<- UpdateCurrentVersionMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentVersionMaster(
			&PreUpdateCurrentVersionMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentVersionMasterAsyncResult{
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
				callback <- UpdateCurrentVersionMasterAsyncResult{
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

	go updateCurrentVersionMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) UpdateCurrentVersionMaster(
	request *UpdateCurrentVersionMasterRequest,
) (*UpdateCurrentVersionMasterResult, error) {
	callback := make(chan UpdateCurrentVersionMasterAsyncResult, 1)
	go p.UpdateCurrentVersionMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentVersionMasterFromGitHubAsyncHandler(
	client Gs2VersionRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentVersionMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentVersionMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentVersionMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentVersionMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentVersionMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentVersionMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2VersionRestClient) UpdateCurrentVersionMasterFromGitHubAsync(
	request *UpdateCurrentVersionMasterFromGitHubRequest,
	callback chan<- UpdateCurrentVersionMasterFromGitHubAsyncResult,
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

	go updateCurrentVersionMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("version").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2VersionRestClient) UpdateCurrentVersionMasterFromGitHub(
	request *UpdateCurrentVersionMasterFromGitHubRequest,
) (*UpdateCurrentVersionMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentVersionMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentVersionMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
