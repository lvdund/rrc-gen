package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1540_IEs_measResultReportCGI struct {
	ssbFrequency            ARFCN_ValueNR `madatory`
	cellForWhichToReportCGI PhysCellId    `madatory`
	cgi_Info                CGI_InfoNR    `madatory`
}

func (ie *CG_ConfigInfo_v1540_IEs_measResultReportCGI) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssbFrequency.Encode(w); err != nil {
		return utils.WrapError("Encode ssbFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI.Encode(w); err != nil {
		return utils.WrapError("Encode cellForWhichToReportCGI", err)
	}
	if err = ie.cgi_Info.Encode(w); err != nil {
		return utils.WrapError("Encode cgi_Info", err)
	}
	return nil
}

func (ie *CG_ConfigInfo_v1540_IEs_measResultReportCGI) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssbFrequency.Decode(r); err != nil {
		return utils.WrapError("Decode ssbFrequency", err)
	}
	if err = ie.cellForWhichToReportCGI.Decode(r); err != nil {
		return utils.WrapError("Decode cellForWhichToReportCGI", err)
	}
	if err = ie.cgi_Info.Decode(r); err != nil {
		return utils.WrapError("Decode cgi_Info", err)
	}
	return nil
}
