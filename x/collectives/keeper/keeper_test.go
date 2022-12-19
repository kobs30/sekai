package keeper_test

import (
	"encoding/hex"
	"testing"

	simapp "github.com/KiraCore/sekai/app"
	"github.com/KiraCore/sekai/x/collectives/types"
	govtypes "github.com/KiraCore/sekai/x/gov/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func newPubKey(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	pubkey := &ed25519.PubKey{Key: pkBytes}

	return pubkey
}

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.SekaiApp
}

func (suite *KeeperTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)

	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.app = app
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestIsAllowedAddress() {
	pubkeys := []cryptotypes.PubKey{
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB50"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB51"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB52"),
	}
	addrs := []sdk.AccAddress{
		pubkeys[0].Address().Bytes(),
		pubkeys[1].Address().Bytes(),
		pubkeys[2].Address().Bytes(),
	}

	actor := govtypes.NetworkActor{
		Address: addrs[0],
		Permissions: &govtypes.Permissions{
			Whitelist: []uint32{uint32(govtypes.PermHandleBasketEmergency)},
		},
		Roles: []uint64{1},
	}
	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, actor)
	suite.app.CustomGovKeeper.AssignRoleToActor(suite.ctx, actor, 1)

	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, govtypes.NetworkActor{
		Address: addrs[1],
		Roles:   []uint64{},
	})

	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, govtypes.NetworkActor{
		Address: addrs[2],
		Roles:   []uint64{},
	})

	whitelistInfo := types.OwnersWhitelist{
		Roles:    []uint64{1},
		Accounts: []string{addrs[1].String()},
	}

	isAllowed := suite.app.CollectivesKeeper.IsAllowedAddress(suite.ctx, addrs[0], whitelistInfo)
	suite.Require().True(isAllowed)
	isAllowed = suite.app.CollectivesKeeper.IsAllowedAddress(suite.ctx, addrs[1], whitelistInfo)
	suite.Require().True(isAllowed)
	isAllowed = suite.app.CollectivesKeeper.IsAllowedAddress(suite.ctx, addrs[2], whitelistInfo)
	suite.Require().False(isAllowed)
}

func (suite *KeeperTestSuite) TestAllowedAddresses() {
	pubkeys := []cryptotypes.PubKey{
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB50"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB51"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB52"),
	}
	addrs := []sdk.AccAddress{
		pubkeys[0].Address().Bytes(),
		pubkeys[1].Address().Bytes(),
		pubkeys[2].Address().Bytes(),
	}

	actor := govtypes.NetworkActor{
		Address: addrs[0],
		Permissions: &govtypes.Permissions{
			Whitelist: []uint32{uint32(govtypes.PermHandleBasketEmergency)},
		},
		Roles: []uint64{1},
	}
	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, actor)
	suite.app.CustomGovKeeper.AssignRoleToActor(suite.ctx, actor, 1)

	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, govtypes.NetworkActor{
		Address: addrs[1],
		Roles:   []uint64{},
	})

	suite.app.CustomGovKeeper.SaveNetworkActor(suite.ctx, govtypes.NetworkActor{
		Address: addrs[2],
		Roles:   []uint64{},
	})

	whitelistInfo := types.OwnersWhitelist{
		Roles:    []uint64{1},
		Accounts: []string{addrs[1].String()},
	}

	allowedAddrs := suite.app.CollectivesKeeper.AllowedAddresses(suite.ctx, whitelistInfo)
	suite.Require().Len(allowedAddrs, 2)
}