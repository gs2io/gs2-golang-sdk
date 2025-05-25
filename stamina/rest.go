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

package stamina

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2StaminaRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2StaminaRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeNamespaces(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) CreateNamespaceAsync(
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
	if request.OverflowTriggerScript != nil && *request.OverflowTriggerScript != "" {
		bodies["overflowTriggerScript"] = *request.OverflowTriggerScript
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CreateNamespace(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetNamespaceStatus(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetNamespace(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) UpdateNamespaceAsync(
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
	if request.OverflowTriggerScript != nil && *request.OverflowTriggerScript != "" {
		bodies["overflowTriggerScript"] = *request.OverflowTriggerScript
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateNamespace(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteNamespace(
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

func getServiceVersionAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetServiceVersionResult
	if asyncResult.Err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	path := "/system/version"

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

	go getServiceVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetServiceVersion(
	request *GetServiceVersionRequest,
) (*GetServiceVersionResult, error) {
	callback := make(chan GetServiceVersionAsyncResult, 1)
	go p.GetServiceVersionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func dumpUserDataByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DumpUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CheckDumpUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CleanUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CheckCleanUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) PrepareImportUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ImportUserDataByUserId(
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
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CheckImportUserDataByUserId(
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

func describeStaminaModelMastersAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminaModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeStaminaModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminaModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeStaminaModelMastersAsync(
	request *DescribeStaminaModelMastersRequest,
	callback chan<- DescribeStaminaModelMastersAsyncResult,
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

	go describeStaminaModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeStaminaModelMasters(
	request *DescribeStaminaModelMastersRequest,
) (*DescribeStaminaModelMastersResult, error) {
	callback := make(chan DescribeStaminaModelMastersAsyncResult, 1)
	go p.DescribeStaminaModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStaminaModelMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- CreateStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStaminaModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateStaminaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStaminaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) CreateStaminaModelMasterAsync(
	request *CreateStaminaModelMasterRequest,
	callback chan<- CreateStaminaModelMasterAsyncResult,
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
	if request.RecoverIntervalMinutes != nil {
		bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
	}
	if request.RecoverValue != nil {
		bodies["recoverValue"] = *request.RecoverValue
	}
	if request.InitialCapacity != nil {
		bodies["initialCapacity"] = *request.InitialCapacity
	}
	if request.IsOverflow != nil {
		bodies["isOverflow"] = *request.IsOverflow
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
	}
	if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
		bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
	}
	if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
		bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
	}
	if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
		bodies["recoverValueTableName"] = *request.RecoverValueTableName
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

	go createStaminaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CreateStaminaModelMaster(
	request *CreateStaminaModelMasterRequest,
) (*CreateStaminaModelMasterResult, error) {
	callback := make(chan CreateStaminaModelMasterAsyncResult, 1)
	go p.CreateStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStaminaModelMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetStaminaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetStaminaModelMasterAsync(
	request *GetStaminaModelMasterRequest,
	callback chan<- GetStaminaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go getStaminaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetStaminaModelMaster(
	request *GetStaminaModelMasterRequest,
) (*GetStaminaModelMasterResult, error) {
	callback := make(chan GetStaminaModelMasterAsyncResult, 1)
	go p.GetStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStaminaModelMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStaminaModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateStaminaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStaminaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateStaminaModelMasterAsync(
	request *UpdateStaminaModelMasterRequest,
	callback chan<- UpdateStaminaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.RecoverIntervalMinutes != nil {
		bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
	}
	if request.RecoverValue != nil {
		bodies["recoverValue"] = *request.RecoverValue
	}
	if request.InitialCapacity != nil {
		bodies["initialCapacity"] = *request.InitialCapacity
	}
	if request.IsOverflow != nil {
		bodies["isOverflow"] = *request.IsOverflow
	}
	if request.MaxCapacity != nil {
		bodies["maxCapacity"] = *request.MaxCapacity
	}
	if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
		bodies["maxStaminaTableName"] = *request.MaxStaminaTableName
	}
	if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
		bodies["recoverIntervalTableName"] = *request.RecoverIntervalTableName
	}
	if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
		bodies["recoverValueTableName"] = *request.RecoverValueTableName
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

	go updateStaminaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateStaminaModelMaster(
	request *UpdateStaminaModelMasterRequest,
) (*UpdateStaminaModelMasterResult, error) {
	callback := make(chan UpdateStaminaModelMasterAsyncResult, 1)
	go p.UpdateStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStaminaModelMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStaminaModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStaminaModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStaminaModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteStaminaModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStaminaModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStaminaModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DeleteStaminaModelMasterAsync(
	request *DeleteStaminaModelMasterRequest,
	callback chan<- DeleteStaminaModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go deleteStaminaModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteStaminaModelMaster(
	request *DeleteStaminaModelMasterRequest,
) (*DeleteStaminaModelMasterResult, error) {
	callback := make(chan DeleteStaminaModelMasterAsyncResult, 1)
	go p.DeleteStaminaModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeMaxStaminaTableMastersAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMaxStaminaTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMaxStaminaTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMaxStaminaTableMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeMaxStaminaTableMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMaxStaminaTableMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMaxStaminaTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeMaxStaminaTableMastersAsync(
	request *DescribeMaxStaminaTableMastersRequest,
	callback chan<- DescribeMaxStaminaTableMastersAsyncResult,
) {
	path := "/{namespaceName}/master/maxStaminaTable"
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

	go describeMaxStaminaTableMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeMaxStaminaTableMasters(
	request *DescribeMaxStaminaTableMastersRequest,
) (*DescribeMaxStaminaTableMastersResult, error) {
	callback := make(chan DescribeMaxStaminaTableMastersAsyncResult, 1)
	go p.DescribeMaxStaminaTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createMaxStaminaTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- CreateMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateMaxStaminaTableMasterResult
	if asyncResult.Err != nil {
		callback <- CreateMaxStaminaTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateMaxStaminaTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) CreateMaxStaminaTableMasterAsync(
	request *CreateMaxStaminaTableMasterRequest,
	callback chan<- CreateMaxStaminaTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/maxStaminaTable"
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
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go createMaxStaminaTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CreateMaxStaminaTableMaster(
	request *CreateMaxStaminaTableMasterRequest,
) (*CreateMaxStaminaTableMasterResult, error) {
	callback := make(chan CreateMaxStaminaTableMasterAsyncResult, 1)
	go p.CreateMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getMaxStaminaTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMaxStaminaTableMasterResult
	if asyncResult.Err != nil {
		callback <- GetMaxStaminaTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMaxStaminaTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetMaxStaminaTableMasterAsync(
	request *GetMaxStaminaTableMasterRequest,
	callback chan<- GetMaxStaminaTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/maxStaminaTable/{maxStaminaTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", core.ToString(*request.MaxStaminaTableName))
	} else {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", "null")
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

	go getMaxStaminaTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetMaxStaminaTableMaster(
	request *GetMaxStaminaTableMasterRequest,
) (*GetMaxStaminaTableMasterResult, error) {
	callback := make(chan GetMaxStaminaTableMasterAsyncResult, 1)
	go p.GetMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMaxStaminaTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMaxStaminaTableMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateMaxStaminaTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMaxStaminaTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateMaxStaminaTableMasterAsync(
	request *UpdateMaxStaminaTableMasterRequest,
	callback chan<- UpdateMaxStaminaTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/maxStaminaTable/{maxStaminaTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", core.ToString(*request.MaxStaminaTableName))
	} else {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go updateMaxStaminaTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateMaxStaminaTableMaster(
	request *UpdateMaxStaminaTableMasterRequest,
) (*UpdateMaxStaminaTableMasterResult, error) {
	callback := make(chan UpdateMaxStaminaTableMasterAsyncResult, 1)
	go p.UpdateMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMaxStaminaTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMaxStaminaTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMaxStaminaTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMaxStaminaTableMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteMaxStaminaTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMaxStaminaTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMaxStaminaTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DeleteMaxStaminaTableMasterAsync(
	request *DeleteMaxStaminaTableMasterRequest,
	callback chan<- DeleteMaxStaminaTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/maxStaminaTable/{maxStaminaTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MaxStaminaTableName != nil && *request.MaxStaminaTableName != "" {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", core.ToString(*request.MaxStaminaTableName))
	} else {
		path = strings.ReplaceAll(path, "{maxStaminaTableName}", "null")
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

	go deleteMaxStaminaTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteMaxStaminaTableMaster(
	request *DeleteMaxStaminaTableMasterRequest,
) (*DeleteMaxStaminaTableMasterResult, error) {
	callback := make(chan DeleteMaxStaminaTableMasterAsyncResult, 1)
	go p.DeleteMaxStaminaTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRecoverIntervalTableMastersAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRecoverIntervalTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRecoverIntervalTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRecoverIntervalTableMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeRecoverIntervalTableMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRecoverIntervalTableMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRecoverIntervalTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeRecoverIntervalTableMastersAsync(
	request *DescribeRecoverIntervalTableMastersRequest,
	callback chan<- DescribeRecoverIntervalTableMastersAsyncResult,
) {
	path := "/{namespaceName}/master/recoverIntervalTable"
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

	go describeRecoverIntervalTableMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeRecoverIntervalTableMasters(
	request *DescribeRecoverIntervalTableMastersRequest,
) (*DescribeRecoverIntervalTableMastersResult, error) {
	callback := make(chan DescribeRecoverIntervalTableMastersAsyncResult, 1)
	go p.DescribeRecoverIntervalTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRecoverIntervalTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRecoverIntervalTableMasterResult
	if asyncResult.Err != nil {
		callback <- CreateRecoverIntervalTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRecoverIntervalTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) CreateRecoverIntervalTableMasterAsync(
	request *CreateRecoverIntervalTableMasterRequest,
	callback chan<- CreateRecoverIntervalTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverIntervalTable"
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
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go createRecoverIntervalTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CreateRecoverIntervalTableMaster(
	request *CreateRecoverIntervalTableMasterRequest,
) (*CreateRecoverIntervalTableMasterResult, error) {
	callback := make(chan CreateRecoverIntervalTableMasterAsyncResult, 1)
	go p.CreateRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRecoverIntervalTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRecoverIntervalTableMasterResult
	if asyncResult.Err != nil {
		callback <- GetRecoverIntervalTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRecoverIntervalTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetRecoverIntervalTableMasterAsync(
	request *GetRecoverIntervalTableMasterRequest,
	callback chan<- GetRecoverIntervalTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverIntervalTable/{recoverIntervalTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", core.ToString(*request.RecoverIntervalTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", "null")
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

	go getRecoverIntervalTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetRecoverIntervalTableMaster(
	request *GetRecoverIntervalTableMasterRequest,
) (*GetRecoverIntervalTableMasterResult, error) {
	callback := make(chan GetRecoverIntervalTableMasterAsyncResult, 1)
	go p.GetRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRecoverIntervalTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRecoverIntervalTableMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateRecoverIntervalTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRecoverIntervalTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateRecoverIntervalTableMasterAsync(
	request *UpdateRecoverIntervalTableMasterRequest,
	callback chan<- UpdateRecoverIntervalTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverIntervalTable/{recoverIntervalTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", core.ToString(*request.RecoverIntervalTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go updateRecoverIntervalTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateRecoverIntervalTableMaster(
	request *UpdateRecoverIntervalTableMasterRequest,
) (*UpdateRecoverIntervalTableMasterResult, error) {
	callback := make(chan UpdateRecoverIntervalTableMasterAsyncResult, 1)
	go p.UpdateRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRecoverIntervalTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRecoverIntervalTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRecoverIntervalTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRecoverIntervalTableMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteRecoverIntervalTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRecoverIntervalTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRecoverIntervalTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DeleteRecoverIntervalTableMasterAsync(
	request *DeleteRecoverIntervalTableMasterRequest,
	callback chan<- DeleteRecoverIntervalTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverIntervalTable/{recoverIntervalTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverIntervalTableName != nil && *request.RecoverIntervalTableName != "" {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", core.ToString(*request.RecoverIntervalTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverIntervalTableName}", "null")
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

	go deleteRecoverIntervalTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteRecoverIntervalTableMaster(
	request *DeleteRecoverIntervalTableMasterRequest,
) (*DeleteRecoverIntervalTableMasterResult, error) {
	callback := make(chan DeleteRecoverIntervalTableMasterAsyncResult, 1)
	go p.DeleteRecoverIntervalTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRecoverValueTableMastersAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRecoverValueTableMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRecoverValueTableMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRecoverValueTableMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeRecoverValueTableMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRecoverValueTableMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRecoverValueTableMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeRecoverValueTableMastersAsync(
	request *DescribeRecoverValueTableMastersRequest,
	callback chan<- DescribeRecoverValueTableMastersAsyncResult,
) {
	path := "/{namespaceName}/master/recoverValueTable"
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

	go describeRecoverValueTableMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeRecoverValueTableMasters(
	request *DescribeRecoverValueTableMastersRequest,
) (*DescribeRecoverValueTableMastersResult, error) {
	callback := make(chan DescribeRecoverValueTableMastersAsyncResult, 1)
	go p.DescribeRecoverValueTableMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRecoverValueTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRecoverValueTableMasterResult
	if asyncResult.Err != nil {
		callback <- CreateRecoverValueTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateRecoverValueTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) CreateRecoverValueTableMasterAsync(
	request *CreateRecoverValueTableMasterRequest,
	callback chan<- CreateRecoverValueTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverValueTable"
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
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go createRecoverValueTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) CreateRecoverValueTableMaster(
	request *CreateRecoverValueTableMasterRequest,
) (*CreateRecoverValueTableMasterResult, error) {
	callback := make(chan CreateRecoverValueTableMasterAsyncResult, 1)
	go p.CreateRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRecoverValueTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRecoverValueTableMasterResult
	if asyncResult.Err != nil {
		callback <- GetRecoverValueTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRecoverValueTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetRecoverValueTableMasterAsync(
	request *GetRecoverValueTableMasterRequest,
	callback chan<- GetRecoverValueTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverValueTable/{recoverValueTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", core.ToString(*request.RecoverValueTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", "null")
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

	go getRecoverValueTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetRecoverValueTableMaster(
	request *GetRecoverValueTableMasterRequest,
) (*GetRecoverValueTableMasterResult, error) {
	callback := make(chan GetRecoverValueTableMasterAsyncResult, 1)
	go p.GetRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRecoverValueTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRecoverValueTableMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateRecoverValueTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateRecoverValueTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateRecoverValueTableMasterAsync(
	request *UpdateRecoverValueTableMasterRequest,
	callback chan<- UpdateRecoverValueTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverValueTable/{recoverValueTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", core.ToString(*request.RecoverValueTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ExperienceModelId != nil && *request.ExperienceModelId != "" {
		bodies["experienceModelId"] = *request.ExperienceModelId
	}
	if request.Values != nil {
		var _values []interface{}
		for _, item := range request.Values {
			_values = append(_values, item)
		}
		bodies["values"] = _values
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

	go updateRecoverValueTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateRecoverValueTableMaster(
	request *UpdateRecoverValueTableMasterRequest,
) (*UpdateRecoverValueTableMasterResult, error) {
	callback := make(chan UpdateRecoverValueTableMasterAsyncResult, 1)
	go p.UpdateRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRecoverValueTableMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRecoverValueTableMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRecoverValueTableMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRecoverValueTableMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteRecoverValueTableMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRecoverValueTableMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRecoverValueTableMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DeleteRecoverValueTableMasterAsync(
	request *DeleteRecoverValueTableMasterRequest,
	callback chan<- DeleteRecoverValueTableMasterAsyncResult,
) {
	path := "/{namespaceName}/master/recoverValueTable/{recoverValueTableName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RecoverValueTableName != nil && *request.RecoverValueTableName != "" {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", core.ToString(*request.RecoverValueTableName))
	} else {
		path = strings.ReplaceAll(path, "{recoverValueTableName}", "null")
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

	go deleteRecoverValueTableMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteRecoverValueTableMaster(
	request *DeleteRecoverValueTableMasterRequest,
) (*DeleteRecoverValueTableMasterResult, error) {
	callback := make(chan DeleteRecoverValueTableMasterAsyncResult, 1)
	go p.DeleteRecoverValueTableMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2StaminaRestClient,
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

func (p Gs2StaminaRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ExportMaster(
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

func getCurrentStaminaMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentStaminaMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentStaminaMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentStaminaMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentStaminaMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentStaminaMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentStaminaMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetCurrentStaminaMasterAsync(
	request *GetCurrentStaminaMasterRequest,
	callback chan<- GetCurrentStaminaMasterAsyncResult,
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

	go getCurrentStaminaMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetCurrentStaminaMaster(
	request *GetCurrentStaminaMasterRequest,
) (*GetCurrentStaminaMasterResult, error) {
	callback := make(chan GetCurrentStaminaMasterAsyncResult, 1)
	go p.GetCurrentStaminaMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentStaminaMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentStaminaMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentStaminaMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentStaminaMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentStaminaMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentStaminaMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentStaminaMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) PreUpdateCurrentStaminaMasterAsync(
	request *PreUpdateCurrentStaminaMasterRequest,
	callback chan<- PreUpdateCurrentStaminaMasterAsyncResult,
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

	go preUpdateCurrentStaminaMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) PreUpdateCurrentStaminaMaster(
	request *PreUpdateCurrentStaminaMasterRequest,
) (*PreUpdateCurrentStaminaMasterResult, error) {
	callback := make(chan PreUpdateCurrentStaminaMasterAsyncResult, 1)
	go p.PreUpdateCurrentStaminaMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentStaminaMasterAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentStaminaMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStaminaMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStaminaMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentStaminaMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStaminaMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentStaminaMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateCurrentStaminaMasterAsync(
	request *UpdateCurrentStaminaMasterRequest,
	callback chan<- UpdateCurrentStaminaMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentStaminaMaster(
			&PreUpdateCurrentStaminaMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentStaminaMasterAsyncResult{
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
				callback <- UpdateCurrentStaminaMasterAsyncResult{
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

	go updateCurrentStaminaMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateCurrentStaminaMaster(
	request *UpdateCurrentStaminaMasterRequest,
) (*UpdateCurrentStaminaMasterResult, error) {
	callback := make(chan UpdateCurrentStaminaMasterAsyncResult, 1)
	go p.UpdateCurrentStaminaMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentStaminaMasterFromGitHubAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentStaminaMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentStaminaMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentStaminaMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateCurrentStaminaMasterFromGitHubAsync(
	request *UpdateCurrentStaminaMasterFromGitHubRequest,
	callback chan<- UpdateCurrentStaminaMasterFromGitHubAsyncResult,
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

	go updateCurrentStaminaMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateCurrentStaminaMasterFromGitHub(
	request *UpdateCurrentStaminaMasterFromGitHubRequest,
) (*UpdateCurrentStaminaMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentStaminaMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentStaminaMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStaminaModelsAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminaModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeStaminaModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminaModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeStaminaModelsAsync(
	request *DescribeStaminaModelsRequest,
	callback chan<- DescribeStaminaModelsAsyncResult,
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

	go describeStaminaModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeStaminaModels(
	request *DescribeStaminaModelsRequest,
) (*DescribeStaminaModelsResult, error) {
	callback := make(chan DescribeStaminaModelsAsyncResult, 1)
	go p.DescribeStaminaModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStaminaModelAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaModelResult
	if asyncResult.Err != nil {
		callback <- GetStaminaModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetStaminaModelAsync(
	request *GetStaminaModelRequest,
	callback chan<- GetStaminaModelAsyncResult,
) {
	path := "/{namespaceName}/model/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go getStaminaModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetStaminaModel(
	request *GetStaminaModelRequest,
) (*GetStaminaModelResult, error) {
	callback := make(chan GetStaminaModelAsyncResult, 1)
	go p.GetStaminaModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStaminasAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminasAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminasAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminasResult
	if asyncResult.Err != nil {
		callback <- DescribeStaminasAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminasAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminasAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeStaminasAsync(
	request *DescribeStaminasRequest,
	callback chan<- DescribeStaminasAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina"
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

	go describeStaminasAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeStaminas(
	request *DescribeStaminasRequest,
) (*DescribeStaminasResult, error) {
	callback := make(chan DescribeStaminasAsyncResult, 1)
	go p.DescribeStaminasAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStaminasByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminasByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminasByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminasByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeStaminasByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminasByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminasByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DescribeStaminasByUserIdAsync(
	request *DescribeStaminasByUserIdRequest,
	callback chan<- DescribeStaminasByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina"
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

	go describeStaminasByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DescribeStaminasByUserId(
	request *DescribeStaminasByUserIdRequest,
) (*DescribeStaminasByUserIdResult, error) {
	callback := make(chan DescribeStaminasByUserIdAsyncResult, 1)
	go p.DescribeStaminasByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStaminaAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaResult
	if asyncResult.Err != nil {
		callback <- GetStaminaAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetStaminaAsync(
	request *GetStaminaRequest,
	callback chan<- GetStaminaAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go getStaminaAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetStamina(
	request *GetStaminaRequest,
) (*GetStaminaResult, error) {
	callback := make(chan GetStaminaAsyncResult, 1)
	go p.GetStaminaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) GetStaminaByUserIdAsync(
	request *GetStaminaByUserIdRequest,
	callback chan<- GetStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go getStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) GetStaminaByUserId(
	request *GetStaminaByUserIdRequest,
) (*GetStaminaByUserIdResult, error) {
	callback := make(chan GetStaminaByUserIdAsyncResult, 1)
	go p.GetStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStaminaByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) UpdateStaminaByUserIdAsync(
	request *UpdateStaminaByUserIdRequest,
	callback chan<- UpdateStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MaxValue != nil {
		bodies["maxValue"] = *request.MaxValue
	}
	if request.RecoverIntervalMinutes != nil {
		bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
	}
	if request.RecoverValue != nil {
		bodies["recoverValue"] = *request.RecoverValue
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

	go updateStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) UpdateStaminaByUserId(
	request *UpdateStaminaByUserIdRequest,
) (*UpdateStaminaByUserIdResult, error) {
	callback := make(chan UpdateStaminaByUserIdAsyncResult, 1)
	go p.UpdateStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeStaminaAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeStaminaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "stamina.stamina.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeStaminaAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeStaminaAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeStaminaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) ConsumeStaminaAsync(
	request *ConsumeStaminaRequest,
	callback chan<- ConsumeStaminaAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeValue != nil {
		bodies["consumeValue"] = *request.ConsumeValue
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

	go consumeStaminaAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ConsumeStamina(
	request *ConsumeStaminaRequest,
) (*ConsumeStaminaResult, error) {
	callback := make(chan ConsumeStaminaAsyncResult, 1)
	go p.ConsumeStaminaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "stamina.stamina.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- ConsumeStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) ConsumeStaminaByUserIdAsync(
	request *ConsumeStaminaByUserIdRequest,
	callback chan<- ConsumeStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/consume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ConsumeValue != nil {
		bodies["consumeValue"] = *request.ConsumeValue
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

	go consumeStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ConsumeStaminaByUserId(
	request *ConsumeStaminaByUserIdRequest,
) (*ConsumeStaminaByUserIdResult, error) {
	callback := make(chan ConsumeStaminaByUserIdAsyncResult, 1)
	go p.ConsumeStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func applyStaminaAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- ApplyStaminaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ApplyStaminaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ApplyStaminaResult
	if asyncResult.Err != nil {
		callback <- ApplyStaminaAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ApplyStaminaAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ApplyStaminaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) ApplyStaminaAsync(
	request *ApplyStaminaRequest,
	callback chan<- ApplyStaminaAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/apply"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go applyStaminaAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ApplyStamina(
	request *ApplyStaminaRequest,
) (*ApplyStaminaResult, error) {
	callback := make(chan ApplyStaminaAsyncResult, 1)
	go p.ApplyStaminaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func applyStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- ApplyStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ApplyStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ApplyStaminaByUserIdResult
	if asyncResult.Err != nil {
		callback <- ApplyStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ApplyStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ApplyStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) ApplyStaminaByUserIdAsync(
	request *ApplyStaminaByUserIdRequest,
	callback chan<- ApplyStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/apply"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go applyStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ApplyStaminaByUserId(
	request *ApplyStaminaByUserIdRequest,
) (*ApplyStaminaByUserIdResult, error) {
	callback := make(chan ApplyStaminaByUserIdAsyncResult, 1)
	go p.ApplyStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func recoverStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- RecoverStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecoverStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecoverStaminaByUserIdResult
	if asyncResult.Err != nil {
		callback <- RecoverStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RecoverStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RecoverStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) RecoverStaminaByUserIdAsync(
	request *RecoverStaminaByUserIdRequest,
	callback chan<- RecoverStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/recover"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RecoverValue != nil {
		bodies["recoverValue"] = *request.RecoverValue
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

	go recoverStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) RecoverStaminaByUserId(
	request *RecoverStaminaByUserIdRequest,
) (*RecoverStaminaByUserIdResult, error) {
	callback := make(chan RecoverStaminaByUserIdAsyncResult, 1)
	go p.RecoverStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func raiseMaxValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- RaiseMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RaiseMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RaiseMaxValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- RaiseMaxValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RaiseMaxValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RaiseMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) RaiseMaxValueByUserIdAsync(
	request *RaiseMaxValueByUserIdRequest,
	callback chan<- RaiseMaxValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/raise"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RaiseValue != nil {
		bodies["raiseValue"] = *request.RaiseValue
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

	go raiseMaxValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) RaiseMaxValueByUserId(
	request *RaiseMaxValueByUserIdRequest,
) (*RaiseMaxValueByUserIdResult, error) {
	callback := make(chan RaiseMaxValueByUserIdAsyncResult, 1)
	go p.RaiseMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaxValueAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaxValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaxValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaxValueResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaxValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaxValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaxValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DecreaseMaxValueAsync(
	request *DecreaseMaxValueRequest,
	callback chan<- DecreaseMaxValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DecreaseValue != nil {
		bodies["decreaseValue"] = *request.DecreaseValue
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

	go decreaseMaxValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DecreaseMaxValue(
	request *DecreaseMaxValueRequest,
) (*DecreaseMaxValueResult, error) {
	callback := make(chan DecreaseMaxValueAsyncResult, 1)
	go p.DecreaseMaxValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaxValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaxValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaxValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaxValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DecreaseMaxValueByUserIdAsync(
	request *DecreaseMaxValueByUserIdRequest,
	callback chan<- DecreaseMaxValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DecreaseValue != nil {
		bodies["decreaseValue"] = *request.DecreaseValue
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

	go decreaseMaxValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DecreaseMaxValueByUserId(
	request *DecreaseMaxValueByUserIdRequest,
) (*DecreaseMaxValueByUserIdResult, error) {
	callback := make(chan DecreaseMaxValueByUserIdAsyncResult, 1)
	go p.DecreaseMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMaxValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetMaxValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaxValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetMaxValueByUserIdAsync(
	request *SetMaxValueByUserIdRequest,
	callback chan<- SetMaxValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.MaxValue != nil {
		bodies["maxValue"] = *request.MaxValue
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

	go setMaxValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetMaxValueByUserId(
	request *SetMaxValueByUserIdRequest,
) (*SetMaxValueByUserIdResult, error) {
	callback := make(chan SetMaxValueByUserIdAsyncResult, 1)
	go p.SetMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverIntervalByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverIntervalByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetRecoverIntervalByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverIntervalByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverIntervalByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverIntervalByUserIdAsync(
	request *SetRecoverIntervalByUserIdRequest,
	callback chan<- SetRecoverIntervalByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/recoverInterval/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RecoverIntervalMinutes != nil {
		bodies["recoverIntervalMinutes"] = *request.RecoverIntervalMinutes
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

	go setRecoverIntervalByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverIntervalByUserId(
	request *SetRecoverIntervalByUserIdRequest,
) (*SetRecoverIntervalByUserIdResult, error) {
	callback := make(chan SetRecoverIntervalByUserIdAsyncResult, 1)
	go p.SetRecoverIntervalByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetRecoverValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverValueByUserIdAsync(
	request *SetRecoverValueByUserIdRequest,
	callback chan<- SetRecoverValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/recoverValue/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RecoverValue != nil {
		bodies["recoverValue"] = *request.RecoverValue
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

	go setRecoverValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverValueByUserId(
	request *SetRecoverValueByUserIdRequest,
) (*SetRecoverValueByUserIdResult, error) {
	callback := make(chan SetRecoverValueByUserIdAsyncResult, 1)
	go p.SetRecoverValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMaxValueByStatusAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetMaxValueByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByStatusResult
	if asyncResult.Err != nil {
		callback <- SetMaxValueByStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaxValueByStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMaxValueByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetMaxValueByStatusAsync(
	request *SetMaxValueByStatusRequest,
	callback chan<- SetMaxValueByStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
		bodies["signedStatusBody"] = *request.SignedStatusBody
	}
	if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
		bodies["signedStatusSignature"] = *request.SignedStatusSignature
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

	go setMaxValueByStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetMaxValueByStatus(
	request *SetMaxValueByStatusRequest,
) (*SetMaxValueByStatusResult, error) {
	callback := make(chan SetMaxValueByStatusAsyncResult, 1)
	go p.SetMaxValueByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverIntervalByStatusAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverIntervalByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByStatusResult
	if asyncResult.Err != nil {
		callback <- SetRecoverIntervalByStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverIntervalByStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverIntervalByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverIntervalByStatusAsync(
	request *SetRecoverIntervalByStatusRequest,
	callback chan<- SetRecoverIntervalByStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/recoverInterval/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
		bodies["signedStatusBody"] = *request.SignedStatusBody
	}
	if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
		bodies["signedStatusSignature"] = *request.SignedStatusSignature
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

	go setRecoverIntervalByStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverIntervalByStatus(
	request *SetRecoverIntervalByStatusRequest,
) (*SetRecoverIntervalByStatusResult, error) {
	callback := make(chan SetRecoverIntervalByStatusAsyncResult, 1)
	go p.SetRecoverIntervalByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverValueByStatusAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverValueByStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByStatusResult
	if asyncResult.Err != nil {
		callback <- SetRecoverValueByStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverValueByStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverValueByStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverValueByStatusAsync(
	request *SetRecoverValueByStatusRequest,
	callback chan<- SetRecoverValueByStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/recoverValue/set"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.SignedStatusBody != nil && *request.SignedStatusBody != "" {
		bodies["signedStatusBody"] = *request.SignedStatusBody
	}
	if request.SignedStatusSignature != nil && *request.SignedStatusSignature != "" {
		bodies["signedStatusSignature"] = *request.SignedStatusSignature
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

	go setRecoverValueByStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverValueByStatus(
	request *SetRecoverValueByStatusRequest,
) (*SetRecoverValueByStatusResult, error) {
	callback := make(chan SetRecoverValueByStatusAsyncResult, 1)
	go p.SetRecoverValueByStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStaminaByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStaminaByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStaminaByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStaminaByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteStaminaByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStaminaByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStaminaByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DeleteStaminaByUserIdAsync(
	request *DeleteStaminaByUserIdRequest,
	callback chan<- DeleteStaminaByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
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

	go deleteStaminaByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DeleteStaminaByUserId(
	request *DeleteStaminaByUserIdRequest,
) (*DeleteStaminaByUserIdResult, error) {
	callback := make(chan DeleteStaminaByUserIdAsyncResult, 1)
	go p.DeleteStaminaByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaValueAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaValueResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaValueAsync(
	request *VerifyStaminaValueRequest,
	callback chan<- VerifyStaminaValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/value/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaValue(
	request *VerifyStaminaValueRequest,
) (*VerifyStaminaValueResult, error) {
	callback := make(chan VerifyStaminaValueAsyncResult, 1)
	go p.VerifyStaminaValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaValueByUserIdAsync(
	request *VerifyStaminaValueByUserIdRequest,
	callback chan<- VerifyStaminaValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/value/verify/{verifyType}"
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
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaValueByUserId(
	request *VerifyStaminaValueByUserIdRequest,
) (*VerifyStaminaValueByUserIdResult, error) {
	callback := make(chan VerifyStaminaValueByUserIdAsyncResult, 1)
	go p.VerifyStaminaValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaMaxValueAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaMaxValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaMaxValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaMaxValueResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaMaxValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaMaxValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaMaxValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValueAsync(
	request *VerifyStaminaMaxValueRequest,
	callback chan<- VerifyStaminaMaxValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/max/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaMaxValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValue(
	request *VerifyStaminaMaxValueRequest,
) (*VerifyStaminaMaxValueResult, error) {
	callback := make(chan VerifyStaminaMaxValueAsyncResult, 1)
	go p.VerifyStaminaMaxValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaMaxValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaMaxValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaMaxValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaMaxValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaMaxValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaMaxValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaMaxValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValueByUserIdAsync(
	request *VerifyStaminaMaxValueByUserIdRequest,
	callback chan<- VerifyStaminaMaxValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/max/verify/{verifyType}"
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
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaMaxValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValueByUserId(
	request *VerifyStaminaMaxValueByUserIdRequest,
) (*VerifyStaminaMaxValueByUserIdResult, error) {
	callback := make(chan VerifyStaminaMaxValueByUserIdAsyncResult, 1)
	go p.VerifyStaminaMaxValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverIntervalMinutesAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverIntervalMinutesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverIntervalMinutesResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverIntervalMinutesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverIntervalMinutesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutesAsync(
	request *VerifyStaminaRecoverIntervalMinutesRequest,
	callback chan<- VerifyStaminaRecoverIntervalMinutesAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/recover/interval/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaRecoverIntervalMinutesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutes(
	request *VerifyStaminaRecoverIntervalMinutesRequest,
) (*VerifyStaminaRecoverIntervalMinutesResult, error) {
	callback := make(chan VerifyStaminaRecoverIntervalMinutesAsyncResult, 1)
	go p.VerifyStaminaRecoverIntervalMinutesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverIntervalMinutesByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverIntervalMinutesByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutesByUserIdAsync(
	request *VerifyStaminaRecoverIntervalMinutesByUserIdRequest,
	callback chan<- VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/recover/interval/verify/{verifyType}"
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
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaRecoverIntervalMinutesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutesByUserId(
	request *VerifyStaminaRecoverIntervalMinutesByUserIdRequest,
) (*VerifyStaminaRecoverIntervalMinutesByUserIdResult, error) {
	callback := make(chan VerifyStaminaRecoverIntervalMinutesByUserIdAsyncResult, 1)
	go p.VerifyStaminaRecoverIntervalMinutesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverValueAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverValueResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValueAsync(
	request *VerifyStaminaRecoverValueRequest,
	callback chan<- VerifyStaminaRecoverValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/recover/value/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaRecoverValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValue(
	request *VerifyStaminaRecoverValueRequest,
) (*VerifyStaminaRecoverValueResult, error) {
	callback := make(chan VerifyStaminaRecoverValueAsyncResult, 1)
	go p.VerifyStaminaRecoverValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValueByUserIdAsync(
	request *VerifyStaminaRecoverValueByUserIdRequest,
	callback chan<- VerifyStaminaRecoverValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/recover/value/verify/{verifyType}"
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
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaRecoverValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValueByUserId(
	request *VerifyStaminaRecoverValueByUserIdRequest,
) (*VerifyStaminaRecoverValueByUserIdResult, error) {
	callback := make(chan VerifyStaminaRecoverValueByUserIdAsyncResult, 1)
	go p.VerifyStaminaRecoverValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaOverflowValueAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaOverflowValueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaOverflowValueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaOverflowValueResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaOverflowValueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaOverflowValueAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaOverflowValueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValueAsync(
	request *VerifyStaminaOverflowValueRequest,
	callback chan<- VerifyStaminaOverflowValueAsyncResult,
) {
	path := "/{namespaceName}/user/me/stamina/{staminaName}/overflow/value/verify/{verifyType}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaOverflowValueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValue(
	request *VerifyStaminaOverflowValueRequest,
) (*VerifyStaminaOverflowValueResult, error) {
	callback := make(chan VerifyStaminaOverflowValueAsyncResult, 1)
	go p.VerifyStaminaOverflowValueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaOverflowValueByUserIdAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaOverflowValueByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaOverflowValueByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaOverflowValueByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaOverflowValueByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaOverflowValueByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaOverflowValueByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValueByUserIdAsync(
	request *VerifyStaminaOverflowValueByUserIdRequest,
	callback chan<- VerifyStaminaOverflowValueByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/stamina/{staminaName}/overflow/value/verify/{verifyType}"
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
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}
	if request.VerifyType != nil && *request.VerifyType != "" {
		path = strings.ReplaceAll(path, "{verifyType}", core.ToString(*request.VerifyType))
	} else {
		path = strings.ReplaceAll(path, "{verifyType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.MultiplyValueSpecifyingQuantity != nil {
		bodies["multiplyValueSpecifyingQuantity"] = *request.MultiplyValueSpecifyingQuantity
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

	go verifyStaminaOverflowValueByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValueByUserId(
	request *VerifyStaminaOverflowValueByUserIdRequest,
) (*VerifyStaminaOverflowValueByUserIdResult, error) {
	callback := make(chan VerifyStaminaOverflowValueByUserIdAsyncResult, 1)
	go p.VerifyStaminaOverflowValueByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func recoverStaminaByStampSheetAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- RecoverStaminaByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RecoverStaminaByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RecoverStaminaByStampSheetResult
	if asyncResult.Err != nil {
		callback <- RecoverStaminaByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RecoverStaminaByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RecoverStaminaByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) RecoverStaminaByStampSheetAsync(
	request *RecoverStaminaByStampSheetRequest,
	callback chan<- RecoverStaminaByStampSheetAsyncResult,
) {
	path := "/stamina/recover"

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

	go recoverStaminaByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) RecoverStaminaByStampSheet(
	request *RecoverStaminaByStampSheetRequest,
) (*RecoverStaminaByStampSheetResult, error) {
	callback := make(chan RecoverStaminaByStampSheetAsyncResult, 1)
	go p.RecoverStaminaByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func raiseMaxValueByStampSheetAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- RaiseMaxValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RaiseMaxValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RaiseMaxValueByStampSheetResult
	if asyncResult.Err != nil {
		callback <- RaiseMaxValueByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RaiseMaxValueByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RaiseMaxValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) RaiseMaxValueByStampSheetAsync(
	request *RaiseMaxValueByStampSheetRequest,
	callback chan<- RaiseMaxValueByStampSheetAsyncResult,
) {
	path := "/stamina/raise"

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

	go raiseMaxValueByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) RaiseMaxValueByStampSheet(
	request *RaiseMaxValueByStampSheetRequest,
) (*RaiseMaxValueByStampSheetResult, error) {
	callback := make(chan RaiseMaxValueByStampSheetAsyncResult, 1)
	go p.RaiseMaxValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaxValueByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaxValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaxValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaxValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaxValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaxValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaxValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) DecreaseMaxValueByStampTaskAsync(
	request *DecreaseMaxValueByStampTaskRequest,
	callback chan<- DecreaseMaxValueByStampTaskAsyncResult,
) {
	path := "/stamina/decrease"

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

	go decreaseMaxValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) DecreaseMaxValueByStampTask(
	request *DecreaseMaxValueByStampTaskRequest,
) (*DecreaseMaxValueByStampTaskResult, error) {
	callback := make(chan DecreaseMaxValueByStampTaskAsyncResult, 1)
	go p.DecreaseMaxValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMaxValueByStampSheetAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetMaxValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaxValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaxValueByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetMaxValueByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaxValueByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMaxValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetMaxValueByStampSheetAsync(
	request *SetMaxValueByStampSheetRequest,
	callback chan<- SetMaxValueByStampSheetAsyncResult,
) {
	path := "/stamina/max/set"

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

	go setMaxValueByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetMaxValueByStampSheet(
	request *SetMaxValueByStampSheetRequest,
) (*SetMaxValueByStampSheetResult, error) {
	callback := make(chan SetMaxValueByStampSheetAsyncResult, 1)
	go p.SetMaxValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverIntervalByStampSheetAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverIntervalByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverIntervalByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverIntervalByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetRecoverIntervalByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverIntervalByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverIntervalByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverIntervalByStampSheetAsync(
	request *SetRecoverIntervalByStampSheetRequest,
	callback chan<- SetRecoverIntervalByStampSheetAsyncResult,
) {
	path := "/stamina/recoverInterval/set"

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

	go setRecoverIntervalByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverIntervalByStampSheet(
	request *SetRecoverIntervalByStampSheetRequest,
) (*SetRecoverIntervalByStampSheetResult, error) {
	callback := make(chan SetRecoverIntervalByStampSheetAsyncResult, 1)
	go p.SetRecoverIntervalByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setRecoverValueByStampSheetAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- SetRecoverValueByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRecoverValueByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRecoverValueByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetRecoverValueByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetRecoverValueByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetRecoverValueByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) SetRecoverValueByStampSheetAsync(
	request *SetRecoverValueByStampSheetRequest,
	callback chan<- SetRecoverValueByStampSheetAsyncResult,
) {
	path := "/stamina/recoverValue/set"

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

	go setRecoverValueByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) SetRecoverValueByStampSheet(
	request *SetRecoverValueByStampSheetRequest,
) (*SetRecoverValueByStampSheetResult, error) {
	callback := make(chan SetRecoverValueByStampSheetAsyncResult, 1)
	go p.SetRecoverValueByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func consumeStaminaByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- ConsumeStaminaByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ConsumeStaminaByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ConsumeStaminaByStampTaskResult
	if asyncResult.Err != nil {
		callback <- ConsumeStaminaByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ConsumeStaminaByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ConsumeStaminaByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) ConsumeStaminaByStampTaskAsync(
	request *ConsumeStaminaByStampTaskRequest,
	callback chan<- ConsumeStaminaByStampTaskAsyncResult,
) {
	path := "/stamina/consume"

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

	go consumeStaminaByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) ConsumeStaminaByStampTask(
	request *ConsumeStaminaByStampTaskRequest,
) (*ConsumeStaminaByStampTaskResult, error) {
	callback := make(chan ConsumeStaminaByStampTaskAsyncResult, 1)
	go p.ConsumeStaminaByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaValueByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaValueByStampTaskAsync(
	request *VerifyStaminaValueByStampTaskRequest,
	callback chan<- VerifyStaminaValueByStampTaskAsyncResult,
) {
	path := "/stamina/value/verify"

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

	go verifyStaminaValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaValueByStampTask(
	request *VerifyStaminaValueByStampTaskRequest,
) (*VerifyStaminaValueByStampTaskResult, error) {
	callback := make(chan VerifyStaminaValueByStampTaskAsyncResult, 1)
	go p.VerifyStaminaValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaMaxValueByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaMaxValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaMaxValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaMaxValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaMaxValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaMaxValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaMaxValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValueByStampTaskAsync(
	request *VerifyStaminaMaxValueByStampTaskRequest,
	callback chan<- VerifyStaminaMaxValueByStampTaskAsyncResult,
) {
	path := "/stamina/max/verify"

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

	go verifyStaminaMaxValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaMaxValueByStampTask(
	request *VerifyStaminaMaxValueByStampTaskRequest,
) (*VerifyStaminaMaxValueByStampTaskResult, error) {
	callback := make(chan VerifyStaminaMaxValueByStampTaskAsyncResult, 1)
	go p.VerifyStaminaMaxValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverIntervalMinutesByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverIntervalMinutesByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutesByStampTaskAsync(
	request *VerifyStaminaRecoverIntervalMinutesByStampTaskRequest,
	callback chan<- VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult,
) {
	path := "/stamina/recover/interval/verify"

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

	go verifyStaminaRecoverIntervalMinutesByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverIntervalMinutesByStampTask(
	request *VerifyStaminaRecoverIntervalMinutesByStampTaskRequest,
) (*VerifyStaminaRecoverIntervalMinutesByStampTaskResult, error) {
	callback := make(chan VerifyStaminaRecoverIntervalMinutesByStampTaskAsyncResult, 1)
	go p.VerifyStaminaRecoverIntervalMinutesByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaRecoverValueByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaRecoverValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaRecoverValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaRecoverValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaRecoverValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaRecoverValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaRecoverValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValueByStampTaskAsync(
	request *VerifyStaminaRecoverValueByStampTaskRequest,
	callback chan<- VerifyStaminaRecoverValueByStampTaskAsyncResult,
) {
	path := "/stamina/recover/value/verify"

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

	go verifyStaminaRecoverValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaRecoverValueByStampTask(
	request *VerifyStaminaRecoverValueByStampTaskRequest,
) (*VerifyStaminaRecoverValueByStampTaskResult, error) {
	callback := make(chan VerifyStaminaRecoverValueByStampTaskAsyncResult, 1)
	go p.VerifyStaminaRecoverValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyStaminaOverflowValueByStampTaskAsyncHandler(
	client Gs2StaminaRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyStaminaOverflowValueByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyStaminaOverflowValueByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyStaminaOverflowValueByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyStaminaOverflowValueByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyStaminaOverflowValueByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyStaminaOverflowValueByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValueByStampTaskAsync(
	request *VerifyStaminaOverflowValueByStampTaskRequest,
	callback chan<- VerifyStaminaOverflowValueByStampTaskAsyncResult,
) {
	path := "/stamina/overflow/value/verify"

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

	go verifyStaminaOverflowValueByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("stamina").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2StaminaRestClient) VerifyStaminaOverflowValueByStampTask(
	request *VerifyStaminaOverflowValueByStampTaskRequest,
) (*VerifyStaminaOverflowValueByStampTaskResult, error) {
	callback := make(chan VerifyStaminaOverflowValueByStampTaskAsyncResult, 1)
	go p.VerifyStaminaOverflowValueByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
