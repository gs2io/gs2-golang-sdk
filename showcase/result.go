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

package showcase

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

type DescribeSalesItemMastersResult struct {
    Items []SalesItemMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSalesItemMastersAsyncResult struct {
	result *DescribeSalesItemMastersResult
	err    error
}

func NewDescribeSalesItemMastersResultFromJson(data string) DescribeSalesItemMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSalesItemMastersResultFromDict(dict)
}

func NewDescribeSalesItemMastersResultFromDict(data map[string]interface{}) DescribeSalesItemMastersResult {
    return DescribeSalesItemMastersResult {
        Items: CastSalesItemMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSalesItemMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSalesItemMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSalesItemMastersResult) Pointer() *DescribeSalesItemMastersResult {
    return &p
}

type CreateSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type CreateSalesItemMasterAsyncResult struct {
	result *CreateSalesItemMasterResult
	err    error
}

func NewCreateSalesItemMasterResultFromJson(data string) CreateSalesItemMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSalesItemMasterResultFromDict(dict)
}

func NewCreateSalesItemMasterResultFromDict(data map[string]interface{}) CreateSalesItemMasterResult {
    return CreateSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSalesItemMasterResult) Pointer() *CreateSalesItemMasterResult {
    return &p
}

type GetSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type GetSalesItemMasterAsyncResult struct {
	result *GetSalesItemMasterResult
	err    error
}

func NewGetSalesItemMasterResultFromJson(data string) GetSalesItemMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSalesItemMasterResultFromDict(dict)
}

func NewGetSalesItemMasterResultFromDict(data map[string]interface{}) GetSalesItemMasterResult {
    return GetSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSalesItemMasterResult) Pointer() *GetSalesItemMasterResult {
    return &p
}

type UpdateSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type UpdateSalesItemMasterAsyncResult struct {
	result *UpdateSalesItemMasterResult
	err    error
}

func NewUpdateSalesItemMasterResultFromJson(data string) UpdateSalesItemMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSalesItemMasterResultFromDict(dict)
}

func NewUpdateSalesItemMasterResultFromDict(data map[string]interface{}) UpdateSalesItemMasterResult {
    return UpdateSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSalesItemMasterResult) Pointer() *UpdateSalesItemMasterResult {
    return &p
}

type DeleteSalesItemMasterResult struct {
    Item *SalesItemMaster `json:"item"`
}

type DeleteSalesItemMasterAsyncResult struct {
	result *DeleteSalesItemMasterResult
	err    error
}

func NewDeleteSalesItemMasterResultFromJson(data string) DeleteSalesItemMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSalesItemMasterResultFromDict(dict)
}

func NewDeleteSalesItemMasterResultFromDict(data map[string]interface{}) DeleteSalesItemMasterResult {
    return DeleteSalesItemMasterResult {
        Item: NewSalesItemMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSalesItemMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSalesItemMasterResult) Pointer() *DeleteSalesItemMasterResult {
    return &p
}

type DescribeSalesItemGroupMastersResult struct {
    Items []SalesItemGroupMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeSalesItemGroupMastersAsyncResult struct {
	result *DescribeSalesItemGroupMastersResult
	err    error
}

func NewDescribeSalesItemGroupMastersResultFromJson(data string) DescribeSalesItemGroupMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSalesItemGroupMastersResultFromDict(dict)
}

func NewDescribeSalesItemGroupMastersResultFromDict(data map[string]interface{}) DescribeSalesItemGroupMastersResult {
    return DescribeSalesItemGroupMastersResult {
        Items: CastSalesItemGroupMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeSalesItemGroupMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastSalesItemGroupMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeSalesItemGroupMastersResult) Pointer() *DescribeSalesItemGroupMastersResult {
    return &p
}

type CreateSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type CreateSalesItemGroupMasterAsyncResult struct {
	result *CreateSalesItemGroupMasterResult
	err    error
}

func NewCreateSalesItemGroupMasterResultFromJson(data string) CreateSalesItemGroupMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateSalesItemGroupMasterResultFromDict(dict)
}

func NewCreateSalesItemGroupMasterResultFromDict(data map[string]interface{}) CreateSalesItemGroupMasterResult {
    return CreateSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateSalesItemGroupMasterResult) Pointer() *CreateSalesItemGroupMasterResult {
    return &p
}

type GetSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type GetSalesItemGroupMasterAsyncResult struct {
	result *GetSalesItemGroupMasterResult
	err    error
}

func NewGetSalesItemGroupMasterResultFromJson(data string) GetSalesItemGroupMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSalesItemGroupMasterResultFromDict(dict)
}

