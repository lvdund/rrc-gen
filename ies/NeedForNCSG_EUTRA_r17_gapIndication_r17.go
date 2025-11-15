package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NeedForNCSG_EUTRA_r17_gapIndication_r17_Enum_gap          uper.Enumerated = 0
	NeedForNCSG_EUTRA_r17_gapIndication_r17_Enum_ncsg         uper.Enumerated = 1
	NeedForNCSG_EUTRA_r17_gapIndication_r17_Enum_nogap_noncsg uper.Enumerated = 2
)

type NeedForNCSG_EUTRA_r17_gapIndication_r17 struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *NeedForNCSG_EUTRA_r17_gapIndication_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode NeedForNCSG_EUTRA_r17_gapIndication_r17", err)
	}
	return nil
}

func (ie *NeedForNCSG_EUTRA_r17_gapIndication_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode NeedForNCSG_EUTRA_r17_gapIndication_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
