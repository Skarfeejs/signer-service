package service

import (
	"context"
	"github.com/orbs-network/crypto-lib-go/crypto/ethereum/digest"
	"github.com/orbs-network/crypto-lib-go/test/crypto/ethereum/keys"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/services"
	"github.com/orbs-network/scribe/log"
	"github.com/stretchr/testify/require"
	"testing"
)

type signerServiceConfig struct {
}

func (s *signerServiceConfig) NodePrivateKey() primitives.EcdsaSecp256K1PrivateKey {
	return keys.EcdsaSecp256K1KeyPairForTests(0).PrivateKey()
}

func TestService_NodeSign(t *testing.T) {
	cfg := &signerServiceConfig{}
	pk := cfg.NodePrivateKey()

	testOutput := log.NewTestOutput(t, log.NewHumanReadableFormatter())
	testLogger := log.GetLogger().WithOutput(testOutput)

	service := NewService(cfg, testLogger)

	payload := []byte("payload")

	signed, err := digest.SignAsNode(pk, payload)
	require.NoError(t, err)

	clientSignature, err := service.NodeSign(context.TODO(), (&services.NodeSignInputBuilder{
		Data: payload,
	}).Build())
	require.NoError(t, err)

	require.EqualValues(t, signed, clientSignature.Signature())
}

// Contract values for JS
func Test_Payload(t *testing.T) {
	payload := []byte{1, 2, 3}

	raw := (&services.NodeSignInputBuilder{
		Data: payload,
	}).Build().Raw()

	require.EqualValues(t,[]byte{3, 0, 0, 0, 1, 2, 3}, raw)
}