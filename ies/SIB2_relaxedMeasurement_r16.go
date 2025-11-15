package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r16 struct {
	lowMobilityEvaluation_r16       *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16       `optional`
	cellEdgeEvaluation_r16          *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16          `optional`
	combineRelaxedMeasCondition_r16 *SIB2_relaxedMeasurement_r16_combineRelaxedMeasCondition_r16 `optional`
	highPriorityMeasRelax_r16       *SIB2_relaxedMeasurement_r16_highPriorityMeasRelax_r16       `optional`
}

func (ie *SIB2_relaxedMeasurement_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.lowMobilityEvaluation_r16 != nil, ie.cellEdgeEvaluation_r16 != nil, ie.combineRelaxedMeasCondition_r16 != nil, ie.highPriorityMeasRelax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.lowMobilityEvaluation_r16 != nil {
		if err = ie.lowMobilityEvaluation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode lowMobilityEvaluation_r16", err)
		}
	}
	if ie.cellEdgeEvaluation_r16 != nil {
		if err = ie.cellEdgeEvaluation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode cellEdgeEvaluation_r16", err)
		}
	}
	if ie.combineRelaxedMeasCondition_r16 != nil {
		if err = ie.combineRelaxedMeasCondition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode combineRelaxedMeasCondition_r16", err)
		}
	}
	if ie.highPriorityMeasRelax_r16 != nil {
		if err = ie.highPriorityMeasRelax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode highPriorityMeasRelax_r16", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16) Decode(r *uper.UperReader) error {
	var err error
	var lowMobilityEvaluation_r16Present bool
	if lowMobilityEvaluation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var cellEdgeEvaluation_r16Present bool
	if cellEdgeEvaluation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var combineRelaxedMeasCondition_r16Present bool
	if combineRelaxedMeasCondition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highPriorityMeasRelax_r16Present bool
	if highPriorityMeasRelax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if lowMobilityEvaluation_r16Present {
		ie.lowMobilityEvaluation_r16 = new(SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16)
		if err = ie.lowMobilityEvaluation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode lowMobilityEvaluation_r16", err)
		}
	}
	if cellEdgeEvaluation_r16Present {
		ie.cellEdgeEvaluation_r16 = new(SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16)
		if err = ie.cellEdgeEvaluation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cellEdgeEvaluation_r16", err)
		}
	}
	if combineRelaxedMeasCondition_r16Present {
		ie.combineRelaxedMeasCondition_r16 = new(SIB2_relaxedMeasurement_r16_combineRelaxedMeasCondition_r16)
		if err = ie.combineRelaxedMeasCondition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode combineRelaxedMeasCondition_r16", err)
		}
	}
	if highPriorityMeasRelax_r16Present {
		ie.highPriorityMeasRelax_r16 = new(SIB2_relaxedMeasurement_r16_highPriorityMeasRelax_r16)
		if err = ie.highPriorityMeasRelax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode highPriorityMeasRelax_r16", err)
		}
	}
	return nil
}
