package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LocationInfo_r16 struct {
	commonLocationInfo_r16  *CommonLocationInfo_r16    `optional`
	bt_LocationInfo_r16     *LogMeasResultListBT_r16   `optional`
	wlan_LocationInfo_r16   *LogMeasResultListWLAN_r16 `optional`
	sensor_LocationInfo_r16 *Sensor_LocationInfo_r16   `optional`
}

func (ie *LocationInfo_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.commonLocationInfo_r16 != nil, ie.bt_LocationInfo_r16 != nil, ie.wlan_LocationInfo_r16 != nil, ie.sensor_LocationInfo_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.commonLocationInfo_r16 != nil {
		if err = ie.commonLocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode commonLocationInfo_r16", err)
		}
	}
	if ie.bt_LocationInfo_r16 != nil {
		if err = ie.bt_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bt_LocationInfo_r16", err)
		}
	}
	if ie.wlan_LocationInfo_r16 != nil {
		if err = ie.wlan_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode wlan_LocationInfo_r16", err)
		}
	}
	if ie.sensor_LocationInfo_r16 != nil {
		if err = ie.sensor_LocationInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sensor_LocationInfo_r16", err)
		}
	}
	return nil
}

func (ie *LocationInfo_r16) Decode(r *uper.UperReader) error {
	var err error
	var commonLocationInfo_r16Present bool
	if commonLocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bt_LocationInfo_r16Present bool
	if bt_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var wlan_LocationInfo_r16Present bool
	if wlan_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sensor_LocationInfo_r16Present bool
	if sensor_LocationInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if commonLocationInfo_r16Present {
		ie.commonLocationInfo_r16 = new(CommonLocationInfo_r16)
		if err = ie.commonLocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode commonLocationInfo_r16", err)
		}
	}
	if bt_LocationInfo_r16Present {
		ie.bt_LocationInfo_r16 = new(LogMeasResultListBT_r16)
		if err = ie.bt_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bt_LocationInfo_r16", err)
		}
	}
	if wlan_LocationInfo_r16Present {
		ie.wlan_LocationInfo_r16 = new(LogMeasResultListWLAN_r16)
		if err = ie.wlan_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode wlan_LocationInfo_r16", err)
		}
	}
	if sensor_LocationInfo_r16Present {
		ie.sensor_LocationInfo_r16 = new(Sensor_LocationInfo_r16)
		if err = ie.sensor_LocationInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sensor_LocationInfo_r16", err)
		}
	}
	return nil
}
