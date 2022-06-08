/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package client

// import (
// 	"context"
// 	"testing"

// 	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/identity"
// 	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/internal/test"
// 	"gitee.com/zhaochuninhefei/fabric-gateway-go-gm/pkg/internal/util"
// 	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/gateway"
// 	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/msp"
// 	"gitee.com/zhaochuninhefei/fabric-protos-go-gm/peer"
// 	"gitee.com/zhaochuninhefei/gmgo/grpc"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestIdentity(t *testing.T) {
// 	privateKey, err := test.NewSM2PrivateKey()
// 	require.NoError(t, err)

// 	certificate, err := test.NewCertificate(privateKey)
// 	require.NoError(t, err)

// 	id, err := identity.NewX509Identity("MSP_ID", certificate)
// 	require.NoError(t, err)

// 	serializedIdentity := &msp.SerializedIdentity{
// 		Mspid:   id.MspID(),
// 		IdBytes: id.Credentials(),
// 	}
// 	creator, err := util.Marshal(serializedIdentity)
// 	require.NoError(t, err)

// 	t.Run("Evaluate uses client identity for proposals", func(t *testing.T) {
// 		var actual []byte
// 		mockClient := NewMockGatewayClient(gomock.NewController(t))
// 		evaluateResponse := &gateway.EvaluateResponse{
// 			Result: &peer.Response{
// 				Payload: nil,
// 			},
// 		}
// 		mockClient.EXPECT().Evaluate(gomock.Any(), gomock.Any()).
// 			Do(func(_ context.Context, in *gateway.EvaluateRequest, _ ...grpc.CallOption) {
// 				actual = test.AssertUnmarshalSignatureHeader(t, in.ProposedTransaction).Creator
// 			}).
// 			Return(evaluateResponse, nil).
// 			Times(1)

// 		contract := AssertNewTestContract(t, "contract", WithGatewayClient(mockClient), WithIdentity(id))

// 		_, err := contract.EvaluateTransaction("transaction")
// 		require.NoError(t, err)

// 		require.EqualValues(t, creator, actual)
// 	})

// 	t.Run("Submit uses client identity for proposals", func(t *testing.T) {
// 		var actual []byte
// 		mockClient := NewMockGatewayClient(gomock.NewController(t))
// 		endorseResponse := AssertNewEndorseResponse(t, "result", "channel")
// 		statusResponse := &gateway.CommitStatusResponse{
// 			Result: peer.TxValidationCode_VALID,
// 		}
// 		mockClient.EXPECT().Endorse(gomock.Any(), gomock.Any()).
// 			Do(func(_ context.Context, in *gateway.EndorseRequest, _ ...grpc.CallOption) {
// 				actual = test.AssertUnmarshalSignatureHeader(t, in.ProposedTransaction).Creator
// 			}).
// 			Return(endorseResponse, nil).
// 			Times(1)
// 		mockClient.EXPECT().Submit(gomock.Any(), gomock.Any()).
// 			Return(nil, nil)
// 		mockClient.EXPECT().CommitStatus(gomock.Any(), gomock.Any()).
// 			Return(statusResponse, nil)

// 		contract := AssertNewTestContract(t, "contract", WithGatewayClient(mockClient), WithIdentity(id))

// 		_, err := contract.SubmitTransaction("transaction")
// 		require.NoError(t, err)

// 		require.EqualValues(t, creator, actual)
// 	})
// }
