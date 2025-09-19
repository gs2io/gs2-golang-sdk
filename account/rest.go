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

package account

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2AccountRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2AccountRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeNamespaces(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) CreateNamespaceAsync(
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
	if request.ChangePasswordIfTakeOver != nil {
		bodies["changePasswordIfTakeOver"] = *request.ChangePasswordIfTakeOver
	}
	if request.DifferentUserIdForLoginAndDataRetention != nil {
		bodies["differentUserIdForLoginAndDataRetention"] = *request.DifferentUserIdForLoginAndDataRetention
	}
	if request.CreateAccountScript != nil {
		bodies["createAccountScript"] = request.CreateAccountScript.ToDict()
	}
	if request.AuthenticationScript != nil {
		bodies["authenticationScript"] = request.AuthenticationScript.ToDict()
	}
	if request.CreateTakeOverScript != nil {
		bodies["createTakeOverScript"] = request.CreateTakeOverScript.ToDict()
	}
	if request.DoTakeOverScript != nil {
		bodies["doTakeOverScript"] = request.DoTakeOverScript.ToDict()
	}
	if request.BanScript != nil {
		bodies["banScript"] = request.BanScript.ToDict()
	}
	if request.UnBanScript != nil {
		bodies["unBanScript"] = request.UnBanScript.ToDict()
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateNamespace(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetNamespaceStatus(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetNamespace(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) UpdateNamespaceAsync(
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
	if request.ChangePasswordIfTakeOver != nil {
		bodies["changePasswordIfTakeOver"] = *request.ChangePasswordIfTakeOver
	}
	if request.CreateAccountScript != nil {
		bodies["createAccountScript"] = request.CreateAccountScript.ToDict()
	}
	if request.AuthenticationScript != nil {
		bodies["authenticationScript"] = request.AuthenticationScript.ToDict()
	}
	if request.CreateTakeOverScript != nil {
		bodies["createTakeOverScript"] = request.CreateTakeOverScript.ToDict()
	}
	if request.DoTakeOverScript != nil {
		bodies["doTakeOverScript"] = request.DoTakeOverScript.ToDict()
	}
	if request.BanScript != nil {
		bodies["banScript"] = request.BanScript.ToDict()
	}
	if request.UnBanScript != nil {
		bodies["unBanScript"] = request.UnBanScript.ToDict()
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateNamespace(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteNamespace(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) GetServiceVersionAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetServiceVersion(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DumpUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CheckDumpUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CleanUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CheckCleanUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) PrepareImportUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) ImportUserDataByUserId(
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
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CheckImportUserDataByUserId(
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

func describeAccountsAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAccountsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAccountsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAccountsResult
	if asyncResult.Err != nil {
		callback <- DescribeAccountsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAccountsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAccountsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribeAccountsAsync(
	request *DescribeAccountsRequest,
	callback chan<- DescribeAccountsAsyncResult,
) {
	path := "/{namespaceName}/account"
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

	go describeAccountsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeAccounts(
	request *DescribeAccountsRequest,
) (*DescribeAccountsResult, error) {
	callback := make(chan DescribeAccountsAsyncResult, 1)
	go p.DescribeAccountsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createAccountAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateAccountResult
	if asyncResult.Err != nil {
		callback <- CreateAccountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateAccountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateAccountAsync(
	request *CreateAccountRequest,
	callback chan<- CreateAccountAsyncResult,
) {
	path := "/{namespaceName}/account"
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

	go createAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateAccount(
	request *CreateAccountRequest,
) (*CreateAccountResult, error) {
	callback := make(chan CreateAccountAsyncResult, 1)
	go p.CreateAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateTimeOffsetAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateTimeOffsetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTimeOffsetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTimeOffsetResult
	if asyncResult.Err != nil {
		callback <- UpdateTimeOffsetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateTimeOffsetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateTimeOffsetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateTimeOffsetAsync(
	request *UpdateTimeOffsetRequest,
	callback chan<- UpdateTimeOffsetAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/time_offset"
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
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
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

	go updateTimeOffsetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateTimeOffset(
	request *UpdateTimeOffsetRequest,
) (*UpdateTimeOffsetResult, error) {
	callback := make(chan UpdateTimeOffsetAsyncResult, 1)
	go p.UpdateTimeOffsetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBannedAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBannedAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBannedAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBannedResult
	if asyncResult.Err != nil {
		callback <- UpdateBannedAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBannedAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBannedAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateBannedAsync(
	request *UpdateBannedRequest,
	callback chan<- UpdateBannedAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/banned"
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
	if request.Banned != nil {
		bodies["banned"] = *request.Banned
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

	go updateBannedAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateBanned(
	request *UpdateBannedRequest,
) (*UpdateBannedResult, error) {
	callback := make(chan UpdateBannedAsyncResult, 1)
	go p.UpdateBannedAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func addBanAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- AddBanAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddBanAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddBanResult
	if asyncResult.Err != nil {
		callback <- AddBanAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AddBanAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AddBanAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) AddBanAsync(
	request *AddBanRequest,
	callback chan<- AddBanAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/ban"
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
	if request.BanStatus != nil {
		bodies["banStatus"] = request.BanStatus.ToDict()
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

	go addBanAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) AddBan(
	request *AddBanRequest,
) (*AddBanResult, error) {
	callback := make(chan AddBanAsyncResult, 1)
	go p.AddBanAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func removeBanAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- RemoveBanAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RemoveBanAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RemoveBanResult
	if asyncResult.Err != nil {
		callback <- RemoveBanAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- RemoveBanAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- RemoveBanAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) RemoveBanAsync(
	request *RemoveBanRequest,
	callback chan<- RemoveBanAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/ban/{banStatusName}"
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
	if request.BanStatusName != nil && *request.BanStatusName != "" {
		path = strings.ReplaceAll(path, "{banStatusName}", core.ToString(*request.BanStatusName))
	} else {
		path = strings.ReplaceAll(path, "{banStatusName}", "null")
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

	go removeBanAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) RemoveBan(
	request *RemoveBanRequest,
) (*RemoveBanResult, error) {
	callback := make(chan RemoveBanAsyncResult, 1)
	go p.RemoveBanAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAccountAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAccountResult
	if asyncResult.Err != nil {
		callback <- GetAccountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAccountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetAccountAsync(
	request *GetAccountRequest,
	callback chan<- GetAccountAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}"
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
	if request.IncludeLastAuthenticatedAt != nil {
		queryStrings["includeLastAuthenticatedAt"] = core.ToString(*request.IncludeLastAuthenticatedAt)
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

	go getAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetAccount(
	request *GetAccountRequest,
) (*GetAccountResult, error) {
	callback := make(chan GetAccountAsyncResult, 1)
	go p.GetAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAccountAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteAccountResult
	if asyncResult.Err != nil {
		callback <- DeleteAccountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteAccountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteAccountAsync(
	request *DeleteAccountRequest,
	callback chan<- DeleteAccountAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}"
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

	go deleteAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteAccount(
	request *DeleteAccountRequest,
) (*DeleteAccountResult, error) {
	callback := make(chan DeleteAccountAsyncResult, 1)
	go p.DeleteAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func authenticationAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- AuthenticationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AuthenticationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AuthenticationResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.banned.infinity" {
				asyncResult.Err = gs2err.SetClientError(BannedInfinity{})
			}
		}
		callback <- AuthenticationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AuthenticationAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AuthenticationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) AuthenticationAsync(
	request *AuthenticationRequest,
	callback chan<- AuthenticationAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}"
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
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go authenticationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) Authentication(
	request *AuthenticationRequest,
) (*AuthenticationResult, error) {
	callback := make(chan AuthenticationAsyncResult, 1)
	go p.AuthenticationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTakeOversAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTakeOversAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOversAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOversResult
	if asyncResult.Err != nil {
		callback <- DescribeTakeOversAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeTakeOversAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeTakeOversAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribeTakeOversAsync(
	request *DescribeTakeOversRequest,
	callback chan<- DescribeTakeOversAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover"
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

	go describeTakeOversAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeTakeOvers(
	request *DescribeTakeOversRequest,
) (*DescribeTakeOversResult, error) {
	callback := make(chan DescribeTakeOversAsyncResult, 1)
	go p.DescribeTakeOversAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTakeOversByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTakeOversByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOversByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOversByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeTakeOversByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeTakeOversByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeTakeOversByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribeTakeOversByUserIdAsync(
	request *DescribeTakeOversByUserIdRequest,
	callback chan<- DescribeTakeOversByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover"
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

	go describeTakeOversByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeTakeOversByUserId(
	request *DescribeTakeOversByUserIdRequest,
) (*DescribeTakeOversByUserIdResult, error) {
	callback := make(chan DescribeTakeOversByUserIdAsyncResult, 1)
	go p.DescribeTakeOversByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createTakeOverAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverResult
	if asyncResult.Err != nil {
		callback <- CreateTakeOverAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateTakeOverAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateTakeOverAsync(
	request *CreateTakeOverRequest,
	callback chan<- CreateTakeOverAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		bodies["userIdentifier"] = *request.UserIdentifier
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go createTakeOverAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateTakeOver(
	request *CreateTakeOverRequest,
) (*CreateTakeOverResult, error) {
	callback := make(chan CreateTakeOverAsyncResult, 1)
	go p.CreateTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createTakeOverByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateTakeOverByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateTakeOverByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateTakeOverByUserIdAsync(
	request *CreateTakeOverByUserIdRequest,
	callback chan<- CreateTakeOverByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover"
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
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		bodies["userIdentifier"] = *request.UserIdentifier
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go createTakeOverByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateTakeOverByUserId(
	request *CreateTakeOverByUserIdRequest,
) (*CreateTakeOverByUserIdResult, error) {
	callback := make(chan CreateTakeOverByUserIdAsyncResult, 1)
	go p.CreateTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createTakeOverOpenIdConnectAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateTakeOverOpenIdConnectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverOpenIdConnectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverOpenIdConnectResult
	if asyncResult.Err != nil {
		callback <- CreateTakeOverOpenIdConnectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateTakeOverOpenIdConnectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateTakeOverOpenIdConnectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateTakeOverOpenIdConnectAsync(
	request *CreateTakeOverOpenIdConnectRequest,
	callback chan<- CreateTakeOverOpenIdConnectAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover/openIdConnect"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.IdToken != nil && *request.IdToken != "" {
		bodies["idToken"] = *request.IdToken
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

	go createTakeOverOpenIdConnectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateTakeOverOpenIdConnect(
	request *CreateTakeOverOpenIdConnectRequest,
) (*CreateTakeOverOpenIdConnectResult, error) {
	callback := make(chan CreateTakeOverOpenIdConnectAsyncResult, 1)
	go p.CreateTakeOverOpenIdConnectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createTakeOverOpenIdConnectAndByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverOpenIdConnectAndByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateTakeOverOpenIdConnectAndByUserIdAsync(
	request *CreateTakeOverOpenIdConnectAndByUserIdRequest,
	callback chan<- CreateTakeOverOpenIdConnectAndByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover/openIdConnect"
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
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.IdToken != nil && *request.IdToken != "" {
		bodies["idToken"] = *request.IdToken
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

	go createTakeOverOpenIdConnectAndByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateTakeOverOpenIdConnectAndByUserId(
	request *CreateTakeOverOpenIdConnectAndByUserIdRequest,
) (*CreateTakeOverOpenIdConnectAndByUserIdResult, error) {
	callback := make(chan CreateTakeOverOpenIdConnectAndByUserIdAsyncResult, 1)
	go p.CreateTakeOverOpenIdConnectAndByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTakeOverAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverResult
	if asyncResult.Err != nil {
		callback <- GetTakeOverAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTakeOverAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetTakeOverAsync(
	request *GetTakeOverRequest,
	callback chan<- GetTakeOverAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getTakeOverAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetTakeOver(
	request *GetTakeOverRequest,
) (*GetTakeOverResult, error) {
	callback := make(chan GetTakeOverAsyncResult, 1)
	go p.GetTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTakeOverByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetTakeOverByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTakeOverByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetTakeOverByUserIdAsync(
	request *GetTakeOverByUserIdRequest,
	callback chan<- GetTakeOverByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover/type/{type}"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getTakeOverByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetTakeOverByUserId(
	request *GetTakeOverByUserIdRequest,
) (*GetTakeOverByUserIdResult, error) {
	callback := make(chan GetTakeOverByUserIdAsyncResult, 1)
	go p.GetTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateTakeOverAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTakeOverResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- UpdateTakeOverAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateTakeOverAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateTakeOverAsync(
	request *UpdateTakeOverRequest,
	callback chan<- UpdateTakeOverAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.OldPassword != nil && *request.OldPassword != "" {
		bodies["oldPassword"] = *request.OldPassword
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go updateTakeOverAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateTakeOver(
	request *UpdateTakeOverRequest,
) (*UpdateTakeOverResult, error) {
	callback := make(chan UpdateTakeOverAsyncResult, 1)
	go p.UpdateTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateTakeOverByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTakeOverByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- UpdateTakeOverByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateTakeOverByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateTakeOverByUserIdAsync(
	request *UpdateTakeOverByUserIdRequest,
	callback chan<- UpdateTakeOverByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover/type/{type}"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.OldPassword != nil && *request.OldPassword != "" {
		bodies["oldPassword"] = *request.OldPassword
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go updateTakeOverByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateTakeOverByUserId(
	request *UpdateTakeOverByUserIdRequest,
) (*UpdateTakeOverByUserIdResult, error) {
	callback := make(chan UpdateTakeOverByUserIdAsyncResult, 1)
	go p.UpdateTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTakeOverAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverResult
	if asyncResult.Err != nil {
		callback <- DeleteTakeOverAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteTakeOverAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteTakeOverAsync(
	request *DeleteTakeOverRequest,
	callback chan<- DeleteTakeOverAsyncResult,
) {
	path := "/{namespaceName}/account/me/takeover/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go deleteTakeOverAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteTakeOver(
	request *DeleteTakeOverRequest,
) (*DeleteTakeOverResult, error) {
	callback := make(chan DeleteTakeOverAsyncResult, 1)
	go p.DeleteTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTakeOverByUserIdentifierAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTakeOverByUserIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverByUserIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverByUserIdentifierResult
	if asyncResult.Err != nil {
		callback <- DeleteTakeOverByUserIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteTakeOverByUserIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteTakeOverByUserIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteTakeOverByUserIdentifierAsync(
	request *DeleteTakeOverByUserIdentifierRequest,
	callback chan<- DeleteTakeOverByUserIdentifierAsyncResult,
) {
	path := "/{namespaceName}/takeover/type/{type}/userIdentifier/{userIdentifier}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		path = strings.ReplaceAll(path, "{userIdentifier}", core.ToString(*request.UserIdentifier))
	} else {
		path = strings.ReplaceAll(path, "{userIdentifier}", "null")
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

	go deleteTakeOverByUserIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteTakeOverByUserIdentifier(
	request *DeleteTakeOverByUserIdentifierRequest,
) (*DeleteTakeOverByUserIdentifierResult, error) {
	callback := make(chan DeleteTakeOverByUserIdentifierAsyncResult, 1)
	go p.DeleteTakeOverByUserIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTakeOverByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTakeOverByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteTakeOverByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteTakeOverByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteTakeOverByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteTakeOverByUserIdAsync(
	request *DeleteTakeOverByUserIdRequest,
	callback chan<- DeleteTakeOverByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/takeover/type/{type}/takeover"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go deleteTakeOverByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteTakeOverByUserId(
	request *DeleteTakeOverByUserIdRequest,
) (*DeleteTakeOverByUserIdResult, error) {
	callback := make(chan DeleteTakeOverByUserIdAsyncResult, 1)
	go p.DeleteTakeOverByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doTakeOverAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DoTakeOverAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoTakeOverAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoTakeOverResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "account.password.invalid" {
				asyncResult.Err = gs2err.SetClientError(PasswordIncorrect{})
			}
		}
		callback <- DoTakeOverAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoTakeOverAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoTakeOverAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DoTakeOverAsync(
	request *DoTakeOverRequest,
	callback chan<- DoTakeOverAsyncResult,
) {
	path := "/{namespaceName}/takeover/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		bodies["userIdentifier"] = *request.UserIdentifier
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go doTakeOverAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DoTakeOver(
	request *DoTakeOverRequest,
) (*DoTakeOverResult, error) {
	callback := make(chan DoTakeOverAsyncResult, 1)
	go p.DoTakeOverAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doTakeOverOpenIdConnectAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DoTakeOverOpenIdConnectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoTakeOverOpenIdConnectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoTakeOverOpenIdConnectResult
	if asyncResult.Err != nil {
		callback <- DoTakeOverOpenIdConnectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DoTakeOverOpenIdConnectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DoTakeOverOpenIdConnectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DoTakeOverOpenIdConnectAsync(
	request *DoTakeOverOpenIdConnectRequest,
	callback chan<- DoTakeOverOpenIdConnectAsyncResult,
) {
	path := "/{namespaceName}/takeover/type/{type}/openIdConnect"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.IdToken != nil && *request.IdToken != "" {
		bodies["idToken"] = *request.IdToken
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

	go doTakeOverOpenIdConnectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DoTakeOverOpenIdConnect(
	request *DoTakeOverOpenIdConnectRequest,
) (*DoTakeOverOpenIdConnectResult, error) {
	callback := make(chan DoTakeOverOpenIdConnectAsyncResult, 1)
	go p.DoTakeOverOpenIdConnectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getAuthorizationUrlAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetAuthorizationUrlAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAuthorizationUrlAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAuthorizationUrlResult
	if asyncResult.Err != nil {
		callback <- GetAuthorizationUrlAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAuthorizationUrlAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAuthorizationUrlAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetAuthorizationUrlAsync(
	request *GetAuthorizationUrlRequest,
	callback chan<- GetAuthorizationUrlAsyncResult,
) {
	path := "/{namespaceName}/type/{type}/authorization/url"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getAuthorizationUrlAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetAuthorizationUrl(
	request *GetAuthorizationUrlRequest,
) (*GetAuthorizationUrlResult, error) {
	callback := make(chan GetAuthorizationUrlAsyncResult, 1)
	go p.GetAuthorizationUrlAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePlatformIdsAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePlatformIdsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePlatformIdsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePlatformIdsResult
	if asyncResult.Err != nil {
		callback <- DescribePlatformIdsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePlatformIdsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePlatformIdsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribePlatformIdsAsync(
	request *DescribePlatformIdsRequest,
	callback chan<- DescribePlatformIdsAsyncResult,
) {
	path := "/{namespaceName}/account/me/platformId"
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

	go describePlatformIdsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribePlatformIds(
	request *DescribePlatformIdsRequest,
) (*DescribePlatformIdsResult, error) {
	callback := make(chan DescribePlatformIdsAsyncResult, 1)
	go p.DescribePlatformIdsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePlatformIdsByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePlatformIdsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePlatformIdsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePlatformIdsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribePlatformIdsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePlatformIdsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePlatformIdsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribePlatformIdsByUserIdAsync(
	request *DescribePlatformIdsByUserIdRequest,
	callback chan<- DescribePlatformIdsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/platformId"
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

	go describePlatformIdsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribePlatformIdsByUserId(
	request *DescribePlatformIdsByUserIdRequest,
) (*DescribePlatformIdsByUserIdResult, error) {
	callback := make(chan DescribePlatformIdsByUserIdAsyncResult, 1)
	go p.DescribePlatformIdsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPlatformIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePlatformIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePlatformIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePlatformIdResult
	if asyncResult.Err != nil {
		callback <- CreatePlatformIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreatePlatformIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreatePlatformIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreatePlatformIdAsync(
	request *CreatePlatformIdRequest,
	callback chan<- CreatePlatformIdAsyncResult,
) {
	path := "/{namespaceName}/account/me/platformId"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		bodies["userIdentifier"] = *request.UserIdentifier
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

	go createPlatformIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreatePlatformId(
	request *CreatePlatformIdRequest,
) (*CreatePlatformIdResult, error) {
	callback := make(chan CreatePlatformIdAsyncResult, 1)
	go p.CreatePlatformIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPlatformIdByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePlatformIdByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePlatformIdByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePlatformIdByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreatePlatformIdByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreatePlatformIdByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreatePlatformIdByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreatePlatformIdByUserIdAsync(
	request *CreatePlatformIdByUserIdRequest,
	callback chan<- CreatePlatformIdByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/platformId"
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
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		bodies["userIdentifier"] = *request.UserIdentifier
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

	go createPlatformIdByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreatePlatformIdByUserId(
	request *CreatePlatformIdByUserIdRequest,
) (*CreatePlatformIdByUserIdResult, error) {
	callback := make(chan CreatePlatformIdByUserIdAsyncResult, 1)
	go p.CreatePlatformIdByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPlatformIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetPlatformIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPlatformIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPlatformIdResult
	if asyncResult.Err != nil {
		callback <- GetPlatformIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPlatformIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPlatformIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetPlatformIdAsync(
	request *GetPlatformIdRequest,
	callback chan<- GetPlatformIdAsyncResult,
) {
	path := "/{namespaceName}/account/me/platformId/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getPlatformIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetPlatformId(
	request *GetPlatformIdRequest,
) (*GetPlatformIdResult, error) {
	callback := make(chan GetPlatformIdAsyncResult, 1)
	go p.GetPlatformIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPlatformIdByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetPlatformIdByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPlatformIdByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPlatformIdByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetPlatformIdByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPlatformIdByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPlatformIdByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetPlatformIdByUserIdAsync(
	request *GetPlatformIdByUserIdRequest,
	callback chan<- GetPlatformIdByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/platformId/type/{type}"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getPlatformIdByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetPlatformIdByUserId(
	request *GetPlatformIdByUserIdRequest,
) (*GetPlatformIdByUserIdResult, error) {
	callback := make(chan GetPlatformIdByUserIdAsyncResult, 1)
	go p.GetPlatformIdByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func findPlatformIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- FindPlatformIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FindPlatformIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FindPlatformIdResult
	if asyncResult.Err != nil {
		callback <- FindPlatformIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FindPlatformIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FindPlatformIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) FindPlatformIdAsync(
	request *FindPlatformIdRequest,
	callback chan<- FindPlatformIdAsyncResult,
) {
	path := "/{namespaceName}/account/me/platformId/type/{type}/userIdentifier/{userIdentifier}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		path = strings.ReplaceAll(path, "{userIdentifier}", core.ToString(*request.UserIdentifier))
	} else {
		path = strings.ReplaceAll(path, "{userIdentifier}", "null")
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

	go findPlatformIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) FindPlatformId(
	request *FindPlatformIdRequest,
) (*FindPlatformIdResult, error) {
	callback := make(chan FindPlatformIdAsyncResult, 1)
	go p.FindPlatformIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func findPlatformIdByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- FindPlatformIdByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FindPlatformIdByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FindPlatformIdByUserIdResult
	if asyncResult.Err != nil {
		callback <- FindPlatformIdByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FindPlatformIdByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FindPlatformIdByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) FindPlatformIdByUserIdAsync(
	request *FindPlatformIdByUserIdRequest,
	callback chan<- FindPlatformIdByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/platformId/type/{type}/userIdentifier/{userIdentifier}"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		path = strings.ReplaceAll(path, "{userIdentifier}", core.ToString(*request.UserIdentifier))
	} else {
		path = strings.ReplaceAll(path, "{userIdentifier}", "null")
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

	go findPlatformIdByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) FindPlatformIdByUserId(
	request *FindPlatformIdByUserIdRequest,
) (*FindPlatformIdByUserIdResult, error) {
	callback := make(chan FindPlatformIdByUserIdAsyncResult, 1)
	go p.FindPlatformIdByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePlatformIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePlatformIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePlatformIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePlatformIdResult
	if asyncResult.Err != nil {
		callback <- DeletePlatformIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePlatformIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePlatformIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeletePlatformIdAsync(
	request *DeletePlatformIdRequest,
	callback chan<- DeletePlatformIdAsyncResult,
) {
	path := "/{namespaceName}/account/me/platformId/type/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.UserIdentifier != nil {
		queryStrings["userIdentifier"] = core.ToString(*request.UserIdentifier)
	}
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

	go deletePlatformIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeletePlatformId(
	request *DeletePlatformIdRequest,
) (*DeletePlatformIdResult, error) {
	callback := make(chan DeletePlatformIdAsyncResult, 1)
	go p.DeletePlatformIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePlatformIdByUserIdentifierAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePlatformIdByUserIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePlatformIdByUserIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePlatformIdByUserIdentifierResult
	if asyncResult.Err != nil {
		callback <- DeletePlatformIdByUserIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePlatformIdByUserIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePlatformIdByUserIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeletePlatformIdByUserIdentifierAsync(
	request *DeletePlatformIdByUserIdentifierRequest,
	callback chan<- DeletePlatformIdByUserIdentifierAsyncResult,
) {
	path := "/{namespaceName}/platformId/type/{type}/userIdentifier/{userIdentifier}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}
	if request.UserIdentifier != nil && *request.UserIdentifier != "" {
		path = strings.ReplaceAll(path, "{userIdentifier}", core.ToString(*request.UserIdentifier))
	} else {
		path = strings.ReplaceAll(path, "{userIdentifier}", "null")
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

	go deletePlatformIdByUserIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeletePlatformIdByUserIdentifier(
	request *DeletePlatformIdByUserIdentifierRequest,
) (*DeletePlatformIdByUserIdentifierResult, error) {
	callback := make(chan DeletePlatformIdByUserIdentifierAsyncResult, 1)
	go p.DeletePlatformIdByUserIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePlatformIdByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePlatformIdByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePlatformIdByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePlatformIdByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeletePlatformIdByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePlatformIdByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePlatformIdByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeletePlatformIdByUserIdAsync(
	request *DeletePlatformIdByUserIdRequest,
	callback chan<- DeletePlatformIdByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/platformId/type/{type}/platformId"
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
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go deletePlatformIdByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeletePlatformIdByUserId(
	request *DeletePlatformIdByUserIdRequest,
) (*DeletePlatformIdByUserIdResult, error) {
	callback := make(chan DeletePlatformIdByUserIdAsyncResult, 1)
	go p.DeletePlatformIdByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDataOwnerByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetDataOwnerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataOwnerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataOwnerByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetDataOwnerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDataOwnerByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDataOwnerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetDataOwnerByUserIdAsync(
	request *GetDataOwnerByUserIdRequest,
	callback chan<- GetDataOwnerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/dataOwner"
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

	go getDataOwnerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetDataOwnerByUserId(
	request *GetDataOwnerByUserIdRequest,
) (*GetDataOwnerByUserIdResult, error) {
	callback := make(chan GetDataOwnerByUserIdAsyncResult, 1)
	go p.GetDataOwnerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateDataOwnerByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateDataOwnerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDataOwnerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDataOwnerByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateDataOwnerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateDataOwnerByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateDataOwnerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateDataOwnerByUserIdAsync(
	request *UpdateDataOwnerByUserIdRequest,
	callback chan<- UpdateDataOwnerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/dataOwner"
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
	if request.DataOwnerName != nil && *request.DataOwnerName != "" {
		bodies["dataOwnerName"] = *request.DataOwnerName
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

	go updateDataOwnerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateDataOwnerByUserId(
	request *UpdateDataOwnerByUserIdRequest,
) (*UpdateDataOwnerByUserIdResult, error) {
	callback := make(chan UpdateDataOwnerByUserIdAsyncResult, 1)
	go p.UpdateDataOwnerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteDataOwnerByUserIdAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDataOwnerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataOwnerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataOwnerByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteDataOwnerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteDataOwnerByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteDataOwnerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteDataOwnerByUserIdAsync(
	request *DeleteDataOwnerByUserIdRequest,
	callback chan<- DeleteDataOwnerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/account/{userId}/dataOwner"
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

	go deleteDataOwnerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteDataOwnerByUserId(
	request *DeleteDataOwnerByUserIdRequest,
) (*DeleteDataOwnerByUserIdResult, error) {
	callback := make(chan DeleteDataOwnerByUserIdAsyncResult, 1)
	go p.DeleteDataOwnerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTakeOverTypeModelsAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTakeOverTypeModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOverTypeModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOverTypeModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeTakeOverTypeModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeTakeOverTypeModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeTakeOverTypeModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribeTakeOverTypeModelsAsync(
	request *DescribeTakeOverTypeModelsRequest,
	callback chan<- DescribeTakeOverTypeModelsAsyncResult,
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

	go describeTakeOverTypeModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeTakeOverTypeModels(
	request *DescribeTakeOverTypeModelsRequest,
) (*DescribeTakeOverTypeModelsResult, error) {
	callback := make(chan DescribeTakeOverTypeModelsAsyncResult, 1)
	go p.DescribeTakeOverTypeModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTakeOverTypeModelAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetTakeOverTypeModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverTypeModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverTypeModelResult
	if asyncResult.Err != nil {
		callback <- GetTakeOverTypeModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTakeOverTypeModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTakeOverTypeModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetTakeOverTypeModelAsync(
	request *GetTakeOverTypeModelRequest,
	callback chan<- GetTakeOverTypeModelAsyncResult,
) {
	path := "/{namespaceName}/model/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getTakeOverTypeModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetTakeOverTypeModel(
	request *GetTakeOverTypeModelRequest,
) (*GetTakeOverTypeModelResult, error) {
	callback := make(chan GetTakeOverTypeModelAsyncResult, 1)
	go p.GetTakeOverTypeModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTakeOverTypeModelMastersAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTakeOverTypeModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTakeOverTypeModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTakeOverTypeModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeTakeOverTypeModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeTakeOverTypeModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeTakeOverTypeModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DescribeTakeOverTypeModelMastersAsync(
	request *DescribeTakeOverTypeModelMastersRequest,
	callback chan<- DescribeTakeOverTypeModelMastersAsyncResult,
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

	go describeTakeOverTypeModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DescribeTakeOverTypeModelMasters(
	request *DescribeTakeOverTypeModelMastersRequest,
) (*DescribeTakeOverTypeModelMastersResult, error) {
	callback := make(chan DescribeTakeOverTypeModelMastersAsyncResult, 1)
	go p.DescribeTakeOverTypeModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createTakeOverTypeModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- CreateTakeOverTypeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateTakeOverTypeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateTakeOverTypeModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateTakeOverTypeModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateTakeOverTypeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateTakeOverTypeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) CreateTakeOverTypeModelMasterAsync(
	request *CreateTakeOverTypeModelMasterRequest,
	callback chan<- CreateTakeOverTypeModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Type != nil {
		bodies["type"] = *request.Type
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.OpenIdConnectSetting != nil {
		bodies["openIdConnectSetting"] = request.OpenIdConnectSetting.ToDict()
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

	go createTakeOverTypeModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) CreateTakeOverTypeModelMaster(
	request *CreateTakeOverTypeModelMasterRequest,
) (*CreateTakeOverTypeModelMasterResult, error) {
	callback := make(chan CreateTakeOverTypeModelMasterAsyncResult, 1)
	go p.CreateTakeOverTypeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTakeOverTypeModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetTakeOverTypeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTakeOverTypeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTakeOverTypeModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetTakeOverTypeModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetTakeOverTypeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetTakeOverTypeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetTakeOverTypeModelMasterAsync(
	request *GetTakeOverTypeModelMasterRequest,
	callback chan<- GetTakeOverTypeModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go getTakeOverTypeModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetTakeOverTypeModelMaster(
	request *GetTakeOverTypeModelMasterRequest,
) (*GetTakeOverTypeModelMasterResult, error) {
	callback := make(chan GetTakeOverTypeModelMasterAsyncResult, 1)
	go p.GetTakeOverTypeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateTakeOverTypeModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateTakeOverTypeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateTakeOverTypeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateTakeOverTypeModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateTakeOverTypeModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateTakeOverTypeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateTakeOverTypeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateTakeOverTypeModelMasterAsync(
	request *UpdateTakeOverTypeModelMasterRequest,
	callback chan<- UpdateTakeOverTypeModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.OpenIdConnectSetting != nil {
		bodies["openIdConnectSetting"] = request.OpenIdConnectSetting.ToDict()
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

	go updateTakeOverTypeModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateTakeOverTypeModelMaster(
	request *UpdateTakeOverTypeModelMasterRequest,
) (*UpdateTakeOverTypeModelMasterResult, error) {
	callback := make(chan UpdateTakeOverTypeModelMasterAsyncResult, 1)
	go p.UpdateTakeOverTypeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTakeOverTypeModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTakeOverTypeModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTakeOverTypeModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTakeOverTypeModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteTakeOverTypeModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteTakeOverTypeModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteTakeOverTypeModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) DeleteTakeOverTypeModelMasterAsync(
	request *DeleteTakeOverTypeModelMasterRequest,
	callback chan<- DeleteTakeOverTypeModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{type}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Type != nil {
		path = strings.ReplaceAll(path, "{type}", core.ToString(*request.Type))
	} else {
		path = strings.ReplaceAll(path, "{type}", "null")
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

	go deleteTakeOverTypeModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) DeleteTakeOverTypeModelMaster(
	request *DeleteTakeOverTypeModelMasterRequest,
) (*DeleteTakeOverTypeModelMasterResult, error) {
	callback := make(chan DeleteTakeOverTypeModelMasterAsyncResult, 1)
	go p.DeleteTakeOverTypeModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2AccountRestClient,
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

func (p Gs2AccountRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) ExportMaster(
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

func getCurrentModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCurrentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCurrentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) GetCurrentModelMasterAsync(
	request *GetCurrentModelMasterRequest,
	callback chan<- GetCurrentModelMasterAsyncResult,
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

	go getCurrentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) GetCurrentModelMaster(
	request *GetCurrentModelMasterRequest,
) (*GetCurrentModelMasterResult, error) {
	callback := make(chan GetCurrentModelMasterAsyncResult, 1)
	go p.GetCurrentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateCurrentModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateCurrentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateCurrentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateCurrentModelMasterResult
	if asyncResult.Err != nil {
		callback <- PreUpdateCurrentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateCurrentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateCurrentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) PreUpdateCurrentModelMasterAsync(
	request *PreUpdateCurrentModelMasterRequest,
	callback chan<- PreUpdateCurrentModelMasterAsyncResult,
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

	go preUpdateCurrentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) PreUpdateCurrentModelMaster(
	request *PreUpdateCurrentModelMasterRequest,
) (*PreUpdateCurrentModelMasterResult, error) {
	callback := make(chan PreUpdateCurrentModelMasterAsyncResult, 1)
	go p.PreUpdateCurrentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentModelMasterAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateCurrentModelMasterAsync(
	request *UpdateCurrentModelMasterRequest,
	callback chan<- UpdateCurrentModelMasterAsyncResult,
) {
	if request.Settings != nil {
		res, err := p.PreUpdateCurrentModelMaster(
			&PreUpdateCurrentModelMasterRequest{
				ContextStack:  request.ContextStack,
				NamespaceName: request.NamespaceName,
			},
		)
		if err != nil {
			callback <- UpdateCurrentModelMasterAsyncResult{
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
				callback <- UpdateCurrentModelMasterAsyncResult{
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

	go updateCurrentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateCurrentModelMaster(
	request *UpdateCurrentModelMasterRequest,
) (*UpdateCurrentModelMasterResult, error) {
	callback := make(chan UpdateCurrentModelMasterAsyncResult, 1)
	go p.UpdateCurrentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentModelMasterFromGitHubAsyncHandler(
	client Gs2AccountRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentModelMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentModelMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateCurrentModelMasterFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateCurrentModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AccountRestClient) UpdateCurrentModelMasterFromGitHubAsync(
	request *UpdateCurrentModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentModelMasterFromGitHubAsyncResult,
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

	go updateCurrentModelMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("account").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AccountRestClient) UpdateCurrentModelMasterFromGitHub(
	request *UpdateCurrentModelMasterFromGitHubRequest,
) (*UpdateCurrentModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
