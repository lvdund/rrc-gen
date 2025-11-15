package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_IEs struct {
	scg_CellGroupConfig        *[]byte                     `optional`
	scg_RB_Config              *[]byte                     `optional`
	configRestrictModReq       *ConfigRestrictModReqSCG    `optional`
	drx_InfoSCG                *DRX_Info                   `optional`
	candidateCellInfoListSN    *[]byte                     `optional`
	measConfigSN               *MeasConfigSN               `optional`
	selectedBandCombination    *BandCombinationInfoSN      `optional`
	fr_InfoListSCG             *FR_InfoList                `optional`
	candidateServingFreqListNR *CandidateServingFreqListNR `optional`
	nonCriticalExtension       *CG_Config_v1540_IEs        `optional`
}

func (ie *CG_Config_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scg_CellGroupConfig != nil, ie.scg_RB_Config != nil, ie.configRestrictModReq != nil, ie.drx_InfoSCG != nil, ie.candidateCellInfoListSN != nil, ie.measConfigSN != nil, ie.selectedBandCombination != nil, ie.fr_InfoListSCG != nil, ie.candidateServingFreqListNR != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scg_CellGroupConfig != nil {
		if err = w.WriteOctetString(*ie.scg_CellGroupConfig, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode scg_CellGroupConfig", err)
		}
	}
	if ie.scg_RB_Config != nil {
		if err = w.WriteOctetString(*ie.scg_RB_Config, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode scg_RB_Config", err)
		}
	}
	if ie.configRestrictModReq != nil {
		if err = ie.configRestrictModReq.Encode(w); err != nil {
			return utils.WrapError("Encode configRestrictModReq", err)
		}
	}
	if ie.drx_InfoSCG != nil {
		if err = ie.drx_InfoSCG.Encode(w); err != nil {
			return utils.WrapError("Encode drx_InfoSCG", err)
		}
	}
	if ie.candidateCellInfoListSN != nil {
		if err = w.WriteOctetString(*ie.candidateCellInfoListSN, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode candidateCellInfoListSN", err)
		}
	}
	if ie.measConfigSN != nil {
		if err = ie.measConfigSN.Encode(w); err != nil {
			return utils.WrapError("Encode measConfigSN", err)
		}
	}
	if ie.selectedBandCombination != nil {
		if err = ie.selectedBandCombination.Encode(w); err != nil {
			return utils.WrapError("Encode selectedBandCombination", err)
		}
	}
	if ie.fr_InfoListSCG != nil {
		if err = ie.fr_InfoListSCG.Encode(w); err != nil {
			return utils.WrapError("Encode fr_InfoListSCG", err)
		}
	}
	if ie.candidateServingFreqListNR != nil {
		if err = ie.candidateServingFreqListNR.Encode(w); err != nil {
			return utils.WrapError("Encode candidateServingFreqListNR", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_IEs) Decode(r *uper.UperReader) error {
	var err error
	var scg_CellGroupConfigPresent bool
	if scg_CellGroupConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_RB_ConfigPresent bool
	if scg_RB_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var configRestrictModReqPresent bool
	if configRestrictModReqPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_InfoSCGPresent bool
	if drx_InfoSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListSNPresent bool
	if candidateCellInfoListSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measConfigSNPresent bool
	if measConfigSNPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var selectedBandCombinationPresent bool
	if selectedBandCombinationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr_InfoListSCGPresent bool
	if fr_InfoListSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateServingFreqListNRPresent bool
	if candidateServingFreqListNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scg_CellGroupConfigPresent {
		var tmp_os_scg_CellGroupConfig []byte
		if tmp_os_scg_CellGroupConfig, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode scg_CellGroupConfig", err)
		}
		ie.scg_CellGroupConfig = &tmp_os_scg_CellGroupConfig
	}
	if scg_RB_ConfigPresent {
		var tmp_os_scg_RB_Config []byte
		if tmp_os_scg_RB_Config, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode scg_RB_Config", err)
		}
		ie.scg_RB_Config = &tmp_os_scg_RB_Config
	}
	if configRestrictModReqPresent {
		ie.configRestrictModReq = new(ConfigRestrictModReqSCG)
		if err = ie.configRestrictModReq.Decode(r); err != nil {
			return utils.WrapError("Decode configRestrictModReq", err)
		}
	}
	if drx_InfoSCGPresent {
		ie.drx_InfoSCG = new(DRX_Info)
		if err = ie.drx_InfoSCG.Decode(r); err != nil {
			return utils.WrapError("Decode drx_InfoSCG", err)
		}
	}
	if candidateCellInfoListSNPresent {
		var tmp_os_candidateCellInfoListSN []byte
		if tmp_os_candidateCellInfoListSN, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode candidateCellInfoListSN", err)
		}
		ie.candidateCellInfoListSN = &tmp_os_candidateCellInfoListSN
	}
	if measConfigSNPresent {
		ie.measConfigSN = new(MeasConfigSN)
		if err = ie.measConfigSN.Decode(r); err != nil {
			return utils.WrapError("Decode measConfigSN", err)
		}
	}
	if selectedBandCombinationPresent {
		ie.selectedBandCombination = new(BandCombinationInfoSN)
		if err = ie.selectedBandCombination.Decode(r); err != nil {
			return utils.WrapError("Decode selectedBandCombination", err)
		}
	}
	if fr_InfoListSCGPresent {
		ie.fr_InfoListSCG = new(FR_InfoList)
		if err = ie.fr_InfoListSCG.Decode(r); err != nil {
			return utils.WrapError("Decode fr_InfoListSCG", err)
		}
	}
	if candidateServingFreqListNRPresent {
		ie.candidateServingFreqListNR = new(CandidateServingFreqListNR)
		if err = ie.candidateServingFreqListNR.Decode(r); err != nil {
			return utils.WrapError("Decode candidateServingFreqListNR", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1540_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
