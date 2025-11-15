package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type LogMeasReport_r16 struct {
	absoluteTimeStamp_r16        AbsoluteTimeInfo_r16                        `madatory`
	traceReference_r16           TraceReference_r16                          `madatory`
	traceRecordingSessionRef_r16 []byte                                      `lb:2,ub:2,madatory`
	tce_Id_r16                   []byte                                      `lb:1,ub:1,madatory`
	logMeasInfoList_r16          LogMeasInfoList_r16                         `madatory`
	logMeasAvailable_r16         *LogMeasReport_r16_logMeasAvailable_r16     `optional`
	logMeasAvailableBT_r16       *LogMeasReport_r16_logMeasAvailableBT_r16   `optional`
	logMeasAvailableWLAN_r16     *LogMeasReport_r16_logMeasAvailableWLAN_r16 `optional`
}

func (ie *LogMeasReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.logMeasAvailable_r16 != nil, ie.logMeasAvailableBT_r16 != nil, ie.logMeasAvailableWLAN_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.absoluteTimeStamp_r16.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteTimeStamp_r16", err)
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
	if err = ie.logMeasInfoList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode logMeasInfoList_r16", err)
	}
	if ie.logMeasAvailable_r16 != nil {
		if err = ie.logMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailable_r16", err)
		}
	}
	if ie.logMeasAvailableBT_r16 != nil {
		if err = ie.logMeasAvailableBT_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailableBT_r16", err)
		}
	}
	if ie.logMeasAvailableWLAN_r16 != nil {
		if err = ie.logMeasAvailableWLAN_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasAvailableWLAN_r16", err)
		}
	}
	return nil
}

func (ie *LogMeasReport_r16) Decode(r *uper.UperReader) error {
	var err error
	var logMeasAvailable_r16Present bool
	if logMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasAvailableBT_r16Present bool
	if logMeasAvailableBT_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasAvailableWLAN_r16Present bool
	if logMeasAvailableWLAN_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.absoluteTimeStamp_r16.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteTimeStamp_r16", err)
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
	if err = ie.logMeasInfoList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode logMeasInfoList_r16", err)
	}
	if logMeasAvailable_r16Present {
		ie.logMeasAvailable_r16 = new(LogMeasReport_r16_logMeasAvailable_r16)
		if err = ie.logMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailable_r16", err)
		}
	}
	if logMeasAvailableBT_r16Present {
		ie.logMeasAvailableBT_r16 = new(LogMeasReport_r16_logMeasAvailableBT_r16)
		if err = ie.logMeasAvailableBT_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailableBT_r16", err)
		}
	}
	if logMeasAvailableWLAN_r16Present {
		ie.logMeasAvailableWLAN_r16 = new(LogMeasReport_r16_logMeasAvailableWLAN_r16)
		if err = ie.logMeasAvailableWLAN_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasAvailableWLAN_r16", err)
		}
	}
	return nil
}
