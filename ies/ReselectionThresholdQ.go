package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReselectionThresholdQ struct {
	Value uint64 `lb:0,ub:31,madatory`
}

func (ie *ReselectionThresholdQ) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode ReselectionThresholdQ", err)
	}
	return nil
}

func (ie *ReselectionThresholdQ) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode ReselectionThresholdQ", err)
	}
	ie.Value = uint64(v)
	return nil
}
