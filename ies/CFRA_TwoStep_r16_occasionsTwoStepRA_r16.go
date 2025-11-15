package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CFRA_TwoStep_r16_occasionsTwoStepRA_r16 struct {
	rach_ConfigGenericTwoStepRA_r16   RACH_ConfigGenericTwoStepRA_r16                                           `madatory`
	ssb_PerRACH_OccasionTwoStepRA_r16 CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16 `madatory`
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rach_ConfigGenericTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rach_ConfigGenericTwoStepRA_r16", err)
	}
	if err = ie.ssb_PerRACH_OccasionTwoStepRA_r16.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rach_ConfigGenericTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rach_ConfigGenericTwoStepRA_r16", err)
	}
	if err = ie.ssb_PerRACH_OccasionTwoStepRA_r16.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}
