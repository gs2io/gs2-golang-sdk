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

package formation

import "core"

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromDict(data map[string]interface{}) DescribeNamespacesRequest {
    return DescribeNamespacesRequest {
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeNamespacesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeNamespacesRequest) Pointer() *DescribeNamespacesRequest {
    return &p
}

type CreateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    UpdateMoldScript *ScriptSetting `json:"updateMoldScript"`
    UpdateFormScript *ScriptSetting `json:"updateFormScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        UpdateMoldScript: NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
        UpdateFormScript: NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "updateMoldScript": p.UpdateMoldScript.ToDict(),
        "updateFormScript": p.UpdateFormScript.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromDict(data map[string]interface{}) GetNamespaceStatusRequest {
    return GetNamespaceStatusRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceStatusRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceStatusRequest) Pointer() *GetNamespaceStatusRequest {
    return &p
}

type GetNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromDict(data map[string]interface{}) GetNamespaceRequest {
    return GetNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetNamespaceRequest) Pointer() *GetNamespaceRequest {
    return &p
}

type UpdateNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    UpdateMoldScript *ScriptSetting `json:"updateMoldScript"`
    UpdateFormScript *ScriptSetting `json:"updateFormScript"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        UpdateMoldScript: NewScriptSettingFromDict(core.CastMap(data["updateMoldScript"])).Pointer(),
        UpdateFormScript: NewScriptSettingFromDict(core.CastMap(data["updateFormScript"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "updateMoldScript": p.UpdateMoldScript.ToDict(),
        "updateFormScript": p.UpdateFormScript.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromDict(data map[string]interface{}) DeleteNamespaceRequest {
    return DeleteNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DeleteNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DeleteNamespaceRequest) Pointer() *DeleteNamespaceRequest {
    return &p
}

type DescribeFormModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFormModelMastersRequestFromDict(data map[string]interface{}) DescribeFormModelMastersRequest {
    return DescribeFormModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFormModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFormModelMastersRequest) Pointer() *DescribeFormModelMastersRequest {
    return &p
}

type CreateFormModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Slots []SlotModel `json:"slots"`
}

func NewCreateFormModelMasterRequestFromDict(data map[string]interface{}) CreateFormModelMasterRequest {
    return CreateFormModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Slots: CastSlotModels(core.CastArray(data["slots"])),
    }
}

func (p CreateFormModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "slots": CastSlotModelsFromDict(
            p.Slots,
        ),
    }
}

func (p CreateFormModelMasterRequest) Pointer() *CreateFormModelMasterRequest {
    return &p
}

type GetFormModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    FormModelName *string `json:"formModelName"`
}

func NewGetFormModelMasterRequestFromDict(data map[string]interface{}) GetFormModelMasterRequest {
    return GetFormModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        FormModelName: core.CastString(data["formModelName"]),
    }
}

func (p GetFormModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "formModelName": p.FormModelName,
    }
}

func (p GetFormModelMasterRequest) Pointer() *GetFormModelMasterRequest {
    return &p
}

type UpdateFormModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    FormModelName *string `json:"formModelName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Slots []SlotModel `json:"slots"`
}

func NewUpdateFormModelMasterRequestFromDict(data map[string]interface{}) UpdateFormModelMasterRequest {
    return UpdateFormModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        FormModelName: core.CastString(data["formModelName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Slots: CastSlotModels(core.CastArray(data["slots"])),
    }
}

func (p UpdateFormModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "formModelName": p.FormModelName,
        "description": p.Description,
        "metadata": p.Metadata,
        "slots": CastSlotModelsFromDict(
            p.Slots,
        ),
    }
}

func (p UpdateFormModelMasterRequest) Pointer() *UpdateFormModelMasterRequest {
    return &p
}

type DeleteFormModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    FormModelName *string `json:"formModelName"`
}

func NewDeleteFormModelMasterRequestFromDict(data map[string]interface{}) DeleteFormModelMasterRequest {
    return DeleteFormModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        FormModelName: core.CastString(data["formModelName"]),
    }
}

func (p DeleteFormModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "formModelName": p.FormModelName,
    }
}

func (p DeleteFormModelMasterRequest) Pointer() *DeleteFormModelMasterRequest {
    return &p
}

type DescribeMoldModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeMoldModelsRequestFromDict(data map[string]interface{}) DescribeMoldModelsRequest {
    return DescribeMoldModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeMoldModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeMoldModelsRequest) Pointer() *DescribeMoldModelsRequest {
    return &p
}

type GetMoldModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
}

func NewGetMoldModelRequestFromDict(data map[string]interface{}) GetMoldModelRequest {
    return GetMoldModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p GetMoldModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
    }
}

func (p GetMoldModelRequest) Pointer() *GetMoldModelRequest {
    return &p
}

type DescribeMoldModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeMoldModelMastersRequestFromDict(data map[string]interface{}) DescribeMoldModelMastersRequest {
    return DescribeMoldModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeMoldModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeMoldModelMastersRequest) Pointer() *DescribeMoldModelMastersRequest {
    return &p
}

type CreateMoldModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    FormModelName *string `json:"formModelName"`
    InitialMaxCapacity *int32 `json:"initialMaxCapacity"`
    MaxCapacity *int32 `json:"maxCapacity"`
}

func NewCreateMoldModelMasterRequestFromDict(data map[string]interface{}) CreateMoldModelMasterRequest {
    return CreateMoldModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        FormModelName: core.CastString(data["formModelName"]),
        InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
    }
}

func (p CreateMoldModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "formModelName": p.FormModelName,
        "initialMaxCapacity": p.InitialMaxCapacity,
        "maxCapacity": p.MaxCapacity,
    }
}

func (p CreateMoldModelMasterRequest) Pointer() *CreateMoldModelMasterRequest {
    return &p
}

type GetMoldModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
}

func NewGetMoldModelMasterRequestFromDict(data map[string]interface{}) GetMoldModelMasterRequest {
    return GetMoldModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p GetMoldModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
    }
}

func (p GetMoldModelMasterRequest) Pointer() *GetMoldModelMasterRequest {
    return &p
}

type UpdateMoldModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    FormModelName *string `json:"formModelName"`
    InitialMaxCapacity *int32 `json:"initialMaxCapacity"`
    MaxCapacity *int32 `json:"maxCapacity"`
}

func NewUpdateMoldModelMasterRequestFromDict(data map[string]interface{}) UpdateMoldModelMasterRequest {
    return UpdateMoldModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        FormModelName: core.CastString(data["formModelName"]),
        InitialMaxCapacity: core.CastInt32(data["initialMaxCapacity"]),
        MaxCapacity: core.CastInt32(data["maxCapacity"]),
    }
}

func (p UpdateMoldModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
        "description": p.Description,
        "metadata": p.Metadata,
        "formModelName": p.FormModelName,
        "initialMaxCapacity": p.InitialMaxCapacity,
        "maxCapacity": p.MaxCapacity,
    }
}

func (p UpdateMoldModelMasterRequest) Pointer() *UpdateMoldModelMasterRequest {
    return &p
}

type DeleteMoldModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
}

func NewDeleteMoldModelMasterRequestFromDict(data map[string]interface{}) DeleteMoldModelMasterRequest {
    return DeleteMoldModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p DeleteMoldModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
    }
}

func (p DeleteMoldModelMasterRequest) Pointer() *DeleteMoldModelMasterRequest {
    return &p
}

type ExportMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewExportMasterRequestFromDict(data map[string]interface{}) ExportMasterRequest {
    return ExportMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p ExportMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p ExportMasterRequest) Pointer() *ExportMasterRequest {
    return &p
}

type GetCurrentFormMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentFormMasterRequestFromDict(data map[string]interface{}) GetCurrentFormMasterRequest {
    return GetCurrentFormMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentFormMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentFormMasterRequest) Pointer() *GetCurrentFormMasterRequest {
    return &p
}

type UpdateCurrentFormMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentFormMasterRequestFromDict(data map[string]interface{}) UpdateCurrentFormMasterRequest {
    return UpdateCurrentFormMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentFormMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentFormMasterRequest) Pointer() *UpdateCurrentFormMasterRequest {
    return &p
}

type UpdateCurrentFormMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentFormMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentFormMasterFromGitHubRequest {
    return UpdateCurrentFormMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentFormMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentFormMasterFromGitHubRequest) Pointer() *UpdateCurrentFormMasterFromGitHubRequest {
    return &p
}

type DescribeMoldsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeMoldsRequestFromDict(data map[string]interface{}) DescribeMoldsRequest {
    return DescribeMoldsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeMoldsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeMoldsRequest) Pointer() *DescribeMoldsRequest {
    return &p
}

type DescribeMoldsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeMoldsByUserIdRequestFromDict(data map[string]interface{}) DescribeMoldsByUserIdRequest {
    return DescribeMoldsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeMoldsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeMoldsByUserIdRequest) Pointer() *DescribeMoldsByUserIdRequest {
    return &p
}

type GetMoldRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
}

func NewGetMoldRequestFromDict(data map[string]interface{}) GetMoldRequest {
    return GetMoldRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p GetMoldRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
    }
}

func (p GetMoldRequest) Pointer() *GetMoldRequest {
    return &p
}

type GetMoldByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
}

func NewGetMoldByUserIdRequestFromDict(data map[string]interface{}) GetMoldByUserIdRequest {
    return GetMoldByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p GetMoldByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
    }
}

func (p GetMoldByUserIdRequest) Pointer() *GetMoldByUserIdRequest {
    return &p
}

type SetMoldCapacityByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Capacity *int32 `json:"capacity"`
}

func NewSetMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) SetMoldCapacityByUserIdRequest {
    return SetMoldCapacityByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Capacity: core.CastInt32(data["capacity"]),
    }
}

func (p SetMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "capacity": p.Capacity,
    }
}

func (p SetMoldCapacityByUserIdRequest) Pointer() *SetMoldCapacityByUserIdRequest {
    return &p
}

type AddMoldCapacityByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Capacity *int32 `json:"capacity"`
}

func NewAddMoldCapacityByUserIdRequestFromDict(data map[string]interface{}) AddMoldCapacityByUserIdRequest {
    return AddMoldCapacityByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Capacity: core.CastInt32(data["capacity"]),
    }
}

func (p AddMoldCapacityByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "capacity": p.Capacity,
    }
}

func (p AddMoldCapacityByUserIdRequest) Pointer() *AddMoldCapacityByUserIdRequest {
    return &p
}

type DeleteMoldRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
}

func NewDeleteMoldRequestFromDict(data map[string]interface{}) DeleteMoldRequest {
    return DeleteMoldRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p DeleteMoldRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
    }
}

func (p DeleteMoldRequest) Pointer() *DeleteMoldRequest {
    return &p
}

type DeleteMoldByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
}

func NewDeleteMoldByUserIdRequestFromDict(data map[string]interface{}) DeleteMoldByUserIdRequest {
    return DeleteMoldByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
    }
}

func (p DeleteMoldByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
    }
}

func (p DeleteMoldByUserIdRequest) Pointer() *DeleteMoldByUserIdRequest {
    return &p
}

type AddCapacityByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewAddCapacityByStampSheetRequestFromDict(data map[string]interface{}) AddCapacityByStampSheetRequest {
    return AddCapacityByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p AddCapacityByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p AddCapacityByStampSheetRequest) Pointer() *AddCapacityByStampSheetRequest {
    return &p
}

type SetCapacityByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewSetCapacityByStampSheetRequestFromDict(data map[string]interface{}) SetCapacityByStampSheetRequest {
    return SetCapacityByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p SetCapacityByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p SetCapacityByStampSheetRequest) Pointer() *SetCapacityByStampSheetRequest {
    return &p
}

type DescribeFormsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFormsRequestFromDict(data map[string]interface{}) DescribeFormsRequest {
    return DescribeFormsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFormsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFormsRequest) Pointer() *DescribeFormsRequest {
    return &p
}

type DescribeFormsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    MoldName *string `json:"moldName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFormsByUserIdRequestFromDict(data map[string]interface{}) DescribeFormsByUserIdRequest {
    return DescribeFormsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        MoldName: core.CastString(data["moldName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFormsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "moldName": p.MoldName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFormsByUserIdRequest) Pointer() *DescribeFormsByUserIdRequest {
    return &p
}

type GetFormRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
}

