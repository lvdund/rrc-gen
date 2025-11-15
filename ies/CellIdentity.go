package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellIdentity struct {
	Value uper.BitString `lb:36,ub:36,madatory`
}

func (ie *CellIdentity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return utils.WrapError("Encode CellIdentity", err)
	}
	return nil
}

func (ie *CellIdentity) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 36, Ub: 36}, false); err != nil {
		return utils.WrapError("Decode CellIdentity", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
