package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16 struct {
	maxNumberPUSCH_Tx_r16 FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_maxNumberPUSCH_Tx_r16 `madatory`
	hoppingScheme_r16     FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16_hoppingScheme_r16     `madatory`
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.maxNumberPUSCH_Tx_r16.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberPUSCH_Tx_r16", err)
	}
	if err = ie.hoppingScheme_r16.Encode(w); err != nil {
		return utils.WrapError("Encode hoppingScheme_r16", err)
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_pusch_RepetitionTypeB_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.maxNumberPUSCH_Tx_r16.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberPUSCH_Tx_r16", err)
	}
	if err = ie.hoppingScheme_r16.Decode(r); err != nil {
		return utils.WrapError("Decode hoppingScheme_r16", err)
	}
	return nil
}
