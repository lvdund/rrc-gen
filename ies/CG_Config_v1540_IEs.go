package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs struct {
	pSCellFrequency      *ARFCN_ValueNR                           `optional`
	reportCGI_RequestNR  *CG_Config_v1540_IEs_reportCGI_RequestNR `optional`
	ph_InfoSCG           *PH_TypeListSCG                          `optional`
	nonCriticalExtension *CG_Config_v1560_IEs                     `optional`
}

func (ie *CG_Config_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pSCellFrequency != nil, ie.reportCGI_RequestNR != nil, ie.ph_InfoSCG != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pSCellFrequency != nil {
		if err = ie.pSCellFrequency.Encode(w); err != nil {
			return utils.WrapError("Encode pSCellFrequency", err)
		}
	}
	if ie.reportCGI_RequestNR != nil {
		if err = ie.reportCGI_RequestNR.Encode(w); err != nil {
			return utils.WrapError("Encode reportCGI_RequestNR", err)
		}
	}
	if ie.ph_InfoSCG != nil {
		if err = ie.ph_InfoSCG.Encode(w); err != nil {
			return utils.WrapError("Encode ph_InfoSCG", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_Config_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var pSCellFrequencyPresent bool
	if pSCellFrequencyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var reportCGI_RequestNRPresent bool
	if reportCGI_RequestNRPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ph_InfoSCGPresent bool
	if ph_InfoSCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pSCellFrequencyPresent {
		ie.pSCellFrequency = new(ARFCN_ValueNR)
		if err = ie.pSCellFrequency.Decode(r); err != nil {
			return utils.WrapError("Decode pSCellFrequency", err)
		}
	}
	if reportCGI_RequestNRPresent {
		ie.reportCGI_RequestNR = new(CG_Config_v1540_IEs_reportCGI_RequestNR)
		if err = ie.reportCGI_RequestNR.Decode(r); err != nil {
			return utils.WrapError("Decode reportCGI_RequestNR", err)
		}
	}
	if ph_InfoSCGPresent {
		ie.ph_InfoSCG = new(PH_TypeListSCG)
		if err = ie.ph_InfoSCG.Decode(r); err != nil {
			return utils.WrapError("Decode ph_InfoSCG", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_Config_v1560_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
