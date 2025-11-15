package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationRequest_r16_IEs struct {
	idleModeMeasurementReq_r16   *UEInformationRequest_r16_IEs_idleModeMeasurementReq_r16   `optional`
	logMeasReportReq_r16         *UEInformationRequest_r16_IEs_logMeasReportReq_r16         `optional`
	connEstFailReportReq_r16     *UEInformationRequest_r16_IEs_connEstFailReportReq_r16     `optional`
	ra_ReportReq_r16             *UEInformationRequest_r16_IEs_ra_ReportReq_r16             `optional`
	rlf_ReportReq_r16            *UEInformationRequest_r16_IEs_rlf_ReportReq_r16            `optional`
	mobilityHistoryReportReq_r16 *UEInformationRequest_r16_IEs_mobilityHistoryReportReq_r16 `optional`
	lateNonCriticalExtension     *[]byte                                                    `optional`
	nonCriticalExtension         *UEInformationRequest_v1700_IEs                            `optional`
}

func (ie *UEInformationRequest_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idleModeMeasurementReq_r16 != nil, ie.logMeasReportReq_r16 != nil, ie.connEstFailReportReq_r16 != nil, ie.ra_ReportReq_r16 != nil, ie.rlf_ReportReq_r16 != nil, ie.mobilityHistoryReportReq_r16 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idleModeMeasurementReq_r16 != nil {
		if err = ie.idleModeMeasurementReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleModeMeasurementReq_r16", err)
		}
	}
	if ie.logMeasReportReq_r16 != nil {
		if err = ie.logMeasReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasReportReq_r16", err)
		}
	}
	if ie.connEstFailReportReq_r16 != nil {
		if err = ie.connEstFailReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode connEstFailReportReq_r16", err)
		}
	}
	if ie.ra_ReportReq_r16 != nil {
		if err = ie.ra_ReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ra_ReportReq_r16", err)
		}
	}
	if ie.rlf_ReportReq_r16 != nil {
		if err = ie.rlf_ReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlf_ReportReq_r16", err)
		}
	}
	if ie.mobilityHistoryReportReq_r16 != nil {
		if err = ie.mobilityHistoryReportReq_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityHistoryReportReq_r16", err)
		}
	}
	if ie.lateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.lateNonCriticalExtension, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode lateNonCriticalExtension", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *UEInformationRequest_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var idleModeMeasurementReq_r16Present bool
	if idleModeMeasurementReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasReportReq_r16Present bool
	if logMeasReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var connEstFailReportReq_r16Present bool
	if connEstFailReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_ReportReq_r16Present bool
	if ra_ReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlf_ReportReq_r16Present bool
	if rlf_ReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityHistoryReportReq_r16Present bool
	if mobilityHistoryReportReq_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var lateNonCriticalExtensionPresent bool
	if lateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if idleModeMeasurementReq_r16Present {
		ie.idleModeMeasurementReq_r16 = new(UEInformationRequest_r16_IEs_idleModeMeasurementReq_r16)
		if err = ie.idleModeMeasurementReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleModeMeasurementReq_r16", err)
		}
	}
	if logMeasReportReq_r16Present {
		ie.logMeasReportReq_r16 = new(UEInformationRequest_r16_IEs_logMeasReportReq_r16)
		if err = ie.logMeasReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasReportReq_r16", err)
		}
	}
	if connEstFailReportReq_r16Present {
		ie.connEstFailReportReq_r16 = new(UEInformationRequest_r16_IEs_connEstFailReportReq_r16)
		if err = ie.connEstFailReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode connEstFailReportReq_r16", err)
		}
	}
	if ra_ReportReq_r16Present {
		ie.ra_ReportReq_r16 = new(UEInformationRequest_r16_IEs_ra_ReportReq_r16)
		if err = ie.ra_ReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ra_ReportReq_r16", err)
		}
	}
	if rlf_ReportReq_r16Present {
		ie.rlf_ReportReq_r16 = new(UEInformationRequest_r16_IEs_rlf_ReportReq_r16)
		if err = ie.rlf_ReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlf_ReportReq_r16", err)
		}
	}
	if mobilityHistoryReportReq_r16Present {
		ie.mobilityHistoryReportReq_r16 = new(UEInformationRequest_r16_IEs_mobilityHistoryReportReq_r16)
		if err = ie.mobilityHistoryReportReq_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityHistoryReportReq_r16", err)
		}
	}
	if lateNonCriticalExtensionPresent {
		var tmp_os_lateNonCriticalExtension []byte
		if tmp_os_lateNonCriticalExtension, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode lateNonCriticalExtension", err)
		}
		ie.lateNonCriticalExtension = &tmp_os_lateNonCriticalExtension
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(UEInformationRequest_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
