package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LTE_CRS_PatternList_r16 struct {
	Value []RateMatchPatternLTE_CRS `lb:1,ub:maxLTE_CRS_Patterns_r16,madatory`
}

func (ie *LTE_CRS_PatternList_r16) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*RateMatchPatternLTE_CRS]([]*RateMatchPatternLTE_CRS{}, uper.Constraint{Lb: 1, Ub: maxLTE_CRS_Patterns_r16}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode LTE_CRS_PatternList_r16", err)
	}
	return nil
}

func (ie *LTE_CRS_PatternList_r16) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*RateMatchPatternLTE_CRS]([]*RateMatchPatternLTE_CRS{}, uper.Constraint{Lb: 1, Ub: maxLTE_CRS_Patterns_r16}, false)
	fn := func() *RateMatchPatternLTE_CRS {
		return new(RateMatchPatternLTE_CRS)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode LTE_CRS_PatternList_r16", err)
	}
	ie.Value = []RateMatchPatternLTE_CRS{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
