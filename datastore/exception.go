package datastore

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

type InvalidStatus struct {
}

func (p InvalidStatus) Type() string {
    return "DatastoreInvalidStatus"
}

func (p InvalidStatus) Code() string {
    return "dataObject.status.invalid"
}

type NotUploaded struct {
}

func (p NotUploaded) Type() string {
    return "DatastoreNotUploaded"
}

func (p NotUploaded) Code() string {
    return "dataObject.file.notUploaded"
}