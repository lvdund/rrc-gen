package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarLogMeasConfig_r16_IEs struct {
	areaConfiguration_r16   *AreaConfiguration_r16                            `optional`
	bt_NameList_r16         *BT_NameList_r16                                  `optional`
	wlan_NameList_r16       *WLAN_NameList_r16                                `optional`
	sensor_NameList_r16     *Sensor_NameList_r16                              `optional`
	loggingDuration_r16     LoggingDuration_r16                               `madatory`
	reportType              VarLogMeasConfig_r16_IEs_reportType               `madatory`
	earlyMeasIndication_r17 *VarLogMeasConfig_r16_IEs_earlyMeasIndication_r17 `optional`
	areaConfiguration_v1700 *AreaConfiguration_v1700                          `optional`
}

func (ie *VarLogMeasConfig_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.areaConfiguration_r16 != nil, ie.bt_NameList_r16 != nil, ie.wlan_NameList_r16 != nil, ie.sensor_NameList_r16 != nil, ie.earlyMeasIndication_r17 != nil, ie.areaConfiguration_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.areaConfiguration_r16 != nil {
		if err = ie.areaConfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode areaConfiguration_r16", err)
		}
	}
	if ie.bt_NameList_r16 != nil {
		if err = ie.bt_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bt_NameList_r16", err)
		}
	}
	if ie.wlan_NameList_r16 != nil {
		if err = ie.wlan_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode wlan_NameList_r16", err)
		}
	}
	if ie.sensor_NameList_r16 != nil {
		if err = ie.sensor_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sensor_NameList_r16", err)
		}
	}
	if err = ie.loggingDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode loggingDuration_r16", err)
	}
	if err = ie.reportType.Encode(w); err != nil {
		return utils.WrapError("Encode reportType", err)
	}
	if ie.earlyMeasIndication_r17 != nil {
		if err = ie.earlyMeasIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode earlyMeasIndication_r17", err)
		}
	}
	if ie.areaConfiguration_v1700 != nil {
		if err = ie.areaConfiguration_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode areaConfiguration_v1700", err)
		}
	}
	return nil
}

func (ie *VarLogMeasConfig_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var areaConfiguration_r16Present bool
	if areaConfiguration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var bt_NameList_r16Present bool
	if bt_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var wlan_NameList_r16Present bool
	if wlan_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var sensor_NameList_r16Present bool
	if sensor_NameList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var earlyMeasIndication_r17Present bool
	if earlyMeasIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var areaConfiguration_v1700Present bool
	if areaConfiguration_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if areaConfiguration_r16Present {
		ie.areaConfiguration_r16 = new(AreaConfiguration_r16)
		if err = ie.areaConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode areaConfiguration_r16", err)
		}
	}
	if bt_NameList_r16Present {
		ie.bt_NameList_r16 = new(BT_NameList_r16)
		if err = ie.bt_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bt_NameList_r16", err)
		}
	}
	if wlan_NameList_r16Present {
		ie.wlan_NameList_r16 = new(WLAN_NameList_r16)
		if err = ie.wlan_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode wlan_NameList_r16", err)
		}
	}
	if sensor_NameList_r16Present {
		ie.sensor_NameList_r16 = new(Sensor_NameList_r16)
		if err = ie.sensor_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sensor_NameList_r16", err)
		}
	}
	if err = ie.loggingDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode loggingDuration_r16", err)
	}
	if err = ie.reportType.Decode(r); err != nil {
		return utils.WrapError("Decode reportType", err)
	}
	if earlyMeasIndication_r17Present {
		ie.earlyMeasIndication_r17 = new(VarLogMeasConfig_r16_IEs_earlyMeasIndication_r17)
		if err = ie.earlyMeasIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode earlyMeasIndication_r17", err)
		}
	}
	if areaConfiguration_v1700Present {
		ie.areaConfiguration_v1700 = new(AreaConfiguration_v1700)
		if err = ie.areaConfiguration_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode areaConfiguration_v1700", err)
		}
	}
	return nil
}
