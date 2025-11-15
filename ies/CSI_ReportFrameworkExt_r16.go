package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportFrameworkExt_r16 struct {
	maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 int64 `lb:5,ub:8,madatory`
}

func (ie *CSI_ReportFrameworkExt_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16, &uper.Constraint{Lb: 5, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16", err)
	}
	return nil
}

func (ie *CSI_ReportFrameworkExt_r16) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 int64
	if tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16, err = r.ReadInteger(&uper.Constraint{Lb: 5, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16", err)
	}
	ie.maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16 = tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_ReportExt_r16
	return nil
}
