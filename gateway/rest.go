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

package gateway

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2GatewayRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2GatewayRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DescribeNamespaces(
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
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) CreateNamespaceAsync(
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
    if request.FirebaseSecret != nil && *request.FirebaseSecret != "" {
        bodies["firebaseSecret"] = *request.FirebaseSecret
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) CreateNamespace(
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
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) GetNamespaceStatus(
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
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) GetNamespace(
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
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) UpdateNamespaceAsync(
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
    if request.FirebaseSecret != nil && *request.FirebaseSecret != "" {
        bodies["firebaseSecret"] = *request.FirebaseSecret
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) UpdateNamespace(
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
	client Gs2GatewayRestClient,
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

func (p Gs2GatewayRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DeleteNamespace(
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

func describeWebSocketSessionsAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeWebSocketSessionsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWebSocketSessionsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWebSocketSessionsResult
	if asyncResult.Err != nil {
		callback <- DescribeWebSocketSessionsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeWebSocketSessionsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeWebSocketSessionsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DescribeWebSocketSessionsAsync(
	request *DescribeWebSocketSessionsRequest,
	callback chan<- DescribeWebSocketSessionsAsyncResult,
) {
	path := "/{namespaceName}/session/user/me"
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go describeWebSocketSessionsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DescribeWebSocketSessions(
	request *DescribeWebSocketSessionsRequest,
) (*DescribeWebSocketSessionsResult, error) {
	callback := make(chan DescribeWebSocketSessionsAsyncResult, 1)
	go p.DescribeWebSocketSessionsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeWebSocketSessionsByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeWebSocketSessionsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeWebSocketSessionsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeWebSocketSessionsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeWebSocketSessionsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeWebSocketSessionsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeWebSocketSessionsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DescribeWebSocketSessionsByUserIdAsync(
	request *DescribeWebSocketSessionsByUserIdRequest,
	callback chan<- DescribeWebSocketSessionsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/session/user/{userId}"
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

	go describeWebSocketSessionsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DescribeWebSocketSessionsByUserId(
	request *DescribeWebSocketSessionsByUserIdRequest,
) (*DescribeWebSocketSessionsByUserIdResult, error) {
	callback := make(chan DescribeWebSocketSessionsByUserIdAsyncResult, 1)
	go p.DescribeWebSocketSessionsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SetUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetUserIdResult
	if asyncResult.Err != nil {
		callback <- SetUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SetUserIdAsync(
	request *SetUserIdRequest,
	callback chan<- SetUserIdAsyncResult,
) {
	path := "/{namespaceName}/session/user/me/user"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AllowConcurrentAccess != nil {
        bodies["allowConcurrentAccess"] = *request.AllowConcurrentAccess
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

	go setUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SetUserId(
	request *SetUserIdRequest,
) (*SetUserIdResult, error) {
	callback := make(chan SetUserIdAsyncResult, 1)
	go p.SetUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setUserIdByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SetUserIdByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetUserIdByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetUserIdByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetUserIdByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetUserIdByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetUserIdByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SetUserIdByUserIdAsync(
	request *SetUserIdByUserIdRequest,
	callback chan<- SetUserIdByUserIdAsyncResult,
) {
	path := "/{namespaceName}/session/user/{userId}/user"
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
    if request.AllowConcurrentAccess != nil {
        bodies["allowConcurrentAccess"] = *request.AllowConcurrentAccess
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

	go setUserIdByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SetUserIdByUserId(
	request *SetUserIdByUserIdRequest,
) (*SetUserIdByUserIdResult, error) {
	callback := make(chan SetUserIdByUserIdAsyncResult, 1)
	go p.SetUserIdByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendNotificationAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SendNotificationAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendNotificationAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendNotificationResult
	if asyncResult.Err != nil {
		callback <- SendNotificationAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SendNotificationAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SendNotificationAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SendNotificationAsync(
	request *SendNotificationRequest,
	callback chan<- SendNotificationAsyncResult,
) {
	path := "/{namespaceName}/session/user/{userId}/notification"
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
    if request.Subject != nil && *request.Subject != "" {
        bodies["subject"] = *request.Subject
    }
    if request.Payload != nil && *request.Payload != "" {
        bodies["payload"] = *request.Payload
    }
    if request.EnableTransferMobileNotification != nil {
        bodies["enableTransferMobileNotification"] = *request.EnableTransferMobileNotification
    }
    if request.Sound != nil && *request.Sound != "" {
        bodies["sound"] = *request.Sound
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

	go sendNotificationAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SendNotification(
	request *SendNotificationRequest,
) (*SendNotificationResult, error) {
	callback := make(chan SendNotificationAsyncResult, 1)
	go p.SendNotificationAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func disconnectByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DisconnectByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DisconnectByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DisconnectByUserIdResult
	if asyncResult.Err != nil {
		callback <- DisconnectByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DisconnectByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DisconnectByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DisconnectByUserIdAsync(
	request *DisconnectByUserIdRequest,
	callback chan<- DisconnectByUserIdAsyncResult,
) {
	path := "/{namespaceName}/session/user/{userId}"
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go disconnectByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DisconnectByUserId(
	request *DisconnectByUserIdRequest,
) (*DisconnectByUserIdResult, error) {
	callback := make(chan DisconnectByUserIdAsyncResult, 1)
	go p.DisconnectByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func disconnectAllAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DisconnectAllAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DisconnectAllAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DisconnectAllResult
	if asyncResult.Err != nil {
		callback <- DisconnectAllAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DisconnectAllAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DisconnectAllAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DisconnectAllAsync(
	request *DisconnectAllRequest,
	callback chan<- DisconnectAllAsyncResult,
) {
	path := "/{namespaceName}/session"
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

	go disconnectAllAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DisconnectAll(
	request *DisconnectAllRequest,
) (*DisconnectAllResult, error) {
	callback := make(chan DisconnectAllAsyncResult, 1)
	go p.DisconnectAllAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setFirebaseTokenAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SetFirebaseTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFirebaseTokenAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFirebaseTokenResult
	if asyncResult.Err != nil {
		callback <- SetFirebaseTokenAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetFirebaseTokenAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetFirebaseTokenAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SetFirebaseTokenAsync(
	request *SetFirebaseTokenRequest,
	callback chan<- SetFirebaseTokenAsyncResult,
) {
	path := "/{namespaceName}/user/me/firebase/token"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Token != nil && *request.Token != "" {
        bodies["token"] = *request.Token
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

	go setFirebaseTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SetFirebaseToken(
	request *SetFirebaseTokenRequest,
) (*SetFirebaseTokenResult, error) {
	callback := make(chan SetFirebaseTokenAsyncResult, 1)
	go p.SetFirebaseTokenAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func setFirebaseTokenByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SetFirebaseTokenByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetFirebaseTokenByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetFirebaseTokenByUserIdResult
	if asyncResult.Err != nil {
		callback <- SetFirebaseTokenByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetFirebaseTokenByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetFirebaseTokenByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SetFirebaseTokenByUserIdAsync(
	request *SetFirebaseTokenByUserIdRequest,
	callback chan<- SetFirebaseTokenByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/firebase/token"
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
    if request.Token != nil && *request.Token != "" {
        bodies["token"] = *request.Token
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

	go setFirebaseTokenByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SetFirebaseTokenByUserId(
	request *SetFirebaseTokenByUserIdRequest,
) (*SetFirebaseTokenByUserIdResult, error) {
	callback := make(chan SetFirebaseTokenByUserIdAsyncResult, 1)
	go p.SetFirebaseTokenByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFirebaseTokenAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- GetFirebaseTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFirebaseTokenAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFirebaseTokenResult
	if asyncResult.Err != nil {
		callback <- GetFirebaseTokenAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFirebaseTokenAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFirebaseTokenAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) GetFirebaseTokenAsync(
	request *GetFirebaseTokenRequest,
	callback chan<- GetFirebaseTokenAsyncResult,
) {
	path := "/{namespaceName}/user/me/firebase/token"
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go getFirebaseTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) GetFirebaseToken(
	request *GetFirebaseTokenRequest,
) (*GetFirebaseTokenResult, error) {
	callback := make(chan GetFirebaseTokenAsyncResult, 1)
	go p.GetFirebaseTokenAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getFirebaseTokenByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- GetFirebaseTokenByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFirebaseTokenByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFirebaseTokenByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetFirebaseTokenByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetFirebaseTokenByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetFirebaseTokenByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) GetFirebaseTokenByUserIdAsync(
	request *GetFirebaseTokenByUserIdRequest,
	callback chan<- GetFirebaseTokenByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/firebase/token"
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getFirebaseTokenByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) GetFirebaseTokenByUserId(
	request *GetFirebaseTokenByUserIdRequest,
) (*GetFirebaseTokenByUserIdResult, error) {
	callback := make(chan GetFirebaseTokenByUserIdAsyncResult, 1)
	go p.GetFirebaseTokenByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteFirebaseTokenAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFirebaseTokenAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFirebaseTokenAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFirebaseTokenResult
	if asyncResult.Err != nil {
		callback <- DeleteFirebaseTokenAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteFirebaseTokenAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteFirebaseTokenAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DeleteFirebaseTokenAsync(
	request *DeleteFirebaseTokenRequest,
	callback chan<- DeleteFirebaseTokenAsyncResult,
) {
	path := "/{namespaceName}/user/me/firebase/token"
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go deleteFirebaseTokenAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DeleteFirebaseToken(
	request *DeleteFirebaseTokenRequest,
) (*DeleteFirebaseTokenResult, error) {
	callback := make(chan DeleteFirebaseTokenAsyncResult, 1)
	go p.DeleteFirebaseTokenAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteFirebaseTokenByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteFirebaseTokenByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteFirebaseTokenByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteFirebaseTokenByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteFirebaseTokenByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteFirebaseTokenByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteFirebaseTokenByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) DeleteFirebaseTokenByUserIdAsync(
	request *DeleteFirebaseTokenByUserIdRequest,
	callback chan<- DeleteFirebaseTokenByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/firebase/token"
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go deleteFirebaseTokenByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) DeleteFirebaseTokenByUserId(
	request *DeleteFirebaseTokenByUserIdRequest,
) (*DeleteFirebaseTokenByUserIdResult, error) {
	callback := make(chan DeleteFirebaseTokenByUserIdAsyncResult, 1)
	go p.DeleteFirebaseTokenByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func sendMobileNotificationByUserIdAsyncHandler(
	client Gs2GatewayRestClient,
	job *core.NetworkJob,
	callback chan<- SendMobileNotificationByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SendMobileNotificationByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SendMobileNotificationByUserIdResult
	if asyncResult.Err != nil {
		callback <- SendMobileNotificationByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SendMobileNotificationByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SendMobileNotificationByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2GatewayRestClient) SendMobileNotificationByUserIdAsync(
	request *SendMobileNotificationByUserIdRequest,
	callback chan<- SendMobileNotificationByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/firebase/token/notification"
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
    if request.Subject != nil && *request.Subject != "" {
        bodies["subject"] = *request.Subject
    }
    if request.Payload != nil && *request.Payload != "" {
        bodies["payload"] = *request.Payload
    }
    if request.Sound != nil && *request.Sound != "" {
        bodies["sound"] = *request.Sound
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

	go sendMobileNotificationByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("gateway").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2GatewayRestClient) SendMobileNotificationByUserId(
	request *SendMobileNotificationByUserIdRequest,
) (*SendMobileNotificationByUserIdResult, error) {
	callback := make(chan SendMobileNotificationByUserIdAsyncResult, 1)
	go p.SendMobileNotificationByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
