/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package test

import (
	cb "github.com/hyperledger/fabric-protos-go/common"
	mspproto "github.com/hyperledger/fabric-protos-go/msp"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"github.com/m4ru1/fabric-gm-bdais/common/channelconfig"
	"github.com/m4ru1/fabric-gm-bdais/common/flogging"
	"github.com/m4ru1/fabric-gm-bdais/common/genesis"
	"github.com/m4ru1/fabric-gm-bdais/core/config/configtest"
	"github.com/m4ru1/fabric-gm-bdais/internal/configtxgen/encoder"
	"github.com/m4ru1/fabric-gm-bdais/internal/configtxgen/genesisconfig"
	"github.com/m4ru1/fabric-gm-bdais/internal/pkg/txflags"
	"github.com/m4ru1/fabric-gm-bdais/protoutil"
)

var logger = flogging.MustGetLogger("common.configtx.test")

// MakeGenesisBlock creates a genesis block using the test templates for the given channelID
func MakeGenesisBlock(channelID string) (*cb.Block, error) {
	profile := genesisconfig.Load(genesisconfig.SampleDevModeSoloProfile, configtest.GetDevConfigDir())
	channelGroup, err := encoder.NewChannelGroup(profile)
	if err != nil {
		logger.Panicf("Error creating channel config: %s", err)
	}

	gb := genesis.NewFactoryImpl(channelGroup).Block(channelID)
	if gb == nil {
		return gb, nil
	}

	txsFilter := txflags.NewWithValues(len(gb.Data.Data), pb.TxValidationCode_VALID)
	gb.Metadata.Metadata[cb.BlockMetadataIndex_TRANSACTIONS_FILTER] = txsFilter

	return gb, nil
}

func MakeChannelConfig(channelID string) (*cb.Config, error) {
	profile := genesisconfig.Load(genesisconfig.SampleDevModeSoloProfile, configtest.GetDevConfigDir())
	channelGroup, err := encoder.NewChannelGroup(profile)
	if err != nil {
		return nil, err
	}
	return &cb.Config{ChannelGroup: channelGroup}, nil
}

// MakeGenesisBlockWithMSPs creates a genesis block using the MSPs provided for the given channelID
func MakeGenesisBlockFromMSPs(channelID string, appMSPConf, ordererMSPConf *mspproto.MSPConfig, appOrgID, ordererOrgID string) (*cb.Block, error) {
	profile := genesisconfig.Load(genesisconfig.SampleDevModeSoloProfile, configtest.GetDevConfigDir())
	profile.Orderer.Organizations = nil
	channelGroup, err := encoder.NewChannelGroup(profile)
	if err != nil {
		logger.Panicf("Error creating channel config: %s", err)
	}

	ordererOrg := protoutil.NewConfigGroup()
	ordererOrg.ModPolicy = channelconfig.AdminsPolicyKey
	ordererOrg.Values[channelconfig.MSPKey] = &cb.ConfigValue{
		Value:     protoutil.MarshalOrPanic(channelconfig.MSPValue(ordererMSPConf).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}

	applicationOrg := protoutil.NewConfigGroup()
	applicationOrg.ModPolicy = channelconfig.AdminsPolicyKey
	applicationOrg.Values[channelconfig.MSPKey] = &cb.ConfigValue{
		Value:     protoutil.MarshalOrPanic(channelconfig.MSPValue(appMSPConf).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}
	applicationOrg.Values[channelconfig.AnchorPeersKey] = &cb.ConfigValue{
		Value:     protoutil.MarshalOrPanic(channelconfig.AnchorPeersValue([]*pb.AnchorPeer{}).Value()),
		ModPolicy: channelconfig.AdminsPolicyKey,
	}

	channelGroup.Groups[channelconfig.OrdererGroupKey].Groups[ordererOrgID] = ordererOrg
	channelGroup.Groups[channelconfig.ApplicationGroupKey].Groups[appOrgID] = applicationOrg

	return genesis.NewFactoryImpl(channelGroup).Block(channelID), nil
}
