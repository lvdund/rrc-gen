package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SpCellConfig_lowMobilityEvaluationConnected_r17 struct {
	s_SearchDeltaP_Connected_r17 SpCellConfig_lowMobilityEvaluationConnected_r17_s_SearchDeltaP_Connected_r17 `madatory`
	t_SearchDeltaP_Connected_r17 SpCellConfig_lowMobilityEvaluationConnected_r17_t_SearchDeltaP_Connected_r17 `madatory`
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.s_SearchDeltaP_Connected_r17.Encode(w); err != nil {
		return utils.WrapError("Encode s_SearchDeltaP_Connected_r17", err)
	}
	if err = ie.t_SearchDeltaP_Connected_r17.Encode(w); err != nil {
		return utils.WrapError("Encode t_SearchDeltaP_Connected_r17", err)
	}
	return nil
}

func (ie *SpCellConfig_lowMobilityEvaluationConnected_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.s_SearchDeltaP_Connected_r17.Decode(r); err != nil {
		return utils.WrapError("Decode s_SearchDeltaP_Connected_r17", err)
	}
	if err = ie.t_SearchDeltaP_Connected_r17.Decode(r); err != nil {
		return utils.WrapError("Decode t_SearchDeltaP_Connected_r17", err)
	}
	return nil
}
