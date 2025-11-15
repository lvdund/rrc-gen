package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n3   uper.Enumerated = 0
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n7   uper.Enumerated = 1
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n15  uper.Enumerated = 2
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n31  uper.Enumerated = 3
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n63  uper.Enumerated = 4
	CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC_Enum_n128 uper.Enumerated = 5
)

type CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	return nil
}

func (ie *CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
