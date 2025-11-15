package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SINR_RangeEUTRA struct {
	Value uint64 `lb:0,ub:127,madatory`
}

func (ie *SINR_RangeEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Encode SINR_RangeEUTRA", err)
	}
	return nil
}

func (ie *SINR_RangeEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 127}, false); err != nil {
		return utils.WrapError("Decode SINR_RangeEUTRA", err)
	}
	ie.Value = uint64(v)
	return nil
}
