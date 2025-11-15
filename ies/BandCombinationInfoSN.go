package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombinationInfoSN struct {
	bandCombinationIndex BandCombinationIndex `madatory`
	requestedFeatureSets FeatureSetEntryIndex `madatory`
}

func (ie *BandCombinationInfoSN) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.bandCombinationIndex.Encode(w); err != nil {
		return utils.WrapError("Encode bandCombinationIndex", err)
	}
	if err = ie.requestedFeatureSets.Encode(w); err != nil {
		return utils.WrapError("Encode requestedFeatureSets", err)
	}
	return nil
}

func (ie *BandCombinationInfoSN) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.bandCombinationIndex.Decode(r); err != nil {
		return utils.WrapError("Decode bandCombinationIndex", err)
	}
	if err = ie.requestedFeatureSets.Decode(r); err != nil {
		return utils.WrapError("Decode requestedFeatureSets", err)
	}
	return nil
}
