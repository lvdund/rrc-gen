package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Fr1_r17 struct {
	scs_15kHz_r17 *uper.BitString `lb:16,ub:16,optional`
	scs_30kHz_r17 *uper.BitString `lb:16,ub:16,optional`
	scs_60kHz_r17 *uper.BitString `lb:16,ub:16,optional`
}

func (ie *Fr1_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz_r17 != nil, ie.scs_30kHz_r17 != nil, ie.scs_60kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs_15kHz_r17.Bytes, uint(ie.scs_15kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_15kHz_r17", err)
		}
	}
	if ie.scs_30kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs_30kHz_r17.Bytes, uint(ie.scs_30kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_30kHz_r17", err)
		}
	}
	if ie.scs_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.scs_60kHz_r17.Bytes, uint(ie.scs_60kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_60kHz_r17", err)
		}
	}
	return nil
}

func (ie *Fr1_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHz_r17Present bool
	if scs_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHz_r17Present bool
	if scs_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_60kHz_r17Present bool
	if scs_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHz_r17Present {
		var tmp_bs_scs_15kHz_r17 []byte
		var n_scs_15kHz_r17 uint
		if tmp_bs_scs_15kHz_r17, n_scs_15kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_15kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_15kHz_r17,
			NumBits: uint64(n_scs_15kHz_r17),
		}
		ie.scs_15kHz_r17 = &tmp_bitstring
	}
	if scs_30kHz_r17Present {
		var tmp_bs_scs_30kHz_r17 []byte
		var n_scs_30kHz_r17 uint
		if tmp_bs_scs_30kHz_r17, n_scs_30kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_30kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_30kHz_r17,
			NumBits: uint64(n_scs_30kHz_r17),
		}
		ie.scs_30kHz_r17 = &tmp_bitstring
	}
	if scs_60kHz_r17Present {
		var tmp_bs_scs_60kHz_r17 []byte
		var n_scs_60kHz_r17 uint
		if tmp_bs_scs_60kHz_r17, n_scs_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_60kHz_r17,
			NumBits: uint64(n_scs_60kHz_r17),
		}
		ie.scs_60kHz_r17 = &tmp_bitstring
	}
	return nil
}
