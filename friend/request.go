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

package friend

import (
    "encoding/json"
    "github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeNamespacesRequestFromJson(data string) DescribeNamespacesRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeNamespacesRequestFromDict(dict)
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    Name *string `json:"name"`
    Description *string `json:"description"`
    FollowScript *ScriptSetting `json:"followScript"`
    UnfollowScript *ScriptSetting `json:"unfollowScript"`
    SendRequestScript *ScriptSetting `json:"sendRequestScript"`
    CancelRequestScript *ScriptSetting `json:"cancelRequestScript"`
    AcceptRequestScript *ScriptSetting `json:"acceptRequestScript"`
    RejectRequestScript *ScriptSetting `json:"rejectRequestScript"`
    DeleteFriendScript *ScriptSetting `json:"deleteFriendScript"`
    UpdateProfileScript *ScriptSetting `json:"updateProfileScript"`
    FollowNotification *NotificationSetting `json:"followNotification"`
    ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
    AcceptRequestNotification *NotificationSetting `json:"acceptRequestNotification"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewCreateNamespaceRequestFromJson(data string) CreateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewCreateNamespaceRequestFromDict(dict)
}

func NewCreateNamespaceRequestFromDict(data map[string]interface{}) CreateNamespaceRequest {
    return CreateNamespaceRequest {
        Name: core.CastString(data["name"]),
        Description: core.CastString(data["description"]),
        FollowScript: NewScriptSettingFromDict(core.CastMap(data["followScript"])).Pointer(),
        UnfollowScript: NewScriptSettingFromDict(core.CastMap(data["unfollowScript"])).Pointer(),
        SendRequestScript: NewScriptSettingFromDict(core.CastMap(data["sendRequestScript"])).Pointer(),
        CancelRequestScript: NewScriptSettingFromDict(core.CastMap(data["cancelRequestScript"])).Pointer(),
        AcceptRequestScript: NewScriptSettingFromDict(core.CastMap(data["acceptRequestScript"])).Pointer(),
        RejectRequestScript: NewScriptSettingFromDict(core.CastMap(data["rejectRequestScript"])).Pointer(),
        DeleteFriendScript: NewScriptSettingFromDict(core.CastMap(data["deleteFriendScript"])).Pointer(),
        UpdateProfileScript: NewScriptSettingFromDict(core.CastMap(data["updateProfileScript"])).Pointer(),
        FollowNotification: NewNotificationSettingFromDict(core.CastMap(data["followNotification"])).Pointer(),
        ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
        AcceptRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["acceptRequestNotification"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p CreateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "name": p.Name,
        "description": p.Description,
        "followScript": p.FollowScript.ToDict(),
        "unfollowScript": p.UnfollowScript.ToDict(),
        "sendRequestScript": p.SendRequestScript.ToDict(),
        "cancelRequestScript": p.CancelRequestScript.ToDict(),
        "acceptRequestScript": p.AcceptRequestScript.ToDict(),
        "rejectRequestScript": p.RejectRequestScript.ToDict(),
        "deleteFriendScript": p.DeleteFriendScript.ToDict(),
        "updateProfileScript": p.UpdateProfileScript.ToDict(),
        "followNotification": p.FollowNotification.ToDict(),
        "receiveRequestNotification": p.ReceiveRequestNotification.ToDict(),
        "acceptRequestNotification": p.AcceptRequestNotification.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p CreateNamespaceRequest) Pointer() *CreateNamespaceRequest {
    return &p
}

type GetNamespaceStatusRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceStatusRequestFromJson(data string) GetNamespaceStatusRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceStatusRequestFromDict(dict)
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewGetNamespaceRequestFromJson(data string) GetNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetNamespaceRequestFromDict(dict)
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
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    Description *string `json:"description"`
    FollowScript *ScriptSetting `json:"followScript"`
    UnfollowScript *ScriptSetting `json:"unfollowScript"`
    SendRequestScript *ScriptSetting `json:"sendRequestScript"`
    CancelRequestScript *ScriptSetting `json:"cancelRequestScript"`
    AcceptRequestScript *ScriptSetting `json:"acceptRequestScript"`
    RejectRequestScript *ScriptSetting `json:"rejectRequestScript"`
    DeleteFriendScript *ScriptSetting `json:"deleteFriendScript"`
    UpdateProfileScript *ScriptSetting `json:"updateProfileScript"`
    FollowNotification *NotificationSetting `json:"followNotification"`
    ReceiveRequestNotification *NotificationSetting `json:"receiveRequestNotification"`
    AcceptRequestNotification *NotificationSetting `json:"acceptRequestNotification"`
    LogSetting *LogSetting `json:"logSetting"`
}

func NewUpdateNamespaceRequestFromJson(data string) UpdateNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateNamespaceRequestFromDict(dict)
}

func NewUpdateNamespaceRequestFromDict(data map[string]interface{}) UpdateNamespaceRequest {
    return UpdateNamespaceRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        Description: core.CastString(data["description"]),
        FollowScript: NewScriptSettingFromDict(core.CastMap(data["followScript"])).Pointer(),
        UnfollowScript: NewScriptSettingFromDict(core.CastMap(data["unfollowScript"])).Pointer(),
        SendRequestScript: NewScriptSettingFromDict(core.CastMap(data["sendRequestScript"])).Pointer(),
        CancelRequestScript: NewScriptSettingFromDict(core.CastMap(data["cancelRequestScript"])).Pointer(),
        AcceptRequestScript: NewScriptSettingFromDict(core.CastMap(data["acceptRequestScript"])).Pointer(),
        RejectRequestScript: NewScriptSettingFromDict(core.CastMap(data["rejectRequestScript"])).Pointer(),
        DeleteFriendScript: NewScriptSettingFromDict(core.CastMap(data["deleteFriendScript"])).Pointer(),
        UpdateProfileScript: NewScriptSettingFromDict(core.CastMap(data["updateProfileScript"])).Pointer(),
        FollowNotification: NewNotificationSettingFromDict(core.CastMap(data["followNotification"])).Pointer(),
        ReceiveRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["receiveRequestNotification"])).Pointer(),
        AcceptRequestNotification: NewNotificationSettingFromDict(core.CastMap(data["acceptRequestNotification"])).Pointer(),
        LogSetting: NewLogSettingFromDict(core.CastMap(data["logSetting"])).Pointer(),
    }
}

func (p UpdateNamespaceRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "description": p.Description,
        "followScript": p.FollowScript.ToDict(),
        "unfollowScript": p.UnfollowScript.ToDict(),
        "sendRequestScript": p.SendRequestScript.ToDict(),
        "cancelRequestScript": p.CancelRequestScript.ToDict(),
        "acceptRequestScript": p.AcceptRequestScript.ToDict(),
        "rejectRequestScript": p.RejectRequestScript.ToDict(),
        "deleteFriendScript": p.DeleteFriendScript.ToDict(),
        "updateProfileScript": p.UpdateProfileScript.ToDict(),
        "followNotification": p.FollowNotification.ToDict(),
        "receiveRequestNotification": p.ReceiveRequestNotification.ToDict(),
        "acceptRequestNotification": p.AcceptRequestNotification.ToDict(),
        "logSetting": p.LogSetting.ToDict(),
    }
}

func (p UpdateNamespaceRequest) Pointer() *UpdateNamespaceRequest {
    return &p
}

type DeleteNamespaceRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
}

func NewDeleteNamespaceRequestFromJson(data string) DeleteNamespaceRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteNamespaceRequestFromDict(dict)
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

type GetProfileRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
}

func NewGetProfileRequestFromJson(data string) GetProfileRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetProfileRequestFromDict(dict)
}

func NewGetProfileRequestFromDict(data map[string]interface{}) GetProfileRequest {
    return GetProfileRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
    }
}

func (p GetProfileRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
    }
}

func (p GetProfileRequest) Pointer() *GetProfileRequest {
    return &p
}

type GetProfileByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewGetProfileByUserIdRequestFromJson(data string) GetProfileByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetProfileByUserIdRequestFromDict(dict)
}

func NewGetProfileByUserIdRequestFromDict(data map[string]interface{}) GetProfileByUserIdRequest {
    return GetProfileByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetProfileByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p GetProfileByUserIdRequest) Pointer() *GetProfileByUserIdRequest {
    return &p
}

type UpdateProfileRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PublicProfile *string `json:"publicProfile"`
    FollowerProfile *string `json:"followerProfile"`
    FriendProfile *string `json:"friendProfile"`
}

func NewUpdateProfileRequestFromJson(data string) UpdateProfileRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateProfileRequestFromDict(dict)
}

func NewUpdateProfileRequestFromDict(data map[string]interface{}) UpdateProfileRequest {
    return UpdateProfileRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PublicProfile: core.CastString(data["publicProfile"]),
        FollowerProfile: core.CastString(data["followerProfile"]),
        FriendProfile: core.CastString(data["friendProfile"]),
    }
}

func (p UpdateProfileRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "publicProfile": p.PublicProfile,
        "followerProfile": p.FollowerProfile,
        "friendProfile": p.FriendProfile,
    }
}

func (p UpdateProfileRequest) Pointer() *UpdateProfileRequest {
    return &p
}

type UpdateProfileByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PublicProfile *string `json:"publicProfile"`
    FollowerProfile *string `json:"followerProfile"`
    FriendProfile *string `json:"friendProfile"`
}

func NewUpdateProfileByUserIdRequestFromJson(data string) UpdateProfileByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUpdateProfileByUserIdRequestFromDict(dict)
}

func NewUpdateProfileByUserIdRequestFromDict(data map[string]interface{}) UpdateProfileByUserIdRequest {
    return UpdateProfileByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PublicProfile: core.CastString(data["publicProfile"]),
        FollowerProfile: core.CastString(data["followerProfile"]),
        FriendProfile: core.CastString(data["friendProfile"]),
    }
}

func (p UpdateProfileByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "publicProfile": p.PublicProfile,
        "followerProfile": p.FollowerProfile,
        "friendProfile": p.FriendProfile,
    }
}

func (p UpdateProfileByUserIdRequest) Pointer() *UpdateProfileByUserIdRequest {
    return &p
}

type DeleteProfileByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewDeleteProfileByUserIdRequestFromJson(data string) DeleteProfileByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteProfileByUserIdRequestFromDict(dict)
}

func NewDeleteProfileByUserIdRequestFromDict(data map[string]interface{}) DeleteProfileByUserIdRequest {
    return DeleteProfileByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p DeleteProfileByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p DeleteProfileByUserIdRequest) Pointer() *DeleteProfileByUserIdRequest {
    return &p
}

type DescribeFriendsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    WithProfile *bool `json:"withProfile"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFriendsRequestFromJson(data string) DescribeFriendsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFriendsRequestFromDict(dict)
}

