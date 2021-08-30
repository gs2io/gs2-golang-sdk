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

package identifier

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2IdentifierRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2IdentifierRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeUsersAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeUsersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeUsersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeUsersResult
	if asyncResult.Err != nil {
		callback <- DescribeUsersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeUsersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeUsersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DescribeUsersAsync(
	request *DescribeUsersRequest,
	callback chan<- DescribeUsersAsyncResult,
) {
	path := "/user"

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

	go describeUsersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DescribeUsers(
	request *DescribeUsersRequest,
) (*DescribeUsersResult, error) {
	callback := make(chan DescribeUsersAsyncResult, 1)
	go p.DescribeUsersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createUserAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- CreateUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateUserResult
	if asyncResult.Err != nil {
		callback <- CreateUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) CreateUserAsync(
	request *CreateUserRequest,
	callback chan<- CreateUserAsyncResult,
) {
	path := "/user"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) CreateUser(
	request *CreateUserRequest,
) (*CreateUserResult, error) {
	callback := make(chan CreateUserAsyncResult, 1)
	go p.CreateUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateUserAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateUserResult
	if asyncResult.Err != nil {
		callback <- UpdateUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) UpdateUserAsync(
	request *UpdateUserRequest,
	callback chan<- UpdateUserAsyncResult,
) {
	path := "/user/{userName}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) UpdateUser(
	request *UpdateUserRequest,
) (*UpdateUserResult, error) {
	callback := make(chan UpdateUserAsyncResult, 1)
	go p.UpdateUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getUserAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- GetUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetUserResult
	if asyncResult.Err != nil {
		callback <- GetUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) GetUserAsync(
	request *GetUserRequest,
	callback chan<- GetUserAsyncResult,
) {
	path := "/user/{userName}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) GetUser(
	request *GetUserRequest,
) (*GetUserResult, error) {
	callback := make(chan GetUserAsyncResult, 1)
	go p.GetUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteUserAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteUserResult
	if asyncResult.Err != nil {
		callback <- DeleteUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DeleteUserAsync(
	request *DeleteUserRequest,
	callback chan<- DeleteUserAsyncResult,
) {
	path := "/user/{userName}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DeleteUser(
	request *DeleteUserRequest,
) (*DeleteUserResult, error) {
	callback := make(chan DeleteUserAsyncResult, 1)
	go p.DeleteUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSecurityPoliciesAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSecurityPoliciesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSecurityPoliciesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSecurityPoliciesResult
	if asyncResult.Err != nil {
		callback <- DescribeSecurityPoliciesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeSecurityPoliciesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeSecurityPoliciesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DescribeSecurityPoliciesAsync(
	request *DescribeSecurityPoliciesRequest,
	callback chan<- DescribeSecurityPoliciesAsyncResult,
) {
	path := "/securityPolicy"

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

	go describeSecurityPoliciesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DescribeSecurityPolicies(
	request *DescribeSecurityPoliciesRequest,
) (*DescribeSecurityPoliciesResult, error) {
	callback := make(chan DescribeSecurityPoliciesAsyncResult, 1)
	go p.DescribeSecurityPoliciesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCommonSecurityPoliciesAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCommonSecurityPoliciesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCommonSecurityPoliciesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCommonSecurityPoliciesResult
	if asyncResult.Err != nil {
		callback <- DescribeCommonSecurityPoliciesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeCommonSecurityPoliciesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeCommonSecurityPoliciesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DescribeCommonSecurityPoliciesAsync(
	request *DescribeCommonSecurityPoliciesRequest,
	callback chan<- DescribeCommonSecurityPoliciesAsyncResult,
) {
	path := "/securityPolicy/common"

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

	go describeCommonSecurityPoliciesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DescribeCommonSecurityPolicies(
	request *DescribeCommonSecurityPoliciesRequest,
) (*DescribeCommonSecurityPoliciesResult, error) {
	callback := make(chan DescribeCommonSecurityPoliciesAsyncResult, 1)
	go p.DescribeCommonSecurityPoliciesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- CreateSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- CreateSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) CreateSecurityPolicyAsync(
	request *CreateSecurityPolicyRequest,
	callback chan<- CreateSecurityPolicyAsyncResult,
) {
	path := "/securityPolicy"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Policy != nil && *request.Policy != "" {
		bodies["policy"] = *request.Policy
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) CreateSecurityPolicy(
	request *CreateSecurityPolicyRequest,
) (*CreateSecurityPolicyResult, error) {
	callback := make(chan CreateSecurityPolicyAsyncResult, 1)
	go p.CreateSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- UpdateSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) UpdateSecurityPolicyAsync(
	request *UpdateSecurityPolicyRequest,
	callback chan<- UpdateSecurityPolicyAsyncResult,
) {
	path := "/securityPolicy/{securityPolicyName}"
	if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
		path = strings.ReplaceAll(path, "{securityPolicyName}", core.ToString(*request.SecurityPolicyName))
	} else {
		path = strings.ReplaceAll(path, "{securityPolicyName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Policy != nil && *request.Policy != "" {
		bodies["policy"] = *request.Policy
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go updateSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) UpdateSecurityPolicy(
	request *UpdateSecurityPolicyRequest,
) (*UpdateSecurityPolicyResult, error) {
	callback := make(chan UpdateSecurityPolicyAsyncResult, 1)
	go p.UpdateSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- GetSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- GetSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) GetSecurityPolicyAsync(
	request *GetSecurityPolicyRequest,
	callback chan<- GetSecurityPolicyAsyncResult,
) {
	path := "/securityPolicy/{securityPolicyName}"
	if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
		path = strings.ReplaceAll(path, "{securityPolicyName}", core.ToString(*request.SecurityPolicyName))
	} else {
		path = strings.ReplaceAll(path, "{securityPolicyName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) GetSecurityPolicy(
	request *GetSecurityPolicyRequest,
) (*GetSecurityPolicyResult, error) {
	callback := make(chan GetSecurityPolicyAsyncResult, 1)
	go p.GetSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- DeleteSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DeleteSecurityPolicyAsync(
	request *DeleteSecurityPolicyRequest,
	callback chan<- DeleteSecurityPolicyAsyncResult,
) {
	path := "/securityPolicy/{securityPolicyName}"
	if request.SecurityPolicyName != nil && *request.SecurityPolicyName != "" {
		path = strings.ReplaceAll(path, "{securityPolicyName}", core.ToString(*request.SecurityPolicyName))
	} else {
		path = strings.ReplaceAll(path, "{securityPolicyName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DeleteSecurityPolicy(
	request *DeleteSecurityPolicyRequest,
) (*DeleteSecurityPolicyResult, error) {
	callback := make(chan DeleteSecurityPolicyAsyncResult, 1)
	go p.DeleteSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeIdentifiersAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeIdentifiersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIdentifiersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIdentifiersResult
	if asyncResult.Err != nil {
		callback <- DescribeIdentifiersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeIdentifiersAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeIdentifiersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DescribeIdentifiersAsync(
	request *DescribeIdentifiersRequest,
	callback chan<- DescribeIdentifiersAsyncResult,
) {
	path := "/user/{userName}/identifier"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
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

	go describeIdentifiersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DescribeIdentifiers(
	request *DescribeIdentifiersRequest,
) (*DescribeIdentifiersResult, error) {
	callback := make(chan DescribeIdentifiersAsyncResult, 1)
	go p.DescribeIdentifiersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createIdentifierAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- CreateIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateIdentifierResult
	if asyncResult.Err != nil {
		callback <- CreateIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) CreateIdentifierAsync(
	request *CreateIdentifierRequest,
	callback chan<- CreateIdentifierAsyncResult,
) {
	path := "/user/{userName}/identifier"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
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

	go createIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) CreateIdentifier(
	request *CreateIdentifierRequest,
) (*CreateIdentifierResult, error) {
	callback := make(chan CreateIdentifierAsyncResult, 1)
	go p.CreateIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getIdentifierAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- GetIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIdentifierResult
	if asyncResult.Err != nil {
		callback <- GetIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) GetIdentifierAsync(
	request *GetIdentifierRequest,
	callback chan<- GetIdentifierAsyncResult,
) {
	path := "/user/{userName}/identifier/{clientId}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}
	if request.ClientId != nil && *request.ClientId != "" {
		path = strings.ReplaceAll(path, "{clientId}", core.ToString(*request.ClientId))
	} else {
		path = strings.ReplaceAll(path, "{clientId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) GetIdentifier(
	request *GetIdentifierRequest,
) (*GetIdentifierResult, error) {
	callback := make(chan GetIdentifierAsyncResult, 1)
	go p.GetIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteIdentifierAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteIdentifierAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteIdentifierAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteIdentifierResult
	if asyncResult.Err != nil {
		callback <- DeleteIdentifierAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteIdentifierAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteIdentifierAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DeleteIdentifierAsync(
	request *DeleteIdentifierRequest,
	callback chan<- DeleteIdentifierAsyncResult,
) {
	path := "/user/{userName}/identifier/{clientId}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}
	if request.ClientId != nil && *request.ClientId != "" {
		path = strings.ReplaceAll(path, "{clientId}", core.ToString(*request.ClientId))
	} else {
		path = strings.ReplaceAll(path, "{clientId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deleteIdentifierAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DeleteIdentifier(
	request *DeleteIdentifierRequest,
) (*DeleteIdentifierResult, error) {
	callback := make(chan DeleteIdentifierAsyncResult, 1)
	go p.DeleteIdentifierAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describePasswordsAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DescribePasswordsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribePasswordsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribePasswordsResult
	if asyncResult.Err != nil {
		callback <- DescribePasswordsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribePasswordsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribePasswordsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DescribePasswordsAsync(
	request *DescribePasswordsRequest,
	callback chan<- DescribePasswordsAsyncResult,
) {
	path := "/user/{userName}/password"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
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

	go describePasswordsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DescribePasswords(
	request *DescribePasswordsRequest,
) (*DescribePasswordsResult, error) {
	callback := make(chan DescribePasswordsAsyncResult, 1)
	go p.DescribePasswordsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createPasswordAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- CreatePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreatePasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreatePasswordResult
	if asyncResult.Err != nil {
		callback <- CreatePasswordAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreatePasswordAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreatePasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) CreatePasswordAsync(
	request *CreatePasswordRequest,
	callback chan<- CreatePasswordAsyncResult,
) {
	path := "/user/{userName}/password"
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
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go createPasswordAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) CreatePassword(
	request *CreatePasswordRequest,
) (*CreatePasswordResult, error) {
	callback := make(chan CreatePasswordAsyncResult, 1)
	go p.CreatePasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getPasswordAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- GetPasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetPasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetPasswordResult
	if asyncResult.Err != nil {
		callback <- GetPasswordAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetPasswordAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetPasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) GetPasswordAsync(
	request *GetPasswordRequest,
	callback chan<- GetPasswordAsyncResult,
) {
	path := "/user/{userName}/password/entity"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getPasswordAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) GetPassword(
	request *GetPasswordRequest,
) (*GetPasswordResult, error) {
	callback := make(chan GetPasswordAsyncResult, 1)
	go p.GetPasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deletePasswordAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DeletePasswordAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeletePasswordAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeletePasswordResult
	if asyncResult.Err != nil {
		callback <- DeletePasswordAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeletePasswordAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeletePasswordAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DeletePasswordAsync(
	request *DeletePasswordRequest,
	callback chan<- DeletePasswordAsyncResult,
) {
	path := "/user/{userName}/password/entity"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go deletePasswordAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DeletePassword(
	request *DeletePasswordRequest,
) (*DeletePasswordResult, error) {
	callback := make(chan DeletePasswordAsyncResult, 1)
	go p.DeletePasswordAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getHasSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- GetHasSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetHasSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetHasSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- GetHasSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetHasSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetHasSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) GetHasSecurityPolicyAsync(
	request *GetHasSecurityPolicyRequest,
	callback chan<- GetHasSecurityPolicyAsyncResult,
) {
	path := "/user/{userName}/securityPolicy"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go getHasSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) GetHasSecurityPolicy(
	request *GetHasSecurityPolicyRequest,
) (*GetHasSecurityPolicyResult, error) {
	callback := make(chan GetHasSecurityPolicyAsyncResult, 1)
	go p.GetHasSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func attachSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- AttachSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AttachSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AttachSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- AttachSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- AttachSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- AttachSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) AttachSecurityPolicyAsync(
	request *AttachSecurityPolicyRequest,
	callback chan<- AttachSecurityPolicyAsyncResult,
) {
	path := "/user/{userName}/securityPolicy"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.SecurityPolicyId != nil && *request.SecurityPolicyId != "" {
		bodies["securityPolicyId"] = *request.SecurityPolicyId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go attachSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) AttachSecurityPolicy(
	request *AttachSecurityPolicyRequest,
) (*AttachSecurityPolicyResult, error) {
	callback := make(chan AttachSecurityPolicyAsyncResult, 1)
	go p.AttachSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func detachSecurityPolicyAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- DetachSecurityPolicyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DetachSecurityPolicyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DetachSecurityPolicyResult
	if asyncResult.Err != nil {
		callback <- DetachSecurityPolicyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DetachSecurityPolicyAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DetachSecurityPolicyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) DetachSecurityPolicyAsync(
	request *DetachSecurityPolicyRequest,
	callback chan<- DetachSecurityPolicyAsyncResult,
) {
	path := "/user/{userName}/securityPolicy/{securityPolicyId}"
	if request.UserName != nil && *request.UserName != "" {
		path = strings.ReplaceAll(path, "{userName}", core.ToString(*request.UserName))
	} else {
		path = strings.ReplaceAll(path, "{userName}", "null")
	}
	if request.SecurityPolicyId != nil && *request.SecurityPolicyId != "" {
		path = strings.ReplaceAll(path, "{securityPolicyId}", core.ToString(*request.SecurityPolicyId))
	} else {
		path = strings.ReplaceAll(path, "{securityPolicyId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go detachSecurityPolicyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) DetachSecurityPolicy(
	request *DetachSecurityPolicyRequest,
) (*DetachSecurityPolicyResult, error) {
	callback := make(chan DetachSecurityPolicyAsyncResult, 1)
	go p.DetachSecurityPolicyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func loginAsyncHandler(
	client Gs2IdentifierRestClient,
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

func (p Gs2IdentifierRestClient) LoginAsync(
	request *LoginRequest,
	callback chan<- LoginAsyncResult,
) {
	path := "/projectToken/login"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ClientId != nil && *request.ClientId != "" {
		bodies["client_id"] = *request.ClientId
	}
	if request.ClientSecret != nil && *request.ClientSecret != "" {
		bodies["client_secret"] = *request.ClientSecret
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.RequestId != nil {
		headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
	}

	go loginAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) Login(
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

func loginByUserAsyncHandler(
	client Gs2IdentifierRestClient,
	job *core.NetworkJob,
	callback chan<- LoginByUserAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- LoginByUserAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result LoginByUserResult
	if asyncResult.Err != nil {
		callback <- LoginByUserAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- LoginByUserAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- LoginByUserAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2IdentifierRestClient) LoginByUserAsync(
	request *LoginByUserRequest,
	callback chan<- LoginByUserAsyncResult,
) {
	path := "/projectToken/login/user"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.UserName != nil && *request.UserName != "" {
		bodies["userName"] = *request.UserName
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

	go loginByUserAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("identifier").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2IdentifierRestClient) LoginByUser(
	request *LoginByUserRequest,
) (*LoginByUserResult, error) {
	callback := make(chan LoginByUserAsyncResult, 1)
	go p.LoginByUserAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
