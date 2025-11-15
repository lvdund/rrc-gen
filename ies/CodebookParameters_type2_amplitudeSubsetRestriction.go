package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CodebookParameters_type2_amplitudeSubsetRestriction_Enum_supported uper.Enumerated = 0
)

type CodebookParameters_type2_amplitudeSubsetRestriction struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *CodebookParameters_type2_amplitudeSubsetRestriction) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode CodebookParameters_type2_amplitudeSubsetRestriction", err)
	}
	return nil
}

func (ie *CodebookParameters_type2_amplitudeSubsetRestriction) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode CodebookParameters_type2_amplitudeSubsetRestriction", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
