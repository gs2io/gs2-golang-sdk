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
	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2ProjectWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2ProjectWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2ProjectWebSocketClient) createAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) CreateAccountAsync(
	request *CreateAccountRequest,
	callback chan<- CreateAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "createAccount",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) CreateAccount(
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

func (p Gs2ProjectWebSocketClient) verifyAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- VerifyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) VerifyAsync(
	request *VerifyRequest,
	callback chan<- VerifyAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "verify",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.VerifyToken != nil && *request.VerifyToken != "" {
        bodies["verifyToken"] = *request.VerifyToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.verifyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) Verify(
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

func (p Gs2ProjectWebSocketClient) signInAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SignInAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) SignInAsync(
	request *SignInRequest,
	callback chan<- SignInAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "signIn",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.Email != nil && *request.Email != "" {
        bodies["email"] = *request.Email
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.signInAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) SignIn(
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

func (p Gs2ProjectWebSocketClient) issueAccountTokenAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IssueAccountTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) IssueAccountTokenAsync(
	request *IssueAccountTokenRequest,
	callback chan<- IssueAccountTokenAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "issueAccountToken",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountName != nil && *request.AccountName != "" {
        bodies["accountName"] = *request.AccountName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.issueAccountTokenAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) IssueAccountToken(
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

func (p Gs2ProjectWebSocketClient) forgetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ForgetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) ForgetAsync(
	request *ForgetRequest,
	callback chan<- ForgetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "forget",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.Email != nil && *request.Email != "" {
        bodies["email"] = *request.Email
    }
    if request.Lang != nil && *request.Lang != "" {
        bodies["lang"] = *request.Lang
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.forgetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) Forget(
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

func (p Gs2ProjectWebSocketClient) issuePasswordAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IssuePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) IssuePasswordAsync(
	request *IssuePasswordRequest,
	callback chan<- IssuePasswordAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "issuePassword",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.IssuePasswordToken != nil && *request.IssuePasswordToken != "" {
        bodies["issuePasswordToken"] = *request.IssuePasswordToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.issuePasswordAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) IssuePassword(
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

func (p Gs2ProjectWebSocketClient) updateAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) UpdateAccountAsync(
	request *UpdateAccountRequest,
	callback chan<- UpdateAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "updateAccount",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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

	go p.updateAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) UpdateAccount(
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

func (p Gs2ProjectWebSocketClient) deleteAccountAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteAccountAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DeleteAccountAsync(
	request *DeleteAccountRequest,
	callback chan<- DeleteAccountAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "account",
    		"function": "deleteAccount",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DeleteAccount(
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

func (p Gs2ProjectWebSocketClient) describeProjectsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeProjectsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DescribeProjectsAsync(
	request *DescribeProjectsRequest,
	callback chan<- DescribeProjectsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "describeProjects",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeProjectsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeProjects(
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

func (p Gs2ProjectWebSocketClient) createProjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) CreateProjectAsync(
	request *CreateProjectRequest,
	callback chan<- CreateProjectAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "createProject",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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

	go p.createProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) CreateProject(
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

func (p Gs2ProjectWebSocketClient) getProjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) GetProjectAsync(
	request *GetProjectRequest,
	callback chan<- GetProjectAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "getProject",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetProject(
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

func (p Gs2ProjectWebSocketClient) getProjectTokenAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProjectTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) GetProjectTokenAsync(
	request *GetProjectTokenRequest,
	callback chan<- GetProjectTokenAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "getProjectToken",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
    }
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getProjectTokenAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetProjectToken(
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

func (p Gs2ProjectWebSocketClient) getProjectTokenByIdentifierAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetProjectTokenByIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) GetProjectTokenByIdentifierAsync(
	request *GetProjectTokenByIdentifierRequest,
	callback chan<- GetProjectTokenByIdentifierAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "getProjectTokenByIdentifier",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountName != nil && *request.AccountName != "" {
        bodies["accountName"] = *request.AccountName
    }
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
    }
    if request.UserName != nil && *request.UserName != "" {
        bodies["userName"] = *request.UserName
    }
    if request.Password != nil && *request.Password != "" {
        bodies["password"] = *request.Password
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getProjectTokenByIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetProjectTokenByIdentifier(
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

func (p Gs2ProjectWebSocketClient) updateProjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) UpdateProjectAsync(
	request *UpdateProjectRequest,
	callback chan<- UpdateProjectAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "updateProject",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
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

	go p.updateProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) UpdateProject(
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

func (p Gs2ProjectWebSocketClient) deleteProjectAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteProjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DeleteProjectAsync(
	request *DeleteProjectRequest,
	callback chan<- DeleteProjectAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "project",
    		"function": "deleteProject",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DeleteProject(
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

func (p Gs2ProjectWebSocketClient) describeBillingMethodsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBillingMethodsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DescribeBillingMethodsAsync(
	request *DescribeBillingMethodsRequest,
	callback chan<- DescribeBillingMethodsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billingMethod",
    		"function": "describeBillingMethods",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeBillingMethodsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeBillingMethods(
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

func (p Gs2ProjectWebSocketClient) createBillingMethodAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) CreateBillingMethodAsync(
	request *CreateBillingMethodRequest,
	callback chan<- CreateBillingMethodAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billingMethod",
    		"function": "createBillingMethod",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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

	go p.createBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) CreateBillingMethod(
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

func (p Gs2ProjectWebSocketClient) getBillingMethodAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) GetBillingMethodAsync(
	request *GetBillingMethodRequest,
	callback chan<- GetBillingMethodAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billingMethod",
    		"function": "getBillingMethod",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.BillingMethodName != nil && *request.BillingMethodName != "" {
        bodies["billingMethodName"] = *request.BillingMethodName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetBillingMethod(
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

func (p Gs2ProjectWebSocketClient) updateBillingMethodAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) UpdateBillingMethodAsync(
	request *UpdateBillingMethodRequest,
	callback chan<- UpdateBillingMethodAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billingMethod",
    		"function": "updateBillingMethod",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.BillingMethodName != nil && *request.BillingMethodName != "" {
        bodies["billingMethodName"] = *request.BillingMethodName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) UpdateBillingMethod(
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

func (p Gs2ProjectWebSocketClient) deleteBillingMethodAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteBillingMethodAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DeleteBillingMethodAsync(
	request *DeleteBillingMethodRequest,
	callback chan<- DeleteBillingMethodAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billingMethod",
    		"function": "deleteBillingMethod",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.BillingMethodName != nil && *request.BillingMethodName != "" {
        bodies["billingMethodName"] = *request.BillingMethodName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DeleteBillingMethod(
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

func (p Gs2ProjectWebSocketClient) describeReceiptsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeReceiptsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DescribeReceiptsAsync(
	request *DescribeReceiptsRequest,
	callback chan<- DescribeReceiptsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "receipt",
    		"function": "describeReceipts",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeReceiptsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeReceipts(
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

func (p Gs2ProjectWebSocketClient) describeBillingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBillingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2ProjectWebSocketClient) DescribeBillingsAsync(
	request *DescribeBillingsRequest,
	callback chan<- DescribeBillingsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "project",
    		"component": "billing",
    		"function": "describeBillings",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.AccountToken != nil && *request.AccountToken != "" {
        bodies["accountToken"] = *request.AccountToken
    }
    if request.ProjectName != nil && *request.ProjectName != "" {
        bodies["projectName"] = *request.ProjectName
    }
    if request.Year != nil {
        bodies["year"] = *request.Year
    }
    if request.Month != nil {
        bodies["month"] = *request.Month
    }
    if request.Region != nil && *request.Region != "" {
        bodies["region"] = *request.Region
    }
    if request.Service != nil && *request.Service != "" {
        bodies["service"] = *request.Service
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeBillingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeBillings(
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
