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
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go verifyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go signInAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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

func issueAccountTokenAsyncHandler(
	client Gs2ProjectRestClient,
	job *core.NetworkJob,
	callback chan<- IssueAccountTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IssueAccountTokenAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IssueAccountTokenResult
	if asyncResult.Err != nil {
		callback <- IssueAccountTokenAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- IssueAccountTokenAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- IssueAccountTokenAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectRestClient) IssueAccountTokenAsync(
	request *IssueAccountTokenRequest,
	callback chan<- IssueAccountTokenAsyncResult,
) {
	path := "/account/accountToken"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AccountName != nil && *request.AccountName != "" {
        bodies["accountName"] = *request.AccountName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go issueAccountTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectRestClient) IssueAccountToken(
	request *IssueAccountTokenRequest,
) (*IssueAccountTokenResult, error) {
	callback := make(chan IssueAccountTokenAsyncResult, 1)
	go p.IssueAccountTokenAsync(
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
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go forgetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go issuePasswordAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateAccountAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
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

    headers := p.CreateAuthorizedHeaders()
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

    headers := p.CreateAuthorizedHeaders()
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.ProjectName != nil && *request.ProjectName != ""  {
        path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
    } else {
        path = strings.ReplaceAll(path, "{projectName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
	}

    headers := p.CreateAuthorizedHeaders()
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
    if request.ProjectName != nil && *request.ProjectName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getProjectTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.AccountName != nil && *request.AccountName != ""  {
        path = strings.ReplaceAll(path, "{accountName}", core.ToString(*request.AccountName))
    } else {
        path = strings.ReplaceAll(path, "{accountName}", "null")
    }
    if request.ProjectName != nil && *request.ProjectName != ""  {
        path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
    } else {
        path = strings.ReplaceAll(path, "{projectName}", "null")
    }
    if request.UserName != nil && *request.UserName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getProjectTokenByIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.ProjectName != nil && *request.ProjectName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateProjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
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
    if request.ProjectName != nil && *request.ProjectName != ""  {
        path = strings.ReplaceAll(path, "{projectName}", core.ToString(*request.ProjectName))
    } else {
        path = strings.ReplaceAll(path, "{projectName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
	}

    headers := p.CreateAuthorizedHeaders()
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

    headers := p.CreateAuthorizedHeaders()
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.BillingMethodName != nil && *request.BillingMethodName != ""  {
        path = strings.ReplaceAll(path, "{billingMethodName}", core.ToString(*request.BillingMethodName))
    } else {
        path = strings.ReplaceAll(path, "{billingMethodName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
	}

    headers := p.CreateAuthorizedHeaders()
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
    if request.BillingMethodName != nil && *request.BillingMethodName != ""  {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateBillingMethodAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("project").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
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
    if request.BillingMethodName != nil && *request.BillingMethodName != ""  {
        path = strings.ReplaceAll(path, "{billingMethodName}", core.ToString(*request.BillingMethodName))
    } else {
        path = strings.ReplaceAll(path, "{billingMethodName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.AccountToken != nil {
		queryStrings["accountToken"] = core.ToString(*request.AccountToken)
	}

    headers := p.CreateAuthorizedHeaders()
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

    headers := p.CreateAuthorizedHeaders()
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
    if request.ProjectName != nil && *request.ProjectName != ""  {
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

    headers := p.CreateAuthorizedHeaders()
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
