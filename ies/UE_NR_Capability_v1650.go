package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UE_NR_Capability_v1650 struct {
	mpsPriorityIndication_r16 *UE_NR_Capability_v1650_mpsPriorityIndication_r16 `optional`
	highSpeedParameters_v1650 *HighSpeedParameters_v1650                        `optional`
	nonCriticalExtension      *UE_NR_Capability_v1690                           `optional`
}

func (ie *UE_NR_Capability_v1650) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mpsPriorityIndication_r16 != nil, ie.highSpeedParameters_v1650 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mpsPriorityIndication_r16 != nil {
		if err = ie.mpsPriorityIndication_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mpsPriorityIndication_r16", err)
		}
	}
	if ie.highSpeedParameters_v1650 != nil {
		if err = ie.highSpeedParameters_v1650.Encode(w); err != nil {
			return utils.WrapError("Encode highSpeedParameters_v1650", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UE_NR_Capability_v1650) Decode(r *uper.UperReader) error {
	var err error
	var mpsPriorityIndication_r16Present bool
	if mpsPriorityIndication_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var highSpeedParameters_v1650Present bool
	if highSpeedParameters_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mpsPriorityIndication_r16Present {
		ie.mpsPriorityIndication_r16 = new(UE_NR_Capability_v1650_mpsPriorityIndication_r16)
		if err = ie.mpsPriorityIndication_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mpsPriorityIndication_r16", err)
		}
	}
	if highSpeedParameters_v1650Present {
		ie.highSpeedParameters_v1650 = new(HighSpeedParameters_v1650)
		if err = ie.highSpeedParameters_v1650.Decode(r); err != nil {
			return utils.WrapError("Decode highSpeedParameters_v1650", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UE_NR_Capability_v1690)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
