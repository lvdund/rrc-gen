package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_uac_BarringInfo struct {
	uac_BarringForCommon                        *UAC_BarringPerCatList                                            `optional`
	uac_BarringPerPLMN_List                     *UAC_BarringPerPLMN_List                                          `optional`
	uac_BarringInfoSetList                      UAC_BarringInfoSetList                                            `madatory`
	uac_AccessCategory1_SelectionAssistanceInfo *SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo `lb:2,ub:maxPLMN,optional`
}

func (ie *SIB1_uac_BarringInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.uac_BarringForCommon != nil, ie.uac_BarringPerPLMN_List != nil, ie.uac_AccessCategory1_SelectionAssistanceInfo != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.uac_BarringForCommon != nil {
		if err = ie.uac_BarringForCommon.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringForCommon", err)
		}
	}
	if ie.uac_BarringPerPLMN_List != nil {
		if err = ie.uac_BarringPerPLMN_List.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringPerPLMN_List", err)
		}
	}
	if err = ie.uac_BarringInfoSetList.Encode(w); err != nil {
		return utils.WrapError("Encode uac_BarringInfoSetList", err)
	}
	if ie.uac_AccessCategory1_SelectionAssistanceInfo != nil {
		if err = ie.uac_AccessCategory1_SelectionAssistanceInfo.Encode(w); err != nil {
			return utils.WrapError("Encode uac_AccessCategory1_SelectionAssistanceInfo", err)
		}
	}
	return nil
}

func (ie *SIB1_uac_BarringInfo) Decode(r *uper.UperReader) error {
	var err error
	var uac_BarringForCommonPresent bool
	if uac_BarringForCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uac_BarringPerPLMN_ListPresent bool
	if uac_BarringPerPLMN_ListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uac_AccessCategory1_SelectionAssistanceInfoPresent bool
	if uac_AccessCategory1_SelectionAssistanceInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if uac_BarringForCommonPresent {
		ie.uac_BarringForCommon = new(UAC_BarringPerCatList)
		if err = ie.uac_BarringForCommon.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringForCommon", err)
		}
	}
	if uac_BarringPerPLMN_ListPresent {
		ie.uac_BarringPerPLMN_List = new(UAC_BarringPerPLMN_List)
		if err = ie.uac_BarringPerPLMN_List.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringPerPLMN_List", err)
		}
	}
	if err = ie.uac_BarringInfoSetList.Decode(r); err != nil {
		return utils.WrapError("Decode uac_BarringInfoSetList", err)
	}
	if uac_AccessCategory1_SelectionAssistanceInfoPresent {
		ie.uac_AccessCategory1_SelectionAssistanceInfo = new(SIB1_uac_BarringInfo_uac_AccessCategory1_SelectionAssistanceInfo)
		if err = ie.uac_AccessCategory1_SelectionAssistanceInfo.Decode(r); err != nil {
			return utils.WrapError("Decode uac_AccessCategory1_SelectionAssistanceInfo", err)
		}
	}
	return nil
}
