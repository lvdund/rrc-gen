package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDU_SessionID struct {
	Value uint64 `lb:0,ub:255,madatory`
}

func (ie *PDU_SessionID) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("Encode PDU_SessionID", err)
	}
	return nil
}

func (ie *PDU_SessionID) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return utils.WrapError("Decode PDU_SessionID", err)
	}
	ie.Value = uint64(v)
	return nil
}
