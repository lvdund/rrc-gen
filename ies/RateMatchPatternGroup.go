package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RateMatchPatternGroup struct {
	Value []RateMatchPatternGroupItem `lb:1,ub:maxNrofRateMatchPatternsPerGroup,madatory`
}

func (ie *RateMatchPatternGroup) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*RateMatchPatternGroupItem]([]*RateMatchPatternGroupItem{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatternsPerGroup}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode RateMatchPatternGroup", err)
	}
	return nil
}

func (ie *RateMatchPatternGroup) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*RateMatchPatternGroupItem]([]*RateMatchPatternGroupItem{}, uper.Constraint{Lb: 1, Ub: maxNrofRateMatchPatternsPerGroup}, false)
	fn := func() *RateMatchPatternGroupItem {
		return new(RateMatchPatternGroupItem)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode RateMatchPatternGroup", err)
	}
	ie.Value = []RateMatchPatternGroupItem{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