func NewGetSalesItemGroupMasterResultFromDict(data map[string]interface{}) GetSalesItemGroupMasterResult {
    return GetSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetSalesItemGroupMasterResult) Pointer() *GetSalesItemGroupMasterResult {
    return &p
}

type UpdateSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type UpdateSalesItemGroupMasterAsyncResult struct {
	result *UpdateSalesItemGroupMasterResult
	err    error
}

func NewUpdateSalesItemGroupMasterResultFromJson(data string) UpdateSalesItemGroupMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateSalesItemGroupMasterResultFromDict(dict)
}

func NewUpdateSalesItemGroupMasterResultFromDict(data map[string]interface{}) UpdateSalesItemGroupMasterResult {
    return UpdateSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateSalesItemGroupMasterResult) Pointer() *UpdateSalesItemGroupMasterResult {
    return &p
}

type DeleteSalesItemGroupMasterResult struct {
    Item *SalesItemGroupMaster `json:"item"`
}

type DeleteSalesItemGroupMasterAsyncResult struct {
	result *DeleteSalesItemGroupMasterResult
	err    error
}

func NewDeleteSalesItemGroupMasterResultFromJson(data string) DeleteSalesItemGroupMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteSalesItemGroupMasterResultFromDict(dict)
}

func NewDeleteSalesItemGroupMasterResultFromDict(data map[string]interface{}) DeleteSalesItemGroupMasterResult {
    return DeleteSalesItemGroupMasterResult {
        Item: NewSalesItemGroupMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteSalesItemGroupMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteSalesItemGroupMasterResult) Pointer() *DeleteSalesItemGroupMasterResult {
    return &p
}

type DescribeShowcaseMastersResult struct {
    Items []ShowcaseMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeShowcaseMastersAsyncResult struct {
	result *DescribeShowcaseMastersResult
	err    error
}

func NewDescribeShowcaseMastersResultFromJson(data string) DescribeShowcaseMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcaseMastersResultFromDict(dict)
}

func NewDescribeShowcaseMastersResultFromDict(data map[string]interface{}) DescribeShowcaseMastersResult {
    return DescribeShowcaseMastersResult {
        Items: CastShowcaseMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeShowcaseMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcaseMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeShowcaseMastersResult) Pointer() *DescribeShowcaseMastersResult {
    return &p
}

type CreateShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type CreateShowcaseMasterAsyncResult struct {
	result *CreateShowcaseMasterResult
	err    error
}

func NewCreateShowcaseMasterResultFromJson(data string) CreateShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateShowcaseMasterResultFromDict(dict)
}

func NewCreateShowcaseMasterResultFromDict(data map[string]interface{}) CreateShowcaseMasterResult {
    return CreateShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateShowcaseMasterResult) Pointer() *CreateShowcaseMasterResult {
    return &p
}

type GetShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type GetShowcaseMasterAsyncResult struct {
	result *GetShowcaseMasterResult
	err    error
}

func NewGetShowcaseMasterResultFromJson(data string) GetShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseMasterResultFromDict(dict)
}

func NewGetShowcaseMasterResultFromDict(data map[string]interface{}) GetShowcaseMasterResult {
    return GetShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseMasterResult) Pointer() *GetShowcaseMasterResult {
    return &p
}

type UpdateShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type UpdateShowcaseMasterAsyncResult struct {
	result *UpdateShowcaseMasterResult
	err    error
}

func NewUpdateShowcaseMasterResultFromJson(data string) UpdateShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateShowcaseMasterResultFromDict(dict)
}

func NewUpdateShowcaseMasterResultFromDict(data map[string]interface{}) UpdateShowcaseMasterResult {
    return UpdateShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateShowcaseMasterResult) Pointer() *UpdateShowcaseMasterResult {
    return &p
}

type DeleteShowcaseMasterResult struct {
    Item *ShowcaseMaster `json:"item"`
}

type DeleteShowcaseMasterAsyncResult struct {
	result *DeleteShowcaseMasterResult
	err    error
}

func NewDeleteShowcaseMasterResultFromJson(data string) DeleteShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteShowcaseMasterResultFromDict(dict)
}

