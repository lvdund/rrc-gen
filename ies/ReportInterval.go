package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportInterval_Enum_ms120   uper.Enumerated = 0
	ReportInterval_Enum_ms240   uper.Enumerated = 1
	ReportInterval_Enum_ms480   uper.Enumerated = 2
	ReportInterval_Enum_ms640   uper.Enumerated = 3
	ReportInterval_Enum_ms1024  uper.Enumerated = 4
	ReportInterval_Enum_ms2048  uper.Enumerated = 5
	ReportInterval_Enum_ms5120  uper.Enumerated = 6
	ReportInterval_Enum_ms10240 uper.Enumerated = 7
	ReportInterval_Enum_ms20480 uper.Enumerated = 8
	ReportInterval_Enum_ms40960 uper.Enumerated = 9
	ReportInterval_Enum_min1    uper.Enumerated = 10
	ReportInterval_Enum_min6    uper.Enumerated = 11
	ReportInterval_Enum_min12   uper.Enumerated = 12
	ReportInterval_Enum_min30   uper.Enumerated = 13
)

type ReportInterval struct {
	Value uper.Enumerated `lb:0,ub:13,madatory`
}

func (ie *ReportInterval) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Encode ReportInterval", err)
	}
	return nil
}

func (ie *ReportInterval) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 13}, false); err != nil {
		return utils.WrapError("Decode ReportInterval", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
