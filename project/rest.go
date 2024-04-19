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

package project

import (
	"encoding/json"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ProjectRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ProjectRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func createAccountAsyncHandler(
	client Gs2ProjectRestClient,
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

func (p Gs2ProjectRestClient) CreateAccountAsync(
	request *CreateAccountRequest,
	callback chan<- CreateAccountAsyncResult,
) {
	path := "/account"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Email != nil && *request.Email != "" {
		bodies["email"] = *request.Email
	}
	if request.FullName != nil && *request.FullName != "" {
		bodies["fullName"] = *request.FullName
	}
	if request.CompanyName != nil && *request.CompanyName != "" {
		bodies["companyName"] = *request.CompanyName
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.Lang != nil && *request.Lang != "" {
		bodies["lang"] = *request.Lang
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

	go createAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) CreateAccount(
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

func verifyAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- VerifyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VerifyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VerifyResult
	if asyncResult.Err != nil {
		callback <- VerifyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- VerifyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- VerifyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) VerifyAsync(
	request *VerifyRequest,
	callback chan<- VerifyAsyncResult,
) {
	path := "/account/verify"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.VerifyToken != nil && *request.VerifyToken != "" {
		bodies["verifyToken"] = *request.VerifyToken
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

	go verifyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) Verify(
	request *VerifyRequest,
) (*VerifyResult, error) {
	callback := make(chan VerifyAsyncResult, 1)
	go p.VerifyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func signInAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- SignInAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SignInAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SignInResult
	if asyncResult.Err != nil {
		callback <- SignInAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- SignInAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- SignInAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) SignInAsync(
	request *SignInRequest,
	callback chan<- SignInAsyncResult,
) {
	path := "/account/signIn"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Email != nil && *request.Email != "" {
		bodies["email"] = *request.Email
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go signInAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) SignIn(
	request *SignInRequest,
) (*SignInResult, error) {
	callback := make(chan SignInAsyncResult, 1)
	go p.SignInAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func forgetAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- ForgetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ForgetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ForgetResult
	if asyncResult.Err != nil {
		callback <- ForgetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ForgetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ForgetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) ForgetAsync(
	request *ForgetRequest,
	callback chan<- ForgetAsyncResult,
) {
	path := "/account/forget"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Email != nil && *request.Email != "" {
		bodies["email"] = *request.Email
	}
	if request.Lang != nil && *request.Lang != "" {
		bodies["lang"] = *request.Lang
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

	go forgetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) Forget(
	request *ForgetRequest,
) (*ForgetResult, error) {
	callback := make(chan ForgetAsyncResult, 1)
	go p.ForgetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func issuePasswordAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- IssuePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IssuePasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IssuePasswordResult
	if asyncResult.Err != nil {
		callback <- IssuePasswordAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IssuePasswordAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IssuePasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) IssuePasswordAsync(
	request *IssuePasswordRequest,
	callback chan<- IssuePasswordAsyncResult,
) {
	path := "/account/password/issue"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.IssuePasswordToken != nil && *request.IssuePasswordToken != "" {
		bodies["issuePasswordToken"] = *request.IssuePasswordToken
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

	go issuePasswordAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) IssuePassword(
	request *IssuePasswordRequest,
) (*IssuePasswordResult, error) {
	callback := make(chan IssuePasswordAsyncResult, 1)
	go p.IssuePasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateAccountAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateAccountAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateAccountResult
	if asyncResult.Err != nil {
		callback <- UpdateAccountAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateAccountAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateAccountAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) UpdateAccountAsync(
	request *UpdateAccountRequest,
	callback chan<- UpdateAccountAsyncResult,
) {
	path := "/account"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Email != nil && *request.Email != "" {
		bodies["email"] = *request.Email
	}
	if request.FullName != nil && *request.FullName != "" {
		bodies["fullName"] = *request.FullName
	}
	if request.CompanyName != nil && *request.CompanyName != "" {
		bodies["companyName"] = *request.CompanyName
	}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
	}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
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

	go updateAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) UpdateAccount(
	request *UpdateAccountRequest,
) (*UpdateAccountResult, error) {
	callback := make(chan UpdateAccountAsyncResult, 1)
	go p.UpdateAccountAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteAccountAsyncHandler(
	client Gs2ProjectRestClient,
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

func (p Gs2ProjectRestClient) DeleteAccountAsync(
	request *DeleteAccountRequest,
	callback chan<- DeleteAccountAsyncResult,
) {
	path := "/account"

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

	go deleteAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DeleteAccount(
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

func describeProjectsAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeProjectsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeProjectsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeProjectsResult
	if asyncResult.Err != nil {
		callback <- DescribeProjectsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeProjectsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeProjectsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeProjectsAsync(
	request *DescribeProjectsRequest,
	callback chan<- DescribeProjectsAsyncResult,
) {
	path := "/account/me/project"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeProjectsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeProjects(
	request *DescribeProjectsRequest,
) (*DescribeProjectsResult, error) {
	callback := make(chan DescribeProjectsAsyncResult, 1)
	go p.DescribeProjectsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createProjectAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- CreateProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateProjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateProjectResult
	if asyncResult.Err != nil {
		callback <- CreateProjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateProjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateProjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) CreateProjectAsync(
	request *CreateProjectRequest,
	callback chan<- CreateProjectAsyncResult,
) {
	path := "/account/me/project"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Plan != nil && *request.Plan != "" {
		bodies["plan"] = *request.Plan
	}
	if request.Currency != nil && *request.Currency != "" {
		bodies["currency"] = *request.Currency
	}
	if request.ActivateRegionName != nil && *request.ActivateRegionName != "" {
		bodies["activateRegionName"] = *request.ActivateRegionName
	}
	if request.BillingMethodName != nil && *request.BillingMethodName != "" {
		bodies["billingMethodName"] = *request.BillingMethodName
	}
	if request.EnableEventBridge != nil && *request.EnableEventBridge != "" {
		bodies["enableEventBridge"] = *request.EnableEventBridge
	}
	if request.EventBridgeAwsAccountId != nil && *request.EventBridgeAwsAccountId != "" {
		bodies["eventBridgeAwsAccountId"] = *request.EventBridgeAwsAccountId
	}
	if request.EventBridgeAwsRegion != nil && *request.EventBridgeAwsRegion != "" {
		bodies["eventBridgeAwsRegion"] = *request.EventBridgeAwsRegion
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

	go createProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) CreateProject(
	request *CreateProjectRequest,
) (*CreateProjectResult, error) {
	callback := make(chan CreateProjectAsyncResult, 1)
	go p.CreateProjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getProjectAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProjectResult
	if asyncResult.Err != nil {
		callback <- GetProjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetProjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetProjectAsync(
	request *GetProjectRequest,
	callback chan<- GetProjectAsyncResult,
) {
	path := "/account/me/project/{projectName}"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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

	go getProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetProject(
	request *GetProjectRequest,
) (*GetProjectResult, error) {
	callback := make(chan GetProjectAsyncResult, 1)
	go p.GetProjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getProjectTokenAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetProjectTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProjectTokenAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProjectTokenResult
	if asyncResult.Err != nil {
		callback <- GetProjectTokenAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProjectTokenAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetProjectTokenAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetProjectTokenAsync(
	request *GetProjectTokenRequest,
	callback chan<- GetProjectTokenAsyncResult,
) {
	path := "/project/{projectName}/projectToken"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
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

	go getProjectTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetProjectToken(
	request *GetProjectTokenRequest,
) (*GetProjectTokenResult, error) {
	callback := make(chan GetProjectTokenAsyncResult, 1)
	go p.GetProjectTokenAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getProjectTokenByIdentifierAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetProjectTokenByIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetProjectTokenByIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetProjectTokenByIdentifierResult
	if asyncResult.Err != nil {
		callback <- GetProjectTokenByIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetProjectTokenByIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetProjectTokenByIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetProjectTokenByIdentifierAsync(
	request *GetProjectTokenByIdentifierRequest,
	callback chan<- GetProjectTokenByIdentifierAsyncResult,
) {
	path := "/account/{accountName}/project/{projectName}/user/{userName}/projectToken"
	if request.AccountName != nil && *request.AccountName != "" {
		path = strings.ReplaceAll(path, "{accountName}", core.ToString(*request.AccountName))
	} else {
		path = strings.ReplaceAll(path, "{accountName}", "null")
	}
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Password != nil && *request.Password != "" {
		bodies["password"] = *request.Password
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

	go getProjectTokenByIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetProjectTokenByIdentifier(
	request *GetProjectTokenByIdentifierRequest,
) (*GetProjectTokenByIdentifierResult, error) {
	callback := make(chan GetProjectTokenByIdentifierAsyncResult, 1)
	go p.GetProjectTokenByIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateProjectAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateProjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateProjectResult
	if asyncResult.Err != nil {
		callback <- UpdateProjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateProjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateProjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) UpdateProjectAsync(
	request *UpdateProjectRequest,
	callback chan<- UpdateProjectAsyncResult,
) {
	path := "/account/me/project/{projectName}"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Plan != nil && *request.Plan != "" {
		bodies["plan"] = *request.Plan
	}
	if request.BillingMethodName != nil && *request.BillingMethodName != "" {
		bodies["billingMethodName"] = *request.BillingMethodName
	}
	if request.EnableEventBridge != nil && *request.EnableEventBridge != "" {
		bodies["enableEventBridge"] = *request.EnableEventBridge
	}
	if request.EventBridgeAwsAccountId != nil && *request.EventBridgeAwsAccountId != "" {
		bodies["eventBridgeAwsAccountId"] = *request.EventBridgeAwsAccountId
	}
	if request.EventBridgeAwsRegion != nil && *request.EventBridgeAwsRegion != "" {
		bodies["eventBridgeAwsRegion"] = *request.EventBridgeAwsRegion
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

	go updateProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) UpdateProject(
	request *UpdateProjectRequest,
) (*UpdateProjectResult, error) {
	callback := make(chan UpdateProjectAsyncResult, 1)
	go p.UpdateProjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func activateRegionAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- ActivateRegionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ActivateRegionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ActivateRegionResult
	if asyncResult.Err != nil {
		callback <- ActivateRegionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ActivateRegionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ActivateRegionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) ActivateRegionAsync(
	request *ActivateRegionRequest,
	callback chan<- ActivateRegionAsyncResult,
) {
	path := "/account/me/project/{projectName}/region/{regionName}/activate"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}
	if request.RegionName != nil && *request.RegionName != "" {
		path = strings.ReplaceAll(path, "{regionName}", core.ToString(*request.RegionName))
	} else {
		path = strings.ReplaceAll(path, "{regionName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
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

	go activateRegionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) ActivateRegion(
	request *ActivateRegionRequest,
) (*ActivateRegionResult, error) {
	callback := make(chan ActivateRegionAsyncResult, 1)
	go p.ActivateRegionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func waitActivateRegionAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- WaitActivateRegionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WaitActivateRegionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WaitActivateRegionResult
	if asyncResult.Err != nil {
		callback <- WaitActivateRegionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitActivateRegionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WaitActivateRegionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) WaitActivateRegionAsync(
	request *WaitActivateRegionRequest,
	callback chan<- WaitActivateRegionAsyncResult,
) {
	path := "/account/me/project/{projectName}/region/{regionName}/activate/wait"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}
	if request.RegionName != nil && *request.RegionName != "" {
		path = strings.ReplaceAll(path, "{regionName}", core.ToString(*request.RegionName))
	} else {
		path = strings.ReplaceAll(path, "{regionName}", "null")
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

	go waitActivateRegionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) WaitActivateRegion(
	request *WaitActivateRegionRequest,
) (*WaitActivateRegionResult, error) {
	callback := make(chan WaitActivateRegionAsyncResult, 1)
	go p.WaitActivateRegionAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteProjectAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteProjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteProjectResult
	if asyncResult.Err != nil {
		callback <- DeleteProjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteProjectAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteProjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DeleteProjectAsync(
	request *DeleteProjectRequest,
	callback chan<- DeleteProjectAsyncResult,
) {
	path := "/account/me/project/{projectName}"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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

	go deleteProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DeleteProject(
	request *DeleteProjectRequest,
) (*DeleteProjectResult, error) {
	callback := make(chan DeleteProjectAsyncResult, 1)
	go p.DeleteProjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBillingMethodsAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBillingMethodsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBillingMethodsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBillingMethodsResult
	if asyncResult.Err != nil {
		callback <- DescribeBillingMethodsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBillingMethodsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBillingMethodsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeBillingMethodsAsync(
	request *DescribeBillingMethodsRequest,
	callback chan<- DescribeBillingMethodsAsyncResult,
) {
	path := "/account/me/billingMethod"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeBillingMethodsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeBillingMethods(
	request *DescribeBillingMethodsRequest,
) (*DescribeBillingMethodsResult, error) {
	callback := make(chan DescribeBillingMethodsAsyncResult, 1)
	go p.DescribeBillingMethodsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createBillingMethodAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- CreateBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateBillingMethodAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateBillingMethodResult
	if asyncResult.Err != nil {
		callback <- CreateBillingMethodAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateBillingMethodAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateBillingMethodAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) CreateBillingMethodAsync(
	request *CreateBillingMethodRequest,
	callback chan<- CreateBillingMethodAsyncResult,
) {
	path := "/account/me/billingMethod"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.MethodType != nil && *request.MethodType != "" {
		bodies["methodType"] = *request.MethodType
	}
	if request.CardCustomerId != nil && *request.CardCustomerId != "" {
		bodies["cardCustomerId"] = *request.CardCustomerId
	}
	if request.PartnerId != nil && *request.PartnerId != "" {
		bodies["partnerId"] = *request.PartnerId
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

	go createBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) CreateBillingMethod(
	request *CreateBillingMethodRequest,
) (*CreateBillingMethodResult, error) {
	callback := make(chan CreateBillingMethodAsyncResult, 1)
	go p.CreateBillingMethodAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBillingMethodAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBillingMethodAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBillingMethodResult
	if asyncResult.Err != nil {
		callback <- GetBillingMethodAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetBillingMethodAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetBillingMethodAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetBillingMethodAsync(
	request *GetBillingMethodRequest,
	callback chan<- GetBillingMethodAsyncResult,
) {
	path := "/account/me/billingMethod/{billingMethodName}"
	if request.BillingMethodName != nil && *request.BillingMethodName != "" {
		path = strings.ReplaceAll(path, "{billingMethodName}", core.ToString(*request.BillingMethodName))
	} else {
		path = strings.ReplaceAll(path, "{billingMethodName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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

	go getBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetBillingMethod(
	request *GetBillingMethodRequest,
) (*GetBillingMethodResult, error) {
	callback := make(chan GetBillingMethodAsyncResult, 1)
	go p.GetBillingMethodAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateBillingMethodAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateBillingMethodAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateBillingMethodResult
	if asyncResult.Err != nil {
		callback <- UpdateBillingMethodAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateBillingMethodAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateBillingMethodAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) UpdateBillingMethodAsync(
	request *UpdateBillingMethodRequest,
	callback chan<- UpdateBillingMethodAsyncResult,
) {
	path := "/account/me/billingMethod/{billingMethodName}"
	if request.BillingMethodName != nil && *request.BillingMethodName != "" {
		path = strings.ReplaceAll(path, "{billingMethodName}", core.ToString(*request.BillingMethodName))
	} else {
		path = strings.ReplaceAll(path, "{billingMethodName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
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

	go updateBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) UpdateBillingMethod(
	request *UpdateBillingMethodRequest,
) (*UpdateBillingMethodResult, error) {
	callback := make(chan UpdateBillingMethodAsyncResult, 1)
	go p.UpdateBillingMethodAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteBillingMethodAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteBillingMethodAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteBillingMethodResult
	if asyncResult.Err != nil {
		callback <- DeleteBillingMethodAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteBillingMethodAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteBillingMethodAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DeleteBillingMethodAsync(
	request *DeleteBillingMethodRequest,
	callback chan<- DeleteBillingMethodAsyncResult,
) {
	path := "/account/me/billingMethod/{billingMethodName}"
	if request.BillingMethodName != nil && *request.BillingMethodName != "" {
		path = strings.ReplaceAll(path, "{billingMethodName}", core.ToString(*request.BillingMethodName))
	} else {
		path = strings.ReplaceAll(path, "{billingMethodName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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

	go deleteBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DeleteBillingMethod(
	request *DeleteBillingMethodRequest,
) (*DeleteBillingMethodResult, error) {
	callback := make(chan DeleteBillingMethodAsyncResult, 1)
	go p.DeleteBillingMethodAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeReceiptsAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeReceiptsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeReceiptsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeReceiptsResult
	if asyncResult.Err != nil {
		callback <- DescribeReceiptsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeReceiptsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeReceiptsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeReceiptsAsync(
	request *DescribeReceiptsRequest,
	callback chan<- DescribeReceiptsAsyncResult,
) {
	path := "/account/me/receipt"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go describeReceiptsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeReceipts(
	request *DescribeReceiptsRequest,
) (*DescribeReceiptsResult, error) {
	callback := make(chan DescribeReceiptsAsyncResult, 1)
	go p.DescribeReceiptsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeBillingsAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBillingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBillingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBillingsResult
	if asyncResult.Err != nil {
		callback <- DescribeBillingsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeBillingsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeBillingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeBillingsAsync(
	request *DescribeBillingsRequest,
	callback chan<- DescribeBillingsAsyncResult,
) {
	path := "/account/me/billing/{projectName}/{year}/{month}"
	if request.ProjectName != nil && *request.ProjectName != "" {
		path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
	} else {
		path = strings.ReplaceAll(path, "{projectName}", "null")
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
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
	}
	if request.Region != nil {
		queryStrings["region"] = core.ToString(*request.Region)
	}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
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

	go describeBillingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeBillings(
	request *DescribeBillingsRequest,
) (*DescribeBillingsResult, error) {
	callback := make(chan DescribeBillingsAsyncResult, 1)
	go p.DescribeBillingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDumpProgressesAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDumpProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDumpProgressesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDumpProgressesResult
	if asyncResult.Err != nil {
		callback <- DescribeDumpProgressesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDumpProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDumpProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeDumpProgressesAsync(
	request *DescribeDumpProgressesRequest,
	callback chan<- DescribeDumpProgressesAsyncResult,
) {
	path := "/account/me/project/dump/progress"

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

	go describeDumpProgressesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeDumpProgresses(
	request *DescribeDumpProgressesRequest,
) (*DescribeDumpProgressesResult, error) {
	callback := make(chan DescribeDumpProgressesAsyncResult, 1)
	go p.DescribeDumpProgressesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDumpProgressAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetDumpProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDumpProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDumpProgressResult
	if asyncResult.Err != nil {
		callback <- GetDumpProgressAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDumpProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDumpProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetDumpProgressAsync(
	request *GetDumpProgressRequest,
	callback chan<- GetDumpProgressAsyncResult,
) {
	path := "/account/me/project/dump/progress/{transactionId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getDumpProgressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetDumpProgress(
	request *GetDumpProgressRequest,
) (*GetDumpProgressResult, error) {
	callback := make(chan GetDumpProgressAsyncResult, 1)
	go p.GetDumpProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func waitDumpUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- WaitDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WaitDumpUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WaitDumpUserDataResult
	if asyncResult.Err != nil {
		callback <- WaitDumpUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WaitDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) WaitDumpUserDataAsync(
	request *WaitDumpUserDataRequest,
	callback chan<- WaitDumpUserDataAsyncResult,
) {
	path := "/account/me/project/dump/progress/{transactionId}/wait"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
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

	go waitDumpUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) WaitDumpUserData(
	request *WaitDumpUserDataRequest,
) (*WaitDumpUserDataResult, error) {
	callback := make(chan WaitDumpUserDataAsyncResult, 1)
	go p.WaitDumpUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func archiveDumpUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- ArchiveDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ArchiveDumpUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ArchiveDumpUserDataResult
	if asyncResult.Err != nil {
		callback <- ArchiveDumpUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ArchiveDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ArchiveDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) ArchiveDumpUserDataAsync(
	request *ArchiveDumpUserDataRequest,
	callback chan<- ArchiveDumpUserDataAsyncResult,
) {
	path := "/account/me/project/dump/progress/{transactionId}/archive"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
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

	go archiveDumpUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) ArchiveDumpUserData(
	request *ArchiveDumpUserDataRequest,
) (*ArchiveDumpUserDataResult, error) {
	callback := make(chan ArchiveDumpUserDataAsyncResult, 1)
	go p.ArchiveDumpUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func dumpUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DumpUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DumpUserDataResult
	if asyncResult.Err != nil {
		callback <- DumpUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DumpUserDataAsync(
	request *DumpUserDataRequest,
	callback chan<- DumpUserDataAsyncResult,
) {
	path := "/account/me/project/dump/{userId}"
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

	go dumpUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DumpUserData(
	request *DumpUserDataRequest,
) (*DumpUserDataResult, error) {
	callback := make(chan DumpUserDataAsyncResult, 1)
	go p.DumpUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDumpUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDumpUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDumpUserDataResult
	if asyncResult.Err != nil {
		callback <- GetDumpUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetDumpUserDataAsync(
	request *GetDumpUserDataRequest,
	callback chan<- GetDumpUserDataAsyncResult,
) {
	path := "/account/me/project/dump/{transactionId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getDumpUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetDumpUserData(
	request *GetDumpUserDataRequest,
) (*GetDumpUserDataResult, error) {
	callback := make(chan GetDumpUserDataAsyncResult, 1)
	go p.GetDumpUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCleanProgressesAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCleanProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCleanProgressesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCleanProgressesResult
	if asyncResult.Err != nil {
		callback <- DescribeCleanProgressesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCleanProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCleanProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeCleanProgressesAsync(
	request *DescribeCleanProgressesRequest,
	callback chan<- DescribeCleanProgressesAsyncResult,
) {
	path := "/account/me/project/clean/progress"

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

	go describeCleanProgressesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeCleanProgresses(
	request *DescribeCleanProgressesRequest,
) (*DescribeCleanProgressesResult, error) {
	callback := make(chan DescribeCleanProgressesAsyncResult, 1)
	go p.DescribeCleanProgressesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCleanProgressAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetCleanProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCleanProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCleanProgressResult
	if asyncResult.Err != nil {
		callback <- GetCleanProgressAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCleanProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetCleanProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetCleanProgressAsync(
	request *GetCleanProgressRequest,
	callback chan<- GetCleanProgressAsyncResult,
) {
	path := "/account/me/project/clean/progress/{transactionId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getCleanProgressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetCleanProgress(
	request *GetCleanProgressRequest,
) (*GetCleanProgressResult, error) {
	callback := make(chan GetCleanProgressAsyncResult, 1)
	go p.GetCleanProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func waitCleanUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- WaitCleanUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WaitCleanUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WaitCleanUserDataResult
	if asyncResult.Err != nil {
		callback <- WaitCleanUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitCleanUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WaitCleanUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) WaitCleanUserDataAsync(
	request *WaitCleanUserDataRequest,
	callback chan<- WaitCleanUserDataAsyncResult,
) {
	path := "/account/me/project/clean/progress/{transactionId}/wait"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
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

	go waitCleanUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) WaitCleanUserData(
	request *WaitCleanUserDataRequest,
) (*WaitCleanUserDataResult, error) {
	callback := make(chan WaitCleanUserDataAsyncResult, 1)
	go p.WaitCleanUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func cleanUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- CleanUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CleanUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CleanUserDataResult
	if asyncResult.Err != nil {
		callback <- CleanUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CleanUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) CleanUserDataAsync(
	request *CleanUserDataRequest,
	callback chan<- CleanUserDataAsyncResult,
) {
	path := "/account/me/project/clean/{userId}"
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

	go cleanUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) CleanUserData(
	request *CleanUserDataRequest,
) (*CleanUserDataResult, error) {
	callback := make(chan CleanUserDataAsyncResult, 1)
	go p.CleanUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeImportProgressesAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeImportProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeImportProgressesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeImportProgressesResult
	if asyncResult.Err != nil {
		callback <- DescribeImportProgressesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeImportProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeImportProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeImportProgressesAsync(
	request *DescribeImportProgressesRequest,
	callback chan<- DescribeImportProgressesAsyncResult,
) {
	path := "/account/me/project/import/progress"

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

	go describeImportProgressesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeImportProgresses(
	request *DescribeImportProgressesRequest,
) (*DescribeImportProgressesResult, error) {
	callback := make(chan DescribeImportProgressesAsyncResult, 1)
	go p.DescribeImportProgressesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getImportProgressAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetImportProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetImportProgressAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetImportProgressResult
	if asyncResult.Err != nil {
		callback <- GetImportProgressAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetImportProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetImportProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetImportProgressAsync(
	request *GetImportProgressRequest,
	callback chan<- GetImportProgressAsyncResult,
) {
	path := "/account/me/project/import/progress/{transactionId}"
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
	if request.SourceRequestId != nil {
		headers["X-GS2-SOURCE-REQUEST-ID"] = string(*request.SourceRequestId)
	}
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getImportProgressAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetImportProgress(
	request *GetImportProgressRequest,
) (*GetImportProgressResult, error) {
	callback := make(chan GetImportProgressAsyncResult, 1)
	go p.GetImportProgressAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func waitImportUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- WaitImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- WaitImportUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result WaitImportUserDataResult
	if asyncResult.Err != nil {
		callback <- WaitImportUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- WaitImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) WaitImportUserDataAsync(
	request *WaitImportUserDataRequest,
	callback chan<- WaitImportUserDataAsyncResult,
) {
	path := "/account/me/project/import/progress/{transactionId}/wait"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
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

	go waitImportUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) WaitImportUserData(
	request *WaitImportUserDataRequest,
) (*WaitImportUserDataResult, error) {
	callback := make(chan WaitImportUserDataAsyncResult, 1)
	go p.WaitImportUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareImportUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareImportUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareImportUserDataResult
	if asyncResult.Err != nil {
		callback <- PrepareImportUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PrepareImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) PrepareImportUserDataAsync(
	request *PrepareImportUserDataRequest,
	callback chan<- PrepareImportUserDataAsyncResult,
) {
	path := "/account/me/project/import/{userId}/prepare"
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

	go prepareImportUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) PrepareImportUserData(
	request *PrepareImportUserDataRequest,
) (*PrepareImportUserDataResult, error) {
	callback := make(chan PrepareImportUserDataAsyncResult, 1)
	go p.PrepareImportUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func importUserDataAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- ImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ImportUserDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ImportUserDataResult
	if asyncResult.Err != nil {
		callback <- ImportUserDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) ImportUserDataAsync(
	request *ImportUserDataRequest,
	callback chan<- ImportUserDataAsyncResult,
) {
	path := "/account/me/project/import/{userId}"
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

	go importUserDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) ImportUserData(
	request *ImportUserDataRequest,
) (*ImportUserDataResult, error) {
	callback := make(chan ImportUserDataAsyncResult, 1)
	go p.ImportUserDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeImportErrorLogsAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeImportErrorLogsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeImportErrorLogsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeImportErrorLogsResult
	if asyncResult.Err != nil {
		callback <- DescribeImportErrorLogsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeImportErrorLogsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeImportErrorLogsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) DescribeImportErrorLogsAsync(
	request *DescribeImportErrorLogsRequest,
	callback chan<- DescribeImportErrorLogsAsyncResult,
) {
	path := "/account/me/project/import/progress/{transactionId}/log"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
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

	go describeImportErrorLogsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) DescribeImportErrorLogs(
	request *DescribeImportErrorLogsRequest,
) (*DescribeImportErrorLogsResult, error) {
	callback := make(chan DescribeImportErrorLogsAsyncResult, 1)
	go p.DescribeImportErrorLogsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getImportErrorLogAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- GetImportErrorLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetImportErrorLogAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetImportErrorLogResult
	if asyncResult.Err != nil {
		callback <- GetImportErrorLogAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetImportErrorLogAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetImportErrorLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) GetImportErrorLogAsync(
	request *GetImportErrorLogRequest,
	callback chan<- GetImportErrorLogAsyncResult,
) {
	path := "/account/me/project/import/progress/{transactionId}/log/{errorLogName}"
	if request.TransactionId != nil && *request.TransactionId != "" {
		path = strings.ReplaceAll(path, "{transactionId}", core.ToString(*request.TransactionId))
	} else {
		path = strings.ReplaceAll(path, "{transactionId}", "null")
	}
	if request.ErrorLogName != nil && *request.ErrorLogName != "" {
		path = strings.ReplaceAll(path, "{errorLogName}", core.ToString(*request.ErrorLogName))
	} else {
		path = strings.ReplaceAll(path, "{errorLogName}", "null")
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

	go getImportErrorLogAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) GetImportErrorLog(
	request *GetImportErrorLogRequest,
) (*GetImportErrorLogResult, error) {
	callback := make(chan GetImportErrorLogAsyncResult, 1)
	go p.GetImportErrorLogAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
