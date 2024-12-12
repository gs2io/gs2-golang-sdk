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
	"strings"

	"github.com/gs2io/gs2-golang-sdk/core"
)

type Gs2WatchRestClient struct {
	Session *core.Gs2RestSession
}

func (p Gs2WatchRestClient) CreateAuthorizedHeaders() map[string]string {
	return p.Session.CreateAuthorizationHeader()
}

func getChartAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetChartAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetChartAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2WatchRestClient) GetChartAsync(
	request *GetChartRequest,
	callback chan<- GetChartAsyncResult,
) {
	path := "/chart/{measure}"
	if request.Measure != nil && *request.Measure != "" {
		path = strings.ReplaceAll(path, "{measure}", core.ToString(*request.Measure))
	} else {
		path = strings.ReplaceAll(path, "{measure}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getChartAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetChart(
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

func getDistributionAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetDistributionAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetDistributionAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDistributionAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDistributionAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetDistributionAsync(
	request *GetDistributionRequest,
	callback chan<- GetDistributionAsyncResult,
) {
	path := "/distribution/{measure}"
	if request.Measure != nil && *request.Measure != "" {
		path = strings.ReplaceAll(path, "{measure}", core.ToString(*request.Measure))
	} else {
		path = strings.ReplaceAll(path, "{measure}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
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

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getDistributionAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetDistribution(
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

func getCumulativeAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetCumulativeAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetCumulativeAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2WatchRestClient) GetCumulativeAsync(
	request *GetCumulativeRequest,
	callback chan<- GetCumulativeAsyncResult,
) {
	path := "/cumulative/{name}"
	if request.Name != nil && *request.Name != "" {
		path = strings.ReplaceAll(path, "{name}", core.ToString(*request.Name))
	} else {
		path = strings.ReplaceAll(path, "{name}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ResourceGrn != nil && *request.ResourceGrn != "" {
		bodies["resourceGrn"] = *request.ResourceGrn
	}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getCumulativeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetCumulative(
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

func describeBillingActivitiesAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeBillingActivitiesAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeBillingActivitiesAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2WatchRestClient) DescribeBillingActivitiesAsync(
	request *DescribeBillingActivitiesRequest,
	callback chan<- DescribeBillingActivitiesAsyncResult,
) {
	path := "/billingActivity/{year}/{month}"
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}
	if request.Month != nil {
		path = strings.ReplaceAll(path, "{month}", core.ToString(*request.Month))
	} else {
		path = strings.ReplaceAll(path, "{month}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.Service != nil {
		queryStrings["service"] = core.ToString(*request.Service)
	}
	if request.PageToken != nil {
		queryStrings["pageToken"] = core.ToString(*request.PageToken)
	}
	if request.Limit != nil {
		queryStrings["limit"] = core.ToString(*request.Limit)
	}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeBillingActivitiesAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeBillingActivities(
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

func getBillingActivityAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetBillingActivityAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetBillingActivityAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
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

func (p Gs2WatchRestClient) GetBillingActivityAsync(
	request *GetBillingActivityRequest,
	callback chan<- GetBillingActivityAsyncResult,
) {
	path := "/billingActivity/{year}/{month}/{service}/{activityType}"
	if request.Year != nil {
		path = strings.ReplaceAll(path, "{year}", core.ToString(*request.Year))
	} else {
		path = strings.ReplaceAll(path, "{year}", "null")
	}
	if request.Month != nil {
		path = strings.ReplaceAll(path, "{month}", core.ToString(*request.Month))
	} else {
		path = strings.ReplaceAll(path, "{month}", "null")
	}
	if request.Service != nil && *request.Service != "" {
		path = strings.ReplaceAll(path, "{service}", core.ToString(*request.Service))
	} else {
		path = strings.ReplaceAll(path, "{service}", "null")
	}
	if request.ActivityType != nil && *request.ActivityType != "" {
		path = strings.ReplaceAll(path, "{activityType}", core.ToString(*request.ActivityType))
	} else {
		path = strings.ReplaceAll(path, "{activityType}", "null")
	}

	replacer := strings.NewReplacer()
	var bodies = core.Bodies{}
	if request.ContextStack != nil {
		bodies["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getBillingActivityAsyncHandler(
		p,
		&core.NetworkJob{
			Url:     p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:  core.Post,
			Headers: headers,
			Bodies:  bodies,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetBillingActivity(
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

func getGeneralMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetGeneralMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetGeneralMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetGeneralMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetGeneralMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetGeneralMetricsAsync(
	request *GetGeneralMetricsRequest,
	callback chan<- GetGeneralMetricsAsyncResult,
) {
	path := "/metrics/general"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getGeneralMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetGeneralMetrics(
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

func describeAccountNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeAccountNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeAccountNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeAccountNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeAccountNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeAccountNamespaceMetricsAsync(
	request *DescribeAccountNamespaceMetricsRequest,
	callback chan<- DescribeAccountNamespaceMetricsAsyncResult,
) {
	path := "/metrics/account/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeAccountNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeAccountNamespaceMetrics(
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

func getAccountNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetAccountNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetAccountNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetAccountNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetAccountNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetAccountNamespaceMetricsAsync(
	request *GetAccountNamespaceMetricsRequest,
	callback chan<- GetAccountNamespaceMetricsAsyncResult,
) {
	path := "/metrics/account/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getAccountNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetAccountNamespaceMetrics(
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

func describeChatNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeChatNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeChatNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeChatNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeChatNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeChatNamespaceMetricsAsync(
	request *DescribeChatNamespaceMetricsRequest,
	callback chan<- DescribeChatNamespaceMetricsAsyncResult,
) {
	path := "/metrics/chat/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeChatNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeChatNamespaceMetrics(
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

func getChatNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetChatNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetChatNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetChatNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetChatNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetChatNamespaceMetricsAsync(
	request *GetChatNamespaceMetricsRequest,
	callback chan<- GetChatNamespaceMetricsAsyncResult,
) {
	path := "/metrics/chat/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getChatNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetChatNamespaceMetrics(
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

func describeDatastoreNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDatastoreNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDatastoreNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeDatastoreNamespaceMetricsAsync(
	request *DescribeDatastoreNamespaceMetricsRequest,
	callback chan<- DescribeDatastoreNamespaceMetricsAsyncResult,
) {
	path := "/metrics/datastore/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeDatastoreNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeDatastoreNamespaceMetrics(
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

func getDatastoreNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetDatastoreNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetDatastoreNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDatastoreNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDatastoreNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetDatastoreNamespaceMetricsAsync(
	request *GetDatastoreNamespaceMetricsRequest,
	callback chan<- GetDatastoreNamespaceMetricsAsyncResult,
) {
	path := "/metrics/datastore/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getDatastoreNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetDatastoreNamespaceMetrics(
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

func describeDictionaryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeDictionaryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeDictionaryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeDictionaryNamespaceMetricsAsync(
	request *DescribeDictionaryNamespaceMetricsRequest,
	callback chan<- DescribeDictionaryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/dictionary/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeDictionaryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeDictionaryNamespaceMetrics(
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

func getDictionaryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetDictionaryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetDictionaryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetDictionaryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetDictionaryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetDictionaryNamespaceMetricsAsync(
	request *GetDictionaryNamespaceMetricsRequest,
	callback chan<- GetDictionaryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/dictionary/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getDictionaryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetDictionaryNamespaceMetrics(
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

func describeExchangeRateModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExchangeRateModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeExchangeRateModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExchangeRateModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExchangeRateModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeExchangeRateModelMetricsAsync(
	request *DescribeExchangeRateModelMetricsRequest,
	callback chan<- DescribeExchangeRateModelMetricsAsyncResult,
) {
	path := "/metrics/exchange/namespace/{namespaceName}/rateModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeExchangeRateModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeExchangeRateModelMetrics(
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

func getExchangeRateModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetExchangeRateModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetExchangeRateModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExchangeRateModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExchangeRateModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetExchangeRateModelMetricsAsync(
	request *GetExchangeRateModelMetricsRequest,
	callback chan<- GetExchangeRateModelMetricsAsyncResult,
) {
	path := "/metrics/exchange/namespace/{namespaceName}/rateModel/{rateName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.RateName != nil && *request.RateName != "" {
		path = strings.ReplaceAll(path, "{rateName}", core.ToString(*request.RateName))
	} else {
		path = strings.ReplaceAll(path, "{rateName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getExchangeRateModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetExchangeRateModelMetrics(
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

func describeExchangeNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExchangeNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeExchangeNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExchangeNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExchangeNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeExchangeNamespaceMetricsAsync(
	request *DescribeExchangeNamespaceMetricsRequest,
	callback chan<- DescribeExchangeNamespaceMetricsAsyncResult,
) {
	path := "/metrics/exchange/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeExchangeNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeExchangeNamespaceMetrics(
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

func getExchangeNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetExchangeNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetExchangeNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExchangeNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExchangeNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetExchangeNamespaceMetricsAsync(
	request *GetExchangeNamespaceMetricsRequest,
	callback chan<- GetExchangeNamespaceMetricsAsyncResult,
) {
	path := "/metrics/exchange/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getExchangeNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetExchangeNamespaceMetrics(
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

func describeExperienceStatusMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExperienceStatusMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeExperienceStatusMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceStatusMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExperienceStatusMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeExperienceStatusMetricsAsync(
	request *DescribeExperienceStatusMetricsRequest,
	callback chan<- DescribeExperienceStatusMetricsAsyncResult,
) {
	path := "/metrics/experience/namespace/{namespaceName}/experienceModel/{experienceName}/status"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeExperienceStatusMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeExperienceStatusMetrics(
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

func describeExperienceExperienceModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExperienceExperienceModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExperienceExperienceModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeExperienceExperienceModelMetricsAsync(
	request *DescribeExperienceExperienceModelMetricsRequest,
	callback chan<- DescribeExperienceExperienceModelMetricsAsyncResult,
) {
	path := "/metrics/experience/namespace/{namespaceName}/experienceModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeExperienceExperienceModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeExperienceExperienceModelMetrics(
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

func getExperienceExperienceModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetExperienceExperienceModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetExperienceExperienceModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceExperienceModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExperienceExperienceModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetExperienceExperienceModelMetricsAsync(
	request *GetExperienceExperienceModelMetricsRequest,
	callback chan<- GetExperienceExperienceModelMetricsAsyncResult,
) {
	path := "/metrics/experience/namespace/{namespaceName}/experienceModel/{experienceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ExperienceName != nil && *request.ExperienceName != "" {
		path = strings.ReplaceAll(path, "{experienceName}", core.ToString(*request.ExperienceName))
	} else {
		path = strings.ReplaceAll(path, "{experienceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getExperienceExperienceModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetExperienceExperienceModelMetrics(
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

func describeExperienceNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeExperienceNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeExperienceNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeExperienceNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeExperienceNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeExperienceNamespaceMetricsAsync(
	request *DescribeExperienceNamespaceMetricsRequest,
	callback chan<- DescribeExperienceNamespaceMetricsAsyncResult,
) {
	path := "/metrics/experience/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeExperienceNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeExperienceNamespaceMetrics(
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

func getExperienceNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetExperienceNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetExperienceNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetExperienceNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetExperienceNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetExperienceNamespaceMetricsAsync(
	request *GetExperienceNamespaceMetricsRequest,
	callback chan<- GetExperienceNamespaceMetricsAsyncResult,
) {
	path := "/metrics/experience/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getExperienceNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetExperienceNamespaceMetrics(
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

func describeFormationFormMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormationFormMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeFormationFormMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationFormMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormationFormMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeFormationFormMetricsAsync(
	request *DescribeFormationFormMetricsRequest,
	callback chan<- DescribeFormationFormMetricsAsyncResult,
) {
	path := "/metrics/formation/namespace/{namespaceName}/mold/{moldModelName}/form"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MoldModelName != nil && *request.MoldModelName != "" {
		path = strings.ReplaceAll(path, "{moldModelName}", core.ToString(*request.MoldModelName))
	} else {
		path = strings.ReplaceAll(path, "{moldModelName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeFormationFormMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeFormationFormMetrics(
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

func describeFormationMoldMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormationMoldMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeFormationMoldMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationMoldMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormationMoldMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeFormationMoldMetricsAsync(
	request *DescribeFormationMoldMetricsRequest,
	callback chan<- DescribeFormationMoldMetricsAsyncResult,
) {
	path := "/metrics/formation/namespace/{namespaceName}/mold"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeFormationMoldMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeFormationMoldMetrics(
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

func describeFormationNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFormationNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeFormationNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFormationNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFormationNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeFormationNamespaceMetricsAsync(
	request *DescribeFormationNamespaceMetricsRequest,
	callback chan<- DescribeFormationNamespaceMetricsAsyncResult,
) {
	path := "/metrics/formation/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeFormationNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeFormationNamespaceMetrics(
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

func getFormationNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetFormationNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetFormationNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFormationNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFormationNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetFormationNamespaceMetricsAsync(
	request *GetFormationNamespaceMetricsRequest,
	callback chan<- GetFormationNamespaceMetricsAsyncResult,
) {
	path := "/metrics/formation/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getFormationNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetFormationNamespaceMetrics(
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

func describeFriendNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeFriendNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeFriendNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeFriendNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeFriendNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeFriendNamespaceMetricsAsync(
	request *DescribeFriendNamespaceMetricsRequest,
	callback chan<- DescribeFriendNamespaceMetricsAsyncResult,
) {
	path := "/metrics/friend/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeFriendNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeFriendNamespaceMetrics(
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

func getFriendNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetFriendNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetFriendNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetFriendNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetFriendNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetFriendNamespaceMetricsAsync(
	request *GetFriendNamespaceMetricsRequest,
	callback chan<- GetFriendNamespaceMetricsAsyncResult,
) {
	path := "/metrics/friend/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getFriendNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetFriendNamespaceMetrics(
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

func describeInboxNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInboxNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeInboxNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInboxNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInboxNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeInboxNamespaceMetricsAsync(
	request *DescribeInboxNamespaceMetricsRequest,
	callback chan<- DescribeInboxNamespaceMetricsAsyncResult,
) {
	path := "/metrics/inbox/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeInboxNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeInboxNamespaceMetrics(
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

func getInboxNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetInboxNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetInboxNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInboxNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInboxNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetInboxNamespaceMetricsAsync(
	request *GetInboxNamespaceMetricsRequest,
	callback chan<- GetInboxNamespaceMetricsAsyncResult,
) {
	path := "/metrics/inbox/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getInboxNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetInboxNamespaceMetrics(
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

func describeInventoryItemSetMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryItemSetMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeInventoryItemSetMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryItemSetMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoryItemSetMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeInventoryItemSetMetricsAsync(
	request *DescribeInventoryItemSetMetricsRequest,
	callback chan<- DescribeInventoryItemSetMetricsAsyncResult,
) {
	path := "/metrics/inventory/namespace/{namespaceName}/inventory/{inventoryName}/itemSet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.InventoryName != nil && *request.InventoryName != "" {
		path = strings.ReplaceAll(path, "{inventoryName}", core.ToString(*request.InventoryName))
	} else {
		path = strings.ReplaceAll(path, "{inventoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeInventoryItemSetMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeInventoryItemSetMetrics(
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

func describeInventoryInventoryMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryInventoryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeInventoryInventoryMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryInventoryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoryInventoryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeInventoryInventoryMetricsAsync(
	request *DescribeInventoryInventoryMetricsRequest,
	callback chan<- DescribeInventoryInventoryMetricsAsyncResult,
) {
	path := "/metrics/inventory/namespace/{namespaceName}/inventory"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeInventoryInventoryMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeInventoryInventoryMetrics(
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

func describeInventoryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeInventoryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeInventoryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeInventoryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeInventoryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeInventoryNamespaceMetricsAsync(
	request *DescribeInventoryNamespaceMetricsRequest,
	callback chan<- DescribeInventoryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/inventory/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeInventoryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeInventoryNamespaceMetrics(
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

func getInventoryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetInventoryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetInventoryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetInventoryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetInventoryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetInventoryNamespaceMetricsAsync(
	request *GetInventoryNamespaceMetricsRequest,
	callback chan<- GetInventoryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/inventory/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getInventoryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetInventoryNamespaceMetrics(
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

func describeKeyNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeKeyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeKeyNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeKeyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeKeyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeKeyNamespaceMetricsAsync(
	request *DescribeKeyNamespaceMetricsRequest,
	callback chan<- DescribeKeyNamespaceMetricsAsyncResult,
) {
	path := "/metrics/key/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeKeyNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeKeyNamespaceMetrics(
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

func getKeyNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetKeyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetKeyNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetKeyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetKeyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetKeyNamespaceMetricsAsync(
	request *GetKeyNamespaceMetricsRequest,
	callback chan<- GetKeyNamespaceMetricsAsyncResult,
) {
	path := "/metrics/key/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getKeyNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetKeyNamespaceMetrics(
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

func describeLimitCounterMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLimitCounterMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLimitCounterMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitCounterMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLimitCounterMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeLimitCounterMetricsAsync(
	request *DescribeLimitCounterMetricsRequest,
	callback chan<- DescribeLimitCounterMetricsAsyncResult,
) {
	path := "/metrics/limit/namespace/{namespaceName}/limitModel/{limitName}/counter"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLimitCounterMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeLimitCounterMetrics(
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

func describeLimitLimitModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLimitLimitModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLimitLimitModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitLimitModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLimitLimitModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeLimitLimitModelMetricsAsync(
	request *DescribeLimitLimitModelMetricsRequest,
	callback chan<- DescribeLimitLimitModelMetricsAsyncResult,
) {
	path := "/metrics/limit/namespace/{namespaceName}/limitModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLimitLimitModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeLimitLimitModelMetrics(
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

func getLimitLimitModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetLimitLimitModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetLimitLimitModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitLimitModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLimitLimitModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetLimitLimitModelMetricsAsync(
	request *GetLimitLimitModelMetricsRequest,
	callback chan<- GetLimitLimitModelMetricsAsyncResult,
) {
	path := "/metrics/limit/namespace/{namespaceName}/limitModel/{limitName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LimitName != nil && *request.LimitName != "" {
		path = strings.ReplaceAll(path, "{limitName}", core.ToString(*request.LimitName))
	} else {
		path = strings.ReplaceAll(path, "{limitName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLimitLimitModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetLimitLimitModelMetrics(
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

func describeLimitNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLimitNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLimitNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLimitNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLimitNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeLimitNamespaceMetricsAsync(
	request *DescribeLimitNamespaceMetricsRequest,
	callback chan<- DescribeLimitNamespaceMetricsAsyncResult,
) {
	path := "/metrics/limit/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLimitNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeLimitNamespaceMetrics(
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

func getLimitNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetLimitNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetLimitNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLimitNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLimitNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetLimitNamespaceMetricsAsync(
	request *GetLimitNamespaceMetricsRequest,
	callback chan<- GetLimitNamespaceMetricsAsyncResult,
) {
	path := "/metrics/limit/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLimitNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetLimitNamespaceMetrics(
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

func describeLotteryLotteryMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryLotteryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLotteryLotteryMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryLotteryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLotteryLotteryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeLotteryLotteryMetricsAsync(
	request *DescribeLotteryLotteryMetricsRequest,
	callback chan<- DescribeLotteryLotteryMetricsAsyncResult,
) {
	path := "/metrics/lottery/namespace/{namespaceName}/lottery"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLotteryLotteryMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeLotteryLotteryMetrics(
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

func getLotteryLotteryMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryLotteryMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetLotteryLotteryMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryLotteryMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLotteryLotteryMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetLotteryLotteryMetricsAsync(
	request *GetLotteryLotteryMetricsRequest,
	callback chan<- GetLotteryLotteryMetricsAsyncResult,
) {
	path := "/metrics/lottery/namespace/{namespaceName}/lotteryModel/{lotteryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.LotteryName != nil && *request.LotteryName != "" {
		path = strings.ReplaceAll(path, "{lotteryName}", core.ToString(*request.LotteryName))
	} else {
		path = strings.ReplaceAll(path, "{lotteryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLotteryLotteryMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetLotteryLotteryMetrics(
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

func describeLotteryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeLotteryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeLotteryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeLotteryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeLotteryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeLotteryNamespaceMetricsAsync(
	request *DescribeLotteryNamespaceMetricsRequest,
	callback chan<- DescribeLotteryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/lottery/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeLotteryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeLotteryNamespaceMetrics(
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

func getLotteryNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetLotteryNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetLotteryNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetLotteryNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetLotteryNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetLotteryNamespaceMetricsAsync(
	request *GetLotteryNamespaceMetricsRequest,
	callback chan<- GetLotteryNamespaceMetricsAsyncResult,
) {
	path := "/metrics/lottery/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getLotteryNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetLotteryNamespaceMetrics(
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

func describeMatchmakingNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMatchmakingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMatchmakingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMatchmakingNamespaceMetricsAsync(
	request *DescribeMatchmakingNamespaceMetricsRequest,
	callback chan<- DescribeMatchmakingNamespaceMetricsAsyncResult,
) {
	path := "/metrics/matchmaking/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMatchmakingNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMatchmakingNamespaceMetrics(
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

func getMatchmakingNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetMatchmakingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMatchmakingNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMatchmakingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMatchmakingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetMatchmakingNamespaceMetricsAsync(
	request *GetMatchmakingNamespaceMetricsRequest,
	callback chan<- GetMatchmakingNamespaceMetricsAsyncResult,
) {
	path := "/metrics/matchmaking/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getMatchmakingNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetMatchmakingNamespaceMetrics(
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

func describeMissionCounterMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionCounterMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionCounterMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionCounterMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionCounterMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMissionCounterMetricsAsync(
	request *DescribeMissionCounterMetricsRequest,
	callback chan<- DescribeMissionCounterMetricsAsyncResult,
) {
	path := "/metrics/mission/namespace/{namespaceName}/counter"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMissionCounterMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMissionCounterMetrics(
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

func describeMissionMissionGroupModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionMissionGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionMissionGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMissionMissionGroupModelMetricsAsync(
	request *DescribeMissionMissionGroupModelMetricsRequest,
	callback chan<- DescribeMissionMissionGroupModelMetricsAsyncResult,
) {
	path := "/metrics/mission/namespace/{namespaceName}/missionGroupModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMissionMissionGroupModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMissionMissionGroupModelMetrics(
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

func getMissionMissionGroupModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionMissionGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionMissionGroupModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionMissionGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionMissionGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetMissionMissionGroupModelMetricsAsync(
	request *GetMissionMissionGroupModelMetricsRequest,
	callback chan<- GetMissionMissionGroupModelMetricsAsyncResult,
) {
	path := "/metrics/mission/namespace/{namespaceName}/missionGroupModel/{missionGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.MissionGroupName != nil && *request.MissionGroupName != "" {
		path = strings.ReplaceAll(path, "{missionGroupName}", core.ToString(*request.MissionGroupName))
	} else {
		path = strings.ReplaceAll(path, "{missionGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getMissionMissionGroupModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetMissionMissionGroupModelMetrics(
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

func describeMissionNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMissionNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMissionNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMissionNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMissionNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMissionNamespaceMetricsAsync(
	request *DescribeMissionNamespaceMetricsRequest,
	callback chan<- DescribeMissionNamespaceMetricsAsyncResult,
) {
	path := "/metrics/mission/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMissionNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMissionNamespaceMetrics(
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

func getMissionNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetMissionNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMissionNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMissionNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMissionNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetMissionNamespaceMetricsAsync(
	request *GetMissionNamespaceMetricsRequest,
	callback chan<- GetMissionNamespaceMetricsAsyncResult,
) {
	path := "/metrics/mission/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getMissionNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetMissionNamespaceMetrics(
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

func describeMoneyWalletMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoneyWalletMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMoneyWalletMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyWalletMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoneyWalletMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMoneyWalletMetricsAsync(
	request *DescribeMoneyWalletMetricsRequest,
	callback chan<- DescribeMoneyWalletMetricsAsyncResult,
) {
	path := "/metrics/money/namespace/{namespaceName}/wallet"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMoneyWalletMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMoneyWalletMetrics(
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

func describeMoneyReceiptMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoneyReceiptMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMoneyReceiptMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyReceiptMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoneyReceiptMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMoneyReceiptMetricsAsync(
	request *DescribeMoneyReceiptMetricsRequest,
	callback chan<- DescribeMoneyReceiptMetricsAsyncResult,
) {
	path := "/metrics/money/namespace/{namespaceName}/receipt"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMoneyReceiptMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMoneyReceiptMetrics(
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

func describeMoneyNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeMoneyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeMoneyNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeMoneyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeMoneyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeMoneyNamespaceMetricsAsync(
	request *DescribeMoneyNamespaceMetricsRequest,
	callback chan<- DescribeMoneyNamespaceMetricsAsyncResult,
) {
	path := "/metrics/money/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeMoneyNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeMoneyNamespaceMetrics(
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

func getMoneyNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetMoneyNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetMoneyNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetMoneyNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetMoneyNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetMoneyNamespaceMetricsAsync(
	request *GetMoneyNamespaceMetricsRequest,
	callback chan<- GetMoneyNamespaceMetricsAsyncResult,
) {
	path := "/metrics/money/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getMoneyNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetMoneyNamespaceMetrics(
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

func describeQuestQuestModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestQuestModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeQuestQuestModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestQuestModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestQuestModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeQuestQuestModelMetricsAsync(
	request *DescribeQuestQuestModelMetricsRequest,
	callback chan<- DescribeQuestQuestModelMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace/{namespaceName}/questGroupModel/{questGroupName}/questModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeQuestQuestModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeQuestQuestModelMetrics(
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

func getQuestQuestModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestQuestModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetQuestQuestModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestQuestModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestQuestModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetQuestQuestModelMetricsAsync(
	request *GetQuestQuestModelMetricsRequest,
	callback chan<- GetQuestQuestModelMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace/{namespaceName}/questGroupModel/{questGroupName}/questModel/{questName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}
	if request.QuestName != nil && *request.QuestName != "" {
		path = strings.ReplaceAll(path, "{questName}", core.ToString(*request.QuestName))
	} else {
		path = strings.ReplaceAll(path, "{questName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getQuestQuestModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetQuestQuestModelMetrics(
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

func describeQuestQuestGroupModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestQuestGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestQuestGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeQuestQuestGroupModelMetricsAsync(
	request *DescribeQuestQuestGroupModelMetricsRequest,
	callback chan<- DescribeQuestQuestGroupModelMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace/{namespaceName}/questGroupModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeQuestQuestGroupModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeQuestQuestGroupModelMetrics(
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

func getQuestQuestGroupModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestQuestGroupModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetQuestQuestGroupModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestQuestGroupModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestQuestGroupModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetQuestQuestGroupModelMetricsAsync(
	request *GetQuestQuestGroupModelMetricsRequest,
	callback chan<- GetQuestQuestGroupModelMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace/{namespaceName}/questGroupModel/{questGroupName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.QuestGroupName != nil && *request.QuestGroupName != "" {
		path = strings.ReplaceAll(path, "{questGroupName}", core.ToString(*request.QuestGroupName))
	} else {
		path = strings.ReplaceAll(path, "{questGroupName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getQuestQuestGroupModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetQuestQuestGroupModelMetrics(
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

func describeQuestNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeQuestNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeQuestNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeQuestNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeQuestNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeQuestNamespaceMetricsAsync(
	request *DescribeQuestNamespaceMetricsRequest,
	callback chan<- DescribeQuestNamespaceMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeQuestNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeQuestNamespaceMetrics(
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

func getQuestNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetQuestNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetQuestNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetQuestNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetQuestNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetQuestNamespaceMetricsAsync(
	request *GetQuestNamespaceMetricsRequest,
	callback chan<- GetQuestNamespaceMetricsAsyncResult,
) {
	path := "/metrics/quest/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getQuestNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetQuestNamespaceMetrics(
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

func describeRankingCategoryModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRankingCategoryModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeRankingCategoryModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingCategoryModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRankingCategoryModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeRankingCategoryModelMetricsAsync(
	request *DescribeRankingCategoryModelMetricsRequest,
	callback chan<- DescribeRankingCategoryModelMetricsAsyncResult,
) {
	path := "/metrics/ranking/namespace/{namespaceName}/categoryModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeRankingCategoryModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeRankingCategoryModelMetrics(
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

func getRankingCategoryModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetRankingCategoryModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetRankingCategoryModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingCategoryModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRankingCategoryModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetRankingCategoryModelMetricsAsync(
	request *GetRankingCategoryModelMetricsRequest,
	callback chan<- GetRankingCategoryModelMetricsAsyncResult,
) {
	path := "/metrics/ranking/namespace/{namespaceName}/categoryModel/{categoryName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.CategoryName != nil && *request.CategoryName != "" {
		path = strings.ReplaceAll(path, "{categoryName}", core.ToString(*request.CategoryName))
	} else {
		path = strings.ReplaceAll(path, "{categoryName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getRankingCategoryModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetRankingCategoryModelMetrics(
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

func describeRankingNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeRankingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeRankingNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeRankingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeRankingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeRankingNamespaceMetricsAsync(
	request *DescribeRankingNamespaceMetricsRequest,
	callback chan<- DescribeRankingNamespaceMetricsAsyncResult,
) {
	path := "/metrics/ranking/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeRankingNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeRankingNamespaceMetrics(
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

func getRankingNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetRankingNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetRankingNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetRankingNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetRankingNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetRankingNamespaceMetricsAsync(
	request *GetRankingNamespaceMetricsRequest,
	callback chan<- GetRankingNamespaceMetricsAsyncResult,
) {
	path := "/metrics/ranking/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getRankingNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetRankingNamespaceMetrics(
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

func describeShowcaseDisplayItemMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcaseDisplayItemMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeShowcaseDisplayItemMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeShowcaseDisplayItemMetricsAsync(
	request *DescribeShowcaseDisplayItemMetricsRequest,
	callback chan<- DescribeShowcaseDisplayItemMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace/{namespaceName}/showcase/{showcaseName}/displayItem"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
	} else {
		path = strings.ReplaceAll(path, "{showcaseName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeShowcaseDisplayItemMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeShowcaseDisplayItemMetrics(
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

func getShowcaseDisplayItemMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseDisplayItemMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetShowcaseDisplayItemMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseDisplayItemMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetShowcaseDisplayItemMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetShowcaseDisplayItemMetricsAsync(
	request *GetShowcaseDisplayItemMetricsRequest,
	callback chan<- GetShowcaseDisplayItemMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace/{namespaceName}/showcase/{showcaseName}/displayItem/{displayItemId}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
	} else {
		path = strings.ReplaceAll(path, "{showcaseName}", "null")
	}
	if request.DisplayItemId != nil && *request.DisplayItemId != "" {
		path = strings.ReplaceAll(path, "{displayItemId}", core.ToString(*request.DisplayItemId))
	} else {
		path = strings.ReplaceAll(path, "{displayItemId}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getShowcaseDisplayItemMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetShowcaseDisplayItemMetrics(
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

func describeShowcaseShowcaseMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcaseShowcaseMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeShowcaseShowcaseMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeShowcaseShowcaseMetricsAsync(
	request *DescribeShowcaseShowcaseMetricsRequest,
	callback chan<- DescribeShowcaseShowcaseMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace/{namespaceName}/showcase"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeShowcaseShowcaseMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeShowcaseShowcaseMetrics(
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

func getShowcaseShowcaseMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseShowcaseMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetShowcaseShowcaseMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseShowcaseMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetShowcaseShowcaseMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetShowcaseShowcaseMetricsAsync(
	request *GetShowcaseShowcaseMetricsRequest,
	callback chan<- GetShowcaseShowcaseMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace/{namespaceName}/showcase/{showcaseName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.ShowcaseName != nil && *request.ShowcaseName != "" {
		path = strings.ReplaceAll(path, "{showcaseName}", core.ToString(*request.ShowcaseName))
	} else {
		path = strings.ReplaceAll(path, "{showcaseName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getShowcaseShowcaseMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetShowcaseShowcaseMetrics(
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

func describeShowcaseNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeShowcaseNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeShowcaseNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeShowcaseNamespaceMetricsAsync(
	request *DescribeShowcaseNamespaceMetricsRequest,
	callback chan<- DescribeShowcaseNamespaceMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeShowcaseNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeShowcaseNamespaceMetrics(
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

func getShowcaseNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetShowcaseNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetShowcaseNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetShowcaseNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetShowcaseNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetShowcaseNamespaceMetricsAsync(
	request *GetShowcaseNamespaceMetricsRequest,
	callback chan<- GetShowcaseNamespaceMetricsAsyncResult,
) {
	path := "/metrics/showcase/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getShowcaseNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetShowcaseNamespaceMetrics(
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

func describeStaminaStaminaModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminaStaminaModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminaStaminaModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeStaminaStaminaModelMetricsAsync(
	request *DescribeStaminaStaminaModelMetricsRequest,
	callback chan<- DescribeStaminaStaminaModelMetricsAsyncResult,
) {
	path := "/metrics/stamina/namespace/{namespaceName}/staminaModel"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeStaminaStaminaModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeStaminaStaminaModelMetrics(
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

func getStaminaStaminaModelMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaStaminaModelMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetStaminaStaminaModelMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaStaminaModelMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaStaminaModelMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetStaminaStaminaModelMetricsAsync(
	request *GetStaminaStaminaModelMetricsRequest,
	callback chan<- GetStaminaStaminaModelMetricsAsyncResult,
) {
	path := "/metrics/stamina/namespace/{namespaceName}/staminaModel/{staminaName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}
	if request.StaminaName != nil && *request.StaminaName != "" {
		path = strings.ReplaceAll(path, "{staminaName}", core.ToString(*request.StaminaName))
	} else {
		path = strings.ReplaceAll(path, "{staminaName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getStaminaStaminaModelMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetStaminaStaminaModelMetrics(
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

func describeStaminaNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- DescribeStaminaNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- DescribeStaminaNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- DescribeStaminaNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- DescribeStaminaNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) DescribeStaminaNamespaceMetricsAsync(
	request *DescribeStaminaNamespaceMetricsRequest,
	callback chan<- DescribeStaminaNamespaceMetricsAsyncResult,
) {
	path := "/metrics/stamina/namespace"

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go describeStaminaNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) DescribeStaminaNamespaceMetrics(
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

func getStaminaNamespaceMetricsAsyncHandler(
	client Gs2WatchRestClient,
	job *core.NetworkJob,
	callback chan<- GetStaminaNamespaceMetricsAsyncResult,
) {
	internalCallback := make(chan core.AsyncResult, 1)
	job.Callback = internalCallback
	err := client.Session.Send(
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
	if asyncResult.Err != nil {
		callback <- GetStaminaNamespaceMetricsAsyncResult{
			err: asyncResult.Err,
		}
		return
	}
	if asyncResult.Payload != "" {
		err = json.Unmarshal([]byte(asyncResult.Payload), &result)
		if err != nil {
			callback <- GetStaminaNamespaceMetricsAsyncResult{
				err: err,
			}
			return
		}
	}
	callback <- GetStaminaNamespaceMetricsAsyncResult{
		result: &result,
		err:    asyncResult.Err,
	}

}

func (p Gs2WatchRestClient) GetStaminaNamespaceMetricsAsync(
	request *GetStaminaNamespaceMetricsRequest,
	callback chan<- GetStaminaNamespaceMetricsAsyncResult,
) {
	path := "/metrics/stamina/namespace/{namespaceName}"
	if request.NamespaceName != nil && *request.NamespaceName != "" {
		path = strings.ReplaceAll(path, "{namespaceName}", core.ToString(*request.NamespaceName))
	} else {
		path = strings.ReplaceAll(path, "{namespaceName}", "null")
	}

	replacer := strings.NewReplacer()
	queryStrings := core.QueryStrings{}
	if request.ContextStack != nil {
		queryStrings["contextStack"] = *request.ContextStack
	}

	headers := p.CreateAuthorizedHeaders()
	if request.DryRun != nil {
		if *request.DryRun {
			headers["X-GS2-DRY-RUN"] = "true"
		} else {
			headers["X-GS2-DRY-RUN"] = "false"
		}
	}

	go getStaminaNamespaceMetricsAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Get,
			Headers:      headers,
			QueryStrings: queryStrings,
		},
		callback,
	)
}

func (p Gs2WatchRestClient) GetStaminaNamespaceMetrics(
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
