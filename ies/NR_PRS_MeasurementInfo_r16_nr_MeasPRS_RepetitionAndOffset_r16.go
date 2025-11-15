package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_nothing uint64 = iota
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms20_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms40_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms80_r16
	NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms160_r16
)

type NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16 struct {
	Choice    uint64
	ms20_r16  int64 `lb:0,ub:19,madatory`
	ms40_r16  int64 `lb:0,ub:39,madatory`
	ms80_r16  int64 `lb:0,ub:79,madatory`
	ms160_r16 int64 `lb:0,ub:159,madatory`
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms20_r16:
		if err = w.WriteInteger(int64(ie.ms20_r16), &uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			err = utils.WrapError("Encode ms20_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms40_r16:
		if err = w.WriteInteger(int64(ie.ms40_r16), &uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			err = utils.WrapError("Encode ms40_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms80_r16:
		if err = w.WriteInteger(int64(ie.ms80_r16), &uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			err = utils.WrapError("Encode ms80_r16", err)
		}
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms160_r16:
		if err = w.WriteInteger(int64(ie.ms160_r16), &uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			err = utils.WrapError("Encode ms160_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms20_r16:
		var tmp_int_ms20_r16 int64
		if tmp_int_ms20_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 19}, false); err != nil {
			return utils.WrapError("Decode ms20_r16", err)
		}
		ie.ms20_r16 = tmp_int_ms20_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms40_r16:
		var tmp_int_ms40_r16 int64
		if tmp_int_ms40_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 39}, false); err != nil {
			return utils.WrapError("Decode ms40_r16", err)
		}
		ie.ms40_r16 = tmp_int_ms40_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms80_r16:
		var tmp_int_ms80_r16 int64
		if tmp_int_ms80_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 79}, false); err != nil {
			return utils.WrapError("Decode ms80_r16", err)
		}
		ie.ms80_r16 = tmp_int_ms80_r16
	case NR_PRS_MeasurementInfo_r16_nr_MeasPRS_RepetitionAndOffset_r16_Choice_ms160_r16:
		var tmp_int_ms160_r16 int64
		if tmp_int_ms160_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 159}, false); err != nil {
			return utils.WrapError("Decode ms160_r16", err)
		}
		ie.ms160_r16 = tmp_int_ms160_r16
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
