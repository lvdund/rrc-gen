package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB1 struct {
	cellSelectionInfo        *SIB1_cellSelectionInfo     `lb:1,ub:8,optional`
	cellAccessRelatedInfo    CellAccessRelatedInfo       `madatory`
	connEstFailureControl    *ConnEstFailureControl      `optional`
	si_SchedulingInfo        *SI_SchedulingInfo          `optional`
	servingCellConfigCommon  *ServingCellConfigCommonSIB `optional`
	ims_EmergencySupport     *SIB1_ims_EmergencySupport  `optional`
	eCallOverIMS_Support     *SIB1_eCallOverIMS_Support  `optional`
	ue_TimersAndConstants    *UE_TimersAndConstants      `optional`
	uac_BarringInfo          *SIB1_uac_BarringInfo       `lb:2,ub:maxPLMN,optional`
	useFullResumeID          *SIB1_useFullResumeID       `optional`
	lateNonCriticalExtension *[]byte                     `optional`
	nonCriticalExtension     *SIB1_v1610_IEs             `optional`
}

func (ie *SIB1) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellSelectionInfo != nil, ie.connEstFailureControl != nil, ie.si_SchedulingInfo != nil, ie.servingCellConfigCommon != nil, ie.ims_EmergencySupport != nil, ie.eCallOverIMS_Support != nil, ie.ue_TimersAndConstants != nil, ie.uac_BarringInfo != nil, ie.useFullResumeID != nil, ie.lateNonCriticalExtension != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellSelectionInfo != nil {
		if err = ie.cellSelectionInfo.Encode(w); err != nil {
			return utils.WrapError("Encode cellSelectionInfo", err)
		}
	}
	if err = ie.cellAccessRelatedInfo.Encode(w); err != nil {
		return utils.WrapError("Encode cellAccessRelatedInfo", err)
	}
	if ie.connEstFailureControl != nil {
		if err = ie.connEstFailureControl.Encode(w); err != nil {
			return utils.WrapError("Encode connEstFailureControl", err)
		}
	}
	if ie.si_SchedulingInfo != nil {
		if err = ie.si_SchedulingInfo.Encode(w); err != nil {
			return utils.WrapError("Encode si_SchedulingInfo", err)
		}
	}
	if ie.servingCellConfigCommon != nil {
		if err = ie.servingCellConfigCommon.Encode(w); err != nil {
			return utils.WrapError("Encode servingCellConfigCommon", err)
		}
	}
	if ie.ims_EmergencySupport != nil {
		if err = ie.ims_EmergencySupport.Encode(w); err != nil {
			return utils.WrapError("Encode ims_EmergencySupport", err)
		}
	}
	if ie.eCallOverIMS_Support != nil {
		if err = ie.eCallOverIMS_Support.Encode(w); err != nil {
			return utils.WrapError("Encode eCallOverIMS_Support", err)
		}
	}
	if ie.ue_TimersAndConstants != nil {
		if err = ie.ue_TimersAndConstants.Encode(w); err != nil {
			return utils.WrapError("Encode ue_TimersAndConstants", err)
		}
	}
	if ie.uac_BarringInfo != nil {
		if err = ie.uac_BarringInfo.Encode(w); err != nil {
			return utils.WrapError("Encode uac_BarringInfo", err)
		}
	}
	if ie.useFullResumeID != nil {
		if err = ie.useFullResumeID.Encode(w); err != nil {
			return utils.WrapError("Encode useFullResumeID", err)
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

func (ie *SIB1) Decode(r *uper.UperReader) error {
	var err error
	var cellSelectionInfoPresent bool
	if cellSelectionInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var connEstFailureControlPresent bool
	if connEstFailureControlPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var si_SchedulingInfoPresent bool
	if si_SchedulingInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var servingCellConfigCommonPresent bool
	if servingCellConfigCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ims_EmergencySupportPresent bool
	if ims_EmergencySupportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var eCallOverIMS_SupportPresent bool
	if eCallOverIMS_SupportPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ue_TimersAndConstantsPresent bool
	if ue_TimersAndConstantsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var uac_BarringInfoPresent bool
	if uac_BarringInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var useFullResumeIDPresent bool
	if useFullResumeIDPresent, err = r.ReadBool(); err != nil {
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
	if cellSelectionInfoPresent {
		ie.cellSelectionInfo = new(SIB1_cellSelectionInfo)
		if err = ie.cellSelectionInfo.Decode(r); err != nil {
			return utils.WrapError("Decode cellSelectionInfo", err)
		}
	}
	if err = ie.cellAccessRelatedInfo.Decode(r); err != nil {
		return utils.WrapError("Decode cellAccessRelatedInfo", err)
	}
	if connEstFailureControlPresent {
		ie.connEstFailureControl = new(ConnEstFailureControl)
		if err = ie.connEstFailureControl.Decode(r); err != nil {
			return utils.WrapError("Decode connEstFailureControl", err)
		}
	}
	if si_SchedulingInfoPresent {
		ie.si_SchedulingInfo = new(SI_SchedulingInfo)
		if err = ie.si_SchedulingInfo.Decode(r); err != nil {
			return utils.WrapError("Decode si_SchedulingInfo", err)
		}
	}
	if servingCellConfigCommonPresent {
		ie.servingCellConfigCommon = new(ServingCellConfigCommonSIB)
		if err = ie.servingCellConfigCommon.Decode(r); err != nil {
			return utils.WrapError("Decode servingCellConfigCommon", err)
		}
	}
	if ims_EmergencySupportPresent {
		ie.ims_EmergencySupport = new(SIB1_ims_EmergencySupport)
		if err = ie.ims_EmergencySupport.Decode(r); err != nil {
			return utils.WrapError("Decode ims_EmergencySupport", err)
		}
	}
	if eCallOverIMS_SupportPresent {
		ie.eCallOverIMS_Support = new(SIB1_eCallOverIMS_Support)
		if err = ie.eCallOverIMS_Support.Decode(r); err != nil {
			return utils.WrapError("Decode eCallOverIMS_Support", err)
		}
	}
	if ue_TimersAndConstantsPresent {
		ie.ue_TimersAndConstants = new(UE_TimersAndConstants)
		if err = ie.ue_TimersAndConstants.Decode(r); err != nil {
			return utils.WrapError("Decode ue_TimersAndConstants", err)
		}
	}
	if uac_BarringInfoPresent {
		ie.uac_BarringInfo = new(SIB1_uac_BarringInfo)
		if err = ie.uac_BarringInfo.Decode(r); err != nil {
			return utils.WrapError("Decode uac_BarringInfo", err)
		}
	}
	if useFullResumeIDPresent {
		ie.useFullResumeID = new(SIB1_useFullResumeID)
		if err = ie.useFullResumeID.Decode(r); err != nil {
			return utils.WrapError("Decode useFullResumeID", err)
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
		ie.nonCriticalExtension = new(SIB1_v1610_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
