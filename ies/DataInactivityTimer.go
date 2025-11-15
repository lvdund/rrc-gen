package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	DataInactivityTimer_Enum_s1   uper.Enumerated = 0
	DataInactivityTimer_Enum_s2   uper.Enumerated = 1
	DataInactivityTimer_Enum_s3   uper.Enumerated = 2
	DataInactivityTimer_Enum_s5   uper.Enumerated = 3
	DataInactivityTimer_Enum_s7   uper.Enumerated = 4
	DataInactivityTimer_Enum_s10  uper.Enumerated = 5
	DataInactivityTimer_Enum_s15  uper.Enumerated = 6
	DataInactivityTimer_Enum_s20  uper.Enumerated = 7
	DataInactivityTimer_Enum_s40  uper.Enumerated = 8
	DataInactivityTimer_Enum_s50  uper.Enumerated = 9
	DataInactivityTimer_Enum_s60  uper.Enumerated = 10
	DataInactivityTimer_Enum_s80  uper.Enumerated = 11
	DataInactivityTimer_Enum_s100 uper.Enumerated = 12
	DataInactivityTimer_Enum_s120 uper.Enumerated = 13
	DataInactivityTimer_Enum_s150 uper.Enumerated = 14
	DataInactivityTimer_Enum_s180 uper.Enumerated = 15
)

type DataInactivityTimer struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *DataInactivityTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode DataInactivityTimer", err)
	}
	return nil
}

func (ie *DataInactivityTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode DataInactivityTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