func NewDeleteShowcaseMasterResultFromDict(data map[string]interface{}) DeleteShowcaseMasterResult {
    return DeleteShowcaseMasterResult {
        Item: NewShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteShowcaseMasterResult) Pointer() *DeleteShowcaseMasterResult {
    return &p
}

type ExportMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
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
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
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

type GetCurrentShowcaseMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type GetCurrentShowcaseMasterAsyncResult struct {
	result *GetCurrentShowcaseMasterResult
	err    error
}

func NewGetCurrentShowcaseMasterResultFromJson(data string) GetCurrentShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetCurrentShowcaseMasterResultFromDict(dict)
}

func NewGetCurrentShowcaseMasterResultFromDict(data map[string]interface{}) GetCurrentShowcaseMasterResult {
    return GetCurrentShowcaseMasterResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetCurrentShowcaseMasterResult) Pointer() *GetCurrentShowcaseMasterResult {
    return &p
}

type UpdateCurrentShowcaseMasterResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type UpdateCurrentShowcaseMasterAsyncResult struct {
	result *UpdateCurrentShowcaseMasterResult
	err    error
}

func NewUpdateCurrentShowcaseMasterResultFromJson(data string) UpdateCurrentShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentShowcaseMasterResultFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterResult {
    return UpdateCurrentShowcaseMasterResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentShowcaseMasterResult) Pointer() *UpdateCurrentShowcaseMasterResult {
    return &p
}

type UpdateCurrentShowcaseMasterFromGitHubResult struct {
    Item *CurrentShowcaseMaster `json:"item"`
}

type UpdateCurrentShowcaseMasterFromGitHubAsyncResult struct {
	result *UpdateCurrentShowcaseMasterFromGitHubResult
	err    error
}

func NewUpdateCurrentShowcaseMasterFromGitHubResultFromJson(data string) UpdateCurrentShowcaseMasterFromGitHubResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateCurrentShowcaseMasterFromGitHubResultFromDict(dict)
}

func NewUpdateCurrentShowcaseMasterFromGitHubResultFromDict(data map[string]interface{}) UpdateCurrentShowcaseMasterFromGitHubResult {
    return UpdateCurrentShowcaseMasterFromGitHubResult {
        Item: NewCurrentShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateCurrentShowcaseMasterFromGitHubResult) Pointer() *UpdateCurrentShowcaseMasterFromGitHubResult {
    return &p
}

type DescribeShowcasesResult struct {
    Items []Showcase `json:"items"`
}

type DescribeShowcasesAsyncResult struct {
	result *DescribeShowcasesResult
	err    error
}

func NewDescribeShowcasesResultFromJson(data string) DescribeShowcasesResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcasesResultFromDict(dict)
}

func NewDescribeShowcasesResultFromDict(data map[string]interface{}) DescribeShowcasesResult {
    return DescribeShowcasesResult {
        Items: CastShowcases(core.CastArray(data["items"])),
    }
}

func (p DescribeShowcasesResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcasesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeShowcasesResult) Pointer() *DescribeShowcasesResult {
    return &p
}

type DescribeShowcasesByUserIdResult struct {
    Items []Showcase `json:"items"`
}

type DescribeShowcasesByUserIdAsyncResult struct {
	result *DescribeShowcasesByUserIdResult
	err    error
}

func NewDescribeShowcasesByUserIdResultFromJson(data string) DescribeShowcasesByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeShowcasesByUserIdResultFromDict(dict)
}

func NewDescribeShowcasesByUserIdResultFromDict(data map[string]interface{}) DescribeShowcasesByUserIdResult {
    return DescribeShowcasesByUserIdResult {
        Items: CastShowcases(core.CastArray(data["items"])),
    }
}

func (p DescribeShowcasesByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastShowcasesFromDict(
            p.Items,
        ),
    }
}

func (p DescribeShowcasesByUserIdResult) Pointer() *DescribeShowcasesByUserIdResult {
    return &p
}

type GetShowcaseResult struct {
    Item *Showcase `json:"item"`
}

type GetShowcaseAsyncResult struct {
	result *GetShowcaseResult
	err    error
}

func NewGetShowcaseResultFromJson(data string) GetShowcaseResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseResultFromDict(dict)
}

