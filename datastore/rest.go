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

package datastore

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2DatastoreRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2DatastoreRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DescribeNamespaces(
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
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) CreateNamespaceAsync(
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
    if request.DoneUploadScript != nil {
        bodies["doneUploadScript"] = request.DoneUploadScript.ToDict()
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) CreateNamespace(
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
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) GetNamespaceStatus(
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
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) GetNamespace(
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
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) UpdateNamespaceAsync(
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
    if request.DoneUploadScript != nil {
        bodies["doneUploadScript"] = request.DoneUploadScript.ToDict()
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) UpdateNamespace(
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
	client Gs2DatastoreRestClient,
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

func (p Gs2DatastoreRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DeleteNamespace(
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

func describeDataObjectsAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDataObjectsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectsResult
	if asyncResult.Err != nil {
		callback <- DescribeDataObjectsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeDataObjectsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeDataObjectsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DescribeDataObjectsAsync(
	request *DescribeDataObjectsRequest,
	callback chan<- DescribeDataObjectsAsyncResult,
) {
	path := "/{namespaceName}/user/me/data"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Status != nil {
		queryStrings["status"] = core.ToString(*request.Status)
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go describeDataObjectsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DescribeDataObjects(
	request *DescribeDataObjectsRequest,
) (*DescribeDataObjectsResult, error) {
	callback := make(chan DescribeDataObjectsAsyncResult, 1)
	go p.DescribeDataObjectsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDataObjectsByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDataObjectsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeDataObjectsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeDataObjectsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeDataObjectsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DescribeDataObjectsByUserIdAsync(
	request *DescribeDataObjectsByUserIdRequest,
	callback chan<- DescribeDataObjectsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data"
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
	if request.Status != nil {
		queryStrings["status"] = core.ToString(*request.Status)
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

	go describeDataObjectsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DescribeDataObjectsByUserId(
	request *DescribeDataObjectsByUserIdRequest,
) (*DescribeDataObjectsByUserIdResult, error) {
	callback := make(chan DescribeDataObjectsByUserIdAsyncResult, 1)
	go p.DescribeDataObjectsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareUploadAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareUploadResult
	if asyncResult.Err != nil {
		callback <- PrepareUploadAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareUploadAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareUploadAsync(
	request *PrepareUploadRequest,
	callback chan<- PrepareUploadAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/file"
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
    if request.ContentType != nil && *request.ContentType != "" {
        bodies["contentType"] = *request.ContentType
    }
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []string
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
    }
    if request.UpdateIfExists != nil {
        bodies["updateIfExists"] = *request.UpdateIfExists
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

	go prepareUploadAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareUpload(
	request *PrepareUploadRequest,
) (*PrepareUploadResult, error) {
	callback := make(chan PrepareUploadAsyncResult, 1)
	go p.PrepareUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareUploadByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareUploadByUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareUploadByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareUploadByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareUploadByUserIdAsync(
	request *PrepareUploadByUserIdRequest,
	callback chan<- PrepareUploadByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/file"
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
    if request.Name != nil && *request.Name != "" {
        bodies["name"] = *request.Name
    }
    if request.ContentType != nil && *request.ContentType != "" {
        bodies["contentType"] = *request.ContentType
    }
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []string
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
    }
    if request.UpdateIfExists != nil {
        bodies["updateIfExists"] = *request.UpdateIfExists
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go prepareUploadByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareUploadByUserId(
	request *PrepareUploadByUserIdRequest,
) (*PrepareUploadByUserIdResult, error) {
	callback := make(chan PrepareUploadByUserIdAsyncResult, 1)
	go p.PrepareUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateDataObjectAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDataObjectResult
	if asyncResult.Err != nil {
		callback <- UpdateDataObjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateDataObjectAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) UpdateDataObjectAsync(
	request *UpdateDataObjectRequest,
	callback chan<- UpdateDataObjectAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []string
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
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

	go updateDataObjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) UpdateDataObject(
	request *UpdateDataObjectRequest,
) (*UpdateDataObjectResult, error) {
	callback := make(chan UpdateDataObjectAsyncResult, 1)
	go p.UpdateDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateDataObjectByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateDataObjectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateDataObjectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateDataObjectByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateDataObjectByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateDataObjectByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateDataObjectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) UpdateDataObjectByUserIdAsync(
	request *UpdateDataObjectByUserIdRequest,
	callback chan<- UpdateDataObjectByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []string
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateDataObjectByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) UpdateDataObjectByUserId(
	request *UpdateDataObjectByUserIdRequest,
) (*UpdateDataObjectByUserIdResult, error) {
	callback := make(chan UpdateDataObjectByUserIdAsyncResult, 1)
	go p.UpdateDataObjectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareReUploadAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareReUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareReUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareReUploadResult
	if asyncResult.Err != nil {
		callback <- PrepareReUploadAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareReUploadAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareReUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareReUploadAsync(
	request *PrepareReUploadRequest,
	callback chan<- PrepareReUploadAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}/file"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ContentType != nil && *request.ContentType != "" {
        bodies["contentType"] = *request.ContentType
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

	go prepareReUploadAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareReUpload(
	request *PrepareReUploadRequest,
) (*PrepareReUploadResult, error) {
	callback := make(chan PrepareReUploadAsyncResult, 1)
	go p.PrepareReUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareReUploadByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareReUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareReUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareReUploadByUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareReUploadByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareReUploadByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareReUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareReUploadByUserIdAsync(
	request *PrepareReUploadByUserIdRequest,
	callback chan<- PrepareReUploadByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/file"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.ContentType != nil && *request.ContentType != "" {
        bodies["contentType"] = *request.ContentType
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go prepareReUploadByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareReUploadByUserId(
	request *PrepareReUploadByUserIdRequest,
) (*PrepareReUploadByUserIdResult, error) {
	callback := make(chan PrepareReUploadByUserIdAsyncResult, 1)
	go p.PrepareReUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doneUploadAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DoneUploadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoneUploadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoneUploadResult
	if asyncResult.Err != nil {
		callback <- DoneUploadAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoneUploadAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DoneUploadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DoneUploadAsync(
	request *DoneUploadRequest,
	callback chan<- DoneUploadAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}/done"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
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

	go doneUploadAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DoneUpload(
	request *DoneUploadRequest,
) (*DoneUploadResult, error) {
	callback := make(chan DoneUploadAsyncResult, 1)
	go p.DoneUploadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doneUploadByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DoneUploadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoneUploadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoneUploadByUserIdResult
	if asyncResult.Err != nil {
		callback <- DoneUploadByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoneUploadByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DoneUploadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DoneUploadByUserIdAsync(
	request *DoneUploadByUserIdRequest,
	callback chan<- DoneUploadByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/done"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
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

	go doneUploadByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DoneUploadByUserId(
	request *DoneUploadByUserIdRequest,
) (*DoneUploadByUserIdResult, error) {
	callback := make(chan DoneUploadByUserIdAsyncResult, 1)
	go p.DoneUploadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteDataObjectAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataObjectResult
	if asyncResult.Err != nil {
		callback <- DeleteDataObjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteDataObjectAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DeleteDataObjectAsync(
	request *DeleteDataObjectRequest,
	callback chan<- DeleteDataObjectAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go deleteDataObjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DeleteDataObject(
	request *DeleteDataObjectRequest,
) (*DeleteDataObjectResult, error) {
	callback := make(chan DeleteDataObjectAsyncResult, 1)
	go p.DeleteDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteDataObjectByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteDataObjectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteDataObjectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteDataObjectByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteDataObjectByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteDataObjectByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteDataObjectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DeleteDataObjectByUserIdAsync(
	request *DeleteDataObjectByUserIdRequest,
	callback chan<- DeleteDataObjectByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}"
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
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteDataObjectByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DeleteDataObjectByUserId(
	request *DeleteDataObjectByUserIdRequest,
) (*DeleteDataObjectByUserIdResult, error) {
	callback := make(chan DeleteDataObjectByUserIdAsyncResult, 1)
	go p.DeleteDataObjectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadAsync(
	request *PrepareDownloadRequest,
	callback chan<- PrepareDownloadAsyncResult,
) {
	path := "/{namespaceName}/file"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.DataObjectId != nil && *request.DataObjectId != "" {
        bodies["dataObjectId"] = *request.DataObjectId
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

	go prepareDownloadAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownload(
	request *PrepareDownloadRequest,
) (*PrepareDownloadResult, error) {
	callback := make(chan PrepareDownloadAsyncResult, 1)
	go p.PrepareDownloadAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserIdAsync(
	request *PrepareDownloadByUserIdRequest,
	callback chan<- PrepareDownloadByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/file"
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
    if request.DataObjectId != nil && *request.DataObjectId != "" {
        bodies["dataObjectId"] = *request.DataObjectId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go prepareDownloadByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserId(
	request *PrepareDownloadByUserIdRequest,
) (*PrepareDownloadByUserIdResult, error) {
	callback := make(chan PrepareDownloadByUserIdAsyncResult, 1)
	go p.PrepareDownloadByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadByGenerationAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadByGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByGenerationResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadByGenerationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadByGenerationAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadByGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadByGenerationAsync(
	request *PrepareDownloadByGenerationRequest,
	callback chan<- PrepareDownloadByGenerationAsyncResult,
) {
	path := "/{namespaceName}/file/generation/{generation}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.DataObjectId != nil && *request.DataObjectId != "" {
        bodies["dataObjectId"] = *request.DataObjectId
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

	go prepareDownloadByGenerationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadByGeneration(
	request *PrepareDownloadByGenerationRequest,
) (*PrepareDownloadByGenerationResult, error) {
	callback := make(chan PrepareDownloadByGenerationAsyncResult, 1)
	go p.PrepareDownloadByGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadByGenerationAndUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadByGenerationAndUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByGenerationAndUserIdResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadByGenerationAndUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadByGenerationAndUserIdAsync(
	request *PrepareDownloadByGenerationAndUserIdRequest,
	callback chan<- PrepareDownloadByGenerationAndUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/file/generation/{generation}"
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
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.DataObjectId != nil && *request.DataObjectId != "" {
        bodies["dataObjectId"] = *request.DataObjectId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go prepareDownloadByGenerationAndUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadByGenerationAndUserId(
	request *PrepareDownloadByGenerationAndUserIdRequest,
) (*PrepareDownloadByGenerationAndUserIdResult, error) {
	callback := make(chan PrepareDownloadByGenerationAndUserIdAsyncResult, 1)
	go p.PrepareDownloadByGenerationAndUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadOwnDataAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadOwnDataAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadOwnDataAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadOwnDataResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadOwnDataAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadOwnDataAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadOwnDataAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadOwnDataAsync(
	request *PrepareDownloadOwnDataRequest,
	callback chan<- PrepareDownloadOwnDataAsyncResult,
) {
	path := "/{namespaceName}/user/me/file"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
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

	go prepareDownloadOwnDataAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadOwnData(
	request *PrepareDownloadOwnDataRequest,
) (*PrepareDownloadOwnDataResult, error) {
	callback := make(chan PrepareDownloadOwnDataAsyncResult, 1)
	go p.PrepareDownloadOwnDataAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadByUserIdAndDataObjectNameAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdAndDataObjectNameResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadByUserIdAndDataObjectNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserIdAndDataObjectNameAsync(
	request *PrepareDownloadByUserIdAndDataObjectNameRequest,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/file"
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
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go prepareDownloadByUserIdAndDataObjectNameAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserIdAndDataObjectName(
	request *PrepareDownloadByUserIdAndDataObjectNameRequest,
) (*PrepareDownloadByUserIdAndDataObjectNameResult, error) {
	callback := make(chan PrepareDownloadByUserIdAndDataObjectNameAsyncResult, 1)
	go p.PrepareDownloadByUserIdAndDataObjectNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadOwnDataByGenerationAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadOwnDataByGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadOwnDataByGenerationResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadOwnDataByGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadOwnDataByGenerationAsync(
	request *PrepareDownloadOwnDataByGenerationRequest,
	callback chan<- PrepareDownloadOwnDataByGenerationAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}/generation/{generation}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
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

	go prepareDownloadOwnDataByGenerationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadOwnDataByGeneration(
	request *PrepareDownloadOwnDataByGenerationRequest,
) (*PrepareDownloadOwnDataByGenerationResult, error) {
	callback := make(chan PrepareDownloadOwnDataByGenerationAsyncResult, 1)
	go p.PrepareDownloadOwnDataByGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func prepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult
	if asyncResult.Err != nil {
		callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsync(
	request *PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest,
	callback chan<- PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/generation/{generation}"
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
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
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

	go prepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) PrepareDownloadByUserIdAndDataObjectNameAndGeneration(
	request *PrepareDownloadByUserIdAndDataObjectNameAndGenerationRequest,
) (*PrepareDownloadByUserIdAndDataObjectNameAndGenerationResult, error) {
	callback := make(chan PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsyncResult, 1)
	go p.PrepareDownloadByUserIdAndDataObjectNameAndGenerationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func restoreDataObjectAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- RestoreDataObjectAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- RestoreDataObjectAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result RestoreDataObjectResult
	if asyncResult.Err != nil {
		callback <- RestoreDataObjectAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- RestoreDataObjectAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- RestoreDataObjectAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) RestoreDataObjectAsync(
	request *RestoreDataObjectRequest,
	callback chan<- RestoreDataObjectAsyncResult,
) {
	path := "/{namespaceName}/file/restore"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.DataObjectId != nil && *request.DataObjectId != "" {
        bodies["dataObjectId"] = *request.DataObjectId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go restoreDataObjectAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) RestoreDataObject(
	request *RestoreDataObjectRequest,
) (*RestoreDataObjectResult, error) {
	callback := make(chan RestoreDataObjectAsyncResult, 1)
	go p.RestoreDataObjectAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDataObjectHistoriesAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDataObjectHistoriesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectHistoriesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectHistoriesResult
	if asyncResult.Err != nil {
		callback <- DescribeDataObjectHistoriesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeDataObjectHistoriesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeDataObjectHistoriesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DescribeDataObjectHistoriesAsync(
	request *DescribeDataObjectHistoriesRequest,
	callback chan<- DescribeDataObjectHistoriesAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}/history"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go describeDataObjectHistoriesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DescribeDataObjectHistories(
	request *DescribeDataObjectHistoriesRequest,
) (*DescribeDataObjectHistoriesResult, error) {
	callback := make(chan DescribeDataObjectHistoriesAsyncResult, 1)
	go p.DescribeDataObjectHistoriesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeDataObjectHistoriesByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDataObjectHistoriesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDataObjectHistoriesByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeDataObjectHistoriesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) DescribeDataObjectHistoriesByUserIdAsync(
	request *DescribeDataObjectHistoriesByUserIdRequest,
	callback chan<- DescribeDataObjectHistoriesByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/history"
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
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
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

	go describeDataObjectHistoriesByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) DescribeDataObjectHistoriesByUserId(
	request *DescribeDataObjectHistoriesByUserIdRequest,
) (*DescribeDataObjectHistoriesByUserIdResult, error) {
	callback := make(chan DescribeDataObjectHistoriesByUserIdAsyncResult, 1)
	go p.DescribeDataObjectHistoriesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDataObjectHistoryAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- GetDataObjectHistoryAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataObjectHistoryAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataObjectHistoryResult
	if asyncResult.Err != nil {
		callback <- GetDataObjectHistoryAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetDataObjectHistoryAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetDataObjectHistoryAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) GetDataObjectHistoryAsync(
	request *GetDataObjectHistoryRequest,
	callback chan<- GetDataObjectHistoryAsyncResult,
) {
	path := "/{namespaceName}/user/me/data/{dataObjectName}/history/{generation}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
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

	go getDataObjectHistoryAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) GetDataObjectHistory(
	request *GetDataObjectHistoryRequest,
) (*GetDataObjectHistoryResult, error) {
	callback := make(chan GetDataObjectHistoryAsyncResult, 1)
	go p.GetDataObjectHistoryAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getDataObjectHistoryByUserIdAsyncHandler(
	client Gs2DatastoreRestClient,
	job *core.NetworkJob,
	callback chan<- GetDataObjectHistoryByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDataObjectHistoryByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDataObjectHistoryByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetDataObjectHistoryByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetDataObjectHistoryByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetDataObjectHistoryByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2DatastoreRestClient) GetDataObjectHistoryByUserIdAsync(
	request *GetDataObjectHistoryByUserIdRequest,
	callback chan<- GetDataObjectHistoryByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/data/{dataObjectName}/history/{generation}"
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
    if request.DataObjectName != nil && *request.DataObjectName != ""  {
        path = strings.ReplaceAll(path, "{dataObjectName}", core.ToString(*request.DataObjectName))
    } else {
        path = strings.ReplaceAll(path, "{dataObjectName}", "null")
    }
    if request.Generation != nil && *request.Generation != ""  {
        path = strings.ReplaceAll(path, "{generation}", core.ToString(*request.Generation))
    } else {
        path = strings.ReplaceAll(path, "{generation}", "null")
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

	go getDataObjectHistoryByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("datastore").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2DatastoreRestClient) GetDataObjectHistoryByUserId(
	request *GetDataObjectHistoryByUserIdRequest,
) (*GetDataObjectHistoryByUserIdResult, error) {
	callback := make(chan GetDataObjectHistoryByUserIdAsyncResult, 1)
	go p.GetDataObjectHistoryByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
