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
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go loginAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go loginBySignatureAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("auth").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
