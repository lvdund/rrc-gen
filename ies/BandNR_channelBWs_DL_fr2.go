package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_channelBWs_DL_fr2 struct {
	scs_60kHz  *uper.BitString `lb:3,ub:3,optional`
	scs_120kHz *uper.BitString `lb:3,ub:3,optional`
}

func (ie *BandNR_channelBWs_DL_fr2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_60kHz != nil, ie.scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_60kHz != nil {
		if err = w.WriteBitString(ie.scs_60kHz.Bytes, uint(ie.scs_60kHz.NumBits), &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode scs_60kHz", err)
		}
	}
	if ie.scs_120kHz != nil {
		if err = w.WriteBitString(ie.scs_120kHz.Bytes, uint(ie.scs_120kHz.NumBits), &uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Encode scs_120kHz", err)
		}
	}
	return nil
}

func (ie *BandNR_channelBWs_DL_fr2) Decode(r *uper.UperReader) error {
	var err error
	var scs_60kHzPresent bool
	if scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_120kHzPresent bool
	if scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_60kHzPresent {
		var tmp_bs_scs_60kHz []byte
		var n_scs_60kHz uint
		if tmp_bs_scs_60kHz, n_scs_60kHz, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_60kHz,
			NumBits: uint64(n_scs_60kHz),
		}
		ie.scs_60kHz = &tmp_bitstring
	}
	if scs_120kHzPresent {
		var tmp_bs_scs_120kHz []byte
		var n_scs_120kHz uint
		if tmp_bs_scs_120kHz, n_scs_120kHz, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode scs_120kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_scs_120kHz,
			NumBits: uint64(n_scs_120kHz),
		}
		ie.scs_120kHz = &tmp_bitstring
	}
	return nil
}
