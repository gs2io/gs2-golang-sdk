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
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         *[]*Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

type CreateNamespaceResult struct {
    /** 作成したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *CreateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

type GetNamespaceStatusResult struct {
    /** None */
	Status         *core.String	`json:"status"`
}

func (p *GetNamespaceStatusResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Status != nil {
     data["Status"] = p.Status
    }
    return &data
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

type GetNamespaceResult struct {
    /** ネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *GetNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

type UpdateNamespaceResult struct {
    /** 更新したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *UpdateNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeEventMastersResult struct {
    /** イベントマスターのリスト */
	Items         *[]*EventMaster	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeEventMastersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeEventMastersAsyncResult struct {
	result *DescribeEventMastersResult
	err    error
}

type CreateEventMasterResult struct {
    /** 作成したイベントマスター */
	Item         *EventMaster	`json:"item"`
}

func (p *CreateEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type CreateEventMasterAsyncResult struct {
	result *CreateEventMasterResult
	err    error
}

type GetEventMasterResult struct {
    /** イベントマスター */
	Item         *EventMaster	`json:"item"`
}

func (p *GetEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetEventMasterAsyncResult struct {
	result *GetEventMasterResult
	err    error
}

type UpdateEventMasterResult struct {
    /** 更新したイベントマスター */
	Item         *EventMaster	`json:"item"`
}

func (p *UpdateEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateEventMasterAsyncResult struct {
	result *UpdateEventMasterResult
	err    error
}

type DeleteEventMasterResult struct {
    /** 削除したイベントマスター */
	Item         *EventMaster	`json:"item"`
}

func (p *DeleteEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteEventMasterAsyncResult struct {
	result *DeleteEventMasterResult
	err    error
}

type DescribeTriggersResult struct {
    /** トリガーのリスト */
	Items         *[]*Trigger	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeTriggersResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeTriggersAsyncResult struct {
	result *DescribeTriggersResult
	err    error
}

type DescribeTriggersByUserIdResult struct {
    /** トリガーのリスト */
	Items         *[]*Trigger	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeTriggersByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    if p.NextPageToken != nil {
     data["NextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeTriggersByUserIdAsyncResult struct {
	result *DescribeTriggersByUserIdResult
	err    error
}

type GetTriggerResult struct {
    /** トリガー */
	Item         *Trigger	`json:"item"`
}

func (p *GetTriggerResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetTriggerAsyncResult struct {
	result *GetTriggerResult
	err    error
}

type GetTriggerByUserIdResult struct {
    /** トリガー */
	Item         *Trigger	`json:"item"`
}

func (p *GetTriggerByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetTriggerByUserIdAsyncResult struct {
	result *GetTriggerByUserIdResult
	err    error
}

type TriggerByUserIdResult struct {
    /** 引いたトリガー */
	Item         *Trigger	`json:"item"`
}

func (p *TriggerByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type TriggerByUserIdAsyncResult struct {
	result *TriggerByUserIdResult
	err    error
}

type DeleteTriggerResult struct {
    /** トリガー */
	Item         *Trigger	`json:"item"`
}

func (p *DeleteTriggerResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteTriggerAsyncResult struct {
	result *DeleteTriggerResult
	err    error
}

type DeleteTriggerByUserIdResult struct {
    /** トリガー */
	Item         *Trigger	`json:"item"`
}

func (p *DeleteTriggerByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteTriggerByUserIdAsyncResult struct {
	result *DeleteTriggerByUserIdResult
	err    error
}

type DescribeEventsResult struct {
    /** イベントのリスト */
	Items         *[]*Event	`json:"items"`
}

func (p *DescribeEventsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

type DescribeEventsByUserIdResult struct {
    /** イベントのリスト */
	Items         *[]*Event	`json:"items"`
}

func (p *DescribeEventsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeEventsByUserIdAsyncResult struct {
	result *DescribeEventsByUserIdResult
	err    error
}

type DescribeRawEventsResult struct {
    /** イベントのリスト */
	Items         *[]*Event	`json:"items"`
}

func (p *DescribeRawEventsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    }
    return &data
}

type DescribeRawEventsAsyncResult struct {
	result *DescribeRawEventsResult
	err    error
}

type GetEventResult struct {
    /** イベント */
	Item         *Event	`json:"item"`
}

func (p *GetEventResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

type GetEventByUserIdResult struct {
    /** イベント */
	Item         *Event	`json:"item"`
}

func (p *GetEventByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetEventByUserIdAsyncResult struct {
	result *GetEventByUserIdResult
	err    error
}

type GetRawEventResult struct {
    /** イベント */
	Item         *Event	`json:"item"`
}

func (p *GetRawEventResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetRawEventAsyncResult struct {
	result *GetRawEventResult
	err    error
}

type ExportMasterResult struct {
    /** 現在有効なイベントスケジュールマスター */
	Item         *CurrentEventMaster	`json:"item"`
}

func (p *ExportMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

type GetCurrentEventMasterResult struct {
    /** 現在有効なイベントスケジュールマスター */
	Item         *CurrentEventMaster	`json:"item"`
}

func (p *GetCurrentEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type GetCurrentEventMasterAsyncResult struct {
	result *GetCurrentEventMasterResult
	err    error
}

type UpdateCurrentEventMasterResult struct {
    /** 更新した現在有効なイベントスケジュールマスター */
	Item         *CurrentEventMaster	`json:"item"`
}

func (p *UpdateCurrentEventMasterResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentEventMasterAsyncResult struct {
	result *UpdateCurrentEventMasterResult
	err    error
}

type UpdateCurrentEventMasterFromGitHubResult struct {
    /** 更新した現在有効なイベントスケジュールマスター */
	Item         *CurrentEventMaster	`json:"item"`
}

func (p *UpdateCurrentEventMasterFromGitHubResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
     data["Item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateCurrentEventMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEventMasterFromGitHubResult
	err    error
}
