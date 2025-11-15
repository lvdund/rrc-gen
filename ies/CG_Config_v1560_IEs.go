package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs struct {
	pSCellFrequencyEUTRA          *ARFCN_ValueEUTRA                           `optional`
	scg_CellGroupConfigEUTRA      *[]byte                                     `optional`
	candidateCellInfoListSN_EUTRA *[]byte                                     `optional`
	candidateServingFreqListEUTRA *CandidateServingFreqListEUTRA              `optional`
	needForGaps                   *CG_Config_v1560_IEs_needForGaps            `optional`
	drx_ConfigSCG                 *DRX_Config                                 `optional`
	reportCGI_RequestEUTRA        *CG_Config_v1560_IEs_reportCGI_RequestEUTRA `optional`
	nonCriticalExtension          *CG_Config_v1590_IEs                        `optional`
}

func (ie *CG_Config_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pSCellFrequencyEUTRA != nil, ie.scg_CellGroupConfigEUTRA != nil, ie.candidateCellInfoListSN_EUTRA != nil, ie.candidateServingFreqListEUTRA != nil, ie.needForGaps != nil, ie.drx_ConfigSCG != nil, ie.reportCGI_RequestEUTRA != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pSCellFrequencyEUTRA != nil {
		if err = ie.pSCellFrequencyEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode pSCellFrequencyEUTRA", err)
		}
	}
	if ie.scg_CellGroupConfigEUTRA != nil {
		if err = w.WriteOctetString(*ie.scg_CellGroupConfigEUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode scg_CellGroupConfigEUTRA", err)
		}
	}
	if ie.candidateCellInfoListSN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.candidateCellInfoListSN_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode candidateCellInfoListSN_EUTRA", err)
		}
	}
	if ie.candidateServingFreqListEUTRA != nil {
		if err = ie.candidateServingFreqListEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode candidateServingFreqListEUTRA", err)
		}
	}
	if ie.needForGaps != nil {
		if err = ie.needForGaps.Encode(w); err != nil {
			return utils.WrapError("Encode needForGaps", err)
		}
	}
	if ie.drx_ConfigSCG != nil {
		if err = ie.drx_ConfigSCG.Encode(w); err != nil {
			return utils.WrapError("Encode drx_ConfigSCG", err)
		}
	}
	if ie.reportCGI_RequestEUTRA != nil {
		if err = ie.reportCGI_RequestEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode reportCGI_RequestEUTRA", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var pSCellFrequencyEUTRAPresent bool
	if pSCellFrequencyEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scg_CellGroupConfigEUTRAPresent bool
	if scg_CellGroupConfigEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListSN_EUTRAPresent bool
	if candidateCellInfoListSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateServingFreqListEUTRAPresent bool
	if candidateServingFreqListEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var needForGapsPresent bool
	if needForGapsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_ConfigSCGPresent bool
	if drx_ConfigSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportCGI_RequestEUTRAPresent bool
	if reportCGI_RequestEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pSCellFrequencyEUTRAPresent {
		ie.pSCellFrequencyEUTRA = new(ARFCN_ValueEUTRA)
		if err = ie.pSCellFrequencyEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode pSCellFrequencyEUTRA", err)
		}
	}
	if scg_CellGroupConfigEUTRAPresent {
		var tmp_os_scg_CellGroupConfigEUTRA []byte
		if tmp_os_scg_CellGroupConfigEUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode scg_CellGroupConfigEUTRA", err)
		}
		ie.scg_CellGroupConfigEUTRA = &tmp_os_scg_CellGroupConfigEUTRA
	}
	if candidateCellInfoListSN_EUTRAPresent {
		var tmp_os_candidateCellInfoListSN_EUTRA []byte
		if tmp_os_candidateCellInfoListSN_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode candidateCellInfoListSN_EUTRA", err)
		}
		ie.candidateCellInfoListSN_EUTRA = &tmp_os_candidateCellInfoListSN_EUTRA
	}
	if candidateServingFreqListEUTRAPresent {
		ie.candidateServingFreqListEUTRA = new(CandidateServingFreqListEUTRA)
		if err = ie.candidateServingFreqListEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode candidateServingFreqListEUTRA", err)
		}
	}
	if needForGapsPresent {
		ie.needForGaps = new(CG_Config_v1560_IEs_needForGaps)
		if err = ie.needForGaps.Decode(r); err != nil {
			return utils.WrapError("Decode needForGaps", err)
		}
	}
	if drx_ConfigSCGPresent {
		ie.drx_ConfigSCG = new(DRX_Config)
		if err = ie.drx_ConfigSCG.Decode(r); err != nil {
			return utils.WrapError("Decode drx_ConfigSCG", err)
		}
	}
	if reportCGI_RequestEUTRAPresent {
		ie.reportCGI_RequestEUTRA = new(CG_Config_v1560_IEs_reportCGI_RequestEUTRA)
		if err = ie.reportCGI_RequestEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode reportCGI_RequestEUTRA", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1590_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
