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

package enchant

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    Items []Namespace `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeNamespacesAsyncResult struct {
	result *DescribeNamespacesResult
	err    error
}

func NewDescribeNamespacesResultFromJson(data string) DescribeNamespacesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesResultFromDict(dict)
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

func NewCreateNamespaceResultFromJson(data string) CreateNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceResultFromDict(dict)
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

func NewGetNamespaceStatusResultFromJson(data string) GetNamespaceStatusResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusResultFromDict(dict)
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

func NewGetNamespaceResultFromJson(data string) GetNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceResultFromDict(dict)
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

func NewUpdateNamespaceResultFromJson(data string) UpdateNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceResultFromDict(dict)
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

func NewDeleteNamespaceResultFromJson(data string) DeleteNamespaceResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceResultFromDict(dict)
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

type DescribeBalanceParameterModelsResult struct {
    Items []BalanceParameterModel `json:"items"`
}

type DescribeBalanceParameterModelsAsyncResult struct {
	result *DescribeBalanceParameterModelsResult
	err    error
}

func NewDescribeBalanceParameterModelsResultFromJson(data string) DescribeBalanceParameterModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBalanceParameterModelsResultFromDict(dict)
}

func NewDescribeBalanceParameterModelsResultFromDict(data map[string]interface{}) DescribeBalanceParameterModelsResult {
    return DescribeBalanceParameterModelsResult {
        Items: CastBalanceParameterModels(core.CastArray(data["items"])),
    }
}

func (p DescribeBalanceParameterModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBalanceParameterModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeBalanceParameterModelsResult) Pointer() *DescribeBalanceParameterModelsResult {
    return &p
}

type GetBalanceParameterModelResult struct {
    Item *BalanceParameterModel `json:"item"`
}

type GetBalanceParameterModelAsyncResult struct {
	result *GetBalanceParameterModelResult
	err    error
}

func NewGetBalanceParameterModelResultFromJson(data string) GetBalanceParameterModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBalanceParameterModelResultFromDict(dict)
}

func NewGetBalanceParameterModelResultFromDict(data map[string]interface{}) GetBalanceParameterModelResult {
    return GetBalanceParameterModelResult {
        Item: NewBalanceParameterModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBalanceParameterModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBalanceParameterModelResult) Pointer() *GetBalanceParameterModelResult {
    return &p
}

type DescribeBalanceParameterModelMastersResult struct {
    Items []BalanceParameterModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBalanceParameterModelMastersAsyncResult struct {
	result *DescribeBalanceParameterModelMastersResult
	err    error
}

func NewDescribeBalanceParameterModelMastersResultFromJson(data string) DescribeBalanceParameterModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBalanceParameterModelMastersResultFromDict(dict)
}

func NewDescribeBalanceParameterModelMastersResultFromDict(data map[string]interface{}) DescribeBalanceParameterModelMastersResult {
    return DescribeBalanceParameterModelMastersResult {
        Items: CastBalanceParameterModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBalanceParameterModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBalanceParameterModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBalanceParameterModelMastersResult) Pointer() *DescribeBalanceParameterModelMastersResult {
    return &p
}

type CreateBalanceParameterModelMasterResult struct {
    Item *BalanceParameterModelMaster `json:"item"`
}

type CreateBalanceParameterModelMasterAsyncResult struct {
	result *CreateBalanceParameterModelMasterResult
	err    error
}

func NewCreateBalanceParameterModelMasterResultFromJson(data string) CreateBalanceParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateBalanceParameterModelMasterResultFromDict(dict)
}

func NewCreateBalanceParameterModelMasterResultFromDict(data map[string]interface{}) CreateBalanceParameterModelMasterResult {
    return CreateBalanceParameterModelMasterResult {
        Item: NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateBalanceParameterModelMasterResult) Pointer() *CreateBalanceParameterModelMasterResult {
    return &p
}

type GetBalanceParameterModelMasterResult struct {
    Item *BalanceParameterModelMaster `json:"item"`
}

type GetBalanceParameterModelMasterAsyncResult struct {
	result *GetBalanceParameterModelMasterResult
	err    error
}

func NewGetBalanceParameterModelMasterResultFromJson(data string) GetBalanceParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBalanceParameterModelMasterResultFromDict(dict)
}

func NewGetBalanceParameterModelMasterResultFromDict(data map[string]interface{}) GetBalanceParameterModelMasterResult {
    return GetBalanceParameterModelMasterResult {
        Item: NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBalanceParameterModelMasterResult) Pointer() *GetBalanceParameterModelMasterResult {
    return &p
}

type UpdateBalanceParameterModelMasterResult struct {
    Item *BalanceParameterModelMaster `json:"item"`
}

type UpdateBalanceParameterModelMasterAsyncResult struct {
	result *UpdateBalanceParameterModelMasterResult
	err    error
}

func NewUpdateBalanceParameterModelMasterResultFromJson(data string) UpdateBalanceParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateBalanceParameterModelMasterResultFromDict(dict)
}

func NewUpdateBalanceParameterModelMasterResultFromDict(data map[string]interface{}) UpdateBalanceParameterModelMasterResult {
    return UpdateBalanceParameterModelMasterResult {
        Item: NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateBalanceParameterModelMasterResult) Pointer() *UpdateBalanceParameterModelMasterResult {
    return &p
}

type DeleteBalanceParameterModelMasterResult struct {
    Item *BalanceParameterModelMaster `json:"item"`
}

type DeleteBalanceParameterModelMasterAsyncResult struct {
	result *DeleteBalanceParameterModelMasterResult
	err    error
}

func NewDeleteBalanceParameterModelMasterResultFromJson(data string) DeleteBalanceParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBalanceParameterModelMasterResultFromDict(dict)
}

func NewDeleteBalanceParameterModelMasterResultFromDict(data map[string]interface{}) DeleteBalanceParameterModelMasterResult {
    return DeleteBalanceParameterModelMasterResult {
        Item: NewBalanceParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteBalanceParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteBalanceParameterModelMasterResult) Pointer() *DeleteBalanceParameterModelMasterResult {
    return &p
}

type DescribeRarityParameterModelsResult struct {
    Items []RarityParameterModel `json:"items"`
}

type DescribeRarityParameterModelsAsyncResult struct {
	result *DescribeRarityParameterModelsResult
	err    error
}

func NewDescribeRarityParameterModelsResultFromJson(data string) DescribeRarityParameterModelsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRarityParameterModelsResultFromDict(dict)
}

func NewDescribeRarityParameterModelsResultFromDict(data map[string]interface{}) DescribeRarityParameterModelsResult {
    return DescribeRarityParameterModelsResult {
        Items: CastRarityParameterModels(core.CastArray(data["items"])),
    }
}

func (p DescribeRarityParameterModelsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRarityParameterModelsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeRarityParameterModelsResult) Pointer() *DescribeRarityParameterModelsResult {
    return &p
}

type GetRarityParameterModelResult struct {
    Item *RarityParameterModel `json:"item"`
}

type GetRarityParameterModelAsyncResult struct {
	result *GetRarityParameterModelResult
	err    error
}

func NewGetRarityParameterModelResultFromJson(data string) GetRarityParameterModelResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRarityParameterModelResultFromDict(dict)
}

func NewGetRarityParameterModelResultFromDict(data map[string]interface{}) GetRarityParameterModelResult {
    return GetRarityParameterModelResult {
        Item: NewRarityParameterModelFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRarityParameterModelResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRarityParameterModelResult) Pointer() *GetRarityParameterModelResult {
    return &p
}

type DescribeRarityParameterModelMastersResult struct {
    Items []RarityParameterModelMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeRarityParameterModelMastersAsyncResult struct {
	result *DescribeRarityParameterModelMastersResult
	err    error
}

func NewDescribeRarityParameterModelMastersResultFromJson(data string) DescribeRarityParameterModelMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRarityParameterModelMastersResultFromDict(dict)
}

func NewDescribeRarityParameterModelMastersResultFromDict(data map[string]interface{}) DescribeRarityParameterModelMastersResult {
    return DescribeRarityParameterModelMastersResult {
        Items: CastRarityParameterModelMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeRarityParameterModelMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRarityParameterModelMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeRarityParameterModelMastersResult) Pointer() *DescribeRarityParameterModelMastersResult {
    return &p
}

type CreateRarityParameterModelMasterResult struct {
    Item *RarityParameterModelMaster `json:"item"`
}

type CreateRarityParameterModelMasterAsyncResult struct {
	result *CreateRarityParameterModelMasterResult
	err    error
}

func NewCreateRarityParameterModelMasterResultFromJson(data string) CreateRarityParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateRarityParameterModelMasterResultFromDict(dict)
}

func NewCreateRarityParameterModelMasterResultFromDict(data map[string]interface{}) CreateRarityParameterModelMasterResult {
    return CreateRarityParameterModelMasterResult {
        Item: NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateRarityParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateRarityParameterModelMasterResult) Pointer() *CreateRarityParameterModelMasterResult {
    return &p
}

type GetRarityParameterModelMasterResult struct {
    Item *RarityParameterModelMaster `json:"item"`
}

type GetRarityParameterModelMasterAsyncResult struct {
	result *GetRarityParameterModelMasterResult
	err    error
}

func NewGetRarityParameterModelMasterResultFromJson(data string) GetRarityParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRarityParameterModelMasterResultFromDict(dict)
}

func NewGetRarityParameterModelMasterResultFromDict(data map[string]interface{}) GetRarityParameterModelMasterResult {
    return GetRarityParameterModelMasterResult {
        Item: NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRarityParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRarityParameterModelMasterResult) Pointer() *GetRarityParameterModelMasterResult {
    return &p
}

type UpdateRarityParameterModelMasterResult struct {
    Item *RarityParameterModelMaster `json:"item"`
}

type UpdateRarityParameterModelMasterAsyncResult struct {
	result *UpdateRarityParameterModelMasterResult
	err    error
}

func NewUpdateRarityParameterModelMasterResultFromJson(data string) UpdateRarityParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateRarityParameterModelMasterResultFromDict(dict)
}

func NewUpdateRarityParameterModelMasterResultFromDict(data map[string]interface{}) UpdateRarityParameterModelMasterResult {
    return UpdateRarityParameterModelMasterResult {
        Item: NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateRarityParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateRarityParameterModelMasterResult) Pointer() *UpdateRarityParameterModelMasterResult {
    return &p
}

type DeleteRarityParameterModelMasterResult struct {
    Item *RarityParameterModelMaster `json:"item"`
}

type DeleteRarityParameterModelMasterAsyncResult struct {
	result *DeleteRarityParameterModelMasterResult
	err    error
}

func NewDeleteRarityParameterModelMasterResultFromJson(data string) DeleteRarityParameterModelMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteRarityParameterModelMasterResultFromDict(dict)
}

func NewDeleteRarityParameterModelMasterResultFromDict(data map[string]interface{}) DeleteRarityParameterModelMasterResult {
    return DeleteRarityParameterModelMasterResult {
        Item: NewRarityParameterModelMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRarityParameterModelMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRarityParameterModelMasterResult) Pointer() *DeleteRarityParameterModelMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentParameterMaster `json:"item"`
}

type ExportMasterAsyncResult struct {
	result *ExportMasterResult
	err    error
}

func NewExportMasterResultFromJson(data string) ExportMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewExportMasterResultFromDict(dict)
}

func NewExportMasterResultFromDict(data map[string]interface{}) ExportMasterResult {
    return ExportMasterResult {
        Item: NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentParameterMasterResult struct {
    Item *CurrentParameterMaster `json:"item"`
}

type GetCurrentParameterMasterAsyncResult struct {
	result *GetCurrentParameterMasterResult
	err    error
}

func NewGetCurrentParameterMasterResultFromJson(data string) GetCurrentParameterMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentParameterMasterResultFromDict(dict)
}

func NewGetCurrentParameterMasterResultFromDict(data map[string]interface{}) GetCurrentParameterMasterResult {
    return GetCurrentParameterMasterResult {
        Item: NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentParameterMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentParameterMasterResult) Pointer() *GetCurrentParameterMasterResult {
    return &p
}

type UpdateCurrentParameterMasterResult struct {
    Item *CurrentParameterMaster `json:"item"`
}

type UpdateCurrentParameterMasterAsyncResult struct {
	result *UpdateCurrentParameterMasterResult
	err    error
}

func NewUpdateCurrentParameterMasterResultFromJson(data string) UpdateCurrentParameterMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentParameterMasterResultFromDict(dict)
}

func NewUpdateCurrentParameterMasterResultFromDict(data map[string]interface{}) UpdateCurrentParameterMasterResult {
    return UpdateCurrentParameterMasterResult {
        Item: NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentParameterMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentParameterMasterResult) Pointer() *UpdateCurrentParameterMasterResult {
    return &p
}

type UpdateCurrentParameterMasterFromGitHubResult struct {
    Item *CurrentParameterMaster `json:"item"`
}

type UpdateCurrentParameterMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentParameterMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentParameterMasterFromGitHubResultFromJson(data string) UpdateCurrentParameterMasterFromGitHubResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentParameterMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentParameterMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentParameterMasterFromGitHubResult {
    return UpdateCurrentParameterMasterFromGitHubResult {
        Item: NewCurrentParameterMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentParameterMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentParameterMasterFromGitHubResult) Pointer() *UpdateCurrentParameterMasterFromGitHubResult {
    return &p
}

type DescribeBalanceParameterStatusesResult struct {
    Items []BalanceParameterStatus `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBalanceParameterStatusesAsyncResult struct {
	result *DescribeBalanceParameterStatusesResult
	err    error
}

