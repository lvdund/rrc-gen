package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n0     uper.Enumerated = 0
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n0dot5 uper.Enumerated = 1
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n1     uper.Enumerated = 2
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n1dot5 uper.Enumerated = 3
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n2     uper.Enumerated = 4
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n2dot5 uper.Enumerated = 5
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n3     uper.Enumerated = 6
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n3dot5 uper.Enumerated = 7
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n4     uper.Enumerated = 8
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n4dot5 uper.Enumerated = 9
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n5     uper.Enumerated = 10
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n5dot5 uper.Enumerated = 11
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n6     uper.Enumerated = 12
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n6dot5 uper.Enumerated = 13
	SRS_SwitchingTimeEUTRA_switchingTimeUL_Enum_n7     uper.Enumerated = 14
)

type SRS_SwitchingTimeEUTRA_switchingTimeUL struct {
	Value uper.Enumerated `lb:0,ub:14,madatory`
}

func (ie *SRS_SwitchingTimeEUTRA_switchingTimeUL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Encode SRS_SwitchingTimeEUTRA_switchingTimeUL", err)
	}
	return nil
}

func (ie *SRS_SwitchingTimeEUTRA_switchingTimeUL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 14}, false); err != nil {
		return utils.WrapError("Decode SRS_SwitchingTimeEUTRA_switchingTimeUL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
