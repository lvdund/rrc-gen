package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneEighth uper.Enumerated = 0
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneFourth uper.Enumerated = 1
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_oneHalf   uper.Enumerated = 2
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_one       uper.Enumerated = 3
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_two       uper.Enumerated = 4
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_four      uper.Enumerated = 5
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_eight     uper.Enumerated = 6
	CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16_Enum_sixteen   uper.Enumerated = 7
)

type CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	return nil
}

func (ie *CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode CFRA_TwoStep_r16_occasionsTwoStepRA_r16_ssb_PerRACH_OccasionTwoStepRA_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
