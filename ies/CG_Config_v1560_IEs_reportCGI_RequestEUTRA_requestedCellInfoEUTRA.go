package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA struct {
	eutraFrequency                ARFCN_ValueEUTRA `madatory`
	cellForWhichToReportCGI_EUTRA EUTRA_PhysCellId `madatory`
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.eutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode eutraFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode cellForWhichToReportCGI_EUTRA", err)
	}
	return nil
}

func (ie *CG_Config_v1560_IEs_reportCGI_RequestEUTRA_requestedCellInfoEUTRA) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.eutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode eutraFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode cellForWhichToReportCGI_EUTRA", err)
	}
	return nil
}
