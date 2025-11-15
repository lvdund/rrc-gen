package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeComplete_v1610_IEs struct {
	idleMeasAvailable_r16        *RRCResumeComplete_v1610_IEs_idleMeasAvailable_r16    `optional`
	measResultIdleEUTRA_r16      *MeasResultIdleEUTRA_r16                              `optional`
	measResultIdleNR_r16         *MeasResultIdleNR_r16                                 `optional`
	scg_Response_r16             *RRCResumeComplete_v1610_IEs_scg_Response_r16         `optional`
	ue_MeasurementsAvailable_r16 *UE_MeasurementsAvailable_r16                         `optional`
	mobilityHistoryAvail_r16     *RRCResumeComplete_v1610_IEs_mobilityHistoryAvail_r16 `optional`
	mobilityState_r16            *RRCResumeComplete_v1610_IEs_mobilityState_r16        `optional`
	needForGapsInfoNR_r16        *NeedForGapsInfoNR_r16                                `optional`
	nonCriticalExtension         *RRCResumeComplete_v1640_IEs                          `optional`
}

func (ie *RRCResumeComplete_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.idleMeasAvailable_r16 != nil, ie.measResultIdleEUTRA_r16 != nil, ie.measResultIdleNR_r16 != nil, ie.scg_Response_r16 != nil, ie.ue_MeasurementsAvailable_r16 != nil, ie.mobilityHistoryAvail_r16 != nil, ie.mobilityState_r16 != nil, ie.needForGapsInfoNR_r16 != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.idleMeasAvailable_r16 != nil {
		if err = ie.idleMeasAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode idleMeasAvailable_r16", err)
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
	if ie.scg_Response_r16 != nil {
		if err = ie.scg_Response_r16.Encode(w); err != nil {
			return utils.WrapError("Encode scg_Response_r16", err)
		}
	}
	if ie.ue_MeasurementsAvailable_r16 != nil {
		if err = ie.ue_MeasurementsAvailable_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ue_MeasurementsAvailable_r16", err)
		}
	}
	if ie.mobilityHistoryAvail_r16 != nil {
		if err = ie.mobilityHistoryAvail_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityHistoryAvail_r16", err)
		}
	}
	if ie.mobilityState_r16 != nil {
		if err = ie.mobilityState_r16.Encode(w); err != nil {
			return utils.WrapError("Encode mobilityState_r16", err)
		}
	}
	if ie.needForGapsInfoNR_r16 != nil {
		if err = ie.needForGapsInfoNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode needForGapsInfoNR_r16", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCResumeComplete_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var idleMeasAvailable_r16Present bool
	if idleMeasAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultIdleEUTRA_r16Present bool
	if measResultIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultIdleNR_r16Present bool
	if measResultIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_Response_r16Present bool
	if scg_Response_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_MeasurementsAvailable_r16Present bool
	if ue_MeasurementsAvailable_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityHistoryAvail_r16Present bool
	if mobilityHistoryAvail_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var mobilityState_r16Present bool
	if mobilityState_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapsInfoNR_r16Present bool
	if needForGapsInfoNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if idleMeasAvailable_r16Present {
		ie.idleMeasAvailable_r16 = new(RRCResumeComplete_v1610_IEs_idleMeasAvailable_r16)
		if err = ie.idleMeasAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode idleMeasAvailable_r16", err)
		}
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
	if scg_Response_r16Present {
		ie.scg_Response_r16 = new(RRCResumeComplete_v1610_IEs_scg_Response_r16)
		if err = ie.scg_Response_r16.Decode(r); err != nil {
			return utils.WrapError("Decode scg_Response_r16", err)
		}
	}
	if ue_MeasurementsAvailable_r16Present {
		ie.ue_MeasurementsAvailable_r16 = new(UE_MeasurementsAvailable_r16)
		if err = ie.ue_MeasurementsAvailable_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ue_MeasurementsAvailable_r16", err)
		}
	}
	if mobilityHistoryAvail_r16Present {
		ie.mobilityHistoryAvail_r16 = new(RRCResumeComplete_v1610_IEs_mobilityHistoryAvail_r16)
		if err = ie.mobilityHistoryAvail_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityHistoryAvail_r16", err)
		}
	}
	if mobilityState_r16Present {
		ie.mobilityState_r16 = new(RRCResumeComplete_v1610_IEs_mobilityState_r16)
		if err = ie.mobilityState_r16.Decode(r); err != nil {
			return utils.WrapError("Decode mobilityState_r16", err)
		}
	}
	if needForGapsInfoNR_r16Present {
		ie.needForGapsInfoNR_r16 = new(NeedForGapsInfoNR_r16)
		if err = ie.needForGapsInfoNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode needForGapsInfoNR_r16", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(RRCResumeComplete_v1640_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
