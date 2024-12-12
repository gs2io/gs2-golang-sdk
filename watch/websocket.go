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
	"github.com/gs2io/gs2-golang-sdk/core"
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "chart",
			"function":    "getChart",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.Measure != nil && *request.Measure != "" {
		bodies["measure"] = *request.Measure
	}
	if request.Grn != nil && *request.Grn != "" {
		bodies["grn"] = *request.Grn
	}
	if request.Round != nil && *request.Round != "" {
		bodies["round"] = *request.Round
	}
	if request.Filters != nil {
		var _filters []interface{}
		for _, item := range request.Filters {
			_filters = append(_filters, item)
		}
		bodies["filters"] = _filters
	}
	if request.GroupBys != nil {
		var _groupBys []interface{}
		for _, item := range request.GroupBys {
			_groupBys = append(_groupBys, item)
		}
		bodies["groupBys"] = _groupBys
	}
	if request.CountBy != nil && *request.CountBy != "" {
		bodies["countBy"] = *request.CountBy
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getChartAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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

func (p Gs2WatchWebSocketClient) getDistributionAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDistributionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDistributionAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDistributionResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDistributionAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDistributionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetDistributionAsync(
	request *GetDistributionRequest,
	callback chan<- GetDistributionAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "distribution",
			"function":    "getDistribution",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.Measure != nil && *request.Measure != "" {
		bodies["measure"] = *request.Measure
	}
	if request.Grn != nil && *request.Grn != "" {
		bodies["grn"] = *request.Grn
	}
	if request.Filters != nil {
		var _filters []interface{}
		for _, item := range request.Filters {
			_filters = append(_filters, item)
		}
		bodies["filters"] = _filters
	}
	if request.GroupBys != nil {
		var _groupBys []interface{}
		for _, item := range request.GroupBys {
			_groupBys = append(_groupBys, item)
		}
		bodies["groupBys"] = _groupBys
	}
	if request.Begin != nil {
		bodies["begin"] = *request.Begin
	}
	if request.End != nil {
		bodies["end"] = *request.End
	}
	if request.PageToken != nil && *request.PageToken != "" {
		bodies["pageToken"] = *request.PageToken
	}
	if request.Limit != nil {
		bodies["limit"] = *request.Limit
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getDistributionAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetDistribution(
	request *GetDistributionRequest,
) (*GetDistributionResult, error) {
	callback := make(chan GetDistributionAsyncResult, 1)
	go p.GetDistributionAsync(
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "cumulative",
			"function":    "getCumulative",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getCumulativeAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "billingActivity",
			"function":    "describeBillingActivities",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeBillingActivitiesAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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
	if asyncResult.Err != nil {
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
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "billingActivity",
			"function":    "getBillingActivity",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getBillingActivityAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
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

func (p Gs2WatchWebSocketClient) getGeneralMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetGeneralMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetGeneralMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetGeneralMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGeneralMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetGeneralMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetGeneralMetricsAsync(
	request *GetGeneralMetricsRequest,
	callback chan<- GetGeneralMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "generalMetrics",
			"function":    "getGeneralMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getGeneralMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetGeneralMetrics(
	request *GetGeneralMetricsRequest,
) (*GetGeneralMetricsResult, error) {
	callback := make(chan GetGeneralMetricsAsyncResult, 1)
	go p.GetGeneralMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeAccountNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeAccountNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeAccountNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeAccountNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAccountNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeAccountNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeAccountNamespaceMetricsAsync(
	request *DescribeAccountNamespaceMetricsRequest,
	callback chan<- DescribeAccountNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "accountNamespace",
			"function":    "describeAccountNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeAccountNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeAccountNamespaceMetrics(
	request *DescribeAccountNamespaceMetricsRequest,
) (*DescribeAccountNamespaceMetricsResult, error) {
	callback := make(chan DescribeAccountNamespaceMetricsAsyncResult, 1)
	go p.DescribeAccountNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getAccountNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetAccountNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetAccountNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetAccountNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAccountNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetAccountNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetAccountNamespaceMetricsAsync(
	request *GetAccountNamespaceMetricsRequest,
	callback chan<- GetAccountNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "accountNamespace",
			"function":    "getAccountNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getAccountNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetAccountNamespaceMetrics(
	request *GetAccountNamespaceMetricsRequest,
) (*GetAccountNamespaceMetricsResult, error) {
	callback := make(chan GetAccountNamespaceMetricsAsyncResult, 1)
	go p.GetAccountNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeChatNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeChatNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeChatNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeChatNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeChatNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeChatNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeChatNamespaceMetricsAsync(
	request *DescribeChatNamespaceMetricsRequest,
	callback chan<- DescribeChatNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "chatNamespace",
			"function":    "describeChatNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeChatNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeChatNamespaceMetrics(
	request *DescribeChatNamespaceMetricsRequest,
) (*DescribeChatNamespaceMetricsResult, error) {
	callback := make(chan DescribeChatNamespaceMetricsAsyncResult, 1)
	go p.DescribeChatNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getChatNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetChatNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetChatNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetChatNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetChatNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetChatNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetChatNamespaceMetricsAsync(
	request *GetChatNamespaceMetricsRequest,
	callback chan<- GetChatNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "chatNamespace",
			"function":    "getChatNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getChatNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetChatNamespaceMetrics(
	request *GetChatNamespaceMetricsRequest,
) (*GetChatNamespaceMetricsResult, error) {
	callback := make(chan GetChatNamespaceMetricsAsyncResult, 1)
	go p.GetChatNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeDatastoreNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDatastoreNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDatastoreNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeDatastoreNamespaceMetricsAsync(
	request *DescribeDatastoreNamespaceMetricsRequest,
	callback chan<- DescribeDatastoreNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "datastoreNamespace",
			"function":    "describeDatastoreNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeDatastoreNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeDatastoreNamespaceMetrics(
	request *DescribeDatastoreNamespaceMetricsRequest,
) (*DescribeDatastoreNamespaceMetricsResult, error) {
	callback := make(chan DescribeDatastoreNamespaceMetricsAsyncResult, 1)
	go p.DescribeDatastoreNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getDatastoreNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDatastoreNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDatastoreNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDatastoreNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDatastoreNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDatastoreNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetDatastoreNamespaceMetricsAsync(
	request *GetDatastoreNamespaceMetricsRequest,
	callback chan<- GetDatastoreNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "datastoreNamespace",
			"function":    "getDatastoreNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getDatastoreNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetDatastoreNamespaceMetrics(
	request *GetDatastoreNamespaceMetricsRequest,
) (*GetDatastoreNamespaceMetricsResult, error) {
	callback := make(chan GetDatastoreNamespaceMetricsAsyncResult, 1)
	go p.GetDatastoreNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeDictionaryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeDictionaryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeDictionaryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeDictionaryNamespaceMetricsAsync(
	request *DescribeDictionaryNamespaceMetricsRequest,
	callback chan<- DescribeDictionaryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "dictionaryNamespace",
			"function":    "describeDictionaryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeDictionaryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeDictionaryNamespaceMetrics(
	request *DescribeDictionaryNamespaceMetricsRequest,
) (*DescribeDictionaryNamespaceMetricsResult, error) {
	callback := make(chan DescribeDictionaryNamespaceMetricsAsyncResult, 1)
	go p.DescribeDictionaryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getDictionaryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetDictionaryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetDictionaryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetDictionaryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDictionaryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetDictionaryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetDictionaryNamespaceMetricsAsync(
	request *GetDictionaryNamespaceMetricsRequest,
	callback chan<- GetDictionaryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "dictionaryNamespace",
			"function":    "getDictionaryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getDictionaryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetDictionaryNamespaceMetrics(
	request *GetDictionaryNamespaceMetricsRequest,
) (*GetDictionaryNamespaceMetricsResult, error) {
	callback := make(chan GetDictionaryNamespaceMetricsAsyncResult, 1)
	go p.GetDictionaryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeExchangeRateModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExchangeRateModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExchangeRateModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExchangeRateModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExchangeRateModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeExchangeRateModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeExchangeRateModelMetricsAsync(
	request *DescribeExchangeRateModelMetricsRequest,
	callback chan<- DescribeExchangeRateModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "exchangeRateModel",
			"function":    "describeExchangeRateModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeExchangeRateModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeExchangeRateModelMetrics(
	request *DescribeExchangeRateModelMetricsRequest,
) (*DescribeExchangeRateModelMetricsResult, error) {
	callback := make(chan DescribeExchangeRateModelMetricsAsyncResult, 1)
	go p.DescribeExchangeRateModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getExchangeRateModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExchangeRateModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExchangeRateModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExchangeRateModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExchangeRateModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetExchangeRateModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetExchangeRateModelMetricsAsync(
	request *GetExchangeRateModelMetricsRequest,
	callback chan<- GetExchangeRateModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "exchangeRateModel",
			"function":    "getExchangeRateModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.RateName != nil && *request.RateName != "" {
		bodies["rateName"] = *request.RateName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getExchangeRateModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetExchangeRateModelMetrics(
	request *GetExchangeRateModelMetricsRequest,
) (*GetExchangeRateModelMetricsResult, error) {
	callback := make(chan GetExchangeRateModelMetricsAsyncResult, 1)
	go p.GetExchangeRateModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeExchangeNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExchangeNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExchangeNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExchangeNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExchangeNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeExchangeNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeExchangeNamespaceMetricsAsync(
	request *DescribeExchangeNamespaceMetricsRequest,
	callback chan<- DescribeExchangeNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "exchangeNamespace",
			"function":    "describeExchangeNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeExchangeNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeExchangeNamespaceMetrics(
	request *DescribeExchangeNamespaceMetricsRequest,
) (*DescribeExchangeNamespaceMetricsResult, error) {
	callback := make(chan DescribeExchangeNamespaceMetricsAsyncResult, 1)
	go p.DescribeExchangeNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getExchangeNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExchangeNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExchangeNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExchangeNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExchangeNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetExchangeNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetExchangeNamespaceMetricsAsync(
	request *GetExchangeNamespaceMetricsRequest,
	callback chan<- GetExchangeNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "exchangeNamespace",
			"function":    "getExchangeNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getExchangeNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetExchangeNamespaceMetrics(
	request *GetExchangeNamespaceMetricsRequest,
) (*GetExchangeNamespaceMetricsResult, error) {
	callback := make(chan GetExchangeNamespaceMetricsAsyncResult, 1)
	go p.GetExchangeNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeExperienceStatusMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExperienceStatusMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceStatusMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceStatusMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceStatusMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeExperienceStatusMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeExperienceStatusMetricsAsync(
	request *DescribeExperienceStatusMetricsRequest,
	callback chan<- DescribeExperienceStatusMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "experienceStatus",
			"function":    "describeExperienceStatusMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeExperienceStatusMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeExperienceStatusMetrics(
	request *DescribeExperienceStatusMetricsRequest,
) (*DescribeExperienceStatusMetricsResult, error) {
	callback := make(chan DescribeExperienceStatusMetricsAsyncResult, 1)
	go p.DescribeExperienceStatusMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeExperienceExperienceModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExperienceExperienceModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceExperienceModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeExperienceExperienceModelMetricsAsync(
	request *DescribeExperienceExperienceModelMetricsRequest,
	callback chan<- DescribeExperienceExperienceModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "experienceExperienceModel",
			"function":    "describeExperienceExperienceModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeExperienceExperienceModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeExperienceExperienceModelMetrics(
	request *DescribeExperienceExperienceModelMetricsRequest,
) (*DescribeExperienceExperienceModelMetricsResult, error) {
	callback := make(chan DescribeExperienceExperienceModelMetricsAsyncResult, 1)
	go p.DescribeExperienceExperienceModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getExperienceExperienceModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExperienceExperienceModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceExperienceModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceExperienceModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceExperienceModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetExperienceExperienceModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetExperienceExperienceModelMetricsAsync(
	request *GetExperienceExperienceModelMetricsRequest,
	callback chan<- GetExperienceExperienceModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "experienceExperienceModel",
			"function":    "getExperienceExperienceModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getExperienceExperienceModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetExperienceExperienceModelMetrics(
	request *GetExperienceExperienceModelMetricsRequest,
) (*GetExperienceExperienceModelMetricsResult, error) {
	callback := make(chan GetExperienceExperienceModelMetricsAsyncResult, 1)
	go p.GetExperienceExperienceModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeExperienceNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeExperienceNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeExperienceNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeExperienceNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeExperienceNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeExperienceNamespaceMetricsAsync(
	request *DescribeExperienceNamespaceMetricsRequest,
	callback chan<- DescribeExperienceNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "experienceNamespace",
			"function":    "describeExperienceNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeExperienceNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeExperienceNamespaceMetrics(
	request *DescribeExperienceNamespaceMetricsRequest,
) (*DescribeExperienceNamespaceMetricsResult, error) {
	callback := make(chan DescribeExperienceNamespaceMetricsAsyncResult, 1)
	go p.DescribeExperienceNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getExperienceNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetExperienceNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetExperienceNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetExperienceNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetExperienceNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetExperienceNamespaceMetricsAsync(
	request *GetExperienceNamespaceMetricsRequest,
	callback chan<- GetExperienceNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "experienceNamespace",
			"function":    "getExperienceNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getExperienceNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetExperienceNamespaceMetrics(
	request *GetExperienceNamespaceMetricsRequest,
) (*GetExperienceNamespaceMetricsResult, error) {
	callback := make(chan GetExperienceNamespaceMetricsAsyncResult, 1)
	go p.GetExperienceNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeFormationFormMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormationFormMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormationFormMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormationFormMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationFormMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFormationFormMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeFormationFormMetricsAsync(
	request *DescribeFormationFormMetricsRequest,
	callback chan<- DescribeFormationFormMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "formationForm",
			"function":    "describeFormationFormMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		bodies["moldModelName"] = *request.MoldModelName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFormationFormMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeFormationFormMetrics(
	request *DescribeFormationFormMetricsRequest,
) (*DescribeFormationFormMetricsResult, error) {
	callback := make(chan DescribeFormationFormMetricsAsyncResult, 1)
	go p.DescribeFormationFormMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeFormationMoldMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormationMoldMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormationMoldMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormationMoldMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationMoldMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFormationMoldMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeFormationMoldMetricsAsync(
	request *DescribeFormationMoldMetricsRequest,
	callback chan<- DescribeFormationMoldMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "formationMold",
			"function":    "describeFormationMoldMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFormationMoldMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeFormationMoldMetrics(
	request *DescribeFormationMoldMetricsRequest,
) (*DescribeFormationMoldMetricsResult, error) {
	callback := make(chan DescribeFormationMoldMetricsAsyncResult, 1)
	go p.DescribeFormationMoldMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeFormationNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFormationNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFormationNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFormationNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFormationNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeFormationNamespaceMetricsAsync(
	request *DescribeFormationNamespaceMetricsRequest,
	callback chan<- DescribeFormationNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "formationNamespace",
			"function":    "describeFormationNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFormationNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeFormationNamespaceMetrics(
	request *DescribeFormationNamespaceMetricsRequest,
) (*DescribeFormationNamespaceMetricsResult, error) {
	callback := make(chan DescribeFormationNamespaceMetricsAsyncResult, 1)
	go p.DescribeFormationNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getFormationNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFormationNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFormationNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFormationNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormationNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFormationNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetFormationNamespaceMetricsAsync(
	request *GetFormationNamespaceMetricsRequest,
	callback chan<- GetFormationNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "formationNamespace",
			"function":    "getFormationNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFormationNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetFormationNamespaceMetrics(
	request *GetFormationNamespaceMetricsRequest,
) (*GetFormationNamespaceMetricsResult, error) {
	callback := make(chan GetFormationNamespaceMetricsAsyncResult, 1)
	go p.GetFormationNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeFriendNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeFriendNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeFriendNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeFriendNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFriendNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeFriendNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeFriendNamespaceMetricsAsync(
	request *DescribeFriendNamespaceMetricsRequest,
	callback chan<- DescribeFriendNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "friendNamespace",
			"function":    "describeFriendNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeFriendNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeFriendNamespaceMetrics(
	request *DescribeFriendNamespaceMetricsRequest,
) (*DescribeFriendNamespaceMetricsResult, error) {
	callback := make(chan DescribeFriendNamespaceMetricsAsyncResult, 1)
	go p.DescribeFriendNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getFriendNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetFriendNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetFriendNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetFriendNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFriendNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetFriendNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetFriendNamespaceMetricsAsync(
	request *GetFriendNamespaceMetricsRequest,
	callback chan<- GetFriendNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "friendNamespace",
			"function":    "getFriendNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getFriendNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetFriendNamespaceMetrics(
	request *GetFriendNamespaceMetricsRequest,
) (*GetFriendNamespaceMetricsResult, error) {
	callback := make(chan GetFriendNamespaceMetricsAsyncResult, 1)
	go p.GetFriendNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeInboxNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInboxNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInboxNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInboxNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInboxNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeInboxNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeInboxNamespaceMetricsAsync(
	request *DescribeInboxNamespaceMetricsRequest,
	callback chan<- DescribeInboxNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inboxNamespace",
			"function":    "describeInboxNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeInboxNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeInboxNamespaceMetrics(
	request *DescribeInboxNamespaceMetricsRequest,
) (*DescribeInboxNamespaceMetricsResult, error) {
	callback := make(chan DescribeInboxNamespaceMetricsAsyncResult, 1)
	go p.DescribeInboxNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getInboxNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInboxNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInboxNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInboxNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInboxNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetInboxNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetInboxNamespaceMetricsAsync(
	request *GetInboxNamespaceMetricsRequest,
	callback chan<- GetInboxNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inboxNamespace",
			"function":    "getInboxNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getInboxNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetInboxNamespaceMetrics(
	request *GetInboxNamespaceMetricsRequest,
) (*GetInboxNamespaceMetricsResult, error) {
	callback := make(chan GetInboxNamespaceMetricsAsyncResult, 1)
	go p.GetInboxNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeInventoryItemSetMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoryItemSetMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryItemSetMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryItemSetMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryItemSetMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeInventoryItemSetMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeInventoryItemSetMetricsAsync(
	request *DescribeInventoryItemSetMetricsRequest,
	callback chan<- DescribeInventoryItemSetMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inventoryItemSet",
			"function":    "describeInventoryItemSetMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		bodies["inventoryName"] = *request.InventoryName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeInventoryItemSetMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeInventoryItemSetMetrics(
	request *DescribeInventoryItemSetMetricsRequest,
) (*DescribeInventoryItemSetMetricsResult, error) {
	callback := make(chan DescribeInventoryItemSetMetricsAsyncResult, 1)
	go p.DescribeInventoryItemSetMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeInventoryInventoryMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoryInventoryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryInventoryMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryInventoryMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryInventoryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeInventoryInventoryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeInventoryInventoryMetricsAsync(
	request *DescribeInventoryInventoryMetricsRequest,
	callback chan<- DescribeInventoryInventoryMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inventoryInventory",
			"function":    "describeInventoryInventoryMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeInventoryInventoryMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeInventoryInventoryMetrics(
	request *DescribeInventoryInventoryMetricsRequest,
) (*DescribeInventoryInventoryMetricsResult, error) {
	callback := make(chan DescribeInventoryInventoryMetricsAsyncResult, 1)
	go p.DescribeInventoryInventoryMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeInventoryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeInventoryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeInventoryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeInventoryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeInventoryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeInventoryNamespaceMetricsAsync(
	request *DescribeInventoryNamespaceMetricsRequest,
	callback chan<- DescribeInventoryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inventoryNamespace",
			"function":    "describeInventoryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeInventoryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeInventoryNamespaceMetrics(
	request *DescribeInventoryNamespaceMetricsRequest,
) (*DescribeInventoryNamespaceMetricsResult, error) {
	callback := make(chan DescribeInventoryNamespaceMetricsAsyncResult, 1)
	go p.DescribeInventoryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getInventoryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetInventoryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetInventoryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetInventoryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetInventoryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetInventoryNamespaceMetricsAsync(
	request *GetInventoryNamespaceMetricsRequest,
	callback chan<- GetInventoryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "inventoryNamespace",
			"function":    "getInventoryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getInventoryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetInventoryNamespaceMetrics(
	request *GetInventoryNamespaceMetricsRequest,
) (*GetInventoryNamespaceMetricsResult, error) {
	callback := make(chan GetInventoryNamespaceMetricsAsyncResult, 1)
	go p.GetInventoryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeKeyNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeKeyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeKeyNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeKeyNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeKeyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeKeyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeKeyNamespaceMetricsAsync(
	request *DescribeKeyNamespaceMetricsRequest,
	callback chan<- DescribeKeyNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "keyNamespace",
			"function":    "describeKeyNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeKeyNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeKeyNamespaceMetrics(
	request *DescribeKeyNamespaceMetricsRequest,
) (*DescribeKeyNamespaceMetricsResult, error) {
	callback := make(chan DescribeKeyNamespaceMetricsAsyncResult, 1)
	go p.DescribeKeyNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getKeyNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetKeyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetKeyNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetKeyNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetKeyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetKeyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetKeyNamespaceMetricsAsync(
	request *GetKeyNamespaceMetricsRequest,
	callback chan<- GetKeyNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "keyNamespace",
			"function":    "getKeyNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getKeyNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetKeyNamespaceMetrics(
	request *GetKeyNamespaceMetricsRequest,
) (*GetKeyNamespaceMetricsResult, error) {
	callback := make(chan GetKeyNamespaceMetricsAsyncResult, 1)
	go p.GetKeyNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeLimitCounterMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLimitCounterMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLimitCounterMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLimitCounterMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitCounterMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLimitCounterMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeLimitCounterMetricsAsync(
	request *DescribeLimitCounterMetricsRequest,
	callback chan<- DescribeLimitCounterMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "limitCounter",
			"function":    "describeLimitCounterMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.LimitName != nil && *request.LimitName != "" {
		bodies["limitName"] = *request.LimitName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeLimitCounterMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeLimitCounterMetrics(
	request *DescribeLimitCounterMetricsRequest,
) (*DescribeLimitCounterMetricsResult, error) {
	callback := make(chan DescribeLimitCounterMetricsAsyncResult, 1)
	go p.DescribeLimitCounterMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeLimitLimitModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLimitLimitModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLimitLimitModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLimitLimitModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitLimitModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLimitLimitModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeLimitLimitModelMetricsAsync(
	request *DescribeLimitLimitModelMetricsRequest,
	callback chan<- DescribeLimitLimitModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "limitLimitModel",
			"function":    "describeLimitLimitModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeLimitLimitModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeLimitLimitModelMetrics(
	request *DescribeLimitLimitModelMetricsRequest,
) (*DescribeLimitLimitModelMetricsResult, error) {
	callback := make(chan DescribeLimitLimitModelMetricsAsyncResult, 1)
	go p.DescribeLimitLimitModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getLimitLimitModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLimitLimitModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLimitLimitModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLimitLimitModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitLimitModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLimitLimitModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetLimitLimitModelMetricsAsync(
	request *GetLimitLimitModelMetricsRequest,
	callback chan<- GetLimitLimitModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "limitLimitModel",
			"function":    "getLimitLimitModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.LimitName != nil && *request.LimitName != "" {
		bodies["limitName"] = *request.LimitName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getLimitLimitModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetLimitLimitModelMetrics(
	request *GetLimitLimitModelMetricsRequest,
) (*GetLimitLimitModelMetricsResult, error) {
	callback := make(chan GetLimitLimitModelMetricsAsyncResult, 1)
	go p.GetLimitLimitModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeLimitNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLimitNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLimitNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLimitNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLimitNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeLimitNamespaceMetricsAsync(
	request *DescribeLimitNamespaceMetricsRequest,
	callback chan<- DescribeLimitNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "limitNamespace",
			"function":    "describeLimitNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeLimitNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeLimitNamespaceMetrics(
	request *DescribeLimitNamespaceMetricsRequest,
) (*DescribeLimitNamespaceMetricsResult, error) {
	callback := make(chan DescribeLimitNamespaceMetricsAsyncResult, 1)
	go p.DescribeLimitNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getLimitNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLimitNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLimitNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLimitNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLimitNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetLimitNamespaceMetricsAsync(
	request *GetLimitNamespaceMetricsRequest,
	callback chan<- GetLimitNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "limitNamespace",
			"function":    "getLimitNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getLimitNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetLimitNamespaceMetrics(
	request *GetLimitNamespaceMetricsRequest,
) (*GetLimitNamespaceMetricsResult, error) {
	callback := make(chan GetLimitNamespaceMetricsAsyncResult, 1)
	go p.GetLimitNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeLotteryLotteryMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLotteryLotteryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryLotteryMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryLotteryMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryLotteryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLotteryLotteryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeLotteryLotteryMetricsAsync(
	request *DescribeLotteryLotteryMetricsRequest,
	callback chan<- DescribeLotteryLotteryMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "lotteryLottery",
			"function":    "describeLotteryLotteryMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeLotteryLotteryMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeLotteryLotteryMetrics(
	request *DescribeLotteryLotteryMetricsRequest,
) (*DescribeLotteryLotteryMetricsResult, error) {
	callback := make(chan DescribeLotteryLotteryMetricsAsyncResult, 1)
	go p.DescribeLotteryLotteryMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getLotteryLotteryMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLotteryLotteryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryLotteryMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryLotteryMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryLotteryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLotteryLotteryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetLotteryLotteryMetricsAsync(
	request *GetLotteryLotteryMetricsRequest,
	callback chan<- GetLotteryLotteryMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "lotteryLottery",
			"function":    "getLotteryLotteryMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		bodies["lotteryName"] = *request.LotteryName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getLotteryLotteryMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetLotteryLotteryMetrics(
	request *GetLotteryLotteryMetricsRequest,
) (*GetLotteryLotteryMetricsResult, error) {
	callback := make(chan GetLotteryLotteryMetricsAsyncResult, 1)
	go p.GetLotteryLotteryMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeLotteryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeLotteryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeLotteryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeLotteryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeLotteryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeLotteryNamespaceMetricsAsync(
	request *DescribeLotteryNamespaceMetricsRequest,
	callback chan<- DescribeLotteryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "lotteryNamespace",
			"function":    "describeLotteryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeLotteryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeLotteryNamespaceMetrics(
	request *DescribeLotteryNamespaceMetricsRequest,
) (*DescribeLotteryNamespaceMetricsResult, error) {
	callback := make(chan DescribeLotteryNamespaceMetricsAsyncResult, 1)
	go p.DescribeLotteryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getLotteryNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetLotteryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetLotteryNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetLotteryNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetLotteryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetLotteryNamespaceMetricsAsync(
	request *GetLotteryNamespaceMetricsRequest,
	callback chan<- GetLotteryNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "lotteryNamespace",
			"function":    "getLotteryNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getLotteryNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetLotteryNamespaceMetrics(
	request *GetLotteryNamespaceMetricsRequest,
) (*GetLotteryNamespaceMetricsResult, error) {
	callback := make(chan GetLotteryNamespaceMetricsAsyncResult, 1)
	go p.GetLotteryNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMatchmakingNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMatchmakingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMatchmakingNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMatchmakingNamespaceMetricsAsync(
	request *DescribeMatchmakingNamespaceMetricsRequest,
	callback chan<- DescribeMatchmakingNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "matchmakingNamespace",
			"function":    "describeMatchmakingNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMatchmakingNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMatchmakingNamespaceMetrics(
	request *DescribeMatchmakingNamespaceMetricsRequest,
) (*DescribeMatchmakingNamespaceMetricsResult, error) {
	callback := make(chan DescribeMatchmakingNamespaceMetricsAsyncResult, 1)
	go p.DescribeMatchmakingNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getMatchmakingNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMatchmakingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMatchmakingNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMatchmakingNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMatchmakingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMatchmakingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetMatchmakingNamespaceMetricsAsync(
	request *GetMatchmakingNamespaceMetricsRequest,
	callback chan<- GetMatchmakingNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "matchmakingNamespace",
			"function":    "getMatchmakingNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getMatchmakingNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetMatchmakingNamespaceMetrics(
	request *GetMatchmakingNamespaceMetricsRequest,
) (*GetMatchmakingNamespaceMetricsResult, error) {
	callback := make(chan GetMatchmakingNamespaceMetricsAsyncResult, 1)
	go p.GetMatchmakingNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMissionCounterMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionCounterMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionCounterMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionCounterMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionCounterMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionCounterMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMissionCounterMetricsAsync(
	request *DescribeMissionCounterMetricsRequest,
	callback chan<- DescribeMissionCounterMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "missionCounter",
			"function":    "describeMissionCounterMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMissionCounterMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMissionCounterMetrics(
	request *DescribeMissionCounterMetricsRequest,
) (*DescribeMissionCounterMetricsResult, error) {
	callback := make(chan DescribeMissionCounterMetricsAsyncResult, 1)
	go p.DescribeMissionCounterMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMissionMissionGroupModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionMissionGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionMissionGroupModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMissionMissionGroupModelMetricsAsync(
	request *DescribeMissionMissionGroupModelMetricsRequest,
	callback chan<- DescribeMissionMissionGroupModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "missionMissionGroupModel",
			"function":    "describeMissionMissionGroupModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMissionMissionGroupModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMissionMissionGroupModelMetrics(
	request *DescribeMissionMissionGroupModelMetricsRequest,
) (*DescribeMissionMissionGroupModelMetricsResult, error) {
	callback := make(chan DescribeMissionMissionGroupModelMetricsAsyncResult, 1)
	go p.DescribeMissionMissionGroupModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getMissionMissionGroupModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionMissionGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionMissionGroupModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionMissionGroupModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionMissionGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionMissionGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetMissionMissionGroupModelMetricsAsync(
	request *GetMissionMissionGroupModelMetricsRequest,
	callback chan<- GetMissionMissionGroupModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "missionMissionGroupModel",
			"function":    "getMissionMissionGroupModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		bodies["missionGroupName"] = *request.MissionGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getMissionMissionGroupModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetMissionMissionGroupModelMetrics(
	request *GetMissionMissionGroupModelMetricsRequest,
) (*GetMissionMissionGroupModelMetricsResult, error) {
	callback := make(chan GetMissionMissionGroupModelMetricsAsyncResult, 1)
	go p.GetMissionMissionGroupModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMissionNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMissionNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMissionNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMissionNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMissionNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMissionNamespaceMetricsAsync(
	request *DescribeMissionNamespaceMetricsRequest,
	callback chan<- DescribeMissionNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "missionNamespace",
			"function":    "describeMissionNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMissionNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMissionNamespaceMetrics(
	request *DescribeMissionNamespaceMetricsRequest,
) (*DescribeMissionNamespaceMetricsResult, error) {
	callback := make(chan DescribeMissionNamespaceMetricsAsyncResult, 1)
	go p.DescribeMissionNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getMissionNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMissionNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMissionNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMissionNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMissionNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetMissionNamespaceMetricsAsync(
	request *GetMissionNamespaceMetricsRequest,
	callback chan<- GetMissionNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "missionNamespace",
			"function":    "getMissionNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getMissionNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetMissionNamespaceMetrics(
	request *GetMissionNamespaceMetricsRequest,
) (*GetMissionNamespaceMetricsResult, error) {
	callback := make(chan GetMissionNamespaceMetricsAsyncResult, 1)
	go p.GetMissionNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMoneyWalletMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoneyWalletMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoneyWalletMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoneyWalletMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyWalletMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMoneyWalletMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMoneyWalletMetricsAsync(
	request *DescribeMoneyWalletMetricsRequest,
	callback chan<- DescribeMoneyWalletMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "moneyWallet",
			"function":    "describeMoneyWalletMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMoneyWalletMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMoneyWalletMetrics(
	request *DescribeMoneyWalletMetricsRequest,
) (*DescribeMoneyWalletMetricsResult, error) {
	callback := make(chan DescribeMoneyWalletMetricsAsyncResult, 1)
	go p.DescribeMoneyWalletMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMoneyReceiptMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoneyReceiptMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoneyReceiptMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoneyReceiptMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyReceiptMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMoneyReceiptMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMoneyReceiptMetricsAsync(
	request *DescribeMoneyReceiptMetricsRequest,
	callback chan<- DescribeMoneyReceiptMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "moneyReceipt",
			"function":    "describeMoneyReceiptMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMoneyReceiptMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMoneyReceiptMetrics(
	request *DescribeMoneyReceiptMetricsRequest,
) (*DescribeMoneyReceiptMetricsResult, error) {
	callback := make(chan DescribeMoneyReceiptMetricsAsyncResult, 1)
	go p.DescribeMoneyReceiptMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeMoneyNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeMoneyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeMoneyNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeMoneyNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeMoneyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeMoneyNamespaceMetricsAsync(
	request *DescribeMoneyNamespaceMetricsRequest,
	callback chan<- DescribeMoneyNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "moneyNamespace",
			"function":    "describeMoneyNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeMoneyNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeMoneyNamespaceMetrics(
	request *DescribeMoneyNamespaceMetricsRequest,
) (*DescribeMoneyNamespaceMetricsResult, error) {
	callback := make(chan DescribeMoneyNamespaceMetricsAsyncResult, 1)
	go p.DescribeMoneyNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getMoneyNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetMoneyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetMoneyNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetMoneyNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoneyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetMoneyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetMoneyNamespaceMetricsAsync(
	request *GetMoneyNamespaceMetricsRequest,
	callback chan<- GetMoneyNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "moneyNamespace",
			"function":    "getMoneyNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getMoneyNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetMoneyNamespaceMetrics(
	request *GetMoneyNamespaceMetricsRequest,
) (*GetMoneyNamespaceMetricsResult, error) {
	callback := make(chan GetMoneyNamespaceMetricsAsyncResult, 1)
	go p.GetMoneyNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeQuestQuestModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestQuestModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestQuestModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestQuestModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestQuestModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeQuestQuestModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeQuestQuestModelMetricsAsync(
	request *DescribeQuestQuestModelMetricsRequest,
	callback chan<- DescribeQuestQuestModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questQuestModel",
			"function":    "describeQuestQuestModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		bodies["questGroupName"] = *request.QuestGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeQuestQuestModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeQuestQuestModelMetrics(
	request *DescribeQuestQuestModelMetricsRequest,
) (*DescribeQuestQuestModelMetricsResult, error) {
	callback := make(chan DescribeQuestQuestModelMetricsAsyncResult, 1)
	go p.DescribeQuestQuestModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getQuestQuestModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestQuestModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestQuestModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestQuestModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestQuestModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetQuestQuestModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetQuestQuestModelMetricsAsync(
	request *GetQuestQuestModelMetricsRequest,
	callback chan<- GetQuestQuestModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questQuestModel",
			"function":    "getQuestQuestModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		bodies["questGroupName"] = *request.QuestGroupName
	}
	if request.QuestName != nil && *request.QuestName != "" {
		bodies["questName"] = *request.QuestName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getQuestQuestModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetQuestQuestModelMetrics(
	request *GetQuestQuestModelMetricsRequest,
) (*GetQuestQuestModelMetricsResult, error) {
	callback := make(chan GetQuestQuestModelMetricsAsyncResult, 1)
	go p.GetQuestQuestModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeQuestQuestGroupModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestQuestGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestQuestGroupModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeQuestQuestGroupModelMetricsAsync(
	request *DescribeQuestQuestGroupModelMetricsRequest,
	callback chan<- DescribeQuestQuestGroupModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questQuestGroupModel",
			"function":    "describeQuestQuestGroupModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeQuestQuestGroupModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeQuestQuestGroupModelMetrics(
	request *DescribeQuestQuestGroupModelMetricsRequest,
) (*DescribeQuestQuestGroupModelMetricsResult, error) {
	callback := make(chan DescribeQuestQuestGroupModelMetricsAsyncResult, 1)
	go p.DescribeQuestQuestGroupModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getQuestQuestGroupModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestQuestGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestQuestGroupModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestQuestGroupModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestQuestGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetQuestQuestGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetQuestQuestGroupModelMetricsAsync(
	request *GetQuestQuestGroupModelMetricsRequest,
	callback chan<- GetQuestQuestGroupModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questQuestGroupModel",
			"function":    "getQuestQuestGroupModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		bodies["questGroupName"] = *request.QuestGroupName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getQuestQuestGroupModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetQuestQuestGroupModelMetrics(
	request *GetQuestQuestGroupModelMetricsRequest,
) (*GetQuestQuestGroupModelMetricsResult, error) {
	callback := make(chan GetQuestQuestGroupModelMetricsAsyncResult, 1)
	go p.GetQuestQuestGroupModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeQuestNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeQuestNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeQuestNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeQuestNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeQuestNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeQuestNamespaceMetricsAsync(
	request *DescribeQuestNamespaceMetricsRequest,
	callback chan<- DescribeQuestNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questNamespace",
			"function":    "describeQuestNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeQuestNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeQuestNamespaceMetrics(
	request *DescribeQuestNamespaceMetricsRequest,
) (*DescribeQuestNamespaceMetricsResult, error) {
	callback := make(chan DescribeQuestNamespaceMetricsAsyncResult, 1)
	go p.DescribeQuestNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getQuestNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetQuestNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetQuestNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetQuestNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetQuestNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetQuestNamespaceMetricsAsync(
	request *GetQuestNamespaceMetricsRequest,
	callback chan<- GetQuestNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "questNamespace",
			"function":    "getQuestNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getQuestNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetQuestNamespaceMetrics(
	request *GetQuestNamespaceMetricsRequest,
) (*GetQuestNamespaceMetricsResult, error) {
	callback := make(chan GetQuestNamespaceMetricsAsyncResult, 1)
	go p.GetQuestNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeRankingCategoryModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRankingCategoryModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingCategoryModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingCategoryModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingCategoryModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeRankingCategoryModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeRankingCategoryModelMetricsAsync(
	request *DescribeRankingCategoryModelMetricsRequest,
	callback chan<- DescribeRankingCategoryModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "rankingCategoryModel",
			"function":    "describeRankingCategoryModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeRankingCategoryModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeRankingCategoryModelMetrics(
	request *DescribeRankingCategoryModelMetricsRequest,
) (*DescribeRankingCategoryModelMetricsResult, error) {
	callback := make(chan DescribeRankingCategoryModelMetricsAsyncResult, 1)
	go p.DescribeRankingCategoryModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getRankingCategoryModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRankingCategoryModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingCategoryModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingCategoryModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingCategoryModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetRankingCategoryModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetRankingCategoryModelMetricsAsync(
	request *GetRankingCategoryModelMetricsRequest,
	callback chan<- GetRankingCategoryModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "rankingCategoryModel",
			"function":    "getRankingCategoryModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
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
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getRankingCategoryModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetRankingCategoryModelMetrics(
	request *GetRankingCategoryModelMetricsRequest,
) (*GetRankingCategoryModelMetricsResult, error) {
	callback := make(chan GetRankingCategoryModelMetricsAsyncResult, 1)
	go p.GetRankingCategoryModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeRankingNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeRankingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeRankingNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeRankingNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeRankingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeRankingNamespaceMetricsAsync(
	request *DescribeRankingNamespaceMetricsRequest,
	callback chan<- DescribeRankingNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "rankingNamespace",
			"function":    "describeRankingNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeRankingNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeRankingNamespaceMetrics(
	request *DescribeRankingNamespaceMetricsRequest,
) (*DescribeRankingNamespaceMetricsResult, error) {
	callback := make(chan DescribeRankingNamespaceMetricsAsyncResult, 1)
	go p.DescribeRankingNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getRankingNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetRankingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetRankingNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetRankingNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetRankingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetRankingNamespaceMetricsAsync(
	request *GetRankingNamespaceMetricsRequest,
	callback chan<- GetRankingNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "rankingNamespace",
			"function":    "getRankingNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getRankingNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetRankingNamespaceMetrics(
	request *GetRankingNamespaceMetricsRequest,
) (*GetRankingNamespaceMetricsResult, error) {
	callback := make(chan GetRankingNamespaceMetricsAsyncResult, 1)
	go p.GetRankingNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeShowcaseDisplayItemMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcaseDisplayItemMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcaseDisplayItemMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeShowcaseDisplayItemMetricsAsync(
	request *DescribeShowcaseDisplayItemMetricsRequest,
	callback chan<- DescribeShowcaseDisplayItemMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseDisplayItem",
			"function":    "describeShowcaseDisplayItemMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		bodies["showcaseName"] = *request.ShowcaseName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeShowcaseDisplayItemMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeShowcaseDisplayItemMetrics(
	request *DescribeShowcaseDisplayItemMetricsRequest,
) (*DescribeShowcaseDisplayItemMetricsResult, error) {
	callback := make(chan DescribeShowcaseDisplayItemMetricsAsyncResult, 1)
	go p.DescribeShowcaseDisplayItemMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getShowcaseDisplayItemMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseDisplayItemMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseDisplayItemMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseDisplayItemMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseDisplayItemMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetShowcaseDisplayItemMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetShowcaseDisplayItemMetricsAsync(
	request *GetShowcaseDisplayItemMetricsRequest,
	callback chan<- GetShowcaseDisplayItemMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseDisplayItem",
			"function":    "getShowcaseDisplayItemMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		bodies["showcaseName"] = *request.ShowcaseName
	}
	if request.DisplayItemId != nil && *request.DisplayItemId != "" {
		bodies["displayItemId"] = *request.DisplayItemId
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getShowcaseDisplayItemMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetShowcaseDisplayItemMetrics(
	request *GetShowcaseDisplayItemMetricsRequest,
) (*GetShowcaseDisplayItemMetricsResult, error) {
	callback := make(chan GetShowcaseDisplayItemMetricsAsyncResult, 1)
	go p.GetShowcaseDisplayItemMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeShowcaseShowcaseMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcaseShowcaseMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcaseShowcaseMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeShowcaseShowcaseMetricsAsync(
	request *DescribeShowcaseShowcaseMetricsRequest,
	callback chan<- DescribeShowcaseShowcaseMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseShowcase",
			"function":    "describeShowcaseShowcaseMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeShowcaseShowcaseMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeShowcaseShowcaseMetrics(
	request *DescribeShowcaseShowcaseMetricsRequest,
) (*DescribeShowcaseShowcaseMetricsResult, error) {
	callback := make(chan DescribeShowcaseShowcaseMetricsAsyncResult, 1)
	go p.DescribeShowcaseShowcaseMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getShowcaseShowcaseMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseShowcaseMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseShowcaseMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseShowcaseMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseShowcaseMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetShowcaseShowcaseMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetShowcaseShowcaseMetricsAsync(
	request *GetShowcaseShowcaseMetricsRequest,
	callback chan<- GetShowcaseShowcaseMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseShowcase",
			"function":    "getShowcaseShowcaseMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		bodies["showcaseName"] = *request.ShowcaseName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getShowcaseShowcaseMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetShowcaseShowcaseMetrics(
	request *GetShowcaseShowcaseMetricsRequest,
) (*GetShowcaseShowcaseMetricsResult, error) {
	callback := make(chan GetShowcaseShowcaseMetricsAsyncResult, 1)
	go p.GetShowcaseShowcaseMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeShowcaseNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeShowcaseNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeShowcaseNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeShowcaseNamespaceMetricsAsync(
	request *DescribeShowcaseNamespaceMetricsRequest,
	callback chan<- DescribeShowcaseNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseNamespace",
			"function":    "describeShowcaseNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeShowcaseNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeShowcaseNamespaceMetrics(
	request *DescribeShowcaseNamespaceMetricsRequest,
) (*DescribeShowcaseNamespaceMetricsResult, error) {
	callback := make(chan DescribeShowcaseNamespaceMetricsAsyncResult, 1)
	go p.DescribeShowcaseNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getShowcaseNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetShowcaseNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetShowcaseNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetShowcaseNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetShowcaseNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetShowcaseNamespaceMetricsAsync(
	request *GetShowcaseNamespaceMetricsRequest,
	callback chan<- GetShowcaseNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "showcaseNamespace",
			"function":    "getShowcaseNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getShowcaseNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetShowcaseNamespaceMetrics(
	request *GetShowcaseNamespaceMetricsRequest,
) (*GetShowcaseNamespaceMetricsResult, error) {
	callback := make(chan GetShowcaseNamespaceMetricsAsyncResult, 1)
	go p.GetShowcaseNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeStaminaStaminaModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminaStaminaModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaStaminaModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeStaminaStaminaModelMetricsAsync(
	request *DescribeStaminaStaminaModelMetricsRequest,
	callback chan<- DescribeStaminaStaminaModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "staminaStaminaModel",
			"function":    "describeStaminaStaminaModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeStaminaStaminaModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeStaminaStaminaModelMetrics(
	request *DescribeStaminaStaminaModelMetricsRequest,
) (*DescribeStaminaStaminaModelMetricsResult, error) {
	callback := make(chan DescribeStaminaStaminaModelMetricsAsyncResult, 1)
	go p.DescribeStaminaStaminaModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getStaminaStaminaModelMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaStaminaModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaStaminaModelMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaStaminaModelMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaStaminaModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStaminaStaminaModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetStaminaStaminaModelMetricsAsync(
	request *GetStaminaStaminaModelMetricsRequest,
	callback chan<- GetStaminaStaminaModelMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "staminaStaminaModel",
			"function":    "getStaminaStaminaModelMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		bodies["staminaName"] = *request.StaminaName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getStaminaStaminaModelMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetStaminaStaminaModelMetrics(
	request *GetStaminaStaminaModelMetricsRequest,
) (*GetStaminaStaminaModelMetricsResult, error) {
	callback := make(chan GetStaminaStaminaModelMetricsAsyncResult, 1)
	go p.GetStaminaStaminaModelMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) describeStaminaNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- DescribeStaminaNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- DescribeStaminaNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result DescribeStaminaNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- DescribeStaminaNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) DescribeStaminaNamespaceMetricsAsync(
	request *DescribeStaminaNamespaceMetricsRequest,
	callback chan<- DescribeStaminaNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "staminaNamespace",
			"function":    "describeStaminaNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.describeStaminaNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) DescribeStaminaNamespaceMetrics(
	request *DescribeStaminaNamespaceMetricsRequest,
) (*DescribeStaminaNamespaceMetricsResult, error) {
	callback := make(chan DescribeStaminaNamespaceMetricsAsyncResult, 1)
	go p.DescribeStaminaNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}

func (p Gs2WatchWebSocketClient) getStaminaNamespaceMetricsAsyncHandler(
	job *core.WebSocketNetworkJob,
	callback chan<- GetStaminaNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := p.Session.Send(
		job,
		false,
	)
	if err != nil {
		callback <- GetStaminaNamespaceMetricsAsyncResult{
			err: err,
		}
		return
	}
	asyncResult := <-internalCallback
	var result GetStaminaNamespaceMetricsResult
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	if asyncResult.Err != nil {
	}
	callback <- GetStaminaNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchWebSocketClient) GetStaminaNamespaceMetricsAsync(
	request *GetStaminaNamespaceMetricsRequest,
	callback chan<- GetStaminaNamespaceMetricsAsyncResult,
) {
	requestId := core.WebSocketRequestId(uuid.New().String())
	var bodies = core.WebSocketBodies{
		"x_gs2": map[string]interface{}{
			"service":     "watch",
			"component":   "staminaNamespace",
			"function":    "getStaminaNamespaceMetrics",
			"contentType": "application/json",
			"requestId":   requestId,
		},
	}
	for k, v := range p.Session.CreateAuthorizationHeader() {
		bodies[k] = v
	}
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		bodies["namespaceName"] = *request.NamespaceName
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}
	if request.DryRun != nil {
		if *request.DryRun {
			bodies["xGs2DryRun"] = "true"
		} else {
			bodies["xGs2DryRun"] = "false"
		}
	}

	go p.getStaminaNamespaceMetricsAsyncHandler(
		&core.WebSocketNetworkJob{
			RequestId: requestId,
			Bodies:    bodies,
		},
		callback,
	)
}

func (p Gs2WatchWebSocketClient) GetStaminaNamespaceMetrics(
	request *GetStaminaNamespaceMetricsRequest,
) (*GetStaminaNamespaceMetricsResult, error) {
	callback := make(chan GetStaminaNamespaceMetricsAsyncResult, 1)
	go p.GetStaminaNamespaceMetricsAsync(
		request,
		callback,
	)
	asyncResult := <-callback
	return asyncResult.result, asyncResult.err
}
