package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ARFCN_ValueEUTRA struct {
	Value uint64 `lb:0,ub:maxEARFCN,madatory`
}

func (ie *ARFCN_ValueEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxEARFCN}, false); err != nil {
		return utils.WrapError("Encode ARFCN_ValueEUTRA", err)
	}
	return nil
}

func (ie *ARFCN_ValueEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxEARFCN}, false); err != nil {
		return utils.WrapError("Decode ARFCN_ValueEUTRA", err)
	}
	ie.Value = uint64(v)
	return nil
}
