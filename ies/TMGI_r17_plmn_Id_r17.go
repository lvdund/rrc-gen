package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	TMGI_r17_plmn_Id_r17_Choice_nothing uint64 = iota
	TMGI_r17_plmn_Id_r17_Choice_plmn_Index
	TMGI_r17_plmn_Id_r17_Choice_explicitValue
)

type TMGI_r17_plmn_Id_r17 struct {
	Choice        uint64
	plmn_Index    int64 `lb:1,ub:maxPLMN,madatory`
	explicitValue *PLMN_Identity
}

func (ie *TMGI_r17_plmn_Id_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TMGI_r17_plmn_Id_r17_Choice_plmn_Index:
		if err = w.WriteInteger(int64(ie.plmn_Index), &uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			err = utils.WrapError("Encode plmn_Index", err)
		}
	case TMGI_r17_plmn_Id_r17_Choice_explicitValue:
		if err = ie.explicitValue.Encode(w); err != nil {
			err = utils.WrapError("Encode explicitValue", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *TMGI_r17_plmn_Id_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case TMGI_r17_plmn_Id_r17_Choice_plmn_Index:
		var tmp_int_plmn_Index int64
		if tmp_int_plmn_Index, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxPLMN}, false); err != nil {
			return utils.WrapError("Decode plmn_Index", err)
		}
		ie.plmn_Index = tmp_int_plmn_Index
	case TMGI_r17_plmn_Id_r17_Choice_explicitValue:
		ie.explicitValue = new(PLMN_Identity)
		if err = ie.explicitValue.Decode(r); err != nil {
			return utils.WrapError("Decode explicitValue", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
