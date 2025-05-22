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

deny overwrite
*/

package auth

import (
    "encoding/json"

    "github.com/gs2io/gs2-golang-sdk/core"
)

type LoginResult struct {
    Token    *core.AccessToken    `json:"token"`
    UserId   *string              `json:"userId"`
    Expire   *int64               `json:"expire"`
    Metadata *core.ResultMetadata `json:"metadata"`
}

type LoginAsyncResult struct {
    result *LoginResult
    err    error
}

func NewLoginResultFromJson(data string) LoginResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLoginResultFromDict(dict)
}

func NewLoginResultFromDict(data map[string]interface{}) LoginResult {
    return LoginResult{
        Token: func() *string {
            v, ok := data["token"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["token"])
        }(),
        UserId: func() *string {
            v, ok := data["userId"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["userId"])
        }(),
        Expire: func() *int64 {
            v, ok := data["expire"]
            if !ok || v == nil {
                return nil
            }
            return core.CastInt64(data["expire"])
        }(),
        Metadata: func() *core.ResultMetadata {
            if data["metadata"] == nil {
                return nil
            }
            v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
            return &v
        }(),
    }
}

func (p LoginResult) ToDict() map[string]interface{} {
    return map[string]interface{}{
        "token":  p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
        "metadata": func() map[string]interface{} {
            if p.Metadata == nil {
                return nil
            }
            return p.Metadata.ToDict()
        }(),
    }
}

func (p LoginResult) Pointer() *LoginResult {
    return &p
}

type LoginBySignatureResult struct {
    Token    *core.AccessToken    `json:"token"`
    UserId   *string              `json:"userId"`
    Expire   *int64               `json:"expire"`
    Metadata *core.ResultMetadata `json:"metadata"`
}

type LoginBySignatureAsyncResult struct {
    result *LoginBySignatureResult
    err    error
}

func NewLoginBySignatureResultFromJson(data string) LoginBySignatureResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewLoginBySignatureResultFromDict(dict)
}

func NewLoginBySignatureResultFromDict(data map[string]interface{}) LoginBySignatureResult {
    return LoginBySignatureResult{
        Token: func() *string {
            v, ok := data["token"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["token"])
        }(),
        UserId: func() *string {
            v, ok := data["userId"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["userId"])
        }(),
        Expire: func() *int64 {
            v, ok := data["expire"]
            if !ok || v == nil {
                return nil
            }
            return core.CastInt64(data["expire"])
        }(),
        Metadata: func() *core.ResultMetadata {
            if data["metadata"] == nil {
                return nil
            }
            v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
            return &v
        }(),
    }
}

func (p LoginBySignatureResult) ToDict() map[string]interface{} {
    return map[string]interface{}{
        "token":  p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
        "metadata": func() map[string]interface{} {
            if p.Metadata == nil {
                return nil
            }
            return p.Metadata.ToDict()
        }(),
    }
}

func (p LoginBySignatureResult) Pointer() *LoginBySignatureResult {
    return &p
}

type FederationResult struct {
    Token    *core.AccessToken    `json:"token"`
    UserId   *string              `json:"userId"`
    Expire   *int64               `json:"expire"`
    Metadata *core.ResultMetadata `json:"metadata"`
}

type FederationAsyncResult struct {
    result *FederationResult
    err    error
}

func NewFederationResultFromJson(data string) FederationResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewFederationResultFromDict(dict)
}

func NewFederationResultFromDict(data map[string]interface{}) FederationResult {
    return FederationResult{
        Token: func() *string {
            v, ok := data["token"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["token"])
        }(),
        UserId: func() *string {
            v, ok := data["userId"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["userId"])
        }(),
        Expire: func() *int64 {
            v, ok := data["expire"]
            if !ok || v == nil {
                return nil
            }
            return core.CastInt64(data["expire"])
        }(),
        Metadata: func() *core.ResultMetadata {
            if data["metadata"] == nil {
                return nil
            }
            v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
            return &v
        }(),
    }
}

func (p FederationResult) ToDict() map[string]interface{} {
    return map[string]interface{}{
        "token":  p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
        "metadata": func() map[string]interface{} {
            if p.Metadata == nil {
                return nil
            }
            return p.Metadata.ToDict()
        }(),
    }
}

func (p FederationResult) Pointer() *FederationResult {
    return &p
}

type IssueTimeOffsetTokenByUserIdResult struct {
    Token    *string              `json:"token"`
    UserId   *string              `json:"userId"`
    Expire   *int64               `json:"expire"`
    Metadata *core.ResultMetadata `json:"metadata"`
}

type IssueTimeOffsetTokenByUserIdAsyncResult struct {
    result *IssueTimeOffsetTokenByUserIdResult
    err    error
}

func NewIssueTimeOffsetTokenByUserIdResultFromJson(data string) IssueTimeOffsetTokenByUserIdResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewIssueTimeOffsetTokenByUserIdResultFromDict(dict)
}

func NewIssueTimeOffsetTokenByUserIdResultFromDict(data map[string]interface{}) IssueTimeOffsetTokenByUserIdResult {
    return IssueTimeOffsetTokenByUserIdResult{
        Token: func() *string {
            v, ok := data["token"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["token"])
        }(),
        UserId: func() *string {
            v, ok := data["userId"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["userId"])
        }(),
        Expire: func() *int64 {
            v, ok := data["expire"]
            if !ok || v == nil {
                return nil
            }
            return core.CastInt64(data["expire"])
        }(),
        Metadata: func() *core.ResultMetadata {
            if data["metadata"] == nil {
                return nil
            }
            v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
            return &v
        }(),
    }
}

func (p IssueTimeOffsetTokenByUserIdResult) ToDict() map[string]interface{} {
    return map[string]interface{}{
        "token":  p.Token,
        "userId": p.UserId,
        "expire": p.Expire,
        "metadata": func() map[string]interface{} {
            if p.Metadata == nil {
                return nil
            }
            return p.Metadata.ToDict()
        }(),
    }
}

func (p IssueTimeOffsetTokenByUserIdResult) Pointer() *IssueTimeOffsetTokenByUserIdResult {
    return &p
}

type GetServiceVersionResult struct {
    Status   *string              `json:"status"`
    Metadata *core.ResultMetadata `json:"metadata"`
}

type GetServiceVersionAsyncResult struct {
    result *GetServiceVersionResult
    err    error
}

func NewGetServiceVersionResultFromJson(data string) GetServiceVersionResult {
    dict := map[string]interface{}{}
    _ = json.Unmarshal([]byte(data), &dict)
    return NewGetServiceVersionResultFromDict(dict)
}

func NewGetServiceVersionResultFromDict(data map[string]interface{}) GetServiceVersionResult {
    return GetServiceVersionResult{
        Status: func() *string {
            v, ok := data["status"]
            if !ok || v == nil {
                return nil
            }
            return core.CastString(data["status"])
        }(),
        Metadata: func() *core.ResultMetadata {
            if data["metadata"] == nil {
                return nil
            }
            v := core.NewResultMetadataFromDict(core.CastMap(data["metadata"]))
            return &v
        }(),
    }
}

func (p GetServiceVersionResult) ToDict() map[string]interface{} {
    return map[string]interface{}{
        "status": p.Status,
        "metadata": func() map[string]interface{} {
            if p.Metadata == nil {
                return nil
            }
            return p.Metadata.ToDict()
        }(),
    }
}

func (p GetServiceVersionResult) Pointer() *GetServiceVersionResult {
    return &p
}
