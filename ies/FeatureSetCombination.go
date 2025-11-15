package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetCombination struct {
	Value []FeatureSetsPerBand `lb:1,ub:maxSimultaneousBands,madatory`
}

func (ie *FeatureSetCombination) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FeatureSetsPerBand]([]*FeatureSetsPerBand{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureSetCombination", err)
	}
	return nil
}

func (ie *FeatureSetCombination) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FeatureSetsPerBand]([]*FeatureSetsPerBand{}, uper.Constraint{Lb: 1, Ub: maxSimultaneousBands}, false)
	fn := func() *FeatureSetsPerBand {
		return new(FeatureSetsPerBand)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FeatureSetCombination", err)
	}
	ie.Value = []FeatureSetsPerBand{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
