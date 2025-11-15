package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSetIndex struct {
	Value uint64 `lb:1,ub:maxBarringInfoSet,madatory`
}

func (ie *UAC_BarringInfoSetIndex) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSetIndex", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSetIndex) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxBarringInfoSet}, false); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSetIndex", err)
	}
	ie.Value = uint64(v)
	return nil
}
