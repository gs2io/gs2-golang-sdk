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
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2AuthRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2AuthRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func loginAsyncHandler(
	client Gs2AuthRestClient,
	job *core.NetworkJob,
	callback chan<- LoginAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- LoginAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- LoginAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- LoginAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthRestClient) LoginAsync(
	request *LoginRequest,
	callback chan<- LoginAsyncResult,
) {
	path := "/login"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
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

	go loginAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AuthRestClient) Login(
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

func loginBySignatureAsyncHandler(
	client Gs2AuthRestClient,
	job *core.NetworkJob,
	callback chan<- LoginBySignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- LoginBySignatureAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- LoginBySignatureAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- LoginBySignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthRestClient) LoginBySignatureAsync(
	request *LoginBySignatureRequest,
	callback chan<- LoginBySignatureAsyncResult,
) {
	path := "/login/signed"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go loginBySignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AuthRestClient) LoginBySignature(
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

func federationAsyncHandler(
	client Gs2AuthRestClient,
	job *core.NetworkJob,
	callback chan<- FederationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- FederationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- FederationAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- FederationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthRestClient) FederationAsync(
	request *FederationRequest,
	callback chan<- FederationAsyncResult,
) {
	path := "/federation"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go federationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AuthRestClient) Federation(
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

func issueTimeOffsetTokenByUserIdAsyncHandler(
	client Gs2AuthRestClient,
	job *core.NetworkJob,
	callback chan<- IssueTimeOffsetTokenByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- IssueTimeOffsetTokenByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2AuthRestClient) IssueTimeOffsetTokenByUserIdAsync(
	request *IssueTimeOffsetTokenByUserIdRequest,
	callback chan<- IssueTimeOffsetTokenByUserIdAsyncResult,
) {
	path := "/timeoffset/token"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserId != nil && *request.UserId != "" {
		bodies["userId"] = *request.UserId
	}
	if request.TimeOffset != nil {
		bodies["timeOffset"] = *request.TimeOffset
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

	go issueTimeOffsetTokenByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2AuthRestClient) IssueTimeOffsetTokenByUserId(
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
