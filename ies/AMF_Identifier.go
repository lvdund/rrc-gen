package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AMF_Identifier struct {
	Value uper.BitString `lb:24,ub:24,madatory`
}

func (ie *AMF_Identifier) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Encode AMF_Identifier", err)
	}
	return nil
}

func (ie *AMF_Identifier) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Decode AMF_Identifier", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
