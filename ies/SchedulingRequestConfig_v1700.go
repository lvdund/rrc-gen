package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingRequestConfig_v1700 struct {
	schedulingRequestToAddModListExt_v1700 []SchedulingRequestToAddModExt_v1700 `lb:1,ub:maxNrofSR_ConfigPerCellGroup,optional`
}

func (ie *SchedulingRequestConfig_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.schedulingRequestToAddModListExt_v1700) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.schedulingRequestToAddModListExt_v1700) > 0 {
		tmp_schedulingRequestToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestToAddModExt_v1700]([]*SchedulingRequestToAddModExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		for _, i := range ie.schedulingRequestToAddModListExt_v1700 {
			tmp_schedulingRequestToAddModListExt_v1700.Value = append(tmp_schedulingRequestToAddModListExt_v1700.Value, &i)
		}
		if err = tmp_schedulingRequestToAddModListExt_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode schedulingRequestToAddModListExt_v1700", err)
		}
	}
	return nil
}

func (ie *SchedulingRequestConfig_v1700) Decode(r *uper.UperReader) error {
	var err error
	var schedulingRequestToAddModListExt_v1700Present bool
	if schedulingRequestToAddModListExt_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if schedulingRequestToAddModListExt_v1700Present {
		tmp_schedulingRequestToAddModListExt_v1700 := utils.NewSequence[*SchedulingRequestToAddModExt_v1700]([]*SchedulingRequestToAddModExt_v1700{}, uper.Constraint{Lb: 1, Ub: maxNrofSR_ConfigPerCellGroup}, false)
		fn_schedulingRequestToAddModListExt_v1700 := func() *SchedulingRequestToAddModExt_v1700 {
			return new(SchedulingRequestToAddModExt_v1700)
		}
		if err = tmp_schedulingRequestToAddModListExt_v1700.Decode(r, fn_schedulingRequestToAddModListExt_v1700); err != nil {
			return utils.WrapError("Decode schedulingRequestToAddModListExt_v1700", err)
		}
		ie.schedulingRequestToAddModListExt_v1700 = []SchedulingRequestToAddModExt_v1700{}
		for _, i := range tmp_schedulingRequestToAddModListExt_v1700.Value {
			ie.schedulingRequestToAddModListExt_v1700 = append(ie.schedulingRequestToAddModListExt_v1700, *i)
		}
	}
	return nil
}
