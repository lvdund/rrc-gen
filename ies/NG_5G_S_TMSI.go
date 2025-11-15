package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NG_5G_S_TMSI struct {
	Value uper.BitString `lb:48,ub:48,madatory`
}

func (ie *NG_5G_S_TMSI) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("Encode NG_5G_S_TMSI", err)
	}
	return nil
}

func (ie *NG_5G_S_TMSI) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 48, Ub: 48}, false); err != nil {
		return utils.WrapError("Decode NG_5G_S_TMSI", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
