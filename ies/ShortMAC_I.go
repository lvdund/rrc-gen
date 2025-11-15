package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ShortMAC_I struct {
	Value uper.BitString `lb:16,ub:16,madatory`
}

func (ie *ShortMAC_I) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("Encode ShortMAC_I", err)
	}
	return nil
}

func (ie *ShortMAC_I) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("Decode ShortMAC_I", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
