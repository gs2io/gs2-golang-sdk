package friend

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

type SendRequestCapacityFull struct {
}

func (p SendRequestCapacityFull) Type() string {
	return "FriendSendRequestCapacityFull"
}

func (p SendRequestCapacityFull) Code() string {
	return "friend.sendFriendRequest.capacity.full"
}

type DuplicateFriendRequest struct {
}

func (p DuplicateFriendRequest) Type() string {
	return "FriendDuplicateFriendRequest"
}

func (p DuplicateFriendRequest) Code() string {
	return "friend.sendFriendRequest.targetUserId.duplicate"
}

type SendRequestToSelf struct {
}

func (p SendRequestToSelf) Type() string {
	return "FriendSendRequestToSelf"
}

func (p SendRequestToSelf) Code() string {
	return "friend.sendFriendRequest.targetUserId.self"
}

type AlreadyFriend struct {
}

func (p AlreadyFriend) Type() string {
	return "FriendAlreadyFriend"
}

func (p AlreadyFriend) Code() string {
	return "friend.friend.targetUserId.duplicate"
}

type AlreadyFollowing struct {
}

func (p AlreadyFollowing) Type() string {
	return "FriendAlreadyFollowing"
}

func (p AlreadyFollowing) Code() string {
	return "friend.followUser.targetUserId.duplicate"
}

type AlreadyInBlackList struct {
}

func (p AlreadyInBlackList) Type() string {
	return "FriendAlreadyInBlackList"
}

func (p AlreadyInBlackList) Code() string {
	return "friend.blackList.targetUserId.duplicate"
}
