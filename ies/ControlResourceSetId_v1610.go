package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSetId_v1610 struct {
	Value uint64 `lb:maxNrofControlResourceSets,ub:maxNrofControlResourceSets_1_r16,madatory`
}

func (ie *ControlResourceSetId_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: maxNrofControlResourceSets, Ub: maxNrofControlResourceSets_1_r16}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSetId_v1610", err)
	}
	return nil
}

func (ie *ControlResourceSetId_v1610) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: maxNrofControlResourceSets, Ub: maxNrofControlResourceSets_1_r16}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSetId_v1610", err)
	}
	ie.Value = uint64(v)
	return nil
}
