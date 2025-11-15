package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarLogMeasReport_r16 struct {
	absoluteTimeInfo_r16         AbsoluteTimeInfo_r16                       `madatory`
	traceReference_r16           TraceReference_r16                         `madatory`
	traceRecordingSessionRef_r16 []byte                                     `lb:2,ub:2,madatory`
	tce_Id_r16                   []byte                                     `lb:1,ub:1,madatory`
	logMeasInfoList_r16          LogMeasInfoList_r16                        `madatory`
	plmn_IdentityList_r16        PLMN_IdentityList2_r16                     `madatory`
	sigLoggedMeasType_r17        VarLogMeasReport_r16_sigLoggedMeasType_r17 `madatory`
}

func (ie *VarLogMeasReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.absoluteTimeInfo_r16.Encode(w); err != nil {
		return utils.WrapError("Encode absoluteTimeInfo_r16", err)
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
	if err = ie.plmn_IdentityList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode plmn_IdentityList_r16", err)
	}
	if err = ie.sigLoggedMeasType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode sigLoggedMeasType_r17", err)
	}
	return nil
}

func (ie *VarLogMeasReport_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.absoluteTimeInfo_r16.Decode(r); err != nil {
		return utils.WrapError("Decode absoluteTimeInfo_r16", err)
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
	if err = ie.plmn_IdentityList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode plmn_IdentityList_r16", err)
	}
	if err = ie.sigLoggedMeasType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode sigLoggedMeasType_r17", err)
	}
	return nil
}
