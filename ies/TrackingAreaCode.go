package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TrackingAreaCode struct {
	Value uper.BitString `lb:24,ub:24,madatory`
}

func (ie *TrackingAreaCode) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Encode TrackingAreaCode", err)
	}
	return nil
}

func (ie *TrackingAreaCode) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
		return utils.WrapError("Decode TrackingAreaCode", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
