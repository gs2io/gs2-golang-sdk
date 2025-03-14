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

package money2

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2Money2RestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2Money2RestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeNamespaces(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
	path := "/"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.CurrencyUsagePriority != nil && *request.CurrencyUsagePriority != "" {
		bodies["currencyUsagePriority"] = *request.CurrencyUsagePriority
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.SharedFreeCurrency != nil {
		bodies["sharedFreeCurrency"] = *request.SharedFreeCurrency
	}
	if request.PlatformSetting != nil {
		bodies["platformSetting"] = request.PlatformSetting.ToDict()
	}
	if request.DepositBalanceScript != nil {
		bodies["depositBalanceScript"] = request.DepositBalanceScript.ToDict()
	}
	if request.WithdrawBalanceScript != nil {
		bodies["withdrawBalanceScript"] = request.WithdrawBalanceScript.ToDict()
	}
	if request.SubscribeScript != nil && *request.SubscribeScript != "" {
		bodies["subscribeScript"] = *request.SubscribeScript
	}
	if request.RenewScript != nil && *request.RenewScript != "" {
		bodies["renewScript"] = *request.RenewScript
	}
	if request.UnsubscribeScript != nil && *request.UnsubscribeScript != "" {
		bodies["unsubscribeScript"] = *request.UnsubscribeScript
	}
	if request.TakeOverScript != nil {
		bodies["takeOverScript"] = request.TakeOverScript.ToDict()
	}
	if request.ChangeSubscriptionStatusNotification != nil {
		bodies["changeSubscriptionStatusNotification"] = request.ChangeSubscriptionStatusNotification.ToDict()
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CreateNamespace(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetNamespaceStatus(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetNamespace(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) UpdateNamespaceAsync(
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
	if request.CurrencyUsagePriority != nil && *request.CurrencyUsagePriority != "" {
		bodies["currencyUsagePriority"] = *request.CurrencyUsagePriority
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.PlatformSetting != nil {
		bodies["platformSetting"] = request.PlatformSetting.ToDict()
	}
	if request.DepositBalanceScript != nil {
		bodies["depositBalanceScript"] = request.DepositBalanceScript.ToDict()
	}
	if request.WithdrawBalanceScript != nil {
		bodies["withdrawBalanceScript"] = request.WithdrawBalanceScript.ToDict()
	}
	if request.SubscribeScript != nil && *request.SubscribeScript != "" {
		bodies["subscribeScript"] = *request.SubscribeScript
	}
	if request.RenewScript != nil && *request.RenewScript != "" {
		bodies["renewScript"] = *request.RenewScript
	}
	if request.UnsubscribeScript != nil && *request.UnsubscribeScript != "" {
		bodies["unsubscribeScript"] = *request.UnsubscribeScript
	}
	if request.TakeOverScript != nil {
		bodies["takeOverScript"] = request.TakeOverScript.ToDict()
	}
	if request.ChangeSubscriptionStatusNotification != nil {
		bodies["changeSubscriptionStatusNotification"] = request.ChangeSubscriptionStatusNotification.ToDict()
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) UpdateNamespace(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DeleteNamespace(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) DumpUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DumpUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) CheckDumpUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CheckDumpUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) CleanUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CleanUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) CheckCleanUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CheckCleanUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) PrepareImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) PrepareImportUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) ImportUserDataByUserIdAsync(
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) ImportUserDataByUserId(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) CheckImportUserDataByUserIdAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CheckImportUserDataByUserId(
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

func describeWalletsAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeWalletsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWalletsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWalletsResult
	if asyncResult.Err != nil {
		callback <- DescribeWalletsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeWalletsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeWalletsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeWalletsAsync(
	request *DescribeWalletsRequest,
	callback chan<- DescribeWalletsAsyncResult,
) {
	path := "/{namespaceName}/user/me/wallet"
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

	go describeWalletsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeWallets(
	request *DescribeWalletsRequest,
) (*DescribeWalletsResult, error) {
	callback := make(chan DescribeWalletsAsyncResult, 1)
	go p.DescribeWalletsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeWalletsByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeWalletsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWalletsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWalletsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeWalletsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeWalletsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeWalletsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeWalletsByUserIdAsync(
	request *DescribeWalletsByUserIdRequest,
	callback chan<- DescribeWalletsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/wallet"
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

	go describeWalletsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeWalletsByUserId(
	request *DescribeWalletsByUserIdRequest,
) (*DescribeWalletsByUserIdResult, error) {
	callback := make(chan DescribeWalletsByUserIdAsyncResult, 1)
	go p.DescribeWalletsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getWalletAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetWalletAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetWalletAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetWalletResult
	if asyncResult.Err != nil {
		callback <- GetWalletAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetWalletAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetWalletAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetWalletAsync(
	request *GetWalletRequest,
	callback chan<- GetWalletAsyncResult,
) {
	path := "/{namespaceName}/user/me/wallet/{slot}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Slot != nil {
		path = strings.ReplaceAll(path, "{slot}", core.ToString(*request.Slot))
	} else {
		path = strings.ReplaceAll(path, "{slot}", "null")
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

	go getWalletAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetWallet(
	request *GetWalletRequest,
) (*GetWalletResult, error) {
	callback := make(chan GetWalletAsyncResult, 1)
	go p.GetWalletAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getWalletByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetWalletByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetWalletByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetWalletByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetWalletByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetWalletByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetWalletByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetWalletByUserIdAsync(
	request *GetWalletByUserIdRequest,
	callback chan<- GetWalletByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/wallet/{slot}"
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
	if request.Slot != nil {
		path = strings.ReplaceAll(path, "{slot}", core.ToString(*request.Slot))
	} else {
		path = strings.ReplaceAll(path, "{slot}", "null")
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

	go getWalletByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetWalletByUserId(
	request *GetWalletByUserIdRequest,
) (*GetWalletByUserIdResult, error) {
	callback := make(chan GetWalletByUserIdAsyncResult, 1)
	go p.GetWalletByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func depositByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DepositByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DepositByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DepositByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "wallet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
		}
		callback <- DepositByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DepositByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DepositByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DepositByUserIdAsync(
	request *DepositByUserIdRequest,
	callback chan<- DepositByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/wallet/{slot}/deposit"
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
	if request.Slot != nil {
		path = strings.ReplaceAll(path, "{slot}", core.ToString(*request.Slot))
	} else {
		path = strings.ReplaceAll(path, "{slot}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.DepositTransactions != nil {
		var _depositTransactions []interface{}
		for _, item := range request.DepositTransactions {
			_depositTransactions = append(_depositTransactions, item)
		}
		bodies["depositTransactions"] = _depositTransactions
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

	go depositByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DepositByUserId(
	request *DepositByUserIdRequest,
) (*DepositByUserIdResult, error) {
	callback := make(chan DepositByUserIdAsyncResult, 1)
	go p.DepositByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func withdrawAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- WithdrawAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "wallet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "wallet.balance.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- WithdrawAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WithdrawAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) WithdrawAsync(
	request *WithdrawRequest,
	callback chan<- WithdrawAsyncResult,
) {
	path := "/{namespaceName}/user/me/wallet/{slot}/withdraw"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Slot != nil {
		path = strings.ReplaceAll(path, "{slot}", core.ToString(*request.Slot))
	} else {
		path = strings.ReplaceAll(path, "{slot}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.WithdrawCount != nil {
		bodies["withdrawCount"] = *request.WithdrawCount
	}
	if request.PaidOnly != nil {
		bodies["paidOnly"] = *request.PaidOnly
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

	go withdrawAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) Withdraw(
	request *WithdrawRequest,
) (*WithdrawResult, error) {
	callback := make(chan WithdrawAsyncResult, 1)
	go p.WithdrawAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func withdrawByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- WithdrawByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "wallet.operation.conflict" {
				asyncResult.Err = gs2err.SetClientError(Conflict{})
			}
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "wallet.balance.insufficient" {
				asyncResult.Err = gs2err.SetClientError(Insufficient{})
			}
		}
		callback <- WithdrawByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WithdrawByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) WithdrawByUserIdAsync(
	request *WithdrawByUserIdRequest,
	callback chan<- WithdrawByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/wallet/{slot}/withdraw"
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
	if request.Slot != nil {
		path = strings.ReplaceAll(path, "{slot}", core.ToString(*request.Slot))
	} else {
		path = strings.ReplaceAll(path, "{slot}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.WithdrawCount != nil {
		bodies["withdrawCount"] = *request.WithdrawCount
	}
	if request.PaidOnly != nil {
		bodies["paidOnly"] = *request.PaidOnly
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

	go withdrawByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) WithdrawByUserId(
	request *WithdrawByUserIdRequest,
) (*WithdrawByUserIdResult, error) {
	callback := make(chan WithdrawByUserIdAsyncResult, 1)
	go p.WithdrawByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func depositByStampSheetAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DepositByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DepositByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DepositByStampSheetResult
	if asyncResult.Err != nil {
		callback <- DepositByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DepositByStampSheetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DepositByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DepositByStampSheetAsync(
	request *DepositByStampSheetRequest,
	callback chan<- DepositByStampSheetAsyncResult,
) {
	path := "/stamp/deposit"

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

	go depositByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DepositByStampSheet(
	request *DepositByStampSheetRequest,
) (*DepositByStampSheetResult, error) {
	callback := make(chan DepositByStampSheetAsyncResult, 1)
	go p.DepositByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func withdrawByStampTaskAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- WithdrawByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WithdrawByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WithdrawByStampTaskResult
	if asyncResult.Err != nil {
		callback <- WithdrawByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WithdrawByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WithdrawByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) WithdrawByStampTaskAsync(
	request *WithdrawByStampTaskRequest,
	callback chan<- WithdrawByStampTaskAsyncResult,
) {
	path := "/stamp/withdraw"

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

	go withdrawByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) WithdrawByStampTask(
	request *WithdrawByStampTaskRequest,
) (*WithdrawByStampTaskResult, error) {
	callback := make(chan WithdrawByStampTaskAsyncResult, 1)
	go p.WithdrawByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEventsByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEventsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEventsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEventsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeEventsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEventsByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEventsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeEventsByUserIdAsync(
	request *DescribeEventsByUserIdRequest,
	callback chan<- DescribeEventsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/event/user/{userId}"
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
	if request.Begin != nil {
		queryStrings["begin"] = core.ToString(*request.Begin)
	}
	if request.End != nil {
		queryStrings["end"] = core.ToString(*request.End)
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

	go describeEventsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeEventsByUserId(
	request *DescribeEventsByUserIdRequest,
) (*DescribeEventsByUserIdResult, error) {
	callback := make(chan DescribeEventsByUserIdAsyncResult, 1)
	go p.DescribeEventsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEventByTransactionIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetEventByTransactionIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEventByTransactionIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEventByTransactionIdResult
	if asyncResult.Err != nil {
		callback <- GetEventByTransactionIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEventByTransactionIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEventByTransactionIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetEventByTransactionIdAsync(
	request *GetEventByTransactionIdRequest,
	callback chan<- GetEventByTransactionIdAsyncResult,
) {
	path := "/{namespaceName}/event/{transactionId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
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

	go getEventByTransactionIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetEventByTransactionId(
	request *GetEventByTransactionIdRequest,
) (*GetEventByTransactionIdResult, error) {
	callback := make(chan GetEventByTransactionIdAsyncResult, 1)
	go p.GetEventByTransactionIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReceiptAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReceiptAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReceiptAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReceiptResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "receipt.payload.invalid" {
				asyncResult.Err = gs2err.SetClientError(ReceiptInvalid{})
			}
		}
		callback <- VerifyReceiptAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReceiptAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReceiptAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) VerifyReceiptAsync(
	request *VerifyReceiptRequest,
	callback chan<- VerifyReceiptAsyncResult,
) {
	path := "/{namespaceName}/user/me/content/{contentName}/receipt/verify"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Receipt != nil {
		bodies["receipt"] = request.Receipt.ToDict()
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

	go verifyReceiptAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) VerifyReceipt(
	request *VerifyReceiptRequest,
) (*VerifyReceiptResult, error) {
	callback := make(chan VerifyReceiptAsyncResult, 1)
	go p.VerifyReceiptAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReceiptByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReceiptByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReceiptByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReceiptByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "receipt.payload.invalid" {
				asyncResult.Err = gs2err.SetClientError(ReceiptInvalid{})
			}
		}
		callback <- VerifyReceiptByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReceiptByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReceiptByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) VerifyReceiptByUserIdAsync(
	request *VerifyReceiptByUserIdRequest,
	callback chan<- VerifyReceiptByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/content/{contentName}/receipt/verify"
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
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Receipt != nil {
		bodies["receipt"] = request.Receipt.ToDict()
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

	go verifyReceiptByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) VerifyReceiptByUserId(
	request *VerifyReceiptByUserIdRequest,
) (*VerifyReceiptByUserIdResult, error) {
	callback := make(chan VerifyReceiptByUserIdAsyncResult, 1)
	go p.VerifyReceiptByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func verifyReceiptByStampTaskAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- VerifyReceiptByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyReceiptByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyReceiptByStampTaskResult
	if asyncResult.Err != nil {
		callback <- VerifyReceiptByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyReceiptByStampTaskAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyReceiptByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) VerifyReceiptByStampTaskAsync(
	request *VerifyReceiptByStampTaskRequest,
	callback chan<- VerifyReceiptByStampTaskAsyncResult,
) {
	path := "/stamp/receipt/verify"

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

	go verifyReceiptByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) VerifyReceiptByStampTask(
	request *VerifyReceiptByStampTaskRequest,
) (*VerifyReceiptByStampTaskResult, error) {
	callback := make(chan VerifyReceiptByStampTaskAsyncResult, 1)
	go p.VerifyReceiptByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscriptionStatusesAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscriptionStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscriptionStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscriptionStatusesResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscriptionStatusesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscriptionStatusesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscriptionStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeSubscriptionStatusesAsync(
	request *DescribeSubscriptionStatusesRequest,
	callback chan<- DescribeSubscriptionStatusesAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscription"
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
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeSubscriptionStatusesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeSubscriptionStatuses(
	request *DescribeSubscriptionStatusesRequest,
) (*DescribeSubscriptionStatusesResult, error) {
	callback := make(chan DescribeSubscriptionStatusesAsyncResult, 1)
	go p.DescribeSubscriptionStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSubscriptionStatusesByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSubscriptionStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscriptionStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscriptionStatusesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeSubscriptionStatusesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSubscriptionStatusesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSubscriptionStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeSubscriptionStatusesByUserIdAsync(
	request *DescribeSubscriptionStatusesByUserIdRequest,
	callback chan<- DescribeSubscriptionStatusesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscription"
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

	go describeSubscriptionStatusesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeSubscriptionStatusesByUserId(
	request *DescribeSubscriptionStatusesByUserIdRequest,
) (*DescribeSubscriptionStatusesByUserIdResult, error) {
	callback := make(chan DescribeSubscriptionStatusesByUserIdAsyncResult, 1)
	go p.DescribeSubscriptionStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscriptionStatusAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscriptionStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscriptionStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscriptionStatusResult
	if asyncResult.Err != nil {
		callback <- GetSubscriptionStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscriptionStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscriptionStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetSubscriptionStatusAsync(
	request *GetSubscriptionStatusRequest,
	callback chan<- GetSubscriptionStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/subscription/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getSubscriptionStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetSubscriptionStatus(
	request *GetSubscriptionStatusRequest,
) (*GetSubscriptionStatusResult, error) {
	callback := make(chan GetSubscriptionStatusAsyncResult, 1)
	go p.GetSubscriptionStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSubscriptionStatusByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetSubscriptionStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscriptionStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscriptionStatusByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetSubscriptionStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSubscriptionStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSubscriptionStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetSubscriptionStatusByUserIdAsync(
	request *GetSubscriptionStatusByUserIdRequest,
	callback chan<- GetSubscriptionStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/subscription/{contentName}"
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
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getSubscriptionStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetSubscriptionStatusByUserId(
	request *GetSubscriptionStatusByUserIdRequest,
) (*GetSubscriptionStatusByUserIdResult, error) {
	callback := make(chan GetSubscriptionStatusByUserIdAsyncResult, 1)
	go p.GetSubscriptionStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func allocateSubscriptionStatusAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- AllocateSubscriptionStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AllocateSubscriptionStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AllocateSubscriptionStatusResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "subscription.transaction.used" {
				asyncResult.Err = gs2err.SetClientError(AlreadyUsed{})
			}
		}
		callback <- AllocateSubscriptionStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AllocateSubscriptionStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AllocateSubscriptionStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) AllocateSubscriptionStatusAsync(
	request *AllocateSubscriptionStatusRequest,
	callback chan<- AllocateSubscriptionStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/allocate/subscription"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Receipt != nil && *request.Receipt != "" {
		bodies["receipt"] = *request.Receipt
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

	go allocateSubscriptionStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) AllocateSubscriptionStatus(
	request *AllocateSubscriptionStatusRequest,
) (*AllocateSubscriptionStatusResult, error) {
	callback := make(chan AllocateSubscriptionStatusAsyncResult, 1)
	go p.AllocateSubscriptionStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func allocateSubscriptionStatusByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- AllocateSubscriptionStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AllocateSubscriptionStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AllocateSubscriptionStatusByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "subscription.transaction.used" {
				asyncResult.Err = gs2err.SetClientError(AlreadyUsed{})
			}
		}
		callback <- AllocateSubscriptionStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AllocateSubscriptionStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AllocateSubscriptionStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) AllocateSubscriptionStatusByUserIdAsync(
	request *AllocateSubscriptionStatusByUserIdRequest,
	callback chan<- AllocateSubscriptionStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/allocate/subscription"
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
	if request.Receipt != nil && *request.Receipt != "" {
		bodies["receipt"] = *request.Receipt
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

	go allocateSubscriptionStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) AllocateSubscriptionStatusByUserId(
	request *AllocateSubscriptionStatusByUserIdRequest,
) (*AllocateSubscriptionStatusByUserIdResult, error) {
	callback := make(chan AllocateSubscriptionStatusByUserIdAsyncResult, 1)
	go p.AllocateSubscriptionStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func takeoverSubscriptionStatusAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- TakeoverSubscriptionStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- TakeoverSubscriptionStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result TakeoverSubscriptionStatusResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "subscription.transaction.used" {
				asyncResult.Err = gs2err.SetClientError(LockPeriodNotElapsed{})
			}
		}
		callback <- TakeoverSubscriptionStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- TakeoverSubscriptionStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- TakeoverSubscriptionStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) TakeoverSubscriptionStatusAsync(
	request *TakeoverSubscriptionStatusRequest,
	callback chan<- TakeoverSubscriptionStatusAsyncResult,
) {
	path := "/{namespaceName}/user/me/takeover/subscription"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Receipt != nil && *request.Receipt != "" {
		bodies["receipt"] = *request.Receipt
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

	go takeoverSubscriptionStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) TakeoverSubscriptionStatus(
	request *TakeoverSubscriptionStatusRequest,
) (*TakeoverSubscriptionStatusResult, error) {
	callback := make(chan TakeoverSubscriptionStatusAsyncResult, 1)
	go p.TakeoverSubscriptionStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func takeoverSubscriptionStatusByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- TakeoverSubscriptionStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- TakeoverSubscriptionStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result TakeoverSubscriptionStatusByUserIdResult
	if asyncResult.Err != nil {
		gs2err, ok := asyncResult.Err.(core.Gs2Exception)
		if ok {
			if len(gs2err.RequestErrors()) > 0 && gs2err.RequestErrors()[0].Code != nil && *gs2err.RequestErrors()[0].Code == "subscription.transaction.used" {
				asyncResult.Err = gs2err.SetClientError(LockPeriodNotElapsed{})
			}
		}
		callback <- TakeoverSubscriptionStatusByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- TakeoverSubscriptionStatusByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- TakeoverSubscriptionStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) TakeoverSubscriptionStatusByUserIdAsync(
	request *TakeoverSubscriptionStatusByUserIdRequest,
	callback chan<- TakeoverSubscriptionStatusByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/takeover/subscription"
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
	if request.Receipt != nil && *request.Receipt != "" {
		bodies["receipt"] = *request.Receipt
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

	go takeoverSubscriptionStatusByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) TakeoverSubscriptionStatusByUserId(
	request *TakeoverSubscriptionStatusByUserIdRequest,
) (*TakeoverSubscriptionStatusByUserIdResult, error) {
	callback := make(chan TakeoverSubscriptionStatusByUserIdAsyncResult, 1)
	go p.TakeoverSubscriptionStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRefundHistoriesByUserIdAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRefundHistoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRefundHistoriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRefundHistoriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeRefundHistoriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRefundHistoriesByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRefundHistoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeRefundHistoriesByUserIdAsync(
	request *DescribeRefundHistoriesByUserIdRequest,
	callback chan<- DescribeRefundHistoriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/refund/user/{userId}"
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

	go describeRefundHistoriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeRefundHistoriesByUserId(
	request *DescribeRefundHistoriesByUserIdRequest,
) (*DescribeRefundHistoriesByUserIdResult, error) {
	callback := make(chan DescribeRefundHistoriesByUserIdAsyncResult, 1)
	go p.DescribeRefundHistoriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRefundHistoriesByDateAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRefundHistoriesByDateAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRefundHistoriesByDateAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRefundHistoriesByDateResult
	if asyncResult.Err != nil {
		callback <- DescribeRefundHistoriesByDateAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRefundHistoriesByDateAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRefundHistoriesByDateAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeRefundHistoriesByDateAsync(
	request *DescribeRefundHistoriesByDateRequest,
	callback chan<- DescribeRefundHistoriesByDateAsyncResult,
) {
	path := "/{namespaceName}/refund/date/{year}/{month}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}
	if request.Month != nil {
		path = strings.ReplaceAll(path, "{month}", core.ToString(*request.Month))
	} else {
		path = strings.ReplaceAll(path, "{month}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Day != nil {
		queryStrings["day"] = core.ToString(*request.Day)
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

	go describeRefundHistoriesByDateAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeRefundHistoriesByDate(
	request *DescribeRefundHistoriesByDateRequest,
) (*DescribeRefundHistoriesByDateResult, error) {
	callback := make(chan DescribeRefundHistoriesByDateAsyncResult, 1)
	go p.DescribeRefundHistoriesByDateAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRefundHistoryAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetRefundHistoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRefundHistoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRefundHistoryResult
	if asyncResult.Err != nil {
		callback <- GetRefundHistoryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRefundHistoryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRefundHistoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetRefundHistoryAsync(
	request *GetRefundHistoryRequest,
	callback chan<- GetRefundHistoryAsyncResult,
) {
	path := "/{namespaceName}/refund/{transactionId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
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

	go getRefundHistoryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetRefundHistory(
	request *GetRefundHistoryRequest,
) (*GetRefundHistoryResult, error) {
	callback := make(chan GetRefundHistoryAsyncResult, 1)
	go p.GetRefundHistoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStoreContentModelsAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStoreContentModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStoreContentModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStoreContentModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeStoreContentModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStoreContentModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStoreContentModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeStoreContentModelsAsync(
	request *DescribeStoreContentModelsRequest,
	callback chan<- DescribeStoreContentModelsAsyncResult,
) {
	path := "/{namespaceName}/model/content"
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

	go describeStoreContentModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeStoreContentModels(
	request *DescribeStoreContentModelsRequest,
) (*DescribeStoreContentModelsResult, error) {
	callback := make(chan DescribeStoreContentModelsAsyncResult, 1)
	go p.DescribeStoreContentModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStoreContentModelAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetStoreContentModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStoreContentModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStoreContentModelResult
	if asyncResult.Err != nil {
		callback <- GetStoreContentModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStoreContentModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStoreContentModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetStoreContentModelAsync(
	request *GetStoreContentModelRequest,
	callback chan<- GetStoreContentModelAsyncResult,
) {
	path := "/{namespaceName}/model/content/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getStoreContentModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetStoreContentModel(
	request *GetStoreContentModelRequest,
) (*GetStoreContentModelResult, error) {
	callback := make(chan GetStoreContentModelAsyncResult, 1)
	go p.GetStoreContentModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStoreContentModelMastersAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStoreContentModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStoreContentModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStoreContentModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeStoreContentModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStoreContentModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStoreContentModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeStoreContentModelMastersAsync(
	request *DescribeStoreContentModelMastersRequest,
	callback chan<- DescribeStoreContentModelMastersAsyncResult,
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

	go describeStoreContentModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeStoreContentModelMasters(
	request *DescribeStoreContentModelMastersRequest,
) (*DescribeStoreContentModelMastersResult, error) {
	callback := make(chan DescribeStoreContentModelMastersAsyncResult, 1)
	go p.DescribeStoreContentModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStoreContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateStoreContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStoreContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStoreContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateStoreContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStoreContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStoreContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) CreateStoreContentModelMasterAsync(
	request *CreateStoreContentModelMasterRequest,
	callback chan<- CreateStoreContentModelMasterAsyncResult,
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
	if request.AppleAppStore != nil {
		bodies["appleAppStore"] = request.AppleAppStore.ToDict()
	}
	if request.GooglePlay != nil {
		bodies["googlePlay"] = request.GooglePlay.ToDict()
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

	go createStoreContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CreateStoreContentModelMaster(
	request *CreateStoreContentModelMasterRequest,
) (*CreateStoreContentModelMasterResult, error) {
	callback := make(chan CreateStoreContentModelMasterAsyncResult, 1)
	go p.CreateStoreContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStoreContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetStoreContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStoreContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStoreContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetStoreContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStoreContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStoreContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetStoreContentModelMasterAsync(
	request *GetStoreContentModelMasterRequest,
	callback chan<- GetStoreContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getStoreContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetStoreContentModelMaster(
	request *GetStoreContentModelMasterRequest,
) (*GetStoreContentModelMasterResult, error) {
	callback := make(chan GetStoreContentModelMasterAsyncResult, 1)
	go p.GetStoreContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStoreContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStoreContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStoreContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStoreContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateStoreContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStoreContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStoreContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) UpdateStoreContentModelMasterAsync(
	request *UpdateStoreContentModelMasterRequest,
	callback chan<- UpdateStoreContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.AppleAppStore != nil {
		bodies["appleAppStore"] = request.AppleAppStore.ToDict()
	}
	if request.GooglePlay != nil {
		bodies["googlePlay"] = request.GooglePlay.ToDict()
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

	go updateStoreContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) UpdateStoreContentModelMaster(
	request *UpdateStoreContentModelMasterRequest,
) (*UpdateStoreContentModelMasterResult, error) {
	callback := make(chan UpdateStoreContentModelMasterAsyncResult, 1)
	go p.UpdateStoreContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStoreContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStoreContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStoreContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStoreContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteStoreContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStoreContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStoreContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DeleteStoreContentModelMasterAsync(
	request *DeleteStoreContentModelMasterRequest,
	callback chan<- DeleteStoreContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go deleteStoreContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DeleteStoreContentModelMaster(
	request *DeleteStoreContentModelMasterRequest,
) (*DeleteStoreContentModelMasterResult, error) {
	callback := make(chan DeleteStoreContentModelMasterAsyncResult, 1)
	go p.DeleteStoreContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStoreSubscriptionContentModelsAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStoreSubscriptionContentModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStoreSubscriptionContentModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStoreSubscriptionContentModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeStoreSubscriptionContentModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStoreSubscriptionContentModelsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStoreSubscriptionContentModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeStoreSubscriptionContentModelsAsync(
	request *DescribeStoreSubscriptionContentModelsRequest,
	callback chan<- DescribeStoreSubscriptionContentModelsAsyncResult,
) {
	path := "/{namespaceName}/model/subscription/content"
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

	go describeStoreSubscriptionContentModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeStoreSubscriptionContentModels(
	request *DescribeStoreSubscriptionContentModelsRequest,
) (*DescribeStoreSubscriptionContentModelsResult, error) {
	callback := make(chan DescribeStoreSubscriptionContentModelsAsyncResult, 1)
	go p.DescribeStoreSubscriptionContentModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStoreSubscriptionContentModelAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetStoreSubscriptionContentModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStoreSubscriptionContentModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStoreSubscriptionContentModelResult
	if asyncResult.Err != nil {
		callback <- GetStoreSubscriptionContentModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStoreSubscriptionContentModelAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStoreSubscriptionContentModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetStoreSubscriptionContentModelAsync(
	request *GetStoreSubscriptionContentModelRequest,
	callback chan<- GetStoreSubscriptionContentModelAsyncResult,
) {
	path := "/{namespaceName}/model/subscription/content/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getStoreSubscriptionContentModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetStoreSubscriptionContentModel(
	request *GetStoreSubscriptionContentModelRequest,
) (*GetStoreSubscriptionContentModelResult, error) {
	callback := make(chan GetStoreSubscriptionContentModelAsyncResult, 1)
	go p.GetStoreSubscriptionContentModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeStoreSubscriptionContentModelMastersAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStoreSubscriptionContentModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStoreSubscriptionContentModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStoreSubscriptionContentModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeStoreSubscriptionContentModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStoreSubscriptionContentModelMastersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStoreSubscriptionContentModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeStoreSubscriptionContentModelMastersAsync(
	request *DescribeStoreSubscriptionContentModelMastersRequest,
	callback chan<- DescribeStoreSubscriptionContentModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/model/subscription"
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

	go describeStoreSubscriptionContentModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeStoreSubscriptionContentModelMasters(
	request *DescribeStoreSubscriptionContentModelMastersRequest,
) (*DescribeStoreSubscriptionContentModelMastersResult, error) {
	callback := make(chan DescribeStoreSubscriptionContentModelMastersAsyncResult, 1)
	go p.DescribeStoreSubscriptionContentModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStoreSubscriptionContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- CreateStoreSubscriptionContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStoreSubscriptionContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStoreSubscriptionContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateStoreSubscriptionContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStoreSubscriptionContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStoreSubscriptionContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) CreateStoreSubscriptionContentModelMasterAsync(
	request *CreateStoreSubscriptionContentModelMasterRequest,
	callback chan<- CreateStoreSubscriptionContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/subscription"
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
	if request.ScheduleNamespaceId != nil && *request.ScheduleNamespaceId != "" {
		bodies["scheduleNamespaceId"] = *request.ScheduleNamespaceId
	}
	if request.TriggerName != nil && *request.TriggerName != "" {
		bodies["triggerName"] = *request.TriggerName
	}
	if request.TriggerExtendMode != nil && *request.TriggerExtendMode != "" {
		bodies["triggerExtendMode"] = *request.TriggerExtendMode
	}
	if request.RollupHour != nil {
		bodies["rollupHour"] = *request.RollupHour
	}
	if request.ReallocateSpanDays != nil {
		bodies["reallocateSpanDays"] = *request.ReallocateSpanDays
	}
	if request.AppleAppStore != nil {
		bodies["appleAppStore"] = request.AppleAppStore.ToDict()
	}
	if request.GooglePlay != nil {
		bodies["googlePlay"] = request.GooglePlay.ToDict()
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

	go createStoreSubscriptionContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) CreateStoreSubscriptionContentModelMaster(
	request *CreateStoreSubscriptionContentModelMasterRequest,
) (*CreateStoreSubscriptionContentModelMasterResult, error) {
	callback := make(chan CreateStoreSubscriptionContentModelMasterAsyncResult, 1)
	go p.CreateStoreSubscriptionContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStoreSubscriptionContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetStoreSubscriptionContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStoreSubscriptionContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStoreSubscriptionContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetStoreSubscriptionContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStoreSubscriptionContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStoreSubscriptionContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetStoreSubscriptionContentModelMasterAsync(
	request *GetStoreSubscriptionContentModelMasterRequest,
	callback chan<- GetStoreSubscriptionContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/subscription/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go getStoreSubscriptionContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetStoreSubscriptionContentModelMaster(
	request *GetStoreSubscriptionContentModelMasterRequest,
) (*GetStoreSubscriptionContentModelMasterResult, error) {
	callback := make(chan GetStoreSubscriptionContentModelMasterAsyncResult, 1)
	go p.GetStoreSubscriptionContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStoreSubscriptionContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStoreSubscriptionContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStoreSubscriptionContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStoreSubscriptionContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateStoreSubscriptionContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStoreSubscriptionContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStoreSubscriptionContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) UpdateStoreSubscriptionContentModelMasterAsync(
	request *UpdateStoreSubscriptionContentModelMasterRequest,
	callback chan<- UpdateStoreSubscriptionContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/subscription/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Metadata != nil && *request.Metadata != "" {
		bodies["metadata"] = *request.Metadata
	}
	if request.ScheduleNamespaceId != nil && *request.ScheduleNamespaceId != "" {
		bodies["scheduleNamespaceId"] = *request.ScheduleNamespaceId
	}
	if request.TriggerName != nil && *request.TriggerName != "" {
		bodies["triggerName"] = *request.TriggerName
	}
	if request.TriggerExtendMode != nil && *request.TriggerExtendMode != "" {
		bodies["triggerExtendMode"] = *request.TriggerExtendMode
	}
	if request.RollupHour != nil {
		bodies["rollupHour"] = *request.RollupHour
	}
	if request.ReallocateSpanDays != nil {
		bodies["reallocateSpanDays"] = *request.ReallocateSpanDays
	}
	if request.AppleAppStore != nil {
		bodies["appleAppStore"] = request.AppleAppStore.ToDict()
	}
	if request.GooglePlay != nil {
		bodies["googlePlay"] = request.GooglePlay.ToDict()
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

	go updateStoreSubscriptionContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) UpdateStoreSubscriptionContentModelMaster(
	request *UpdateStoreSubscriptionContentModelMasterRequest,
) (*UpdateStoreSubscriptionContentModelMasterResult, error) {
	callback := make(chan UpdateStoreSubscriptionContentModelMasterAsyncResult, 1)
	go p.UpdateStoreSubscriptionContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStoreSubscriptionContentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStoreSubscriptionContentModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStoreSubscriptionContentModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStoreSubscriptionContentModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteStoreSubscriptionContentModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStoreSubscriptionContentModelMasterAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStoreSubscriptionContentModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DeleteStoreSubscriptionContentModelMasterAsync(
	request *DeleteStoreSubscriptionContentModelMasterRequest,
	callback chan<- DeleteStoreSubscriptionContentModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/model/subscription/{contentName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ContentName != nil && *request.ContentName != "" {
		path = strings.ReplaceAll(path, "{contentName}", core.ToString(*request.ContentName))
	} else {
		path = strings.ReplaceAll(path, "{contentName}", "null")
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

	go deleteStoreSubscriptionContentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DeleteStoreSubscriptionContentModelMaster(
	request *DeleteStoreSubscriptionContentModelMasterRequest,
) (*DeleteStoreSubscriptionContentModelMasterResult, error) {
	callback := make(chan DeleteStoreSubscriptionContentModelMasterAsyncResult, 1)
	go p.DeleteStoreSubscriptionContentModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) ExportMaster(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) GetCurrentModelMasterAsync(
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
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetCurrentModelMaster(
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

func updateCurrentModelMasterAsyncHandler(
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) UpdateCurrentModelMasterAsync(
	request *UpdateCurrentModelMasterRequest,
	callback chan<- UpdateCurrentModelMasterAsyncResult,
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

	go updateCurrentModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) UpdateCurrentModelMaster(
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
	client Gs2Money2RestClient,
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

func (p Gs2Money2RestClient) UpdateCurrentModelMasterFromGitHubAsync(
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
			Url:     p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) UpdateCurrentModelMasterFromGitHub(
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

func describeDailyTransactionHistoriesByCurrencyAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDailyTransactionHistoriesByCurrencyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDailyTransactionHistoriesByCurrencyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDailyTransactionHistoriesByCurrencyResult
	if asyncResult.Err != nil {
		callback <- DescribeDailyTransactionHistoriesByCurrencyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDailyTransactionHistoriesByCurrencyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDailyTransactionHistoriesByCurrencyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeDailyTransactionHistoriesByCurrencyAsync(
	request *DescribeDailyTransactionHistoriesByCurrencyRequest,
	callback chan<- DescribeDailyTransactionHistoriesByCurrencyAsyncResult,
) {
	path := "/{namespaceName}/transaction/daily/currency/{currency}/date/{year}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Currency != nil && *request.Currency != "" {
		path = strings.ReplaceAll(path, "{currency}", core.ToString(*request.Currency))
	} else {
		path = strings.ReplaceAll(path, "{currency}", "null")
	}
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Month != nil {
		queryStrings["month"] = core.ToString(*request.Month)
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

	go describeDailyTransactionHistoriesByCurrencyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeDailyTransactionHistoriesByCurrency(
	request *DescribeDailyTransactionHistoriesByCurrencyRequest,
) (*DescribeDailyTransactionHistoriesByCurrencyResult, error) {
	callback := make(chan DescribeDailyTransactionHistoriesByCurrencyAsyncResult, 1)
	go p.DescribeDailyTransactionHistoriesByCurrencyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDailyTransactionHistoriesAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDailyTransactionHistoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDailyTransactionHistoriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDailyTransactionHistoriesResult
	if asyncResult.Err != nil {
		callback <- DescribeDailyTransactionHistoriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDailyTransactionHistoriesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDailyTransactionHistoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeDailyTransactionHistoriesAsync(
	request *DescribeDailyTransactionHistoriesRequest,
	callback chan<- DescribeDailyTransactionHistoriesAsyncResult,
) {
	path := "/{namespaceName}/transaction/daily/{year}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Month != nil {
		queryStrings["month"] = core.ToString(*request.Month)
	}
	if request.Day != nil {
		queryStrings["day"] = core.ToString(*request.Day)
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

	go describeDailyTransactionHistoriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeDailyTransactionHistories(
	request *DescribeDailyTransactionHistoriesRequest,
) (*DescribeDailyTransactionHistoriesResult, error) {
	callback := make(chan DescribeDailyTransactionHistoriesAsyncResult, 1)
	go p.DescribeDailyTransactionHistoriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDailyTransactionHistoryAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetDailyTransactionHistoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDailyTransactionHistoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDailyTransactionHistoryResult
	if asyncResult.Err != nil {
		callback <- GetDailyTransactionHistoryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDailyTransactionHistoryAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDailyTransactionHistoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetDailyTransactionHistoryAsync(
	request *GetDailyTransactionHistoryRequest,
	callback chan<- GetDailyTransactionHistoryAsyncResult,
) {
	path := "/{namespaceName}/transaction/daily/{year}/{month}/{day}/currency/{currency}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}
	if request.Month != nil {
		path = strings.ReplaceAll(path, "{month}", core.ToString(*request.Month))
	} else {
		path = strings.ReplaceAll(path, "{month}", "null")
	}
	if request.Day != nil {
		path = strings.ReplaceAll(path, "{day}", core.ToString(*request.Day))
	} else {
		path = strings.ReplaceAll(path, "{day}", "null")
	}
	if request.Currency != nil && *request.Currency != "" {
		path = strings.ReplaceAll(path, "{currency}", core.ToString(*request.Currency))
	} else {
		path = strings.ReplaceAll(path, "{currency}", "null")
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

	go getDailyTransactionHistoryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetDailyTransactionHistory(
	request *GetDailyTransactionHistoryRequest,
) (*GetDailyTransactionHistoryResult, error) {
	callback := make(chan GetDailyTransactionHistoryAsyncResult, 1)
	go p.GetDailyTransactionHistoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeUnusedBalancesAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- DescribeUnusedBalancesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeUnusedBalancesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeUnusedBalancesResult
	if asyncResult.Err != nil {
		callback <- DescribeUnusedBalancesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeUnusedBalancesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeUnusedBalancesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) DescribeUnusedBalancesAsync(
	request *DescribeUnusedBalancesRequest,
	callback chan<- DescribeUnusedBalancesAsyncResult,
) {
	path := "/{namespaceName}/balance/unused"
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

	go describeUnusedBalancesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) DescribeUnusedBalances(
	request *DescribeUnusedBalancesRequest,
) (*DescribeUnusedBalancesResult, error) {
	callback := make(chan DescribeUnusedBalancesAsyncResult, 1)
	go p.DescribeUnusedBalancesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getUnusedBalanceAsyncHandler(
	client Gs2Money2RestClient,
	job *core.NetworkJob,
	callback chan<- GetUnusedBalanceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetUnusedBalanceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetUnusedBalanceResult
	if asyncResult.Err != nil {
		callback <- GetUnusedBalanceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetUnusedBalanceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetUnusedBalanceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2Money2RestClient) GetUnusedBalanceAsync(
	request *GetUnusedBalanceRequest,
	callback chan<- GetUnusedBalanceAsyncResult,
) {
	path := "/{namespaceName}/balance/unused/{currency}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.Currency != nil && *request.Currency != "" {
		path = strings.ReplaceAll(path, "{currency}", core.ToString(*request.Currency))
	} else {
		path = strings.ReplaceAll(path, "{currency}", "null")
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

	go getUnusedBalanceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("money2").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2Money2RestClient) GetUnusedBalance(
	request *GetUnusedBalanceRequest,
) (*GetUnusedBalanceResult, error) {
	callback := make(chan GetUnusedBalanceAsyncResult, 1)
	go p.GetUnusedBalanceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