func NewGetShowcaseResultFromDict(data map[string]interface{}) GetShowcaseResult {
    return GetShowcaseResult {
        Item: NewShowcaseFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseResult) Pointer() *GetShowcaseResult {
    return &p
}

type GetShowcaseByUserIdResult struct {
    Item *Showcase `json:"item"`
}

type GetShowcaseByUserIdAsyncResult struct {
	result *GetShowcaseByUserIdResult
	err    error
}

func NewGetShowcaseByUserIdResultFromJson(data string) GetShowcaseByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetShowcaseByUserIdResultFromDict(dict)
}

func NewGetShowcaseByUserIdResultFromDict(data map[string]interface{}) GetShowcaseByUserIdResult {
    return GetShowcaseByUserIdResult {
        Item: NewShowcaseFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetShowcaseByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetShowcaseByUserIdResult) Pointer() *GetShowcaseByUserIdResult {
    return &p
}

type BuyResult struct {
    Item *SalesItem `json:"item"`
    TransactionId *string `json:"transactionId"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    AutoRunStampSheet *bool `json:"autoRunStampSheet"`
}

type BuyAsyncResult struct {
	result *BuyResult
	err    error
}

func NewBuyResultFromJson(data string) BuyResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBuyResultFromDict(dict)
}

func NewBuyResultFromDict(data map[string]interface{}) BuyResult {
    return BuyResult {
        Item: NewSalesItemFromDict(core.CastMap(data["item"])).Pointer(),
        TransactionId: core.CastString(data["transactionId"]),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        AutoRunStampSheet: core.CastBool(data["autoRunStampSheet"]),
    }
}

func (p BuyResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "transactionId": p.TransactionId,
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "autoRunStampSheet": p.AutoRunStampSheet,
    }
}

func (p BuyResult) Pointer() *BuyResult {
    return &p
}

type BuyByUserIdResult struct {
    Item *SalesItem `json:"item"`
    TransactionId *string `json:"transactionId"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    AutoRunStampSheet *bool `json:"autoRunStampSheet"`
}

type BuyByUserIdAsyncResult struct {
	result *BuyByUserIdResult
	err    error
}

func NewBuyByUserIdResultFromJson(data string) BuyByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewBuyByUserIdResultFromDict(dict)
}

func NewBuyByUserIdResultFromDict(data map[string]interface{}) BuyByUserIdResult {
    return BuyByUserIdResult {
        Item: NewSalesItemFromDict(core.CastMap(data["item"])).Pointer(),
        TransactionId: core.CastString(data["transactionId"]),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        AutoRunStampSheet: core.CastBool(data["autoRunStampSheet"]),
    }
}

func (p BuyByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "transactionId": p.TransactionId,
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "autoRunStampSheet": p.AutoRunStampSheet,
    }
}

func (p BuyByUserIdResult) Pointer() *BuyByUserIdResult {
    return &p
}

type DescribeRandomShowcaseMastersResult struct {
    Items []RandomShowcaseMaster `json:"items"`
    NextPageToken *string `json:"nextPageToken"`
}

type DescribeRandomShowcaseMastersAsyncResult struct {
	result *DescribeRandomShowcaseMastersResult
	err    error
}

func NewDescribeRandomShowcaseMastersResultFromJson(data string) DescribeRandomShowcaseMastersResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRandomShowcaseMastersResultFromDict(dict)
}

func NewDescribeRandomShowcaseMastersResultFromDict(data map[string]interface{}) DescribeRandomShowcaseMastersResult {
    return DescribeRandomShowcaseMastersResult {
        Items: CastRandomShowcaseMasters(core.CastArray(data["items"])),
        NextPageToken: core.CastString(data["nextPageToken"]),
    }
}

func (p DescribeRandomShowcaseMastersResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRandomShowcaseMastersFromDict(
            p.Items,
        ),
        "nextPageToken": p.NextPageToken,
    }
}

func (p DescribeRandomShowcaseMastersResult) Pointer() *DescribeRandomShowcaseMastersResult {
    return &p
}

type CreateRandomShowcaseMasterResult struct {
    Item *RandomShowcaseMaster `json:"item"`
}

type CreateRandomShowcaseMasterAsyncResult struct {
	result *CreateRandomShowcaseMasterResult
	err    error
}

func NewCreateRandomShowcaseMasterResultFromJson(data string) CreateRandomShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateRandomShowcaseMasterResultFromDict(dict)
}

