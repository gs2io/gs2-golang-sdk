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

type DescribeNamespacesResult struct {
    /** ネームスペースのリスト */
	Items         []Namespace	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
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
	Status         *string	`json:"status"`
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

type DescribeProfilesResult struct {
    /** プロフィールのリスト */
	Items         []Profile	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeProfilesResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]Profile, 0)
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

type DescribeProfilesAsyncResult struct {
	result *DescribeProfilesResult
	err    error
}

type GetProfileResult struct {
    /** プロフィール */
	Item         *Profile	`json:"item"`
}

func (p *GetProfileResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetProfileAsyncResult struct {
	result *GetProfileResult
	err    error
}

type GetProfileByUserIdResult struct {
    /** プロフィール */
	Item         *Profile	`json:"item"`
}

func (p *GetProfileByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetProfileByUserIdAsyncResult struct {
	result *GetProfileByUserIdResult
	err    error
}

type UpdateProfileResult struct {
    /** 更新したプロフィール */
	Item         *Profile	`json:"item"`
}

func (p *UpdateProfileResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateProfileAsyncResult struct {
	result *UpdateProfileResult
	err    error
}

type UpdateProfileByUserIdResult struct {
    /** 更新したプロフィール */
	Item         *Profile	`json:"item"`
}

func (p *UpdateProfileByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UpdateProfileByUserIdAsyncResult struct {
	result *UpdateProfileByUserIdResult
	err    error
}

type DeleteProfileByUserIdResult struct {
}

func (p *DeleteProfileByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    return &data
}

type DeleteProfileByUserIdAsyncResult struct {
	result *DeleteProfileByUserIdResult
	err    error
}

type GetPublicProfileResult struct {
    /** 公開プロフィール */
	Item         *PublicProfile	`json:"item"`
}

func (p *GetPublicProfileResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetPublicProfileAsyncResult struct {
	result *GetPublicProfileResult
	err    error
}

type DescribeFollowsResult struct {
    /** フォローしているユーザーのリスト */
	Items         []FollowUser	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeFollowsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FollowUser, 0)
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

type DescribeFollowsAsyncResult struct {
	result *DescribeFollowsResult
	err    error
}

type DescribeFollowsByUserIdResult struct {
    /** フォローしているユーザーのリスト */
	Items         []FollowUser	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeFollowsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FollowUser, 0)
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

type DescribeFollowsByUserIdAsyncResult struct {
	result *DescribeFollowsByUserIdResult
	err    error
}

type GetFollowResult struct {
    /** フォローしているユーザー */
	Item         *FollowUser	`json:"item"`
}

func (p *GetFollowResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFollowAsyncResult struct {
	result *GetFollowResult
	err    error
}

type GetFollowByUserIdResult struct {
    /** フォローしているユーザー */
	Item         *FollowUser	`json:"item"`
}

func (p *GetFollowByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFollowByUserIdAsyncResult struct {
	result *GetFollowByUserIdResult
	err    error
}

type FollowResult struct {
    /** フォローしたユーザ */
	Item         *FollowUser	`json:"item"`
}

func (p *FollowResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type FollowAsyncResult struct {
	result *FollowResult
	err    error
}

type FollowByUserIdResult struct {
    /** フォローしたユーザ */
	Item         *FollowUser	`json:"item"`
}

func (p *FollowByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type FollowByUserIdAsyncResult struct {
	result *FollowByUserIdResult
	err    error
}

type UnfollowResult struct {
    /** アンフォローしたユーザ */
	Item         *FollowUser	`json:"item"`
}

func (p *UnfollowResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnfollowAsyncResult struct {
	result *UnfollowResult
	err    error
}

type UnfollowByUserIdResult struct {
    /** アンフォローしたユーザ */
	Item         *FollowUser	`json:"item"`
}

func (p *UnfollowByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnfollowByUserIdAsyncResult struct {
	result *UnfollowByUserIdResult
	err    error
}

type DescribeFriendsResult struct {
    /** フレンドのユーザーのリスト */
	Items         []FriendUser	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeFriendsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendUser, 0)
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

type DescribeFriendsAsyncResult struct {
	result *DescribeFriendsResult
	err    error
}

type DescribeFriendsByUserIdResult struct {
    /** フレンドのユーザーのリスト */
	Items         []FriendUser	`json:"items"`
    /** リストの続きを取得するためのページトークン */
	NextPageToken         *string	`json:"nextPageToken"`
}

func (p *DescribeFriendsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendUser, 0)
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

type DescribeFriendsByUserIdAsyncResult struct {
	result *DescribeFriendsByUserIdResult
	err    error
}

type GetFriendResult struct {
    /** フレンドのユーザー */
	Item         *FriendUser	`json:"item"`
}

func (p *GetFriendResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFriendAsyncResult struct {
	result *GetFriendResult
	err    error
}

type GetFriendByUserIdResult struct {
    /** フレンドのユーザー */
	Item         *FriendUser	`json:"item"`
}

func (p *GetFriendByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetFriendByUserIdAsyncResult struct {
	result *GetFriendByUserIdResult
	err    error
}

type DeleteFriendResult struct {
    /** フレンドのユーザー */
	Item         *FriendUser	`json:"item"`
}

func (p *DeleteFriendResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteFriendAsyncResult struct {
	result *DeleteFriendResult
	err    error
}

type DeleteFriendByUserIdResult struct {
    /** フレンドのユーザー */
	Item         *FriendUser	`json:"item"`
}

func (p *DeleteFriendByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteFriendByUserIdAsyncResult struct {
	result *DeleteFriendByUserIdResult
	err    error
}

type DescribeSendRequestsResult struct {
    /** フレンドリクエストのリスト */
	Items         []FriendRequest	`json:"items"`
}

func (p *DescribeSendRequestsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendRequest, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeSendRequestsAsyncResult struct {
	result *DescribeSendRequestsResult
	err    error
}

type DescribeSendRequestsByUserIdResult struct {
    /** フレンドリクエストのリスト */
	Items         []FriendRequest	`json:"items"`
}

func (p *DescribeSendRequestsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendRequest, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeSendRequestsByUserIdAsyncResult struct {
	result *DescribeSendRequestsByUserIdResult
	err    error
}

type GetSendRequestResult struct {
    /** フレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *GetSendRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSendRequestAsyncResult struct {
	result *GetSendRequestResult
	err    error
}

type GetSendRequestByUserIdResult struct {
    /** フレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *GetSendRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetSendRequestByUserIdAsyncResult struct {
	result *GetSendRequestByUserIdResult
	err    error
}

type SendRequestResult struct {
    /** 送信したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *SendRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SendRequestAsyncResult struct {
	result *SendRequestResult
	err    error
}

type SendRequestByUserIdResult struct {
    /** 送信したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *SendRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type SendRequestByUserIdAsyncResult struct {
	result *SendRequestByUserIdResult
	err    error
}

type DeleteRequestResult struct {
    /** 削除したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *DeleteRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRequestAsyncResult struct {
	result *DeleteRequestResult
	err    error
}

type DeleteRequestByUserIdResult struct {
    /** 削除したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *DeleteRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type DeleteRequestByUserIdAsyncResult struct {
	result *DeleteRequestByUserIdResult
	err    error
}

type DescribeReceiveRequestsResult struct {
    /** フレンドリクエストのリスト */
	Items         []FriendRequest	`json:"items"`
}

func (p *DescribeReceiveRequestsResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendRequest, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeReceiveRequestsAsyncResult struct {
	result *DescribeReceiveRequestsResult
	err    error
}

type DescribeReceiveRequestsByUserIdResult struct {
    /** フレンドリクエストのリスト */
	Items         []FriendRequest	`json:"items"`
}

func (p *DescribeReceiveRequestsByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]FriendRequest, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeReceiveRequestsByUserIdAsyncResult struct {
	result *DescribeReceiveRequestsByUserIdResult
	err    error
}

type GetReceiveRequestResult struct {
    /** フレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *GetReceiveRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetReceiveRequestAsyncResult struct {
	result *GetReceiveRequestResult
	err    error
}

type GetReceiveRequestByUserIdResult struct {
    /** フレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *GetReceiveRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type GetReceiveRequestByUserIdAsyncResult struct {
	result *GetReceiveRequestByUserIdResult
	err    error
}

type AcceptRequestResult struct {
    /** 承諾したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *AcceptRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AcceptRequestAsyncResult struct {
	result *AcceptRequestResult
	err    error
}

type AcceptRequestByUserIdResult struct {
    /** 承諾したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *AcceptRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type AcceptRequestByUserIdAsyncResult struct {
	result *AcceptRequestByUserIdResult
	err    error
}

type RejectRequestResult struct {
    /** 拒否したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *RejectRequestResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type RejectRequestAsyncResult struct {
	result *RejectRequestResult
	err    error
}

type RejectRequestByUserIdResult struct {
    /** 拒否したフレンドリクエスト */
	Item         *FriendRequest	`json:"item"`
}

func (p *RejectRequestByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type RejectRequestByUserIdAsyncResult struct {
	result *RejectRequestByUserIdResult
	err    error
}

type DescribeBlackListResult struct {
    /** ブラックリストに登録されたユーザIDリスト */
	Items         []string	`json:"items"`
}

func (p *DescribeBlackListResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]string, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeBlackListAsyncResult struct {
	result *DescribeBlackListResult
	err    error
}

type DescribeBlackListByUserIdResult struct {
    /** ブラックリストに登録されたユーザIDリスト */
	Items         []string	`json:"items"`
}

func (p *DescribeBlackListByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Items != nil {
    	items := make([]string, 0)
    	for _, item := range p.Items {
			items = append(items, item)
		}
		data["items"] = items
    }
    return &data
}

type DescribeBlackListByUserIdAsyncResult struct {
	result *DescribeBlackListByUserIdResult
	err    error
}

type RegisterBlackListResult struct {
    /** ブラックリスト */
	Item         *BlackList	`json:"item"`
}

func (p *RegisterBlackListResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type RegisterBlackListAsyncResult struct {
	result *RegisterBlackListResult
	err    error
}

type RegisterBlackListByUserIdResult struct {
    /** ブラックリスト */
	Item         *BlackList	`json:"item"`
}

func (p *RegisterBlackListByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type RegisterBlackListByUserIdAsyncResult struct {
	result *RegisterBlackListByUserIdResult
	err    error
}

type UnregisterBlackListResult struct {
    /** ブラックリスト */
	Item         *BlackList	`json:"item"`
}

func (p *UnregisterBlackListResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnregisterBlackListAsyncResult struct {
	result *UnregisterBlackListResult
	err    error
}

type UnregisterBlackListByUserIdResult struct {
    /** ブラックリスト */
	Item         *BlackList	`json:"item"`
}

func (p *UnregisterBlackListByUserIdResult) ToDict() *map[string]interface{} {
    var data = map[string]interface{}{}
    if p.Item != nil {
        data["item"] = p.Item.ToDict()
    }
    return &data
}

type UnregisterBlackListByUserIdAsyncResult struct {
	result *UnregisterBlackListByUserIdResult
	err    error
}
