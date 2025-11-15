package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n2_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n4_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n6_r17
	NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n12_r17
)

type NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17 struct {
	Choice  uint64
	n2_r17  int64 `lb:0,ub:1,madatory`
	n4_r17  int64 `lb:0,ub:3,madatory`
	n6_r17  int64 `lb:0,ub:5,madatory`
	n12_r17 int64 `lb:0,ub:11,madatory`
}

func (ie *NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n2_r17:
		if err = w.WriteInteger(int64(ie.n2_r17), &uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			err = utils.WrapError("Encode n2_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n4_r17:
		if err = w.WriteInteger(int64(ie.n4_r17), &uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			err = utils.WrapError("Encode n4_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n6_r17:
		if err = w.WriteInteger(int64(ie.n6_r17), &uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
			err = utils.WrapError("Encode n6_r17", err)
		}
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n12_r17:
		if err = w.WriteInteger(int64(ie.n12_r17), &uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
			err = utils.WrapError("Encode n12_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n2_r17:
		var tmp_int_n2_r17 int64
		if tmp_int_n2_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
			return utils.WrapError("Decode n2_r17", err)
		}
		ie.n2_r17 = tmp_int_n2_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n4_r17:
		var tmp_int_n4_r17 int64
		if tmp_int_n4_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 3}, false); err != nil {
			return utils.WrapError("Decode n4_r17", err)
		}
		ie.n4_r17 = tmp_int_n4_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n6_r17:
		var tmp_int_n6_r17 int64
		if tmp_int_n6_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
			return utils.WrapError("Decode n6_r17", err)
		}
		ie.n6_r17 = tmp_int_n6_r17
	case NR_DL_PRS_Resource_r17_dl_PRS_CombSizeN_AndReOffset_r17_Choice_n12_r17:
		var tmp_int_n12_r17 int64
		if tmp_int_n12_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 11}, false); err != nil {
			return utils.WrapError("Decode n12_r17", err)
		}
		ie.n12_r17 = tmp_int_n12_r17
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
