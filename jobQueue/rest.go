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

package jobQueue

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2JobQueueRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2JobQueueRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DescribeNamespaces(
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
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) CreateNamespaceAsync(
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
    if request.PushNotification != nil {
        bodies["pushNotification"] = request.PushNotification.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) CreateNamespace(
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
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) GetNamespaceStatus(
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
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) GetNamespace(
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
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.PushNotification != nil {
        bodies["pushNotification"] = request.PushNotification.ToDict()
    }
    if request.LogSetting != nil {
        bodies["logSetting"] = request.LogSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateNamespaceAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) UpdateNamespace(
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
	client Gs2JobQueueRestClient,
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

func (p Gs2JobQueueRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
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
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DeleteNamespace(
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

func describeJobsByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeJobsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeJobsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeJobsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeJobsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeJobsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeJobsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) DescribeJobsByUserIdAsync(
	request *DescribeJobsByUserIdRequest,
	callback chan<- DescribeJobsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/job"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go describeJobsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DescribeJobsByUserId(
	request *DescribeJobsByUserIdRequest,
) (*DescribeJobsByUserIdResult, error) {
	callback := make(chan DescribeJobsByUserIdAsyncResult, 1)
	go p.DescribeJobsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getJobByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- GetJobByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetJobByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetJobByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetJobByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetJobByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetJobByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) GetJobByUserIdAsync(
	request *GetJobByUserIdRequest,
	callback chan<- GetJobByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/job/{jobName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.JobName != nil && *request.JobName != ""  {
        path = strings.ReplaceAll(path, "{jobName}", core.ToString(*request.JobName))
    } else {
        path = strings.ReplaceAll(path, "{jobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getJobByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) GetJobByUserId(
	request *GetJobByUserIdRequest,
) (*GetJobByUserIdResult, error) {
	callback := make(chan GetJobByUserIdAsyncResult, 1)
	go p.GetJobByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func pushByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- PushByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PushByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PushByUserIdResult
	if asyncResult.Err != nil {
		callback <- PushByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PushByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PushByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) PushByUserIdAsync(
	request *PushByUserIdRequest,
	callback chan<- PushByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/job"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Jobs != nil {
        var _jobs []interface {}
        for _, item := range request.Jobs {
            _jobs = append(_jobs, item)
        }
        bodies["jobs"] = _jobs
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go pushByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) PushByUserId(
	request *PushByUserIdRequest,
) (*PushByUserIdResult, error) {
	callback := make(chan PushByUserIdAsyncResult, 1)
	go p.PushByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- RunAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunResult
	if asyncResult.Err != nil {
		callback <- RunAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RunAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RunAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) RunAsync(
	request *RunRequest,
	callback chan<- RunAsyncResult,
) {
	path := "/{namespaceName}/user/me/job/run"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go runAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) Run(
	request *RunRequest,
) (*RunResult, error) {
	callback := make(chan RunAsyncResult, 1)
	go p.RunAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func runByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- RunByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RunByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RunByUserIdResult
	if asyncResult.Err != nil {
		callback <- RunByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RunByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RunByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) RunByUserIdAsync(
	request *RunByUserIdRequest,
	callback chan<- RunByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/job/run"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go runByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) RunByUserId(
	request *RunByUserIdRequest,
) (*RunByUserIdResult, error) {
	callback := make(chan RunByUserIdAsyncResult, 1)
	go p.RunByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteJobByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteJobByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteJobByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteJobByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteJobByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteJobByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteJobByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) DeleteJobByUserIdAsync(
	request *DeleteJobByUserIdRequest,
	callback chan<- DeleteJobByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/job/{jobName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.JobName != nil && *request.JobName != ""  {
        path = strings.ReplaceAll(path, "{jobName}", core.ToString(*request.JobName))
    } else {
        path = strings.ReplaceAll(path, "{jobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteJobByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DeleteJobByUserId(
	request *DeleteJobByUserIdRequest,
) (*DeleteJobByUserIdResult, error) {
	callback := make(chan DeleteJobByUserIdAsyncResult, 1)
	go p.DeleteJobByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func pushByStampSheetAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- PushByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PushByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PushByStampSheetResult
	if asyncResult.Err != nil {
		callback <- PushByStampSheetAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PushByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PushByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) PushByStampSheetAsync(
	request *PushByStampSheetRequest,
	callback chan<- PushByStampSheetAsyncResult,
) {
	path := "/stamp/job"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go pushByStampSheetAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) PushByStampSheet(
	request *PushByStampSheetRequest,
) (*PushByStampSheetResult, error) {
	callback := make(chan PushByStampSheetAsyncResult, 1)
	go p.PushByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDeadLetterJobsByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDeadLetterJobsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDeadLetterJobsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDeadLetterJobsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeDeadLetterJobsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeDeadLetterJobsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeDeadLetterJobsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) DescribeDeadLetterJobsByUserIdAsync(
	request *DescribeDeadLetterJobsByUserIdRequest,
	callback chan<- DescribeDeadLetterJobsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/dead"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
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

	go describeDeadLetterJobsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DescribeDeadLetterJobsByUserId(
	request *DescribeDeadLetterJobsByUserIdRequest,
) (*DescribeDeadLetterJobsByUserIdResult, error) {
	callback := make(chan DescribeDeadLetterJobsByUserIdAsyncResult, 1)
	go p.DescribeDeadLetterJobsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDeadLetterJobByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- GetDeadLetterJobByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDeadLetterJobByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDeadLetterJobByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetDeadLetterJobByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetDeadLetterJobByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetDeadLetterJobByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) GetDeadLetterJobByUserIdAsync(
	request *GetDeadLetterJobByUserIdRequest,
	callback chan<- GetDeadLetterJobByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/dead/{deadLetterJobName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.DeadLetterJobName != nil && *request.DeadLetterJobName != ""  {
        path = strings.ReplaceAll(path, "{deadLetterJobName}", core.ToString(*request.DeadLetterJobName))
    } else {
        path = strings.ReplaceAll(path, "{deadLetterJobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getDeadLetterJobByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) GetDeadLetterJobByUserId(
	request *GetDeadLetterJobByUserIdRequest,
) (*GetDeadLetterJobByUserIdResult, error) {
	callback := make(chan GetDeadLetterJobByUserIdAsyncResult, 1)
	go p.GetDeadLetterJobByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteDeadLetterJobByUserIdAsyncHandler(
	client Gs2JobQueueRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDeadLetterJobByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDeadLetterJobByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDeadLetterJobByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteDeadLetterJobByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteDeadLetterJobByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteDeadLetterJobByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2JobQueueRestClient) DeleteDeadLetterJobByUserIdAsync(
	request *DeleteDeadLetterJobByUserIdRequest,
	callback chan<- DeleteDeadLetterJobByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/dead/{deadLetterJobName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.DeadLetterJobName != nil && *request.DeadLetterJobName != ""  {
        path = strings.ReplaceAll(path, "{deadLetterJobName}", core.ToString(*request.DeadLetterJobName))
    } else {
        path = strings.ReplaceAll(path, "{deadLetterJobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteDeadLetterJobByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("job-queue").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2JobQueueRestClient) DeleteDeadLetterJobByUserId(
	request *DeleteDeadLetterJobByUserIdRequest,
) (*DeleteDeadLetterJobByUserIdResult, error) {
	callback := make(chan DeleteDeadLetterJobByUserIdAsyncResult, 1)
	go p.DeleteDeadLetterJobByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
