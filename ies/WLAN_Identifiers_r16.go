package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type WLAN_Identifiers_r16 struct {
	ssid_r16   *[]byte `lb:1,ub:32,optional`
	bssid_r16  *[]byte `lb:6,ub:6,optional`
	hessid_r16 *[]byte `lb:6,ub:6,optional`
}

func (ie *WLAN_Identifiers_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssid_r16 != nil, ie.bssid_r16 != nil, ie.hessid_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ssid_r16 != nil {
		if err = w.WriteOctetString(*ie.ssid_r16, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode ssid_r16", err)
		}
	}
	if ie.bssid_r16 != nil {
		if err = w.WriteOctetString(*ie.bssid_r16, &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode bssid_r16", err)
		}
	}
	if ie.hessid_r16 != nil {
		if err = w.WriteOctetString(*ie.hessid_r16, &uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode hessid_r16", err)
		}
	}
	return nil
}

func (ie *WLAN_Identifiers_r16) Decode(r *uper.UperReader) error {
	var err error
	var ssid_r16Present bool
	if ssid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bssid_r16Present bool
	if bssid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var hessid_r16Present bool
	if hessid_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ssid_r16Present {
		var tmp_os_ssid_r16 []byte
		if tmp_os_ssid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode ssid_r16", err)
		}
		ie.ssid_r16 = &tmp_os_ssid_r16
	}
	if bssid_r16Present {
		var tmp_os_bssid_r16 []byte
		if tmp_os_bssid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode bssid_r16", err)
		}
		ie.bssid_r16 = &tmp_os_bssid_r16
	}
	if hessid_r16Present {
		var tmp_os_hessid_r16 []byte
		if tmp_os_hessid_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 6, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode hessid_r16", err)
		}
		ie.hessid_r16 = &tmp_os_hessid_r16
	}
	return nil
}
