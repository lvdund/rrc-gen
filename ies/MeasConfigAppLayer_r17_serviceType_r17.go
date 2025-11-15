package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MeasConfigAppLayer_r17_serviceType_r17_Enum_streaming uper.Enumerated = 0
	MeasConfigAppLayer_r17_serviceType_r17_Enum_mtsi      uper.Enumerated = 1
	MeasConfigAppLayer_r17_serviceType_r17_Enum_vr        uper.Enumerated = 2
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare5    uper.Enumerated = 3
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare4    uper.Enumerated = 4
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare3    uper.Enumerated = 5
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare2    uper.Enumerated = 6
	MeasConfigAppLayer_r17_serviceType_r17_Enum_spare1    uper.Enumerated = 7
)

type MeasConfigAppLayer_r17_serviceType_r17 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MeasConfigAppLayer_r17_serviceType_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MeasConfigAppLayer_r17_serviceType_r17", err)
	}
	return nil
}

func (ie *MeasConfigAppLayer_r17_serviceType_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MeasConfigAppLayer_r17_serviceType_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
