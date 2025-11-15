package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RxTxPeriodical_r17_reportAmount_r17_Enum_r1       uper.Enumerated = 0
	RxTxPeriodical_r17_reportAmount_r17_Enum_infinity uper.Enumerated = 1
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare6   uper.Enumerated = 2
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare5   uper.Enumerated = 3
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare4   uper.Enumerated = 4
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare3   uper.Enumerated = 5
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare2   uper.Enumerated = 6
	RxTxPeriodical_r17_reportAmount_r17_Enum_spare1   uper.Enumerated = 7
)

type RxTxPeriodical_r17_reportAmount_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *RxTxPeriodical_r17_reportAmount_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode RxTxPeriodical_r17_reportAmount_r17", err)
	}
	return nil
}

func (ie *RxTxPeriodical_r17_reportAmount_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode RxTxPeriodical_r17_reportAmount_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
