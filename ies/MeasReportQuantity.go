package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantity struct {
	rsrp bool `madatory`
	rsrq bool `madatory`
	sinr bool `madatory`
}

func (ie *MeasReportQuantity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.rsrp); err != nil {
		return utils.WrapError("WriteBoolean rsrp", err)
	}
	if err = w.WriteBoolean(ie.rsrq); err != nil {
		return utils.WrapError("WriteBoolean rsrq", err)
	}
	if err = w.WriteBoolean(ie.sinr); err != nil {
		return utils.WrapError("WriteBoolean sinr", err)
	}
	return nil
}

func (ie *MeasReportQuantity) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_rsrp bool
	if tmp_bool_rsrp, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean rsrp", err)
	}
	ie.rsrp = tmp_bool_rsrp
	var tmp_bool_rsrq bool
	if tmp_bool_rsrq, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean rsrq", err)
	}
	ie.rsrq = tmp_bool_rsrq
	var tmp_bool_sinr bool
	if tmp_bool_sinr, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean sinr", err)
	}
	ie.sinr = tmp_bool_sinr
	return nil
}
