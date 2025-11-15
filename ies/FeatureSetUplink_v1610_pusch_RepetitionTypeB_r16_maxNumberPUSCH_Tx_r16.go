package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n2  uper.Enumerated = 0
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n3  uper.Enumerated = 1
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n4  uper.Enumerated = 2
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n7  uper.Enumerated = 3
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n8  uper.Enumerated = 4
	FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16_Enum_n12 uper.Enumerated = 5
)

type FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16 struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
