package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VictimSystemType struct {
	gps       *VictimSystemType_gps       `optional`
	glonass   *VictimSystemType_glonass   `optional`
	bds       *VictimSystemType_bds       `optional`
	galileo   *VictimSystemType_galileo   `optional`
	wlan      *VictimSystemType_wlan      `optional`
	bluetooth *VictimSystemType_bluetooth `optional`
}

func (ie *VictimSystemType) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gps != nil, ie.glonass != nil, ie.bds != nil, ie.galileo != nil, ie.wlan != nil, ie.bluetooth != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gps != nil {
		if err = ie.gps.Encode(w); err != nil {
			return utils.WrapError("Encode gps", err)
		}
	}
	if ie.glonass != nil {
		if err = ie.glonass.Encode(w); err != nil {
			return utils.WrapError("Encode glonass", err)
		}
	}
	if ie.bds != nil {
		if err = ie.bds.Encode(w); err != nil {
			return utils.WrapError("Encode bds", err)
		}
	}
	if ie.galileo != nil {
		if err = ie.galileo.Encode(w); err != nil {
			return utils.WrapError("Encode galileo", err)
		}
	}
	if ie.wlan != nil {
		if err = ie.wlan.Encode(w); err != nil {
			return utils.WrapError("Encode wlan", err)
		}
	}
	if ie.bluetooth != nil {
		if err = ie.bluetooth.Encode(w); err != nil {
			return utils.WrapError("Encode bluetooth", err)
		}
	}
	return nil
}

func (ie *VictimSystemType) Decode(r *uper.UperReader) error {
	var err error
	var gpsPresent bool
	if gpsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var glonassPresent bool
	if glonassPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bdsPresent bool
	if bdsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var galileoPresent bool
	if galileoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var wlanPresent bool
	if wlanPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var bluetoothPresent bool
	if bluetoothPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if gpsPresent {
		ie.gps = new(VictimSystemType_gps)
		if err = ie.gps.Decode(r); err != nil {
			return utils.WrapError("Decode gps", err)
		}
	}
	if glonassPresent {
		ie.glonass = new(VictimSystemType_glonass)
		if err = ie.glonass.Decode(r); err != nil {
			return utils.WrapError("Decode glonass", err)
		}
	}
	if bdsPresent {
		ie.bds = new(VictimSystemType_bds)
		if err = ie.bds.Decode(r); err != nil {
			return utils.WrapError("Decode bds", err)
		}
	}
	if galileoPresent {
		ie.galileo = new(VictimSystemType_galileo)
		if err = ie.galileo.Decode(r); err != nil {
			return utils.WrapError("Decode galileo", err)
		}
	}
	if wlanPresent {
		ie.wlan = new(VictimSystemType_wlan)
		if err = ie.wlan.Decode(r); err != nil {
			return utils.WrapError("Decode wlan", err)
		}
	}
	if bluetoothPresent {
		ie.bluetooth = new(VictimSystemType_bluetooth)
		if err = ie.bluetooth.Decode(r); err != nil {
			return utils.WrapError("Decode bluetooth", err)
		}
	}
	return nil
}
