package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AdditionalSpectrumEmission struct {
	Value uint64 `lb:0,ub:7,madatory`
}

func (ie *AdditionalSpectrumEmission) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode AdditionalSpectrumEmission", err)
	}
	return nil
}

func (ie *AdditionalSpectrumEmission) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode AdditionalSpectrumEmission", err)
	}
	ie.Value = uint64(v)
	return nil
}
