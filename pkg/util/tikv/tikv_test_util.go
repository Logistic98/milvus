// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tikv

import (
	"github.com/tikv/client-go/v2/testutils"
	tilib "github.com/tikv/client-go/v2/tikv"
	"github.com/tikv/client-go/v2/txnkv"
)

func SetupLocalTxn() *txnkv.Client {
	client, cluster, pdClient, err := testutils.NewMockTiKV("", nil)
	if err != nil {
		panic(err)
	}
	testutils.BootstrapWithSingleStore(cluster)
	store, err := tilib.NewTestTiKVStore(client, pdClient, nil, nil, 0)
	if err != nil {
		panic(err)
	}
	return &txnkv.Client{KVStore: store}
}
