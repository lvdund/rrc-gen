package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetCombinationId struct {
	Value uint64 `lb:0,ub:maxFeatureSetCombinations,madatory`
}

func (ie *FeatureSetCombinationId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxFeatureSetCombinations}, false); err != nil {
		return utils.WrapError("Encode FeatureSetCombinationId", err)
	}
	return nil
}

func (ie *FeatureSetCombinationId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxFeatureSetCombinations}, false); err != nil {
		return utils.WrapError("Decode FeatureSetCombinationId", err)
	}
	ie.Value = uint64(v)
	return nil
}
