package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasReportQuantityUTRA_FDD_r16 struct {
	cpich_RSCP bool `madatory`
	cpich_EcN0 bool `madatory`
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBoolean(ie.cpich_RSCP); err != nil {
		return utils.WrapError("WriteBoolean cpich_RSCP", err)
	}
	if err = w.WriteBoolean(ie.cpich_EcN0); err != nil {
		return utils.WrapError("WriteBoolean cpich_EcN0", err)
	}
	return nil
}

func (ie *MeasReportQuantityUTRA_FDD_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bool_cpich_RSCP bool
	if tmp_bool_cpich_RSCP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean cpich_RSCP", err)
	}
	ie.cpich_RSCP = tmp_bool_cpich_RSCP
	var tmp_bool_cpich_EcN0 bool
	if tmp_bool_cpich_EcN0, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean cpich_EcN0", err)
	}
	ie.cpich_EcN0 = tmp_bool_cpich_EcN0
	return nil
}
