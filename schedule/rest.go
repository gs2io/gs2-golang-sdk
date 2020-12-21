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

package schedule

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2ScheduleRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2ScheduleRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeNamespaces(
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
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) CreateNamespaceAsync(
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) CreateNamespace(
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
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	path := "/{namespaceName}/status"
    if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetNamespaceStatus(
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
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetNamespace(
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
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) UpdateNamespace(
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
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	path := "/{namespaceName}"
    if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DeleteNamespace(
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

func describeEventMastersAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEventMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEventMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEventMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeEventMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeEventMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeEventMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DescribeEventMastersAsync(
	request *DescribeEventMastersRequest,
	callback chan<- DescribeEventMastersAsyncResult,
) {
	path := "/{namespaceName}/master/event"
    if request.NamespaceName != nil {
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

	go describeEventMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeEventMasters(
	request *DescribeEventMastersRequest,
) (*DescribeEventMastersResult, error) {
	callback := make(chan DescribeEventMastersAsyncResult, 1)
	go p.DescribeEventMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- CreateEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateEventMasterResult
	if asyncResult.Err != nil {
		callback <- CreateEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) CreateEventMasterAsync(
	request *CreateEventMasterRequest,
	callback chan<- CreateEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master/event"
    if request.NamespaceName != nil {
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
    if request.ScheduleType != nil && *request.ScheduleType != "" {
        bodies["scheduleType"] = *request.ScheduleType
    }
    if request.AbsoluteBegin != nil {
        bodies["absoluteBegin"] = *request.AbsoluteBegin
    }
    if request.AbsoluteEnd != nil {
        bodies["absoluteEnd"] = *request.AbsoluteEnd
    }
    if request.RepeatType != nil && *request.RepeatType != "" {
        bodies["repeatType"] = *request.RepeatType
    }
    if request.RepeatBeginDayOfMonth != nil {
        bodies["repeatBeginDayOfMonth"] = *request.RepeatBeginDayOfMonth
    }
    if request.RepeatEndDayOfMonth != nil {
        bodies["repeatEndDayOfMonth"] = *request.RepeatEndDayOfMonth
    }
    if request.RepeatBeginDayOfWeek != nil && *request.RepeatBeginDayOfWeek != "" {
        bodies["repeatBeginDayOfWeek"] = *request.RepeatBeginDayOfWeek
    }
    if request.RepeatEndDayOfWeek != nil && *request.RepeatEndDayOfWeek != "" {
        bodies["repeatEndDayOfWeek"] = *request.RepeatEndDayOfWeek
    }
    if request.RepeatBeginHour != nil {
        bodies["repeatBeginHour"] = *request.RepeatBeginHour
    }
    if request.RepeatEndHour != nil {
        bodies["repeatEndHour"] = *request.RepeatEndHour
    }
    if request.RelativeTriggerName != nil && *request.RelativeTriggerName != "" {
        bodies["relativeTriggerName"] = *request.RelativeTriggerName
    }
    if request.RelativeDuration != nil {
        bodies["relativeDuration"] = *request.RelativeDuration
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) CreateEventMaster(
	request *CreateEventMasterRequest,
) (*CreateEventMasterResult, error) {
	callback := make(chan CreateEventMasterAsyncResult, 1)
	go p.CreateEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEventMasterResult
	if asyncResult.Err != nil {
		callback <- GetEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetEventMasterAsync(
	request *GetEventMasterRequest,
	callback chan<- GetEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetEventMaster(
	request *GetEventMasterRequest,
) (*GetEventMasterResult, error) {
	callback := make(chan GetEventMasterAsyncResult, 1)
	go p.GetEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateEventMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) UpdateEventMasterAsync(
	request *UpdateEventMasterRequest,
	callback chan<- UpdateEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.ScheduleType != nil && *request.ScheduleType != "" {
        bodies["scheduleType"] = *request.ScheduleType
    }
    if request.AbsoluteBegin != nil {
        bodies["absoluteBegin"] = *request.AbsoluteBegin
    }
    if request.AbsoluteEnd != nil {
        bodies["absoluteEnd"] = *request.AbsoluteEnd
    }
    if request.RepeatType != nil && *request.RepeatType != "" {
        bodies["repeatType"] = *request.RepeatType
    }
    if request.RepeatBeginDayOfMonth != nil {
        bodies["repeatBeginDayOfMonth"] = *request.RepeatBeginDayOfMonth
    }
    if request.RepeatEndDayOfMonth != nil {
        bodies["repeatEndDayOfMonth"] = *request.RepeatEndDayOfMonth
    }
    if request.RepeatBeginDayOfWeek != nil && *request.RepeatBeginDayOfWeek != "" {
        bodies["repeatBeginDayOfWeek"] = *request.RepeatBeginDayOfWeek
    }
    if request.RepeatEndDayOfWeek != nil && *request.RepeatEndDayOfWeek != "" {
        bodies["repeatEndDayOfWeek"] = *request.RepeatEndDayOfWeek
    }
    if request.RepeatBeginHour != nil {
        bodies["repeatBeginHour"] = *request.RepeatBeginHour
    }
    if request.RepeatEndHour != nil {
        bodies["repeatEndHour"] = *request.RepeatEndHour
    }
    if request.RelativeTriggerName != nil && *request.RelativeTriggerName != "" {
        bodies["relativeTriggerName"] = *request.RelativeTriggerName
    }
    if request.RelativeDuration != nil {
        bodies["relativeDuration"] = *request.RelativeDuration
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) UpdateEventMaster(
	request *UpdateEventMasterRequest,
) (*UpdateEventMasterResult, error) {
	callback := make(chan UpdateEventMasterAsyncResult, 1)
	go p.UpdateEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteEventMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DeleteEventMasterAsync(
	request *DeleteEventMasterRequest,
	callback chan<- DeleteEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DeleteEventMaster(
	request *DeleteEventMasterRequest,
) (*DeleteEventMasterResult, error) {
	callback := make(chan DeleteEventMasterAsyncResult, 1)
	go p.DeleteEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTriggersAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTriggersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTriggersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTriggersResult
	if asyncResult.Err != nil {
		callback <- DescribeTriggersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeTriggersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeTriggersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DescribeTriggersAsync(
	request *DescribeTriggersRequest,
	callback chan<- DescribeTriggersAsyncResult,
) {
	path := "/{namespaceName}/user/me/trigger"
    if request.NamespaceName != nil {
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

	go describeTriggersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeTriggers(
	request *DescribeTriggersRequest,
) (*DescribeTriggersResult, error) {
	callback := make(chan DescribeTriggersAsyncResult, 1)
	go p.DescribeTriggersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeTriggersByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeTriggersByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeTriggersByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeTriggersByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeTriggersByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeTriggersByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeTriggersByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DescribeTriggersByUserIdAsync(
	request *DescribeTriggersByUserIdRequest,
	callback chan<- DescribeTriggersByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/trigger"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil {
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

	go describeTriggersByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeTriggersByUserId(
	request *DescribeTriggersByUserIdRequest,
) (*DescribeTriggersByUserIdResult, error) {
	callback := make(chan DescribeTriggersByUserIdAsyncResult, 1)
	go p.DescribeTriggersByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTriggerAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetTriggerAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTriggerAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTriggerResult
	if asyncResult.Err != nil {
		callback <- GetTriggerAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetTriggerAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetTriggerAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetTriggerAsync(
	request *GetTriggerRequest,
	callback chan<- GetTriggerAsyncResult,
) {
	path := "/{namespaceName}/user/me/trigger/{triggerName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.TriggerName != nil {
        path = strings.ReplaceAll(path, "{triggerName}", core.ToString(*request.TriggerName))
    } else {
        path = strings.ReplaceAll(path, "{triggerName}", "null")
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

	go getTriggerAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetTrigger(
	request *GetTriggerRequest,
) (*GetTriggerResult, error) {
	callback := make(chan GetTriggerAsyncResult, 1)
	go p.GetTriggerAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getTriggerByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetTriggerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetTriggerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetTriggerByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetTriggerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetTriggerByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetTriggerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetTriggerByUserIdAsync(
	request *GetTriggerByUserIdRequest,
	callback chan<- GetTriggerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/trigger/{triggerName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.TriggerName != nil {
        path = strings.ReplaceAll(path, "{triggerName}", core.ToString(*request.TriggerName))
    } else {
        path = strings.ReplaceAll(path, "{triggerName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getTriggerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetTriggerByUserId(
	request *GetTriggerByUserIdRequest,
) (*GetTriggerByUserIdResult, error) {
	callback := make(chan GetTriggerByUserIdAsyncResult, 1)
	go p.GetTriggerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func triggerByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- TriggerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- TriggerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result TriggerByUserIdResult
	if asyncResult.Err != nil {
		callback <- TriggerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- TriggerByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- TriggerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) TriggerByUserIdAsync(
	request *TriggerByUserIdRequest,
	callback chan<- TriggerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/trigger/{triggerName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.TriggerName != nil {
        path = strings.ReplaceAll(path, "{triggerName}", core.ToString(*request.TriggerName))
    } else {
        path = strings.ReplaceAll(path, "{triggerName}", "null")
    }
    if request.UserId != nil {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.TriggerStrategy != nil && *request.TriggerStrategy != "" {
        bodies["triggerStrategy"] = *request.TriggerStrategy
    }
    if request.Ttl != nil {
        bodies["ttl"] = *request.Ttl
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go triggerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) TriggerByUserId(
	request *TriggerByUserIdRequest,
) (*TriggerByUserIdResult, error) {
	callback := make(chan TriggerByUserIdAsyncResult, 1)
	go p.TriggerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTriggerAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTriggerAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTriggerAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTriggerResult
	if asyncResult.Err != nil {
		callback <- DeleteTriggerAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteTriggerAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteTriggerAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DeleteTriggerAsync(
	request *DeleteTriggerRequest,
	callback chan<- DeleteTriggerAsyncResult,
) {
	path := "/{namespaceName}/user/me/trigger/{triggerName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.TriggerName != nil {
        path = strings.ReplaceAll(path, "{triggerName}", core.ToString(*request.TriggerName))
    } else {
        path = strings.ReplaceAll(path, "{triggerName}", "null")
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

	go deleteTriggerAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DeleteTrigger(
	request *DeleteTriggerRequest,
) (*DeleteTriggerResult, error) {
	callback := make(chan DeleteTriggerAsyncResult, 1)
	go p.DeleteTriggerAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteTriggerByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteTriggerByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteTriggerByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteTriggerByUserIdResult
	if asyncResult.Err != nil {
		callback <- DeleteTriggerByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteTriggerByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteTriggerByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DeleteTriggerByUserIdAsync(
	request *DeleteTriggerByUserIdRequest,
	callback chan<- DeleteTriggerByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/trigger/{triggerName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }
    if request.TriggerName != nil {
        path = strings.ReplaceAll(path, "{triggerName}", core.ToString(*request.TriggerName))
    } else {
        path = strings.ReplaceAll(path, "{triggerName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteTriggerByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DeleteTriggerByUserId(
	request *DeleteTriggerByUserIdRequest,
) (*DeleteTriggerByUserIdResult, error) {
	callback := make(chan DeleteTriggerByUserIdAsyncResult, 1)
	go p.DeleteTriggerByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeEventsAsyncHandler(
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) DescribeEventsAsync(
	request *DescribeEventsRequest,
	callback chan<- DescribeEventsAsyncResult,
) {
	path := "/{namespaceName}/user/me/event"
    if request.NamespaceName != nil {
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

	go describeEventsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeEvents(
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

func describeEventsByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeEventsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeEventsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeEventsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeEventsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeEventsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeEventsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DescribeEventsByUserIdAsync(
	request *DescribeEventsByUserIdRequest,
	callback chan<- DescribeEventsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/event"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.UserId != nil {
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

	go describeEventsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeEventsByUserId(
	request *DescribeEventsByUserIdRequest,
) (*DescribeEventsByUserIdResult, error) {
	callback := make(chan DescribeEventsByUserIdAsyncResult, 1)
	go p.DescribeEventsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRawEventsAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRawEventsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRawEventsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRawEventsResult
	if asyncResult.Err != nil {
		callback <- DescribeRawEventsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRawEventsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRawEventsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) DescribeRawEventsAsync(
	request *DescribeRawEventsRequest,
	callback chan<- DescribeRawEventsAsyncResult,
) {
	path := "/{namespaceName}/event"
    if request.NamespaceName != nil {
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

	go describeRawEventsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) DescribeRawEvents(
	request *DescribeRawEventsRequest,
) (*DescribeRawEventsResult, error) {
	callback := make(chan DescribeRawEventsAsyncResult, 1)
	go p.DescribeRawEventsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getEventAsyncHandler(
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) GetEventAsync(
	request *GetEventRequest,
	callback chan<- GetEventAsyncResult,
) {
	path := "/{namespaceName}/user/me/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
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

	go getEventAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetEvent(
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

func getEventByUserIdAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetEventByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetEventByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetEventByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetEventByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetEventByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetEventByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetEventByUserIdAsync(
	request *GetEventByUserIdRequest,
	callback chan<- GetEventByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
    }
    if request.UserId != nil {
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

	go getEventByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetEventByUserId(
	request *GetEventByUserIdRequest,
) (*GetEventByUserIdResult, error) {
	callback := make(chan GetEventByUserIdAsyncResult, 1)
	go p.GetEventByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRawEventAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetRawEventAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRawEventAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRawEventResult
	if asyncResult.Err != nil {
		callback <- GetRawEventAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRawEventAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRawEventAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetRawEventAsync(
	request *GetRawEventRequest,
	callback chan<- GetRawEventAsyncResult,
) {
	path := "/{namespaceName}/event/{eventName}"
    if request.NamespaceName != nil {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.EventName != nil {
        path = strings.ReplaceAll(path, "{eventName}", core.ToString(*request.EventName))
    } else {
        path = strings.ReplaceAll(path, "{eventName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRawEventAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetRawEvent(
	request *GetRawEventRequest,
) (*GetRawEventResult, error) {
	callback := make(chan GetRawEventAsyncResult, 1)
	go p.GetRawEventAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2ScheduleRestClient,
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

func (p Gs2ScheduleRestClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
	path := "/{namespaceName}/master/export"
    if request.NamespaceName != nil {
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
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) ExportMaster(
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

func getCurrentEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentEventMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) GetCurrentEventMasterAsync(
	request *GetCurrentEventMasterRequest,
	callback chan<- GetCurrentEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
    if request.NamespaceName != nil {
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

	go getCurrentEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) GetCurrentEventMaster(
	request *GetCurrentEventMasterRequest,
) (*GetCurrentEventMasterResult, error) {
	callback := make(chan GetCurrentEventMasterAsyncResult, 1)
	go p.GetCurrentEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentEventMasterAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentEventMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentEventMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentEventMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentEventMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentEventMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentEventMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) UpdateCurrentEventMasterAsync(
	request *UpdateCurrentEventMasterRequest,
	callback chan<- UpdateCurrentEventMasterAsyncResult,
) {
	path := "/{namespaceName}/master"
    if request.NamespaceName != nil {
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

	go updateCurrentEventMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) UpdateCurrentEventMaster(
	request *UpdateCurrentEventMasterRequest,
) (*UpdateCurrentEventMasterResult, error) {
	callback := make(chan UpdateCurrentEventMasterAsyncResult, 1)
	go p.UpdateCurrentEventMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentEventMasterFromGitHubAsyncHandler(
	client Gs2ScheduleRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentEventMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentEventMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentEventMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentEventMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentEventMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentEventMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ScheduleRestClient) UpdateCurrentEventMasterFromGitHubAsync(
	request *UpdateCurrentEventMasterFromGitHubRequest,
	callback chan<- UpdateCurrentEventMasterFromGitHubAsyncResult,
) {
	path := "/{namespaceName}/master/from_git_hub"
    if request.NamespaceName != nil {
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

	go updateCurrentEventMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("schedule").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ScheduleRestClient) UpdateCurrentEventMasterFromGitHub(
	request *UpdateCurrentEventMasterFromGitHubRequest,
) (*UpdateCurrentEventMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentEventMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentEventMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
