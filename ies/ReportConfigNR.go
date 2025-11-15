package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigNR struct {
	reportType ReportConfigNR_reportType `madatory`
}

func (ie *ReportConfigNR) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportType.Encode(w); err != nil {
		return utils.WrapError("Encode reportType", err)
	}
	return nil
}

func (ie *ReportConfigNR) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportType.Decode(r); err != nil {
		return utils.WrapError("Decode reportType", err)
	}
	return nil
}
