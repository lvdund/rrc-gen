package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_neither uper.Enumerated = 0
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_enable  uper.Enumerated = 1
	PUCCH_ConfigCommon_pucch_GroupHopping_Enum_disable uper.Enumerated = 2
)

type PUCCH_ConfigCommon_pucch_GroupHopping struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *PUCCH_ConfigCommon_pucch_GroupHopping) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode PUCCH_ConfigCommon_pucch_GroupHopping", err)
	}
	return nil
}

func (ie *PUCCH_ConfigCommon_pucch_GroupHopping) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode PUCCH_ConfigCommon_pucch_GroupHopping", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
