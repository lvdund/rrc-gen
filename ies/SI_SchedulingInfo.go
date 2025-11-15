package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_SchedulingInfo struct {
	schedulingInfoList      []SchedulingInfo                  `lb:1,ub:maxSI_Message,madatory`
	si_WindowLength         SI_SchedulingInfo_si_WindowLength `madatory`
	si_RequestConfig        *SI_RequestConfig                 `optional`
	si_RequestConfigSUL     *SI_RequestConfig                 `optional`
	systemInformationAreaID *uper.BitString                   `lb:24,ub:24,optional`
}

func (ie *SI_SchedulingInfo) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.si_RequestConfig != nil, ie.si_RequestConfigSUL != nil, ie.systemInformationAreaID != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_schedulingInfoList := utils.NewSequence[*SchedulingInfo]([]*SchedulingInfo{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.schedulingInfoList {
		tmp_schedulingInfoList.Value = append(tmp_schedulingInfoList.Value, &i)
	}
	if err = tmp_schedulingInfoList.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingInfoList", err)
	}
	if err = ie.si_WindowLength.Encode(w); err != nil {
		return utils.WrapError("Encode si_WindowLength", err)
	}
	if ie.si_RequestConfig != nil {
		if err = ie.si_RequestConfig.Encode(w); err != nil {
			return utils.WrapError("Encode si_RequestConfig", err)
		}
	}
	if ie.si_RequestConfigSUL != nil {
		if err = ie.si_RequestConfigSUL.Encode(w); err != nil {
			return utils.WrapError("Encode si_RequestConfigSUL", err)
		}
	}
	if ie.systemInformationAreaID != nil {
		if err = w.WriteBitString(ie.systemInformationAreaID.Bytes, uint(ie.systemInformationAreaID.NumBits), &uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Encode systemInformationAreaID", err)
		}
	}
	return nil
}

func (ie *SI_SchedulingInfo) Decode(r *uper.UperReader) error {
	var err error
	var si_RequestConfigPresent bool
	if si_RequestConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var si_RequestConfigSULPresent bool
	if si_RequestConfigSULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var systemInformationAreaIDPresent bool
	if systemInformationAreaIDPresent, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_schedulingInfoList := utils.NewSequence[*SchedulingInfo]([]*SchedulingInfo{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_schedulingInfoList := func() *SchedulingInfo {
		return new(SchedulingInfo)
	}
	if err = tmp_schedulingInfoList.Decode(r, fn_schedulingInfoList); err != nil {
		return utils.WrapError("Decode schedulingInfoList", err)
	}
	ie.schedulingInfoList = []SchedulingInfo{}
	for _, i := range tmp_schedulingInfoList.Value {
		ie.schedulingInfoList = append(ie.schedulingInfoList, *i)
	}
	if err = ie.si_WindowLength.Decode(r); err != nil {
		return utils.WrapError("Decode si_WindowLength", err)
	}
	if si_RequestConfigPresent {
		ie.si_RequestConfig = new(SI_RequestConfig)
		if err = ie.si_RequestConfig.Decode(r); err != nil {
			return utils.WrapError("Decode si_RequestConfig", err)
		}
	}
	if si_RequestConfigSULPresent {
		ie.si_RequestConfigSUL = new(SI_RequestConfig)
		if err = ie.si_RequestConfigSUL.Decode(r); err != nil {
			return utils.WrapError("Decode si_RequestConfigSUL", err)
		}
	}
	if systemInformationAreaIDPresent {
		var tmp_bs_systemInformationAreaID []byte
		var n_systemInformationAreaID uint
		if tmp_bs_systemInformationAreaID, n_systemInformationAreaID, err = r.ReadBitString(&uper.Constraint{Lb: 24, Ub: 24}, false); err != nil {
			return utils.WrapError("Decode systemInformationAreaID", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_systemInformationAreaID,
			NumBits: uint64(n_systemInformationAreaID),
		}
		ie.systemInformationAreaID = &tmp_bitstring
	}
	return nil
}
