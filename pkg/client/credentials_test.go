/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/identity"
	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/internal/test"
)

var TestCredentials *signingIdentity

func init() {
	privateKey, err := test.NewSM2PrivateKey()
	if err != nil {
		panic(err)
	}

	certificate, err := test.NewCertificate(privateKey)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity("mspID", certificate)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	TestCredentials = newSigningIdentity(id)
	TestCredentials.sign = sign
}
