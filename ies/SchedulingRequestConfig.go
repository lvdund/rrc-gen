package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestConfig struct {
	schedulingRequestToAddModList  []SchedulingRequestToAddMod `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
	schedulingRequestToReleaseList []SchedulingRequestId       `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
}

func (ie *SchedulingRequestConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.schedulingRequestToAddModList) > 0, len(ie.schedulingRequestToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.schedulingRequestToAddModList) > 0 {
		tmp_schedulingRequestToAddModList := utils.NewSequence[*SchedulingRequestToAddMod]([]*SchedulingRequestToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.schedulingRequestToAddModList {
			tmp_schedulingRequestToAddModList.Value = append(tmp_schedulingRequestToAddModList.Value, &i)
		}
		if err = tmp_schedulingRequestToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestToAddModList", err)
		}
	}
	if len(ie.schedulingRequestToReleaseList) > 0 {
		tmp_schedulingRequestToReleaseList := utils.NewSequence[*SchedulingRequestId]([]*SchedulingRequestId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.schedulingRequestToReleaseList {
			tmp_schedulingRequestToReleaseList.Value = append(tmp_schedulingRequestToReleaseList.Value, &i)
		}
		if err = tmp_schedulingRequestToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestToReleaseList", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestConfig) Decode(r *uper.UperReader) error {
	var err error
	var schedulingRequestToAddModListPresent bool
	if schedulingRequestToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var schedulingRequestToReleaseListPresent bool
	if schedulingRequestToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if schedulingRequestToAddModListPresent {
		tmp_schedulingRequestToAddModList := utils.NewSequence[*SchedulingRequestToAddMod]([]*SchedulingRequestToAddMod{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_schedulingRequestToAddModList := func() *SchedulingRequestToAddMod {
			return new(SchedulingRequestToAddMod)
		}
		if err = tmp_schedulingRequestToAddModList.Decode(r, fn_schedulingRequestToAddModList); err != nil {
			return utils.WrapError("Decode schedulingRequestToAddModList", err)
		}
		ie.schedulingRequestToAddModList = []SchedulingRequestToAddMod{}
		for _, i := range tmp_schedulingRequestToAddModList.Value {
			ie.schedulingRequestToAddModList = append(ie.schedulingRequestToAddModList, *i)
		}
	}
	if schedulingRequestToReleaseListPresent {
		tmp_schedulingRequestToReleaseList := utils.NewSequence[*SchedulingRequestId]([]*SchedulingRequestId{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_schedulingRequestToReleaseList := func() *SchedulingRequestId {
			return new(SchedulingRequestId)
		}
		if err = tmp_schedulingRequestToReleaseList.Decode(r, fn_schedulingRequestToReleaseList); err != nil {
			return utils.WrapError("Decode schedulingRequestToReleaseList", err)
		}
		ie.schedulingRequestToReleaseList = []SchedulingRequestId{}
		for _, i := range tmp_schedulingRequestToReleaseList.Value {
			ie.schedulingRequestToReleaseList = append(ie.schedulingRequestToReleaseList, *i)
		}
	}
	return nil
}
