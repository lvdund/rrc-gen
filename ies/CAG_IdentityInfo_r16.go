package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CAG_IdentityInfo_r16 struct {
	cag_Identity_r16              uper.BitString                                      `lb:32,ub:32,madatory`
	manualCAGselectionAllowed_r16 *CAG_IdentityInfo_r16_manualCAGselectionAllowed_r16 `optional`
}

func (ie *CAG_IdentityInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.manualCAGselectionAllowed_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.cag_Identity_r16.Bytes, uint(ie.cag_Identity_r16.NumBits), &uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
		return utils.WrapError("WriteBitString cag_Identity_r16", err)
	}
	if ie.manualCAGselectionAllowed_r16 != nil {
		if err = ie.manualCAGselectionAllowed_r16.Encode(w); err != nil {
			return utils.WrapError("Encode manualCAGselectionAllowed_r16", err)
		}
	}
	return nil
}

func (ie *CAG_IdentityInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var manualCAGselectionAllowed_r16Present bool
	if manualCAGselectionAllowed_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_cag_Identity_r16 []byte
	var n_cag_Identity_r16 uint
	if tmp_bs_cag_Identity_r16, n_cag_Identity_r16, err = r.ReadBitString(&uper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
		return utils.WrapError("ReadBitString cag_Identity_r16", err)
	}
	ie.cag_Identity_r16 = uper.BitString{
		Bytes:   tmp_bs_cag_Identity_r16,
		NumBits: uint64(n_cag_Identity_r16),
	}
	if manualCAGselectionAllowed_r16Present {
		ie.manualCAGselectionAllowed_r16 = new(CAG_IdentityInfo_r16_manualCAGselectionAllowed_r16)
		if err = ie.manualCAGselectionAllowed_r16.Decode(r); err != nil {
			return utils.WrapError("Decode manualCAGselectionAllowed_r16", err)
		}
	}
	return nil
}
