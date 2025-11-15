package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommonSIB_ssb_PositionsInBurst struct {
	inOneGroup    uper.BitString  `lb:8,ub:8,madatory`
	groupPresence *uper.BitString `lb:8,ub:8,optional`
}

func (ie *ServingCellConfigCommonSIB_ssb_PositionsInBurst) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.groupPresence != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.inOneGroup.Bytes, uint(ie.inOneGroup.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString inOneGroup", err)
	}
	if ie.groupPresence != nil {
		if err = w.WriteBitString(ie.groupPresence.Bytes, uint(ie.groupPresence.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode groupPresence", err)
		}
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_ssb_PositionsInBurst) Decode(r *uper.UperReader) error {
	var err error
	var groupPresencePresent bool
	if groupPresencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_inOneGroup []byte
	var n_inOneGroup uint
	if tmp_bs_inOneGroup, n_inOneGroup, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString inOneGroup", err)
	}
	ie.inOneGroup = uper.BitString{
		Bytes:   tmp_bs_inOneGroup,
		NumBits: uint64(n_inOneGroup),
	}
	if groupPresencePresent {
		var tmp_bs_groupPresence []byte
		var n_groupPresence uint
		if tmp_bs_groupPresence, n_groupPresence, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode groupPresence", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_groupPresence,
			NumBits: uint64(n_groupPresence),
		}
		ie.groupPresence = &tmp_bitstring
	}
	return nil
}
