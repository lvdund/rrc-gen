package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PCCH_Config_nAndPagingFrameOffset_Choice_nothing uint64 = iota
	PCCH_Config_nAndPagingFrameOffset_Choice_oneT
	PCCH_Config_nAndPagingFrameOffset_Choice_halfT
	PCCH_Config_nAndPagingFrameOffset_Choice_quarterT
	PCCH_Config_nAndPagingFrameOffset_Choice_oneEighthT
	PCCH_Config_nAndPagingFrameOffset_Choice_oneSixteenthT
)

type PCCH_Config_nAndPagingFrameOffset struct {
	Choice        uint64
	oneT          uper.NULL `madatory`
	halfT         int64     `lb:0,ub:1,madatory`
	quarterT      int64     `lb:0,ub:3,madatory`
	oneEighthT    int64     `lb:0,ub:7,madatory`
	oneSixteenthT int64     `lb:0,ub:15,madatory`
}

func (ie *PCCH_Config_nAndPagingFrameOffset) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneT:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode oneT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_halfT:
		if err = w.WriteInteger(int64(ie.halfT), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode halfT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_quarterT:
		if err = w.WriteInteger(int64(ie.quarterT), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode quarterT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneEighthT:
		if err = w.WriteInteger(int64(ie.oneEighthT), &uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			err = utils.WrapError("Encode oneEighthT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneSixteenthT:
		if err = w.WriteInteger(int64(ie.oneSixteenthT), &uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			err = utils.WrapError("Encode oneSixteenthT", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PCCH_Config_nAndPagingFrameOffset) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(5, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneT:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode oneT", err)
		}
	case PCCH_Config_nAndPagingFrameOffset_Choice_halfT:
		var tmp_int_halfT int64
		if tmp_int_halfT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode halfT", err)
		}
		ie.halfT = tmp_int_halfT
	case PCCH_Config_nAndPagingFrameOffset_Choice_quarterT:
		var tmp_int_quarterT int64
		if tmp_int_quarterT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode quarterT", err)
		}
		ie.quarterT = tmp_int_quarterT
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneEighthT:
		var tmp_int_oneEighthT int64
		if tmp_int_oneEighthT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
			return utils.WrapError("Decode oneEighthT", err)
		}
		ie.oneEighthT = tmp_int_oneEighthT
	case PCCH_Config_nAndPagingFrameOffset_Choice_oneSixteenthT:
		var tmp_int_oneSixteenthT int64
		if tmp_int_oneSixteenthT, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode oneSixteenthT", err)
		}
		ie.oneSixteenthT = tmp_int_oneSixteenthT
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
