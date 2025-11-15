package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReselectionThreshold struct {
	Value uint64 `lb:0,ub:31,madatory`
}

func (ie *ReselectionThreshold) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode ReselectionThreshold", err)
	}
	return nil
}

func (ie *ReselectionThreshold) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode ReselectionThreshold", err)
	}
	ie.Value = uint64(v)
	return nil
}
