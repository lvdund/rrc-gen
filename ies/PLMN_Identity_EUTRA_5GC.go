package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PLMN_Identity_EUTRA_5GC_Choice_nothing uint64 = iota
	PLMN_Identity_EUTRA_5GC_Choice_plmn_Identity_EUTRA_5GC
	PLMN_Identity_EUTRA_5GC_Choice_plmn_index
)

type PLMN_Identity_EUTRA_5GC struct {
	Choice                  uint64
	plmn_Identity_EUTRA_5GC *PLMN_Identity
	plmn_index              int64 `lb:1,ub:maxPLMN,madatory`
}

func (ie *PLMN_Identity_EUTRA_5GC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PLMN_Identity_EUTRA_5GC_Choice_plmn_Identity_EUTRA_5GC:
		if err = ie.plmn_Identity_EUTRA_5GC.Encode(w); err != nil {
			err = utils.WrapError("Encode plmn_Identity_EUTRA_5GC", err)
		}
	case PLMN_Identity_EUTRA_5GC_Choice_plmn_index:
		if err = w.WriteInteger(int64(ie.plmn_index), &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			err = utils.WrapError("Encode plmn_index", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PLMN_Identity_EUTRA_5GC) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PLMN_Identity_EUTRA_5GC_Choice_plmn_Identity_EUTRA_5GC:
		ie.plmn_Identity_EUTRA_5GC = new(PLMN_Identity)
		if err = ie.plmn_Identity_EUTRA_5GC.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_Identity_EUTRA_5GC", err)
		}
	case PLMN_Identity_EUTRA_5GC_Choice_plmn_index:
		var tmp_int_plmn_index int64
		if tmp_int_plmn_index, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode plmn_index", err)
		}
		ie.plmn_index = tmp_int_plmn_index
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
