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

import "core"

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromDict(data map[string]interface{}) DescribeNamespacesResult {
    return DescribeNamespacesResult {
        Items: CastNamespaces(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeNamespacesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastNamespacesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeNamespacesResult) Pointer() *DescribeNamespacesResult {
    return &p
}

type CreateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type CreateNamespaceAsyncResult struct {
	result *CreateNamespaceResult
	err    error
}

func NewCreateNamespaceResultFromDict(data map[string]interface{}) CreateNamespaceResult {
    return CreateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateNamespaceResult) Pointer() *CreateNamespaceResult {
    return &p
}

type GetNamespaceStatusResult struct {
    Status *string `json:"status"`
}

type GetNamespaceStatusAsyncResult struct {
	result *GetNamespaceStatusResult
	err    error
}

func NewGetNamespaceStatusResultFromDict(data map[string]interface{}) GetNamespaceStatusResult {
    return GetNamespaceStatusResult {
        Status: core.CastString(data["status"]),
    }
}

func (p GetNamespaceStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "status": p.Status,
    }
}

func (p GetNamespaceStatusResult) Pointer() *GetNamespaceStatusResult {
    return &p
}

type GetNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type GetNamespaceAsyncResult struct {
	result *GetNamespaceResult
	err    error
}

func NewGetNamespaceResultFromDict(data map[string]interface{}) GetNamespaceResult {
    return GetNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetNamespaceResult) Pointer() *GetNamespaceResult {
    return &p
}

type UpdateNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

func NewUpdateNamespaceResultFromDict(data map[string]interface{}) UpdateNamespaceResult {
    return UpdateNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateNamespaceResult) Pointer() *UpdateNamespaceResult {
    return &p
}

type DeleteNamespaceResult struct {
    Item *Namespace `json:"item"`
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

func NewDeleteNamespaceResultFromDict(data map[string]interface{}) DeleteNamespaceResult {
    return DeleteNamespaceResult {
        Item: NewNamespaceFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteNamespaceResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteNamespaceResult) Pointer() *DeleteNamespaceResult {
    return &p
}

type DescribeEventMastersResult struct {
    Items []EventMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeEventMastersAsyncResult struct {
	result *DescribeEventMastersResult
	err    error
}

func NewDescribeEventMastersResultFromDict(data map[string]interface{}) DescribeEventMastersResult {
    return DescribeEventMastersResult {
        Items: CastEventMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeEventMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastEventMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeEventMastersResult) Pointer() *DescribeEventMastersResult {
    return &p
}

type CreateEventMasterResult struct {
    Item *EventMaster `json:"item"`
}

type CreateEventMasterAsyncResult struct {
	result *CreateEventMasterResult
	err    error
}

func NewCreateEventMasterResultFromDict(data map[string]interface{}) CreateEventMasterResult {
    return CreateEventMasterResult {
        Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateEventMasterResult) Pointer() *CreateEventMasterResult {
    return &p
}

type GetEventMasterResult struct {
    Item *EventMaster `json:"item"`
}

type GetEventMasterAsyncResult struct {
	result *GetEventMasterResult
	err    error
}

func NewGetEventMasterResultFromDict(data map[string]interface{}) GetEventMasterResult {
    return GetEventMasterResult {
        Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetEventMasterResult) Pointer() *GetEventMasterResult {
    return &p
}

type UpdateEventMasterResult struct {
    Item *EventMaster `json:"item"`
}

type UpdateEventMasterAsyncResult struct {
	result *UpdateEventMasterResult
	err    error
}

func NewUpdateEventMasterResultFromDict(data map[string]interface{}) UpdateEventMasterResult {
    return UpdateEventMasterResult {
        Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateEventMasterResult) Pointer() *UpdateEventMasterResult {
    return &p
}

type DeleteEventMasterResult struct {
    Item *EventMaster `json:"item"`
}

type DeleteEventMasterAsyncResult struct {
	result *DeleteEventMasterResult
	err    error
}

func NewDeleteEventMasterResultFromDict(data map[string]interface{}) DeleteEventMasterResult {
    return DeleteEventMasterResult {
        Item: NewEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteEventMasterResult) Pointer() *DeleteEventMasterResult {
    return &p
}

type DescribeTriggersResult struct {
    Items []Trigger `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeTriggersAsyncResult struct {
	result *DescribeTriggersResult
	err    error
}

func NewDescribeTriggersResultFromDict(data map[string]interface{}) DescribeTriggersResult {
    return DescribeTriggersResult {
        Items: CastTriggers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeTriggersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastTriggersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeTriggersResult) Pointer() *DescribeTriggersResult {
    return &p
}

type DescribeTriggersByUserIdResult struct {
    Items []Trigger `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeTriggersByUserIdAsyncResult struct {
	result *DescribeTriggersByUserIdResult
	err    error
}

func NewDescribeTriggersByUserIdResultFromDict(data map[string]interface{}) DescribeTriggersByUserIdResult {
    return DescribeTriggersByUserIdResult {
        Items: CastTriggers(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeTriggersByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastTriggersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeTriggersByUserIdResult) Pointer() *DescribeTriggersByUserIdResult {
    return &p
}

type GetTriggerResult struct {
    Item *Trigger `json:"item"`
}

type GetTriggerAsyncResult struct {
	result *GetTriggerResult
	err    error
}

func NewGetTriggerResultFromDict(data map[string]interface{}) GetTriggerResult {
    return GetTriggerResult {
        Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetTriggerResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetTriggerResult) Pointer() *GetTriggerResult {
    return &p
}

type GetTriggerByUserIdResult struct {
    Item *Trigger `json:"item"`
}

type GetTriggerByUserIdAsyncResult struct {
	result *GetTriggerByUserIdResult
	err    error
}

func NewGetTriggerByUserIdResultFromDict(data map[string]interface{}) GetTriggerByUserIdResult {
    return GetTriggerByUserIdResult {
        Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetTriggerByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetTriggerByUserIdResult) Pointer() *GetTriggerByUserIdResult {
    return &p
}

type TriggerByUserIdResult struct {
    Item *Trigger `json:"item"`
}

type TriggerByUserIdAsyncResult struct {
	result *TriggerByUserIdResult
	err    error
}

func NewTriggerByUserIdResultFromDict(data map[string]interface{}) TriggerByUserIdResult {
    return TriggerByUserIdResult {
        Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p TriggerByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p TriggerByUserIdResult) Pointer() *TriggerByUserIdResult {
    return &p
}

type DeleteTriggerResult struct {
    Item *Trigger `json:"item"`
}

type DeleteTriggerAsyncResult struct {
	result *DeleteTriggerResult
	err    error
}

func NewDeleteTriggerResultFromDict(data map[string]interface{}) DeleteTriggerResult {
    return DeleteTriggerResult {
        Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteTriggerResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteTriggerResult) Pointer() *DeleteTriggerResult {
    return &p
}

type DeleteTriggerByUserIdResult struct {
    Item *Trigger `json:"item"`
}

type DeleteTriggerByUserIdAsyncResult struct {
	result *DeleteTriggerByUserIdResult
	err    error
}

func NewDeleteTriggerByUserIdResultFromDict(data map[string]interface{}) DeleteTriggerByUserIdResult {
    return DeleteTriggerByUserIdResult {
        Item: NewTriggerFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteTriggerByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteTriggerByUserIdResult) Pointer() *DeleteTriggerByUserIdResult {
    return &p
}

type DescribeEventsResult struct {
    Items []Event `json:"items"`
}

type DescribeEventsAsyncResult struct {
	result *DescribeEventsResult
	err    error
}

func NewDescribeEventsResultFromDict(data map[string]interface{}) DescribeEventsResult {
    return DescribeEventsResult {
        Items: CastEvents(core.CastArray(data["items"])),
    }
}

func (p DescribeEventsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastEventsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeEventsResult) Pointer() *DescribeEventsResult {
    return &p
}

type DescribeEventsByUserIdResult struct {
    Items []Event `json:"items"`
}

type DescribeEventsByUserIdAsyncResult struct {
	result *DescribeEventsByUserIdResult
	err    error
}

func NewDescribeEventsByUserIdResultFromDict(data map[string]interface{}) DescribeEventsByUserIdResult {
    return DescribeEventsByUserIdResult {
        Items: CastEvents(core.CastArray(data["items"])),
    }
}

func (p DescribeEventsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastEventsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeEventsByUserIdResult) Pointer() *DescribeEventsByUserIdResult {
    return &p
}

type DescribeRawEventsResult struct {
    Items []Event `json:"items"`
}

type DescribeRawEventsAsyncResult struct {
	result *DescribeRawEventsResult
	err    error
}

func NewDescribeRawEventsResultFromDict(data map[string]interface{}) DescribeRawEventsResult {
    return DescribeRawEventsResult {
        Items: CastEvents(core.CastArray(data["items"])),
    }
}

func (p DescribeRawEventsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastEventsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeRawEventsResult) Pointer() *DescribeRawEventsResult {
    return &p
}

type GetEventResult struct {
    Item *Event `json:"item"`
}

type GetEventAsyncResult struct {
	result *GetEventResult
	err    error
}

func NewGetEventResultFromDict(data map[string]interface{}) GetEventResult {
    return GetEventResult {
        Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetEventResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetEventResult) Pointer() *GetEventResult {
    return &p
}

type GetEventByUserIdResult struct {
    Item *Event `json:"item"`
}

type GetEventByUserIdAsyncResult struct {
	result *GetEventByUserIdResult
	err    error
}

func NewGetEventByUserIdResultFromDict(data map[string]interface{}) GetEventByUserIdResult {
    return GetEventByUserIdResult {
        Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetEventByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetEventByUserIdResult) Pointer() *GetEventByUserIdResult {
    return &p
}

type GetRawEventResult struct {
    Item *Event `json:"item"`
}

type GetRawEventAsyncResult struct {
	result *GetRawEventResult
	err    error
}

func NewGetRawEventResultFromDict(data map[string]interface{}) GetRawEventResult {
    return GetRawEventResult {
        Item: NewEventFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRawEventResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRawEventResult) Pointer() *GetRawEventResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentEventMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p ExportMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p ExportMasterResult) Pointer() *ExportMasterResult {
    return &p
}

type GetCurrentEventMasterResult struct {
    Item *CurrentEventMaster `json:"item"`
}

type GetCurrentEventMasterAsyncResult struct {
	result *GetCurrentEventMasterResult
	err    error
}

func NewGetCurrentEventMasterResultFromDict(data map[string]interface{}) GetCurrentEventMasterResult {
    return GetCurrentEventMasterResult {
        Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentEventMasterResult) Pointer() *GetCurrentEventMasterResult {
    return &p
}

type UpdateCurrentEventMasterResult struct {
    Item *CurrentEventMaster `json:"item"`
}

type UpdateCurrentEventMasterAsyncResult struct {
	result *UpdateCurrentEventMasterResult
	err    error
}

func NewUpdateCurrentEventMasterResultFromDict(data map[string]interface{}) UpdateCurrentEventMasterResult {
    return UpdateCurrentEventMasterResult {
        Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentEventMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentEventMasterResult) Pointer() *UpdateCurrentEventMasterResult {
    return &p
}

type UpdateCurrentEventMasterFromGitHubResult struct {
    Item *CurrentEventMaster `json:"item"`
}

type UpdateCurrentEventMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentEventMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentEventMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentEventMasterFromGitHubResult {
    return UpdateCurrentEventMasterFromGitHubResult {
        Item: NewCurrentEventMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentEventMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentEventMasterFromGitHubResult) Pointer() *UpdateCurrentEventMasterFromGitHubResult {
    return &p
}