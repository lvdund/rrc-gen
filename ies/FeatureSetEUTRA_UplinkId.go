package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetEUTRA_UplinkId struct {
	Value uint64 `lb:0,ub:maxEUTRA_UL_FeatureSets,madatory`
}

func (ie *FeatureSetEUTRA_UplinkId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxEUTRA_UL_FeatureSets}, false); err != nil {
		return utils.WrapError("Encode FeatureSetEUTRA_UplinkId", err)
	}
	return nil
}

func (ie *FeatureSetEUTRA_UplinkId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxEUTRA_UL_FeatureSets}, false); err != nil {
		return utils.WrapError("Decode FeatureSetEUTRA_UplinkId", err)
	}
	ie.Value = uint64(v)
	return nil
}
