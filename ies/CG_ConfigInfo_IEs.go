package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_IEs struct {
	ue_CapabilityInfo         *[]byte                           `optional`
	candidateCellInfoListMN   *MeasResultList2NR                `optional`
	candidateCellInfoListSN   *[]byte                           `optional`
	measResultCellListSFTD_NR *MeasResultCellListSFTD_NR        `optional`
	scgFailureInfo            *CG_ConfigInfo_IEs_scgFailureInfo `optional`
	configRestrictInfo        *ConfigRestrictInfoSCG            `optional`
	drx_InfoMCG               *DRX_Info                         `optional`
	measConfigMN              *MeasConfigMN                     `optional`
	sourceConfigSCG           *[]byte                           `optional`
	scg_RB_Config             *[]byte                           `optional`
	mcg_RB_Config             *[]byte                           `optional`
	mrdc_AssistanceInfo       *MRDC_AssistanceInfo              `optional`
	nonCriticalExtension      *CG_ConfigInfo_v1540_IEs          `optional`
}

func (ie *CG_ConfigInfo_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ue_CapabilityInfo != nil, ie.candidateCellInfoListMN != nil, ie.candidateCellInfoListSN != nil, ie.measResultCellListSFTD_NR != nil, ie.scgFailureInfo != nil, ie.configRestrictInfo != nil, ie.drx_InfoMCG != nil, ie.measConfigMN != nil, ie.sourceConfigSCG != nil, ie.scg_RB_Config != nil, ie.mcg_RB_Config != nil, ie.mrdc_AssistanceInfo != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ue_CapabilityInfo != nil {
		if err = w.WriteOctetString(*ie.ue_CapabilityInfo, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode ue_CapabilityInfo", err)
		}
	}
	if ie.candidateCellInfoListMN != nil {
		if err = ie.candidateCellInfoListMN.Encode(w); err != nil {
			return utils.WrapError("Encode candidateCellInfoListMN", err)
		}
	}
	if ie.candidateCellInfoListSN != nil {
		if err = w.WriteOctetString(*ie.candidateCellInfoListSN, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode candidateCellInfoListSN", err)
		}
	}
	if ie.measResultCellListSFTD_NR != nil {
		if err = ie.measResultCellListSFTD_NR.Encode(w); err != nil {
			return utils.WrapError("Encode measResultCellListSFTD_NR", err)
		}
	}
	if ie.scgFailureInfo != nil {
		if err = ie.scgFailureInfo.Encode(w); err != nil {
			return utils.WrapError("Encode scgFailureInfo", err)
		}
	}
	if ie.configRestrictInfo != nil {
		if err = ie.configRestrictInfo.Encode(w); err != nil {
			return utils.WrapError("Encode configRestrictInfo", err)
		}
	}
	if ie.drx_InfoMCG != nil {
		if err = ie.drx_InfoMCG.Encode(w); err != nil {
			return utils.WrapError("Encode drx_InfoMCG", err)
		}
	}
	if ie.measConfigMN != nil {
		if err = ie.measConfigMN.Encode(w); err != nil {
			return utils.WrapError("Encode measConfigMN", err)
		}
	}
	if ie.sourceConfigSCG != nil {
		if err = w.WriteOctetString(*ie.sourceConfigSCG, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sourceConfigSCG", err)
		}
	}
	if ie.scg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.scg_RB_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode scg_RB_Config", err)
		}
	}
	if ie.mcg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.mcg_RB_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode mcg_RB_Config", err)
		}
	}
	if ie.mrdc_AssistanceInfo != nil {
		if err = ie.mrdc_AssistanceInfo.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_AssistanceInfo", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ue_CapabilityInfoPresent bool
	if ue_CapabilityInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListMNPresent bool
	if candidateCellInfoListMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListSNPresent bool
	if candidateCellInfoListSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultCellListSFTD_NRPresent bool
	if measResultCellListSFTD_NRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scgFailureInfoPresent bool
	if scgFailureInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configRestrictInfoPresent bool
	if configRestrictInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_InfoMCGPresent bool
	if drx_InfoMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measConfigMNPresent bool
	if measConfigMNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sourceConfigSCGPresent bool
	if sourceConfigSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_RB_ConfigPresent bool
	if scg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mcg_RB_ConfigPresent bool
	if mcg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var mrdc_AssistanceInfoPresent bool
	if mrdc_AssistanceInfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ue_CapabilityInfoPresent {
		var tmp_os_ue_CapabilityInfo []byte
		if tmp_os_ue_CapabilityInfo, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode ue_CapabilityInfo", err)
		}
		ie.ue_CapabilityInfo = &tmp_os_ue_CapabilityInfo
	}
	if candidateCellInfoListMNPresent {
		ie.candidateCellInfoListMN = new(MeasResultList2NR)
		if err = ie.candidateCellInfoListMN.Decode(r); err != nil {
			return utils.WrapError("Decode candidateCellInfoListMN", err)
		}
	}
	if candidateCellInfoListSNPresent {
		var tmp_os_candidateCellInfoListSN []byte
		if tmp_os_candidateCellInfoListSN, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode candidateCellInfoListSN", err)
		}
		ie.candidateCellInfoListSN = &tmp_os_candidateCellInfoListSN
	}
	if measResultCellListSFTD_NRPresent {
		ie.measResultCellListSFTD_NR = new(MeasResultCellListSFTD_NR)
		if err = ie.measResultCellListSFTD_NR.Decode(r); err != nil {
			return utils.WrapError("Decode measResultCellListSFTD_NR", err)
		}
	}
	if scgFailureInfoPresent {
		ie.scgFailureInfo = new(CG_ConfigInfo_IEs_scgFailureInfo)
		if err = ie.scgFailureInfo.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInfo", err)
		}
	}
	if configRestrictInfoPresent {
		ie.configRestrictInfo = new(ConfigRestrictInfoSCG)
		if err = ie.configRestrictInfo.Decode(r); err != nil {
			return utils.WrapError("Decode configRestrictInfo", err)
		}
	}
	if drx_InfoMCGPresent {
		ie.drx_InfoMCG = new(DRX_Info)
		if err = ie.drx_InfoMCG.Decode(r); err != nil {
			return utils.WrapError("Decode drx_InfoMCG", err)
		}
	}
	if measConfigMNPresent {
		ie.measConfigMN = new(MeasConfigMN)
		if err = ie.measConfigMN.Decode(r); err != nil {
			return utils.WrapError("Decode measConfigMN", err)
		}
	}
	if sourceConfigSCGPresent {
		var tmp_os_sourceConfigSCG []byte
		if tmp_os_sourceConfigSCG, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sourceConfigSCG", err)
		}
		ie.sourceConfigSCG = &tmp_os_sourceConfigSCG
	}
	if scg_RB_ConfigPresent {
		var tmp_os_scg_RB_Config []byte
		if tmp_os_scg_RB_Config, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode scg_RB_Config", err)
		}
		ie.scg_RB_Config = &tmp_os_scg_RB_Config
	}
	if mcg_RB_ConfigPresent {
		var tmp_os_mcg_RB_Config []byte
		if tmp_os_mcg_RB_Config, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode mcg_RB_Config", err)
		}
		ie.mcg_RB_Config = &tmp_os_mcg_RB_Config
	}
	if mrdc_AssistanceInfoPresent {
		ie.mrdc_AssistanceInfo = new(MRDC_AssistanceInfo)
		if err = ie.mrdc_AssistanceInfo.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_AssistanceInfo", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1540_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
