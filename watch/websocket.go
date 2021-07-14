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

package watch

import (
	"encoding/json"
	"github.com/google/uuid"
	"core"
)

type Gs2WatchWebSocketClient struct {
	Session *core.Gs2WebSocketSession
}

func (p Gs2WatchWebSocketClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func (p Gs2WatchWebSocketClient) getChartAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetChartAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetChartAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetChartResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetChartAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetChartAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetChartAsync(
	request *GetChartRequest,
	callback chan<- GetChartAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "watch",
    		"component": "chart",
    		"function": "getChart",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.Metrics != nil && *request.Metrics != "" {
        bodies["metrics"] = *request.Metrics
    }
    if request.Grn != nil && *request.Grn != "" {
        bodies["grn"] = *request.Grn
    }
    if request.Queries != nil {
        var _queries []interface {}
        for _, item := range request.Queries {
            _queries = append(_queries, item)
        }
        bodies["queries"] = _queries
    }
    if request.By != nil && *request.By != "" {
        bodies["by"] = *request.By
    }
    if request.Timeframe != nil && *request.Timeframe != "" {
        bodies["timeframe"] = *request.Timeframe
    }
    if request.Size != nil && *request.Size != "" {
        bodies["size"] = *request.Size
    }
    if request.Format != nil && *request.Format != "" {
        bodies["format"] = *request.Format
    }
    if request.Aggregator != nil && *request.Aggregator != "" {
        bodies["aggregator"] = *request.Aggregator
    }
    if request.Style != nil && *request.Style != "" {
        bodies["style"] = *request.Style
    }
    if request.Title != nil && *request.Title != "" {
        bodies["title"] = *request.Title
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getChartAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetChart(
	request *GetChartRequest,
) (*GetChartResult, error) {
	callback := make(chan GetChartAsyncResult, 1)
	go p.GetChartAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getCumulativeAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetCumulativeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetCumulativeAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetCumulativeResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetCumulativeAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetCumulativeAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetCumulativeAsync(
	request *GetCumulativeRequest,
	callback chan<- GetCumulativeAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "watch",
    		"component": "cumulative",
    		"function": "getCumulative",
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
    if request.ResourceGrn != nil && *request.ResourceGrn != "" {
        bodies["resourceGrn"] = *request.ResourceGrn
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getCumulativeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetCumulative(
	request *GetCumulativeRequest,
) (*GetCumulativeResult, error) {
	callback := make(chan GetCumulativeAsyncResult, 1)
	go p.GetCumulativeAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeBillingActivitiesAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeBillingActivitiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeBillingActivitiesAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeBillingActivitiesResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- DescribeBillingActivitiesAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- DescribeBillingActivitiesAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeBillingActivitiesAsync(
	request *DescribeBillingActivitiesRequest,
	callback chan<- DescribeBillingActivitiesAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "watch",
    		"component": "billingActivity",
    		"function": "describeBillingActivities",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.Year != nil {
        bodies["year"] = *request.Year
    }
    if request.Month != nil {
        bodies["month"] = *request.Month
    }
    if request.Service != nil && *request.Service != "" {
        bodies["service"] = *request.Service
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

	go p.describeBillingActivitiesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeBillingActivities(
	request *DescribeBillingActivitiesRequest,
) (*DescribeBillingActivitiesResult, error) {
	callback := make(chan DescribeBillingActivitiesAsyncResult, 1)
	go p.DescribeBillingActivitiesAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getBillingActivityAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetBillingActivityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetBillingActivityAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetBillingActivityResult
	if asyncResult.Payload != "" {
        err = json.Unmarshal([]byte(asyncResult.Payload), &result)
        if err != nil {
            callback <- GetBillingActivityAsyncResult{
                err: err,
            }
            return
        }
	}
	callback <- GetBillingActivityAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetBillingActivityAsync(
	request *GetBillingActivityRequest,
	callback chan<- GetBillingActivityAsyncResult,
) {
    requestId := core.WebSocketRequestId(uuid.New().String())
    var bodies = core.WebSocketBodies{
    	"x_gs2": map[string]interface{} {
    		"service": "watch",
    		"component": "billingActivity",
    		"function": "getBillingActivity",
            "contentType": "application/json",
    		"requestId": requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
    if request.Year != nil {
        bodies["year"] = *request.Year
    }
    if request.Month != nil {
        bodies["month"] = *request.Month
    }
    if request.Service != nil && *request.Service != "" {
        bodies["service"] = *request.Service
    }
    if request.ActivityType != nil && *request.ActivityType != "" {
        bodies["activityType"] = *request.ActivityType
    }
	if request.ContextStack != nil {
    	bodies["contextStack"] = *request.ContextStack;
	}

	go p.getBillingActivityAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies: bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetBillingActivity(
	request *GetBillingActivityRequest,
) (*GetBillingActivityResult, error) {
	callback := make(chan GetBillingActivityAsyncResult, 1)
	go p.GetBillingActivityAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
