package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetEUTRA_DownlinkId struct {
	Value uint64 `lb:0,ub:maxEUTRA_DL_FeatureSets,madatory`
}

func (ie *FeatureSetEUTRA_DownlinkId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxEUTRA_DL_FeatureSets}, false); err != nil {
		return utils.WrapError("Encode FeatureSetEUTRA_DownlinkId", err)
	}
	return nil
}

func (ie *FeatureSetEUTRA_DownlinkId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxEUTRA_DL_FeatureSets}, false); err != nil {
		return utils.WrapError("Decode FeatureSetEUTRA_DownlinkId", err)
	}
	ie.Value = uint64(v)
	return nil
}
