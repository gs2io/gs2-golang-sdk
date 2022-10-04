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

package serialKey

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2SerialKeyRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2SerialKeyRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DescribeNamespaces(
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
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) CreateNamespaceAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) CreateNamespace(
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
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetNamespaceStatus(
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
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetNamespace(
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
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) UpdateNamespaceAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UpdateNamespace(
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
	client Gs2SerialKeyRestClient,
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

func (p Gs2SerialKeyRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DeleteNamespace(
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

func describeIssueJobsAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeIssueJobsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeIssueJobsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeIssueJobsResult
	if asyncResult.Err != nil {
		callback <- DescribeIssueJobsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeIssueJobsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeIssueJobsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DescribeIssueJobsAsync(
	request *DescribeIssueJobsRequest,
	callback chan<- DescribeIssueJobsAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}/issue"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
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

	go describeIssueJobsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DescribeIssueJobs(
	request *DescribeIssueJobsRequest,
) (*DescribeIssueJobsResult, error) {
	callback := make(chan DescribeIssueJobsAsyncResult, 1)
	go p.DescribeIssueJobsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getIssueJobAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- GetIssueJobAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetIssueJobAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetIssueJobResult
	if asyncResult.Err != nil {
		callback <- GetIssueJobAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetIssueJobAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetIssueJobAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) GetIssueJobAsync(
	request *GetIssueJobRequest,
	callback chan<- GetIssueJobAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}/issue/{issueJobName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }
    if request.IssueJobName != nil && *request.IssueJobName != ""  {
        path = strings.ReplaceAll(path, "{issueJobName}", core.ToString(*request.IssueJobName))
    } else {
        path = strings.ReplaceAll(path, "{issueJobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getIssueJobAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetIssueJob(
	request *GetIssueJobRequest,
) (*GetIssueJobResult, error) {
	callback := make(chan GetIssueJobAsyncResult, 1)
	go p.GetIssueJobAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func issueAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- IssueAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- IssueAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result IssueResult
	if asyncResult.Err != nil {
		callback <- IssueAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- IssueAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- IssueAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) IssueAsync(
	request *IssueRequest,
	callback chan<- IssueAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.IssueRequestCount != nil {
        bodies["issueRequestCount"] = *request.IssueRequestCount
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go issueAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) Issue(
	request *IssueRequest,
) (*IssueResult, error) {
	callback := make(chan IssueAsyncResult, 1)
	go p.IssueAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeSerialKeysAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeSerialKeysAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSerialKeysAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSerialKeysResult
	if asyncResult.Err != nil {
		callback <- DescribeSerialKeysAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSerialKeysAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSerialKeysAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DescribeSerialKeysAsync(
	request *DescribeSerialKeysRequest,
	callback chan<- DescribeSerialKeysAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}/issue/{issueJobName}/serialKey"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }
    if request.IssueJobName != nil && *request.IssueJobName != ""  {
        path = strings.ReplaceAll(path, "{issueJobName}", core.ToString(*request.IssueJobName))
    } else {
        path = strings.ReplaceAll(path, "{issueJobName}", "null")
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

	go describeSerialKeysAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DescribeSerialKeys(
	request *DescribeSerialKeysRequest,
) (*DescribeSerialKeysResult, error) {
	callback := make(chan DescribeSerialKeysAsyncResult, 1)
	go p.DescribeSerialKeysAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func downloadSerialCodesAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DownloadSerialCodesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DownloadSerialCodesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DownloadSerialCodesResult
	if asyncResult.Err != nil {
		callback <- DownloadSerialCodesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DownloadSerialCodesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DownloadSerialCodesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DownloadSerialCodesAsync(
	request *DownloadSerialCodesRequest,
	callback chan<- DownloadSerialCodesAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}/issue/{issueJobName}/serialCode/download"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }
    if request.IssueJobName != nil && *request.IssueJobName != ""  {
        path = strings.ReplaceAll(path, "{issueJobName}", core.ToString(*request.IssueJobName))
    } else {
        path = strings.ReplaceAll(path, "{issueJobName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go downloadSerialCodesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DownloadSerialCodes(
	request *DownloadSerialCodesRequest,
) (*DownloadSerialCodesResult, error) {
	callback := make(chan DownloadSerialCodesAsyncResult, 1)
	go p.DownloadSerialCodesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getSerialKeyAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- GetSerialKeyAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSerialKeyAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSerialKeyResult
	if asyncResult.Err != nil {
		callback <- GetSerialKeyAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSerialKeyAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSerialKeyAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) GetSerialKeyAsync(
	request *GetSerialKeyRequest,
	callback chan<- GetSerialKeyAsyncResult,
) {
	path := "/{namespaceName}/serialKey/{code}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.Code != nil && *request.Code != ""  {
        path = strings.ReplaceAll(path, "{code}", core.ToString(*request.Code))
    } else {
        path = strings.ReplaceAll(path, "{code}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getSerialKeyAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetSerialKey(
	request *GetSerialKeyRequest,
) (*GetSerialKeyResult, error) {
	callback := make(chan GetSerialKeyAsyncResult, 1)
	go p.GetSerialKeyAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func useAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UseAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UseAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UseResult
	if asyncResult.Err != nil {
		callback <- UseAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UseAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UseAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UseAsync(
	request *UseRequest,
	callback chan<- UseAsyncResult,
) {
	path := "/{namespaceName}/user/me/serialKey"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Code != nil && *request.Code != "" {
        bodies["code"] = *request.Code
    }
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

	go useAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) Use(
	request *UseRequest,
) (*UseResult, error) {
	callback := make(chan UseAsyncResult, 1)
	go p.UseAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func useByUserIdAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UseByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UseByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UseByUserIdResult
	if asyncResult.Err != nil {
		callback <- UseByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UseByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UseByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UseByUserIdAsync(
	request *UseByUserIdRequest,
	callback chan<- UseByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/serialKey"
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
    if request.Code != nil && *request.Code != "" {
        bodies["code"] = *request.Code
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go useByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UseByUserId(
	request *UseByUserIdRequest,
) (*UseByUserIdResult, error) {
	callback := make(chan UseByUserIdAsyncResult, 1)
	go p.UseByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func useByStampTaskAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UseByStampTaskAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UseByStampTaskAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UseByStampTaskResult
	if asyncResult.Err != nil {
		callback <- UseByStampTaskAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UseByStampTaskAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UseByStampTaskAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UseByStampTaskAsync(
	request *UseByStampTaskRequest,
	callback chan<- UseByStampTaskAsyncResult,
) {
	path := "/serialKey/use"

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.StampTask != nil && *request.StampTask != "" {
        bodies["stampTask"] = *request.StampTask
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

	go useByStampTaskAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UseByStampTask(
	request *UseByStampTaskRequest,
) (*UseByStampTaskResult, error) {
	callback := make(chan UseByStampTaskAsyncResult, 1)
	go p.UseByStampTaskAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCampaignModelsAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCampaignModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCampaignModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCampaignModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeCampaignModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCampaignModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCampaignModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DescribeCampaignModelsAsync(
	request *DescribeCampaignModelsRequest,
	callback chan<- DescribeCampaignModelsAsyncResult,
) {
	path := "/{namespaceName}/campaign"
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

	go describeCampaignModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DescribeCampaignModels(
	request *DescribeCampaignModelsRequest,
) (*DescribeCampaignModelsResult, error) {
	callback := make(chan DescribeCampaignModelsAsyncResult, 1)
	go p.DescribeCampaignModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCampaignModelAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- GetCampaignModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCampaignModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCampaignModelResult
	if asyncResult.Err != nil {
		callback <- GetCampaignModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCampaignModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCampaignModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) GetCampaignModelAsync(
	request *GetCampaignModelRequest,
	callback chan<- GetCampaignModelAsyncResult,
) {
	path := "/{namespaceName}/campaign/{campaignModelName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getCampaignModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetCampaignModel(
	request *GetCampaignModelRequest,
) (*GetCampaignModelResult, error) {
	callback := make(chan GetCampaignModelAsyncResult, 1)
	go p.GetCampaignModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeCampaignModelMastersAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeCampaignModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCampaignModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCampaignModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeCampaignModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCampaignModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCampaignModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DescribeCampaignModelMastersAsync(
	request *DescribeCampaignModelMastersRequest,
	callback chan<- DescribeCampaignModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/campaign"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
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

	go describeCampaignModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DescribeCampaignModelMasters(
	request *DescribeCampaignModelMastersRequest,
) (*DescribeCampaignModelMastersResult, error) {
	callback := make(chan DescribeCampaignModelMastersAsyncResult, 1)
	go p.DescribeCampaignModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createCampaignModelMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- CreateCampaignModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateCampaignModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateCampaignModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateCampaignModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateCampaignModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateCampaignModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) CreateCampaignModelMasterAsync(
	request *CreateCampaignModelMasterRequest,
	callback chan<- CreateCampaignModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/campaign"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Name != nil && *request.Name != "" {
        bodies["name"] = *request.Name
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.EnableCampaignCode != nil {
        bodies["enableCampaignCode"] = *request.EnableCampaignCode
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createCampaignModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) CreateCampaignModelMaster(
	request *CreateCampaignModelMasterRequest,
) (*CreateCampaignModelMasterResult, error) {
	callback := make(chan CreateCampaignModelMasterAsyncResult, 1)
	go p.CreateCampaignModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCampaignModelMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- GetCampaignModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCampaignModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCampaignModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetCampaignModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCampaignModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCampaignModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) GetCampaignModelMasterAsync(
	request *GetCampaignModelMasterRequest,
	callback chan<- GetCampaignModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/campaign/{campaignModelName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getCampaignModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetCampaignModelMaster(
	request *GetCampaignModelMasterRequest,
) (*GetCampaignModelMasterResult, error) {
	callback := make(chan GetCampaignModelMasterAsyncResult, 1)
	go p.GetCampaignModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCampaignModelMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCampaignModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCampaignModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCampaignModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCampaignModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCampaignModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCampaignModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UpdateCampaignModelMasterAsync(
	request *UpdateCampaignModelMasterRequest,
	callback chan<- UpdateCampaignModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/campaign/{campaignModelName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.EnableCampaignCode != nil {
        bodies["enableCampaignCode"] = *request.EnableCampaignCode
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateCampaignModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UpdateCampaignModelMaster(
	request *UpdateCampaignModelMasterRequest,
) (*UpdateCampaignModelMasterResult, error) {
	callback := make(chan UpdateCampaignModelMasterAsyncResult, 1)
	go p.UpdateCampaignModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteCampaignModelMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteCampaignModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCampaignModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCampaignModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteCampaignModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteCampaignModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteCampaignModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) DeleteCampaignModelMasterAsync(
	request *DeleteCampaignModelMasterRequest,
	callback chan<- DeleteCampaignModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/campaign/{campaignModelName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.CampaignModelName != nil && *request.CampaignModelName != ""  {
        path = strings.ReplaceAll(path, "{campaignModelName}", core.ToString(*request.CampaignModelName))
    } else {
        path = strings.ReplaceAll(path, "{campaignModelName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteCampaignModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) DeleteCampaignModelMaster(
	request *DeleteCampaignModelMasterRequest,
) (*DeleteCampaignModelMasterResult, error) {
	callback := make(chan DeleteCampaignModelMasterAsyncResult, 1)
	go p.DeleteCampaignModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- ExportMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result ExportMasterResult
	if asyncResult.Err != nil {
		callback <- ExportMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- ExportMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- ExportMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
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

	go exportMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) ExportMaster(
	request *ExportMasterRequest,
) (*ExportMasterResult, error) {
	callback := make(chan ExportMasterAsyncResult, 1)
	go p.ExportMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getCurrentCampaignMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentCampaignMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentCampaignMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentCampaignMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentCampaignMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentCampaignMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentCampaignMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) GetCurrentCampaignMasterAsync(
	request *GetCurrentCampaignMasterRequest,
	callback chan<- GetCurrentCampaignMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
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

	go getCurrentCampaignMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) GetCurrentCampaignMaster(
	request *GetCurrentCampaignMasterRequest,
) (*GetCurrentCampaignMasterResult, error) {
	callback := make(chan GetCurrentCampaignMasterAsyncResult, 1)
	go p.GetCurrentCampaignMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentCampaignMasterAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentCampaignMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentCampaignMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentCampaignMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentCampaignMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentCampaignMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentCampaignMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UpdateCurrentCampaignMasterAsync(
	request *UpdateCurrentCampaignMasterRequest,
	callback chan<- UpdateCurrentCampaignMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Settings != nil && *request.Settings != "" {
        bodies["settings"] = *request.Settings
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateCurrentCampaignMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UpdateCurrentCampaignMaster(
	request *UpdateCurrentCampaignMasterRequest,
) (*UpdateCurrentCampaignMasterResult, error) {
	callback := make(chan UpdateCurrentCampaignMasterAsyncResult, 1)
	go p.UpdateCurrentCampaignMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentCampaignMasterFromGitHubAsyncHandler(
	client Gs2SerialKeyRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentCampaignMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentCampaignMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentCampaignMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentCampaignMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentCampaignMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentCampaignMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2SerialKeyRestClient) UpdateCurrentCampaignMasterFromGitHubAsync(
	request *UpdateCurrentCampaignMasterFromGitHubRequest,
	callback chan<- UpdateCurrentCampaignMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.CheckoutSetting != nil {
        bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateCurrentCampaignMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("serial-key").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2SerialKeyRestClient) UpdateCurrentCampaignMasterFromGitHub(
	request *UpdateCurrentCampaignMasterFromGitHubRequest,
) (*UpdateCurrentCampaignMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentCampaignMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentCampaignMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}