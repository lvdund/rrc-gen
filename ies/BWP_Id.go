package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BWP_Id struct {
	Value uint64 `lb:0,ub:maxNrofBWPs,madatory`
}

func (ie *BWP_Id) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofBWPs}, false); err != nil {
		return utils.WrapError("Encode BWP_Id", err)
	}
	return nil
}

func (ie *BWP_Id) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofBWPs}, false); err != nil {
		return utils.WrapError("Decode BWP_Id", err)
	}
	ie.Value = uint64(v)
	return nil
}
