package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RNTI_Value struct {
	Value uint64 `lb:0,ub:65535,madatory`
}

func (ie *RNTI_Value) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return utils.WrapError("Encode RNTI_Value", err)
	}
	return nil
}

func (ie *RNTI_Value) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return utils.WrapError("Decode RNTI_Value", err)
	}
	ie.Value = uint64(v)
	return nil
}
