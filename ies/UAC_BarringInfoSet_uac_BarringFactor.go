package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p00 uper.Enumerated = 0
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p05 uper.Enumerated = 1
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p10 uper.Enumerated = 2
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p15 uper.Enumerated = 3
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p20 uper.Enumerated = 4
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p25 uper.Enumerated = 5
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p30 uper.Enumerated = 6
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p40 uper.Enumerated = 7
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p50 uper.Enumerated = 8
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p60 uper.Enumerated = 9
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p70 uper.Enumerated = 10
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p75 uper.Enumerated = 11
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p80 uper.Enumerated = 12
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p85 uper.Enumerated = 13
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p90 uper.Enumerated = 14
	UAC_BarringInfoSet_uac_BarringFactor_Enum_p95 uper.Enumerated = 15
)

type UAC_BarringInfoSet_uac_BarringFactor struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *UAC_BarringInfoSet_uac_BarringFactor) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSet_uac_BarringFactor", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet_uac_BarringFactor) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSet_uac_BarringFactor", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
