package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB17_IEs_r17_validityDuration_r17_Enum_t1       uper.Enumerated = 0
	SIB17_IEs_r17_validityDuration_r17_Enum_t2       uper.Enumerated = 1
	SIB17_IEs_r17_validityDuration_r17_Enum_t4       uper.Enumerated = 2
	SIB17_IEs_r17_validityDuration_r17_Enum_t8       uper.Enumerated = 3
	SIB17_IEs_r17_validityDuration_r17_Enum_t16      uper.Enumerated = 4
	SIB17_IEs_r17_validityDuration_r17_Enum_t32      uper.Enumerated = 5
	SIB17_IEs_r17_validityDuration_r17_Enum_t64      uper.Enumerated = 6
	SIB17_IEs_r17_validityDuration_r17_Enum_t128     uper.Enumerated = 7
	SIB17_IEs_r17_validityDuration_r17_Enum_t256     uper.Enumerated = 8
	SIB17_IEs_r17_validityDuration_r17_Enum_t512     uper.Enumerated = 9
	SIB17_IEs_r17_validityDuration_r17_Enum_infinity uper.Enumerated = 10
	SIB17_IEs_r17_validityDuration_r17_Enum_spare5   uper.Enumerated = 11
	SIB17_IEs_r17_validityDuration_r17_Enum_spare4   uper.Enumerated = 12
	SIB17_IEs_r17_validityDuration_r17_Enum_spare3   uper.Enumerated = 13
	SIB17_IEs_r17_validityDuration_r17_Enum_spare2   uper.Enumerated = 14
	SIB17_IEs_r17_validityDuration_r17_Enum_spare1   uper.Enumerated = 15
)

type SIB17_IEs_r17_validityDuration_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *SIB17_IEs_r17_validityDuration_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode SIB17_IEs_r17_validityDuration_r17", err)
	}
	return nil
}

func (ie *SIB17_IEs_r17_validityDuration_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode SIB17_IEs_r17_validityDuration_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
