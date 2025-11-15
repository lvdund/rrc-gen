package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasGapSharingScheme_Enum_scheme00 uper.Enumerated = 0
	MeasGapSharingScheme_Enum_scheme01 uper.Enumerated = 1
	MeasGapSharingScheme_Enum_scheme10 uper.Enumerated = 2
	MeasGapSharingScheme_Enum_scheme11 uper.Enumerated = 3
)

type MeasGapSharingScheme struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *MeasGapSharingScheme) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode MeasGapSharingScheme", err)
	}
	return nil
}

func (ie *MeasGapSharingScheme) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode MeasGapSharingScheme", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
