package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasObjectNR_measCycleSCell_Enum_sf160  uper.Enumerated = 0
	MeasObjectNR_measCycleSCell_Enum_sf256  uper.Enumerated = 1
	MeasObjectNR_measCycleSCell_Enum_sf320  uper.Enumerated = 2
	MeasObjectNR_measCycleSCell_Enum_sf512  uper.Enumerated = 3
	MeasObjectNR_measCycleSCell_Enum_sf640  uper.Enumerated = 4
	MeasObjectNR_measCycleSCell_Enum_sf1024 uper.Enumerated = 5
	MeasObjectNR_measCycleSCell_Enum_sf1280 uper.Enumerated = 6
)

type MeasObjectNR_measCycleSCell struct {
	Value uper.Enumerated `lb:0,ub:6,madatory`
}

func (ie *MeasObjectNR_measCycleSCell) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Encode MeasObjectNR_measCycleSCell", err)
	}
	return nil
}

func (ie *MeasObjectNR_measCycleSCell) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
		return utils.WrapError("Decode MeasObjectNR_measCycleSCell", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
