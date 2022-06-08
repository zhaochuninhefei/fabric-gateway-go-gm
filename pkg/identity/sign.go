/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package identity

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"gitee.com/zhaochuninhefei/gmgo/sm2"
)

// Sign function generates a digital signature of the supplied digest.
type Sign = func(digest []byte) ([]byte, error)

// NewPrivateKeySign returns a Sign function that uses the supplied private key.
func NewPrivateKeySign(privateKey crypto.PrivateKey) (Sign, error) {
	switch key := privateKey.(type) {
	case *sm2.PrivateKey:
		return sm2PrivateKeySign(key), nil
	default:
		return nil, fmt.Errorf("unsupported key type: %T", privateKey)
	}
}

func sm2PrivateKeySign(privateKey *sm2.PrivateKey) Sign {
	return func(digest []byte) ([]byte, error) {
		r, s, err := sm2.Sign(rand.Reader, privateKey, digest)
		if err != nil {
			return nil, err
		}
		// sm2不考虑将s转为低值
		// s, err = toLowSByKey(&privateKey.PublicKey, s)
		// if err != nil {
		// 	return nil, err
		// }

		return marshalSM2Signature(r, s)
	}
}
