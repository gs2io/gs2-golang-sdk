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

package matchmaking

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
    EnableRating *bool `json:"enableRating"`
    CreateGatheringTriggerType *string `json:"createGatheringTriggerType"`
    CreateGatheringTriggerRealtimeNamespaceId *string `json:"createGatheringTriggerRealtimeNamespaceId"`
    CreateGatheringTriggerScriptId *string `json:"createGatheringTriggerScriptId"`
    CompleteMatchmakingTriggerType *string `json:"completeMatchmakingTriggerType"`
    CompleteMatchmakingTriggerRealtimeNamespaceId *string `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
    CompleteMatchmakingTriggerScriptId *string `json:"completeMatchmakingTriggerScriptId"`
    JoinNotification *NotificationSetting `json:"joinNotification"`
    LeaveNotification *NotificationSetting `json:"leaveNotification"`
    CompleteNotification *NotificationSetting `json:"completeNotification"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        EnableRating: core.CastBool(data["enableRating"]),
        CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
        CreateGatheringTriggerRealtimeNamespaceId: core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
        CreateGatheringTriggerScriptId: core.CastString(data["createGatheringTriggerScriptId"]),
        CompleteMatchmakingTriggerType: core.CastString(data["completeMatchmakingTriggerType"]),
        CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
        CompleteMatchmakingTriggerScriptId: core.CastString(data["completeMatchmakingTriggerScriptId"]),
        JoinNotification: NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
        LeaveNotification: NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
        CompleteNotification: NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "enableRating": p.EnableRating,
        "createGatheringTriggerType": p.CreateGatheringTriggerType,
        "createGatheringTriggerRealtimeNamespaceId": p.CreateGatheringTriggerRealtimeNamespaceId,
        "createGatheringTriggerScriptId": p.CreateGatheringTriggerScriptId,
        "completeMatchmakingTriggerType": p.CompleteMatchmakingTriggerType,
        "completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
        "completeMatchmakingTriggerScriptId": p.CompleteMatchmakingTriggerScriptId,
        "joinNotification": p.JoinNotification.ToDict(),
        "leaveNotification": p.LeaveNotification.ToDict(),
        "completeNotification": p.CompleteNotification.ToDict(),
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
    EnableRating *bool `json:"enableRating"`
    CreateGatheringTriggerType *string `json:"createGatheringTriggerType"`
    CreateGatheringTriggerRealtimeNamespaceId *string `json:"createGatheringTriggerRealtimeNamespaceId"`
    CreateGatheringTriggerScriptId *string `json:"createGatheringTriggerScriptId"`
    CompleteMatchmakingTriggerType *string `json:"completeMatchmakingTriggerType"`
    CompleteMatchmakingTriggerRealtimeNamespaceId *string `json:"completeMatchmakingTriggerRealtimeNamespaceId"`
    CompleteMatchmakingTriggerScriptId *string `json:"completeMatchmakingTriggerScriptId"`
    JoinNotification *NotificationSetting `json:"joinNotification"`
    LeaveNotification *NotificationSetting `json:"leaveNotification"`
    CompleteNotification *NotificationSetting `json:"completeNotification"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        EnableRating: core.CastBool(data["enableRating"]),
        CreateGatheringTriggerType: core.CastString(data["createGatheringTriggerType"]),
        CreateGatheringTriggerRealtimeNamespaceId: core.CastString(data["createGatheringTriggerRealtimeNamespaceId"]),
        CreateGatheringTriggerScriptId: core.CastString(data["createGatheringTriggerScriptId"]),
        CompleteMatchmakingTriggerType: core.CastString(data["completeMatchmakingTriggerType"]),
        CompleteMatchmakingTriggerRealtimeNamespaceId: core.CastString(data["completeMatchmakingTriggerRealtimeNamespaceId"]),
        CompleteMatchmakingTriggerScriptId: core.CastString(data["completeMatchmakingTriggerScriptId"]),
        JoinNotification: NewNotificationSettingFromDict(core.CastMap(data["joinNotification"])).Pointer(),
        LeaveNotification: NewNotificationSettingFromDict(core.CastMap(data["leaveNotification"])).Pointer(),
        CompleteNotification: NewNotificationSettingFromDict(core.CastMap(data["completeNotification"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "enableRating": p.EnableRating,
        "createGatheringTriggerType": p.CreateGatheringTriggerType,
        "createGatheringTriggerRealtimeNamespaceId": p.CreateGatheringTriggerRealtimeNamespaceId,
        "createGatheringTriggerScriptId": p.CreateGatheringTriggerScriptId,
        "completeMatchmakingTriggerType": p.CompleteMatchmakingTriggerType,
        "completeMatchmakingTriggerRealtimeNamespaceId": p.CompleteMatchmakingTriggerRealtimeNamespaceId,
        "completeMatchmakingTriggerScriptId": p.CompleteMatchmakingTriggerScriptId,
        "joinNotification": p.JoinNotification.ToDict(),
        "leaveNotification": p.LeaveNotification.ToDict(),
        "completeNotification": p.CompleteNotification.ToDict(),
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

type DescribeGatheringsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeGatheringsRequestFromDict(data map[string]interface{}) DescribeGatheringsRequest {
    return DescribeGatheringsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeGatheringsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeGatheringsRequest) Pointer() *DescribeGatheringsRequest {
    return &p
}

type CreateGatheringRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Player *Player `json:"player"`
    AttributeRanges []AttributeRange `json:"attributeRanges"`
    CapacityOfRoles []CapacityOfRole `json:"capacityOfRoles"`
    AllowUserIds []string `json:"allowUserIds"`
    ExpiresAt *int64 `json:"expiresAt"`
}

func NewCreateGatheringRequestFromDict(data map[string]interface{}) CreateGatheringRequest {
    return CreateGatheringRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Player: NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
        AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
        CapacityOfRoles: CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        ExpiresAt: core.CastInt64(data["expiresAt"]),
    }
}

func (p CreateGatheringRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "player": p.Player.ToDict(),
        "attributeRanges": CastAttributeRangesFromDict(
            p.AttributeRanges,
        ),
        "capacityOfRoles": CastCapacityOfRolesFromDict(
            p.CapacityOfRoles,
        ),
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
        "expiresAt": p.ExpiresAt,
    }
}

func (p CreateGatheringRequest) Pointer() *CreateGatheringRequest {
    return &p
}

type CreateGatheringByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    Player *Player `json:"player"`
    AttributeRanges []AttributeRange `json:"attributeRanges"`
    CapacityOfRoles []CapacityOfRole `json:"capacityOfRoles"`
    AllowUserIds []string `json:"allowUserIds"`
    ExpiresAt *int64 `json:"expiresAt"`
}

func NewCreateGatheringByUserIdRequestFromDict(data map[string]interface{}) CreateGatheringByUserIdRequest {
    return CreateGatheringByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        Player: NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
        AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
        CapacityOfRoles: CastCapacityOfRoles(core.CastArray(data["capacityOfRoles"])),
        AllowUserIds: core.CastStrings(core.CastArray(data["allowUserIds"])),
        ExpiresAt: core.CastInt64(data["expiresAt"]),
    }
}

func (p CreateGatheringByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "player": p.Player.ToDict(),
        "attributeRanges": CastAttributeRangesFromDict(
            p.AttributeRanges,
        ),
        "capacityOfRoles": CastCapacityOfRolesFromDict(
            p.CapacityOfRoles,
        ),
        "allowUserIds": core.CastStringsFromDict(
            p.AllowUserIds,
        ),
        "expiresAt": p.ExpiresAt,
    }
}

func (p CreateGatheringByUserIdRequest) Pointer() *CreateGatheringByUserIdRequest {
    return &p
}

type UpdateGatheringRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
    AccessToken *string `json:"accessToken"`
    AttributeRanges []AttributeRange `json:"attributeRanges"`
}

func NewUpdateGatheringRequestFromDict(data map[string]interface{}) UpdateGatheringRequest {
    return UpdateGatheringRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        AccessToken: core.CastString(data["accessToken"]),
        AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
    }
}

func (p UpdateGatheringRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
        "accessToken": p.AccessToken,
        "attributeRanges": CastAttributeRangesFromDict(
            p.AttributeRanges,
        ),
    }
}

func (p UpdateGatheringRequest) Pointer() *UpdateGatheringRequest {
    return &p
}

type UpdateGatheringByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
    UserId *string `json:"userId"`
    AttributeRanges []AttributeRange `json:"attributeRanges"`
}

func NewUpdateGatheringByUserIdRequestFromDict(data map[string]interface{}) UpdateGatheringByUserIdRequest {
    return UpdateGatheringByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        UserId: core.CastString(data["userId"]),
        AttributeRanges: CastAttributeRanges(core.CastArray(data["attributeRanges"])),
    }
}

func (p UpdateGatheringByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
        "userId": p.UserId,
        "attributeRanges": CastAttributeRangesFromDict(
            p.AttributeRanges,
        ),
    }
}

func (p UpdateGatheringByUserIdRequest) Pointer() *UpdateGatheringByUserIdRequest {
    return &p
}

type DoMatchmakingByPlayerRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Player *Player `json:"player"`
    MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func NewDoMatchmakingByPlayerRequestFromDict(data map[string]interface{}) DoMatchmakingByPlayerRequest {
    return DoMatchmakingByPlayerRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Player: NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
        MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
    }
}

func (p DoMatchmakingByPlayerRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "player": p.Player.ToDict(),
        "matchmakingContextToken": p.MatchmakingContextToken,
    }
}

func (p DoMatchmakingByPlayerRequest) Pointer() *DoMatchmakingByPlayerRequest {
    return &p
}

type DoMatchmakingRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    Player *Player `json:"player"`
    MatchmakingContextToken *string `json:"matchmakingContextToken"`
}

func NewDoMatchmakingRequestFromDict(data map[string]interface{}) DoMatchmakingRequest {
    return DoMatchmakingRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        Player: NewPlayerFromDict(core.CastMap(data["player"])).Pointer(),
        MatchmakingContextToken: core.CastString(data["matchmakingContextToken"]),
    }
}

func (p DoMatchmakingRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "player": p.Player.ToDict(),
        "matchmakingContextToken": p.MatchmakingContextToken,
    }
}

func (p DoMatchmakingRequest) Pointer() *DoMatchmakingRequest {
    return &p
}

type GetGatheringRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
}

func NewGetGatheringRequestFromDict(data map[string]interface{}) GetGatheringRequest {
    return GetGatheringRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
    }
}

func (p GetGatheringRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
    }
}

func (p GetGatheringRequest) Pointer() *GetGatheringRequest {
    return &p
}

type CancelMatchmakingRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
    AccessToken *string `json:"accessToken"`
}

func NewCancelMatchmakingRequestFromDict(data map[string]interface{}) CancelMatchmakingRequest {
    return CancelMatchmakingRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p CancelMatchmakingRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
        "accessToken": p.AccessToken,
    }
}

func (p CancelMatchmakingRequest) Pointer() *CancelMatchmakingRequest {
    return &p
}

type CancelMatchmakingByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
    UserId *string `json:"userId"`
}

func NewCancelMatchmakingByUserIdRequestFromDict(data map[string]interface{}) CancelMatchmakingByUserIdRequest {
    return CancelMatchmakingByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p CancelMatchmakingByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
        "userId": p.UserId,
    }
}

func (p CancelMatchmakingByUserIdRequest) Pointer() *CancelMatchmakingByUserIdRequest {
    return &p
}

type DeleteGatheringRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    GatheringName *string `json:"gatheringName"`
}

func NewDeleteGatheringRequestFromDict(data map[string]interface{}) DeleteGatheringRequest {
    return DeleteGatheringRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        GatheringName: core.CastString(data["gatheringName"]),
    }
}

func (p DeleteGatheringRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "gatheringName": p.GatheringName,
    }
}

func (p DeleteGatheringRequest) Pointer() *DeleteGatheringRequest {
    return &p
}

type DescribeRatingModelMastersRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRatingModelMastersRequestFromDict(data map[string]interface{}) DescribeRatingModelMastersRequest {
    return DescribeRatingModelMastersRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRatingModelMastersRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRatingModelMastersRequest) Pointer() *DescribeRatingModelMastersRequest {
    return &p
}

type CreateRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Volatility *int32 `json:"volatility"`
}

func NewCreateRatingModelMasterRequestFromDict(data map[string]interface{}) CreateRatingModelMasterRequest {
    return CreateRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Volatility: core.CastInt32(data["volatility"]),
    }
}

func (p CreateRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "name": p.Name,
        "description": p.Description,
        "metadata": p.Metadata,
        "volatility": p.Volatility,
    }
}

func (p CreateRatingModelMasterRequest) Pointer() *CreateRatingModelMasterRequest {
    return &p
}

type GetRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
}

func NewGetRatingModelMasterRequestFromDict(data map[string]interface{}) GetRatingModelMasterRequest {
    return GetRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p GetRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
    }
}

func (p GetRatingModelMasterRequest) Pointer() *GetRatingModelMasterRequest {
    return &p
}

type UpdateRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
    Description *string `json:"description"`
    Metadata *string `json:"metadata"`
    Volatility *int32 `json:"volatility"`
}

func NewUpdateRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateRatingModelMasterRequest {
    return UpdateRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
        Description: core.CastString(data["description"]),
        Metadata: core.CastString(data["metadata"]),
        Volatility: core.CastInt32(data["volatility"]),
    }
}

func (p UpdateRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
        "description": p.Description,
        "metadata": p.Metadata,
        "volatility": p.Volatility,
    }
}

func (p UpdateRatingModelMasterRequest) Pointer() *UpdateRatingModelMasterRequest {
    return &p
}

type DeleteRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
}

func NewDeleteRatingModelMasterRequestFromDict(data map[string]interface{}) DeleteRatingModelMasterRequest {
    return DeleteRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p DeleteRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
    }
}

func (p DeleteRatingModelMasterRequest) Pointer() *DeleteRatingModelMasterRequest {
    return &p
}

type DescribeRatingModelsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDescribeRatingModelsRequestFromDict(data map[string]interface{}) DescribeRatingModelsRequest {
    return DescribeRatingModelsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p DescribeRatingModelsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p DescribeRatingModelsRequest) Pointer() *DescribeRatingModelsRequest {
    return &p
}

type GetRatingModelRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
}

func NewGetRatingModelRequestFromDict(data map[string]interface{}) GetRatingModelRequest {
    return GetRatingModelRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p GetRatingModelRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
    }
}

func (p GetRatingModelRequest) Pointer() *GetRatingModelRequest {
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

type GetCurrentRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) GetCurrentRatingModelMasterRequest {
    return GetCurrentRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
    }
}

func (p GetCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
    }
}

func (p GetCurrentRatingModelMasterRequest) Pointer() *GetCurrentRatingModelMasterRequest {
    return &p
}

type UpdateCurrentRatingModelMasterRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    Settings *string `json:"settings"`
}

func NewUpdateCurrentRatingModelMasterRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterRequest {
    return UpdateCurrentRatingModelMasterRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Settings: core.CastString(data["settings"]),
    }
}

func (p UpdateCurrentRatingModelMasterRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "settings": p.Settings,
    }
}

func (p UpdateCurrentRatingModelMasterRequest) Pointer() *UpdateCurrentRatingModelMasterRequest {
    return &p
}

type UpdateCurrentRatingModelMasterFromGitHubRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    CheckoutSetting *GitHubCheckoutSetting `json:"checkoutSetting"`
}

func NewUpdateCurrentRatingModelMasterFromGitHubRequestFromDict(data map[string]interface{}) UpdateCurrentRatingModelMasterFromGitHubRequest {
    return UpdateCurrentRatingModelMasterFromGitHubRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        CheckoutSetting: NewGitHubCheckoutSettingFromDict(core.CastMap(data["checkoutSetting"])).Pointer(),
    }
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "checkoutSetting": p.CheckoutSetting.ToDict(),
    }
}

func (p UpdateCurrentRatingModelMasterFromGitHubRequest) Pointer() *UpdateCurrentRatingModelMasterFromGitHubRequest {
    return &p
}

type DescribeRatingsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRatingsRequestFromDict(data map[string]interface{}) DescribeRatingsRequest {
    return DescribeRatingsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRatingsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRatingsRequest) Pointer() *DescribeRatingsRequest {
    return &p
}

type DescribeRatingsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeRatingsByUserIdRequestFromDict(data map[string]interface{}) DescribeRatingsByUserIdRequest {
    return DescribeRatingsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeRatingsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeRatingsByUserIdRequest) Pointer() *DescribeRatingsByUserIdRequest {
    return &p
}

type GetRatingRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    RatingName *string `json:"ratingName"`
}

func NewGetRatingRequestFromDict(data map[string]interface{}) GetRatingRequest {
    return GetRatingRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p GetRatingRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "ratingName": p.RatingName,
    }
}

func (p GetRatingRequest) Pointer() *GetRatingRequest {
    return &p
}

type GetRatingByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RatingName *string `json:"ratingName"`
}

func NewGetRatingByUserIdRequestFromDict(data map[string]interface{}) GetRatingByUserIdRequest {
    return GetRatingByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p GetRatingByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "ratingName": p.RatingName,
    }
}

func (p GetRatingByUserIdRequest) Pointer() *GetRatingByUserIdRequest {
    return &p
}

type PutResultRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
    GameResults []GameResult `json:"gameResults"`
}

func NewPutResultRequestFromDict(data map[string]interface{}) PutResultRequest {
    return PutResultRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
        GameResults: CastGameResults(core.CastArray(data["gameResults"])),
    }
}

func (p PutResultRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
        "gameResults": CastGameResultsFromDict(
            p.GameResults,
        ),
    }
}

func (p PutResultRequest) Pointer() *PutResultRequest {
    return &p
}

type DeleteRatingRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    RatingName *string `json:"ratingName"`
}

func NewDeleteRatingRequestFromDict(data map[string]interface{}) DeleteRatingRequest {
    return DeleteRatingRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        RatingName: core.CastString(data["ratingName"]),
    }
}

func (p DeleteRatingRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "ratingName": p.RatingName,
    }
}

func (p DeleteRatingRequest) Pointer() *DeleteRatingRequest {
    return &p
}

type GetBallotRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
    GatheringName *string `json:"gatheringName"`
    AccessToken *string `json:"accessToken"`
    NumberOfPlayer *int32 `json:"numberOfPlayer"`
    KeyId *string `json:"keyId"`
}

func NewGetBallotRequestFromDict(data map[string]interface{}) GetBallotRequest {
    return GetBallotRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        AccessToken: core.CastString(data["accessToken"]),
        NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetBallotRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
        "gatheringName": p.GatheringName,
        "accessToken": p.AccessToken,
        "numberOfPlayer": p.NumberOfPlayer,
        "keyId": p.KeyId,
    }
}

func (p GetBallotRequest) Pointer() *GetBallotRequest {
    return &p
}

type GetBallotByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
    GatheringName *string `json:"gatheringName"`
    UserId *string `json:"userId"`
    NumberOfPlayer *int32 `json:"numberOfPlayer"`
    KeyId *string `json:"keyId"`
}

func NewGetBallotByUserIdRequestFromDict(data map[string]interface{}) GetBallotByUserIdRequest {
    return GetBallotByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
        GatheringName: core.CastString(data["gatheringName"]),
        UserId: core.CastString(data["userId"]),
        NumberOfPlayer: core.CastInt32(data["numberOfPlayer"]),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p GetBallotByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
        "gatheringName": p.GatheringName,
        "userId": p.UserId,
        "numberOfPlayer": p.NumberOfPlayer,
        "keyId": p.KeyId,
    }
}

func (p GetBallotByUserIdRequest) Pointer() *GetBallotByUserIdRequest {
    return &p
}

type VoteRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    BallotBody *string `json:"ballotBody"`
    BallotSignature *string `json:"ballotSignature"`
    GameResults []GameResult `json:"gameResults"`
    KeyId *string `json:"keyId"`
}

func NewVoteRequestFromDict(data map[string]interface{}) VoteRequest {
    return VoteRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        BallotBody: core.CastString(data["ballotBody"]),
        BallotSignature: core.CastString(data["ballotSignature"]),
        GameResults: CastGameResults(core.CastArray(data["gameResults"])),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p VoteRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ballotBody": p.BallotBody,
        "ballotSignature": p.BallotSignature,
        "gameResults": CastGameResultsFromDict(
            p.GameResults,
        ),
        "keyId": p.KeyId,
    }
}

func (p VoteRequest) Pointer() *VoteRequest {
    return &p
}

type VoteMultipleRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    SignedBallots []SignedBallot `json:"signedBallots"`
    GameResults []GameResult `json:"gameResults"`
    KeyId *string `json:"keyId"`
}

func NewVoteMultipleRequestFromDict(data map[string]interface{}) VoteMultipleRequest {
    return VoteMultipleRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        SignedBallots: CastSignedBallots(core.CastArray(data["signedBallots"])),
        GameResults: CastGameResults(core.CastArray(data["gameResults"])),
        KeyId: core.CastString(data["keyId"]),
    }
}

func (p VoteMultipleRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "signedBallots": CastSignedBallotsFromDict(
            p.SignedBallots,
        ),
        "gameResults": CastGameResultsFromDict(
            p.GameResults,
        ),
        "keyId": p.KeyId,
    }
}

func (p VoteMultipleRequest) Pointer() *VoteMultipleRequest {
    return &p
}

type CommitVoteRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    NamespaceName *string `json:"namespaceName"`
    RatingName *string `json:"ratingName"`
    GatheringName *string `json:"gatheringName"`
}

func NewCommitVoteRequestFromDict(data map[string]interface{}) CommitVoteRequest {
    return CommitVoteRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        RatingName: core.CastString(data["ratingName"]),
        GatheringName: core.CastString(data["gatheringName"]),
    }
}

func (p CommitVoteRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "ratingName": p.RatingName,
        "gatheringName": p.GatheringName,
    }
}

func (p CommitVoteRequest) Pointer() *CommitVoteRequest {
    return &p
}