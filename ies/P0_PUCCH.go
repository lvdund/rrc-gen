package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type P0_PUCCH struct {
	p0_PUCCH_Id    P0_PUCCH_Id `madatory`
	p0_PUCCH_Value int64       `lb:-16,ub:15,madatory`
}

func (ie *P0_PUCCH) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.p0_PUCCH_Id.Encode(w); err != nil {
		return utils.WrapError("Encode p0_PUCCH_Id", err)
	}
	if err = w.WriteInteger(ie.p0_PUCCH_Value, &uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
		return utils.WrapError("WriteInteger p0_PUCCH_Value", err)
	}
	return nil
}

func (ie *P0_PUCCH) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.p0_PUCCH_Id.Decode(r); err != nil {
		return utils.WrapError("Decode p0_PUCCH_Id", err)
	}
	var tmp_int_p0_PUCCH_Value int64
	if tmp_int_p0_PUCCH_Value, err = r.ReadInteger(&uper.Constraint{Lb: -16, Ub: 15}, false); err != nil {
		return utils.WrapError("ReadInteger p0_PUCCH_Value", err)
	}
	ie.p0_PUCCH_Value = tmp_int_p0_PUCCH_Value
	return nil
}
