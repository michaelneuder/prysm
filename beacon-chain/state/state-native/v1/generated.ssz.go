// Code generated by fastssz. DO NOT EDIT.
// Hash: 2e923b42b8e4fcc278301da6506b212334a78169cb32c70e0d66a636435b8925
package v1

import (
	ssz "github.com/ferranbt/fastssz"
	eth2types "github.com/prysmaticlabs/eth2-types"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(2687377)

	// Field (0) 'genesisTime'
	dst = ssz.MarshalUint64(dst, b.genesisTime)

	// Field (1) 'genesisValidatorsRoot'
	dst = append(dst, b.genesisValidatorsRoot[:]...)

	// Field (2) 'slot'
	dst = ssz.MarshalUint64(dst, uint64(b.slot))

	// Field (3) 'fork'
	if b.fork == nil {
		b.fork = new(ethpb.Fork)
	}
	if dst, err = b.fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'latestBlockHeader'
	if b.latestBlockHeader == nil {
		b.latestBlockHeader = new(ethpb.BeaconBlockHeader)
	}
	if dst, err = b.latestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'blockRoots'
	for ii := 0; ii < 8192; ii++ {
		dst = append(dst, b.blockRoots[ii][:]...)
	}

	// Field (6) 'stateRoots'
	for ii := 0; ii < 8192; ii++ {
		dst = append(dst, b.stateRoots[ii][:]...)
	}

	// Offset (7) 'historicalRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.historicalRoots) * 32

	// Field (8) 'eth1Data'
	if b.eth1Data == nil {
		b.eth1Data = new(ethpb.Eth1Data)
	}
	if dst, err = b.eth1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (9) 'eth1DataVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.eth1DataVotes) * 72

	// Field (10) 'eth1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.eth1DepositIndex)

	// Offset (11) 'validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.validators) * 121

	// Offset (12) 'balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.balances) * 8

	// Field (13) 'randaoMixes'
	for ii := 0; ii < 65536; ii++ {
		dst = append(dst, b.randaoMixes[ii][:]...)
	}

	// Field (14) 'slashings'
	if len(b.slashings) != 8192 {
		err = ssz.ErrVectorLength
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = ssz.MarshalUint64(dst, b.slashings[ii])
	}

	// Offset (15) 'previousEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.previousEpochAttestations); ii++ {
		offset += 4
		offset += b.previousEpochAttestations[ii].SizeSSZ()
	}

	// Offset (16) 'currentEpochAttestations'
	dst = ssz.WriteOffset(dst, offset)
	for ii := 0; ii < len(b.currentEpochAttestations); ii++ {
		offset += 4
		offset += b.currentEpochAttestations[ii].SizeSSZ()
	}

	// Field (17) 'justificationBits'
	if len(b.justificationBits) != 1 {
		err = ssz.ErrBytesLength
		return
	}
	dst = append(dst, b.justificationBits...)

	// Field (18) 'previousJustifiedCheckpoint'
	if b.previousJustifiedCheckpoint == nil {
		b.previousJustifiedCheckpoint = new(ethpb.Checkpoint)
	}
	if dst, err = b.previousJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (19) 'currentJustifiedCheckpoint'
	if b.currentJustifiedCheckpoint == nil {
		b.currentJustifiedCheckpoint = new(ethpb.Checkpoint)
	}
	if dst, err = b.currentJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (20) 'finalizedCheckpoint'
	if b.finalizedCheckpoint == nil {
		b.finalizedCheckpoint = new(ethpb.Checkpoint)
	}
	if dst, err = b.finalizedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (7) 'historicalRoots'
	if len(b.historicalRoots) > 16777216 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.historicalRoots); ii++ {
		dst = append(dst, b.historicalRoots[ii][:]...)
	}

	// Field (9) 'eth1DataVotes'
	if len(b.eth1DataVotes) > 2048 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.eth1DataVotes); ii++ {
		if dst, err = b.eth1DataVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (11) 'validators'
	if len(b.validators) > 1099511627776 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.validators); ii++ {
		if dst, err = b.validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (12) 'balances'
	if len(b.balances) > 1099511627776 {
		err = ssz.ErrListTooBig
		return
	}
	for ii := 0; ii < len(b.balances); ii++ {
		dst = ssz.MarshalUint64(dst, b.balances[ii])
	}

	// Field (15) 'previousEpochAttestations'
	if len(b.previousEpochAttestations) > 4096 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.previousEpochAttestations)
		for ii := 0; ii < len(b.previousEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.previousEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.previousEpochAttestations); ii++ {
		if dst, err = b.previousEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (16) 'currentEpochAttestations'
	if len(b.currentEpochAttestations) > 4096 {
		err = ssz.ErrListTooBig
		return
	}
	{
		offset = 4 * len(b.currentEpochAttestations)
		for ii := 0; ii < len(b.currentEpochAttestations); ii++ {
			dst = ssz.WriteOffset(dst, offset)
			offset += b.currentEpochAttestations[ii].SizeSSZ()
		}
	}
	for ii := 0; ii < len(b.currentEpochAttestations); ii++ {
		if dst, err = b.currentEpochAttestations[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 2687377 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o9, o11, o12, o15, o16 uint64

	// Field (0) 'genesisTime'
	b.genesisTime = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'genesisValidatorsRoot'
	copy(b.genesisValidatorsRoot[:], buf[8:40])

	// Field (2) 'slot'
	b.slot = eth2types.Slot(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'fork'
	if b.fork == nil {
		b.fork = new(ethpb.Fork)
	}
	if err = b.fork.UnmarshalSSZ(buf[48:64]); err != nil {
		return err
	}

	// Field (4) 'latestBlockHeader'
	if b.latestBlockHeader == nil {
		b.latestBlockHeader = new(ethpb.BeaconBlockHeader)
	}
	if err = b.latestBlockHeader.UnmarshalSSZ(buf[64:176]); err != nil {
		return err
	}

	// Field (5) 'blockRoots'

	for ii := 0; ii < 8192; ii++ {
		copy(b.blockRoots[ii][:], buf[176:262320][ii*32:(ii+1)*32])
	}

	// Field (6) 'stateRoots'

	for ii := 0; ii < 8192; ii++ {
		copy(b.stateRoots[ii][:], buf[262320:524464][ii*32:(ii+1)*32])
	}

	// Offset (7) 'historicalRoots'
	if o7 = ssz.ReadOffset(buf[524464:524468]); o7 > size {
		return ssz.ErrOffset
	}

	if o7 < 2687377 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (8) 'eth1Data'
	if b.eth1Data == nil {
		b.eth1Data = new(ethpb.Eth1Data)
	}
	if err = b.eth1Data.UnmarshalSSZ(buf[524468:524540]); err != nil {
		return err
	}

	// Offset (9) 'eth1DataVotes'
	if o9 = ssz.ReadOffset(buf[524540:524544]); o9 > size || o7 > o9 {
		return ssz.ErrOffset
	}

	// Field (10) 'eth1DepositIndex'
	b.eth1DepositIndex = ssz.UnmarshallUint64(buf[524544:524552])

	// Offset (11) 'validators'
	if o11 = ssz.ReadOffset(buf[524552:524556]); o11 > size || o9 > o11 {
		return ssz.ErrOffset
	}

	// Offset (12) 'balances'
	if o12 = ssz.ReadOffset(buf[524556:524560]); o12 > size || o11 > o12 {
		return ssz.ErrOffset
	}

	// Field (13) 'randaoMixes'

	for ii := 0; ii < 65536; ii++ {
		copy(b.randaoMixes[ii][:], buf[524560:2621712][ii*32:(ii+1)*32])
	}

	// Field (14) 'slashings'
	b.slashings = ssz.ExtendUint64(b.slashings, 8192)
	for ii := 0; ii < 8192; ii++ {
		b.slashings[ii] = ssz.UnmarshallUint64(buf[2621712:2687248][ii*8 : (ii+1)*8])
	}

	// Offset (15) 'previousEpochAttestations'
	if o15 = ssz.ReadOffset(buf[2687248:2687252]); o15 > size || o12 > o15 {
		return ssz.ErrOffset
	}

	// Offset (16) 'currentEpochAttestations'
	if o16 = ssz.ReadOffset(buf[2687252:2687256]); o16 > size || o15 > o16 {
		return ssz.ErrOffset
	}

	// Field (17) 'justificationBits'
	if cap(b.justificationBits) == 0 {
		b.justificationBits = make([]byte, 0, len(buf[2687256:2687257]))
	}
	b.justificationBits = append(b.justificationBits, buf[2687256:2687257]...)

	// Field (18) 'previousJustifiedCheckpoint'
	if b.previousJustifiedCheckpoint == nil {
		b.previousJustifiedCheckpoint = new(ethpb.Checkpoint)
	}
	if err = b.previousJustifiedCheckpoint.UnmarshalSSZ(buf[2687257:2687297]); err != nil {
		return err
	}

	// Field (19) 'currentJustifiedCheckpoint'
	if b.currentJustifiedCheckpoint == nil {
		b.currentJustifiedCheckpoint = new(ethpb.Checkpoint)
	}
	if err = b.currentJustifiedCheckpoint.UnmarshalSSZ(buf[2687297:2687337]); err != nil {
		return err
	}

	// Field (20) 'finalizedCheckpoint'
	if b.finalizedCheckpoint == nil {
		b.finalizedCheckpoint = new(ethpb.Checkpoint)
	}
	if err = b.finalizedCheckpoint.UnmarshalSSZ(buf[2687337:2687377]); err != nil {
		return err
	}

	// Field (7) 'historicalRoots'
	{
		buf = tail[o7:o9]
		num, err := ssz.DivideInt2(len(buf), 32, 16777216)
		if err != nil {
			return err
		}
		b.historicalRoots = make([][32]byte, num)
		for ii := 0; ii < num; ii++ {
			copy(b.historicalRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (9) 'eth1DataVotes'
	{
		buf = tail[o9:o11]
		num, err := ssz.DivideInt2(len(buf), 72, 2048)
		if err != nil {
			return err
		}
		b.eth1DataVotes = make([]*ethpb.Eth1Data, num)
		for ii := 0; ii < num; ii++ {
			if b.eth1DataVotes[ii] == nil {
				b.eth1DataVotes[ii] = new(ethpb.Eth1Data)
			}
			if err = b.eth1DataVotes[ii].UnmarshalSSZ(buf[ii*72 : (ii+1)*72]); err != nil {
				return err
			}
		}
	}

	// Field (11) 'validators'
	{
		buf = tail[o11:o12]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.validators = make([]*ethpb.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.validators[ii] == nil {
				b.validators[ii] = new(ethpb.Validator)
			}
			if err = b.validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (12) 'balances'
	{
		buf = tail[o12:o15]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.balances = ssz.ExtendUint64(b.balances, num)
		for ii := 0; ii < num; ii++ {
			b.balances[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (15) 'previousEpochAttestations'
	{
		buf = tail[o15:o16]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.previousEpochAttestations = make([]*ethpb.PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.previousEpochAttestations[indx] == nil {
				b.previousEpochAttestations[indx] = new(ethpb.PendingAttestation)
			}
			if err = b.previousEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	// Field (16) 'currentEpochAttestations'
	{
		buf = tail[o16:]
		num, err := ssz.DecodeDynamicLength(buf, 4096)
		if err != nil {
			return err
		}
		b.currentEpochAttestations = make([]*ethpb.PendingAttestation, num)
		err = ssz.UnmarshalDynamic(buf, num, func(indx int, buf []byte) (err error) {
			if b.currentEpochAttestations[indx] == nil {
				b.currentEpochAttestations[indx] = new(ethpb.PendingAttestation)
			}
			if err = b.currentEpochAttestations[indx].UnmarshalSSZ(buf); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 2687377

	// Field (7) 'historicalRoots'
	size += len(b.historicalRoots) * 32

	// Field (9) 'eth1DataVotes'
	size += len(b.eth1DataVotes) * 72

	// Field (11) 'validators'
	size += len(b.validators) * 121

	// Field (12) 'balances'
	size += len(b.balances) * 8

	// Field (15) 'previousEpochAttestations'
	for ii := 0; ii < len(b.previousEpochAttestations); ii++ {
		size += 4
		size += b.previousEpochAttestations[ii].SizeSSZ()
	}

	// Field (16) 'currentEpochAttestations'
	for ii := 0; ii < len(b.currentEpochAttestations); ii++ {
		size += 4
		size += b.currentEpochAttestations[ii].SizeSSZ()
	}

	return
}
