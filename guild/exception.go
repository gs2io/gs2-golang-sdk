package guild

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

type MaximumMembersReached struct {
}

func (p MaximumMembersReached) Type() string {
	return "GuildMaximumMembersReached"
}

func (p MaximumMembersReached) Code() string {
	return "guild.members.tooMany"
}

type GuildMasterRequired struct {
}

func (p GuildMasterRequired) Type() string {
	return "GuildGuildMasterRequired"
}

func (p GuildMasterRequired) Code() string {
	return "guild.member.master.require"
}

type GuildJoinRequestLimitReached struct {
}

func (p GuildJoinRequestLimitReached) Type() string {
	return "GuildGuildJoinRequestLimitReached"
}

func (p GuildJoinRequestLimitReached) Code() string {
	return "guild.request.tooMany"
}
