package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SI_SchedulingInfo_v1700 struct {
	schedulingInfoList2_r17    []SchedulingInfo2_r17 `lb:1,ub:maxSI_Message,madatory`
	si_RequestConfigRedCap_r17 *SI_RequestConfig     `optional`
}

func (ie *SI_SchedulingInfo_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.si_RequestConfigRedCap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	tmp_schedulingInfoList2_r17 := utils.NewSequence[*SchedulingInfo2_r17]([]*SchedulingInfo2_r17{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	for _, i := range ie.schedulingInfoList2_r17 {
		tmp_schedulingInfoList2_r17.Value = append(tmp_schedulingInfoList2_r17.Value, &i)
	}
	if err = tmp_schedulingInfoList2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode schedulingInfoList2_r17", err)
	}
	if ie.si_RequestConfigRedCap_r17 != nil {
		if err = ie.si_RequestConfigRedCap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode si_RequestConfigRedCap_r17", err)
		}
	}
	return nil
}

func (ie *SI_SchedulingInfo_v1700) Decode(r *uper.UperReader) error {
	var err error
	var si_RequestConfigRedCap_r17Present bool
	if si_RequestConfigRedCap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	tmp_schedulingInfoList2_r17 := utils.NewSequence[*SchedulingInfo2_r17]([]*SchedulingInfo2_r17{}, uper.Constraint{Lb: 1, Ub: maxSI_Message}, false)
	fn_schedulingInfoList2_r17 := func() *SchedulingInfo2_r17 {
		return new(SchedulingInfo2_r17)
	}
	if err = tmp_schedulingInfoList2_r17.Decode(r, fn_schedulingInfoList2_r17); err != nil {
		return utils.WrapError("Decode schedulingInfoList2_r17", err)
	}
	ie.schedulingInfoList2_r17 = []SchedulingInfo2_r17{}
	for _, i := range tmp_schedulingInfoList2_r17.Value {
		ie.schedulingInfoList2_r17 = append(ie.schedulingInfoList2_r17, *i)
	}
	if si_RequestConfigRedCap_r17Present {
		ie.si_RequestConfigRedCap_r17 = new(SI_RequestConfig)
		if err = ie.si_RequestConfigRedCap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode si_RequestConfigRedCap_r17", err)
		}
	}
	return nil
}