func NewDescribeFriendsRequestFromDict(data map[string]interface{}) DescribeFriendsRequest {
    return DescribeFriendsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        WithProfile: core.CastBool(data["withProfile"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFriendsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "withProfile": p.WithProfile,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFriendsRequest) Pointer() *DescribeFriendsRequest {
    return &p
}

type DescribeFriendsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    WithProfile *bool `json:"withProfile"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFriendsByUserIdRequestFromJson(data string) DescribeFriendsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFriendsByUserIdRequestFromDict(dict)
}

func NewDescribeFriendsByUserIdRequestFromDict(data map[string]interface{}) DescribeFriendsByUserIdRequest {
    return DescribeFriendsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        WithProfile: core.CastBool(data["withProfile"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFriendsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "withProfile": p.WithProfile,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFriendsByUserIdRequest) Pointer() *DescribeFriendsByUserIdRequest {
    return &p
}

type DescribeBlackListRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeBlackListRequestFromJson(data string) DescribeBlackListRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBlackListRequestFromDict(dict)
}

func NewDescribeBlackListRequestFromDict(data map[string]interface{}) DescribeBlackListRequest {
    return DescribeBlackListRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeBlackListRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeBlackListRequest) Pointer() *DescribeBlackListRequest {
    return &p
}

type DescribeBlackListByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeBlackListByUserIdRequestFromJson(data string) DescribeBlackListByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeBlackListByUserIdRequestFromDict(dict)
}

func NewDescribeBlackListByUserIdRequestFromDict(data map[string]interface{}) DescribeBlackListByUserIdRequest {
    return DescribeBlackListByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeBlackListByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeBlackListByUserIdRequest) Pointer() *DescribeBlackListByUserIdRequest {
    return &p
}

type RegisterBlackListRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewRegisterBlackListRequestFromJson(data string) RegisterBlackListRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRegisterBlackListRequestFromDict(dict)
}

func NewRegisterBlackListRequestFromDict(data map[string]interface{}) RegisterBlackListRequest {
    return RegisterBlackListRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p RegisterBlackListRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p RegisterBlackListRequest) Pointer() *RegisterBlackListRequest {
    return &p
}

type RegisterBlackListByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewRegisterBlackListByUserIdRequestFromJson(data string) RegisterBlackListByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRegisterBlackListByUserIdRequestFromDict(dict)
}

func NewRegisterBlackListByUserIdRequestFromDict(data map[string]interface{}) RegisterBlackListByUserIdRequest {
    return RegisterBlackListByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p RegisterBlackListByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p RegisterBlackListByUserIdRequest) Pointer() *RegisterBlackListByUserIdRequest {
    return &p
}

type UnregisterBlackListRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnregisterBlackListRequestFromJson(data string) UnregisterBlackListRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnregisterBlackListRequestFromDict(dict)
}

func NewUnregisterBlackListRequestFromDict(data map[string]interface{}) UnregisterBlackListRequest {
    return UnregisterBlackListRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnregisterBlackListRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnregisterBlackListRequest) Pointer() *UnregisterBlackListRequest {
    return &p
}

type UnregisterBlackListByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnregisterBlackListByUserIdRequestFromJson(data string) UnregisterBlackListByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnregisterBlackListByUserIdRequestFromDict(dict)
}

func NewUnregisterBlackListByUserIdRequestFromDict(data map[string]interface{}) UnregisterBlackListByUserIdRequest {
    return UnregisterBlackListByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnregisterBlackListByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnregisterBlackListByUserIdRequest) Pointer() *UnregisterBlackListByUserIdRequest {
    return &p
}

type DescribeFollowsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    WithProfile *bool `json:"withProfile"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFollowsRequestFromJson(data string) DescribeFollowsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFollowsRequestFromDict(dict)
}

func NewDescribeFollowsRequestFromDict(data map[string]interface{}) DescribeFollowsRequest {
    return DescribeFollowsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        WithProfile: core.CastBool(data["withProfile"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFollowsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "withProfile": p.WithProfile,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFollowsRequest) Pointer() *DescribeFollowsRequest {
    return &p
}

type DescribeFollowsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    WithProfile *bool `json:"withProfile"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeFollowsByUserIdRequestFromJson(data string) DescribeFollowsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeFollowsByUserIdRequestFromDict(dict)
}

func NewDescribeFollowsByUserIdRequestFromDict(data map[string]interface{}) DescribeFollowsByUserIdRequest {
    return DescribeFollowsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        WithProfile: core.CastBool(data["withProfile"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeFollowsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "withProfile": p.WithProfile,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeFollowsByUserIdRequest) Pointer() *DescribeFollowsByUserIdRequest {
    return &p
}

type GetFollowRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
    WithProfile *bool `json:"withProfile"`
}

func NewGetFollowRequestFromJson(data string) GetFollowRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFollowRequestFromDict(dict)
}

func NewGetFollowRequestFromDict(data map[string]interface{}) GetFollowRequest {
    return GetFollowRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
        WithProfile: core.CastBool(data["withProfile"]),
    }
}

func (p GetFollowRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
        "withProfile": p.WithProfile,
    }
}

func (p GetFollowRequest) Pointer() *GetFollowRequest {
    return &p
}

type GetFollowByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
    WithProfile *bool `json:"withProfile"`
}

func NewGetFollowByUserIdRequestFromJson(data string) GetFollowByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFollowByUserIdRequestFromDict(dict)
}

func NewGetFollowByUserIdRequestFromDict(data map[string]interface{}) GetFollowByUserIdRequest {
    return GetFollowByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
        WithProfile: core.CastBool(data["withProfile"]),
    }
}

func (p GetFollowByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
        "withProfile": p.WithProfile,
    }
}

func (p GetFollowByUserIdRequest) Pointer() *GetFollowByUserIdRequest {
    return &p
}

type FollowRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewFollowRequestFromJson(data string) FollowRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewFollowRequestFromDict(dict)
}

func NewFollowRequestFromDict(data map[string]interface{}) FollowRequest {
    return FollowRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p FollowRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p FollowRequest) Pointer() *FollowRequest {
    return &p
}

type FollowByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewFollowByUserIdRequestFromJson(data string) FollowByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewFollowByUserIdRequestFromDict(dict)
}

func NewFollowByUserIdRequestFromDict(data map[string]interface{}) FollowByUserIdRequest {
    return FollowByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p FollowByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p FollowByUserIdRequest) Pointer() *FollowByUserIdRequest {
    return &p
}

type UnfollowRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnfollowRequestFromJson(data string) UnfollowRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnfollowRequestFromDict(dict)
}

func NewUnfollowRequestFromDict(data map[string]interface{}) UnfollowRequest {
    return UnfollowRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnfollowRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnfollowRequest) Pointer() *UnfollowRequest {
    return &p
}

type UnfollowByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewUnfollowByUserIdRequestFromJson(data string) UnfollowByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewUnfollowByUserIdRequestFromDict(dict)
}

func NewUnfollowByUserIdRequestFromDict(data map[string]interface{}) UnfollowByUserIdRequest {
    return UnfollowByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p UnfollowByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p UnfollowByUserIdRequest) Pointer() *UnfollowByUserIdRequest {
    return &p
}

type GetFriendRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
    WithProfile *bool `json:"withProfile"`
}

func NewGetFriendRequestFromJson(data string) GetFriendRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFriendRequestFromDict(dict)
}

func NewGetFriendRequestFromDict(data map[string]interface{}) GetFriendRequest {
    return GetFriendRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
        WithProfile: core.CastBool(data["withProfile"]),
    }
}

func (p GetFriendRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
        "withProfile": p.WithProfile,
    }
}

func (p GetFriendRequest) Pointer() *GetFriendRequest {
    return &p
}

type GetFriendByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
    WithProfile *bool `json:"withProfile"`
}

func NewGetFriendByUserIdRequestFromJson(data string) GetFriendByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetFriendByUserIdRequestFromDict(dict)
}

func NewGetFriendByUserIdRequestFromDict(data map[string]interface{}) GetFriendByUserIdRequest {
    return GetFriendByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
        WithProfile: core.CastBool(data["withProfile"]),
    }
}

func (p GetFriendByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
        "withProfile": p.WithProfile,
    }
}

func (p GetFriendByUserIdRequest) Pointer() *GetFriendByUserIdRequest {
    return &p
}

type DeleteFriendRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewDeleteFriendRequestFromJson(data string) DeleteFriendRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteFriendRequestFromDict(dict)
}

func NewDeleteFriendRequestFromDict(data map[string]interface{}) DeleteFriendRequest {
    return DeleteFriendRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p DeleteFriendRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p DeleteFriendRequest) Pointer() *DeleteFriendRequest {
    return &p
}

type DeleteFriendByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewDeleteFriendByUserIdRequestFromJson(data string) DeleteFriendByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteFriendByUserIdRequestFromDict(dict)
}

func NewDeleteFriendByUserIdRequestFromDict(data map[string]interface{}) DeleteFriendByUserIdRequest {
    return DeleteFriendByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p DeleteFriendByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p DeleteFriendByUserIdRequest) Pointer() *DeleteFriendByUserIdRequest {
    return &p
}

type DescribeSendRequestsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeSendRequestsRequestFromJson(data string) DescribeSendRequestsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSendRequestsRequestFromDict(dict)
}

func NewDescribeSendRequestsRequestFromDict(data map[string]interface{}) DescribeSendRequestsRequest {
    return DescribeSendRequestsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeSendRequestsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeSendRequestsRequest) Pointer() *DescribeSendRequestsRequest {
    return &p
}

type DescribeSendRequestsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeSendRequestsByUserIdRequestFromJson(data string) DescribeSendRequestsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeSendRequestsByUserIdRequestFromDict(dict)
}

func NewDescribeSendRequestsByUserIdRequestFromDict(data map[string]interface{}) DescribeSendRequestsByUserIdRequest {
    return DescribeSendRequestsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeSendRequestsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeSendRequestsByUserIdRequest) Pointer() *DescribeSendRequestsByUserIdRequest {
    return &p
}

type GetSendRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewGetSendRequestRequestFromJson(data string) GetSendRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSendRequestRequestFromDict(dict)
}

func NewGetSendRequestRequestFromDict(data map[string]interface{}) GetSendRequestRequest {
    return GetSendRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p GetSendRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p GetSendRequestRequest) Pointer() *GetSendRequestRequest {
    return &p
}

type GetSendRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewGetSendRequestByUserIdRequestFromJson(data string) GetSendRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetSendRequestByUserIdRequestFromDict(dict)
}

func NewGetSendRequestByUserIdRequestFromDict(data map[string]interface{}) GetSendRequestByUserIdRequest {
    return GetSendRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p GetSendRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p GetSendRequestByUserIdRequest) Pointer() *GetSendRequestByUserIdRequest {
    return &p
}

type SendRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewSendRequestRequestFromJson(data string) SendRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSendRequestRequestFromDict(dict)
}

func NewSendRequestRequestFromDict(data map[string]interface{}) SendRequestRequest {
    return SendRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p SendRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p SendRequestRequest) Pointer() *SendRequestRequest {
    return &p
}

type SendRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewSendRequestByUserIdRequestFromJson(data string) SendRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewSendRequestByUserIdRequestFromDict(dict)
}

func NewSendRequestByUserIdRequestFromDict(data map[string]interface{}) SendRequestByUserIdRequest {
    return SendRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p SendRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p SendRequestByUserIdRequest) Pointer() *SendRequestByUserIdRequest {
    return &p
}

type DeleteRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    TargetUserId *string `json:"targetUserId"`
}

func NewDeleteRequestRequestFromJson(data string) DeleteRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteRequestRequestFromDict(dict)
}

func NewDeleteRequestRequestFromDict(data map[string]interface{}) DeleteRequestRequest {
    return DeleteRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p DeleteRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "targetUserId": p.TargetUserId,
    }
}

func (p DeleteRequestRequest) Pointer() *DeleteRequestRequest {
    return &p
}

type DeleteRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    TargetUserId *string `json:"targetUserId"`
}

func NewDeleteRequestByUserIdRequestFromJson(data string) DeleteRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDeleteRequestByUserIdRequestFromDict(dict)
}

func NewDeleteRequestByUserIdRequestFromDict(data map[string]interface{}) DeleteRequestByUserIdRequest {
    return DeleteRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        TargetUserId: core.CastString(data["targetUserId"]),
    }
}

func (p DeleteRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "targetUserId": p.TargetUserId,
    }
}

func (p DeleteRequestByUserIdRequest) Pointer() *DeleteRequestByUserIdRequest {
    return &p
}

type DescribeReceiveRequestsRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiveRequestsRequestFromJson(data string) DescribeReceiveRequestsRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReceiveRequestsRequestFromDict(dict)
}

func NewDescribeReceiveRequestsRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsRequest {
    return DescribeReceiveRequestsRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiveRequestsRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiveRequestsRequest) Pointer() *DescribeReceiveRequestsRequest {
    return &p
}

type DescribeReceiveRequestsByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    PageToken *string `json:"pageToken"`
    Limit *int32 `json:"limit"`
}

func NewDescribeReceiveRequestsByUserIdRequestFromJson(data string) DescribeReceiveRequestsByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewDescribeReceiveRequestsByUserIdRequestFromDict(dict)
}

func NewDescribeReceiveRequestsByUserIdRequestFromDict(data map[string]interface{}) DescribeReceiveRequestsByUserIdRequest {
    return DescribeReceiveRequestsByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        PageToken: core.CastString(data["pageToken"]),
        Limit: core.CastInt32(data["limit"]),
    }
}

func (p DescribeReceiveRequestsByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "pageToken": p.PageToken,
        "limit": p.Limit,
    }
}

func (p DescribeReceiveRequestsByUserIdRequest) Pointer() *DescribeReceiveRequestsByUserIdRequest {
    return &p
}

type GetReceiveRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    FromUserId *string `json:"fromUserId"`
}

func NewGetReceiveRequestRequestFromJson(data string) GetReceiveRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReceiveRequestRequestFromDict(dict)
}

func NewGetReceiveRequestRequestFromDict(data map[string]interface{}) GetReceiveRequestRequest {
    return GetReceiveRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p GetReceiveRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "fromUserId": p.FromUserId,
    }
}

func (p GetReceiveRequestRequest) Pointer() *GetReceiveRequestRequest {
    return &p
}

type GetReceiveRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    FromUserId *string `json:"fromUserId"`
}

func NewGetReceiveRequestByUserIdRequestFromJson(data string) GetReceiveRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetReceiveRequestByUserIdRequestFromDict(dict)
}

func NewGetReceiveRequestByUserIdRequestFromDict(data map[string]interface{}) GetReceiveRequestByUserIdRequest {
    return GetReceiveRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p GetReceiveRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "fromUserId": p.FromUserId,
    }
}

func (p GetReceiveRequestByUserIdRequest) Pointer() *GetReceiveRequestByUserIdRequest {
    return &p
}

type AcceptRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    FromUserId *string `json:"fromUserId"`
}

func NewAcceptRequestRequestFromJson(data string) AcceptRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcceptRequestRequestFromDict(dict)
}

func NewAcceptRequestRequestFromDict(data map[string]interface{}) AcceptRequestRequest {
    return AcceptRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p AcceptRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "fromUserId": p.FromUserId,
    }
}

func (p AcceptRequestRequest) Pointer() *AcceptRequestRequest {
    return &p
}

type AcceptRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    FromUserId *string `json:"fromUserId"`
}

func NewAcceptRequestByUserIdRequestFromJson(data string) AcceptRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewAcceptRequestByUserIdRequestFromDict(dict)
}

func NewAcceptRequestByUserIdRequestFromDict(data map[string]interface{}) AcceptRequestByUserIdRequest {
    return AcceptRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p AcceptRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "fromUserId": p.FromUserId,
    }
}

func (p AcceptRequestByUserIdRequest) Pointer() *AcceptRequestByUserIdRequest {
    return &p
}

type RejectRequestRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    AccessToken *string `json:"accessToken"`
    FromUserId *string `json:"fromUserId"`
}

func NewRejectRequestRequestFromJson(data string) RejectRequestRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRejectRequestRequestFromDict(dict)
}

func NewRejectRequestRequestFromDict(data map[string]interface{}) RejectRequestRequest {
    return RejectRequestRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        AccessToken: core.CastString(data["accessToken"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p RejectRequestRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "accessToken": p.AccessToken,
        "fromUserId": p.FromUserId,
    }
}

func (p RejectRequestRequest) Pointer() *RejectRequestRequest {
    return &p
}

type RejectRequestByUserIdRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
    FromUserId *string `json:"fromUserId"`
}

func NewRejectRequestByUserIdRequestFromJson(data string) RejectRequestByUserIdRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewRejectRequestByUserIdRequestFromDict(dict)
}

func NewRejectRequestByUserIdRequestFromDict(data map[string]interface{}) RejectRequestByUserIdRequest {
    return RejectRequestByUserIdRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
        FromUserId: core.CastString(data["fromUserId"]),
    }
}

func (p RejectRequestByUserIdRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
        "fromUserId": p.FromUserId,
    }
}

func (p RejectRequestByUserIdRequest) Pointer() *RejectRequestByUserIdRequest {
    return &p
}

type GetPublicProfileRequest struct {
    RequestId *string `json:"requestId"`
    ContextStack *string `json:"contextStack"`
    DuplicationAvoider *string `json:"duplicationAvoider"`
    NamespaceName *string `json:"namespaceName"`
    UserId *string `json:"userId"`
}

func NewGetPublicProfileRequestFromJson(data string) GetPublicProfileRequest {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetPublicProfileRequestFromDict(dict)
}

func NewGetPublicProfileRequestFromDict(data map[string]interface{}) GetPublicProfileRequest {
    return GetPublicProfileRequest {
        NamespaceName: core.CastString(data["namespaceName"]),
        UserId: core.CastString(data["userId"]),
    }
}

func (p GetPublicProfileRequest) ToDict() map[string]interface{} {
    return map[string]interface{} {
        "namespaceName": p.NamespaceName,
        "userId": p.UserId,
    }
}

func (p GetPublicProfileRequest) Pointer() *GetPublicProfileRequest {
    return &p
}