package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r17 struct {
	stationaryMobilityEvaluation_r17      SIB2_relaxedMeasurement_r17_stationaryMobilityEvaluation_r17       `madatory`
	cellEdgeEvaluationWhileStationary_r17 *SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17 `optional`
	combineRelaxedMeasCondition2_r17      *SIB2_relaxedMeasurement_r17_combineRelaxedMeasCondition2_r17      `optional`
}

func (ie *SIB2_relaxedMeasurement_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellEdgeEvaluationWhileStationary_r17 != nil, ie.combineRelaxedMeasCondition2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.stationaryMobilityEvaluation_r17.Encode(w); err != nil {
		return utils.WrapError("Encode stationaryMobilityEvaluation_r17", err)
	}
	if ie.cellEdgeEvaluationWhileStationary_r17 != nil {
		if err = ie.cellEdgeEvaluationWhileStationary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode cellEdgeEvaluationWhileStationary_r17", err)
		}
	}
	if ie.combineRelaxedMeasCondition2_r17 != nil {
		if err = ie.combineRelaxedMeasCondition2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode combineRelaxedMeasCondition2_r17", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r17) Decode(r *uper.UperReader) error {
	var err error
	var cellEdgeEvaluationWhileStationary_r17Present bool
	if cellEdgeEvaluationWhileStationary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var combineRelaxedMeasCondition2_r17Present bool
	if combineRelaxedMeasCondition2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.stationaryMobilityEvaluation_r17.Decode(r); err != nil {
		return utils.WrapError("Decode stationaryMobilityEvaluation_r17", err)
	}
	if cellEdgeEvaluationWhileStationary_r17Present {
		ie.cellEdgeEvaluationWhileStationary_r17 = new(SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17)
		if err = ie.cellEdgeEvaluationWhileStationary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode cellEdgeEvaluationWhileStationary_r17", err)
		}
	}
	if combineRelaxedMeasCondition2_r17Present {
		ie.combineRelaxedMeasCondition2_r17 = new(SIB2_relaxedMeasurement_r17_combineRelaxedMeasCondition2_r17)
		if err = ie.combineRelaxedMeasCondition2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode combineRelaxedMeasCondition2_r17", err)
		}
	}
	return nil
}
