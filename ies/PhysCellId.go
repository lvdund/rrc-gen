package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PhysCellId struct {
	Value uint64 `lb:0,ub:1007,madatory`
}

func (ie *PhysCellId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	return nil
}

func (ie *PhysCellId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1007}, false); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	ie.Value = uint64(v)
	return nil
}
