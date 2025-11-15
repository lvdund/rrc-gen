package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportConfigId struct {
	Value uint64 `lb:0,ub:maxNrofCSI_ReportConfigurations_1,madatory`
}

func (ie *CSI_ReportConfigId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 0, Ub: maxNrofCSI_ReportConfigurations_1}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportConfigId", err)
	}
	return nil
}

func (ie *CSI_ReportConfigId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: maxNrofCSI_ReportConfigurations_1}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportConfigId", err)
	}
	ie.Value = uint64(v)
	return nil
}
