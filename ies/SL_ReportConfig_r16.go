package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_ReportConfig_r16 struct {
	sl_ReportType_r16 SL_ReportConfig_r16_sl_ReportType_r16 `madatory`
}

func (ie *SL_ReportConfig_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.sl_ReportType_r16.Encode(w); err != nil {
		return utils.WrapError("Encode sl_ReportType_r16", err)
	}
	return nil
}

func (ie *SL_ReportConfig_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.sl_ReportType_r16.Decode(r); err != nil {
		return utils.WrapError("Decode sl_ReportType_r16", err)
	}
	return nil
}
