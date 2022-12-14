package weshnet_test

import (
	"context"
	"fmt"
	"testing"

	keystore "github.com/ipfs/go-ipfs-keystore"
	"github.com/stretchr/testify/assert"

	"berty.tech/weshnet/internal/cryptoutil"
	"berty.tech/weshnet/internal/testutil"
	"berty.tech/weshnet/pkg/protocoltypes"
)

func TestTestingClient_impl(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, cleanup := testutil.Logger(t)
	defer cleanup()

	client, cleanup := TestingService(ctx, t, Opts{
		Logger:         logger,
		DeviceKeystore: cryptoutil.NewDeviceKeystore(keystore.NewMemKeystore(), nil),
	})
	defer cleanup()

	// test service
	_, _ = client.InstanceGetConfiguration(ctx, &protocoltypes.InstanceGetConfiguration_Request{})
	status := client.Status()
	expected := Status{}
	assert.Equal(t, expected, status)
}

func ExampleNew_basic() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := New(Opts{})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ret, err := client.InstanceGetConfiguration(ctx, &protocoltypes.InstanceGetConfiguration_Request{})
	if err != nil {
		panic(err)
	}

	for _, listener := range ret.Listeners {
		if listener == "/p2p-circuit" {
			fmt.Println(listener)
		}
	}

	// Output:
	// /p2p-circuit
}

// FIXME: create examples that actually use groups and contacts
