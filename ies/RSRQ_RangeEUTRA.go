package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RSRQ_RangeEUTRA struct {
	Value uint64 `lb:0,ub:34,madatory`
}

func (ie *RSRQ_RangeEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 34}, false); err != nil {
		return utils.WrapError("Encode RSRQ_RangeEUTRA", err)
	}
	return nil
}

func (ie *RSRQ_RangeEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 34}, false); err != nil {
		return utils.WrapError("Decode RSRQ_RangeEUTRA", err)
	}
	ie.Value = uint64(v)
	return nil
}
