package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	UAC_BarringInfoSet_uac_BarringTime_Enum_s4   uper.Enumerated = 0
	UAC_BarringInfoSet_uac_BarringTime_Enum_s8   uper.Enumerated = 1
	UAC_BarringInfoSet_uac_BarringTime_Enum_s16  uper.Enumerated = 2
	UAC_BarringInfoSet_uac_BarringTime_Enum_s32  uper.Enumerated = 3
	UAC_BarringInfoSet_uac_BarringTime_Enum_s64  uper.Enumerated = 4
	UAC_BarringInfoSet_uac_BarringTime_Enum_s128 uper.Enumerated = 5
	UAC_BarringInfoSet_uac_BarringTime_Enum_s256 uper.Enumerated = 6
	UAC_BarringInfoSet_uac_BarringTime_Enum_s512 uper.Enumerated = 7
)

type UAC_BarringInfoSet_uac_BarringTime struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *UAC_BarringInfoSet_uac_BarringTime) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode UAC_BarringInfoSet_uac_BarringTime", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet_uac_BarringTime) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode UAC_BarringInfoSet_uac_BarringTime", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
