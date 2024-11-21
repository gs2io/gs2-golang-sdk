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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "createAccount",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.createAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "verify",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.VerifyToken != nil && *request.VerifyToken != "" {
		bodies["verifyToken"] = *request.VerifyToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.verifyAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "signIn",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.Otp != nil && *request.Otp != "" {
		bodies["otp"] = *request.Otp
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.signInAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "forget",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.forgetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "issuePassword",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.IssuePasswordToken != nil && *request.IssuePasswordToken != "" {
		bodies["issuePasswordToken"] = *request.IssuePasswordToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.issuePasswordAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "updateAccount",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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

func (p Gs2ProjectWebSocketClient) enableMfaAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- EnableMfaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- EnableMfaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result EnableMfaResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- EnableMfaAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- EnableMfaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) EnableMfaAsync(
	request *EnableMfaRequest,
	callback chan<- EnableMfaAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "enableMfa",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.enableMfaAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) EnableMfa(
	request *EnableMfaRequest,
) (*EnableMfaResult, error) {
	callback := make(chan EnableMfaAsyncResult, 1)
	go p.EnableMfaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ProjectWebSocketClient) challengeMfaAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ChallengeMfaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ChallengeMfaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ChallengeMfaResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ChallengeMfaAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ChallengeMfaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) ChallengeMfaAsync(
	request *ChallengeMfaRequest,
	callback chan<- ChallengeMfaAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "challengeMfa",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.Passcode != nil && *request.Passcode != "" {
		bodies["passcode"] = *request.Passcode
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.challengeMfaAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) ChallengeMfa(
	request *ChallengeMfaRequest,
) (*ChallengeMfaResult, error) {
	callback := make(chan ChallengeMfaAsyncResult, 1)
	go p.ChallengeMfaAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ProjectWebSocketClient) disableMfaAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DisableMfaAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DisableMfaAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DisableMfaResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DisableMfaAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DisableMfaAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DisableMfaAsync(
	request *DisableMfaRequest,
	callback chan<- DisableMfaAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "disableMfa",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.disableMfaAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DisableMfa(
	request *DisableMfaRequest,
) (*DisableMfaResult, error) {
	callback := make(chan DisableMfaAsyncResult, 1)
	go p.DisableMfaAsync(
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "account",
			"function":    "deleteAccount",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.AccountToken != nil && *request.AccountToken != "" {
		bodies["accountToken"] = *request.AccountToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteAccountAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "describeProjects",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeProjectsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "createProject",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.createProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "getProject",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "getProjectToken",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getProjectTokenAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "getProjectTokenByIdentifier",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.Otp != nil && *request.Otp != "" {
		bodies["otp"] = *request.Otp
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getProjectTokenByIdentifierAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "updateProject",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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

func (p Gs2ProjectWebSocketClient) activateRegionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ActivateRegionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ActivateRegionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ActivateRegionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) ActivateRegionAsync(
	request *ActivateRegionRequest,
	callback chan<- ActivateRegionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "activateRegion",
			"contentType": "application/json",
			"requestId":   requestId,
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
	if request.RegionName != nil && *request.RegionName != "" {
		bodies["regionName"] = *request.RegionName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.activateRegionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) ActivateRegion(
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

func (p Gs2ProjectWebSocketClient) waitActivateRegionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WaitActivateRegionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitActivateRegionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- WaitActivateRegionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) WaitActivateRegionAsync(
	request *WaitActivateRegionRequest,
	callback chan<- WaitActivateRegionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "waitActivateRegion",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ProjectName != nil && *request.ProjectName != "" {
		bodies["projectName"] = *request.ProjectName
	}
	if request.RegionName != nil && *request.RegionName != "" {
		bodies["regionName"] = *request.RegionName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.waitActivateRegionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) WaitActivateRegion(
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "project",
			"function":    "deleteProject",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteProjectAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billingMethod",
			"function":    "describeBillingMethods",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeBillingMethodsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billingMethod",
			"function":    "createBillingMethod",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.createBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billingMethod",
			"function":    "getBillingMethod",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billingMethod",
			"function":    "updateBillingMethod",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.updateBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billingMethod",
			"function":    "deleteBillingMethod",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.deleteBillingMethodAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "receipt",
			"function":    "describeReceipts",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeReceiptsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "billing",
			"function":    "describeBillings",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeBillingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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

func (p Gs2ProjectWebSocketClient) describeDumpProgressesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDumpProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDumpProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeDumpProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DescribeDumpProgressesAsync(
	request *DescribeDumpProgressesRequest,
	callback chan<- DescribeDumpProgressesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "describeDumpProgresses",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
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
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeDumpProgressesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeDumpProgresses(
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

func (p Gs2ProjectWebSocketClient) getDumpProgressAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDumpProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDumpProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDumpProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) GetDumpProgressAsync(
	request *GetDumpProgressRequest,
	callback chan<- GetDumpProgressAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "getDumpProgress",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getDumpProgressAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetDumpProgress(
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

func (p Gs2ProjectWebSocketClient) waitDumpUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WaitDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- WaitDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) WaitDumpUserDataAsync(
	request *WaitDumpUserDataRequest,
	callback chan<- WaitDumpUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "waitDumpUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.waitDumpUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) WaitDumpUserData(
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

func (p Gs2ProjectWebSocketClient) archiveDumpUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ArchiveDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ArchiveDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ArchiveDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) ArchiveDumpUserDataAsync(
	request *ArchiveDumpUserDataRequest,
	callback chan<- ArchiveDumpUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "archiveDumpUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.archiveDumpUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) ArchiveDumpUserData(
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

func (p Gs2ProjectWebSocketClient) dumpUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DumpUserDataAsync(
	request *DumpUserDataRequest,
	callback chan<- DumpUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "dumpUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.dumpUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DumpUserData(
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

func (p Gs2ProjectWebSocketClient) getDumpUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDumpUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDumpUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDumpUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) GetDumpUserDataAsync(
	request *GetDumpUserDataRequest,
	callback chan<- GetDumpUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "dumpProgress",
			"function":    "getDumpUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getDumpUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetDumpUserData(
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

func (p Gs2ProjectWebSocketClient) describeCleanProgressesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCleanProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCleanProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeCleanProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DescribeCleanProgressesAsync(
	request *DescribeCleanProgressesRequest,
	callback chan<- DescribeCleanProgressesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "cleanProgress",
			"function":    "describeCleanProgresses",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
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
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeCleanProgressesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeCleanProgresses(
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

func (p Gs2ProjectWebSocketClient) getCleanProgressAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCleanProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetCleanProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetCleanProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) GetCleanProgressAsync(
	request *GetCleanProgressRequest,
	callback chan<- GetCleanProgressAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "cleanProgress",
			"function":    "getCleanProgress",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getCleanProgressAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetCleanProgress(
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

func (p Gs2ProjectWebSocketClient) waitCleanUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WaitCleanUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitCleanUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- WaitCleanUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) WaitCleanUserDataAsync(
	request *WaitCleanUserDataRequest,
	callback chan<- WaitCleanUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "cleanProgress",
			"function":    "waitCleanUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.waitCleanUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) WaitCleanUserData(
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

func (p Gs2ProjectWebSocketClient) cleanUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CleanUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CleanUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- CleanUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) CleanUserDataAsync(
	request *CleanUserDataRequest,
	callback chan<- CleanUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "cleanProgress",
			"function":    "cleanUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.cleanUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) CleanUserData(
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

func (p Gs2ProjectWebSocketClient) describeImportProgressesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeImportProgressesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeImportProgressesAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeImportProgressesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DescribeImportProgressesAsync(
	request *DescribeImportProgressesRequest,
	callback chan<- DescribeImportProgressesAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importProgress",
			"function":    "describeImportProgresses",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
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
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeImportProgressesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeImportProgresses(
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

func (p Gs2ProjectWebSocketClient) getImportProgressAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetImportProgressAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetImportProgressAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetImportProgressAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) GetImportProgressAsync(
	request *GetImportProgressRequest,
	callback chan<- GetImportProgressAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importProgress",
			"function":    "getImportProgress",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getImportProgressAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetImportProgress(
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

func (p Gs2ProjectWebSocketClient) waitImportUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- WaitImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- WaitImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- WaitImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) WaitImportUserDataAsync(
	request *WaitImportUserDataRequest,
	callback chan<- WaitImportUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importProgress",
			"function":    "waitImportUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.MicroserviceName != nil && *request.MicroserviceName != "" {
		bodies["microserviceName"] = *request.MicroserviceName
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.waitImportUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) WaitImportUserData(
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

func (p Gs2ProjectWebSocketClient) prepareImportUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PrepareImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PrepareImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- PrepareImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) PrepareImportUserDataAsync(
	request *PrepareImportUserDataRequest,
	callback chan<- PrepareImportUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importProgress",
			"function":    "prepareImportUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.prepareImportUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) PrepareImportUserData(
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

func (p Gs2ProjectWebSocketClient) importUserDataAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ImportUserDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ImportUserDataAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- ImportUserDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) ImportUserDataAsync(
	request *ImportUserDataRequest,
	callback chan<- ImportUserDataAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importProgress",
			"function":    "importUserData",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DuplicationAvoider != nil {
		bodies["xGs2DuplicationAvoider"] = string(*request.DuplicationAvoider)
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.importUserDataAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) ImportUserData(
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

func (p Gs2ProjectWebSocketClient) describeImportErrorLogsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeImportErrorLogsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeImportErrorLogsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeImportErrorLogsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) DescribeImportErrorLogsAsync(
	request *DescribeImportErrorLogsRequest,
	callback chan<- DescribeImportErrorLogsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importErrorLog",
			"function":    "describeImportErrorLogs",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
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
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeImportErrorLogsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) DescribeImportErrorLogs(
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

func (p Gs2ProjectWebSocketClient) getImportErrorLogAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetImportErrorLogAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetImportErrorLogAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetImportErrorLogAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ProjectWebSocketClient) GetImportErrorLogAsync(
	request *GetImportErrorLogRequest,
	callback chan<- GetImportErrorLogAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "project",
			"component":   "importErrorLog",
			"function":    "getImportErrorLog",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.TransactionId != nil && *request.TransactionId != "" {
		bodies["transactionId"] = *request.TransactionId
	}
	if request.ErrorLogName != nil && *request.ErrorLogName != "" {
		bodies["errorLogName"] = *request.ErrorLogName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getImportErrorLogAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2ProjectWebSocketClient) GetImportErrorLog(
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
