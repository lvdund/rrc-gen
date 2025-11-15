package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC3_r16_duration_r16_Enum_sf1 uper.Enumerated = 0
	SSB_MTC3_r16_duration_r16_Enum_sf2 uper.Enumerated = 1
	SSB_MTC3_r16_duration_r16_Enum_sf3 uper.Enumerated = 2
	SSB_MTC3_r16_duration_r16_Enum_sf4 uper.Enumerated = 3
	SSB_MTC3_r16_duration_r16_Enum_sf5 uper.Enumerated = 4
)

type SSB_MTC3_r16_duration_r16 struct {
	Value uper.Enumerated `lb:0,ub:4,madatory`
}

func (ie *SSB_MTC3_r16_duration_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Encode SSB_MTC3_r16_duration_r16", err)
	}
	return nil
}

func (ie *SSB_MTC3_r16_duration_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("Decode SSB_MTC3_r16_duration_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
