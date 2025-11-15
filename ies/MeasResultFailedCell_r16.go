package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultFailedCell_r16 struct {
	cgi_Info       CGI_Info_Logging_r16                    `madatory`
	measResult_r16 MeasResultFailedCell_r16_measResult_r16 `madatory`
}

func (ie *MeasResultFailedCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.cgi_Info.Encode(w); err != nil {
		return utils.WrapError("Encode cgi_Info", err)
	}
	if err = ie.measResult_r16.Encode(w); err != nil {
		return utils.WrapError("Encode measResult_r16", err)
	}
	return nil
}

func (ie *MeasResultFailedCell_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.cgi_Info.Decode(r); err != nil {
		return utils.WrapError("Decode cgi_Info", err)
	}
	if err = ie.measResult_r16.Decode(r); err != nil {
		return utils.WrapError("Decode measResult_r16", err)
	}
	return nil
}
