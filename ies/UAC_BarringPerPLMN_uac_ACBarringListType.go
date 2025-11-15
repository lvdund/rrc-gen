package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_nothing uint64 = iota
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ImplicitACBarringList
	UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ExplicitACBarringList
)

type UAC_BarringPerPLMN_uac_ACBarringListType struct {
	Choice                    uint64
	uac_ImplicitACBarringList []UAC_BarringInfoSetIndex `lb:maxAccessCat_1,ub:maxAccessCat_1,madatory`
	uac_ExplicitACBarringList *UAC_BarringPerCatList
}

func (ie *UAC_BarringPerPLMN_uac_ACBarringListType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ImplicitACBarringList:
		tmp := utils.NewSequence[*UAC_BarringInfoSetIndex]([]*UAC_BarringInfoSetIndex{}, uper.Constraint{Lb: maxAccessCat_1, Ub: maxAccessCat_1}, false)
		for _, i := range ie.uac_ImplicitACBarringList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode uac_ImplicitACBarringList", err)
		}
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ExplicitACBarringList:
		if err = ie.uac_ExplicitACBarringList.Encode(w); err != nil {
			err = utils.WrapError("Encode uac_ExplicitACBarringList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *UAC_BarringPerPLMN_uac_ACBarringListType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ImplicitACBarringList:
		tmp := utils.NewSequence[*UAC_BarringInfoSetIndex]([]*UAC_BarringInfoSetIndex{}, uper.Constraint{Lb: maxAccessCat_1, Ub: maxAccessCat_1}, false)
		fn := func() *UAC_BarringInfoSetIndex {
			return new(UAC_BarringInfoSetIndex)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode uac_ImplicitACBarringList", err)
		}
		ie.uac_ImplicitACBarringList = []UAC_BarringInfoSetIndex{}
		for _, i := range tmp.Value {
			ie.uac_ImplicitACBarringList = append(ie.uac_ImplicitACBarringList, *i)
		}
	case UAC_BarringPerPLMN_uac_ACBarringListType_Choice_uac_ExplicitACBarringList:
		ie.uac_ExplicitACBarringList = new(UAC_BarringPerCatList)
		if err = ie.uac_ExplicitACBarringList.Decode(r); err != nil {
			return utils.WrapError("Decode uac_ExplicitACBarringList", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
