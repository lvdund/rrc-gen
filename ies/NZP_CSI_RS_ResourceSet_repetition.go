package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NZP_CSI_RS_ResourceSet_repetition_Enum_on  uper.Enumerated = 0
	NZP_CSI_RS_ResourceSet_repetition_Enum_off uper.Enumerated = 1
)

type NZP_CSI_RS_ResourceSet_repetition struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *NZP_CSI_RS_ResourceSet_repetition) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode NZP_CSI_RS_ResourceSet_repetition", err)
	}
	return nil
}

func (ie *NZP_CSI_RS_ResourceSet_repetition) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode NZP_CSI_RS_ResourceSet_repetition", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
