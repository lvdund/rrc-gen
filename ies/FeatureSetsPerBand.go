package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetsPerBand struct {
	Value []FeatureSet `lb:1,ub:maxFeatureSetsPerBand,madatory`
}

func (ie *FeatureSetsPerBand) Encode(w *uper.UperWriter) error {
	var err error
	tmp := utils.NewSequence[*FeatureSet]([]*FeatureSet{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	for _, i := range ie.Value {
		tmp.Value = append(tmp.Value, &i)
	}
	if err = tmp.Encode(w); err != nil {
		return utils.WrapError("Encode FeatureSetsPerBand", err)
	}
	return nil
}

func (ie *FeatureSetsPerBand) Decode(r *uper.UperReader) error {
	var err error
	tmp := utils.NewSequence[*FeatureSet]([]*FeatureSet{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	fn := func() *FeatureSet {
		return new(FeatureSet)
	}
	if err = tmp.Decode(r, fn); err != nil {
		return utils.WrapError("Decode FeatureSetsPerBand", err)
	}
	ie.Value = []FeatureSet{}
	for _, i := range tmp.Value {
		ie.Value = append(ie.Value, *i)
	}
	return err
}
