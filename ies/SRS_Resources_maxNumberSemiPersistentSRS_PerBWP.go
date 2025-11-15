package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resources_maxNumberSemiPersistentSRS_PerBWP_Enum_n1  uper.Enumerated = 0
	SRS_Resources_maxNumberSemiPersistentSRS_PerBWP_Enum_n2  uper.Enumerated = 1
	SRS_Resources_maxNumberSemiPersistentSRS_PerBWP_Enum_n4  uper.Enumerated = 2
	SRS_Resources_maxNumberSemiPersistentSRS_PerBWP_Enum_n8  uper.Enumerated = 3
	SRS_Resources_maxNumberSemiPersistentSRS_PerBWP_Enum_n16 uper.Enumerated = 4
)

type SRS_Resources_maxNumberSemiPersistentSRS_PerBWP struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SRS_Resources_maxNumberSemiPersistentSRS_PerBWP) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SRS_Resources_maxNumberSemiPersistentSRS_PerBWP", err)
	}
	return nil
}

func (ie *SRS_Resources_maxNumberSemiPersistentSRS_PerBWP) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SRS_Resources_maxNumberSemiPersistentSRS_PerBWP", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