func NewCreateRandomShowcaseMasterResultFromDict(data map[string]interface{}) CreateRandomShowcaseMasterResult {
    return CreateRandomShowcaseMasterResult {
        Item: NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p CreateRandomShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p CreateRandomShowcaseMasterResult) Pointer() *CreateRandomShowcaseMasterResult {
    return &p
}

type GetRandomShowcaseMasterResult struct {
    Item *RandomShowcaseMaster `json:"item"`
}

type GetRandomShowcaseMasterAsyncResult struct {
	result *GetRandomShowcaseMasterResult
	err    error
}

func NewGetRandomShowcaseMasterResultFromJson(data string) GetRandomShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRandomShowcaseMasterResultFromDict(dict)
}

func NewGetRandomShowcaseMasterResultFromDict(data map[string]interface{}) GetRandomShowcaseMasterResult {
    return GetRandomShowcaseMasterResult {
        Item: NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRandomShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRandomShowcaseMasterResult) Pointer() *GetRandomShowcaseMasterResult {
    return &p
}

type UpdateRandomShowcaseMasterResult struct {
    Item *RandomShowcaseMaster `json:"item"`
}

type UpdateRandomShowcaseMasterAsyncResult struct {
	result *UpdateRandomShowcaseMasterResult
	err    error
}

func NewUpdateRandomShowcaseMasterResultFromJson(data string) UpdateRandomShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateRandomShowcaseMasterResultFromDict(dict)
}

func NewUpdateRandomShowcaseMasterResultFromDict(data map[string]interface{}) UpdateRandomShowcaseMasterResult {
    return UpdateRandomShowcaseMasterResult {
        Item: NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p UpdateRandomShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p UpdateRandomShowcaseMasterResult) Pointer() *UpdateRandomShowcaseMasterResult {
    return &p
}

type DeleteRandomShowcaseMasterResult struct {
    Item *RandomShowcaseMaster `json:"item"`
}

type DeleteRandomShowcaseMasterAsyncResult struct {
	result *DeleteRandomShowcaseMasterResult
	err    error
}

func NewDeleteRandomShowcaseMasterResultFromJson(data string) DeleteRandomShowcaseMasterResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteRandomShowcaseMasterResultFromDict(dict)
}

func NewDeleteRandomShowcaseMasterResultFromDict(data map[string]interface{}) DeleteRandomShowcaseMasterResult {
    return DeleteRandomShowcaseMasterResult {
        Item: NewRandomShowcaseMasterFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p DeleteRandomShowcaseMasterResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p DeleteRandomShowcaseMasterResult) Pointer() *DeleteRandomShowcaseMasterResult {
    return &p
}

type DescribeRandomShowcaseSalesItemsResult struct {
    Items []RandomDisplayItem `json:"items"`
}

type DescribeRandomShowcaseSalesItemsAsyncResult struct {
	result *DescribeRandomShowcaseSalesItemsResult
	err    error
}

func NewDescribeRandomShowcaseSalesItemsResultFromJson(data string) DescribeRandomShowcaseSalesItemsResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRandomShowcaseSalesItemsResultFromDict(dict)
}

func NewDescribeRandomShowcaseSalesItemsResultFromDict(data map[string]interface{}) DescribeRandomShowcaseSalesItemsResult {
    return DescribeRandomShowcaseSalesItemsResult {
        Items: CastRandomDisplayItems(core.CastArray(data["items"])),
    }
}

func (p DescribeRandomShowcaseSalesItemsResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRandomDisplayItemsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeRandomShowcaseSalesItemsResult) Pointer() *DescribeRandomShowcaseSalesItemsResult {
    return &p
}

type DescribeRandomShowcaseSalesItemsByUserIdResult struct {
    Items []RandomDisplayItem `json:"items"`
}

type DescribeRandomShowcaseSalesItemsByUserIdAsyncResult struct {
	result *DescribeRandomShowcaseSalesItemsByUserIdResult
	err    error
}

func NewDescribeRandomShowcaseSalesItemsByUserIdResultFromJson(data string) DescribeRandomShowcaseSalesItemsByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeRandomShowcaseSalesItemsByUserIdResultFromDict(dict)
}

func NewDescribeRandomShowcaseSalesItemsByUserIdResultFromDict(data map[string]interface{}) DescribeRandomShowcaseSalesItemsByUserIdResult {
    return DescribeRandomShowcaseSalesItemsByUserIdResult {
        Items: CastRandomDisplayItems(core.CastArray(data["items"])),
    }
}

