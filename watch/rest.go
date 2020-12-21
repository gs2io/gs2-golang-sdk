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
	"github.com/gs2io/gs2-golang-sdk/core"
	"strings"
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
	path := "/chart/{metrics}"
    if request.Metrics != nil {
        path = strings.ReplaceAll(path, "{metrics}", core.ToString(*request.Metrics))
    } else {
        path = strings.ReplaceAll(path, "{metrics}", "null")
    }

	replacer := strings.NewReplacer()
    var bodies = core.Bodies{}
    if request.Grn != nil && *request.Grn != "" {
        bodies["grn"] = *request.Grn
    }
    if request.Queries != nil {
        var _queries []core.String
        for _, item := range *request.Queries {
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getChartAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
    if request.Name != nil {
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
    	bodies["contextStack"] = *request.ContextStack;
	}

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
    }

	go getCumulativeAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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

    headers := p.CreateAuthorizedHeaders()
    if request.RequestId != nil {
        headers["X-GS2-REQUEST-ID"] = string(*request.RequestId)
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
    if request.Service != nil {
        path = strings.ReplaceAll(path, "{service}", core.ToString(*request.Service))
    } else {
        path = strings.ReplaceAll(path, "{service}", "null")
    }
    if request.ActivityType != nil {
        path = strings.ReplaceAll(path, "{activityType}", core.ToString(*request.ActivityType))
    } else {
        path = strings.ReplaceAll(path, "{activityType}", "null")
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

	go getBillingActivityAsyncHandler(
		p,
		&core.NetworkJob{
			Url:          p.Session.EndpointHost("watch").AppendPath(path, replacer),
			Method:       core.Post,
			Headers:      headers,
			Bodies: bodies,
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
