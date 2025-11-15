package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ARFCN_ValueNR struct {
	Value uint64 `lb:0,ub:maxNARFCN,madatory`
}

func (ie *ARFCN_ValueNR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNARFCN}, false); err != nil {
		return utils.WrapError("Encode ARFCN_ValueNR", err)
	}
	return nil
}

func (ie *ARFCN_ValueNR) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNARFCN}, false); err != nil {
		return utils.WrapError("Decode ARFCN_ValueNR", err)
	}
	ie.Value = uint64(v)
	return nil
}
