package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s2  uper.Enumerated = 0
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s4  uper.Enumerated = 1
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s5  uper.Enumerated = 2
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s6  uper.Enumerated = 3
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s8  uper.Enumerated = 4
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s10 uper.Enumerated = 5
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s12 uper.Enumerated = 6
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s14 uper.Enumerated = 7
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s16 uper.Enumerated = 8
	DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17_Enum_s20 uper.Enumerated = 9
)

type DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17 struct {
	Value uper.Enumerated `lb:0,ub:9,madatory`
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Encode DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17", err)
	}
	return nil
}

func (ie *DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 9}, false); err != nil {
		return utils.WrapError("Decode DMRS_BundlingPUSCH_Config_r17_pusch_FrequencyHoppingInterval_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
