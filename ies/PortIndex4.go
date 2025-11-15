package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndex4 struct {
	Value uint64 `lb:0,ub:3,madatory`
}

func (ie *PortIndex4) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode PortIndex4", err)
	}
	return nil
}

func (ie *PortIndex4) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode PortIndex4", err)
	}
	ie.Value = uint64(v)
	return nil
}
