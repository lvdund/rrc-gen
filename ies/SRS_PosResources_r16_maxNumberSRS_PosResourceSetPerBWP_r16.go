package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n1  uper.Enumerated = 0
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n2  uper.Enumerated = 1
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n4  uper.Enumerated = 2
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n8  uper.Enumerated = 3
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n12 uper.Enumerated = 4
	SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16_Enum_n16 uper.Enumerated = 5
)

type SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	return nil
}

func (ie *SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode SRS_PosResources_r16_maxNumberSRS_PosResourceSetPerBWP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
