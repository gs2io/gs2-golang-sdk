package ranking2

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

type RewardAlreadyReceived struct {
}

func (p RewardAlreadyReceived) Type() string {
	return "Ranking2RewardAlreadyReceived"
}

func (p RewardAlreadyReceived) Code() string {
	return "ranking2.rankingReward.alreadyReceived"
}

type SeasonNotEnded struct {
}

func (p SeasonNotEnded) Type() string {
	return "Ranking2SeasonNotEnded"
}

func (p SeasonNotEnded) Code() string {
	return "ranking2.rankingReward.inSchedule"
}

type NotIncludedInCluster struct {
}

func (p NotIncludedInCluster) Type() string {
	return "Ranking2NotIncludedInCluster"
}

func (p NotIncludedInCluster) Code() string {
	return "ranking2.cluster.notInclude"
}

type NoRankingReward struct {
}

func (p NoRankingReward) Type() string {
	return "Ranking2NoRankingReward"
}

func (p NoRankingReward) Code() string {
	return "ranking2.rankingReward.noRewards"
}

type SeasonNotStarted struct {
}

func (p SeasonNotStarted) Type() string {
	return "Ranking2SeasonNotStarted"
}

func (p SeasonNotStarted) Code() string {
	return "ranking2.rankingReward.outOfSchedule"
}
