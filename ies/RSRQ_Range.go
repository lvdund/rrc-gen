package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RSRQ_Range struct {
	Value uint64 `lb:0,ub:127,madatory`
}

func (ie *RSRQ_Range) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Encode RSRQ_Range", err)
	}
	return nil
}

func (ie *RSRQ_Range) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Decode RSRQ_Range", err)
	}
	ie.Value = uint64(v)
	return nil
}
