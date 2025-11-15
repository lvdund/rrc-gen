package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16_Enum_kHz1dot25 uper.Enumerated = 0
	RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16_Enum_kHz5      uper.Enumerated = 1
	RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16_Enum_spare2    uper.Enumerated = 2
	RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16_Enum_spare1    uper.Enumerated = 3
)

type RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16 struct {
	Value uper.Enumerated `lb:0,ub:3,madatory`
}

func (ie *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Encode RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
	}
	return nil
}

func (ie *RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
		return utils.WrapError("Decode RA_InformationCommon_r16_msg1_SCS_From_prach_ConfigurationIndexCFRA_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
