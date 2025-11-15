package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m20    uper.Enumerated = 0
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m50    uper.Enumerated = 1
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m80    uper.Enumerated = 2
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m100   uper.Enumerated = 3
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m120   uper.Enumerated = 4
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m150   uper.Enumerated = 5
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m180   uper.Enumerated = 6
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m200   uper.Enumerated = 7
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m220   uper.Enumerated = 8
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m250   uper.Enumerated = 9
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m270   uper.Enumerated = 10
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m300   uper.Enumerated = 11
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m350   uper.Enumerated = 12
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m370   uper.Enumerated = 13
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m400   uper.Enumerated = 14
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m420   uper.Enumerated = 15
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m450   uper.Enumerated = 16
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m480   uper.Enumerated = 17
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m500   uper.Enumerated = 18
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m550   uper.Enumerated = 19
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m600   uper.Enumerated = 20
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m700   uper.Enumerated = 21
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_m1000  uper.Enumerated = 22
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare9 uper.Enumerated = 23
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare8 uper.Enumerated = 24
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare7 uper.Enumerated = 25
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare6 uper.Enumerated = 26
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare5 uper.Enumerated = 27
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare4 uper.Enumerated = 28
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare3 uper.Enumerated = 29
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare2 uper.Enumerated = 30
	SL_ZoneConfigMCR_r16_sl_TransRange_r16_Enum_spare1 uper.Enumerated = 31
)

type SL_ZoneConfigMCR_r16_sl_TransRange_r16 struct {
	Value uper.Enumerated `lb:0,ub:31,madatory`
}

func (ie *SL_ZoneConfigMCR_r16_sl_TransRange_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Encode SL_ZoneConfigMCR_r16_sl_TransRange_r16", err)
	}
	return nil
}

func (ie *SL_ZoneConfigMCR_r16_sl_TransRange_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
		return utils.WrapError("Decode SL_ZoneConfigMCR_r16_sl_TransRange_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
