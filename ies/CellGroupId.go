package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellGroupId struct {
	Value uint64 `lb:0,ub:maxSecondaryCellGroups,madatory`
}

func (ie *CellGroupId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxSecondaryCellGroups}, false); err != nil {
		return utils.WrapError("Encode CellGroupId", err)
	}
	return nil
}

func (ie *CellGroupId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxSecondaryCellGroups}, false); err != nil {
		return utils.WrapError("Decode CellGroupId", err)
	}
	ie.Value = uint64(v)
	return nil
}
