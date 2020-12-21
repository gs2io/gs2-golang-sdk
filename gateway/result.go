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

package gateway

import (
	"github.com/gs2io/gs2-golang-sdk/core"
)

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeNamespacesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Namespace, 0)
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
        data["item"] = p.Item.ToDict()
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
        data["status"] = p.Status
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
        data["item"] = p.Item.ToDict()
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
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNamespaceAsyncResult struct {
	result *UpdateNamespaceResult
	err    error
}

type DeleteNamespaceResult struct {
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeWebSocketSessionsResult struct {
    /** Websocketセッションのリスト */
	Items         []WebSocketSession	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeWebSocketSessionsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]WebSocketSession, 0)
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

type DescribeWebSocketSessionsAsyncResult struct {
	result *DescribeWebSocketSessionsResult
	err    error
}

type DescribeWebSocketSessionsByUserIdResult struct {
    /** Websocketセッションのリスト */
	Items         []WebSocketSession	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeWebSocketSessionsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]WebSocketSession, 0)
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

type DescribeWebSocketSessionsByUserIdAsyncResult struct {
	result *DescribeWebSocketSessionsByUserIdResult
	err    error
}

type SetUserIdResult struct {
    /** 更新したWebsocketセッション */
	Item         *WebSocketSession	`json:"item"`
}

func (p *SetUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetUserIdAsyncResult struct {
	result *SetUserIdResult
	err    error
}

type SetUserIdByUserIdResult struct {
    /** 更新したWebsocketセッション */
	Item         *WebSocketSession	`json:"item"`
}

func (p *SetUserIdByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetUserIdByUserIdAsyncResult struct {
	result *SetUserIdByUserIdResult
	err    error
}

type GetWebSocketSessionResult struct {
    /** 取得したWebsocketセッション */
	Item         *WebSocketSession	`json:"item"`
}

func (p *GetWebSocketSessionResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetWebSocketSessionAsyncResult struct {
	result *GetWebSocketSessionResult
	err    error
}

type GetWebSocketSessionByConnectionIdResult struct {
    /** 取得したWebsocketセッション */
	Item         *WebSocketSession	`json:"item"`
}

func (p *GetWebSocketSessionByConnectionIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetWebSocketSessionByConnectionIdAsyncResult struct {
	result *GetWebSocketSessionByConnectionIdResult
	err    error
}

type SendNotificationResult struct {
    /** 通知に使用したプロトコル */
	Protocol         *core.String	`json:"protocol"`
}

func (p *SendNotificationResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Protocol != nil {
        data["protocol"] = p.Protocol
    }
    return &data
}

type SendNotificationAsyncResult struct {
	result *SendNotificationResult
	err    error
}

type SetFirebaseTokenResult struct {
    /** 作成したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *SetFirebaseTokenResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetFirebaseTokenAsyncResult struct {
	result *SetFirebaseTokenResult
	err    error
}

type SetFirebaseTokenByUserIdResult struct {
    /** 作成したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *SetFirebaseTokenByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SetFirebaseTokenByUserIdAsyncResult struct {
	result *SetFirebaseTokenByUserIdResult
	err    error
}

type GetFirebaseTokenResult struct {
    /** 取得したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *GetFirebaseTokenResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFirebaseTokenAsyncResult struct {
	result *GetFirebaseTokenResult
	err    error
}

type GetFirebaseTokenByUserIdResult struct {
    /** 取得したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *GetFirebaseTokenByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFirebaseTokenByUserIdAsyncResult struct {
	result *GetFirebaseTokenByUserIdResult
	err    error
}

type DeleteFirebaseTokenResult struct {
    /** 削除したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *DeleteFirebaseTokenResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteFirebaseTokenAsyncResult struct {
	result *DeleteFirebaseTokenResult
	err    error
}

type DeleteFirebaseTokenByUserIdResult struct {
    /** 削除したFirebaseデバイストークン */
	Item         *FirebaseToken	`json:"item"`
}

func (p *DeleteFirebaseTokenByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteFirebaseTokenByUserIdAsyncResult struct {
	result *DeleteFirebaseTokenByUserIdResult
	err    error
}

type SendMobileNotificationByUserIdResult struct {
}

func (p *SendMobileNotificationByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type SendMobileNotificationByUserIdAsyncResult struct {
	result *SendMobileNotificationByUserIdResult
	err    error
}
