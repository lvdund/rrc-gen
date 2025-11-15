package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NID_r16 struct {
	Value uper.BitString `lb:44,ub:44,madatory`
}

func (ie *NID_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 44, Ub: 44}, false); err != nil {
		return utils.WrapError("Encode NID_r16", err)
	}
	return nil
}

func (ie *NID_r16) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 44, Ub: 44}, false); err != nil {
		return utils.WrapError("Decode NID_r16", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
