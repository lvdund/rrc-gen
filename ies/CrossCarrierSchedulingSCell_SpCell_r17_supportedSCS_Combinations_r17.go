package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17 struct {
	scs15kHz_15kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_15kHz_r17 `optional`
	scs15kHz_30kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_30kHz_r17 `optional`
	scs15kHz_60kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_60kHz_r17 `optional`
	scs30kHz_30kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
	scs30kHz_60kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
	scs60kHz_60kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs15kHz_15kHz_r17 != nil, ie.scs15kHz_30kHz_r17 != nil, ie.scs15kHz_60kHz_r17 != nil, ie.scs30kHz_30kHz_r17 != nil, ie.scs30kHz_60kHz_r17 != nil, ie.scs60kHz_60kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs15kHz_15kHz_r17 != nil {
		if err = ie.scs15kHz_15kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs15kHz_15kHz_r17", err)
		}
	}
	if ie.scs15kHz_30kHz_r17 != nil {
		if err = ie.scs15kHz_30kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs15kHz_30kHz_r17", err)
		}
	}
	if ie.scs15kHz_60kHz_r17 != nil {
		if err = ie.scs15kHz_60kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs15kHz_60kHz_r17", err)
		}
	}
	if ie.scs30kHz_30kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs30kHz_30kHz_r17.Bytes, uint(ie.scs30kHz_30kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode scs30kHz_30kHz_r17", err)
		}
	}
	if ie.scs30kHz_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs30kHz_60kHz_r17.Bytes, uint(ie.scs30kHz_60kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode scs30kHz_60kHz_r17", err)
		}
	}
	if ie.scs60kHz_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs60kHz_60kHz_r17.Bytes, uint(ie.scs60kHz_60kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode scs60kHz_60kHz_r17", err)
		}
	}
	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs15kHz_15kHz_r17Present bool
	if scs15kHz_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs15kHz_30kHz_r17Present bool
	if scs15kHz_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs15kHz_60kHz_r17Present bool
	if scs15kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs30kHz_30kHz_r17Present bool
	if scs30kHz_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs30kHz_60kHz_r17Present bool
	if scs30kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs60kHz_60kHz_r17Present bool
	if scs60kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs15kHz_15kHz_r17Present {
		ie.scs15kHz_15kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_15kHz_r17)
		if err = ie.scs15kHz_15kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs15kHz_15kHz_r17", err)
		}
	}
	if scs15kHz_30kHz_r17Present {
		ie.scs15kHz_30kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_30kHz_r17)
		if err = ie.scs15kHz_30kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs15kHz_30kHz_r17", err)
		}
	}
	if scs15kHz_60kHz_r17Present {
		ie.scs15kHz_60kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_60kHz_r17)
		if err = ie.scs15kHz_60kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs15kHz_60kHz_r17", err)
		}
	}
	if scs30kHz_30kHz_r17Present {
		var tmp_bs_scs30kHz_30kHz_r17 []byte
		var n_scs30kHz_30kHz_r17 uint
		if tmp_bs_scs30kHz_30kHz_r17, n_scs30kHz_30kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode scs30kHz_30kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs30kHz_30kHz_r17,
			NumBits: uint64(n_scs30kHz_30kHz_r17),
		}
		ie.scs30kHz_30kHz_r17 = &tmp_bitstring
	}
	if scs30kHz_60kHz_r17Present {
		var tmp_bs_scs30kHz_60kHz_r17 []byte
		var n_scs30kHz_60kHz_r17 uint
		if tmp_bs_scs30kHz_60kHz_r17, n_scs30kHz_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode scs30kHz_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs30kHz_60kHz_r17,
			NumBits: uint64(n_scs30kHz_60kHz_r17),
		}
		ie.scs30kHz_60kHz_r17 = &tmp_bitstring
	}
	if scs60kHz_60kHz_r17Present {
		var tmp_bs_scs60kHz_60kHz_r17 []byte
		var n_scs60kHz_60kHz_r17 uint
		if tmp_bs_scs60kHz_60kHz_r17, n_scs60kHz_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode scs60kHz_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs60kHz_60kHz_r17,
			NumBits: uint64(n_scs60kHz_60kHz_r17),
		}
		ie.scs60kHz_60kHz_r17 = &tmp_bitstring
	}
	return nil
}
