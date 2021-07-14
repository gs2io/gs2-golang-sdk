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

package ranking

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2RankingWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2RankingWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2RankingWebSocketClient) describeNamespacesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeNamespacesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "describeNamespaces",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeNamespacesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeNamespaces(
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

func (p Gs2RankingWebSocketClient) createNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "createNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
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

	go p.createNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) CreateNamespace(
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

func (p Gs2RankingWebSocketClient) getNamespaceStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "getNamespaceStatus",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetNamespaceStatus(
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

func (p Gs2RankingWebSocketClient) getNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "getNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetNamespace(
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

func (p Gs2RankingWebSocketClient) updateNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "updateNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
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

	go p.updateNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) UpdateNamespace(
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

func (p Gs2RankingWebSocketClient) deleteNamespaceAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteNamespaceAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "namespace",
    		"function": "deleteNamespace",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteNamespaceAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DeleteNamespace(
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

func (p Gs2RankingWebSocketClient) describeCategoryModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCategoryModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCategoryModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCategoryModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCategoryModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCategoryModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeCategoryModelsAsync(
	request *DescribeCategoryModelsRequest,
	callback chan<- DescribeCategoryModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModel",
    		"function": "describeCategoryModels",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeCategoryModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeCategoryModels(
	request *DescribeCategoryModelsRequest,
) (*DescribeCategoryModelsResult, error) {
	callback := make(chan DescribeCategoryModelsAsyncResult, 1)
	go p.DescribeCategoryModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getCategoryModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCategoryModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCategoryModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCategoryModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCategoryModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCategoryModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetCategoryModelAsync(
	request *GetCategoryModelRequest,
	callback chan<- GetCategoryModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModel",
    		"function": "getCategoryModel",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCategoryModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetCategoryModel(
	request *GetCategoryModelRequest,
) (*GetCategoryModelResult, error) {
	callback := make(chan GetCategoryModelAsyncResult, 1)
	go p.GetCategoryModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeCategoryModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeCategoryModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeCategoryModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeCategoryModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeCategoryModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeCategoryModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeCategoryModelMastersAsync(
	request *DescribeCategoryModelMastersRequest,
	callback chan<- DescribeCategoryModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModelMaster",
    		"function": "describeCategoryModelMasters",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeCategoryModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeCategoryModelMasters(
	request *DescribeCategoryModelMastersRequest,
) (*DescribeCategoryModelMastersResult, error) {
	callback := make(chan DescribeCategoryModelMastersAsyncResult, 1)
	go p.DescribeCategoryModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) createCategoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateCategoryModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateCategoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) CreateCategoryModelMasterAsync(
	request *CreateCategoryModelMasterRequest,
	callback chan<- CreateCategoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModelMaster",
    		"function": "createCategoryModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.Name != nil && *request.Name != "" {
        bodies["name"] = *request.Name
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.MinimumValue != nil {
        bodies["minimumValue"] = *request.MinimumValue
    }
    if request.MaximumValue != nil {
        bodies["maximumValue"] = *request.MaximumValue
    }
    if request.OrderDirection != nil && *request.OrderDirection != "" {
        bodies["orderDirection"] = *request.OrderDirection
    }
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.UniqueByUserId != nil {
        bodies["uniqueByUserId"] = *request.UniqueByUserId
    }
    if request.CalculateFixedTimingHour != nil {
        bodies["calculateFixedTimingHour"] = *request.CalculateFixedTimingHour
    }
    if request.CalculateFixedTimingMinute != nil {
        bodies["calculateFixedTimingMinute"] = *request.CalculateFixedTimingMinute
    }
    if request.CalculateIntervalMinutes != nil {
        bodies["calculateIntervalMinutes"] = *request.CalculateIntervalMinutes
    }
    if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
        bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
    }
    if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
        bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
    }
    if request.Generation != nil && *request.Generation != "" {
        bodies["generation"] = *request.Generation
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createCategoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) CreateCategoryModelMaster(
	request *CreateCategoryModelMasterRequest,
) (*CreateCategoryModelMasterResult, error) {
	callback := make(chan CreateCategoryModelMasterAsyncResult, 1)
	go p.CreateCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getCategoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCategoryModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCategoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetCategoryModelMasterAsync(
	request *GetCategoryModelMasterRequest,
	callback chan<- GetCategoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModelMaster",
    		"function": "getCategoryModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCategoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetCategoryModelMaster(
	request *GetCategoryModelMasterRequest,
) (*GetCategoryModelMasterResult, error) {
	callback := make(chan GetCategoryModelMasterAsyncResult, 1)
	go p.GetCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) updateCategoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCategoryModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCategoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) UpdateCategoryModelMasterAsync(
	request *UpdateCategoryModelMasterRequest,
	callback chan<- UpdateCategoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModelMaster",
    		"function": "updateCategoryModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.MinimumValue != nil {
        bodies["minimumValue"] = *request.MinimumValue
    }
    if request.MaximumValue != nil {
        bodies["maximumValue"] = *request.MaximumValue
    }
    if request.OrderDirection != nil && *request.OrderDirection != "" {
        bodies["orderDirection"] = *request.OrderDirection
    }
    if request.Scope != nil && *request.Scope != "" {
        bodies["scope"] = *request.Scope
    }
    if request.UniqueByUserId != nil {
        bodies["uniqueByUserId"] = *request.UniqueByUserId
    }
    if request.CalculateFixedTimingHour != nil {
        bodies["calculateFixedTimingHour"] = *request.CalculateFixedTimingHour
    }
    if request.CalculateFixedTimingMinute != nil {
        bodies["calculateFixedTimingMinute"] = *request.CalculateFixedTimingMinute
    }
    if request.CalculateIntervalMinutes != nil {
        bodies["calculateIntervalMinutes"] = *request.CalculateIntervalMinutes
    }
    if request.EntryPeriodEventId != nil && *request.EntryPeriodEventId != "" {
        bodies["entryPeriodEventId"] = *request.EntryPeriodEventId
    }
    if request.AccessPeriodEventId != nil && *request.AccessPeriodEventId != "" {
        bodies["accessPeriodEventId"] = *request.AccessPeriodEventId
    }
    if request.Generation != nil && *request.Generation != "" {
        bodies["generation"] = *request.Generation
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCategoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) UpdateCategoryModelMaster(
	request *UpdateCategoryModelMasterRequest,
) (*UpdateCategoryModelMasterResult, error) {
	callback := make(chan UpdateCategoryModelMasterAsyncResult, 1)
	go p.UpdateCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) deleteCategoryModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteCategoryModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteCategoryModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteCategoryModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteCategoryModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteCategoryModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DeleteCategoryModelMasterAsync(
	request *DeleteCategoryModelMasterRequest,
	callback chan<- DeleteCategoryModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "categoryModelMaster",
    		"function": "deleteCategoryModelMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteCategoryModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DeleteCategoryModelMaster(
	request *DeleteCategoryModelMasterRequest,
) (*DeleteCategoryModelMasterResult, error) {
	callback := make(chan DeleteCategoryModelMasterAsyncResult, 1)
	go p.DeleteCategoryModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeSubscribesByCategoryNameAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribesByCategoryNameAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByCategoryNameAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByCategoryNameResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSubscribesByCategoryNameAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSubscribesByCategoryNameAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeSubscribesByCategoryNameAsync(
	request *DescribeSubscribesByCategoryNameRequest,
	callback chan<- DescribeSubscribesByCategoryNameAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "describeSubscribesByCategoryName",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeSubscribesByCategoryNameAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeSubscribesByCategoryName(
	request *DescribeSubscribesByCategoryNameRequest,
) (*DescribeSubscribesByCategoryNameResult, error) {
	callback := make(chan DescribeSubscribesByCategoryNameAsyncResult, 1)
	go p.DescribeSubscribesByCategoryNameAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeSubscribesByCategoryNameAndUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeSubscribesByCategoryNameAndUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeSubscribesByCategoryNameAndUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeSubscribesByCategoryNameAndUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeSubscribesByCategoryNameAndUserIdAsync(
	request *DescribeSubscribesByCategoryNameAndUserIdRequest,
	callback chan<- DescribeSubscribesByCategoryNameAndUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "describeSubscribesByCategoryNameAndUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeSubscribesByCategoryNameAndUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeSubscribesByCategoryNameAndUserId(
	request *DescribeSubscribesByCategoryNameAndUserIdRequest,
) (*DescribeSubscribesByCategoryNameAndUserIdResult, error) {
	callback := make(chan DescribeSubscribesByCategoryNameAndUserIdAsyncResult, 1)
	go p.DescribeSubscribesByCategoryNameAndUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) subscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SubscribeAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) SubscribeAsync(
	request *SubscribeRequest,
	callback chan<- SubscribeAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "subscribe",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.subscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) Subscribe(
	request *SubscribeRequest,
) (*SubscribeResult, error) {
	callback := make(chan SubscribeAsyncResult, 1)
	go p.SubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) subscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SubscribeByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SubscribeByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) SubscribeByUserIdAsync(
	request *SubscribeByUserIdRequest,
	callback chan<- SubscribeByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "subscribeByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.subscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) SubscribeByUserId(
	request *SubscribeByUserIdRequest,
) (*SubscribeByUserIdResult, error) {
	callback := make(chan SubscribeByUserIdAsyncResult, 1)
	go p.SubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getSubscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSubscribeAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetSubscribeAsync(
	request *GetSubscribeRequest,
	callback chan<- GetSubscribeAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "getSubscribe",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getSubscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetSubscribe(
	request *GetSubscribeRequest,
) (*GetSubscribeResult, error) {
	callback := make(chan GetSubscribeAsyncResult, 1)
	go p.GetSubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getSubscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetSubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetSubscribeByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetSubscribeByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetSubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetSubscribeByUserIdAsync(
	request *GetSubscribeByUserIdRequest,
	callback chan<- GetSubscribeByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "getSubscribeByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getSubscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetSubscribeByUserId(
	request *GetSubscribeByUserIdRequest,
) (*GetSubscribeByUserIdResult, error) {
	callback := make(chan GetSubscribeByUserIdAsyncResult, 1)
	go p.GetSubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) unsubscribeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnsubscribeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UnsubscribeAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UnsubscribeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) UnsubscribeAsync(
	request *UnsubscribeRequest,
	callback chan<- UnsubscribeAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "unsubscribe",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.unsubscribeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) Unsubscribe(
	request *UnsubscribeRequest,
) (*UnsubscribeResult, error) {
	callback := make(chan UnsubscribeAsyncResult, 1)
	go p.UnsubscribeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) unsubscribeByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UnsubscribeByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UnsubscribeByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UnsubscribeByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UnsubscribeByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) UnsubscribeByUserIdAsync(
	request *UnsubscribeByUserIdRequest,
	callback chan<- UnsubscribeByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "subscribe",
    		"function": "unsubscribeByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.TargetUserId != nil && *request.TargetUserId != "" {
        bodies["targetUserId"] = *request.TargetUserId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.unsubscribeByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) UnsubscribeByUserId(
	request *UnsubscribeByUserIdRequest,
) (*UnsubscribeByUserIdResult, error) {
	callback := make(chan UnsubscribeByUserIdAsyncResult, 1)
	go p.UnsubscribeByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeScoresAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeScoresAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeScoresAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeScoresResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeScoresAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeScoresAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeScoresAsync(
	request *DescribeScoresRequest,
	callback chan<- DescribeScoresAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "score",
    		"function": "describeScores",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeScoresAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeScores(
	request *DescribeScoresRequest,
) (*DescribeScoresResult, error) {
	callback := make(chan DescribeScoresAsyncResult, 1)
	go p.DescribeScoresAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeScoresByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeScoresByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeScoresByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeScoresByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeScoresByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeScoresByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeScoresByUserIdAsync(
	request *DescribeScoresByUserIdRequest,
	callback chan<- DescribeScoresByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "score",
    		"function": "describeScoresByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeScoresByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeScoresByUserId(
	request *DescribeScoresByUserIdRequest,
) (*DescribeScoresByUserIdResult, error) {
	callback := make(chan DescribeScoresByUserIdAsyncResult, 1)
	go p.DescribeScoresByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetScoreResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetScoreAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetScoreAsync(
	request *GetScoreRequest,
	callback chan<- GetScoreAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "score",
    		"function": "getScore",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.UniqueId != nil && *request.UniqueId != "" {
        bodies["uniqueId"] = *request.UniqueId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetScore(
	request *GetScoreRequest,
) (*GetScoreResult, error) {
	callback := make(chan GetScoreAsyncResult, 1)
	go p.GetScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetScoreByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetScoreByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetScoreByUserIdAsync(
	request *GetScoreByUserIdRequest,
	callback chan<- GetScoreByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "score",
    		"function": "getScoreByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.UniqueId != nil && *request.UniqueId != "" {
        bodies["uniqueId"] = *request.UniqueId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetScoreByUserId(
	request *GetScoreByUserIdRequest,
) (*GetScoreByUserIdResult, error) {
	callback := make(chan GetScoreByUserIdAsyncResult, 1)
	go p.GetScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeRankingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRankingsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeRankingsAsync(
	request *DescribeRankingsRequest,
	callback chan<- DescribeRankingsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "describeRankings",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.StartIndex != nil {
        bodies["startIndex"] = *request.StartIndex
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.describeRankingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeRankings(
	request *DescribeRankingsRequest,
) (*DescribeRankingsResult, error) {
	callback := make(chan DescribeRankingsAsyncResult, 1)
	go p.DescribeRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeRankingssByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRankingssByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingssByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingssByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeRankingssByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeRankingssByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeRankingssByUserIdAsync(
	request *DescribeRankingssByUserIdRequest,
	callback chan<- DescribeRankingssByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "describeRankingssByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.StartIndex != nil {
        bodies["startIndex"] = *request.StartIndex
    }
    if request.PageToken != nil && *request.PageToken != "" {
        bodies["pageToken"] = *request.PageToken
    }
    if request.Limit != nil {
        bodies["limit"] = *request.Limit
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeRankingssByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeRankingssByUserId(
	request *DescribeRankingssByUserIdRequest,
) (*DescribeRankingssByUserIdResult, error) {
	callback := make(chan DescribeRankingssByUserIdAsyncResult, 1)
	go p.DescribeRankingssByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) describeNearRankingsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeNearRankingsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeNearRankingsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeNearRankingsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeNearRankingsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeNearRankingsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) DescribeNearRankingsAsync(
	request *DescribeNearRankingsRequest,
	callback chan<- DescribeNearRankingsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "describeNearRankings",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.Score != nil {
        bodies["score"] = *request.Score
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.describeNearRankingsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) DescribeNearRankings(
	request *DescribeNearRankingsRequest,
) (*DescribeNearRankingsResult, error) {
	callback := make(chan DescribeNearRankingsAsyncResult, 1)
	go p.DescribeNearRankingsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getRankingAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRankingAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRankingAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRankingAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetRankingAsync(
	request *GetRankingRequest,
	callback chan<- GetRankingAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "getRanking",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.UniqueId != nil && *request.UniqueId != "" {
        bodies["uniqueId"] = *request.UniqueId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getRankingAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetRanking(
	request *GetRankingRequest,
) (*GetRankingResult, error) {
	callback := make(chan GetRankingAsyncResult, 1)
	go p.GetRankingAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) getRankingByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRankingByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetRankingByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetRankingByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetRankingByUserIdAsync(
	request *GetRankingByUserIdRequest,
	callback chan<- GetRankingByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "getRankingByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ScorerUserId != nil && *request.ScorerUserId != "" {
        bodies["scorerUserId"] = *request.ScorerUserId
    }
    if request.UniqueId != nil && *request.UniqueId != "" {
        bodies["uniqueId"] = *request.UniqueId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getRankingByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetRankingByUserId(
	request *GetRankingByUserIdRequest,
) (*GetRankingByUserIdResult, error) {
	callback := make(chan GetRankingByUserIdAsyncResult, 1)
	go p.GetRankingByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) putScoreAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutScoreAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutScoreAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutScoreResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PutScoreAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PutScoreAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) PutScoreAsync(
	request *PutScoreRequest,
	callback chan<- PutScoreAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "putScore",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.Score != nil {
        bodies["score"] = *request.Score
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.putScoreAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) PutScore(
	request *PutScoreRequest,
) (*PutScoreResult, error) {
	callback := make(chan PutScoreAsyncResult, 1)
	go p.PutScoreAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) putScoreByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- PutScoreByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- PutScoreByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result PutScoreByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- PutScoreByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- PutScoreByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) PutScoreByUserIdAsync(
	request *PutScoreByUserIdRequest,
	callback chan<- PutScoreByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "ranking",
    		"function": "putScoreByUserId",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CategoryName != nil && *request.CategoryName != "" {
        bodies["categoryName"] = *request.CategoryName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.Score != nil {
        bodies["score"] = *request.Score
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.putScoreByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) PutScoreByUserId(
	request *PutScoreByUserIdRequest,
) (*PutScoreByUserIdResult, error) {
	callback := make(chan PutScoreByUserIdAsyncResult, 1)
	go p.PutScoreByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) exportMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- ExportMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
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

func (p Gs2RankingWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "currentRankingMaster",
    		"function": "exportMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.exportMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) ExportMaster(
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

func (p Gs2RankingWebSocketClient) getCurrentRankingMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentRankingMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentRankingMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) GetCurrentRankingMasterAsync(
	request *GetCurrentRankingMasterRequest,
	callback chan<- GetCurrentRankingMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "currentRankingMaster",
    		"function": "getCurrentRankingMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCurrentRankingMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) GetCurrentRankingMaster(
	request *GetCurrentRankingMasterRequest,
) (*GetCurrentRankingMasterResult, error) {
	callback := make(chan GetCurrentRankingMasterAsyncResult, 1)
	go p.GetCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) updateCurrentRankingMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRankingMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRankingMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) UpdateCurrentRankingMasterAsync(
	request *UpdateCurrentRankingMasterRequest,
	callback chan<- UpdateCurrentRankingMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "currentRankingMaster",
    		"function": "updateCurrentRankingMaster",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.Settings != nil && *request.Settings != "" {
        bodies["settings"] = *request.Settings
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentRankingMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) UpdateCurrentRankingMaster(
	request *UpdateCurrentRankingMasterRequest,
) (*UpdateCurrentRankingMasterResult, error) {
	callback := make(chan UpdateCurrentRankingMasterAsyncResult, 1)
	go p.UpdateCurrentRankingMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2RankingWebSocketClient) updateCurrentRankingMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentRankingMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentRankingMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2RankingWebSocketClient) UpdateCurrentRankingMasterFromGitHubAsync(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
	callback chan<- UpdateCurrentRankingMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "ranking",
    		"component": "currentRankingMaster",
    		"function": "updateCurrentRankingMasterFromGitHub",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.NamespaceName != nil && *request.NamespaceName != "" {
        bodies["namespaceName"] = *request.NamespaceName
    }
    if request.CheckoutSetting != nil {
        bodies["checkoutSetting"] = request.CheckoutSetting.ToDict()
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateCurrentRankingMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2RankingWebSocketClient) UpdateCurrentRankingMasterFromGitHub(
	request *UpdateCurrentRankingMasterFromGitHubRequest,
) (*UpdateCurrentRankingMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentRankingMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentRankingMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
