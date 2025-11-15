package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Fr1_r16 struct {
	scs_15kHz_r16 *uper.BitString `lb:16,ub:16,optional`
	scs_30kHz_r16 *uper.BitString `lb:16,ub:16,optional`
	scs_60kHz_r16 *uper.BitString `lb:16,ub:16,optional`
}

func (ie *Fr1_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz_r16 != nil, ie.scs_30kHz_r16 != nil, ie.scs_60kHz_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz_r16 != nil {
		if err = w.WriteBitString(ie.scs_15kHz_r16.Bytes, uint(ie.scs_15kHz_r16.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_15kHz_r16", err)
		}
	}
	if ie.scs_30kHz_r16 != nil {
		if err = w.WriteBitString(ie.scs_30kHz_r16.Bytes, uint(ie.scs_30kHz_r16.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_30kHz_r16", err)
		}
	}
	if ie.scs_60kHz_r16 != nil {
		if err = w.WriteBitString(ie.scs_60kHz_r16.Bytes, uint(ie.scs_60kHz_r16.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_60kHz_r16", err)
		}
	}
	return nil
}

func (ie *Fr1_r16) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHz_r16Present bool
	if scs_15kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHz_r16Present bool
	if scs_30kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_60kHz_r16Present bool
	if scs_60kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHz_r16Present {
		var tmp_bs_scs_15kHz_r16 []byte
		var n_scs_15kHz_r16 uint
		if tmp_bs_scs_15kHz_r16, n_scs_15kHz_r16, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_15kHz_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_15kHz_r16,
			NumBits: uint64(n_scs_15kHz_r16),
		}
		ie.scs_15kHz_r16 = &tmp_bitstring
	}
	if scs_30kHz_r16Present {
		var tmp_bs_scs_30kHz_r16 []byte
		var n_scs_30kHz_r16 uint
		if tmp_bs_scs_30kHz_r16, n_scs_30kHz_r16, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_30kHz_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_30kHz_r16,
			NumBits: uint64(n_scs_30kHz_r16),
		}
		ie.scs_30kHz_r16 = &tmp_bitstring
	}
	if scs_60kHz_r16Present {
		var tmp_bs_scs_60kHz_r16 []byte
		var n_scs_60kHz_r16 uint
		if tmp_bs_scs_60kHz_r16, n_scs_60kHz_r16, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_60kHz_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_60kHz_r16,
			NumBits: uint64(n_scs_60kHz_r16),
		}
		ie.scs_60kHz_r16 = &tmp_bitstring
	}
	return nil
}
