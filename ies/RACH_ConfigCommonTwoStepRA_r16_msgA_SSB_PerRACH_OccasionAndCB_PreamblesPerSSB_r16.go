package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_nothing uint64 = iota
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneEighth
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneFourth
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneHalf
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_one
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_two
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_four
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_eight
	RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_sixteen
)

type RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16 struct {
	Choice    uint64
	oneEighth *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneEighth
	oneFourth *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneFourth
	oneHalf   *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneHalf
	one       *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_one
	two       *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two
	four      int64 `lb:1,ub:16,madatory`
	eight     int64 `lb:1,ub:8,madatory`
	sixteen   int64 `lb:1,ub:4,madatory`
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneEighth:
		if err = ie.oneEighth.Encode(w); err != nil {
			err = utils.WrapError("Encode oneEighth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneFourth:
		if err = ie.oneFourth.Encode(w); err != nil {
			err = utils.WrapError("Encode oneFourth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneHalf:
		if err = ie.oneHalf.Encode(w); err != nil {
			err = utils.WrapError("Encode oneHalf", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_one:
		if err = ie.one.Encode(w); err != nil {
			err = utils.WrapError("Encode one", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_two:
		if err = ie.two.Encode(w); err != nil {
			err = utils.WrapError("Encode two", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_four:
		if err = w.WriteInteger(int64(ie.four), &uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			err = utils.WrapError("Encode four", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_eight:
		if err = w.WriteInteger(int64(ie.eight), &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			err = utils.WrapError("Encode eight", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_sixteen:
		if err = w.WriteInteger(int64(ie.sixteen), &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			err = utils.WrapError("Encode sixteen", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneEighth:
		ie.oneEighth = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneEighth)
		if err = ie.oneEighth.Decode(r); err != nil {
			return utils.WrapError("Decode oneEighth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneFourth:
		ie.oneFourth = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneFourth)
		if err = ie.oneFourth.Decode(r); err != nil {
			return utils.WrapError("Decode oneFourth", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_oneHalf:
		ie.oneHalf = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_oneHalf)
		if err = ie.oneHalf.Decode(r); err != nil {
			return utils.WrapError("Decode oneHalf", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_one:
		ie.one = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_one)
		if err = ie.one.Decode(r); err != nil {
			return utils.WrapError("Decode one", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_two:
		ie.two = new(RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_two)
		if err = ie.two.Decode(r); err != nil {
			return utils.WrapError("Decode two", err)
		}
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_four:
		var tmp_int_four int64
		if tmp_int_four, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode four", err)
		}
		ie.four = tmp_int_four
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_eight:
		var tmp_int_eight int64
		if tmp_int_eight, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode eight", err)
		}
		ie.eight = tmp_int_eight
	case RACH_ConfigCommonTwoStepRA_r16_msgA_SSB_PerRACH_OccasionAndCB_PreamblesPerSSB_r16_Choice_sixteen:
		var tmp_int_sixteen int64
		if tmp_int_sixteen, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
			return utils.WrapError("Decode sixteen", err)
		}
		ie.sixteen = tmp_int_sixteen
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