func (p DescribeRandomShowcaseSalesItemsByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRandomDisplayItemsFromDict(
            p.Items,
        ),
    }
}

func (p DescribeRandomShowcaseSalesItemsByUserIdResult) Pointer() *DescribeRandomShowcaseSalesItemsByUserIdResult {
    return &p
}

type GetRandomShowcaseSalesItemResult struct {
    Item *RandomDisplayItem `json:"item"`
}

type GetRandomShowcaseSalesItemAsyncResult struct {
	result *GetRandomShowcaseSalesItemResult
	err    error
}

func NewGetRandomShowcaseSalesItemResultFromJson(data string) GetRandomShowcaseSalesItemResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRandomShowcaseSalesItemResultFromDict(dict)
}

func NewGetRandomShowcaseSalesItemResultFromDict(data map[string]interface{}) GetRandomShowcaseSalesItemResult {
    return GetRandomShowcaseSalesItemResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRandomShowcaseSalesItemResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRandomShowcaseSalesItemResult) Pointer() *GetRandomShowcaseSalesItemResult {
    return &p
}

type GetRandomShowcaseSalesItemByUserIdResult struct {
    Item *RandomDisplayItem `json:"item"`
}

type GetRandomShowcaseSalesItemByUserIdAsyncResult struct {
	result *GetRandomShowcaseSalesItemByUserIdResult
	err    error
}

func NewGetRandomShowcaseSalesItemByUserIdResultFromJson(data string) GetRandomShowcaseSalesItemByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetRandomShowcaseSalesItemByUserIdResultFromDict(dict)
}

func NewGetRandomShowcaseSalesItemByUserIdResultFromDict(data map[string]interface{}) GetRandomShowcaseSalesItemByUserIdResult {
    return GetRandomShowcaseSalesItemByUserIdResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p GetRandomShowcaseSalesItemByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p GetRandomShowcaseSalesItemByUserIdResult) Pointer() *GetRandomShowcaseSalesItemByUserIdResult {
    return &p
}

type IncrementPurchaseCountByUserIdResult struct {
    Item *RandomDisplayItem `json:"item"`
}

type IncrementPurchaseCountByUserIdAsyncResult struct {
	result *IncrementPurchaseCountByUserIdResult
	err    error
}

func NewIncrementPurchaseCountByUserIdResultFromJson(data string) IncrementPurchaseCountByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIncrementPurchaseCountByUserIdResultFromDict(dict)
}

func NewIncrementPurchaseCountByUserIdResultFromDict(data map[string]interface{}) IncrementPurchaseCountByUserIdResult {
    return IncrementPurchaseCountByUserIdResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
    }
}

func (p IncrementPurchaseCountByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
    }
}

func (p IncrementPurchaseCountByUserIdResult) Pointer() *IncrementPurchaseCountByUserIdResult {
    return &p
}

type IncrementPurchaseCountByStampTaskResult struct {
    Item *RandomDisplayItem `json:"item"`
    NewContextStack *string `json:"newContextStack"`
}

type IncrementPurchaseCountByStampTaskAsyncResult struct {
	result *IncrementPurchaseCountByStampTaskResult
	err    error
}

func NewIncrementPurchaseCountByStampTaskResultFromJson(data string) IncrementPurchaseCountByStampTaskResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIncrementPurchaseCountByStampTaskResultFromDict(dict)
}

func NewIncrementPurchaseCountByStampTaskResultFromDict(data map[string]interface{}) IncrementPurchaseCountByStampTaskResult {
    return IncrementPurchaseCountByStampTaskResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
        NewContextStack: core.CastString(data["newContextStack"]),
    }
}

func (p IncrementPurchaseCountByStampTaskResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "newContextStack": p.NewContextStack,
    }
}

func (p IncrementPurchaseCountByStampTaskResult) Pointer() *IncrementPurchaseCountByStampTaskResult {
    return &p
}

type ForceReDrawByUserIdResult struct {
    Items []RandomDisplayItem `json:"items"`
}

type ForceReDrawByUserIdAsyncResult struct {
	result *ForceReDrawByUserIdResult
	err    error
}

func NewForceReDrawByUserIdResultFromJson(data string) ForceReDrawByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewForceReDrawByUserIdResultFromDict(dict)
}