func NewDescribeBalanceParameterStatusesResultFromJson(data string) DescribeBalanceParameterStatusesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBalanceParameterStatusesResultFromDict(dict)
}

func NewDescribeBalanceParameterStatusesResultFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesResult {
    return DescribeBalanceParameterStatusesResult {
        Items: CastBalanceParameterStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBalanceParameterStatusesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBalanceParameterStatusesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBalanceParameterStatusesResult) Pointer() *DescribeBalanceParameterStatusesResult {
    return &p
}

type DescribeBalanceParameterStatusesByUserIdResult struct {
    Items []BalanceParameterStatus `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeBalanceParameterStatusesByUserIdAsyncResult struct {
	result *DescribeBalanceParameterStatusesByUserIdResult
	err    error
}

func NewDescribeBalanceParameterStatusesByUserIdResultFromJson(data string) DescribeBalanceParameterStatusesByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBalanceParameterStatusesByUserIdResultFromDict(dict)
}

func NewDescribeBalanceParameterStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeBalanceParameterStatusesByUserIdResult {
    return DescribeBalanceParameterStatusesByUserIdResult {
        Items: CastBalanceParameterStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeBalanceParameterStatusesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastBalanceParameterStatusesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeBalanceParameterStatusesByUserIdResult) Pointer() *DescribeBalanceParameterStatusesByUserIdResult {
    return &p
}

type GetBalanceParameterStatusResult struct {
    Item *BalanceParameterStatus `json:"item"`
}

type GetBalanceParameterStatusAsyncResult struct {
	result *GetBalanceParameterStatusResult
	err    error
}

func NewGetBalanceParameterStatusResultFromJson(data string) GetBalanceParameterStatusResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBalanceParameterStatusResultFromDict(dict)
}

func NewGetBalanceParameterStatusResultFromDict(data map[string]interface{}) GetBalanceParameterStatusResult {
    return GetBalanceParameterStatusResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBalanceParameterStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBalanceParameterStatusResult) Pointer() *GetBalanceParameterStatusResult {
    return &p
}

type GetBalanceParameterStatusByUserIdResult struct {
    Item *BalanceParameterStatus `json:"item"`
}

type GetBalanceParameterStatusByUserIdAsyncResult struct {
	result *GetBalanceParameterStatusByUserIdResult
	err    error
}

func NewGetBalanceParameterStatusByUserIdResultFromJson(data string) GetBalanceParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewGetBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) GetBalanceParameterStatusByUserIdResult {
    return GetBalanceParameterStatusByUserIdResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetBalanceParameterStatusByUserIdResult) Pointer() *GetBalanceParameterStatusByUserIdResult {
    return &p
}

type DeleteBalanceParameterStatusByUserIdResult struct {
    Item *BalanceParameterStatus `json:"item"`
}

type DeleteBalanceParameterStatusByUserIdAsyncResult struct {
	result *DeleteBalanceParameterStatusByUserIdResult
	err    error
}

func NewDeleteBalanceParameterStatusByUserIdResultFromJson(data string) DeleteBalanceParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewDeleteBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) DeleteBalanceParameterStatusByUserIdResult {
    return DeleteBalanceParameterStatusByUserIdResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteBalanceParameterStatusByUserIdResult) Pointer() *DeleteBalanceParameterStatusByUserIdResult {
    return &p
}

type ReDrawBalanceParameterStatusByUserIdResult struct {
    Item *BalanceParameterStatus `json:"item"`
    Old *BalanceParameterStatus `json:"old"`
}

type ReDrawBalanceParameterStatusByUserIdAsyncResult struct {
	result *ReDrawBalanceParameterStatusByUserIdResult
	err    error
}

func NewReDrawBalanceParameterStatusByUserIdResultFromJson(data string) ReDrawBalanceParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReDrawBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewReDrawBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByUserIdResult {
    return ReDrawBalanceParameterStatusByUserIdResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p ReDrawBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p ReDrawBalanceParameterStatusByUserIdResult) Pointer() *ReDrawBalanceParameterStatusByUserIdResult {
    return &p
}

type ReDrawBalanceParameterStatusByStampSheetResult struct {
    Item *BalanceParameterStatus `json:"item"`
    Old *BalanceParameterStatus `json:"old"`
}

type ReDrawBalanceParameterStatusByStampSheetAsyncResult struct {
	result *ReDrawBalanceParameterStatusByStampSheetResult
	err    error
}

func NewReDrawBalanceParameterStatusByStampSheetResultFromJson(data string) ReDrawBalanceParameterStatusByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReDrawBalanceParameterStatusByStampSheetResultFromDict(dict)
}

func NewReDrawBalanceParameterStatusByStampSheetResultFromDict(data map[string]interface{}) ReDrawBalanceParameterStatusByStampSheetResult {
    return ReDrawBalanceParameterStatusByStampSheetResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p ReDrawBalanceParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p ReDrawBalanceParameterStatusByStampSheetResult) Pointer() *ReDrawBalanceParameterStatusByStampSheetResult {
    return &p
}

type SetBalanceParameterStatusByUserIdResult struct {
    Item *BalanceParameterStatus `json:"item"`
    Old *BalanceParameterStatus `json:"old"`
}

type SetBalanceParameterStatusByUserIdAsyncResult struct {
	result *SetBalanceParameterStatusByUserIdResult
	err    error
}

func NewSetBalanceParameterStatusByUserIdResultFromJson(data string) SetBalanceParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetBalanceParameterStatusByUserIdResultFromDict(dict)
}

func NewSetBalanceParameterStatusByUserIdResultFromDict(data map[string]interface{}) SetBalanceParameterStatusByUserIdResult {
    return SetBalanceParameterStatusByUserIdResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p SetBalanceParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p SetBalanceParameterStatusByUserIdResult) Pointer() *SetBalanceParameterStatusByUserIdResult {
    return &p
}

type SetBalanceParameterStatusByStampSheetResult struct {
    Item *BalanceParameterStatus `json:"item"`
    Old *BalanceParameterStatus `json:"old"`
}

type SetBalanceParameterStatusByStampSheetAsyncResult struct {
	result *SetBalanceParameterStatusByStampSheetResult
	err    error
}

func NewSetBalanceParameterStatusByStampSheetResultFromJson(data string) SetBalanceParameterStatusByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetBalanceParameterStatusByStampSheetResultFromDict(dict)
}

func NewSetBalanceParameterStatusByStampSheetResultFromDict(data map[string]interface{}) SetBalanceParameterStatusByStampSheetResult {
    return SetBalanceParameterStatusByStampSheetResult {
        Item: NewBalanceParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewBalanceParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p SetBalanceParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p SetBalanceParameterStatusByStampSheetResult) Pointer() *SetBalanceParameterStatusByStampSheetResult {
    return &p
}

type DescribeRarityParameterStatusesResult struct {
    Items []RarityParameterStatus `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeRarityParameterStatusesAsyncResult struct {
	result *DescribeRarityParameterStatusesResult
	err    error
}

func NewDescribeRarityParameterStatusesResultFromJson(data string) DescribeRarityParameterStatusesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRarityParameterStatusesResultFromDict(dict)
}

func NewDescribeRarityParameterStatusesResultFromDict(data map[string]interface{}) DescribeRarityParameterStatusesResult {
    return DescribeRarityParameterStatusesResult {
        Items: CastRarityParameterStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeRarityParameterStatusesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRarityParameterStatusesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeRarityParameterStatusesResult) Pointer() *DescribeRarityParameterStatusesResult {
    return &p
}

type DescribeRarityParameterStatusesByUserIdResult struct {
    Items []RarityParameterStatus `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeRarityParameterStatusesByUserIdAsyncResult struct {
	result *DescribeRarityParameterStatusesByUserIdResult
	err    error
}

func NewDescribeRarityParameterStatusesByUserIdResultFromJson(data string) DescribeRarityParameterStatusesByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRarityParameterStatusesByUserIdResultFromDict(dict)
}

func NewDescribeRarityParameterStatusesByUserIdResultFromDict(data map[string]interface{}) DescribeRarityParameterStatusesByUserIdResult {
    return DescribeRarityParameterStatusesByUserIdResult {
        Items: CastRarityParameterStatuses(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeRarityParameterStatusesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRarityParameterStatusesFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeRarityParameterStatusesByUserIdResult) Pointer() *DescribeRarityParameterStatusesByUserIdResult {
    return &p
}

type GetRarityParameterStatusResult struct {
    Item *RarityParameterStatus `json:"item"`
}

type GetRarityParameterStatusAsyncResult struct {
	result *GetRarityParameterStatusResult
	err    error
}

func NewGetRarityParameterStatusResultFromJson(data string) GetRarityParameterStatusResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRarityParameterStatusResultFromDict(dict)
}

func NewGetRarityParameterStatusResultFromDict(data map[string]interface{}) GetRarityParameterStatusResult {
    return GetRarityParameterStatusResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRarityParameterStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRarityParameterStatusResult) Pointer() *GetRarityParameterStatusResult {
    return &p
}

type GetRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
}

type GetRarityParameterStatusByUserIdAsyncResult struct {
	result *GetRarityParameterStatusByUserIdResult
	err    error
}

func NewGetRarityParameterStatusByUserIdResultFromJson(data string) GetRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewGetRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) GetRarityParameterStatusByUserIdResult {
    return GetRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRarityParameterStatusByUserIdResult) Pointer() *GetRarityParameterStatusByUserIdResult {
    return &p
}

type DeleteRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
}

type DeleteRarityParameterStatusByUserIdAsyncResult struct {
	result *DeleteRarityParameterStatusByUserIdResult
	err    error
}

func NewDeleteRarityParameterStatusByUserIdResultFromJson(data string) DeleteRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewDeleteRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) DeleteRarityParameterStatusByUserIdResult {
    return DeleteRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRarityParameterStatusByUserIdResult) Pointer() *DeleteRarityParameterStatusByUserIdResult {
    return &p
}

type ReDrawRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type ReDrawRarityParameterStatusByUserIdAsyncResult struct {
	result *ReDrawRarityParameterStatusByUserIdResult
	err    error
}

func NewReDrawRarityParameterStatusByUserIdResultFromJson(data string) ReDrawRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReDrawRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewReDrawRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByUserIdResult {
    return ReDrawRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p ReDrawRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p ReDrawRarityParameterStatusByUserIdResult) Pointer() *ReDrawRarityParameterStatusByUserIdResult {
    return &p
}

type ReDrawRarityParameterStatusByStampSheetResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type ReDrawRarityParameterStatusByStampSheetAsyncResult struct {
	result *ReDrawRarityParameterStatusByStampSheetResult
	err    error
}

func NewReDrawRarityParameterStatusByStampSheetResultFromJson(data string) ReDrawRarityParameterStatusByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewReDrawRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewReDrawRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) ReDrawRarityParameterStatusByStampSheetResult {
    return ReDrawRarityParameterStatusByStampSheetResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p ReDrawRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p ReDrawRarityParameterStatusByStampSheetResult) Pointer() *ReDrawRarityParameterStatusByStampSheetResult {
    return &p
}

type AddRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type AddRarityParameterStatusByUserIdAsyncResult struct {
	result *AddRarityParameterStatusByUserIdResult
	err    error
}

func NewAddRarityParameterStatusByUserIdResultFromJson(data string) AddRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewAddRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) AddRarityParameterStatusByUserIdResult {
    return AddRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p AddRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p AddRarityParameterStatusByUserIdResult) Pointer() *AddRarityParameterStatusByUserIdResult {
    return &p
}

type AddRarityParameterStatusByStampSheetResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type AddRarityParameterStatusByStampSheetAsyncResult struct {
	result *AddRarityParameterStatusByStampSheetResult
	err    error
}

func NewAddRarityParameterStatusByStampSheetResultFromJson(data string) AddRarityParameterStatusByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAddRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewAddRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) AddRarityParameterStatusByStampSheetResult {
    return AddRarityParameterStatusByStampSheetResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p AddRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p AddRarityParameterStatusByStampSheetResult) Pointer() *AddRarityParameterStatusByStampSheetResult {
    return &p
}

type VerifyRarityParameterStatusResult struct {
    Item *RarityParameterStatus `json:"item"`
}

type VerifyRarityParameterStatusAsyncResult struct {
	result *VerifyRarityParameterStatusResult
	err    error
}

func NewVerifyRarityParameterStatusResultFromJson(data string) VerifyRarityParameterStatusResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyRarityParameterStatusResultFromDict(dict)
}

func NewVerifyRarityParameterStatusResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusResult {
    return VerifyRarityParameterStatusResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p VerifyRarityParameterStatusResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p VerifyRarityParameterStatusResult) Pointer() *VerifyRarityParameterStatusResult {
    return &p
}

type VerifyRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
}

type VerifyRarityParameterStatusByUserIdAsyncResult struct {
	result *VerifyRarityParameterStatusByUserIdResult
	err    error
}

func NewVerifyRarityParameterStatusByUserIdResultFromJson(data string) VerifyRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewVerifyRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusByUserIdResult {
    return VerifyRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p VerifyRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p VerifyRarityParameterStatusByUserIdResult) Pointer() *VerifyRarityParameterStatusByUserIdResult {
    return &p
}

type VerifyRarityParameterStatusByStampTaskResult struct {
    Item *RarityParameterStatus `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type VerifyRarityParameterStatusByStampTaskAsyncResult struct {
	result *VerifyRarityParameterStatusByStampTaskResult
	err    error
}

func NewVerifyRarityParameterStatusByStampTaskResultFromJson(data string) VerifyRarityParameterStatusByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewVerifyRarityParameterStatusByStampTaskResultFromDict(dict)
}

func NewVerifyRarityParameterStatusByStampTaskResultFromDict(data map[string]interface{}) VerifyRarityParameterStatusByStampTaskResult {
    return VerifyRarityParameterStatusByStampTaskResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p VerifyRarityParameterStatusByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p VerifyRarityParameterStatusByStampTaskResult) Pointer() *VerifyRarityParameterStatusByStampTaskResult {
    return &p
}

type SetRarityParameterStatusByUserIdResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type SetRarityParameterStatusByUserIdAsyncResult struct {
	result *SetRarityParameterStatusByUserIdResult
	err    error
}

func NewSetRarityParameterStatusByUserIdResultFromJson(data string) SetRarityParameterStatusByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetRarityParameterStatusByUserIdResultFromDict(dict)
}

func NewSetRarityParameterStatusByUserIdResultFromDict(data map[string]interface{}) SetRarityParameterStatusByUserIdResult {
    return SetRarityParameterStatusByUserIdResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p SetRarityParameterStatusByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p SetRarityParameterStatusByUserIdResult) Pointer() *SetRarityParameterStatusByUserIdResult {
    return &p
}

type SetRarityParameterStatusByStampSheetResult struct {
    Item *RarityParameterStatus `json:"item"`
    Old *RarityParameterStatus `json:"old"`
}

type SetRarityParameterStatusByStampSheetAsyncResult struct {
	result *SetRarityParameterStatusByStampSheetResult
	err    error
}

func NewSetRarityParameterStatusByStampSheetResultFromJson(data string) SetRarityParameterStatusByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSetRarityParameterStatusByStampSheetResultFromDict(dict)
}

func NewSetRarityParameterStatusByStampSheetResultFromDict(data map[string]interface{}) SetRarityParameterStatusByStampSheetResult {
    return SetRarityParameterStatusByStampSheetResult {
        Item: NewRarityParameterStatusFromDict(core.CastMap(data["item"])).Pointer(),
        Old: NewRarityParameterStatusFromDict(core.CastMap(data["old"])).Pointer(),
    }
}

func (p SetRarityParameterStatusByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "old": p.Old.ToDict(),
    }
}

func (p SetRarityParameterStatusByStampSheetResult) Pointer() *SetRarityParameterStatusByStampSheetResult {
    return &p
}