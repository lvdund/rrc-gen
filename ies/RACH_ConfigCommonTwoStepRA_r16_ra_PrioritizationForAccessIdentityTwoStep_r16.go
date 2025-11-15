package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16 struct {
	ra_Prioritization_r16      RA_Prioritization `madatory`
	ra_PrioritizationForAI_r16 uper.BitString    `lb:2,ub:2,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ra_Prioritization_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ra_Prioritization_r16", err)
	}
	if err = w.WriteBitString(ie.ra_PrioritizationForAI_r16.Bytes, uint(ie.ra_PrioritizationForAI_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteBitString ra_PrioritizationForAI_r16", err)
	}
	return nil
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_ra_PrioritizationForAccessIdentityTwoStep_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ra_Prioritization_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ra_Prioritization_r16", err)
	}
	var tmp_bs_ra_PrioritizationForAI_r16 []byte
	var n_ra_PrioritizationForAI_r16 uint
	if tmp_bs_ra_PrioritizationForAI_r16, n_ra_PrioritizationForAI_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadBitString ra_PrioritizationForAI_r16", err)
	}
	ie.ra_PrioritizationForAI_r16 = uper.BitString{
		Bytes:   tmp_bs_ra_PrioritizationForAI_r16,
		NumBits: uint64(n_ra_PrioritizationForAI_r16),
	}
	return nil
}
