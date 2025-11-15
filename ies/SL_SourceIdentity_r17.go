package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_SourceIdentity_r17 struct {
	Value uper.BitString `lb:24,ub:24,madatory`
}

func (ie *SL_SourceIdentity_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Encode SL_SourceIdentity_r17", err)
	}
	return nil
}

func (ie *SL_SourceIdentity_r17) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Decode SL_SourceIdentity_r17", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
