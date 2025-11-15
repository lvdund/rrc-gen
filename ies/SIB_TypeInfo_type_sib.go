package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB_TypeInfo_type_sib_Enum_sibType2        uper.Enumerated = 0
	SIB_TypeInfo_type_sib_Enum_sibType3        uper.Enumerated = 1
	SIB_TypeInfo_type_sib_Enum_sibType4        uper.Enumerated = 2
	SIB_TypeInfo_type_sib_Enum_sibType5        uper.Enumerated = 3
	SIB_TypeInfo_type_sib_Enum_sibType6        uper.Enumerated = 4
	SIB_TypeInfo_type_sib_Enum_sibType7        uper.Enumerated = 5
	SIB_TypeInfo_type_sib_Enum_sibType8        uper.Enumerated = 6
	SIB_TypeInfo_type_sib_Enum_sibType9        uper.Enumerated = 7
	SIB_TypeInfo_type_sib_Enum_sibType10_v1610 uper.Enumerated = 8
	SIB_TypeInfo_type_sib_Enum_sibType11_v1610 uper.Enumerated = 9
	SIB_TypeInfo_type_sib_Enum_sibType12_v1610 uper.Enumerated = 10
	SIB_TypeInfo_type_sib_Enum_sibType13_v1610 uper.Enumerated = 11
	SIB_TypeInfo_type_sib_Enum_sibType14_v1610 uper.Enumerated = 12
	SIB_TypeInfo_type_sib_Enum_spare3          uper.Enumerated = 13
	SIB_TypeInfo_type_sib_Enum_spare2          uper.Enumerated = 14
	SIB_TypeInfo_type_sib_Enum_spare1          uper.Enumerated = 15
)

type SIB_TypeInfo_type_sib struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB_TypeInfo_type_sib) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB_TypeInfo_type_sib", err)
	}
	return nil
}

func (ie *SIB_TypeInfo_type_sib) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB_TypeInfo_type_sib", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
