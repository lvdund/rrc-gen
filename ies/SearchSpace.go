package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpace struct {
	searchSpaceId                      SearchSpaceId                                   `madatory`
	controlResourceSetId               *ControlResourceSetId                           `optional`
	monitoringSlotPeriodicityAndOffset *SearchSpace_monitoringSlotPeriodicityAndOffset `lb:0,ub:1,optional`
	duration                           *int64                                          `lb:2,ub:2559,optional`
	monitoringSymbolsWithinSlot        *uper.BitString                                 `lb:14,ub:14,optional`
	nrofCandidates                     *SearchSpace_nrofCandidates                     `optional`
	searchSpaceType                    *SearchSpace_searchSpaceType                    `optional`
}

func (ie *SearchSpace) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.controlResourceSetId != nil, ie.monitoringSlotPeriodicityAndOffset != nil, ie.duration != nil, ie.monitoringSymbolsWithinSlot != nil, ie.nrofCandidates != nil, ie.searchSpaceType != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.searchSpaceId.Encode(w); err != nil {
		return utils.WrapError("Encode searchSpaceId", err)
	}
	if ie.controlResourceSetId != nil {
		if err = ie.controlResourceSetId.Encode(w); err != nil {
			return utils.WrapError("Encode controlResourceSetId", err)
		}
	}
	if ie.monitoringSlotPeriodicityAndOffset != nil {
		if err = ie.monitoringSlotPeriodicityAndOffset.Encode(w); err != nil {
			return utils.WrapError("Encode monitoringSlotPeriodicityAndOffset", err)
		}
	}
	if ie.duration != nil {
		if err = w.WriteInteger(*ie.duration, &uper.Constraint{Lb: 2, Ub: 2559}, false); err != nil {
			return utils.WrapError("Encode duration", err)
		}
	}
	if ie.monitoringSymbolsWithinSlot != nil {
		if err = w.WriteBitString(ie.monitoringSymbolsWithinSlot.Bytes, uint(ie.monitoringSymbolsWithinSlot.NumBits), &uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Encode monitoringSymbolsWithinSlot", err)
		}
	}
	if ie.nrofCandidates != nil {
		if err = ie.nrofCandidates.Encode(w); err != nil {
			return utils.WrapError("Encode nrofCandidates", err)
		}
	}
	if ie.searchSpaceType != nil {
		if err = ie.searchSpaceType.Encode(w); err != nil {
			return utils.WrapError("Encode searchSpaceType", err)
		}
	}
	return nil
}

func (ie *SearchSpace) Decode(r *uper.UperReader) error {
	var err error
	var controlResourceSetIdPresent bool
	if controlResourceSetIdPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var monitoringSlotPeriodicityAndOffsetPresent bool
	if monitoringSlotPeriodicityAndOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var durationPresent bool
	if durationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var monitoringSymbolsWithinSlotPresent bool
	if monitoringSymbolsWithinSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nrofCandidatesPresent bool
	if nrofCandidatesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var searchSpaceTypePresent bool
	if searchSpaceTypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.searchSpaceId.Decode(r); err != nil {
		return utils.WrapError("Decode searchSpaceId", err)
	}
	if controlResourceSetIdPresent {
		ie.controlResourceSetId = new(ControlResourceSetId)
		if err = ie.controlResourceSetId.Decode(r); err != nil {
			return utils.WrapError("Decode controlResourceSetId", err)
		}
	}
	if monitoringSlotPeriodicityAndOffsetPresent {
		ie.monitoringSlotPeriodicityAndOffset = new(SearchSpace_monitoringSlotPeriodicityAndOffset)
		if err = ie.monitoringSlotPeriodicityAndOffset.Decode(r); err != nil {
			return utils.WrapError("Decode monitoringSlotPeriodicityAndOffset", err)
		}
	}
	if durationPresent {
		var tmp_int_duration int64
		if tmp_int_duration, err = r.ReadInteger(&uper.Constraint{Lb: 2, Ub: 2559}, false); err != nil {
			return utils.WrapError("Decode duration", err)
		}
		ie.duration = &tmp_int_duration
	}
	if monitoringSymbolsWithinSlotPresent {
		var tmp_bs_monitoringSymbolsWithinSlot []byte
		var n_monitoringSymbolsWithinSlot uint
		if tmp_bs_monitoringSymbolsWithinSlot, n_monitoringSymbolsWithinSlot, err = r.ReadBitString(&uper.Constraint{Lb: 14, Ub: 14}, false); err != nil {
			return utils.WrapError("Decode monitoringSymbolsWithinSlot", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_monitoringSymbolsWithinSlot,
			NumBits: uint64(n_monitoringSymbolsWithinSlot),
		}
		ie.monitoringSymbolsWithinSlot = &tmp_bitstring
	}
	if nrofCandidatesPresent {
		ie.nrofCandidates = new(SearchSpace_nrofCandidates)
		if err = ie.nrofCandidates.Decode(r); err != nil {
			return utils.WrapError("Decode nrofCandidates", err)
		}
	}
	if searchSpaceTypePresent {
		ie.searchSpaceType = new(SearchSpace_searchSpaceType)
		if err = ie.searchSpaceType.Decode(r); err != nil {
			return utils.WrapError("Decode searchSpaceType", err)
		}
	}
	return nil
}
