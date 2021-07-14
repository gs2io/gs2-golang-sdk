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

package experience

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2ExperienceWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2ExperienceWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2ExperienceWebSocketClient) describeNamespacesAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) DescribeNamespacesAsync(
	request *DescribeNamespacesRequest,
	callback chan<- DescribeNamespacesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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

func (p Gs2ExperienceWebSocketClient) DescribeNamespaces(
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

func (p Gs2ExperienceWebSocketClient) createNamespaceAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) CreateNamespaceAsync(
	request *CreateNamespaceRequest,
	callback chan<- CreateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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
    if request.ExperienceCapScriptId != nil && *request.ExperienceCapScriptId != "" {
        bodies["experienceCapScriptId"] = *request.ExperienceCapScriptId
    }
    if request.ChangeExperienceScript != nil {
        bodies["changeExperienceScript"] = request.ChangeExperienceScript.ToDict()
    }
    if request.ChangeRankScript != nil {
        bodies["changeRankScript"] = request.ChangeRankScript.ToDict()
    }
    if request.ChangeRankCapScript != nil {
        bodies["changeRankCapScript"] = request.ChangeRankCapScript.ToDict()
    }
    if request.OverflowExperienceScript != nil {
        bodies["overflowExperienceScript"] = request.OverflowExperienceScript.ToDict()
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

func (p Gs2ExperienceWebSocketClient) CreateNamespace(
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

func (p Gs2ExperienceWebSocketClient) getNamespaceStatusAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) GetNamespaceStatusAsync(
	request *GetNamespaceStatusRequest,
	callback chan<- GetNamespaceStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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

func (p Gs2ExperienceWebSocketClient) GetNamespaceStatus(
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

func (p Gs2ExperienceWebSocketClient) getNamespaceAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) GetNamespaceAsync(
	request *GetNamespaceRequest,
	callback chan<- GetNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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

func (p Gs2ExperienceWebSocketClient) GetNamespace(
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

func (p Gs2ExperienceWebSocketClient) updateNamespaceAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) UpdateNamespaceAsync(
	request *UpdateNamespaceRequest,
	callback chan<- UpdateNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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
    if request.ExperienceCapScriptId != nil && *request.ExperienceCapScriptId != "" {
        bodies["experienceCapScriptId"] = *request.ExperienceCapScriptId
    }
    if request.ChangeExperienceScript != nil {
        bodies["changeExperienceScript"] = request.ChangeExperienceScript.ToDict()
    }
    if request.ChangeRankScript != nil {
        bodies["changeRankScript"] = request.ChangeRankScript.ToDict()
    }
    if request.ChangeRankCapScript != nil {
        bodies["changeRankCapScript"] = request.ChangeRankCapScript.ToDict()
    }
    if request.OverflowExperienceScript != nil {
        bodies["overflowExperienceScript"] = request.OverflowExperienceScript.ToDict()
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

func (p Gs2ExperienceWebSocketClient) UpdateNamespace(
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

func (p Gs2ExperienceWebSocketClient) deleteNamespaceAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) DeleteNamespaceAsync(
	request *DeleteNamespaceRequest,
	callback chan<- DeleteNamespaceAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
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

func (p Gs2ExperienceWebSocketClient) DeleteNamespace(
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

func (p Gs2ExperienceWebSocketClient) describeExperienceModelMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExperienceModelMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceModelMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceModelMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeExperienceModelMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeExperienceModelMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DescribeExperienceModelMastersAsync(
	request *DescribeExperienceModelMastersRequest,
	callback chan<- DescribeExperienceModelMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModelMaster",
    		"function": "describeExperienceModelMasters",
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

	go p.describeExperienceModelMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DescribeExperienceModelMasters(
	request *DescribeExperienceModelMastersRequest,
) (*DescribeExperienceModelMastersResult, error) {
	callback := make(chan DescribeExperienceModelMastersAsyncResult, 1)
	go p.DescribeExperienceModelMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) createExperienceModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateExperienceModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateExperienceModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) CreateExperienceModelMasterAsync(
	request *CreateExperienceModelMasterRequest,
	callback chan<- CreateExperienceModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModelMaster",
    		"function": "createExperienceModelMaster",
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
    if request.DefaultExperience != nil {
        bodies["defaultExperience"] = *request.DefaultExperience
    }
    if request.DefaultRankCap != nil {
        bodies["defaultRankCap"] = *request.DefaultRankCap
    }
    if request.MaxRankCap != nil {
        bodies["maxRankCap"] = *request.MaxRankCap
    }
    if request.RankThresholdName != nil && *request.RankThresholdName != "" {
        bodies["rankThresholdName"] = *request.RankThresholdName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createExperienceModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) CreateExperienceModelMaster(
	request *CreateExperienceModelMasterRequest,
) (*CreateExperienceModelMasterResult, error) {
	callback := make(chan CreateExperienceModelMasterAsyncResult, 1)
	go p.CreateExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getExperienceModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetExperienceModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetExperienceModelMasterAsync(
	request *GetExperienceModelMasterRequest,
	callback chan<- GetExperienceModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModelMaster",
    		"function": "getExperienceModelMaster",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getExperienceModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetExperienceModelMaster(
	request *GetExperienceModelMasterRequest,
) (*GetExperienceModelMasterResult, error) {
	callback := make(chan GetExperienceModelMasterAsyncResult, 1)
	go p.GetExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) updateExperienceModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateExperienceModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateExperienceModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) UpdateExperienceModelMasterAsync(
	request *UpdateExperienceModelMasterRequest,
	callback chan<- UpdateExperienceModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModelMaster",
    		"function": "updateExperienceModelMaster",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.DefaultExperience != nil {
        bodies["defaultExperience"] = *request.DefaultExperience
    }
    if request.DefaultRankCap != nil {
        bodies["defaultRankCap"] = *request.DefaultRankCap
    }
    if request.MaxRankCap != nil {
        bodies["maxRankCap"] = *request.MaxRankCap
    }
    if request.RankThresholdName != nil && *request.RankThresholdName != "" {
        bodies["rankThresholdName"] = *request.RankThresholdName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateExperienceModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) UpdateExperienceModelMaster(
	request *UpdateExperienceModelMasterRequest,
) (*UpdateExperienceModelMasterResult, error) {
	callback := make(chan UpdateExperienceModelMasterAsyncResult, 1)
	go p.UpdateExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) deleteExperienceModelMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteExperienceModelMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteExperienceModelMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteExperienceModelMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteExperienceModelMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteExperienceModelMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DeleteExperienceModelMasterAsync(
	request *DeleteExperienceModelMasterRequest,
	callback chan<- DeleteExperienceModelMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModelMaster",
    		"function": "deleteExperienceModelMaster",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteExperienceModelMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DeleteExperienceModelMaster(
	request *DeleteExperienceModelMasterRequest,
) (*DeleteExperienceModelMasterResult, error) {
	callback := make(chan DeleteExperienceModelMasterAsyncResult, 1)
	go p.DeleteExperienceModelMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) describeExperienceModelsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExperienceModelsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceModelsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceModelsResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeExperienceModelsAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeExperienceModelsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DescribeExperienceModelsAsync(
	request *DescribeExperienceModelsRequest,
	callback chan<- DescribeExperienceModelsAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModel",
    		"function": "describeExperienceModels",
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

	go p.describeExperienceModelsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DescribeExperienceModels(
	request *DescribeExperienceModelsRequest,
) (*DescribeExperienceModelsResult, error) {
	callback := make(chan DescribeExperienceModelsAsyncResult, 1)
	go p.DescribeExperienceModelsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getExperienceModelAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExperienceModelAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceModelAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceModelResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetExperienceModelAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetExperienceModelAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetExperienceModelAsync(
	request *GetExperienceModelRequest,
	callback chan<- GetExperienceModelAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "experienceModel",
    		"function": "getExperienceModel",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getExperienceModelAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetExperienceModel(
	request *GetExperienceModelRequest,
) (*GetExperienceModelResult, error) {
	callback := make(chan GetExperienceModelAsyncResult, 1)
	go p.GetExperienceModelAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) describeThresholdMastersAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeThresholdMastersAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeThresholdMastersAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeThresholdMastersResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeThresholdMastersAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeThresholdMastersAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DescribeThresholdMastersAsync(
	request *DescribeThresholdMastersRequest,
	callback chan<- DescribeThresholdMastersAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "thresholdMaster",
    		"function": "describeThresholdMasters",
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

	go p.describeThresholdMastersAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DescribeThresholdMasters(
	request *DescribeThresholdMastersRequest,
) (*DescribeThresholdMastersResult, error) {
	callback := make(chan DescribeThresholdMastersAsyncResult, 1)
	go p.DescribeThresholdMastersAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) createThresholdMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- CreateThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- CreateThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result CreateThresholdMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- CreateThresholdMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- CreateThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) CreateThresholdMasterAsync(
	request *CreateThresholdMasterRequest,
	callback chan<- CreateThresholdMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "thresholdMaster",
    		"function": "createThresholdMaster",
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
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.createThresholdMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) CreateThresholdMaster(
	request *CreateThresholdMasterRequest,
) (*CreateThresholdMasterResult, error) {
	callback := make(chan CreateThresholdMasterAsyncResult, 1)
	go p.CreateThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getThresholdMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetThresholdMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetThresholdMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetThresholdMasterAsync(
	request *GetThresholdMasterRequest,
	callback chan<- GetThresholdMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "thresholdMaster",
    		"function": "getThresholdMaster",
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
    if request.ThresholdName != nil && *request.ThresholdName != "" {
        bodies["thresholdName"] = *request.ThresholdName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getThresholdMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetThresholdMaster(
	request *GetThresholdMasterRequest,
) (*GetThresholdMasterResult, error) {
	callback := make(chan GetThresholdMasterAsyncResult, 1)
	go p.GetThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) updateThresholdMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateThresholdMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateThresholdMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) UpdateThresholdMasterAsync(
	request *UpdateThresholdMasterRequest,
	callback chan<- UpdateThresholdMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "thresholdMaster",
    		"function": "updateThresholdMaster",
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
    if request.ThresholdName != nil && *request.ThresholdName != "" {
        bodies["thresholdName"] = *request.ThresholdName
    }
    if request.Description != nil && *request.Description != "" {
        bodies["description"] = *request.Description
    }
    if request.Metadata != nil && *request.Metadata != "" {
        bodies["metadata"] = *request.Metadata
    }
    if request.Values != nil {
        var _values []interface {}
        for _, item := range request.Values {
            _values = append(_values, item)
        }
        bodies["values"] = _values
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.updateThresholdMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) UpdateThresholdMaster(
	request *UpdateThresholdMasterRequest,
) (*UpdateThresholdMasterResult, error) {
	callback := make(chan UpdateThresholdMasterAsyncResult, 1)
	go p.UpdateThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) deleteThresholdMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteThresholdMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteThresholdMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteThresholdMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteThresholdMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteThresholdMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DeleteThresholdMasterAsync(
	request *DeleteThresholdMasterRequest,
	callback chan<- DeleteThresholdMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "thresholdMaster",
    		"function": "deleteThresholdMaster",
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
    if request.ThresholdName != nil && *request.ThresholdName != "" {
        bodies["thresholdName"] = *request.ThresholdName
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteThresholdMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DeleteThresholdMaster(
	request *DeleteThresholdMasterRequest,
) (*DeleteThresholdMasterResult, error) {
	callback := make(chan DeleteThresholdMasterAsyncResult, 1)
	go p.DeleteThresholdMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) exportMasterAsyncHandler(
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

func (p Gs2ExperienceWebSocketClient) ExportMasterAsync(
	request *ExportMasterRequest,
	callback chan<- ExportMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "currentExperienceMaster",
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

func (p Gs2ExperienceWebSocketClient) ExportMaster(
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

func (p Gs2ExperienceWebSocketClient) getCurrentExperienceMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCurrentExperienceMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCurrentExperienceMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCurrentExperienceMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCurrentExperienceMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCurrentExperienceMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetCurrentExperienceMasterAsync(
	request *GetCurrentExperienceMasterRequest,
	callback chan<- GetCurrentExperienceMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "currentExperienceMaster",
    		"function": "getCurrentExperienceMaster",
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

	go p.getCurrentExperienceMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetCurrentExperienceMaster(
	request *GetCurrentExperienceMasterRequest,
) (*GetCurrentExperienceMasterResult, error) {
	callback := make(chan GetCurrentExperienceMasterAsyncResult, 1)
	go p.GetCurrentExperienceMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) updateCurrentExperienceMasterAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentExperienceMasterAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentExperienceMasterAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentExperienceMasterResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentExperienceMasterAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentExperienceMasterAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) UpdateCurrentExperienceMasterAsync(
	request *UpdateCurrentExperienceMasterRequest,
	callback chan<- UpdateCurrentExperienceMasterAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "currentExperienceMaster",
    		"function": "updateCurrentExperienceMaster",
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

	go p.updateCurrentExperienceMasterAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) UpdateCurrentExperienceMaster(
	request *UpdateCurrentExperienceMasterRequest,
) (*UpdateCurrentExperienceMasterResult, error) {
	callback := make(chan UpdateCurrentExperienceMasterAsyncResult, 1)
	go p.UpdateCurrentExperienceMasterAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) updateCurrentExperienceMasterFromGitHubAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- UpdateCurrentExperienceMasterFromGitHubAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result UpdateCurrentExperienceMasterFromGitHubResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- UpdateCurrentExperienceMasterFromGitHubAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) UpdateCurrentExperienceMasterFromGitHubAsync(
	request *UpdateCurrentExperienceMasterFromGitHubRequest,
	callback chan<- UpdateCurrentExperienceMasterFromGitHubAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "currentExperienceMaster",
    		"function": "updateCurrentExperienceMasterFromGitHub",
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

	go p.updateCurrentExperienceMasterFromGitHubAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) UpdateCurrentExperienceMasterFromGitHub(
	request *UpdateCurrentExperienceMasterFromGitHubRequest,
) (*UpdateCurrentExperienceMasterFromGitHubResult, error) {
	callback := make(chan UpdateCurrentExperienceMasterFromGitHubAsyncResult, 1)
	go p.UpdateCurrentExperienceMasterFromGitHubAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) describeStatusesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStatusesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStatusesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DescribeStatusesAsync(
	request *DescribeStatusesRequest,
	callback chan<- DescribeStatusesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "describeStatuses",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
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

	go p.describeStatusesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DescribeStatuses(
	request *DescribeStatusesRequest,
) (*DescribeStatusesResult, error) {
	callback := make(chan DescribeStatusesAsyncResult, 1)
	go p.DescribeStatusesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) describeStatusesByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStatusesByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStatusesByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeStatusesByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeStatusesByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DescribeStatusesByUserIdAsync(
	request *DescribeStatusesByUserIdRequest,
	callback chan<- DescribeStatusesByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "describeStatusesByUserId",
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
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
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

	go p.describeStatusesByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DescribeStatusesByUserId(
	request *DescribeStatusesByUserIdRequest,
) (*DescribeStatusesByUserIdResult, error) {
	callback := make(chan DescribeStatusesByUserIdAsyncResult, 1)
	go p.DescribeStatusesByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getStatusAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStatusAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStatusAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetStatusAsync(
	request *GetStatusRequest,
	callback chan<- GetStatusAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "getStatus",
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
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getStatusAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetStatus(
	request *GetStatusRequest,
) (*GetStatusResult, error) {
	callback := make(chan GetStatusAsyncResult, 1)
	go p.GetStatusAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetStatusByUserIdAsync(
	request *GetStatusByUserIdRequest,
	callback chan<- GetStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "getStatusByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetStatusByUserId(
	request *GetStatusByUserIdRequest,
) (*GetStatusByUserIdResult, error) {
	callback := make(chan GetStatusByUserIdAsyncResult, 1)
	go p.GetStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) getStatusWithSignatureAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStatusWithSignatureAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStatusWithSignatureAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStatusWithSignatureResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetStatusWithSignatureAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetStatusWithSignatureAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) GetStatusWithSignatureAsync(
	request *GetStatusWithSignatureRequest,
	callback chan<- GetStatusWithSignatureAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "getStatusWithSignature",
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
    if request.AccessToken != nil && *request.AccessToken != "" {
        bodies["accessToken"] = *request.AccessToken
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}
    if request.AccessToken != nil {
        bodies["xGs2AccessToken"] = string(*request.AccessToken)
    }

	go p.getStatusWithSignatureAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) GetStatusWithSignature(
	request *GetStatusWithSignatureRequest,
) (*GetStatusWithSignatureResult, error) {
	callback := make(chan GetStatusWithSignatureAsyncResult, 1)
	go p.GetStatusWithSignatureAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) addExperienceByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddExperienceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddExperienceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddExperienceByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddExperienceByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddExperienceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) AddExperienceByUserIdAsync(
	request *AddExperienceByUserIdRequest,
	callback chan<- AddExperienceByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "addExperienceByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.ExperienceValue != nil {
        bodies["experienceValue"] = *request.ExperienceValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addExperienceByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) AddExperienceByUserId(
	request *AddExperienceByUserIdRequest,
) (*AddExperienceByUserIdResult, error) {
	callback := make(chan AddExperienceByUserIdAsyncResult, 1)
	go p.AddExperienceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) setExperienceByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetExperienceByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetExperienceByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetExperienceByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetExperienceByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetExperienceByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) SetExperienceByUserIdAsync(
	request *SetExperienceByUserIdRequest,
	callback chan<- SetExperienceByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "setExperienceByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.ExperienceValue != nil {
        bodies["experienceValue"] = *request.ExperienceValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setExperienceByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) SetExperienceByUserId(
	request *SetExperienceByUserIdRequest,
) (*SetExperienceByUserIdResult, error) {
	callback := make(chan SetExperienceByUserIdAsyncResult, 1)
	go p.SetExperienceByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) addRankCapByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRankCapByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddRankCapByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) AddRankCapByUserIdAsync(
	request *AddRankCapByUserIdRequest,
	callback chan<- AddRankCapByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "addRankCapByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.RankCapValue != nil {
        bodies["rankCapValue"] = *request.RankCapValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addRankCapByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) AddRankCapByUserId(
	request *AddRankCapByUserIdRequest,
) (*AddRankCapByUserIdResult, error) {
	callback := make(chan AddRankCapByUserIdAsyncResult, 1)
	go p.AddRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) setRankCapByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRankCapByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRankCapByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRankCapByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRankCapByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRankCapByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) SetRankCapByUserIdAsync(
	request *SetRankCapByUserIdRequest,
	callback chan<- SetRankCapByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "setRankCapByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
    if request.RankCapValue != nil {
        bodies["rankCapValue"] = *request.RankCapValue
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setRankCapByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) SetRankCapByUserId(
	request *SetRankCapByUserIdRequest,
) (*SetRankCapByUserIdResult, error) {
	callback := make(chan SetRankCapByUserIdAsyncResult, 1)
	go p.SetRankCapByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) deleteStatusByUserIdAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DeleteStatusByUserIdAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DeleteStatusByUserIdResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DeleteStatusByUserIdAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DeleteStatusByUserIdAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) DeleteStatusByUserIdAsync(
	request *DeleteStatusByUserIdRequest,
	callback chan<- DeleteStatusByUserIdAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "deleteStatusByUserId",
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
    if request.UserId != nil && *request.UserId != "" {
        bodies["userId"] = *request.UserId
    }
    if request.ExperienceName != nil && *request.ExperienceName != "" {
        bodies["experienceName"] = *request.ExperienceName
    }
    if request.PropertyId != nil && *request.PropertyId != "" {
        bodies["propertyId"] = *request.PropertyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.deleteStatusByUserIdAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) DeleteStatusByUserId(
	request *DeleteStatusByUserIdRequest,
) (*DeleteStatusByUserIdResult, error) {
	callback := make(chan DeleteStatusByUserIdAsyncResult, 1)
	go p.DeleteStatusByUserIdAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) addExperienceByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddExperienceByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddExperienceByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddExperienceByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddExperienceByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddExperienceByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) AddExperienceByStampSheetAsync(
	request *AddExperienceByStampSheetRequest,
	callback chan<- AddExperienceByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "addExperienceByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addExperienceByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) AddExperienceByStampSheet(
	request *AddExperienceByStampSheetRequest,
) (*AddExperienceByStampSheetResult, error) {
	callback := make(chan AddExperienceByStampSheetAsyncResult, 1)
	go p.AddExperienceByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) addRankCapByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- AddRankCapByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- AddRankCapByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result AddRankCapByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- AddRankCapByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- AddRankCapByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) AddRankCapByStampSheetAsync(
	request *AddRankCapByStampSheetRequest,
	callback chan<- AddRankCapByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "addRankCapByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.addRankCapByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) AddRankCapByStampSheet(
	request *AddRankCapByStampSheetRequest,
) (*AddRankCapByStampSheetResult, error) {
	callback := make(chan AddRankCapByStampSheetAsyncResult, 1)
	go p.AddRankCapByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2ExperienceWebSocketClient) setRankCapByStampSheetAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- SetRankCapByStampSheetAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- SetRankCapByStampSheetAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result SetRankCapByStampSheetResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- SetRankCapByStampSheetAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- SetRankCapByStampSheetAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2ExperienceWebSocketClient) SetRankCapByStampSheetAsync(
	request *SetRankCapByStampSheetRequest,
	callback chan<- SetRankCapByStampSheetAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "experience",
    		"component": "status",
    		"function": "setRankCapByStampSheet",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.StampSheet != nil && *request.StampSheet != "" {
        bodies["stampSheet"] = *request.StampSheet
    }
    if request.KeyId != nil && *request.KeyId != "" {
        bodies["keyId"] = *request.KeyId
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.setRankCapByStampSheetAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2ExperienceWebSocketClient) SetRankCapByStampSheet(
	request *SetRankCapByStampSheetRequest,
) (*SetRankCapByStampSheetResult, error) {
	callback := make(chan SetRankCapByStampSheetAsyncResult, 1)
	go p.SetRankCapByStampSheetAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
