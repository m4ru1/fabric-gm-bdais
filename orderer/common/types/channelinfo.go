/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package types

// ErrorResponse carries the error response an HTTP request.
// This is marshaled into the body of the HTTP response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ChannelList carries the response to an HTTP request to List all the channels.
// This is marshaled into the body of the HTTP response.
// swagger:model channelList
type ChannelList struct {
	// The system channel info, nil if it doesn't exist.
	SystemChannel *ChannelInfoShort `json:"systemChannel"`
	// Application channels only, nil or empty if no channels defined.
	Channels []ChannelInfoShort `json:"channels"`
}

// ChannelInfoShort carries a short info of a single channel.
type ChannelInfoShort struct {
	// The channel name.
	Name string `json:"name"`
	// The channel relative URL (no Host:Port, only path), e.g.: "/participation/v1/channels/my-channel".
	URL string `json:"url"`
}

// ConsensusRelation represents the relationship between the orderer and the channel's consensus cluster.
type ConsensusRelation string

const (
	// The orderer is a cluster consenter of a cluster consensus protocol (e.g. etcdraft) for a specific channel.
	// That is, the orderer is in the consenters set of the channel.
	ConsensusRelationConsenter ConsensusRelation = "consenter"
	// The orderer is following a cluster consensus protocol by pulling blocks from other orderers.
	// The orderer is NOT in the consenters set of the channel.
	ConsensusRelationFollower ConsensusRelation = "follower"
	// The orderer is NOT in the consenters set of the channel, and is just tracking (polling) the last config block
	// of the channel in order to detect when it is added to the channel.
	ConsensusRelationConfigTracker ConsensusRelation = "config-tracker"
<<<<<<< HEAD
	// The orderer runs a non-cluster consensus type, solo or kafka.
=======
	// The orderer runs a non-cluster consensus type, i.e. solo.
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	ConsensusRelationOther ConsensusRelation = "other"
)

// Status represents the degree by which the orderer had caught up with the rest of the cluster after joining the
// channel (either as a consenter or a follower).
type Status string

const (
	// The orderer is active in the channel's consensus protocol, or following the cluster,
	// with block height > the join-block number. (Height is last block number +1).
	StatusActive Status = "active"
	// The orderer is catching up with the cluster by pulling blocks from other orderers,
	// with block height <= the join-block number.
	StatusOnBoarding Status = "onboarding"
	// The orderer is not storing any blocks for this channel.
	StatusInactive Status = "inactive"
	// The last orderer operation against the channel failed.
	StatusFailed Status = "failed"
)

// ChannelInfo carries the response to an HTTP request to List a single channel.
// This is marshaled into the body of the HTTP response.
// swagger:model channelInfo
type ChannelInfo struct {
	// The channel name.
	Name string `json:"name"`
	// The channel relative URL (no Host:Port, only path), e.g.: "/participation/v1/channels/my-channel".
	URL string `json:"url"`
	// Whether the orderer is a “consenter”, ”follower”, or "config-tracker" of
	// the cluster for this channel.
<<<<<<< HEAD
	// For non cluster consensus types (solo, kafka) it is "other".
	// Possible values:  “consenter”, ”follower”, "config-tracker", "other".
	ConsensusRelation ConsensusRelation `json:"consensusRelation"`
	// Whether the orderer is ”onboarding”, ”active”, or "inactive", for this channel.
	// For non cluster consensus types (solo, kafka) it is "active".
=======
	// For non cluster consensus types (solo) it is "other".
	// Possible values:  “consenter”, ”follower”, "config-tracker", "other".
	ConsensusRelation ConsensusRelation `json:"consensusRelation"`
	// Whether the orderer is ”onboarding”, ”active”, or "inactive", for this channel.
	// For non cluster consensus types (solo) it is "active".
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
	// Possible values:  “onboarding”, ”active”, "inactive".
	Status Status `json:"status"`
	// Current block height.
	Height uint64 `json:"height"`
}
