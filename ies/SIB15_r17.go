package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB15_r17 struct {
	commonPLMNsWithDisasterCondition_r17 []PLMN_Identity              `lb:1,ub:maxPLMN,optional`
	applicableDisasterInfoList_r17       []ApplicableDisasterInfo_r17 `lb:1,ub:maxPLMN,optional`
	lateNonCriticalExtension             *[]byte                      `optional`
}

func (ie *SIB15_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.commonPLMNsWithDisasterCondition_r17) > 0, len(ie.applicableDisasterInfoList_r17) > 0, ie.lateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.commonPLMNsWithDisasterCondition_r17) > 0 {
		tmp_commonPLMNsWithDisasterCondition_r17 := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.commonPLMNsWithDisasterCondition_r17 {
			tmp_commonPLMNsWithDisasterCondition_r17.Value = append(tmp_commonPLMNsWithDisasterCondition_r17.Value, &i)
		}
		if err = tmp_commonPLMNsWithDisasterCondition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode commonPLMNsWithDisasterCondition_r17", err)
		}
	}
	if len(ie.applicableDisasterInfoList_r17) > 0 {
		tmp_applicableDisasterInfoList_r17 := utils.NewSequence[*ApplicableDisasterInfo_r17]([]*ApplicableDisasterInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.applicableDisasterInfoList_r17 {
			tmp_applicableDisasterInfoList_r17.Value = append(tmp_applicableDisasterInfoList_r17.Value, &i)
		}
		if err = tmp_applicableDisasterInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode applicableDisasterInfoList_r17", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB15_r17) Decode(r *uper.UperReader) error {
	var err error
	var commonPLMNsWithDisasterCondition_r17Present bool
	if commonPLMNsWithDisasterCondition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var applicableDisasterInfoList_r17Present bool
	if applicableDisasterInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if commonPLMNsWithDisasterCondition_r17Present {
		tmp_commonPLMNsWithDisasterCondition_r17 := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_commonPLMNsWithDisasterCondition_r17 := func() *PLMN_Identity {
			return new(PLMN_Identity)
		}
		if err = tmp_commonPLMNsWithDisasterCondition_r17.Decode(r, fn_commonPLMNsWithDisasterCondition_r17); err != nil {
			return utils.WrapError("Decode commonPLMNsWithDisasterCondition_r17", err)
		}
		ie.commonPLMNsWithDisasterCondition_r17 = []PLMN_Identity{}
		for _, i := range tmp_commonPLMNsWithDisasterCondition_r17.Value {
			ie.commonPLMNsWithDisasterCondition_r17 = append(ie.commonPLMNsWithDisasterCondition_r17, *i)
		}
	}
	if applicableDisasterInfoList_r17Present {
		tmp_applicableDisasterInfoList_r17 := utils.NewSequence[*ApplicableDisasterInfo_r17]([]*ApplicableDisasterInfo_r17{}, uper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_applicableDisasterInfoList_r17 := func() *ApplicableDisasterInfo_r17 {
			return new(ApplicableDisasterInfo_r17)
		}
		if err = tmp_applicableDisasterInfoList_r17.Decode(r, fn_applicableDisasterInfoList_r17); err != nil {
			return utils.WrapError("Decode applicableDisasterInfoList_r17", err)
		}
		ie.applicableDisasterInfoList_r17 = []ApplicableDisasterInfo_r17{}
		for _, i := range tmp_applicableDisasterInfoList_r17.Value {
			ie.applicableDisasterInfoList_r17 = append(ie.applicableDisasterInfoList_r17, *i)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	return nil
}
