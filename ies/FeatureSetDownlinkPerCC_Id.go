package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_Id struct {
	Value uint64 `lb:1,ub:maxPerCC_FeatureSets,madatory`
}

func (ie *FeatureSetDownlinkPerCC_Id) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false); err != nil {
		return utils.WrapError("Encode FeatureSetDownlinkPerCC_Id", err)
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_Id) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPerCC_FeatureSets}, false); err != nil {
		return utils.WrapError("Decode FeatureSetDownlinkPerCC_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
