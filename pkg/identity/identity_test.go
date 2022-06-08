/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"testing"

	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/internal/test"
	"github.com/stretchr/testify/require"
)

func TestIdentity(t *testing.T) {
	const mspID = "mspID"

	privateKey, err := test.NewSM2PrivateKey()
	require.NoError(t, err)

	certificate, err := test.NewCertificate(privateKey)
	require.NoError(t, err)

	t.Run("NewX509Identity", func(t *testing.T) {
		identity, err := NewX509Identity(mspID, certificate)
		require.NoError(t, err)

		require.Equal(t, mspID, identity.MspID())
	})
}