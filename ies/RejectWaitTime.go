package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RejectWaitTime struct {
	Value uint64 `lb:1,ub:16,madatory`
}

func (ie *RejectWaitTime) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode RejectWaitTime", err)
	}
	return nil
}

func (ie *RejectWaitTime) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode RejectWaitTime", err)
	}
	ie.Value = uint64(v)
	return nil
}
