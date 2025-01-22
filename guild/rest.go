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

package guild

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2GuildRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2GuildRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeNamespaces(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) CreateNamespaceAsync(
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
	if request.ChangeNotification != nil {
		bodies["changeNotification"] = request.ChangeNotification.ToDict()
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.ChangeMemberNotification != nil {
		bodies["changeMemberNotification"] = request.ChangeMemberNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.RemoveRequestNotification != nil {
		bodies["removeRequestNotification"] = request.RemoveRequestNotification.ToDict()
	}
	if request.CreateGuildScript != nil {
		bodies["createGuildScript"] = request.CreateGuildScript.ToDict()
	}
	if request.UpdateGuildScript != nil {
		bodies["updateGuildScript"] = request.UpdateGuildScript.ToDict()
	}
	if request.JoinGuildScript != nil {
		bodies["joinGuildScript"] = request.JoinGuildScript.ToDict()
	}
	if request.LeaveGuildScript != nil {
		bodies["leaveGuildScript"] = request.LeaveGuildScript.ToDict()
	}
	if request.ChangeRoleScript != nil {
		bodies["changeRoleScript"] = request.ChangeRoleScript.ToDict()
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CreateNamespace(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetNamespaceStatus(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetNamespace(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) UpdateNamespaceAsync(
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
	if request.ChangeNotification != nil {
		bodies["changeNotification"] = request.ChangeNotification.ToDict()
	}
	if request.JoinNotification != nil {
		bodies["joinNotification"] = request.JoinNotification.ToDict()
	}
	if request.LeaveNotification != nil {
		bodies["leaveNotification"] = request.LeaveNotification.ToDict()
	}
	if request.ChangeMemberNotification != nil {
		bodies["changeMemberNotification"] = request.ChangeMemberNotification.ToDict()
	}
	if request.ReceiveRequestNotification != nil {
		bodies["receiveRequestNotification"] = request.ReceiveRequestNotification.ToDict()
	}
	if request.RemoveRequestNotification != nil {
		bodies["removeRequestNotification"] = request.RemoveRequestNotification.ToDict()
	}
	if request.CreateGuildScript != nil {
		bodies["createGuildScript"] = request.CreateGuildScript.ToDict()
	}
	if request.UpdateGuildScript != nil {
		bodies["updateGuildScript"] = request.UpdateGuildScript.ToDict()
	}
	if request.JoinGuildScript != nil {
		bodies["joinGuildScript"] = request.JoinGuildScript.ToDict()
	}
	if request.LeaveGuildScript != nil {
		bodies["leaveGuildScript"] = request.LeaveGuildScript.ToDict()
	}
	if request.ChangeRoleScript != nil {
		bodies["changeRoleScript"] = request.ChangeRoleScript.ToDict()
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateNamespace(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteNamespace(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DumpUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CheckDumpUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CleanUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CheckCleanUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) PrepareImportUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) ImportUserDataByUserId(
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
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CheckImportUserDataByUserId(
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

func describeGuildModelMastersAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGuildModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGuildModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGuildModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeGuildModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGuildModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGuildModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeGuildModelMastersAsync(
	request *DescribeGuildModelMastersRequest,
	callback chan<- DescribeGuildModelMastersAsyncResult,
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

	go describeGuildModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeGuildModelMasters(
	request *DescribeGuildModelMastersRequest,
) (*DescribeGuildModelMastersResult, error) {
	callback := make(chan DescribeGuildModelMastersAsyncResult, 1)
	go p.DescribeGuildModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGuildModelMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateGuildModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) CreateGuildModelMasterAsync(
	request *CreateGuildModelMasterRequest,
	callback chan<- CreateGuildModelMasterAsyncResult,
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
	if request.DefaultMaximumMemberCount != nil {
		bodies["defaultMaximumMemberCount"] = *request.DefaultMaximumMemberCount
	}
	if request.MaximumMemberCount != nil {
		bodies["maximumMemberCount"] = *request.MaximumMemberCount
	}
	if request.InactivityPeriodDays != nil {
		bodies["inactivityPeriodDays"] = *request.InactivityPeriodDays
	}
	if request.Roles != nil {
		var _roles []interface{}
		for _, item := range request.Roles {
			_roles = append(_roles, item)
		}
		bodies["roles"] = _roles
	}
	if request.GuildMasterRole != nil && *request.GuildMasterRole != "" {
		bodies["guildMasterRole"] = *request.GuildMasterRole
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
	}
	if request.RejoinCoolTimeMinutes != nil {
		bodies["rejoinCoolTimeMinutes"] = *request.RejoinCoolTimeMinutes
	}
	if request.MaxConcurrentJoinGuilds != nil {
		bodies["maxConcurrentJoinGuilds"] = *request.MaxConcurrentJoinGuilds
	}
	if request.MaxConcurrentGuildMasterCount != nil {
		bodies["maxConcurrentGuildMasterCount"] = *request.MaxConcurrentGuildMasterCount
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

	go createGuildModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CreateGuildModelMaster(
	request *CreateGuildModelMasterRequest,
) (*CreateGuildModelMasterResult, error) {
	callback := make(chan CreateGuildModelMasterAsyncResult, 1)
	go p.CreateGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGuildModelMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetGuildModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetGuildModelMasterAsync(
	request *GetGuildModelMasterRequest,
	callback chan<- GetGuildModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go getGuildModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetGuildModelMaster(
	request *GetGuildModelMasterRequest,
) (*GetGuildModelMasterResult, error) {
	callback := make(chan GetGuildModelMasterAsyncResult, 1)
	go p.GetGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGuildModelMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateGuildModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateGuildModelMasterAsync(
	request *UpdateGuildModelMasterRequest,
	callback chan<- UpdateGuildModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.DefaultMaximumMemberCount != nil {
		bodies["defaultMaximumMemberCount"] = *request.DefaultMaximumMemberCount
	}
	if request.MaximumMemberCount != nil {
		bodies["maximumMemberCount"] = *request.MaximumMemberCount
	}
	if request.InactivityPeriodDays != nil {
		bodies["inactivityPeriodDays"] = *request.InactivityPeriodDays
	}
	if request.Roles != nil {
		var _roles []interface{}
		for _, item := range request.Roles {
			_roles = append(_roles, item)
		}
		bodies["roles"] = _roles
	}
	if request.GuildMasterRole != nil && *request.GuildMasterRole != "" {
		bodies["guildMasterRole"] = *request.GuildMasterRole
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
	}
	if request.RejoinCoolTimeMinutes != nil {
		bodies["rejoinCoolTimeMinutes"] = *request.RejoinCoolTimeMinutes
	}
	if request.MaxConcurrentJoinGuilds != nil {
		bodies["maxConcurrentJoinGuilds"] = *request.MaxConcurrentJoinGuilds
	}
	if request.MaxConcurrentGuildMasterCount != nil {
		bodies["maxConcurrentGuildMasterCount"] = *request.MaxConcurrentGuildMasterCount
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

	go updateGuildModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateGuildModelMaster(
	request *UpdateGuildModelMasterRequest,
) (*UpdateGuildModelMasterResult, error) {
	callback := make(chan UpdateGuildModelMasterAsyncResult, 1)
	go p.UpdateGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGuildModelMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGuildModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteGuildModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGuildModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteGuildModelMasterAsync(
	request *DeleteGuildModelMasterRequest,
	callback chan<- DeleteGuildModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go deleteGuildModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteGuildModelMaster(
	request *DeleteGuildModelMasterRequest,
) (*DeleteGuildModelMasterResult, error) {
	callback := make(chan DeleteGuildModelMasterAsyncResult, 1)
	go p.DeleteGuildModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeGuildModelsAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGuildModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGuildModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGuildModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeGuildModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeGuildModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeGuildModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeGuildModelsAsync(
	request *DescribeGuildModelsRequest,
	callback chan<- DescribeGuildModelsAsyncResult,
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

	go describeGuildModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeGuildModels(
	request *DescribeGuildModelsRequest,
) (*DescribeGuildModelsResult, error) {
	callback := make(chan DescribeGuildModelsAsyncResult, 1)
	go p.DescribeGuildModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGuildModelAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetGuildModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildModelResult
	if asyncResult.Err != nil {
		callback <- GetGuildModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGuildModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetGuildModelAsync(
	request *GetGuildModelRequest,
	callback chan<- GetGuildModelAsyncResult,
) {
	path := "/{namespaceName}/model/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go getGuildModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetGuildModel(
	request *GetGuildModelRequest,
) (*GetGuildModelResult, error) {
	callback := make(chan GetGuildModelAsyncResult, 1)
	go p.GetGuildModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func searchGuildsAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SearchGuildsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SearchGuildsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SearchGuildsResult
	if asyncResult.Err != nil {
		callback <- SearchGuildsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SearchGuildsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SearchGuildsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SearchGuildsAsync(
	request *SearchGuildsRequest,
	callback chan<- SearchGuildsAsyncResult,
) {
	path := "/{namespaceName}/user/me/guild/{guildModelName}/search"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attributes1 != nil {
		var _attributes1 []interface{}
		for _, item := range request.Attributes1 {
			_attributes1 = append(_attributes1, item)
		}
		bodies["attributes1"] = _attributes1
	}
	if request.Attributes2 != nil {
		var _attributes2 []interface{}
		for _, item := range request.Attributes2 {
			_attributes2 = append(_attributes2, item)
		}
		bodies["attributes2"] = _attributes2
	}
	if request.Attributes3 != nil {
		var _attributes3 []interface{}
		for _, item := range request.Attributes3 {
			_attributes3 = append(_attributes3, item)
		}
		bodies["attributes3"] = _attributes3
	}
	if request.Attributes4 != nil {
		var _attributes4 []interface{}
		for _, item := range request.Attributes4 {
			_attributes4 = append(_attributes4, item)
		}
		bodies["attributes4"] = _attributes4
	}
	if request.Attributes5 != nil {
		var _attributes5 []interface{}
		for _, item := range request.Attributes5 {
			_attributes5 = append(_attributes5, item)
		}
		bodies["attributes5"] = _attributes5
	}
	if request.JoinPolicies != nil {
		var _joinPolicies []interface{}
		for _, item := range request.JoinPolicies {
			_joinPolicies = append(_joinPolicies, item)
		}
		bodies["joinPolicies"] = _joinPolicies
	}
	if request.IncludeFullMembersGuild != nil {
		bodies["includeFullMembersGuild"] = *request.IncludeFullMembersGuild
	}
	if request.OrderBy != nil && *request.OrderBy != "" {
		bodies["orderBy"] = *request.OrderBy
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
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

	go searchGuildsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SearchGuilds(
	request *SearchGuildsRequest,
) (*SearchGuildsResult, error) {
	callback := make(chan SearchGuildsAsyncResult, 1)
	go p.SearchGuildsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func searchGuildsByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SearchGuildsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SearchGuildsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SearchGuildsByUserIdResult
	if asyncResult.Err != nil {
		callback <- SearchGuildsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SearchGuildsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SearchGuildsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SearchGuildsByUserIdAsync(
	request *SearchGuildsByUserIdRequest,
	callback chan<- SearchGuildsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/guild/{guildModelName}/search"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attributes1 != nil {
		var _attributes1 []interface{}
		for _, item := range request.Attributes1 {
			_attributes1 = append(_attributes1, item)
		}
		bodies["attributes1"] = _attributes1
	}
	if request.Attributes2 != nil {
		var _attributes2 []interface{}
		for _, item := range request.Attributes2 {
			_attributes2 = append(_attributes2, item)
		}
		bodies["attributes2"] = _attributes2
	}
	if request.Attributes3 != nil {
		var _attributes3 []interface{}
		for _, item := range request.Attributes3 {
			_attributes3 = append(_attributes3, item)
		}
		bodies["attributes3"] = _attributes3
	}
	if request.Attributes4 != nil {
		var _attributes4 []interface{}
		for _, item := range request.Attributes4 {
			_attributes4 = append(_attributes4, item)
		}
		bodies["attributes4"] = _attributes4
	}
	if request.Attributes5 != nil {
		var _attributes5 []interface{}
		for _, item := range request.Attributes5 {
			_attributes5 = append(_attributes5, item)
		}
		bodies["attributes5"] = _attributes5
	}
	if request.JoinPolicies != nil {
		var _joinPolicies []interface{}
		for _, item := range request.JoinPolicies {
			_joinPolicies = append(_joinPolicies, item)
		}
		bodies["joinPolicies"] = _joinPolicies
	}
	if request.IncludeFullMembersGuild != nil {
		bodies["includeFullMembersGuild"] = *request.IncludeFullMembersGuild
	}
	if request.OrderBy != nil && *request.OrderBy != "" {
		bodies["orderBy"] = *request.OrderBy
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
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

	go searchGuildsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SearchGuildsByUserId(
	request *SearchGuildsByUserIdRequest,
) (*SearchGuildsByUserIdResult, error) {
	callback := make(chan SearchGuildsByUserIdAsyncResult, 1)
	go p.SearchGuildsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGuildAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildResult
	if asyncResult.Err != nil {
		callback <- CreateGuildAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) CreateGuildAsync(
	request *CreateGuildRequest,
	callback chan<- CreateGuildAsyncResult,
) {
	path := "/{namespaceName}/user/me/guild/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MemberMetadata != nil && *request.MemberMetadata != "" {
		bodies["memberMetadata"] = *request.MemberMetadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go createGuildAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CreateGuild(
	request *CreateGuildRequest,
) (*CreateGuildResult, error) {
	callback := make(chan CreateGuildAsyncResult, 1)
	go p.CreateGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGuildByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGuildByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateGuildByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) CreateGuildByUserIdAsync(
	request *CreateGuildByUserIdRequest,
	callback chan<- CreateGuildByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/guild/{guildModelName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.MemberMetadata != nil && *request.MemberMetadata != "" {
		bodies["memberMetadata"] = *request.MemberMetadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go createGuildByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) CreateGuildByUserId(
	request *CreateGuildByUserIdRequest,
) (*CreateGuildByUserIdResult, error) {
	callback := make(chan CreateGuildByUserIdAsyncResult, 1)
	go p.CreateGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGuildAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildResult
	if asyncResult.Err != nil {
		callback <- GetGuildAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetGuildAsync(
	request *GetGuildRequest,
	callback chan<- GetGuildAsyncResult,
) {
	path := "/{namespaceName}/user/me/guild/{guildModelName}/{guildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getGuildAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetGuild(
	request *GetGuildRequest,
) (*GetGuildResult, error) {
	callback := make(chan GetGuildAsyncResult, 1)
	go p.GetGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGuildByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGuildByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetGuildByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetGuildByUserIdAsync(
	request *GetGuildByUserIdRequest,
	callback chan<- GetGuildByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/guild/{guildModelName}/{guildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getGuildByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetGuildByUserId(
	request *GetGuildByUserIdRequest,
) (*GetGuildByUserIdResult, error) {
	callback := make(chan GetGuildByUserIdAsyncResult, 1)
	go p.GetGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGuildAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildResult
	if asyncResult.Err != nil {
		callback <- UpdateGuildAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateGuildAsync(
	request *UpdateGuildRequest,
	callback chan<- UpdateGuildAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
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

	go updateGuildAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateGuild(
	request *UpdateGuildRequest,
) (*UpdateGuildResult, error) {
	callback := make(chan UpdateGuildAsyncResult, 1)
	go p.UpdateGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGuildByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGuildByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGuildByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGuildByGuildNameResult
	if asyncResult.Err != nil {
		callback <- UpdateGuildByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateGuildByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateGuildByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateGuildByGuildNameAsync(
	request *UpdateGuildByGuildNameRequest,
	callback chan<- UpdateGuildByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DisplayName != nil && *request.DisplayName != "" {
		bodies["displayName"] = *request.DisplayName
	}
	if request.Attribute1 != nil {
		bodies["attribute1"] = *request.Attribute1
	}
	if request.Attribute2 != nil {
		bodies["attribute2"] = *request.Attribute2
	}
	if request.Attribute3 != nil {
		bodies["attribute3"] = *request.Attribute3
	}
	if request.Attribute4 != nil {
		bodies["attribute4"] = *request.Attribute4
	}
	if request.Attribute5 != nil {
		bodies["attribute5"] = *request.Attribute5
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.JoinPolicy != nil && *request.JoinPolicy != "" {
		bodies["joinPolicy"] = *request.JoinPolicy
	}
	if request.CustomRoles != nil {
		var _customRoles []interface{}
		for _, item := range request.CustomRoles {
			_customRoles = append(_customRoles, item)
		}
		bodies["customRoles"] = _customRoles
	}
	if request.GuildMemberDefaultRole != nil && *request.GuildMemberDefaultRole != "" {
		bodies["guildMemberDefaultRole"] = *request.GuildMemberDefaultRole
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go updateGuildByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateGuildByGuildName(
	request *UpdateGuildByGuildNameRequest,
) (*UpdateGuildByGuildNameResult, error) {
	callback := make(chan UpdateGuildByGuildNameAsyncResult, 1)
	go p.UpdateGuildByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMemberAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMemberResult
	if asyncResult.Err != nil {
		callback <- DeleteMemberAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteMemberAsync(
	request *DeleteMemberRequest,
	callback chan<- DeleteMemberAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/member/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go deleteMemberAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteMember(
	request *DeleteMemberRequest,
) (*DeleteMemberResult, error) {
	callback := make(chan DeleteMemberAsyncResult, 1)
	go p.DeleteMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteMemberByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteMemberByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteMemberByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteMemberByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DeleteMemberByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteMemberByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteMemberByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteMemberByGuildNameAsync(
	request *DeleteMemberByGuildNameRequest,
	callback chan<- DeleteMemberByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/member/{targetUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go deleteMemberByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteMemberByGuildName(
	request *DeleteMemberByGuildNameRequest,
) (*DeleteMemberByGuildNameResult, error) {
	callback := make(chan DeleteMemberByGuildNameAsyncResult, 1)
	go p.DeleteMemberByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMemberRoleAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMemberRoleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberRoleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberRoleResult
	if asyncResult.Err != nil {
		callback <- UpdateMemberRoleAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberRoleAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMemberRoleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateMemberRoleAsync(
	request *UpdateMemberRoleRequest,
	callback chan<- UpdateMemberRoleAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/member/{targetUserId}/role"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RoleName != nil && *request.RoleName != "" {
		bodies["roleName"] = *request.RoleName
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

	go updateMemberRoleAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateMemberRole(
	request *UpdateMemberRoleRequest,
) (*UpdateMemberRoleResult, error) {
	callback := make(chan UpdateMemberRoleAsyncResult, 1)
	go p.UpdateMemberRoleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMemberRoleByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMemberRoleByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberRoleByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberRoleByGuildNameResult
	if asyncResult.Err != nil {
		callback <- UpdateMemberRoleByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberRoleByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMemberRoleByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateMemberRoleByGuildNameAsync(
	request *UpdateMemberRoleByGuildNameRequest,
	callback chan<- UpdateMemberRoleByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/member/{targetUserId}/role"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.TargetUserId != nil && *request.TargetUserId != "" {
		path = strings.ReplaceAll(path, "{targetUserId}", core.ToString(*request.TargetUserId))
	} else {
		path = strings.ReplaceAll(path, "{targetUserId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.RoleName != nil && *request.RoleName != "" {
		bodies["roleName"] = *request.RoleName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go updateMemberRoleByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateMemberRoleByGuildName(
	request *UpdateMemberRoleByGuildNameRequest,
) (*UpdateMemberRoleByGuildNameResult, error) {
	callback := make(chan UpdateMemberRoleByGuildNameAsyncResult, 1)
	go p.UpdateMemberRoleByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchUpdateMemberRoleAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- BatchUpdateMemberRoleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchUpdateMemberRoleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchUpdateMemberRoleResult
	if asyncResult.Err != nil {
		callback <- BatchUpdateMemberRoleAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchUpdateMemberRoleAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchUpdateMemberRoleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) BatchUpdateMemberRoleAsync(
	request *BatchUpdateMemberRoleRequest,
	callback chan<- BatchUpdateMemberRoleAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/batch/member/role"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Members != nil {
		var _members []interface{}
		for _, item := range request.Members {
			_members = append(_members, item)
		}
		bodies["members"] = _members
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

	go batchUpdateMemberRoleAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) BatchUpdateMemberRole(
	request *BatchUpdateMemberRoleRequest,
) (*BatchUpdateMemberRoleResult, error) {
	callback := make(chan BatchUpdateMemberRoleAsyncResult, 1)
	go p.BatchUpdateMemberRoleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func batchUpdateMemberRoleByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- BatchUpdateMemberRoleByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result BatchUpdateMemberRoleByGuildNameResult
	if asyncResult.Err != nil {
		callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- BatchUpdateMemberRoleByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) BatchUpdateMemberRoleByGuildNameAsync(
	request *BatchUpdateMemberRoleByGuildNameRequest,
	callback chan<- BatchUpdateMemberRoleByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/batch/member/role"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Members != nil {
		var _members []interface{}
		for _, item := range request.Members {
			_members = append(_members, item)
		}
		bodies["members"] = _members
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go batchUpdateMemberRoleByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) BatchUpdateMemberRoleByGuildName(
	request *BatchUpdateMemberRoleByGuildNameRequest,
) (*BatchUpdateMemberRoleByGuildNameResult, error) {
	callback := make(chan BatchUpdateMemberRoleByGuildNameAsyncResult, 1)
	go p.BatchUpdateMemberRoleByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMemberMetadataAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMemberMetadataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberMetadataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberMetadataResult
	if asyncResult.Err != nil {
		callback <- UpdateMemberMetadataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberMetadataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMemberMetadataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateMemberMetadataAsync(
	request *UpdateMemberMetadataRequest,
	callback chan<- UpdateMemberMetadataAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/guild/{guildName}/member/me/metadata"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

	go updateMemberMetadataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateMemberMetadata(
	request *UpdateMemberMetadataRequest,
) (*UpdateMemberMetadataResult, error) {
	callback := make(chan UpdateMemberMetadataAsyncResult, 1)
	go p.UpdateMemberMetadataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateMemberMetadataByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateMemberMetadataByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateMemberMetadataByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateMemberMetadataByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateMemberMetadataByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateMemberMetadataByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateMemberMetadataByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateMemberMetadataByUserIdAsync(
	request *UpdateMemberMetadataByUserIdRequest,
	callback chan<- UpdateMemberMetadataByUserIdAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/guild/{guildName}/member/{userId}/metadata"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go updateMemberMetadataByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateMemberMetadataByUserId(
	request *UpdateMemberMetadataByUserIdRequest,
) (*UpdateMemberMetadataByUserIdResult, error) {
	callback := make(chan UpdateMemberMetadataByUserIdAsyncResult, 1)
	go p.UpdateMemberMetadataByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGuildAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildResult
	if asyncResult.Err != nil {
		callback <- DeleteGuildAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteGuildAsync(
	request *DeleteGuildRequest,
	callback chan<- DeleteGuildAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go deleteGuildAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteGuild(
	request *DeleteGuildRequest,
) (*DeleteGuildResult, error) {
	callback := make(chan DeleteGuildAsyncResult, 1)
	go p.DeleteGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGuildByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGuildByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGuildByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGuildByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DeleteGuildByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteGuildByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteGuildByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteGuildByGuildNameAsync(
	request *DeleteGuildByGuildNameRequest,
	callback chan<- DeleteGuildByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go deleteGuildByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteGuildByGuildName(
	request *DeleteGuildByGuildNameRequest,
) (*DeleteGuildByGuildNameResult, error) {
	callback := make(chan DeleteGuildByGuildNameAsyncResult, 1)
	go p.DeleteGuildByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func increaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/currentMaximumMemberCount/increase"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go increaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) IncreaseMaximumCurrentMaximumMemberCountByGuildName(
	request *IncreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*IncreaseMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.IncreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaximumCurrentMaximumMemberCountAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCountAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/currentMaximumMemberCount/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
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

	go decreaseMaximumCurrentMaximumMemberCountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCount(
	request *DecreaseMaximumCurrentMaximumMemberCountRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/currentMaximumMemberCount/decrease"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go decreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCountByGuildName(
	request *DecreaseMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCurrentMaximumMemberCountAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountResult
	if asyncResult.Err != nil {
		callback <- VerifyCurrentMaximumMemberCountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCurrentMaximumMemberCountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCountAsync(
	request *VerifyCurrentMaximumMemberCountRequest,
	callback chan<- VerifyCurrentMaximumMemberCountAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/currentMaximumMemberCount/verify"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
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

	go verifyCurrentMaximumMemberCountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCount(
	request *VerifyCurrentMaximumMemberCountRequest,
) (*VerifyCurrentMaximumMemberCountResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCurrentMaximumMemberCountByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Err != nil {
		callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCountByGuildNameAsync(
	request *VerifyCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- VerifyCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/currentMaximumMemberCount/verify"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
	}
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go verifyCurrentMaximumMemberCountByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCountByGuildName(
	request *VerifyCurrentMaximumMemberCountByGuildNameRequest,
) (*VerifyCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyIncludeMemberAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyIncludeMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberResult
	if asyncResult.Err != nil {
		callback <- VerifyIncludeMemberAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyIncludeMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyIncludeMemberAsync(
	request *VerifyIncludeMemberRequest,
	callback chan<- VerifyIncludeMemberAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/member/me/verify"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
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

	go verifyIncludeMemberAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyIncludeMember(
	request *VerifyIncludeMemberRequest,
) (*VerifyIncludeMemberResult, error) {
	callback := make(chan VerifyIncludeMemberAsyncResult, 1)
	go p.VerifyIncludeMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyIncludeMemberByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyIncludeMemberByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberByUserIdResult
	if asyncResult.Err != nil {
		callback <- VerifyIncludeMemberByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyIncludeMemberByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyIncludeMemberByUserIdAsync(
	request *VerifyIncludeMemberByUserIdRequest,
	callback chan<- VerifyIncludeMemberByUserIdAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/member/{userId}/verify"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.UserId != nil && *request.UserId != "" {
		path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
	} else {
		path = strings.ReplaceAll(path, "{userId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyType != nil && *request.VerifyType != "" {
		bodies["verifyType"] = *request.VerifyType
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

	go verifyIncludeMemberByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyIncludeMemberByUserId(
	request *VerifyIncludeMemberByUserIdRequest,
) (*VerifyIncludeMemberByUserIdResult, error) {
	callback := make(chan VerifyIncludeMemberByUserIdAsyncResult, 1)
	go p.VerifyIncludeMemberByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaximumCurrentMaximumMemberCountByGuildNameResult
	if asyncResult.Err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SetMaximumCurrentMaximumMemberCountByGuildNameAsync(
	request *SetMaximumCurrentMaximumMemberCountByGuildNameRequest,
	callback chan<- SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/currentMaximumMemberCount"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Value != nil {
		bodies["value"] = *request.Value
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
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

	go setMaximumCurrentMaximumMemberCountByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SetMaximumCurrentMaximumMemberCountByGuildName(
	request *SetMaximumCurrentMaximumMemberCountByGuildNameRequest,
) (*SetMaximumCurrentMaximumMemberCountByGuildNameResult, error) {
	callback := make(chan SetMaximumCurrentMaximumMemberCountByGuildNameAsyncResult, 1)
	go p.SetMaximumCurrentMaximumMemberCountByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func assumeAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AssumeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AssumeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AssumeResult
	if asyncResult.Err != nil {
		callback <- AssumeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AssumeAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AssumeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AssumeAsync(
	request *AssumeRequest,
	callback chan<- AssumeAsyncResult,
) {
	path := "/{namespaceName}/user/me/guild/{guildModelName}/{guildName}/assume"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go assumeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) Assume(
	request *AssumeRequest,
) (*AssumeResult, error) {
	callback := make(chan AssumeAsyncResult, 1)
	go p.AssumeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func assumeByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AssumeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AssumeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AssumeByUserIdResult
	if asyncResult.Err != nil {
		callback <- AssumeByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AssumeByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AssumeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AssumeByUserIdAsync(
	request *AssumeByUserIdRequest,
	callback chan<- AssumeByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/guild/{guildModelName}/{guildName}/assume"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go assumeByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) AssumeByUserId(
	request *AssumeByUserIdRequest,
) (*AssumeByUserIdResult, error) {
	callback := make(chan AssumeByUserIdAsyncResult, 1)
	go p.AssumeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func increaseMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult
	if asyncResult.Err != nil {
		callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsync(
	request *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest,
	callback chan<- IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	path := "/stamp/guild/currentMaximumMemberCount/add"

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

	go increaseMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) IncreaseMaximumCurrentMaximumMemberCountByStampSheet(
	request *IncreaseMaximumCurrentMaximumMemberCountByStampSheetRequest,
) (*IncreaseMaximumCurrentMaximumMemberCountByStampSheetResult, error) {
	callback := make(chan IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsyncResult, 1)
	go p.IncreaseMaximumCurrentMaximumMemberCountByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func decreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult
	if asyncResult.Err != nil {
		callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsync(
	request *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest,
	callback chan<- DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	path := "/stamp/guild/currentMaximumMemberCount/sub"

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

	go decreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DecreaseMaximumCurrentMaximumMemberCountByStampTask(
	request *DecreaseMaximumCurrentMaximumMemberCountByStampTaskRequest,
) (*DecreaseMaximumCurrentMaximumMemberCountByStampTaskResult, error) {
	callback := make(chan DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsyncResult, 1)
	go p.DecreaseMaximumCurrentMaximumMemberCountByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetMaximumCurrentMaximumMemberCountByStampSheetResult
	if asyncResult.Err != nil {
		callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SetMaximumCurrentMaximumMemberCountByStampSheetAsync(
	request *SetMaximumCurrentMaximumMemberCountByStampSheetRequest,
	callback chan<- SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult,
) {
	path := "/stamp/guild/currentMaximumMemberCount/set"

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

	go setMaximumCurrentMaximumMemberCountByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SetMaximumCurrentMaximumMemberCountByStampSheet(
	request *SetMaximumCurrentMaximumMemberCountByStampSheetRequest,
) (*SetMaximumCurrentMaximumMemberCountByStampSheetResult, error) {
	callback := make(chan SetMaximumCurrentMaximumMemberCountByStampSheetAsyncResult, 1)
	go p.SetMaximumCurrentMaximumMemberCountByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyCurrentMaximumMemberCountByStampTaskAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyCurrentMaximumMemberCountByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCountByStampTaskAsync(
	request *VerifyCurrentMaximumMemberCountByStampTaskRequest,
	callback chan<- VerifyCurrentMaximumMemberCountByStampTaskAsyncResult,
) {
	path := "/stamp/guild/currentMaximumMemberCount/verify"

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

	go verifyCurrentMaximumMemberCountByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyCurrentMaximumMemberCountByStampTask(
	request *VerifyCurrentMaximumMemberCountByStampTaskRequest,
) (*VerifyCurrentMaximumMemberCountByStampTaskResult, error) {
	callback := make(chan VerifyCurrentMaximumMemberCountByStampTaskAsyncResult, 1)
	go p.VerifyCurrentMaximumMemberCountByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyIncludeMemberByStampTaskAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyIncludeMemberByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyIncludeMemberByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyIncludeMemberByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyIncludeMemberByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyIncludeMemberByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyIncludeMemberByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) VerifyIncludeMemberByStampTaskAsync(
	request *VerifyIncludeMemberByStampTaskRequest,
	callback chan<- VerifyIncludeMemberByStampTaskAsyncResult,
) {
	path := "/stamp/guild/member/verify"

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

	go verifyIncludeMemberByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) VerifyIncludeMemberByStampTask(
	request *VerifyIncludeMemberByStampTaskRequest,
) (*VerifyIncludeMemberByStampTaskResult, error) {
	callback := make(chan VerifyIncludeMemberByStampTaskAsyncResult, 1)
	go p.VerifyIncludeMemberByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeJoinedGuildsAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeJoinedGuildsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeJoinedGuildsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeJoinedGuildsResult
	if asyncResult.Err != nil {
		callback <- DescribeJoinedGuildsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeJoinedGuildsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeJoinedGuildsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeJoinedGuildsAsync(
	request *DescribeJoinedGuildsRequest,
	callback chan<- DescribeJoinedGuildsAsyncResult,
) {
	path := "/{namespaceName}/user/me/joined"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.GuildModelName != nil {
		queryStrings["guildModelName"] = core.ToString(*request.GuildModelName)
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

	go describeJoinedGuildsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeJoinedGuilds(
	request *DescribeJoinedGuildsRequest,
) (*DescribeJoinedGuildsResult, error) {
	callback := make(chan DescribeJoinedGuildsAsyncResult, 1)
	go p.DescribeJoinedGuildsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeJoinedGuildsByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeJoinedGuildsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeJoinedGuildsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeJoinedGuildsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeJoinedGuildsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeJoinedGuildsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeJoinedGuildsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeJoinedGuildsByUserIdAsync(
	request *DescribeJoinedGuildsByUserIdRequest,
	callback chan<- DescribeJoinedGuildsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/joined"
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
	if request.GuildModelName != nil {
		queryStrings["guildModelName"] = core.ToString(*request.GuildModelName)
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

	go describeJoinedGuildsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeJoinedGuildsByUserId(
	request *DescribeJoinedGuildsByUserIdRequest,
) (*DescribeJoinedGuildsByUserIdResult, error) {
	callback := make(chan DescribeJoinedGuildsByUserIdAsyncResult, 1)
	go p.DescribeJoinedGuildsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getJoinedGuildAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetJoinedGuildAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetJoinedGuildAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetJoinedGuildResult
	if asyncResult.Err != nil {
		callback <- GetJoinedGuildAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetJoinedGuildAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetJoinedGuildAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetJoinedGuildAsync(
	request *GetJoinedGuildRequest,
	callback chan<- GetJoinedGuildAsyncResult,
) {
	path := "/{namespaceName}/user/me/joined/{guildModelName}/{guildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getJoinedGuildAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetJoinedGuild(
	request *GetJoinedGuildRequest,
) (*GetJoinedGuildResult, error) {
	callback := make(chan GetJoinedGuildAsyncResult, 1)
	go p.GetJoinedGuildAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getJoinedGuildByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetJoinedGuildByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetJoinedGuildByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetJoinedGuildByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetJoinedGuildByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetJoinedGuildByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetJoinedGuildByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetJoinedGuildByUserIdAsync(
	request *GetJoinedGuildByUserIdRequest,
	callback chan<- GetJoinedGuildByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/joined/{guildModelName}/{guildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getJoinedGuildByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetJoinedGuildByUserId(
	request *GetJoinedGuildByUserIdRequest,
) (*GetJoinedGuildByUserIdResult, error) {
	callback := make(chan GetJoinedGuildByUserIdAsyncResult, 1)
	go p.GetJoinedGuildByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func withdrawalAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- WithdrawalAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawalAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawalResult
	if asyncResult.Err != nil {
		callback <- WithdrawalAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawalAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WithdrawalAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) WithdrawalAsync(
	request *WithdrawalRequest,
	callback chan<- WithdrawalAsyncResult,
) {
	path := "/{namespaceName}/user/me/joined/{guildModelName}/{guildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go withdrawalAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) Withdrawal(
	request *WithdrawalRequest,
) (*WithdrawalResult, error) {
	callback := make(chan WithdrawalAsyncResult, 1)
	go p.WithdrawalAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func withdrawalByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- WithdrawalByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawalByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawalByUserIdResult
	if asyncResult.Err != nil {
		callback <- WithdrawalByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawalByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WithdrawalByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) WithdrawalByUserIdAsync(
	request *WithdrawalByUserIdRequest,
	callback chan<- WithdrawalByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/joined/{guildModelName}/{guildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go withdrawalByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) WithdrawalByUserId(
	request *WithdrawalByUserIdRequest,
) (*WithdrawalByUserIdResult, error) {
	callback := make(chan WithdrawalByUserIdAsyncResult, 1)
	go p.WithdrawalByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLastGuildMasterActivityAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetLastGuildMasterActivityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLastGuildMasterActivityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLastGuildMasterActivityResult
	if asyncResult.Err != nil {
		callback <- GetLastGuildMasterActivityAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLastGuildMasterActivityAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLastGuildMasterActivityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetLastGuildMasterActivityAsync(
	request *GetLastGuildMasterActivityRequest,
	callback chan<- GetLastGuildMasterActivityAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/activity/guildMaster/last"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go getLastGuildMasterActivityAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetLastGuildMasterActivity(
	request *GetLastGuildMasterActivityRequest,
) (*GetLastGuildMasterActivityResult, error) {
	callback := make(chan GetLastGuildMasterActivityAsyncResult, 1)
	go p.GetLastGuildMasterActivityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getLastGuildMasterActivityByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetLastGuildMasterActivityByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLastGuildMasterActivityByGuildNameResult
	if asyncResult.Err != nil {
		callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLastGuildMasterActivityByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetLastGuildMasterActivityByGuildNameAsync(
	request *GetLastGuildMasterActivityByGuildNameRequest,
	callback chan<- GetLastGuildMasterActivityByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/activity/guildMaster/last"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getLastGuildMasterActivityByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetLastGuildMasterActivityByGuildName(
	request *GetLastGuildMasterActivityByGuildNameRequest,
) (*GetLastGuildMasterActivityByGuildNameResult, error) {
	callback := make(chan GetLastGuildMasterActivityByGuildNameAsyncResult, 1)
	go p.GetLastGuildMasterActivityByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func promoteSeniorMemberAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- PromoteSeniorMemberAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteSeniorMemberAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteSeniorMemberResult
	if asyncResult.Err != nil {
		callback <- PromoteSeniorMemberAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteSeniorMemberAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PromoteSeniorMemberAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) PromoteSeniorMemberAsync(
	request *PromoteSeniorMemberRequest,
	callback chan<- PromoteSeniorMemberAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/promote"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go promoteSeniorMemberAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) PromoteSeniorMember(
	request *PromoteSeniorMemberRequest,
) (*PromoteSeniorMemberResult, error) {
	callback := make(chan PromoteSeniorMemberAsyncResult, 1)
	go p.PromoteSeniorMemberAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func promoteSeniorMemberByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- PromoteSeniorMemberByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PromoteSeniorMemberByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PromoteSeniorMemberByGuildNameResult
	if asyncResult.Err != nil {
		callback <- PromoteSeniorMemberByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PromoteSeniorMemberByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PromoteSeniorMemberByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) PromoteSeniorMemberByGuildNameAsync(
	request *PromoteSeniorMemberByGuildNameRequest,
	callback chan<- PromoteSeniorMemberByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/promote"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go promoteSeniorMemberByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) PromoteSeniorMemberByGuildName(
	request *PromoteSeniorMemberByGuildNameRequest,
) (*PromoteSeniorMemberByGuildNameResult, error) {
	callback := make(chan PromoteSeniorMemberByGuildNameAsyncResult, 1)
	go p.PromoteSeniorMemberByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2GuildRestClient,
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

func (p Gs2GuildRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) ExportMaster(
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

func getCurrentGuildMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentGuildMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentGuildMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentGuildMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentGuildMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentGuildMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentGuildMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetCurrentGuildMasterAsync(
	request *GetCurrentGuildMasterRequest,
	callback chan<- GetCurrentGuildMasterAsyncResult,
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

	go getCurrentGuildMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetCurrentGuildMaster(
	request *GetCurrentGuildMasterRequest,
) (*GetCurrentGuildMasterResult, error) {
	callback := make(chan GetCurrentGuildMasterAsyncResult, 1)
	go p.GetCurrentGuildMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentGuildMasterAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentGuildMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGuildMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGuildMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentGuildMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGuildMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentGuildMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateCurrentGuildMasterAsync(
	request *UpdateCurrentGuildMasterRequest,
	callback chan<- UpdateCurrentGuildMasterAsyncResult,
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go updateCurrentGuildMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateCurrentGuildMaster(
	request *UpdateCurrentGuildMasterRequest,
) (*UpdateCurrentGuildMasterResult, error) {
	callback := make(chan UpdateCurrentGuildMasterAsyncResult, 1)
	go p.UpdateCurrentGuildMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentGuildMasterFromGitHubAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentGuildMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentGuildMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentGuildMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) UpdateCurrentGuildMasterFromGitHubAsync(
	request *UpdateCurrentGuildMasterFromGitHubRequest,
	callback chan<- UpdateCurrentGuildMasterFromGitHubAsyncResult,
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

	go updateCurrentGuildMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) UpdateCurrentGuildMasterFromGitHub(
	request *UpdateCurrentGuildMasterFromGitHubRequest,
) (*UpdateCurrentGuildMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentGuildMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentGuildMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReceiveRequestsAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReceiveRequestsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveRequestsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveRequestsResult
	if asyncResult.Err != nil {
		callback <- DescribeReceiveRequestsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveRequestsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReceiveRequestsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeReceiveRequestsAsync(
	request *DescribeReceiveRequestsRequest,
	callback chan<- DescribeReceiveRequestsAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/inbox"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go describeReceiveRequestsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeReceiveRequests(
	request *DescribeReceiveRequestsRequest,
) (*DescribeReceiveRequestsResult, error) {
	callback := make(chan DescribeReceiveRequestsAsyncResult, 1)
	go p.DescribeReceiveRequestsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReceiveRequestsByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReceiveRequestsByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiveRequestsByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReceiveRequestsByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeReceiveRequestsByGuildNameAsync(
	request *DescribeReceiveRequestsByGuildNameRequest,
	callback chan<- DescribeReceiveRequestsByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/inbox"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go describeReceiveRequestsByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeReceiveRequestsByGuildName(
	request *DescribeReceiveRequestsByGuildNameRequest,
) (*DescribeReceiveRequestsByGuildNameResult, error) {
	callback := make(chan DescribeReceiveRequestsByGuildNameAsyncResult, 1)
	go p.DescribeReceiveRequestsByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReceiveRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetReceiveRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveRequestResult
	if asyncResult.Err != nil {
		callback <- GetReceiveRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReceiveRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetReceiveRequestAsync(
	request *GetReceiveRequestRequest,
	callback chan<- GetReceiveRequestAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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

	go getReceiveRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetReceiveRequest(
	request *GetReceiveRequestRequest,
) (*GetReceiveRequestResult, error) {
	callback := make(chan GetReceiveRequestAsyncResult, 1)
	go p.GetReceiveRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getReceiveRequestByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetReceiveRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetReceiveRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetReceiveRequestByGuildNameResult
	if asyncResult.Err != nil {
		callback <- GetReceiveRequestByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetReceiveRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetReceiveRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetReceiveRequestByGuildNameAsync(
	request *GetReceiveRequestByGuildNameRequest,
	callback chan<- GetReceiveRequestByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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

	go getReceiveRequestByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetReceiveRequestByGuildName(
	request *GetReceiveRequestByGuildNameRequest,
) (*GetReceiveRequestByGuildNameResult, error) {
	callback := make(chan GetReceiveRequestByGuildNameAsyncResult, 1)
	go p.GetReceiveRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acceptRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AcceptRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptRequestResult
	if asyncResult.Err != nil {
		callback <- AcceptRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcceptRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AcceptRequestAsync(
	request *AcceptRequestRequest,
	callback chan<- AcceptRequestAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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

	go acceptRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) AcceptRequest(
	request *AcceptRequestRequest,
) (*AcceptRequestResult, error) {
	callback := make(chan AcceptRequestAsyncResult, 1)
	go p.AcceptRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func acceptRequestByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AcceptRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AcceptRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AcceptRequestByGuildNameResult
	if asyncResult.Err != nil {
		callback <- AcceptRequestByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AcceptRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AcceptRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AcceptRequestByGuildNameAsync(
	request *AcceptRequestByGuildNameRequest,
	callback chan<- AcceptRequestByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go acceptRequestByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) AcceptRequestByGuildName(
	request *AcceptRequestByGuildNameRequest,
) (*AcceptRequestByGuildNameResult, error) {
	callback := make(chan AcceptRequestByGuildNameAsyncResult, 1)
	go p.AcceptRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func rejectRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- RejectRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectRequestResult
	if asyncResult.Err != nil {
		callback <- RejectRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RejectRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) RejectRequestAsync(
	request *RejectRequestRequest,
	callback chan<- RejectRequestAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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

	go rejectRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) RejectRequest(
	request *RejectRequestRequest,
) (*RejectRequestResult, error) {
	callback := make(chan RejectRequestAsyncResult, 1)
	go p.RejectRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func rejectRequestByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- RejectRequestByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RejectRequestByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RejectRequestByGuildNameResult
	if asyncResult.Err != nil {
		callback <- RejectRequestByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RejectRequestByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RejectRequestByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) RejectRequestByGuildNameAsync(
	request *RejectRequestByGuildNameRequest,
	callback chan<- RejectRequestByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/inbox/{fromUserId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
	}
	if request.FromUserId != nil && *request.FromUserId != "" {
		path = strings.ReplaceAll(path, "{fromUserId}", core.ToString(*request.FromUserId))
	} else {
		path = strings.ReplaceAll(path, "{fromUserId}", "null")
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go rejectRequestByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) RejectRequestByGuildName(
	request *RejectRequestByGuildNameRequest,
) (*RejectRequestByGuildNameResult, error) {
	callback := make(chan RejectRequestByGuildNameAsyncResult, 1)
	go p.RejectRequestByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSendRequestsAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSendRequestsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSendRequestsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSendRequestsResult
	if asyncResult.Err != nil {
		callback <- DescribeSendRequestsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSendRequestsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSendRequestsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeSendRequestsAsync(
	request *DescribeSendRequestsRequest,
	callback chan<- DescribeSendRequestsAsyncResult,
) {
	path := "/{namespaceName}/user/me/sendBox/guild/{guildModelName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go describeSendRequestsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeSendRequests(
	request *DescribeSendRequestsRequest,
) (*DescribeSendRequestsResult, error) {
	callback := make(chan DescribeSendRequestsAsyncResult, 1)
	go p.DescribeSendRequestsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSendRequestsByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSendRequestsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSendRequestsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSendRequestsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSendRequestsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSendRequestsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSendRequestsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeSendRequestsByUserIdAsync(
	request *DescribeSendRequestsByUserIdRequest,
	callback chan<- DescribeSendRequestsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/sendBox/guild/{guildModelName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go describeSendRequestsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeSendRequestsByUserId(
	request *DescribeSendRequestsByUserIdRequest,
) (*DescribeSendRequestsByUserIdResult, error) {
	callback := make(chan DescribeSendRequestsByUserIdAsyncResult, 1)
	go p.DescribeSendRequestsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSendRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetSendRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSendRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSendRequestResult
	if asyncResult.Err != nil {
		callback <- GetSendRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSendRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSendRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetSendRequestAsync(
	request *GetSendRequestRequest,
	callback chan<- GetSendRequestAsyncResult,
) {
	path := "/{namespaceName}/user/me/sendBox/guild/{guildModelName}/{targetGuildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
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

	go getSendRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetSendRequest(
	request *GetSendRequestRequest,
) (*GetSendRequestResult, error) {
	callback := make(chan GetSendRequestAsyncResult, 1)
	go p.GetSendRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSendRequestByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetSendRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSendRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSendRequestByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSendRequestByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSendRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSendRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetSendRequestByUserIdAsync(
	request *GetSendRequestByUserIdRequest,
	callback chan<- GetSendRequestByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/sendBox/guild/{guildModelName}/{targetGuildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
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

	go getSendRequestByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetSendRequestByUserId(
	request *GetSendRequestByUserIdRequest,
) (*GetSendRequestByUserIdResult, error) {
	callback := make(chan GetSendRequestByUserIdAsyncResult, 1)
	go p.GetSendRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SendRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendRequestResult
	if asyncResult.Err != nil {
		callback <- SendRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SendRequestAsync(
	request *SendRequestRequest,
	callback chan<- SendRequestAsyncResult,
) {
	path := "/{namespaceName}/user/me/sendBox/guild/{guildModelName}/{targetGuildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

	go sendRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SendRequest(
	request *SendRequestRequest,
) (*SendRequestResult, error) {
	callback := make(chan SendRequestAsyncResult, 1)
	go p.SendRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendRequestByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- SendRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendRequestByUserIdResult
	if asyncResult.Err != nil {
		callback <- SendRequestByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SendRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SendRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) SendRequestByUserIdAsync(
	request *SendRequestByUserIdRequest,
	callback chan<- SendRequestByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/sendBox/guild/{guildModelName}/{targetGuildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
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

	go sendRequestByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) SendRequestByUserId(
	request *SendRequestByUserIdRequest,
) (*SendRequestByUserIdResult, error) {
	callback := make(chan SendRequestByUserIdAsyncResult, 1)
	go p.SendRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRequestAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRequestAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRequestAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRequestResult
	if asyncResult.Err != nil {
		callback <- DeleteRequestAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRequestAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRequestAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteRequestAsync(
	request *DeleteRequestRequest,
	callback chan<- DeleteRequestAsyncResult,
) {
	path := "/{namespaceName}/user/me/sendBox/guild/{guildModelName}/{targetGuildName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
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

	go deleteRequestAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteRequest(
	request *DeleteRequestRequest,
) (*DeleteRequestResult, error) {
	callback := make(chan DeleteRequestAsyncResult, 1)
	go p.DeleteRequestAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRequestByUserIdAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRequestByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRequestByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRequestByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteRequestByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteRequestByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteRequestByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteRequestByUserIdAsync(
	request *DeleteRequestByUserIdRequest,
	callback chan<- DeleteRequestByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/sendBox/guild/{guildModelName}/{targetGuildName}"
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
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.TargetGuildName != nil && *request.TargetGuildName != "" {
		path = strings.ReplaceAll(path, "{targetGuildName}", core.ToString(*request.TargetGuildName))
	} else {
		path = strings.ReplaceAll(path, "{targetGuildName}", "null")
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

	go deleteRequestByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteRequestByUserId(
	request *DeleteRequestByUserIdRequest,
) (*DeleteRequestByUserIdResult, error) {
	callback := make(chan DeleteRequestByUserIdAsyncResult, 1)
	go p.DeleteRequestByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeIgnoreUsersAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeIgnoreUsersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIgnoreUsersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIgnoreUsersResult
	if asyncResult.Err != nil {
		callback <- DescribeIgnoreUsersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeIgnoreUsersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeIgnoreUsersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeIgnoreUsersAsync(
	request *DescribeIgnoreUsersRequest,
	callback chan<- DescribeIgnoreUsersAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/ignore/user"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go describeIgnoreUsersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeIgnoreUsers(
	request *DescribeIgnoreUsersRequest,
) (*DescribeIgnoreUsersResult, error) {
	callback := make(chan DescribeIgnoreUsersAsyncResult, 1)
	go p.DescribeIgnoreUsersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeIgnoreUsersByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeIgnoreUsersByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIgnoreUsersByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeIgnoreUsersByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DescribeIgnoreUsersByGuildNameAsync(
	request *DescribeIgnoreUsersByGuildNameRequest,
	callback chan<- DescribeIgnoreUsersByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/ignore/user"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go describeIgnoreUsersByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DescribeIgnoreUsersByGuildName(
	request *DescribeIgnoreUsersByGuildNameRequest,
) (*DescribeIgnoreUsersByGuildNameResult, error) {
	callback := make(chan DescribeIgnoreUsersByGuildNameAsyncResult, 1)
	go p.DescribeIgnoreUsersByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getIgnoreUserAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIgnoreUserResult
	if asyncResult.Err != nil {
		callback <- GetIgnoreUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetIgnoreUserAsync(
	request *GetIgnoreUserRequest,
	callback chan<- GetIgnoreUserAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go getIgnoreUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetIgnoreUser(
	request *GetIgnoreUserRequest,
) (*GetIgnoreUserResult, error) {
	callback := make(chan GetIgnoreUserAsyncResult, 1)
	go p.GetIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getIgnoreUserByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- GetIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIgnoreUserByGuildNameResult
	if asyncResult.Err != nil {
		callback <- GetIgnoreUserByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) GetIgnoreUserByGuildNameAsync(
	request *GetIgnoreUserByGuildNameRequest,
	callback chan<- GetIgnoreUserByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go getIgnoreUserByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) GetIgnoreUserByGuildName(
	request *GetIgnoreUserByGuildNameRequest,
) (*GetIgnoreUserByGuildNameResult, error) {
	callback := make(chan GetIgnoreUserByGuildNameAsyncResult, 1)
	go p.GetIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addIgnoreUserAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AddIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddIgnoreUserResult
	if asyncResult.Err != nil {
		callback <- AddIgnoreUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AddIgnoreUserAsync(
	request *AddIgnoreUserRequest,
	callback chan<- AddIgnoreUserAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go addIgnoreUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) AddIgnoreUser(
	request *AddIgnoreUserRequest,
) (*AddIgnoreUserResult, error) {
	callback := make(chan AddIgnoreUserAsyncResult, 1)
	go p.AddIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addIgnoreUserByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- AddIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddIgnoreUserByGuildNameResult
	if asyncResult.Err != nil {
		callback <- AddIgnoreUserByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) AddIgnoreUserByGuildNameAsync(
	request *AddIgnoreUserByGuildNameRequest,
	callback chan<- AddIgnoreUserByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go addIgnoreUserByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) AddIgnoreUserByGuildName(
	request *AddIgnoreUserByGuildNameRequest,
) (*AddIgnoreUserByGuildNameResult, error) {
	callback := make(chan AddIgnoreUserByGuildNameAsyncResult, 1)
	go p.AddIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteIgnoreUserAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteIgnoreUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIgnoreUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIgnoreUserResult
	if asyncResult.Err != nil {
		callback <- DeleteIgnoreUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteIgnoreUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteIgnoreUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteIgnoreUserAsync(
	request *DeleteIgnoreUserRequest,
	callback chan<- DeleteIgnoreUserAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/me/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
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

	go deleteIgnoreUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteIgnoreUser(
	request *DeleteIgnoreUserRequest,
) (*DeleteIgnoreUserResult, error) {
	callback := make(chan DeleteIgnoreUserAsyncResult, 1)
	go p.DeleteIgnoreUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteIgnoreUserByGuildNameAsyncHandler(
	client Gs2GuildRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteIgnoreUserByGuildNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIgnoreUserByGuildNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIgnoreUserByGuildNameResult
	if asyncResult.Err != nil {
		callback <- DeleteIgnoreUserByGuildNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteIgnoreUserByGuildNameAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteIgnoreUserByGuildNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GuildRestClient) DeleteIgnoreUserByGuildNameAsync(
	request *DeleteIgnoreUserByGuildNameRequest,
	callback chan<- DeleteIgnoreUserByGuildNameAsyncResult,
) {
	path := "/{namespaceName}/guild/{guildModelName}/{guildName}/ignore/user/{userId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.GuildModelName != nil && *request.GuildModelName != "" {
		path = strings.ReplaceAll(path, "{guildModelName}", core.ToString(*request.GuildModelName))
	} else {
		path = strings.ReplaceAll(path, "{guildModelName}", "null")
	}
	if request.GuildName != nil && *request.GuildName != "" {
		path = strings.ReplaceAll(path, "{guildName}", core.ToString(*request.GuildName))
	} else {
		path = strings.ReplaceAll(path, "{guildName}", "null")
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

	go deleteIgnoreUserByGuildNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("guild").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GuildRestClient) DeleteIgnoreUserByGuildName(
	request *DeleteIgnoreUserByGuildNameRequest,
) (*DeleteIgnoreUserByGuildNameResult, error) {
	callback := make(chan DeleteIgnoreUserByGuildNameAsyncResult, 1)
	go p.DeleteIgnoreUserByGuildNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
