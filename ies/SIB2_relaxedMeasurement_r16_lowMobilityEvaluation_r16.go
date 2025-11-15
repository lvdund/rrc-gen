package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16 struct {
	s_SearchDeltaP_r16 SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16 `madatory`
	t_SearchDeltaP_r16 SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_t_SearchDeltaP_r16 `madatory`
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.s_SearchDeltaP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode s_SearchDeltaP_r16", err)
	}
	if err = ie.t_SearchDeltaP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode t_SearchDeltaP_r16", err)
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.s_SearchDeltaP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode s_SearchDeltaP_r16", err)
	}
	if err = ie.t_SearchDeltaP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode t_SearchDeltaP_r16", err)
	}
	return nil
}
