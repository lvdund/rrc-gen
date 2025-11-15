package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReferenceTime_r16 struct {
	refDays_r16           int64 `lb:0,ub:72999,madatory`
	refSeconds_r16        int64 `lb:0,ub:86399,madatory`
	refMilliSeconds_r16   int64 `lb:0,ub:999,madatory`
	refTenNanoSeconds_r16 int64 `lb:0,ub:99999,madatory`
}

func (ie *ReferenceTime_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.refDays_r16, &uper.Constraint{Lb: 0, Ub: 72999}, false); err != nil {
		return utils.WrapError("WriteInteger refDays_r16", err)
	}
	if err = w.WriteInteger(ie.refSeconds_r16, &uper.Constraint{Lb: 0, Ub: 86399}, false); err != nil {
		return utils.WrapError("WriteInteger refSeconds_r16", err)
	}
	if err = w.WriteInteger(ie.refMilliSeconds_r16, &uper.Constraint{Lb: 0, Ub: 999}, false); err != nil {
		return utils.WrapError("WriteInteger refMilliSeconds_r16", err)
	}
	if err = w.WriteInteger(ie.refTenNanoSeconds_r16, &uper.Constraint{Lb: 0, Ub: 99999}, false); err != nil {
		return utils.WrapError("WriteInteger refTenNanoSeconds_r16", err)
	}
	return nil
}

func (ie *ReferenceTime_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_refDays_r16 int64
	if tmp_int_refDays_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 72999}, false); err != nil {
		return utils.WrapError("ReadInteger refDays_r16", err)
	}
	ie.refDays_r16 = tmp_int_refDays_r16
	var tmp_int_refSeconds_r16 int64
	if tmp_int_refSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 86399}, false); err != nil {
		return utils.WrapError("ReadInteger refSeconds_r16", err)
	}
	ie.refSeconds_r16 = tmp_int_refSeconds_r16
	var tmp_int_refMilliSeconds_r16 int64
	if tmp_int_refMilliSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 999}, false); err != nil {
		return utils.WrapError("ReadInteger refMilliSeconds_r16", err)
	}
	ie.refMilliSeconds_r16 = tmp_int_refMilliSeconds_r16
	var tmp_int_refTenNanoSeconds_r16 int64
	if tmp_int_refTenNanoSeconds_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 99999}, false); err != nil {
		return utils.WrapError("ReadInteger refTenNanoSeconds_r16", err)
	}
	ie.refTenNanoSeconds_r16 = tmp_int_refTenNanoSeconds_r16
	return nil
}