func NewGetFormRequestFromDict(data map[string]interface{}) GetFormRequest {
    return GetFormRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
    }
}

func (p GetFormRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
        "index": p.Index,
    }
}

func (p GetFormRequest) Pointer() *GetFormRequest {
    return &p
}

type GetFormByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
}

func NewGetFormByUserIdRequestFromDict(data map[string]interface{}) GetFormByUserIdRequest {
    return GetFormByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
    }
}

func (p GetFormByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "index": p.Index,
    }
}

func (p GetFormByUserIdRequest) Pointer() *GetFormByUserIdRequest {
    return &p
}

type GetFormWithSignatureRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
    KeyId *string `json:"keyId"`
}

func NewGetFormWithSignatureRequestFromDict(data map[string]interface{}) GetFormWithSignatureRequest {
    return GetFormWithSignatureRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetFormWithSignatureRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
        "index": p.Index,
        "keyId": p.KeyId,
    }
}

func (p GetFormWithSignatureRequest) Pointer() *GetFormWithSignatureRequest {
    return &p
}

type GetFormWithSignatureByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
    KeyId *string `json:"keyId"`
}

func NewGetFormWithSignatureByUserIdRequestFromDict(data map[string]interface{}) GetFormWithSignatureByUserIdRequest {
    return GetFormWithSignatureByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetFormWithSignatureByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "index": p.Index,
        "keyId": p.KeyId,
    }
}

