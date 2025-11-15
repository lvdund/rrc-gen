package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs struct {
	candidateCellInfoListMN_EUTRA *[]byte                                            `optional`
	candidateCellInfoListSN_EUTRA *[]byte                                            `optional`
	sourceConfigSCG_EUTRA         *[]byte                                            `optional`
	scgFailureInfoEUTRA           *CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA       `optional`
	drx_ConfigMCG                 *DRX_Config                                        `optional`
	measResultReportCGI_EUTRA     *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA `optional`
	measResultCellListSFTD_EUTRA  *MeasResultCellListSFTD_EUTRA                      `optional`
	fr_InfoListMCG                *FR_InfoList                                       `optional`
	nonCriticalExtension          *CG_ConfigInfo_v1570_IEs                           `optional`
}

func (ie *CG_ConfigInfo_v1560_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.candidateCellInfoListMN_EUTRA != nil, ie.candidateCellInfoListSN_EUTRA != nil, ie.sourceConfigSCG_EUTRA != nil, ie.scgFailureInfoEUTRA != nil, ie.drx_ConfigMCG != nil, ie.measResultReportCGI_EUTRA != nil, ie.measResultCellListSFTD_EUTRA != nil, ie.fr_InfoListMCG != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.candidateCellInfoListMN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.candidateCellInfoListMN_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode candidateCellInfoListMN_EUTRA", err)
		}
	}
	if ie.candidateCellInfoListSN_EUTRA != nil {
		if err = w.WriteOctetString(*ie.candidateCellInfoListSN_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode candidateCellInfoListSN_EUTRA", err)
		}
	}
	if ie.sourceConfigSCG_EUTRA != nil {
		if err = w.WriteOctetString(*ie.sourceConfigSCG_EUTRA, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode sourceConfigSCG_EUTRA", err)
		}
	}
	if ie.scgFailureInfoEUTRA != nil {
		if err = ie.scgFailureInfoEUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode scgFailureInfoEUTRA", err)
		}
	}
	if ie.drx_ConfigMCG != nil {
		if err = ie.drx_ConfigMCG.Encode(w); err != nil {
			return utils.WrapError("Encode drx_ConfigMCG", err)
		}
	}
	if ie.measResultReportCGI_EUTRA != nil {
		if err = ie.measResultReportCGI_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode measResultReportCGI_EUTRA", err)
		}
	}
	if ie.measResultCellListSFTD_EUTRA != nil {
		if err = ie.measResultCellListSFTD_EUTRA.Encode(w); err != nil {
			return utils.WrapError("Encode measResultCellListSFTD_EUTRA", err)
		}
	}
	if ie.fr_InfoListMCG != nil {
		if err = ie.fr_InfoListMCG.Encode(w); err != nil {
			return utils.WrapError("Encode fr_InfoListMCG", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs) Decode(r *uper.UperReader) error {
	var err error
	var candidateCellInfoListMN_EUTRAPresent bool
	if candidateCellInfoListMN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var candidateCellInfoListSN_EUTRAPresent bool
	if candidateCellInfoListSN_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var sourceConfigSCG_EUTRAPresent bool
	if sourceConfigSCG_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scgFailureInfoEUTRAPresent bool
	if scgFailureInfoEUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var drx_ConfigMCGPresent bool
	if drx_ConfigMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultReportCGI_EUTRAPresent bool
	if measResultReportCGI_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultCellListSFTD_EUTRAPresent bool
	if measResultCellListSFTD_EUTRAPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var fr_InfoListMCGPresent bool
	if fr_InfoListMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if candidateCellInfoListMN_EUTRAPresent {
		var tmp_os_candidateCellInfoListMN_EUTRA []byte
		if tmp_os_candidateCellInfoListMN_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode candidateCellInfoListMN_EUTRA", err)
		}
		ie.candidateCellInfoListMN_EUTRA = &tmp_os_candidateCellInfoListMN_EUTRA
	}
	if candidateCellInfoListSN_EUTRAPresent {
		var tmp_os_candidateCellInfoListSN_EUTRA []byte
		if tmp_os_candidateCellInfoListSN_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode candidateCellInfoListSN_EUTRA", err)
		}
		ie.candidateCellInfoListSN_EUTRA = &tmp_os_candidateCellInfoListSN_EUTRA
	}
	if sourceConfigSCG_EUTRAPresent {
		var tmp_os_sourceConfigSCG_EUTRA []byte
		if tmp_os_sourceConfigSCG_EUTRA, err = r.ReadOctetString(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode sourceConfigSCG_EUTRA", err)
		}
		ie.sourceConfigSCG_EUTRA = &tmp_os_sourceConfigSCG_EUTRA
	}
	if scgFailureInfoEUTRAPresent {
		ie.scgFailureInfoEUTRA = new(CG_ConfigInfo_v1560_IEs_scgFailureInfoEUTRA)
		if err = ie.scgFailureInfoEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode scgFailureInfoEUTRA", err)
		}
	}
	if drx_ConfigMCGPresent {
		ie.drx_ConfigMCG = new(DRX_Config)
		if err = ie.drx_ConfigMCG.Decode(r); err != nil {
			return utils.WrapError("Decode drx_ConfigMCG", err)
		}
	}
	if measResultReportCGI_EUTRAPresent {
		ie.measResultReportCGI_EUTRA = new(CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA)
		if err = ie.measResultReportCGI_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode measResultReportCGI_EUTRA", err)
		}
	}
	if measResultCellListSFTD_EUTRAPresent {
		ie.measResultCellListSFTD_EUTRA = new(MeasResultCellListSFTD_EUTRA)
		if err = ie.measResultCellListSFTD_EUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode measResultCellListSFTD_EUTRA", err)
		}
	}
	if fr_InfoListMCGPresent {
		ie.fr_InfoListMCG = new(FR_InfoList)
		if err = ie.fr_InfoListMCG.Decode(r); err != nil {
			return utils.WrapError("Decode fr_InfoListMCG", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1570_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
