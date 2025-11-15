package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LoggedMeasurementConfiguration_r16_IEs struct {
	traceReference_r16           TraceReference_r16                                `madatory`
	traceRecordingSessionRef_r16 []byte                                            `lb:2,ub:2,madatory`
	tce_Id_r16                   []byte                                            `lb:1,ub:1,madatory`
	absoluteTimeInfo_r16         AbsoluteTimeInfo_r16                              `madatory`
	areaConfiguration_r16        *AreaConfiguration_r16                            `optional`
	plmn_IdentityList_r16        *PLMN_IdentityList2_r16                           `optional`
	bt_NameList_r16              *BT_NameList_r16                                  `optional,setuprelease`
	wlan_NameList_r16            *WLAN_NameList_r16                                `optional,setuprelease`
	sensor_NameList_r16          *Sensor_NameList_r16                              `optional,setuprelease`
	loggingDuration_r16          LoggingDuration_r16                               `madatory`
	reportType                   LoggedMeasurementConfiguration_r16_IEs_reportType `madatory`
	lateNonCriticalExtension     *[]byte                                           `optional,ext`
	nonCriticalExtension         *LoggedMeasurementConfiguration_v1700_IEs         `optional,ext`
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.areaConfiguration_r16 != nil, ie.plmn_IdentityList_r16 != nil, ie.bt_NameList_r16 != nil, ie.wlan_NameList_r16 != nil, ie.sensor_NameList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.traceReference_r16.Encode(w); err != nil {
		return utils.WrapError("Encode traceReference_r16", err)
	}
	if err = w.WriteOctetString(ie.traceRecordingSessionRef_r16, &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteOctetString traceRecordingSessionRef_r16", err)
	}
	if err = w.WriteOctetString(ie.tce_Id_r16, &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteOctetString tce_Id_r16", err)
	}
	if err = ie.absoluteTimeInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteTimeInfo_r16", err)
	}
	if ie.areaConfiguration_r16 != nil {
		if err = ie.areaConfiguration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode areaConfiguration_r16", err)
		}
	}
	if ie.plmn_IdentityList_r16 != nil {
		if err = ie.plmn_IdentityList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode plmn_IdentityList_r16", err)
		}
	}
	if ie.bt_NameList_r16 != nil {
		tmp_bt_NameList_r16 := utils.SetupRelease[*BT_NameList_r16]{
			Setup: ie.bt_NameList_r16,
		}
		if err = tmp_bt_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode bt_NameList_r16", err)
		}
	}
	if ie.wlan_NameList_r16 != nil {
		tmp_wlan_NameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{
			Setup: ie.wlan_NameList_r16,
		}
		if err = tmp_wlan_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode wlan_NameList_r16", err)
		}
	}
	if ie.sensor_NameList_r16 != nil {
		tmp_sensor_NameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{
			Setup: ie.sensor_NameList_r16,
		}
		if err = tmp_sensor_NameList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode sensor_NameList_r16", err)
		}
	}
	if err = ie.loggingDuration_r16.Encode(w); err != nil {
		return utils.WrapError("Encode loggingDuration_r16", err)
	}
	if err = ie.reportType.Encode(w); err != nil {
		return utils.WrapError("Encode reportType", err)
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var areaConfiguration_r16Present bool
	if areaConfiguration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var plmn_IdentityList_r16Present bool
	if plmn_IdentityList_r16Present, err = r.ReadBool(); err != nil {
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
	if err = ie.traceReference_r16.Decode(r); err != nil {
		return utils.WrapError("Decode traceReference_r16", err)
	}
	var tmp_os_traceRecordingSessionRef_r16 []byte
	if tmp_os_traceRecordingSessionRef_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadOctetString traceRecordingSessionRef_r16", err)
	}
	ie.traceRecordingSessionRef_r16 = tmp_os_traceRecordingSessionRef_r16
	var tmp_os_tce_Id_r16 []byte
	if tmp_os_tce_Id_r16, err = r.ReadOctetString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadOctetString tce_Id_r16", err)
	}
	ie.tce_Id_r16 = tmp_os_tce_Id_r16
	if err = ie.absoluteTimeInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteTimeInfo_r16", err)
	}
	if areaConfiguration_r16Present {
		ie.areaConfiguration_r16 = new(AreaConfiguration_r16)
		if err = ie.areaConfiguration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode areaConfiguration_r16", err)
		}
	}
	if plmn_IdentityList_r16Present {
		ie.plmn_IdentityList_r16 = new(PLMN_IdentityList2_r16)
		if err = ie.plmn_IdentityList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode plmn_IdentityList_r16", err)
		}
	}
	if bt_NameList_r16Present {
		tmp_bt_NameList_r16 := utils.SetupRelease[*BT_NameList_r16]{}
		if err = tmp_bt_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode bt_NameList_r16", err)
		}
		ie.bt_NameList_r16 = tmp_bt_NameList_r16.Setup
	}
	if wlan_NameList_r16Present {
		tmp_wlan_NameList_r16 := utils.SetupRelease[*WLAN_NameList_r16]{}
		if err = tmp_wlan_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode wlan_NameList_r16", err)
		}
		ie.wlan_NameList_r16 = tmp_wlan_NameList_r16.Setup
	}
	if sensor_NameList_r16Present {
		tmp_sensor_NameList_r16 := utils.SetupRelease[*Sensor_NameList_r16]{}
		if err = tmp_sensor_NameList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode sensor_NameList_r16", err)
		}
		ie.sensor_NameList_r16 = tmp_sensor_NameList_r16.Setup
	}
	if err = ie.loggingDuration_r16.Decode(r); err != nil {
		return utils.WrapError("Decode loggingDuration_r16", err)
	}
	if err = ie.reportType.Decode(r); err != nil {
		return utils.WrapError("Decode reportType", err)
	}
	return nil
}
