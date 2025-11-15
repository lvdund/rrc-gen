package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA struct {
	eutraFrequency                ARFCN_ValueEUTRA `madatory`
	cellForWhichToReportCGI_EUTRA EUTRA_PhysCellId `madatory`
	cgi_InfoEUTRA                 CGI_InfoEUTRA    `madatory`
}

func (ie *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.eutraFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode eutraFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI_EUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode cellForWhichToReportCGI_EUTRA", err)
	}
	if err = ie.cgi_InfoEUTRA.Encode(w); err != nil {
		return utils.WrapError("Encode cgi_InfoEUTRA", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1560_IEs_measResultReportCGI_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.eutraFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode eutraFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI_EUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode cellForWhichToReportCGI_EUTRA", err)
	}
	if err = ie.cgi_InfoEUTRA.Decode(r); err != nil {
		return utils.WrapError("Decode cgi_InfoEUTRA", err)
	}
	return nil
}