func NewForceReDrawByUserIdResultFromDict(data map[string]interface{}) ForceReDrawByUserIdResult {
    return ForceReDrawByUserIdResult {
        Items: CastRandomDisplayItems(core.CastArray(data["items"])),
    }
}

func (p ForceReDrawByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRandomDisplayItemsFromDict(
            p.Items,
        ),
    }
}

func (p ForceReDrawByUserIdResult) Pointer() *ForceReDrawByUserIdResult {
    return &p
}

type ForceReDrawByUserIdByStampSheetResult struct {
    Items []RandomDisplayItem `json:"items"`
}

type ForceReDrawByUserIdByStampSheetAsyncResult struct {
	result *ForceReDrawByUserIdByStampSheetResult
	err    error
}

func NewForceReDrawByUserIdByStampSheetResultFromJson(data string) ForceReDrawByUserIdByStampSheetResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewForceReDrawByUserIdByStampSheetResultFromDict(dict)
}

func NewForceReDrawByUserIdByStampSheetResultFromDict(data map[string]interface{}) ForceReDrawByUserIdByStampSheetResult {
    return ForceReDrawByUserIdByStampSheetResult {
        Items: CastRandomDisplayItems(core.CastArray(data["items"])),
    }
}

func (p ForceReDrawByUserIdByStampSheetResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "items": CastRandomDisplayItemsFromDict(
            p.Items,
        ),
    }
}

func (p ForceReDrawByUserIdByStampSheetResult) Pointer() *ForceReDrawByUserIdByStampSheetResult {
    return &p
}

type RandomShowcaseBuyResult struct {
    Item *RandomDisplayItem `json:"item"`
    TransactionId *string `json:"transactionId"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    AutoRunStampSheet *bool `json:"autoRunStampSheet"`
}

type RandomShowcaseBuyAsyncResult struct {
	result *RandomShowcaseBuyResult
	err    error
}

func NewRandomShowcaseBuyResultFromJson(data string) RandomShowcaseBuyResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRandomShowcaseBuyResultFromDict(dict)
}

func NewRandomShowcaseBuyResultFromDict(data map[string]interface{}) RandomShowcaseBuyResult {
    return RandomShowcaseBuyResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
        TransactionId: core.CastString(data["transactionId"]),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        AutoRunStampSheet: core.CastBool(data["autoRunStampSheet"]),
    }
}

func (p RandomShowcaseBuyResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "transactionId": p.TransactionId,
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "autoRunStampSheet": p.AutoRunStampSheet,
    }
}

func (p RandomShowcaseBuyResult) Pointer() *RandomShowcaseBuyResult {
    return &p
}

type RandomShowcaseBuyByUserIdResult struct {
    Item *RandomDisplayItem `json:"item"`
    TransactionId *string `json:"transactionId"`
    StampSheet *string `json:"stampSheet"`
    StampSheetEncryptionKeyId *string `json:"stampSheetEncryptionKeyId"`
    AutoRunStampSheet *bool `json:"autoRunStampSheet"`
}

type RandomShowcaseBuyByUserIdAsyncResult struct {
	result *RandomShowcaseBuyByUserIdResult
	err    error
}

func NewRandomShowcaseBuyByUserIdResultFromJson(data string) RandomShowcaseBuyByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRandomShowcaseBuyByUserIdResultFromDict(dict)
}

func NewRandomShowcaseBuyByUserIdResultFromDict(data map[string]interface{}) RandomShowcaseBuyByUserIdResult {
    return RandomShowcaseBuyByUserIdResult {
        Item: NewRandomDisplayItemFromDict(core.CastMap(data["item"])).Pointer(),
        TransactionId: core.CastString(data["transactionId"]),
        StampSheet: core.CastString(data["stampSheet"]),
        StampSheetEncryptionKeyId: core.CastString(data["stampSheetEncryptionKeyId"]),
        AutoRunStampSheet: core.CastBool(data["autoRunStampSheet"]),
    }
}

func (p RandomShowcaseBuyByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "item": p.Item.ToDict(),
        "transactionId": p.TransactionId,
        "stampSheet": p.StampSheet,
        "stampSheetEncryptionKeyId": p.StampSheetEncryptionKeyId,
        "autoRunStampSheet": p.AutoRunStampSheet,
    }
}

func (p RandomShowcaseBuyByUserIdResult) Pointer() *RandomShowcaseBuyByUserIdResult {
    return &p
}