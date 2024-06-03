// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api/miscservice"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/util"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := miscservice.NewClient("MiscService", client.WithHostPorts(fmt.Sprintf("0.0.0.0:%v", config.MiscServerPort)))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &sc_misc_api.ReaderLoginRequest{
			Username: "demo_reader",
			Password: "123456",
		}
		resp, err := client.ReaderLogin(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		klog.Info(util.ToJSON(resp))
		time.Sleep(time.Second)
	}
}
