package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VictimSystemType_r16 struct {
	gps_r16       *VictimSystemType_r16_gps_r16       `optional`
	glonass_r16   *VictimSystemType_r16_glonass_r16   `optional`
	bds_r16       *VictimSystemType_r16_bds_r16       `optional`
	galileo_r16   *VictimSystemType_r16_galileo_r16   `optional`
	navIC_r16     *VictimSystemType_r16_navIC_r16     `optional`
	wlan_r16      *VictimSystemType_r16_wlan_r16      `optional`
	bluetooth_r16 *VictimSystemType_r16_bluetooth_r16 `optional`
}

func (ie *VictimSystemType_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.gps_r16 != nil, ie.glonass_r16 != nil, ie.bds_r16 != nil, ie.galileo_r16 != nil, ie.navIC_r16 != nil, ie.wlan_r16 != nil, ie.bluetooth_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.gps_r16 != nil {
		if err = ie.gps_r16.Encode(w); err != nil {
			return utils.WrapError("Encode gps_r16", err)
		}
	}
	if ie.glonass_r16 != nil {
		if err = ie.glonass_r16.Encode(w); err != nil {
			return utils.WrapError("Encode glonass_r16", err)
		}
	}
	if ie.bds_r16 != nil {
		if err = ie.bds_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bds_r16", err)
		}
	}
	if ie.galileo_r16 != nil {
		if err = ie.galileo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode galileo_r16", err)
		}
	}
	if ie.navIC_r16 != nil {
		if err = ie.navIC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode navIC_r16", err)
		}
	}
	if ie.wlan_r16 != nil {
		if err = ie.wlan_r16.Encode(w); err != nil {
			return utils.WrapError("Encode wlan_r16", err)
		}
	}
	if ie.bluetooth_r16 != nil {
		if err = ie.bluetooth_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bluetooth_r16", err)
		}
	}
	return nil
}

func (ie *VictimSystemType_r16) Decode(r *uper.UperReader) error {
	var err error
	var gps_r16Present bool
	if gps_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var glonass_r16Present bool
	if glonass_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bds_r16Present bool
	if bds_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var galileo_r16Present bool
	if galileo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var navIC_r16Present bool
	if navIC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var wlan_r16Present bool
	if wlan_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bluetooth_r16Present bool
	if bluetooth_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if gps_r16Present {
		ie.gps_r16 = new(VictimSystemType_r16_gps_r16)
		if err = ie.gps_r16.Decode(r); err != nil {
			return utils.WrapError("Decode gps_r16", err)
		}
	}
	if glonass_r16Present {
		ie.glonass_r16 = new(VictimSystemType_r16_glonass_r16)
		if err = ie.glonass_r16.Decode(r); err != nil {
			return utils.WrapError("Decode glonass_r16", err)
		}
	}
	if bds_r16Present {
		ie.bds_r16 = new(VictimSystemType_r16_bds_r16)
		if err = ie.bds_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bds_r16", err)
		}
	}
	if galileo_r16Present {
		ie.galileo_r16 = new(VictimSystemType_r16_galileo_r16)
		if err = ie.galileo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode galileo_r16", err)
		}
	}
	if navIC_r16Present {
		ie.navIC_r16 = new(VictimSystemType_r16_navIC_r16)
		if err = ie.navIC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode navIC_r16", err)
		}
	}
	if wlan_r16Present {
		ie.wlan_r16 = new(VictimSystemType_r16_wlan_r16)
		if err = ie.wlan_r16.Decode(r); err != nil {
			return utils.WrapError("Decode wlan_r16", err)
		}
	}
	if bluetooth_r16Present {
		ie.bluetooth_r16 = new(VictimSystemType_r16_bluetooth_r16)
		if err = ie.bluetooth_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bluetooth_r16", err)
		}
	}
	return nil
}
