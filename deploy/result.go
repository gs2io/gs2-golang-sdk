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

package deploy

type DescribeStacksResult struct {
    /** スタックのリスト */
	Items         []Stack	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeStacksResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Stack, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeStacksAsyncResult struct {
	result *DescribeStacksResult
	err    error
}

type CreateStackResult struct {
    /** 作成したスタック */
	Item         *Stack	`json:"item"`
}

func (p *CreateStackResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateStackAsyncResult struct {
	result *CreateStackResult
	err    error
}

type CreateStackFromGitHubResult struct {
    /** 作成したスタック */
	Item         *Stack	`json:"item"`
}

func (p *CreateStackFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateStackFromGitHubAsyncResult struct {
	result *CreateStackFromGitHubResult
	err    error
}

type ValidateResult struct {
}

func (p *ValidateResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type ValidateAsyncResult struct {
	result *ValidateResult
	err    error
}

type GetStackStatusResult struct {
    /** None */
	Status         *string	`json:"status"`
}

func (p *GetStackStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
        data["status"] = p.Status
    }
    return &data
}

type GetStackStatusAsyncResult struct {
	result *GetStackStatusResult
	err    error
}

type GetStackResult struct {
    /** スタック */
	Item         *Stack	`json:"item"`
}

func (p *GetStackResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetStackAsyncResult struct {
	result *GetStackResult
	err    error
}

type UpdateStackResult struct {
    /** 更新したスタック */
	Item         *Stack	`json:"item"`
}

func (p *UpdateStackResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateStackAsyncResult struct {
	result *UpdateStackResult
	err    error
}

type UpdateStackFromGitHubResult struct {
    /** 更新したスタック */
	Item         *Stack	`json:"item"`
}

func (p *UpdateStackFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateStackFromGitHubAsyncResult struct {
	result *UpdateStackFromGitHubResult
	err    error
}

type DeleteStackResult struct {
    /** 削除したスタック */
	Item         *Stack	`json:"item"`
}

func (p *DeleteStackResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteStackAsyncResult struct {
	result *DeleteStackResult
	err    error
}

type ForceDeleteStackResult struct {
    /** 削除したスタック */
	Item         *Stack	`json:"item"`
}

func (p *ForceDeleteStackResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type ForceDeleteStackAsyncResult struct {
	result *ForceDeleteStackResult
	err    error
}

type DeleteStackResourcesResult struct {
    /** リソースを削除したスタック */
	Item         *Stack	`json:"item"`
}

func (p *DeleteStackResourcesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteStackResourcesAsyncResult struct {
	result *DeleteStackResourcesResult
	err    error
}

type DeleteStackEntityResult struct {
    /** 削除したスタック */
	Item         *Stack	`json:"item"`
}

func (p *DeleteStackEntityResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteStackEntityAsyncResult struct {
	result *DeleteStackEntityResult
	err    error
}

type DescribeResourcesResult struct {
    /** 作成されたのリソースのリスト */
	Items         []Resource	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeResourcesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Resource, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeResourcesAsyncResult struct {
	result *DescribeResourcesResult
	err    error
}

type GetResourceResult struct {
    /** 作成されたのリソース */
	Item         *Resource	`json:"item"`
}

func (p *GetResourceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetResourceAsyncResult struct {
	result *GetResourceResult
	err    error
}

type DescribeEventsResult struct {
    /** 発生したイベントのリスト */
	Items         []Event	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeEventsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Event, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

type GetEventResult struct {
    /** 発生したイベント */
	Item         *Event	`json:"item"`
}

func (p *GetEventResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

type DescribeOutputsResult struct {
    /** アウトプットのリスト */
	Items         []Output	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeOutputsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Output, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeOutputsAsyncResult struct {
	result *DescribeOutputsResult
	err    error
}

type GetOutputResult struct {
    /** アウトプット */
	Item         *Output	`json:"item"`
}

func (p *GetOutputResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetOutputAsyncResult struct {
	result *GetOutputResult
	err    error
}
