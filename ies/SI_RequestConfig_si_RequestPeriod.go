package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SI_RequestConfig_si_RequestPeriod_Enum_one     uper.Enumerated = 0
	SI_RequestConfig_si_RequestPeriod_Enum_two     uper.Enumerated = 1
	SI_RequestConfig_si_RequestPeriod_Enum_four    uper.Enumerated = 2
	SI_RequestConfig_si_RequestPeriod_Enum_six     uper.Enumerated = 3
	SI_RequestConfig_si_RequestPeriod_Enum_eight   uper.Enumerated = 4
	SI_RequestConfig_si_RequestPeriod_Enum_ten     uper.Enumerated = 5
	SI_RequestConfig_si_RequestPeriod_Enum_twelve  uper.Enumerated = 6
	SI_RequestConfig_si_RequestPeriod_Enum_sixteen uper.Enumerated = 7
)

type SI_RequestConfig_si_RequestPeriod struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SI_RequestConfig_si_RequestPeriod) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SI_RequestConfig_si_RequestPeriod", err)
	}
	return nil
}

func (ie *SI_RequestConfig_si_RequestPeriod) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SI_RequestConfig_si_RequestPeriod", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
