package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SSB_MTC2_LP_r16_periodicity_Enum_sf10   uper.Enumerated = 0
	SSB_MTC2_LP_r16_periodicity_Enum_sf20   uper.Enumerated = 1
	SSB_MTC2_LP_r16_periodicity_Enum_sf40   uper.Enumerated = 2
	SSB_MTC2_LP_r16_periodicity_Enum_sf80   uper.Enumerated = 3
	SSB_MTC2_LP_r16_periodicity_Enum_sf160  uper.Enumerated = 4
	SSB_MTC2_LP_r16_periodicity_Enum_spare3 uper.Enumerated = 5
	SSB_MTC2_LP_r16_periodicity_Enum_spare2 uper.Enumerated = 6
	SSB_MTC2_LP_r16_periodicity_Enum_spare1 uper.Enumerated = 7
)

type SSB_MTC2_LP_r16_periodicity struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SSB_MTC2_LP_r16_periodicity) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SSB_MTC2_LP_r16_periodicity", err)
	}
	return nil
}

func (ie *SSB_MTC2_LP_r16_periodicity) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SSB_MTC2_LP_r16_periodicity", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
