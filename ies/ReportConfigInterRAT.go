package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigInterRAT struct {
	reportType ReportConfigInterRAT_reportType `madatory`
}

func (ie *ReportConfigInterRAT) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.reportType.Encode(w); err != nil {
		return utils.WrapError("Encode reportType", err)
	}
	return nil
}

func (ie *ReportConfigInterRAT) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.reportType.Decode(r); err != nil {
		return utils.WrapError("Decode reportType", err)
	}
	return nil
}