func (p GetFormWithSignatureByUserIdRequest) Pointer() *GetFormWithSignatureByUserIdRequest {
    return &p
}

type SetFormByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
    Slots []Slot `json:"slots"`
}

func NewSetFormByUserIdRequestFromDict(data map[string]interface{}) SetFormByUserIdRequest {
    return SetFormByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
        Slots: CastSlots(core.CastArray(data["slots"])),
    }
}

func (p SetFormByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "index": p.Index,
        "slots": CastSlotsFromDict(
            p.Slots,
        ),
    }
}

func (p SetFormByUserIdRequest) Pointer() *SetFormByUserIdRequest {
    return &p
}

type SetFormWithSignatureRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
    Slots []SlotWithSignature `json:"slots"`
    KeyId *string `json:"keyId"`
}

func NewSetFormWithSignatureRequestFromDict(data map[string]interface{}) SetFormWithSignatureRequest {
    return SetFormWithSignatureRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
        Slots: CastSlotWithSignatures(core.CastArray(data["slots"])),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p SetFormWithSignatureRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
        "index": p.Index,
        "slots": CastSlotWithSignaturesFromDict(
            p.Slots,
        ),
        "keyId": p.KeyId,
    }
}

func (p SetFormWithSignatureRequest) Pointer() *SetFormWithSignatureRequest {
    return &p
}

type AcquireActionsToFormPropertiesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
    AcquireAction *AcquireAction `json:"acquireAction"`
    QueueNamespaceId *string `json:"queueNamespaceId"`
    KeyId *string `json:"keyId"`
    Config []AcquireActionConfig `json:"config"`
}

func NewAcquireActionsToFormPropertiesRequestFromDict(data map[string]interface{}) AcquireActionsToFormPropertiesRequest {
    return AcquireActionsToFormPropertiesRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
        AcquireAction: NewAcquireActionFromDict(core.CastMap(data["acquireAction"])).Pointer(),
        QueueNamespaceId: core.CastString(data["queueNamespaceId"]),
        KeyId: core.CastString(data["keyId"]),
        Config: CastAcquireActionConfigs(core.CastArray(data["config"])),
    }
}

func (p AcquireActionsToFormPropertiesRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "index": p.Index,
        "acquireAction": p.AcquireAction.ToDict(),
        "queueNamespaceId": p.QueueNamespaceId,
        "keyId": p.KeyId,
        "config": CastAcquireActionConfigsFromDict(
            p.Config,
        ),
    }
}

func (p AcquireActionsToFormPropertiesRequest) Pointer() *AcquireActionsToFormPropertiesRequest {
    return &p
}

type DeleteFormRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
}

func NewDeleteFormRequestFromDict(data map[string]interface{}) DeleteFormRequest {
    return DeleteFormRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
    }
}

func (p DeleteFormRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "moldName": p.MoldName,
        "index": p.Index,
    }
}

func (p DeleteFormRequest) Pointer() *DeleteFormRequest {
    return &p
}

type DeleteFormByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    MoldName *string `json:"moldName"`
    Index *int32 `json:"index"`
}

func NewDeleteFormByUserIdRequestFromDict(data map[string]interface{}) DeleteFormByUserIdRequest {
    return DeleteFormByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        MoldName: core.CastString(data["moldName"]),
        Index: core.CastInt32(data["index"]),
    }
}

func (p DeleteFormByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "moldName": p.MoldName,
        "index": p.Index,
    }
}

func (p DeleteFormByUserIdRequest) Pointer() *DeleteFormByUserIdRequest {
    return &p
}

type AcquireActionToFormPropertiesByStampSheetRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    StampSheet *string `json:"stampSheet"`
    KeyId *string `json:"keyId"`
}

func NewAcquireActionToFormPropertiesByStampSheetRequestFromDict(data map[string]interface{}) AcquireActionToFormPropertiesByStampSheetRequest {
    return AcquireActionToFormPropertiesByStampSheetRequest {
        StampSheet: core.CastString(data["stampSheet"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p AcquireActionToFormPropertiesByStampSheetRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "stampSheet": p.StampSheet,
        "keyId": p.KeyId,
    }
}

func (p AcquireActionToFormPropertiesByStampSheetRequest) Pointer() *AcquireActionToFormPropertiesByStampSheetRequest {
    return &p
}