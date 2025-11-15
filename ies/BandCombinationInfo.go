package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationInfo struct {
	bandCombinationIndex   BandCombinationIndex   `madatory`
	allowedFeatureSetsList []FeatureSetEntryIndex `lb:1,ub:maxFeatureSetsPerBand,madatory`
}

func (ie *BandCombinationInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bandCombinationIndex.Encode(w); err != nil {
		return utils.WrapError("Encode bandCombinationIndex", err)
	}
	tmp_allowedFeatureSetsList := utils.NewSequence[*FeatureSetEntryIndex]([]*FeatureSetEntryIndex{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	for _, i := range ie.allowedFeatureSetsList {
		tmp_allowedFeatureSetsList.Value = append(tmp_allowedFeatureSetsList.Value, &i)
	}
	if err = tmp_allowedFeatureSetsList.Encode(w); err != nil {
		return utils.WrapError("Encode allowedFeatureSetsList", err)
	}
	return nil
}

func (ie *BandCombinationInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bandCombinationIndex.Decode(r); err != nil {
		return utils.WrapError("Decode bandCombinationIndex", err)
	}
	tmp_allowedFeatureSetsList := utils.NewSequence[*FeatureSetEntryIndex]([]*FeatureSetEntryIndex{}, uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false)
	fn_allowedFeatureSetsList := func() *FeatureSetEntryIndex {
		return new(FeatureSetEntryIndex)
	}
	if err = tmp_allowedFeatureSetsList.Decode(r, fn_allowedFeatureSetsList); err != nil {
		return utils.WrapError("Decode allowedFeatureSetsList", err)
	}
	ie.allowedFeatureSetsList = []FeatureSetEntryIndex{}
	for _, i := range tmp_allowedFeatureSetsList.Value {
		ie.allowedFeatureSetsList = append(ie.allowedFeatureSetsList, *i)
	}
	return nil
}
