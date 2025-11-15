package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportConfigId struct {
	Value uint64 `lb:1,ub:maxReportConfigId,madatory`
}

func (ie *ReportConfigId) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(int64(ie.Value), &uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false); err != nil {
		return utils.WrapError("Encode ReportConfigId", err)
	}
	return nil
}

func (ie *ReportConfigId) Decode(r *uper.UperReader) error {
	var err error
	var v int64
	if v, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxReportConfigId}, false); err != nil {
		return utils.WrapError("Decode ReportConfigId", err)
	}
	ie.Value = uint64(v)
	return nil
}
