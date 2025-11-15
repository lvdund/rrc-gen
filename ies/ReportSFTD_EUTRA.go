package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportSFTD_EUTRA struct {
	reportSFTD_Meas bool `madatory`
	reportRSRP      bool `madatory`
}

func (ie *ReportSFTD_EUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.reportSFTD_Meas); err != nil {
		return utils.WrapError("WriteBoolean reportSFTD_Meas", err)
	}
	if err = w.WriteBoolean(ie.reportRSRP); err != nil {
		return utils.WrapError("WriteBoolean reportRSRP", err)
	}
	return nil
}

func (ie *ReportSFTD_EUTRA) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_reportSFTD_Meas bool
	if tmp_bool_reportSFTD_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportSFTD_Meas", err)
	}
	ie.reportSFTD_Meas = tmp_bool_reportSFTD_Meas
	var tmp_bool_reportRSRP bool
	if tmp_bool_reportRSRP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportRSRP", err)
	}
	ie.reportRSRP = tmp_bool_reportRSRP
	return nil
}
