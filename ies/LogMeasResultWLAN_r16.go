package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasResultWLAN_r16 struct {
	wlan_Identifiers_r16 WLAN_Identifiers_r16 `madatory`
	rssiWLAN_r16         *WLAN_RSSI_Range_r16 `optional`
	rtt_WLAN_r16         *WLAN_RTT_r16        `optional`
}

func (ie *LogMeasResultWLAN_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rssiWLAN_r16 != nil, ie.rtt_WLAN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.wlan_Identifiers_r16.Encode(w); err != nil {
		return utils.WrapError("Encode wlan_Identifiers_r16", err)
	}
	if ie.rssiWLAN_r16 != nil {
		if err = ie.rssiWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rssiWLAN_r16", err)
		}
	}
	if ie.rtt_WLAN_r16 != nil {
		if err = ie.rtt_WLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rtt_WLAN_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasResultWLAN_r16) Decode(r *uper.UperReader) error {
	var err error
	var rssiWLAN_r16Present bool
	if rssiWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rtt_WLAN_r16Present bool
	if rtt_WLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.wlan_Identifiers_r16.Decode(r); err != nil {
		return utils.WrapError("Decode wlan_Identifiers_r16", err)
	}
	if rssiWLAN_r16Present {
		ie.rssiWLAN_r16 = new(WLAN_RSSI_Range_r16)
		if err = ie.rssiWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rssiWLAN_r16", err)
		}
	}
	if rtt_WLAN_r16Present {
		ie.rtt_WLAN_r16 = new(WLAN_RTT_r16)
		if err = ie.rtt_WLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rtt_WLAN_r16", err)
		}
	}
	return nil
}
