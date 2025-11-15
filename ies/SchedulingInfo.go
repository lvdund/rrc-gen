package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SchedulingInfo struct {
	si_BroadcastStatus SchedulingInfo_si_BroadcastStatus `madatory`
	si_Periodicity     SchedulingInfo_si_Periodicity     `madatory`
	sib_MappingInfo    SIB_Mapping                       `madatory`
}

func (ie *SchedulingInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.si_BroadcastStatus.Encode(w); err != nil {
		return utils.WrapError("Encode si_BroadcastStatus", err)
	}
	if err = ie.si_Periodicity.Encode(w); err != nil {
		return utils.WrapError("Encode si_Periodicity", err)
	}
	if err = ie.sib_MappingInfo.Encode(w); err != nil {
		return utils.WrapError("Encode sib_MappingInfo", err)
	}
	return nil
}

func (ie *SchedulingInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.si_BroadcastStatus.Decode(r); err != nil {
		return utils.WrapError("Decode si_BroadcastStatus", err)
	}
	if err = ie.si_Periodicity.Decode(r); err != nil {
		return utils.WrapError("Decode si_Periodicity", err)
	}
	if err = ie.sib_MappingInfo.Decode(r); err != nil {
		return utils.WrapError("Decode sib_MappingInfo", err)
	}
	return nil
}
