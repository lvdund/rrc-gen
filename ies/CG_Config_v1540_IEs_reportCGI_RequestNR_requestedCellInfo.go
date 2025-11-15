package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo struct {
	ssbFrequency            ARFCN_ValueNR `madatory`
	cellForWhichToReportCGI PhysCellId    `madatory`
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssbFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode ssbFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI.Encode(w); err != nil {
		return utils.WrapError("Encode cellForWhichToReportCGI", err)
	}
	return nil
}

func (ie *CG_Config_v1540_IEs_reportCGI_RequestNR_requestedCellInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssbFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode ssbFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI.Decode(r); err != nil {
		return utils.WrapError("Decode cellForWhichToReportCGI", err)
	}
	return nil
}
