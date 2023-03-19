/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package consensus

import "github.com/hyperledger/fabric/orderer/common/types"

// StatusReporter is implemented by cluster-type Chain implementations.
// It allows the node to report its cluster relation and its status within that relation.
// This information is used to generate the channelparticipation.ChannelInfo in response
// to a "List" request on a particular channel.
//
<<<<<<< HEAD
// Not all chains must implement this, in particular non-cluster-type (solo, kafka) are
=======
// Not all chains must implement this, in particular non-cluster-type (solo) are
>>>>>>> a5405e2ca41902d62fe0fa9caa102e0d818c2f19
// assigned a StaticStatusReporter at construction time.
type StatusReporter interface {
	// StatusReport provides the cluster relation and status.
	// See:  channelparticipation.ChannelInfo for more details.
	StatusReport() (types.ConsensusRelation, types.Status)
}

// StaticStatusReporter is intended for chains that do not implement the StatusReporter interface.
type StaticStatusReporter struct {
	ConsensusRelation types.ConsensusRelation
	Status            types.Status
}

func (s StaticStatusReporter) StatusReport() (types.ConsensusRelation, types.Status) {
	return s.ConsensusRelation, s.Status
}
