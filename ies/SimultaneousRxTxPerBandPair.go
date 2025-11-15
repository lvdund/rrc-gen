package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SimultaneousRxTxPerBandPair struct {
	Value uper.BitString `lb:3,ub:496,madatory`
}

func (ie *SimultaneousRxTxPerBandPair) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Value.Bytes, uint(ie.Value.NumBits), &uper.Constraint{Lb: 3, Ub: 496}, false); err != nil {
		return utils.WrapError("Encode SimultaneousRxTxPerBandPair", err)
	}
	return nil
}

func (ie *SimultaneousRxTxPerBandPair) Decode(r *uper.UperReader) error {
	var err error
	var v []byte
	var n uint
	if v, n, err = r.ReadBitString(&uper.Constraint{Lb: 3, Ub: 496}, false); err != nil {
		return utils.WrapError("Decode SimultaneousRxTxPerBandPair", err)
	}
	ie.Value = uper.BitString{
		Bytes:   v,
		NumBits: uint64(n),
	}
	return nil
}
