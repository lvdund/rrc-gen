package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingInfo2_r17 struct {
	si_BroadcastStatus_r17 SchedulingInfo2_r17_si_BroadcastStatus_r17 `madatory`
	si_WindowPosition_r17  int64                                      `lb:1,ub:256,madatory`
	si_Periodicity_r17     SchedulingInfo2_r17_si_Periodicity_r17     `madatory`
	sib_MappingInfo_r17    SIB_Mapping_v1700                          `madatory`
}

func (ie *SchedulingInfo2_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.si_BroadcastStatus_r17.Encode(w); err != nil {
		return utils.WrapError("Encode si_BroadcastStatus_r17", err)
	}
	if err = w.WriteInteger(ie.si_WindowPosition_r17, &uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger si_WindowPosition_r17", err)
	}
	if err = ie.si_Periodicity_r17.Encode(w); err != nil {
		return utils.WrapError("Encode si_Periodicity_r17", err)
	}
	if err = ie.sib_MappingInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sib_MappingInfo_r17", err)
	}
	return nil
}

func (ie *SchedulingInfo2_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.si_BroadcastStatus_r17.Decode(r); err != nil {
		return utils.WrapError("Decode si_BroadcastStatus_r17", err)
	}
	var tmp_int_si_WindowPosition_r17 int64
	if tmp_int_si_WindowPosition_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger si_WindowPosition_r17", err)
	}
	ie.si_WindowPosition_r17 = tmp_int_si_WindowPosition_r17
	if err = ie.si_Periodicity_r17.Decode(r); err != nil {
		return utils.WrapError("Decode si_Periodicity_r17", err)
	}
	if err = ie.sib_MappingInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sib_MappingInfo_r17", err)
	}
	return nil
}
