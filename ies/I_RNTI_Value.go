package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type I_RNTI_Value struct {
	Value uper.BitString `lb:40,ub:40,madatory`
}

func (ie *I_RNTI_Value) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
		return utils.WrapError("Encode I_RNTI_Value", err)
	}
	return nil
}

func (ie *I_RNTI_Value) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 40, Ub: 40}, false); err != nil {
		return utils.WrapError("Decode I_RNTI_Value", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
