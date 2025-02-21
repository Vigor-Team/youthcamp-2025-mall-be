// Copyright 2024 CloudWeGo Authors
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

package service

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/checkout/infra/rpc"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/checkout"
	rpcpayment "github.com/Vigor-Team/youthcamp-2025-mall-be/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	rpc.InitClient()
	ctx := context.Background()
	s := NewCheckoutService(ctx)
	req := &checkout.CheckoutReq{
		UserId:    1,
		Firstname: "John",
		Lastname:  "Doe",
		Email:     "1231232@gmail.com",
		Address: &checkout.Address{
			StreetAddress: "7th street",
			City:          "hangzhou",
			State:         "zhejiang",
			Country:       "china",
			ZipCode:       "123131",
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber:          "424242424242424242",
			CreditCardCvv:             123,
			CreditCardExpirationYear:  2030,
			CreditCardExpirationMonth: 12,
		},
		PaymentMethod: "visa",
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	t.Logf("resp: %v", resp)

}
