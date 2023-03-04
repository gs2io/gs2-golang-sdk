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

package matchmaking

import (
	"encoding/json"
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
)

type Gs2MatchmakingRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2MatchmakingRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func describeNamespacesAsyncHandler(
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) DescribeNamespacesAsync(
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeNamespaces(
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
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) CreateNamespaceAsync(
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
    if request.EnableRating != nil {
        bodies["enableRating"] = *request.EnableRating
    }
    if request.CreateGatheringTriggerType != nil && *request.CreateGatheringTriggerType != "" {
        bodies["createGatheringTriggerType"] = *request.CreateGatheringTriggerType
    }
    if request.CreateGatheringTriggerRealtimeNamespaceId != nil && *request.CreateGatheringTriggerRealtimeNamespaceId != "" {
        bodies["createGatheringTriggerRealtimeNamespaceId"] = *request.CreateGatheringTriggerRealtimeNamespaceId
    }
    if request.CreateGatheringTriggerScriptId != nil && *request.CreateGatheringTriggerScriptId != "" {
        bodies["createGatheringTriggerScriptId"] = *request.CreateGatheringTriggerScriptId
    }
    if request.CompleteMatchmakingTriggerType != nil && *request.CompleteMatchmakingTriggerType != "" {
        bodies["completeMatchmakingTriggerType"] = *request.CompleteMatchmakingTriggerType
    }
    if request.CompleteMatchmakingTriggerRealtimeNamespaceId != nil && *request.CompleteMatchmakingTriggerRealtimeNamespaceId != "" {
        bodies["completeMatchmakingTriggerRealtimeNamespaceId"] = *request.CompleteMatchmakingTriggerRealtimeNamespaceId
    }
    if request.CompleteMatchmakingTriggerScriptId != nil && *request.CompleteMatchmakingTriggerScriptId != "" {
        bodies["completeMatchmakingTriggerScriptId"] = *request.CompleteMatchmakingTriggerScriptId
    }
    if request.ChangeRatingScript != nil {
        bodies["changeRatingScript"] = request.ChangeRatingScript.ToDict()
    }
    if request.JoinNotification != nil {
        bodies["joinNotification"] = request.JoinNotification.ToDict()
    }
    if request.LeaveNotification != nil {
        bodies["leaveNotification"] = request.LeaveNotification.ToDict()
    }
    if request.CompleteNotification != nil {
        bodies["completeNotification"] = request.CompleteNotification.ToDict()
    }
    if request.ChangeRatingNotification != nil {
        bodies["changeRatingNotification"] = request.ChangeRatingNotification.ToDict()
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CreateNamespace(
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
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) GetNamespaceStatusAsync(
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetNamespaceStatus(
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
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) GetNamespaceAsync(
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetNamespace(
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
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) UpdateNamespaceAsync(
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
    if request.EnableRating != nil {
        bodies["enableRating"] = *request.EnableRating
    }
    if request.CreateGatheringTriggerType != nil && *request.CreateGatheringTriggerType != "" {
        bodies["createGatheringTriggerType"] = *request.CreateGatheringTriggerType
    }
    if request.CreateGatheringTriggerRealtimeNamespaceId != nil && *request.CreateGatheringTriggerRealtimeNamespaceId != "" {
        bodies["createGatheringTriggerRealtimeNamespaceId"] = *request.CreateGatheringTriggerRealtimeNamespaceId
    }
    if request.CreateGatheringTriggerScriptId != nil && *request.CreateGatheringTriggerScriptId != "" {
        bodies["createGatheringTriggerScriptId"] = *request.CreateGatheringTriggerScriptId
    }
    if request.CompleteMatchmakingTriggerType != nil && *request.CompleteMatchmakingTriggerType != "" {
        bodies["completeMatchmakingTriggerType"] = *request.CompleteMatchmakingTriggerType
    }
    if request.CompleteMatchmakingTriggerRealtimeNamespaceId != nil && *request.CompleteMatchmakingTriggerRealtimeNamespaceId != "" {
        bodies["completeMatchmakingTriggerRealtimeNamespaceId"] = *request.CompleteMatchmakingTriggerRealtimeNamespaceId
    }
    if request.CompleteMatchmakingTriggerScriptId != nil && *request.CompleteMatchmakingTriggerScriptId != "" {
        bodies["completeMatchmakingTriggerScriptId"] = *request.CompleteMatchmakingTriggerScriptId
    }
    if request.ChangeRatingScript != nil {
        bodies["changeRatingScript"] = request.ChangeRatingScript.ToDict()
    }
    if request.JoinNotification != nil {
        bodies["joinNotification"] = request.JoinNotification.ToDict()
    }
    if request.LeaveNotification != nil {
        bodies["leaveNotification"] = request.LeaveNotification.ToDict()
    }
    if request.CompleteNotification != nil {
        bodies["completeNotification"] = request.CompleteNotification.ToDict()
    }
    if request.ChangeRatingNotification != nil {
        bodies["changeRatingNotification"] = request.ChangeRatingNotification.ToDict()
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateNamespace(
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
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) DeleteNamespaceAsync(
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DeleteNamespace(
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

func describeGatheringsAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeGatheringsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeGatheringsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeGatheringsResult
	if asyncResult.Err != nil {
		callback <- DescribeGatheringsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeGatheringsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeGatheringsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DescribeGatheringsAsync(
	request *DescribeGatheringsRequest,
	callback chan<- DescribeGatheringsAsyncResult,
) {
	path := "/{namespaceName}/gathering"
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

	go describeGatheringsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeGatherings(
	request *DescribeGatheringsRequest,
) (*DescribeGatheringsResult, error) {
	callback := make(chan DescribeGatheringsAsyncResult, 1)
	go p.DescribeGatheringsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGatheringAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGatheringResult
	if asyncResult.Err != nil {
		callback <- CreateGatheringAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateGatheringAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CreateGatheringAsync(
	request *CreateGatheringRequest,
	callback chan<- CreateGatheringAsyncResult,
) {
	path := "/{namespaceName}/gathering"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Player != nil {
        bodies["player"] = request.Player.ToDict()
    }
    if request.AttributeRanges != nil {
        var _attributeRanges []interface {}
        for _, item := range request.AttributeRanges {
            _attributeRanges = append(_attributeRanges, item)
        }
        bodies["attributeRanges"] = _attributeRanges
    }
    if request.CapacityOfRoles != nil {
        var _capacityOfRoles []interface {}
        for _, item := range request.CapacityOfRoles {
            _capacityOfRoles = append(_capacityOfRoles, item)
        }
        bodies["capacityOfRoles"] = _capacityOfRoles
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []interface {}
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
    if request.ExpiresAtTimeSpan != nil {
        bodies["expiresAtTimeSpan"] = request.ExpiresAtTimeSpan.ToDict()
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
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go createGatheringAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CreateGathering(
	request *CreateGatheringRequest,
) (*CreateGatheringResult, error) {
	callback := make(chan CreateGatheringAsyncResult, 1)
	go p.CreateGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createGatheringByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CreateGatheringByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateGatheringByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateGatheringByUserIdResult
	if asyncResult.Err != nil {
		callback <- CreateGatheringByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateGatheringByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateGatheringByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CreateGatheringByUserIdAsync(
	request *CreateGatheringByUserIdRequest,
	callback chan<- CreateGatheringByUserIdAsyncResult,
) {
	path := "/{namespaceName}/gathering/user/{userId}"
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
    if request.Player != nil {
        bodies["player"] = request.Player.ToDict()
    }
    if request.AttributeRanges != nil {
        var _attributeRanges []interface {}
        for _, item := range request.AttributeRanges {
            _attributeRanges = append(_attributeRanges, item)
        }
        bodies["attributeRanges"] = _attributeRanges
    }
    if request.CapacityOfRoles != nil {
        var _capacityOfRoles []interface {}
        for _, item := range request.CapacityOfRoles {
            _capacityOfRoles = append(_capacityOfRoles, item)
        }
        bodies["capacityOfRoles"] = _capacityOfRoles
    }
    if request.AllowUserIds != nil {
        var _allowUserIds []interface {}
        for _, item := range request.AllowUserIds {
            _allowUserIds = append(_allowUserIds, item)
        }
        bodies["allowUserIds"] = _allowUserIds
    }
    if request.ExpiresAt != nil {
        bodies["expiresAt"] = *request.ExpiresAt
    }
    if request.ExpiresAtTimeSpan != nil {
        bodies["expiresAtTimeSpan"] = request.ExpiresAtTimeSpan.ToDict()
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

	go createGatheringByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CreateGatheringByUserId(
	request *CreateGatheringByUserIdRequest,
) (*CreateGatheringByUserIdResult, error) {
	callback := make(chan CreateGatheringByUserIdAsyncResult, 1)
	go p.CreateGatheringByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGatheringAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGatheringResult
	if asyncResult.Err != nil {
		callback <- UpdateGatheringAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateGatheringAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) UpdateGatheringAsync(
	request *UpdateGatheringRequest,
	callback chan<- UpdateGatheringAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AttributeRanges != nil {
        var _attributeRanges []interface {}
        for _, item := range request.AttributeRanges {
            _attributeRanges = append(_attributeRanges, item)
        }
        bodies["attributeRanges"] = _attributeRanges
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
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go updateGatheringAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateGathering(
	request *UpdateGatheringRequest,
) (*UpdateGatheringResult, error) {
	callback := make(chan UpdateGatheringAsyncResult, 1)
	go p.UpdateGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateGatheringByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateGatheringByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateGatheringByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateGatheringByUserIdResult
	if asyncResult.Err != nil {
		callback <- UpdateGatheringByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateGatheringByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateGatheringByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) UpdateGatheringByUserIdAsync(
	request *UpdateGatheringByUserIdRequest,
	callback chan<- UpdateGatheringByUserIdAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}/user/{userId}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.AttributeRanges != nil {
        var _attributeRanges []interface {}
        for _, item := range request.AttributeRanges {
            _attributeRanges = append(_attributeRanges, item)
        }
        bodies["attributeRanges"] = _attributeRanges
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

	go updateGatheringByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateGatheringByUserId(
	request *UpdateGatheringByUserIdRequest,
) (*UpdateGatheringByUserIdResult, error) {
	callback := make(chan UpdateGatheringByUserIdAsyncResult, 1)
	go p.UpdateGatheringByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doMatchmakingByPlayerAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DoMatchmakingByPlayerAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingByPlayerAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingByPlayerResult
	if asyncResult.Err != nil {
		callback <- DoMatchmakingByPlayerAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoMatchmakingByPlayerAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DoMatchmakingByPlayerAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DoMatchmakingByPlayerAsync(
	request *DoMatchmakingByPlayerRequest,
	callback chan<- DoMatchmakingByPlayerAsyncResult,
) {
	path := "/{namespaceName}/gathering/player/do"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Player != nil {
        bodies["player"] = request.Player.ToDict()
    }
    if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
        bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go doMatchmakingByPlayerAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DoMatchmakingByPlayer(
	request *DoMatchmakingByPlayerRequest,
) (*DoMatchmakingByPlayerResult, error) {
	callback := make(chan DoMatchmakingByPlayerAsyncResult, 1)
	go p.DoMatchmakingByPlayerAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doMatchmakingAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DoMatchmakingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingResult
	if asyncResult.Err != nil {
		callback <- DoMatchmakingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoMatchmakingAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DoMatchmakingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DoMatchmakingAsync(
	request *DoMatchmakingRequest,
	callback chan<- DoMatchmakingAsyncResult,
) {
	path := "/{namespaceName}/gathering/do"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Player != nil {
        bodies["player"] = request.Player.ToDict()
    }
    if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
        bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
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
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go doMatchmakingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DoMatchmaking(
	request *DoMatchmakingRequest,
) (*DoMatchmakingResult, error) {
	callback := make(chan DoMatchmakingAsyncResult, 1)
	go p.DoMatchmakingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func doMatchmakingByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DoMatchmakingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DoMatchmakingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DoMatchmakingByUserIdResult
	if asyncResult.Err != nil {
		callback <- DoMatchmakingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DoMatchmakingByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DoMatchmakingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DoMatchmakingByUserIdAsync(
	request *DoMatchmakingByUserIdRequest,
	callback chan<- DoMatchmakingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/gathering/do"
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
    if request.Player != nil {
        bodies["player"] = request.Player.ToDict()
    }
    if request.MatchmakingContextToken != nil && *request.MatchmakingContextToken != "" {
        bodies["matchmakingContextToken"] = *request.MatchmakingContextToken
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

	go doMatchmakingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DoMatchmakingByUserId(
	request *DoMatchmakingByUserIdRequest,
) (*DoMatchmakingByUserIdResult, error) {
	callback := make(chan DoMatchmakingByUserIdAsyncResult, 1)
	go p.DoMatchmakingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getGatheringAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGatheringResult
	if asyncResult.Err != nil {
		callback <- GetGatheringAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetGatheringAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetGatheringAsync(
	request *GetGatheringRequest,
	callback chan<- GetGatheringAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getGatheringAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetGathering(
	request *GetGatheringRequest,
) (*GetGatheringResult, error) {
	callback := make(chan GetGatheringAsyncResult, 1)
	go p.GetGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func cancelMatchmakingAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CancelMatchmakingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CancelMatchmakingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CancelMatchmakingResult
	if asyncResult.Err != nil {
		callback <- CancelMatchmakingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CancelMatchmakingAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CancelMatchmakingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CancelMatchmakingAsync(
	request *CancelMatchmakingRequest,
	callback chan<- CancelMatchmakingAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}/user/me"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
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
    if request.DuplicationAvoider != nil {
      headers["X-GS2-DUPLICATION-AVOIDER"] = string(*request.DuplicationAvoider)
    }

	go cancelMatchmakingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CancelMatchmaking(
	request *CancelMatchmakingRequest,
) (*CancelMatchmakingResult, error) {
	callback := make(chan CancelMatchmakingAsyncResult, 1)
	go p.CancelMatchmakingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func cancelMatchmakingByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CancelMatchmakingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CancelMatchmakingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CancelMatchmakingByUserIdResult
	if asyncResult.Err != nil {
		callback <- CancelMatchmakingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CancelMatchmakingByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CancelMatchmakingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CancelMatchmakingByUserIdAsync(
	request *CancelMatchmakingByUserIdRequest,
	callback chan<- CancelMatchmakingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}/user/{userId}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
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

	go cancelMatchmakingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CancelMatchmakingByUserId(
	request *CancelMatchmakingByUserIdRequest,
) (*CancelMatchmakingByUserIdResult, error) {
	callback := make(chan CancelMatchmakingByUserIdAsyncResult, 1)
	go p.CancelMatchmakingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteGatheringAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteGatheringAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteGatheringAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteGatheringResult
	if asyncResult.Err != nil {
		callback <- DeleteGatheringAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteGatheringAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteGatheringAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DeleteGatheringAsync(
	request *DeleteGatheringRequest,
	callback chan<- DeleteGatheringAsyncResult,
) {
	path := "/{namespaceName}/gathering/{gatheringName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteGatheringAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DeleteGathering(
	request *DeleteGatheringRequest,
) (*DeleteGatheringResult, error) {
	callback := make(chan DeleteGatheringAsyncResult, 1)
	go p.DeleteGatheringAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRatingModelMastersAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRatingModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingModelMastersResult
	if asyncResult.Err != nil {
		callback <- DescribeRatingModelMastersAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRatingModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRatingModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DescribeRatingModelMastersAsync(
	request *DescribeRatingModelMastersRequest,
	callback chan<- DescribeRatingModelMastersAsyncResult,
) {
	path := "/{namespaceName}/master/rating"
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

	go describeRatingModelMastersAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeRatingModelMasters(
	request *DescribeRatingModelMastersRequest,
) (*DescribeRatingModelMastersResult, error) {
	callback := make(chan DescribeRatingModelMastersAsyncResult, 1)
	go p.DescribeRatingModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func createRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CreateRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- CreateRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CreateRatingModelMasterAsync(
	request *CreateRatingModelMasterRequest,
	callback chan<- CreateRatingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/rating"
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
    if request.Volatility != nil {
        bodies["volatility"] = *request.Volatility
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go createRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CreateRatingModelMaster(
	request *CreateRatingModelMasterRequest,
) (*CreateRatingModelMasterResult, error) {
	callback := make(chan CreateRatingModelMasterAsyncResult, 1)
	go p.CreateRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetRatingModelMasterAsync(
	request *GetRatingModelMasterRequest,
	callback chan<- GetRatingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/rating/{ratingName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetRatingModelMaster(
	request *GetRatingModelMasterRequest,
) (*GetRatingModelMasterResult, error) {
	callback := make(chan GetRatingModelMasterAsyncResult, 1)
	go p.GetRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) UpdateRatingModelMasterAsync(
	request *UpdateRatingModelMasterRequest,
	callback chan<- UpdateRatingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/rating/{ratingName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.Volatility != nil {
        bodies["volatility"] = *request.Volatility
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go updateRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateRatingModelMaster(
	request *UpdateRatingModelMasterRequest,
) (*UpdateRatingModelMasterResult, error) {
	callback := make(chan UpdateRatingModelMasterAsyncResult, 1)
	go p.UpdateRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- DeleteRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DeleteRatingModelMasterAsync(
	request *DeleteRatingModelMasterRequest,
	callback chan<- DeleteRatingModelMasterAsyncResult,
) {
	path := "/{namespaceName}/master/rating/{ratingName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go deleteRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DeleteRatingModelMaster(
	request *DeleteRatingModelMasterRequest,
) (*DeleteRatingModelMasterResult, error) {
	callback := make(chan DeleteRatingModelMasterAsyncResult, 1)
	go p.DeleteRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRatingModelsAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRatingModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingModelsResult
	if asyncResult.Err != nil {
		callback <- DescribeRatingModelsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRatingModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRatingModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DescribeRatingModelsAsync(
	request *DescribeRatingModelsRequest,
	callback chan<- DescribeRatingModelsAsyncResult,
) {
	path := "/{namespaceName}/rating"
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

	go describeRatingModelsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeRatingModels(
	request *DescribeRatingModelsRequest,
) (*DescribeRatingModelsResult, error) {
	callback := make(chan DescribeRatingModelsAsyncResult, 1)
	go p.DescribeRatingModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRatingModelAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRatingModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingModelResult
	if asyncResult.Err != nil {
		callback <- GetRatingModelAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRatingModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRatingModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetRatingModelAsync(
	request *GetRatingModelRequest,
	callback chan<- GetRatingModelAsyncResult,
) {
	path := "/{namespaceName}/rating/{ratingName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRatingModelAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetRatingModel(
	request *GetRatingModelRequest,
) (*GetRatingModelResult, error) {
	callback := make(chan GetRatingModelAsyncResult, 1)
	go p.GetRatingModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func exportMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
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

func (p Gs2MatchmakingRestClient) ExportMasterAsync(
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
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) ExportMaster(
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

func getCurrentRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetCurrentRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- GetCurrentRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetCurrentRatingModelMasterAsync(
	request *GetCurrentRatingModelMasterRequest,
	callback chan<- GetCurrentRatingModelMasterAsyncResult,
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

	go getCurrentRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetCurrentRatingModelMaster(
	request *GetCurrentRatingModelMasterRequest,
) (*GetCurrentRatingModelMasterResult, error) {
	callback := make(chan GetCurrentRatingModelMasterAsyncResult, 1)
	go p.GetCurrentRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRatingModelMasterAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRatingModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRatingModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRatingModelMasterResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRatingModelMasterAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRatingModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRatingModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) UpdateCurrentRatingModelMasterAsync(
	request *UpdateCurrentRatingModelMasterRequest,
	callback chan<- UpdateCurrentRatingModelMasterAsyncResult,
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

	go updateCurrentRatingModelMasterAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateCurrentRatingModelMaster(
	request *UpdateCurrentRatingModelMasterRequest,
) (*UpdateCurrentRatingModelMasterResult, error) {
	callback := make(chan UpdateCurrentRatingModelMasterAsyncResult, 1)
	go p.UpdateCurrentRatingModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func updateCurrentRatingModelMasterFromGitHubAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- UpdateCurrentRatingModelMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRatingModelMasterFromGitHubResult
	if asyncResult.Err != nil {
		callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRatingModelMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) UpdateCurrentRatingModelMasterFromGitHubAsync(
	request *UpdateCurrentRatingModelMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRatingModelMasterFromGitHubAsyncResult,
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

	go updateCurrentRatingModelMasterFromGitHubAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Put,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) UpdateCurrentRatingModelMasterFromGitHub(
	request *UpdateCurrentRatingModelMasterFromGitHubRequest,
) (*UpdateCurrentRatingModelMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRatingModelMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRatingModelMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRatingsAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRatingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingsResult
	if asyncResult.Err != nil {
		callback <- DescribeRatingsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRatingsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRatingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DescribeRatingsAsync(
	request *DescribeRatingsRequest,
	callback chan<- DescribeRatingsAsyncResult,
) {
	path := "/{namespaceName}/user/me/rating"
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

	go describeRatingsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeRatings(
	request *DescribeRatingsRequest,
) (*DescribeRatingsResult, error) {
	callback := make(chan DescribeRatingsAsyncResult, 1)
	go p.DescribeRatingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func describeRatingsByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRatingsByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRatingsByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRatingsByUserIdResult
	if asyncResult.Err != nil {
		callback <- DescribeRatingsByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRatingsByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRatingsByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DescribeRatingsByUserIdAsync(
	request *DescribeRatingsByUserIdRequest,
	callback chan<- DescribeRatingsByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/rating"
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

	go describeRatingsByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DescribeRatingsByUserId(
	request *DescribeRatingsByUserIdRequest,
) (*DescribeRatingsByUserIdResult, error) {
	callback := make(chan DescribeRatingsByUserIdAsyncResult, 1)
	go p.DescribeRatingsByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRatingAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRatingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingResult
	if asyncResult.Err != nil {
		callback <- GetRatingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRatingAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRatingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetRatingAsync(
	request *GetRatingRequest,
	callback chan<- GetRatingAsyncResult,
) {
	path := "/{namespaceName}/user/me/rating/{ratingName}"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
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

	go getRatingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetRating(
	request *GetRatingRequest,
) (*GetRatingResult, error) {
	callback := make(chan GetRatingAsyncResult, 1)
	go p.GetRatingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getRatingByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetRatingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRatingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRatingByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetRatingByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRatingByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRatingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetRatingByUserIdAsync(
	request *GetRatingByUserIdRequest,
	callback chan<- GetRatingByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/rating/{ratingName}"
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
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getRatingByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetRatingByUserId(
	request *GetRatingByUserIdRequest,
) (*GetRatingByUserIdResult, error) {
	callback := make(chan GetRatingByUserIdAsyncResult, 1)
	go p.GetRatingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func putResultAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- PutResultAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutResultAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutResultResult
	if asyncResult.Err != nil {
		callback <- PutResultAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PutResultAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PutResultAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) PutResultAsync(
	request *PutResultRequest,
	callback chan<- PutResultAsyncResult,
) {
	path := "/{namespaceName}/rating/{ratingName}/vote"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.GameResults != nil {
        var _gameResults []interface {}
        for _, item := range request.GameResults {
            _gameResults = append(_gameResults, item)
        }
        bodies["gameResults"] = _gameResults
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go putResultAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) PutResult(
	request *PutResultRequest,
) (*PutResultResult, error) {
	callback := make(chan PutResultAsyncResult, 1)
	go p.PutResultAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func deleteRatingAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- DeleteRatingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteRatingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteRatingResult
	if asyncResult.Err != nil {
		callback <- DeleteRatingAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteRatingAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteRatingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) DeleteRatingAsync(
	request *DeleteRatingRequest,
	callback chan<- DeleteRatingAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/rating/{ratingName}"
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
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
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

	go deleteRatingAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Delete,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) DeleteRating(
	request *DeleteRatingRequest,
) (*DeleteRatingResult, error) {
	callback := make(chan DeleteRatingAsyncResult, 1)
	go p.DeleteRatingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBallotAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetBallotAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBallotAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBallotResult
	if asyncResult.Err != nil {
		callback <- GetBallotAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBallotAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetBallotAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetBallotAsync(
	request *GetBallotRequest,
	callback chan<- GetBallotAsyncResult,
) {
	path := "/{namespaceName}/user/me/vote/{ratingName}/{gatheringName}/ballot"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.NumberOfPlayer != nil {
        bodies["numberOfPlayer"] = *request.NumberOfPlayer
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
    if request.AccessToken != nil {
        headers["X-GS2-ACCESS-TOKEN"] = string(*request.AccessToken)
    }

	go getBallotAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetBallot(
	request *GetBallotRequest,
) (*GetBallotResult, error) {
	callback := make(chan GetBallotAsyncResult, 1)
	go p.GetBallotAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func getBallotByUserIdAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- GetBallotByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBallotByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBallotByUserIdResult
	if asyncResult.Err != nil {
		callback <- GetBallotByUserIdAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBallotByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetBallotByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) GetBallotByUserIdAsync(
	request *GetBallotByUserIdRequest,
	callback chan<- GetBallotByUserIdAsyncResult,
) {
	path := "/{namespaceName}/user/{userId}/vote/{ratingName}/{gatheringName}/ballot"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
    }
    if request.UserId != nil && *request.UserId != ""  {
        path = strings.ReplaceAll(path, "{userId}", core.ToString(*request.UserId))
    } else {
        path = strings.ReplaceAll(path, "{userId}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.NumberOfPlayer != nil {
        bodies["numberOfPlayer"] = *request.NumberOfPlayer
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

	go getBallotByUserIdAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) GetBallotByUserId(
	request *GetBallotByUserIdRequest,
) (*GetBallotByUserIdResult, error) {
	callback := make(chan GetBallotByUserIdAsyncResult, 1)
	go p.GetBallotByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func voteAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- VoteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VoteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VoteResult
	if asyncResult.Err != nil {
		callback <- VoteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VoteAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- VoteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) VoteAsync(
	request *VoteRequest,
	callback chan<- VoteAsyncResult,
) {
	path := "/{namespaceName}/action/vote"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.BallotBody != nil && *request.BallotBody != "" {
        bodies["ballotBody"] = *request.BallotBody
    }
    if request.BallotSignature != nil && *request.BallotSignature != "" {
        bodies["ballotSignature"] = *request.BallotSignature
    }
    if request.GameResults != nil {
        var _gameResults []interface {}
        for _, item := range request.GameResults {
            _gameResults = append(_gameResults, item)
        }
        bodies["gameResults"] = _gameResults
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

	go voteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) Vote(
	request *VoteRequest,
) (*VoteResult, error) {
	callback := make(chan VoteAsyncResult, 1)
	go p.VoteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func voteMultipleAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- VoteMultipleAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- VoteMultipleAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result VoteMultipleResult
	if asyncResult.Err != nil {
		callback <- VoteMultipleAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- VoteMultipleAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- VoteMultipleAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) VoteMultipleAsync(
	request *VoteMultipleRequest,
	callback chan<- VoteMultipleAsyncResult,
) {
	path := "/{namespaceName}/action/vote/multiple"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.SignedBallots != nil {
        var _signedBallots []interface {}
        for _, item := range request.SignedBallots {
            _signedBallots = append(_signedBallots, item)
        }
        bodies["signedBallots"] = _signedBallots
    }
    if request.GameResults != nil {
        var _gameResults []interface {}
        for _, item := range request.GameResults {
            _gameResults = append(_gameResults, item)
        }
        bodies["gameResults"] = _gameResults
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

	go voteMultipleAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) VoteMultiple(
	request *VoteMultipleRequest,
) (*VoteMultipleResult, error) {
	callback := make(chan VoteMultipleAsyncResult, 1)
	go p.VoteMultipleAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func commitVoteAsyncHandler(
	client Gs2MatchmakingRestClient,
	job *core.NetworkJob,
	callback chan<- CommitVoteAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CommitVoteAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CommitVoteResult
	if asyncResult.Err != nil {
		callback <- CommitVoteAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CommitVoteAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CommitVoteAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2MatchmakingRestClient) CommitVoteAsync(
	request *CommitVoteRequest,
	callback chan<- CommitVoteAsyncResult,
) {
	path := "/{namespaceName}/vote/{ratingName}/{gatheringName}/action/vote/commit"
    if request.NamespaceName != nil && *request.NamespaceName != ""  {
        path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
    } else {
        path = strings.ReplaceAll(path, "{namespaceName}", "null")
    }
    if request.RatingName != nil && *request.RatingName != ""  {
        path = strings.ReplaceAll(path, "{ratingName}", core.ToString(*request.RatingName))
    } else {
        path = strings.ReplaceAll(path, "{ratingName}", "null")
    }
    if request.GatheringName != nil && *request.GatheringName != ""  {
        path = strings.ReplaceAll(path, "{gatheringName}", core.ToString(*request.GatheringName))
    } else {
        path = strings.ReplaceAll(path, "{gatheringName}", "null")
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

	go commitVoteAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("matchmaking").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2MatchmakingRestClient) CommitVote(
	request *CommitVoteRequest,
) (*CommitVoteResult, error) {
	callback := make(chan CommitVoteAsyncResult, 1)
	go p.CommitVoteAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
