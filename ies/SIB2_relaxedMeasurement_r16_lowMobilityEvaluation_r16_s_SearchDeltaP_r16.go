package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB3    uper.Enumerated = 0
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB6    uper.Enumerated = 1
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB9    uper.Enumerated = 2
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB12   uper.Enumerated = 3
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_dB15   uper.Enumerated = 4
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare3 uper.Enumerated = 5
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare2 uper.Enumerated = 6
	SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16_Enum_spare1 uper.Enumerated = 7
)

type SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16", err)
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16_s_SearchDeltaP_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
