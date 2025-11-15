package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n2  uper.Enumerated = 0
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n3  uper.Enumerated = 1
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n4  uper.Enumerated = 2
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n6  uper.Enumerated = 3
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n8  uper.Enumerated = 4
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n9  uper.Enumerated = 5
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n10 uper.Enumerated = 6
	PUCCH_ConfigCommon_additionalPRBOffset_r17_Enum_n12 uper.Enumerated = 7
)

type PUCCH_ConfigCommon_additionalPRBOffset_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PUCCH_ConfigCommon_additionalPRBOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ConfigCommon_additionalPRBOffset_r17", err)
	}
	return nil
}

func (ie *PUCCH_ConfigCommon_additionalPRBOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ConfigCommon_additionalPRBOffset_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
