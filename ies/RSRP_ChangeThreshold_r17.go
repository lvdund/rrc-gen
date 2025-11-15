package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RSRP_ChangeThreshold_r17_Enum_dB4    uper.Enumerated = 0
	RSRP_ChangeThreshold_r17_Enum_dB6    uper.Enumerated = 1
	RSRP_ChangeThreshold_r17_Enum_dB8    uper.Enumerated = 2
	RSRP_ChangeThreshold_r17_Enum_dB10   uper.Enumerated = 3
	RSRP_ChangeThreshold_r17_Enum_dB14   uper.Enumerated = 4
	RSRP_ChangeThreshold_r17_Enum_dB18   uper.Enumerated = 5
	RSRP_ChangeThreshold_r17_Enum_dB22   uper.Enumerated = 6
	RSRP_ChangeThreshold_r17_Enum_dB26   uper.Enumerated = 7
	RSRP_ChangeThreshold_r17_Enum_dB30   uper.Enumerated = 8
	RSRP_ChangeThreshold_r17_Enum_dB34   uper.Enumerated = 9
	RSRP_ChangeThreshold_r17_Enum_spare6 uper.Enumerated = 10
	RSRP_ChangeThreshold_r17_Enum_spare5 uper.Enumerated = 11
	RSRP_ChangeThreshold_r17_Enum_spare4 uper.Enumerated = 12
	RSRP_ChangeThreshold_r17_Enum_spare3 uper.Enumerated = 13
	RSRP_ChangeThreshold_r17_Enum_spare2 uper.Enumerated = 14
	RSRP_ChangeThreshold_r17_Enum_spare1 uper.Enumerated = 15
)

type RSRP_ChangeThreshold_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RSRP_ChangeThreshold_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RSRP_ChangeThreshold_r17", err)
	}
	return nil
}

func (ie *RSRP_ChangeThreshold_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RSRP_ChangeThreshold_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
