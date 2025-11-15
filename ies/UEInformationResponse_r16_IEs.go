package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UEInformationResponse_r16_IEs struct {
	measResultIdleEUTRA_r16   *MeasResultIdleEUTRA_r16         `optional`
	measResultIdleNR_r16      *MeasResultIdleNR_r16            `optional`
	logMeasReport_r16         *LogMeasReport_r16               `optional`
	connEstFailReport_r16     *ConnEstFailReport_r16           `optional`
	ra_ReportList_r16         *RA_ReportList_r16               `optional`
	rlf_Report_r16            *RLF_Report_r16                  `optional`
	mobilityHistoryReport_r16 *MobilityHistoryReport_r16       `optional`
	lateNonCriticalExtension  *[]byte                          `optional`
	nonCriticalExtension      *UEInformationResponse_v1700_IEs `optional`
}

func (ie *UEInformationResponse_r16_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.measResultIdleEUTRA_r16 != nil, ie.measResultIdleNR_r16 != nil, ie.logMeasReport_r16 != nil, ie.connEstFailReport_r16 != nil, ie.ra_ReportList_r16 != nil, ie.rlf_Report_r16 != nil, ie.mobilityHistoryReport_r16 != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.measResultIdleEUTRA_r16 != nil {
		if err = ie.measResultIdleEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultIdleEUTRA_r16", err)
		}
	}
	if ie.measResultIdleNR_r16 != nil {
		if err = ie.measResultIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode measResultIdleNR_r16", err)
		}
	}
	if ie.logMeasReport_r16 != nil {
		if err = ie.logMeasReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode logMeasReport_r16", err)
		}
	}
	if ie.connEstFailReport_r16 != nil {
		if err = ie.connEstFailReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode connEstFailReport_r16", err)
		}
	}
	if ie.ra_ReportList_r16 != nil {
		if err = ie.ra_ReportList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ra_ReportList_r16", err)
		}
	}
	if ie.rlf_Report_r16 != nil {
		if err = ie.rlf_Report_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rlf_Report_r16", err)
		}
	}
	if ie.mobilityHistoryReport_r16 != nil {
		if err = ie.mobilityHistoryReport_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityHistoryReport_r16", err)
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

func (ie *UEInformationResponse_r16_IEs) Decode(r *uper.UperReader) error {
	var err error
	var measResultIdleEUTRA_r16Present bool
	if measResultIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultIdleNR_r16Present bool
	if measResultIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var logMeasReport_r16Present bool
	if logMeasReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var connEstFailReport_r16Present bool
	if connEstFailReport_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ra_ReportList_r16Present bool
	if ra_ReportList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rlf_Report_r16Present bool
	if rlf_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityHistoryReport_r16Present bool
	if mobilityHistoryReport_r16Present, err = r.ReadBool(); err != nil {
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
	if measResultIdleEUTRA_r16Present {
		ie.measResultIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
		if err = ie.measResultIdleEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultIdleEUTRA_r16", err)
		}
	}
	if measResultIdleNR_r16Present {
		ie.measResultIdleNR_r16 = new(MeasResultIdleNR_r16)
		if err = ie.measResultIdleNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode measResultIdleNR_r16", err)
		}
	}
	if logMeasReport_r16Present {
		ie.logMeasReport_r16 = new(LogMeasReport_r16)
		if err = ie.logMeasReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode logMeasReport_r16", err)
		}
	}
	if connEstFailReport_r16Present {
		ie.connEstFailReport_r16 = new(ConnEstFailReport_r16)
		if err = ie.connEstFailReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode connEstFailReport_r16", err)
		}
	}
	if ra_ReportList_r16Present {
		ie.ra_ReportList_r16 = new(RA_ReportList_r16)
		if err = ie.ra_ReportList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ra_ReportList_r16", err)
		}
	}
	if rlf_Report_r16Present {
		ie.rlf_Report_r16 = new(RLF_Report_r16)
		if err = ie.rlf_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rlf_Report_r16", err)
		}
	}
	if mobilityHistoryReport_r16Present {
		ie.mobilityHistoryReport_r16 = new(MobilityHistoryReport_r16)
		if err = ie.mobilityHistoryReport_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityHistoryReport_r16", err)
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
		ie.nonCriticalExtension = new(UEInformationResponse_v1700_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
