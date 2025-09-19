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

package deploy

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2DeployRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2DeployRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeStacksAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStacksAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStacksAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStacksResult
	if asyncResult.Err != nil {
		callback <- DescribeStacksAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStacksAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStacksAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DescribeStacksAsync(
	request *DescribeStacksRequest,
	callback chan<- DescribeStacksAsyncResult,
) {
	path := "/stack"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.NamePrefix != nil {
		queryStrings["namePrefix"] = core.ToString(*request.NamePrefix)
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

	go describeStacksAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DescribeStacks(
	request *DescribeStacksRequest,
) (*DescribeStacksResult, error) {
	callback := make(chan DescribeStacksAsyncResult, 1)
	go p.DescribeStacksAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preCreateStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- PreCreateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreCreateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreCreateStackResult
	if asyncResult.Err != nil {
		callback <- PreCreateStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreCreateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreCreateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) PreCreateStackAsync(
	request *PreCreateStackRequest,
	callback chan<- PreCreateStackAsyncResult,
) {
	path := "/stack/pre"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go preCreateStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) PreCreateStack(
	request *PreCreateStackRequest,
) (*PreCreateStackResult, error) {
	callback := make(chan PreCreateStackAsyncResult, 1)
	go p.PreCreateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- CreateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStackResult
	if asyncResult.Err != nil {
		callback <- CreateStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) CreateStackAsync(
	request *CreateStackRequest,
	callback chan<- CreateStackAsyncResult,
) {
	if request.Template != nil {
		res, err := p.PreCreateStack(
			&PreCreateStackRequest{
				ContextStack: request.ContextStack,
			},
		)
		if err != nil {
			callback <- CreateStackAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Template))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- CreateStackAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Template = nil
	}
	path := "/stack"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go createStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) CreateStack(
	request *CreateStackRequest,
) (*CreateStackResult, error) {
	callback := make(chan CreateStackAsyncResult, 1)
	go p.CreateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createStackFromGitHubAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- CreateStackFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateStackFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateStackFromGitHubResult
	if asyncResult.Err != nil {
		callback <- CreateStackFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- CreateStackFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- CreateStackFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) CreateStackFromGitHubAsync(
	request *CreateStackFromGitHubRequest,
	callback chan<- CreateStackFromGitHubAsyncResult,
) {
	path := "/stack/from_git_hub"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Name != nil && *request.Name != "" {
		bodies["name"] = *request.Name
	}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
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

	go createStackFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) CreateStackFromGitHub(
	request *CreateStackFromGitHubRequest,
) (*CreateStackFromGitHubResult, error) {
	callback := make(chan CreateStackFromGitHubAsyncResult, 1)
	go p.CreateStackFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preValidateAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- PreValidateAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreValidateAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreValidateResult
	if asyncResult.Err != nil {
		callback <- PreValidateAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreValidateAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreValidateAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) PreValidateAsync(
	request *PreValidateRequest,
	callback chan<- PreValidateAsyncResult,
) {
	path := "/stack/validate/pre"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go preValidateAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) PreValidate(
	request *PreValidateRequest,
) (*PreValidateResult, error) {
	callback := make(chan PreValidateAsyncResult, 1)
	go p.PreValidateAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func validateAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- ValidateAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ValidateAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ValidateResult
	if asyncResult.Err != nil {
		callback <- ValidateAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ValidateAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ValidateAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) ValidateAsync(
	request *ValidateRequest,
	callback chan<- ValidateAsyncResult,
) {
	if request.Template != nil {
		res, err := p.PreValidate(
			&PreValidateRequest{
				ContextStack: request.ContextStack,
			},
		)
		if err != nil {
			callback <- ValidateAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Template))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- ValidateAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Template = nil
	}
	path := "/stack/validate"

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go validateAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) Validate(
	request *ValidateRequest,
) (*ValidateResult, error) {
	callback := make(chan ValidateAsyncResult, 1)
	go p.ValidateAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStackStatusAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetStackStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStackStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStackStatusResult
	if asyncResult.Err != nil {
		callback <- GetStackStatusAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStackStatusAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStackStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetStackStatusAsync(
	request *GetStackStatusRequest,
	callback chan<- GetStackStatusAsyncResult,
) {
	path := "/stack/{stackName}/status"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go getStackStatusAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetStackStatus(
	request *GetStackStatusRequest,
) (*GetStackStatusResult, error) {
	callback := make(chan GetStackStatusAsyncResult, 1)
	go p.GetStackStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStackResult
	if asyncResult.Err != nil {
		callback <- GetStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetStackAsync(
	request *GetStackRequest,
	callback chan<- GetStackAsyncResult,
) {
	path := "/stack/{stackName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go getStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetStack(
	request *GetStackRequest,
) (*GetStackResult, error) {
	callback := make(chan GetStackAsyncResult, 1)
	go p.GetStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preUpdateStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- PreUpdateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreUpdateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreUpdateStackResult
	if asyncResult.Err != nil {
		callback <- PreUpdateStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreUpdateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreUpdateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) PreUpdateStackAsync(
	request *PreUpdateStackRequest,
	callback chan<- PreUpdateStackAsyncResult,
) {
	path := "/stack/{stackName}/pre"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go preUpdateStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) PreUpdateStack(
	request *PreUpdateStackRequest,
) (*PreUpdateStackResult, error) {
	callback := make(chan PreUpdateStackAsyncResult, 1)
	go p.PreUpdateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStackResult
	if asyncResult.Err != nil {
		callback <- UpdateStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) UpdateStackAsync(
	request *UpdateStackRequest,
	callback chan<- UpdateStackAsyncResult,
) {
	if request.Template != nil {
		res, err := p.PreUpdateStack(
			&PreUpdateStackRequest{
				ContextStack: request.ContextStack,
				StackName:    request.StackName,
			},
		)
		if err != nil {
			callback <- UpdateStackAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Template))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- UpdateStackAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Template = nil
	}
	path := "/stack/{stackName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go updateStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) UpdateStack(
	request *UpdateStackRequest,
) (*UpdateStackResult, error) {
	callback := make(chan UpdateStackAsyncResult, 1)
	go p.UpdateStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func preChangeSetAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- PreChangeSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PreChangeSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PreChangeSetResult
	if asyncResult.Err != nil {
		callback <- PreChangeSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- PreChangeSetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- PreChangeSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) PreChangeSetAsync(
	request *PreChangeSetRequest,
	callback chan<- PreChangeSetAsyncResult,
) {
	path := "/stack/{stackName}/pre"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	go preChangeSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) PreChangeSet(
	request *PreChangeSetRequest,
) (*PreChangeSetResult, error) {
	callback := make(chan PreChangeSetAsyncResult, 1)
	go p.PreChangeSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func changeSetAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- ChangeSetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ChangeSetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ChangeSetResult
	if asyncResult.Err != nil {
		callback <- ChangeSetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ChangeSetAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ChangeSetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) ChangeSetAsync(
	request *ChangeSetRequest,
	callback chan<- ChangeSetAsyncResult,
) {
	if request.Template != nil {
		res, err := p.PreChangeSet(
			&PreChangeSetRequest{
				ContextStack: request.ContextStack,
				StackName:    request.StackName,
			},
		)
		if err != nil {
			callback <- ChangeSetAsyncResult{
				err: err,
			}
			return
		}
		{
			req, _ := http.NewRequest("PUT", *res.UploadUrl, strings.NewReader(*request.Template))
			req.Header.Set("Content-Type", "application/json")

			client := new(http.Client)
			_, err = client.Do(req)
			if err != nil {
				callback <- ChangeSetAsyncResult{
					err: err,
				}
				return
			}
		}
		v := "preUpload"
		request.Mode = &v
		request.UploadToken = res.UploadToken
		request.Template = nil
	}
	path := "/stack/{stackName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Mode != nil && *request.Mode != "" {
		bodies["mode"] = *request.Mode
	}
	if request.Template != nil && *request.Template != "" {
		bodies["template"] = *request.Template
	}
	if request.UploadToken != nil && *request.UploadToken != "" {
		bodies["uploadToken"] = *request.UploadToken
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

	go changeSetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) ChangeSet(
	request *ChangeSetRequest,
) (*ChangeSetResult, error) {
	callback := make(chan ChangeSetAsyncResult, 1)
	go p.ChangeSetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateStackFromGitHubAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateStackFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateStackFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateStackFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateStackFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- UpdateStackFromGitHubAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- UpdateStackFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) UpdateStackFromGitHubAsync(
	request *UpdateStackFromGitHubRequest,
	callback chan<- UpdateStackFromGitHubAsyncResult,
) {
	path := "/stack/{stackName}/from_git_hub"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.Description != nil && *request.Description != "" {
		bodies["description"] = *request.Description
	}
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

	go updateStackFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:  core.Put,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) UpdateStackFromGitHub(
	request *UpdateStackFromGitHubRequest,
) (*UpdateStackFromGitHubResult, error) {
	callback := make(chan UpdateStackFromGitHubAsyncResult, 1)
	go p.UpdateStackFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackResult
	if asyncResult.Err != nil {
		callback <- DeleteStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DeleteStackAsync(
	request *DeleteStackRequest,
	callback chan<- DeleteStackAsyncResult,
) {
	path := "/stack/{stackName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go deleteStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DeleteStack(
	request *DeleteStackRequest,
) (*DeleteStackResult, error) {
	callback := make(chan DeleteStackAsyncResult, 1)
	go p.DeleteStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func forceDeleteStackAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- ForceDeleteStackAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ForceDeleteStackAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ForceDeleteStackResult
	if asyncResult.Err != nil {
		callback <- ForceDeleteStackAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- ForceDeleteStackAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- ForceDeleteStackAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) ForceDeleteStackAsync(
	request *ForceDeleteStackRequest,
	callback chan<- ForceDeleteStackAsyncResult,
) {
	path := "/stack/{stackName}/force"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go forceDeleteStackAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) ForceDeleteStack(
	request *ForceDeleteStackRequest,
) (*ForceDeleteStackResult, error) {
	callback := make(chan ForceDeleteStackAsyncResult, 1)
	go p.ForceDeleteStackAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStackResourcesAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStackResourcesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackResourcesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackResourcesResult
	if asyncResult.Err != nil {
		callback <- DeleteStackResourcesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackResourcesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStackResourcesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DeleteStackResourcesAsync(
	request *DeleteStackResourcesRequest,
	callback chan<- DeleteStackResourcesAsyncResult,
) {
	path := "/stack/{stackName}/resources"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go deleteStackResourcesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DeleteStackResources(
	request *DeleteStackResourcesRequest,
) (*DeleteStackResourcesResult, error) {
	callback := make(chan DeleteStackResourcesAsyncResult, 1)
	go p.DeleteStackResourcesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteStackEntityAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteStackEntityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStackEntityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStackEntityResult
	if asyncResult.Err != nil {
		callback <- DeleteStackEntityAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DeleteStackEntityAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DeleteStackEntityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DeleteStackEntityAsync(
	request *DeleteStackEntityRequest,
	callback chan<- DeleteStackEntityAsyncResult,
) {
	path := "/stack/{stackName}/entity"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go deleteStackEntityAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DeleteStackEntity(
	request *DeleteStackEntityRequest,
) (*DeleteStackEntityResult, error) {
	callback := make(chan DeleteStackEntityAsyncResult, 1)
	go p.DeleteStackEntityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getServiceVersionAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetServiceVersionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetServiceVersionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetServiceVersionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetServiceVersionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetServiceVersionAsync(
	request *GetServiceVersionRequest,
	callback chan<- GetServiceVersionAsyncResult,
) {
	path := "/system/version"

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

	go getServiceVersionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetServiceVersion(
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

func describeResourcesAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeResourcesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeResourcesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeResourcesResult
	if asyncResult.Err != nil {
		callback <- DescribeResourcesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeResourcesAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeResourcesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DescribeResourcesAsync(
	request *DescribeResourcesRequest,
	callback chan<- DescribeResourcesAsyncResult,
) {
	path := "/stack/{stackName}/resource"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go describeResourcesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DescribeResources(
	request *DescribeResourcesRequest,
) (*DescribeResourcesResult, error) {
	callback := make(chan DescribeResourcesAsyncResult, 1)
	go p.DescribeResourcesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getResourceAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetResourceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetResourceAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetResourceResult
	if asyncResult.Err != nil {
		callback <- GetResourceAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetResourceAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetResourceAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetResourceAsync(
	request *GetResourceRequest,
	callback chan<- GetResourceAsyncResult,
) {
	path := "/stack/{stackName}/resource/{resourceName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}
	if request.ResourceName != nil && *request.ResourceName != "" {
		path = strings.ReplaceAll(path, "{resourceName}", core.ToString(*request.ResourceName))
	} else {
		path = strings.ReplaceAll(path, "{resourceName}", "null")
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

	go getResourceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetResource(
	request *GetResourceRequest,
) (*GetResourceResult, error) {
	callback := make(chan GetResourceAsyncResult, 1)
	go p.GetResourceAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEventsAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEventsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEventsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEventsResult
	if asyncResult.Err != nil {
		callback <- DescribeEventsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeEventsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeEventsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DescribeEventsAsync(
	request *DescribeEventsRequest,
	callback chan<- DescribeEventsAsyncResult,
) {
	path := "/stack/{stackName}/event"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go describeEventsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DescribeEvents(
	request *DescribeEventsRequest,
) (*DescribeEventsResult, error) {
	callback := make(chan DescribeEventsAsyncResult, 1)
	go p.DescribeEventsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEventAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetEventAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEventAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEventResult
	if asyncResult.Err != nil {
		callback <- GetEventAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetEventAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetEventAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetEventAsync(
	request *GetEventRequest,
	callback chan<- GetEventAsyncResult,
) {
	path := "/stack/{stackName}/event/{eventName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}
	if request.EventName != nil && *request.EventName != "" {
		path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
	} else {
		path = strings.ReplaceAll(path, "{eventName}", "null")
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

	go getEventAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetEvent(
	request *GetEventRequest,
) (*GetEventResult, error) {
	callback := make(chan GetEventAsyncResult, 1)
	go p.GetEventAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeOutputsAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeOutputsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeOutputsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeOutputsResult
	if asyncResult.Err != nil {
		callback <- DescribeOutputsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeOutputsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeOutputsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) DescribeOutputsAsync(
	request *DescribeOutputsRequest,
	callback chan<- DescribeOutputsAsyncResult,
) {
	path := "/stack/{stackName}/output"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
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

	go describeOutputsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) DescribeOutputs(
	request *DescribeOutputsRequest,
) (*DescribeOutputsResult, error) {
	callback := make(chan DescribeOutputsAsyncResult, 1)
	go p.DescribeOutputsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getOutputAsyncHandler(
	client Gs2DeployRestClient,
	job *core.NetworkJob,
	callback chan<- GetOutputAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetOutputAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetOutputResult
	if asyncResult.Err != nil {
		callback <- GetOutputAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetOutputAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetOutputAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DeployRestClient) GetOutputAsync(
	request *GetOutputRequest,
	callback chan<- GetOutputAsyncResult,
) {
	path := "/stack/{stackName}/output/{outputName}"
	if request.StackName != nil && *request.StackName != "" {
		path = strings.ReplaceAll(path, "{stackName}", core.ToString(*request.StackName))
	} else {
		path = strings.ReplaceAll(path, "{stackName}", "null")
	}
	if request.OutputName != nil && *request.OutputName != "" {
		path = strings.ReplaceAll(path, "{outputName}", core.ToString(*request.OutputName))
	} else {
		path = strings.ReplaceAll(path, "{outputName}", "null")
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

	go getOutputAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("deploy").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DeployRestClient) GetOutput(
	request *GetOutputRequest,
) (*GetOutputResult, error) {
	callback := make(chan GetOutputAsyncResult, 1)
	go p.GetOutputAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
