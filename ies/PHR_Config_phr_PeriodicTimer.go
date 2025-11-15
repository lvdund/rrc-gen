package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PHR_Config_phr_PeriodicTimer_Enum_sf10     uper.Enumerated = 0
	PHR_Config_phr_PeriodicTimer_Enum_sf20     uper.Enumerated = 1
	PHR_Config_phr_PeriodicTimer_Enum_sf50     uper.Enumerated = 2
	PHR_Config_phr_PeriodicTimer_Enum_sf100    uper.Enumerated = 3
	PHR_Config_phr_PeriodicTimer_Enum_sf200    uper.Enumerated = 4
	PHR_Config_phr_PeriodicTimer_Enum_sf500    uper.Enumerated = 5
	PHR_Config_phr_PeriodicTimer_Enum_sf1000   uper.Enumerated = 6
	PHR_Config_phr_PeriodicTimer_Enum_infinity uper.Enumerated = 7
)

type PHR_Config_phr_PeriodicTimer struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *PHR_Config_phr_PeriodicTimer) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode PHR_Config_phr_PeriodicTimer", err)
	}
	return nil
}

func (ie *PHR_Config_phr_PeriodicTimer) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode PHR_Config_phr_PeriodicTimer", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
