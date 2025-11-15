package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1540_IEs struct {
	ph_InfoMCG           *PH_TypeListMCG                              `optional`
	measResultReportCGI  *CG_ConfigInfo_v1540_IEs_measResultReportCGI `optional`
	nonCriticalExtension *CG_ConfigInfo_v1560_IEs                     `optional`
}

func (ie *CG_ConfigInfo_v1540_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ph_InfoMCG != nil, ie.measResultReportCGI != nil, ie.nonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ph_InfoMCG != nil {
		if err = ie.ph_InfoMCG.Encode(w); err != nil {
			return utils.WrapError("Encode ph_InfoMCG", err)
		}
	}
	if ie.measResultReportCGI != nil {
		if err = ie.measResultReportCGI.Encode(w); err != nil {
			return utils.WrapError("Encode measResultReportCGI", err)
		}
	}
	if ie.nonCriticalExtension != nil {
		if err = ie.nonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode nonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1540_IEs) Decode(r *uper.UperReader) error {
	var err error
	var ph_InfoMCGPresent bool
	if ph_InfoMCGPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var measResultReportCGIPresent bool
	if measResultReportCGIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nonCriticalExtensionPresent bool
	if nonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if ph_InfoMCGPresent {
		ie.ph_InfoMCG = new(PH_TypeListMCG)
		if err = ie.ph_InfoMCG.Decode(r); err != nil {
			return utils.WrapError("Decode ph_InfoMCG", err)
		}
	}
	if measResultReportCGIPresent {
		ie.measResultReportCGI = new(CG_ConfigInfo_v1540_IEs_measResultReportCGI)
		if err = ie.measResultReportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode measResultReportCGI", err)
		}
	}
	if nonCriticalExtensionPresent {
		ie.nonCriticalExtension = new(CG_ConfigInfo_v1560_IEs)
		if err = ie.nonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode nonCriticalExtension", err)
		}
	}
	return nil
}
