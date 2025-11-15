package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantity_r16 struct {
	cbr_r16 bool `madatory`
}

func (ie *MeasReportQuantity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.cbr_r16); err != nil {
		return utils.WrapError("WriteBoolean cbr_r16", err)
	}
	return nil
}

func (ie *MeasReportQuantity_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_cbr_r16 bool
	if tmp_bool_cbr_r16, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean cbr_r16", err)
	}
	ie.cbr_r16 = tmp_bool_cbr_r16
	return nil
}
