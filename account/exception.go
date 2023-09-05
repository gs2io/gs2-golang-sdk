package account

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

type PasswordIncorrect struct {
}

func (p PasswordIncorrect) Type() string {
	return "AccountPasswordIncorrect"
}

func (p PasswordIncorrect) Code() string {
	return "account.password.invalid"
}

type BannedInfinity struct {
}

func (p BannedInfinity) Type() string {
	return "AccountBannedInfinity"
}

func (p BannedInfinity) Code() string {
	return "account.banned.infinity"
}
