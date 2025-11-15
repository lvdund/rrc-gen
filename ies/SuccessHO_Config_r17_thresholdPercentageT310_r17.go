package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_p40    uper.Enumerated = 0
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_p60    uper.Enumerated = 1
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_p80    uper.Enumerated = 2
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_spare5 uper.Enumerated = 3
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_spare4 uper.Enumerated = 4
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_spare3 uper.Enumerated = 5
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_spare2 uper.Enumerated = 6
	SuccessHO_Config_r17_thresholdPercentageT310_r17_Enum_spare1 uper.Enumerated = 7
)

type SuccessHO_Config_r17_thresholdPercentageT310_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SuccessHO_Config_r17_thresholdPercentageT310_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SuccessHO_Config_r17_thresholdPercentageT310_r17", err)
	}
	return nil
}

func (ie *SuccessHO_Config_r17_thresholdPercentageT310_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SuccessHO_Config_r17_thresholdPercentageT310_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
