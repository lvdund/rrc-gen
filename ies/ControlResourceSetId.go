package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ControlResourceSetId struct {
	Value uint64 `lb:0,ub:maxNrofControlResourceSets_1,madatory`
}

func (ie *ControlResourceSetId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofControlResourceSets_1}, false); err != nil {
		return utils.WrapError("Encode ControlResourceSetId", err)
	}
	return nil
}

func (ie *ControlResourceSetId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofControlResourceSets_1}, false); err != nil {
		return utils.WrapError("Decode ControlResourceSetId", err)
	}
	ie.Value = uint64(v)
	return nil
}
