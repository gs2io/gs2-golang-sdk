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
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.LogSetting != nil {
		bodies["logSetting"] = request.LogSetting.ToDict()
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	path := "/{namespaceName}/account/{userId}/ban/{banName}"
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.UserIdentifier != nil {
		queryStrings["userIdentifier"] = core.ToString(*request.UserIdentifier)
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.AccessToken != nil {
		headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}
	if request.DuplicationAvoider != nil {
		headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
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
