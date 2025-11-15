package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf0    uper.Enumerated = 0
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf10   uper.Enumerated = 1
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf20   uper.Enumerated = 2
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf50   uper.Enumerated = 3
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf100  uper.Enumerated = 4
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf200  uper.Enumerated = 5
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf500  uper.Enumerated = 6
	MPE_Config_FR2_r16_mpe_ProhibitTimer_r16_Enum_sf1000 uper.Enumerated = 7
)

type MPE_Config_FR2_r16_mpe_ProhibitTimer_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *MPE_Config_FR2_r16_mpe_ProhibitTimer_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode MPE_Config_FR2_r16_mpe_ProhibitTimer_r16", err)
	}
	return nil
}

func (ie *MPE_Config_FR2_r16_mpe_ProhibitTimer_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode MPE_Config_FR2_r16_mpe_ProhibitTimer_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
