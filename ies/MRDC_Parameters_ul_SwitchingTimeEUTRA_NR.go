package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_ul_SwitchingTimeEUTRA_NR_Enum_type1 uper.Enumerated = 0
	MRDC_Parameters_ul_SwitchingTimeEUTRA_NR_Enum_type2 uper.Enumerated = 1
)

type MRDC_Parameters_ul_SwitchingTimeEUTRA_NR struct {
	Value uper.Enumerated `lb:0,ub:1,madatory`
}

func (ie *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_ul_SwitchingTimeEUTRA_NR", err)
	}
	return nil
}

func (ie *MRDC_Parameters_ul_SwitchingTimeEUTRA_NR) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 1}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_ul_SwitchingTimeEUTRA_NR", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
