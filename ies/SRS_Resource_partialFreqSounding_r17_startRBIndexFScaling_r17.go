package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_nothing uint64 = iota
	SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor2_r17
	SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor4_r17
)

type SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17 struct {
	Choice                                uint64
	startRBIndexAndFreqScalingFactor2_r17 int64 `lb:0,ub:1,madatory`
	startRBIndexAndFreqScalingFactor4_r17 int64 `lb:0,ub:3,madatory`
}

func (ie *SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor2_r17:
		if err = w.WriteInteger(int64(ie.startRBIndexAndFreqScalingFactor2_r17), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode startRBIndexAndFreqScalingFactor2_r17", err)
		}
	case SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor4_r17:
		if err = w.WriteInteger(int64(ie.startRBIndexAndFreqScalingFactor4_r17), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode startRBIndexAndFreqScalingFactor4_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor2_r17:
		var tmp_int_startRBIndexAndFreqScalingFactor2_r17 int64
		if tmp_int_startRBIndexAndFreqScalingFactor2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode startRBIndexAndFreqScalingFactor2_r17", err)
		}
		ie.startRBIndexAndFreqScalingFactor2_r17 = tmp_int_startRBIndexAndFreqScalingFactor2_r17
	case SRS_Resource_partialFreqSounding_r17_startRBIndexFScaling_r17_Choice_startRBIndexAndFreqScalingFactor4_r17:
		var tmp_int_startRBIndexAndFreqScalingFactor4_r17 int64
		if tmp_int_startRBIndexAndFreqScalingFactor4_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode startRBIndexAndFreqScalingFactor4_r17", err)
		}
		ie.startRBIndexAndFreqScalingFactor4_r17 = tmp_int_startRBIndexAndFreqScalingFactor4_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
