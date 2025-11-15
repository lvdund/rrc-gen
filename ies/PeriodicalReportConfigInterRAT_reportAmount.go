package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r1       uper.Enumerated = 0
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r2       uper.Enumerated = 1
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r4       uper.Enumerated = 2
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r8       uper.Enumerated = 3
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r16      uper.Enumerated = 4
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r32      uper.Enumerated = 5
	PeriodicalReportConfigInterRAT_reportAmount_Enum_r64      uper.Enumerated = 6
	PeriodicalReportConfigInterRAT_reportAmount_Enum_infinity uper.Enumerated = 7
)

type PeriodicalReportConfigInterRAT_reportAmount struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PeriodicalReportConfigInterRAT_reportAmount) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PeriodicalReportConfigInterRAT_reportAmount", err)
	}
	return nil
}

func (ie *PeriodicalReportConfigInterRAT_reportAmount) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PeriodicalReportConfigInterRAT_reportAmount", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
