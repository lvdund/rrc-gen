package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_channelBWs_DL_v1590_fr1 struct {
	scs_15kHz *uper.BitString `lb:16,ub:16,optional`
	scs_30kHz *uper.BitString `lb:16,ub:16,optional`
	scs_60kHz *uper.BitString `lb:16,ub:16,optional`
}

func (ie *BandNR_channelBWs_DL_v1590_fr1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz != nil, ie.scs_30kHz != nil, ie.scs_60kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz != nil {
		if err = w.WriteBitString(ie.scs_15kHz.Bytes, uint(ie.scs_15kHz.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_15kHz", err)
		}
	}
	if ie.scs_30kHz != nil {
		if err = w.WriteBitString(ie.scs_30kHz.Bytes, uint(ie.scs_30kHz.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_30kHz", err)
		}
	}
	if ie.scs_60kHz != nil {
		if err = w.WriteBitString(ie.scs_60kHz.Bytes, uint(ie.scs_60kHz.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode scs_60kHz", err)
		}
	}
	return nil
}

func (ie *BandNR_channelBWs_DL_v1590_fr1) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHzPresent bool
	if scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHzPresent bool
	if scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_60kHzPresent bool
	if scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHzPresent {
		var tmp_bs_scs_15kHz []byte
		var n_scs_15kHz uint
		if tmp_bs_scs_15kHz, n_scs_15kHz, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_15kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_15kHz,
			NumBits: uint64(n_scs_15kHz),
		}
		ie.scs_15kHz = &tmp_bitstring
	}
	if scs_30kHzPresent {
		var tmp_bs_scs_30kHz []byte
		var n_scs_30kHz uint
		if tmp_bs_scs_30kHz, n_scs_30kHz, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_30kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_30kHz,
			NumBits: uint64(n_scs_30kHz),
		}
		ie.scs_30kHz = &tmp_bitstring
	}
	if scs_60kHzPresent {
		var tmp_bs_scs_60kHz []byte
		var n_scs_60kHz uint
		if tmp_bs_scs_60kHz, n_scs_60kHz, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_60kHz,
			NumBits: uint64(n_scs_60kHz),
		}
		ie.scs_60kHz = &tmp_bitstring
	}
	return nil
}
