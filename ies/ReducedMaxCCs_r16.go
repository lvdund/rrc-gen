package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReducedMaxCCs_r16 struct {
	reducedCCsDL_r16 int64 `lb:0,ub:31,madatory`
	reducedCCsUL_r16 int64 `lb:0,ub:31,madatory`
}

func (ie *ReducedMaxCCs_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.reducedCCsDL_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger reducedCCsDL_r16", err)
	}
	if err = w.WriteInteger(ie.reducedCCsUL_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("WriteInteger reducedCCsUL_r16", err)
	}
	return nil
}

func (ie *ReducedMaxCCs_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_reducedCCsDL_r16 int64
	if tmp_int_reducedCCsDL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger reducedCCsDL_r16", err)
	}
	ie.reducedCCsDL_r16 = tmp_int_reducedCCsDL_r16
	var tmp_int_reducedCCsUL_r16 int64
	if tmp_int_reducedCCsUL_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("ReadInteger reducedCCsUL_r16", err)
	}
	ie.reducedCCsUL_r16 = tmp_int_reducedCCsUL_r16
	return nil
}
