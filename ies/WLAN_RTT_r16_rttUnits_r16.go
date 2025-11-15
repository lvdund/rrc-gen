package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	WLAN_RTT_r16_rttUnits_r16_Enum_microseconds          uper.Enumerated = 0
	WLAN_RTT_r16_rttUnits_r16_Enum_hundredsofnanoseconds uper.Enumerated = 1
	WLAN_RTT_r16_rttUnits_r16_Enum_tensofnanoseconds     uper.Enumerated = 2
	WLAN_RTT_r16_rttUnits_r16_Enum_nanoseconds           uper.Enumerated = 3
	WLAN_RTT_r16_rttUnits_r16_Enum_tenthsofnanoseconds   uper.Enumerated = 4
)

type WLAN_RTT_r16_rttUnits_r16 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *WLAN_RTT_r16_rttUnits_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode WLAN_RTT_r16_rttUnits_r16", err)
	}
	return nil
}

func (ie *WLAN_RTT_r16_rttUnits_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode WLAN_RTT_r16_rttUnits_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
