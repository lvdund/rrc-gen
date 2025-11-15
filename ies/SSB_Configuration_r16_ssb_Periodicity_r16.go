package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms5    uper.Enumerated = 0
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms10   uper.Enumerated = 1
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms20   uper.Enumerated = 2
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms40   uper.Enumerated = 3
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms80   uper.Enumerated = 4
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_ms160  uper.Enumerated = 5
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_spare2 uper.Enumerated = 6
	SSB_Configuration_r16_ssb_Periodicity_r16_Enum_spare1 uper.Enumerated = 7
)

type SSB_Configuration_r16_ssb_Periodicity_r16 struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SSB_Configuration_r16_ssb_Periodicity_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SSB_Configuration_r16_ssb_Periodicity_r16", err)
	}
	return nil
}

func (ie *SSB_Configuration_r16_ssb_Periodicity_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SSB_Configuration_r16_ssb_Periodicity_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
