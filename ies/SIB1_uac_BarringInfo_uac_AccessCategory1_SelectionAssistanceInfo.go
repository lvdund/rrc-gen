package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_nothing uint64 = iota
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_plmnCommon
	SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_individualPLMNList
)

type SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo struct {
	Choice             uint64
	plmnCommon         *UAC_AccessCategory1_SelectionAssistanceInfo
	individualPLMNList []UAC_AccessCategory1_SelectionAssistanceInfo `lb:2,ub:maxPLMN,madatory`
}

func (ie *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_plmnCommon:
		if err = ie.plmnCommon.Encode(w); err != nil {
			err = utils.WrapError("Encode plmnCommon", err)
		}
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_individualPLMNList:
		tmp := utils.NewSequence[*UAC_AccessCategory1_SelectionAssistanceInfo]([]*UAC_AccessCategory1_SelectionAssistanceInfo{}, uper.Constraint{Lb: 2, Ub: maxPLMN}, false)
		for _, i := range ie.individualPLMNList {
			tmp.Value = append(tmp.Value, &i)
		}
		if err = tmp.Encode(w); err != nil {
			err = utils.WrapError("Encode individualPLMNList", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_plmnCommon:
		ie.plmnCommon = new(UAC_AccessCategory1_SelectionAssistanceInfo)
		if err = ie.plmnCommon.Decode(r); err != nil {
			return utils.WrapError("Decode plmnCommon", err)
		}
	case SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo_Choice_individualPLMNList:
		tmp := utils.NewSequence[*UAC_AccessCategory1_SelectionAssistanceInfo]([]*UAC_AccessCategory1_SelectionAssistanceInfo{}, uper.Constraint{Lb: 2, Ub: maxPLMN}, false)
		fn := func() *UAC_AccessCategory1_SelectionAssistanceInfo {
			return new(UAC_AccessCategory1_SelectionAssistanceInfo)
		}
		if err = tmp.Decode(r, fn); err != nil {
			return utils.WrapError("Decode individualPLMNList", err)
		}
		ie.individualPLMNList = []UAC_AccessCategory1_SelectionAssistanceInfo{}
		for _, i := range tmp.Value {
			ie.individualPLMNList = append(ie.individualPLMNList, *i)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
