package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetEntryIndex struct {
	Value uint64 `lb:1,ub:maxFeatureSetsPerBand,madatory`
}

func (ie *FeatureSetEntryIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false); err != nil {
		return utils.WrapError("Encode FeatureSetEntryIndex", err)
	}
	return nil
}

func (ie *FeatureSetEntryIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxFeatureSetsPerBand}, false); err != nil {
		return utils.WrapError("Decode FeatureSetEntryIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
