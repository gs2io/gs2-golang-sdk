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

package auth

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2AuthWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2AuthWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2AuthWebSocketClient) loginAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- LoginAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- LoginAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result LoginResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- LoginAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- LoginAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthWebSocketClient) LoginAsync(
	request *LoginRequest,
	callback chan<- LoginAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "auth",
			"component":   "accessToken",
			"function":    "login",
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
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
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

	go p.loginAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2AuthWebSocketClient) Login(
	request *LoginRequest,
) (*LoginResult, error) {
	callback := make(chan LoginAsyncResult, 1)
	go p.LoginAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AuthWebSocketClient) loginBySignatureAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- LoginBySignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- LoginBySignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result LoginBySignatureResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- LoginBySignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- LoginBySignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthWebSocketClient) LoginBySignatureAsync(
	request *LoginBySignatureRequest,
	callback chan<- LoginBySignatureAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "auth",
			"component":   "accessToken",
			"function":    "loginBySignature",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.KeyId != nil && *request.KeyId != "" {
		bodies["keyId"] = *request.KeyId
	}
	if request.Body != nil && *request.Body != "" {
		bodies["body"] = *request.Body
	}
	if request.Signature != nil && *request.Signature != "" {
		bodies["signature"] = *request.Signature
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

	go p.loginBySignatureAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2AuthWebSocketClient) LoginBySignature(
	request *LoginBySignatureRequest,
) (*LoginBySignatureResult, error) {
	callback := make(chan LoginBySignatureAsyncResult, 1)
	go p.LoginBySignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AuthWebSocketClient) federationAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- FederationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- FederationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result FederationResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FederationAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- FederationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthWebSocketClient) FederationAsync(
	request *FederationRequest,
	callback chan<- FederationAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "auth",
			"component":   "accessToken",
			"function":    "federation",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.OriginalUserId != nil && *request.OriginalUserId != "" {
		bodies["originalUserId"] = *request.OriginalUserId
	}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.PolicyDocument != nil && *request.PolicyDocument != "" {
		bodies["policyDocument"] = *request.PolicyDocument
	}
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
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

	go p.federationAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2AuthWebSocketClient) Federation(
	request *FederationRequest,
) (*FederationResult, error) {
	callback := make(chan FederationAsyncResult, 1)
	go p.FederationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AuthWebSocketClient) issueTimeOffsetTokenByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- IssueTimeOffsetTokenByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IssueTimeOffsetTokenByUserIdResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthWebSocketClient) IssueTimeOffsetTokenByUserIdAsync(
	request *IssueTimeOffsetTokenByUserIdRequest,
	callback chan<- IssueTimeOffsetTokenByUserIdAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "auth",
			"component":   "accessToken",
			"function":    "issueTimeOffsetTokenByUserId",
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
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
	}
	if request.TimeOffsetToken != nil && *request.TimeOffsetToken != "" {
		bodies["timeOffsetToken"] = *request.TimeOffsetToken
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

	go p.issueTimeOffsetTokenByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2AuthWebSocketClient) IssueTimeOffsetTokenByUserId(
	request *IssueTimeOffsetTokenByUserIdRequest,
) (*IssueTimeOffsetTokenByUserIdResult, error) {
	callback := make(chan IssueTimeOffsetTokenByUserIdAsyncResult, 1)
	go p.IssueTimeOffsetTokenByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2AuthWebSocketClient) getServiceVersionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthWebSocketClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "auth",
			"component":   "accessToken",
			"function":    "getServiceVersion",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
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

	go p.getServiceVersionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2AuthWebSocketClient) GetServiceVersion(
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
