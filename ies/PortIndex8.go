package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndex8 struct {
	Value uint64 `lb:0,ub:7,madatory`
}

func (ie *PortIndex8) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PortIndex8", err)
	}
	return nil
}

func (ie *PortIndex8) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PortIndex8", err)
	}
	ie.Value = uint64(v)
	return nil
}
