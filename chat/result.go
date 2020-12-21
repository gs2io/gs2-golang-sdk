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

package chat

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
    	items := make([]*Namespace, 0)
    	for _, item := range *p.Items {
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
    /** 削除したネームスペース */
	Item         *Namespace	`json:"item"`
}

func (p *DeleteNamespaceResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteNamespaceAsyncResult struct {
	result *DeleteNamespaceResult
	err    error
}

type DescribeRoomsResult struct {
    /** ルームのリスト */
	Items         *[]*Room	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeRoomsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Room, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeRoomsAsyncResult struct {
	result *DescribeRoomsResult
	err    error
}

type CreateRoomResult struct {
    /** 作成したルーム */
	Item         *Room	`json:"item"`
}

func (p *CreateRoomResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRoomAsyncResult struct {
	result *CreateRoomResult
	err    error
}

type CreateRoomFromBackendResult struct {
    /** 作成したルーム */
	Item         *Room	`json:"item"`
}

func (p *CreateRoomFromBackendResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type CreateRoomFromBackendAsyncResult struct {
	result *CreateRoomFromBackendResult
	err    error
}

type GetRoomResult struct {
    /** ルーム */
	Item         *Room	`json:"item"`
}

func (p *GetRoomResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetRoomAsyncResult struct {
	result *GetRoomResult
	err    error
}

type UpdateRoomResult struct {
    /** 更新したルーム */
	Item         *Room	`json:"item"`
}

func (p *UpdateRoomResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateRoomAsyncResult struct {
	result *UpdateRoomResult
	err    error
}

type DeleteRoomResult struct {
    /** 削除したルーム */
	Item         *Room	`json:"item"`
}

func (p *DeleteRoomResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRoomAsyncResult struct {
	result *DeleteRoomResult
	err    error
}

type DeleteRoomFromBackendResult struct {
    /** 削除したルーム */
	Item         *Room	`json:"item"`
}

func (p *DeleteRoomFromBackendResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRoomFromBackendAsyncResult struct {
	result *DeleteRoomFromBackendResult
	err    error
}

type DescribeMessagesResult struct {
    /** メッセージのリスト */
	Items         *[]*Message	`json:"items"`
}

func (p *DescribeMessagesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Message, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeMessagesAsyncResult struct {
	result *DescribeMessagesResult
	err    error
}

type PostResult struct {
    /** 投稿したメッセージ */
	Item         *Message	`json:"item"`
}

func (p *PostResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type PostAsyncResult struct {
	result *PostResult
	err    error
}

type PostByUserIdResult struct {
    /** 投稿したメッセージ */
	Item         *Message	`json:"item"`
}

func (p *PostByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type PostByUserIdAsyncResult struct {
	result *PostByUserIdResult
	err    error
}

type GetMessageResult struct {
    /** メッセージ */
	Item         *Message	`json:"item"`
}

func (p *GetMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetMessageAsyncResult struct {
	result *GetMessageResult
	err    error
}

type DeleteMessageResult struct {
    /** 削除したメッセージ */
	Item         *Message	`json:"item"`
}

func (p *DeleteMessageResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteMessageAsyncResult struct {
	result *DeleteMessageResult
	err    error
}

type DescribeSubscribesResult struct {
    /** 購読のリスト */
	Items         *[]*Subscribe	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeSubscribesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Subscribe, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeSubscribesAsyncResult struct {
	result *DescribeSubscribesResult
	err    error
}

type DescribeSubscribesByUserIdResult struct {
    /** 購読のリスト */
	Items         *[]*Subscribe	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeSubscribesByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Subscribe, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeSubscribesByUserIdAsyncResult struct {
	result *DescribeSubscribesByUserIdResult
	err    error
}

type DescribeSubscribesByRoomNameResult struct {
    /** 購読のリスト */
	Items         *[]*Subscribe	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *core.String	`json:"nextPageToken"`
}

func (p *DescribeSubscribesByRoomNameResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]*Subscribe, 0)
    	for _, item := range *p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    if p.NextPageToken != nil {
        data["nextPageToken"] = p.NextPageToken
    }
    return &data
}

type DescribeSubscribesByRoomNameAsyncResult struct {
	result *DescribeSubscribesByRoomNameResult
	err    error
}

type SubscribeResult struct {
    /** 購読した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *SubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SubscribeAsyncResult struct {
	result *SubscribeResult
	err    error
}

type SubscribeByUserIdResult struct {
    /** 購読した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *SubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SubscribeByUserIdAsyncResult struct {
	result *SubscribeByUserIdResult
	err    error
}

type GetSubscribeResult struct {
    /** 購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *GetSubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSubscribeAsyncResult struct {
	result *GetSubscribeResult
	err    error
}

type GetSubscribeByUserIdResult struct {
    /** 購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *GetSubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSubscribeByUserIdAsyncResult struct {
	result *GetSubscribeByUserIdResult
	err    error
}

type UpdateNotificationTypeResult struct {
    /** 更新した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *UpdateNotificationTypeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNotificationTypeAsyncResult struct {
	result *UpdateNotificationTypeResult
	err    error
}

type UpdateNotificationTypeByUserIdResult struct {
    /** 更新した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *UpdateNotificationTypeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateNotificationTypeByUserIdAsyncResult struct {
	result *UpdateNotificationTypeByUserIdResult
	err    error
}

type UnsubscribeResult struct {
    /** 解除した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *UnsubscribeResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnsubscribeAsyncResult struct {
	result *UnsubscribeResult
	err    error
}

type UnsubscribeByUserIdResult struct {
    /** 解除した購読 */
	Item         *Subscribe	`json:"item"`
}

func (p *UnsubscribeByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnsubscribeByUserIdAsyncResult struct {
	result *UnsubscribeByUserIdResult
	err    error
}
