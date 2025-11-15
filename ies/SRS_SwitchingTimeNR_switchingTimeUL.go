package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n0us   uper.Enumerated = 0
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n30us  uper.Enumerated = 1
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n100us uper.Enumerated = 2
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n140us uper.Enumerated = 3
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n200us uper.Enumerated = 4
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n300us uper.Enumerated = 5
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n500us uper.Enumerated = 6
	SRS_SwitchingTimeNR_switchingTimeUL_Enum_n900us uper.Enumerated = 7
)

type SRS_SwitchingTimeNR_switchingTimeUL struct {
	Value uper.Enumerated `lb:0,ub:7,madatory`
}

func (ie *SRS_SwitchingTimeNR_switchingTimeUL) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Encode SRS_SwitchingTimeNR_switchingTimeUL", err)
	}
	return nil
}

func (ie *SRS_SwitchingTimeNR_switchingTimeUL) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 7}, false); err != nil {
		return utils.WrapError("Decode SRS_SwitchingTimeNR_switchingTimeUL", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
