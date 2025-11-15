package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PreconfigGeneral_r16 struct {
	sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon `optional`
	reservedBits_r16         *uper.BitString         `lb:2,ub:2,optional`
}

func (ie *SL_PreconfigGeneral_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sl_TDD_Configuration_r16 != nil, ie.reservedBits_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.sl_TDD_Configuration_r16 != nil {
		if err = ie.sl_TDD_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sl_TDD_Configuration_r16", err)
		}
	}
	if ie.reservedBits_r16 != nil {
		if err = w.WriteBitString(ie.reservedBits_r16.Bytes, uint(ie.reservedBits_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode reservedBits_r16", err)
		}
	}
	return nil
}

func (ie *SL_PreconfigGeneral_r16) Decode(r *uper.UperReader) error {
	var err error
	var sl_TDD_Configuration_r16Present bool
	if sl_TDD_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var reservedBits_r16Present bool
	if reservedBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if sl_TDD_Configuration_r16Present {
		ie.sl_TDD_Configuration_r16 = new(TDD_UL_DL_ConfigCommon)
		if err = ie.sl_TDD_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sl_TDD_Configuration_r16", err)
		}
	}
	if reservedBits_r16Present {
		var tmp_bs_reservedBits_r16 []byte
		var n_reservedBits_r16 uint
		if tmp_bs_reservedBits_r16, n_reservedBits_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode reservedBits_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_reservedBits_r16,
			NumBits: uint64(n_reservedBits_r16),
		}
		ie.reservedBits_r16 = &tmp_bitstring
	}
	return nil
}
