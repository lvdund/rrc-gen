package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PortIndex2 struct {
	Value uint64 `lb:0,ub:1,madatory`
}

func (ie *PortIndex2) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode PortIndex2", err)
	}
	return nil
}

func (ie *PortIndex2) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode PortIndex2", err)
	}
	ie.Value = uint64(v)
	return nil
}
