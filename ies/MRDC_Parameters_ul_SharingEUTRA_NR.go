package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_tdm  uper.Enumerated = 0
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_fdm  uper.Enumerated = 1
	MRDC_Parameters_ul_SharingEUTRA_NR_Enum_both uper.Enumerated = 2
)

type MRDC_Parameters_ul_SharingEUTRA_NR struct {
	Value uper.Enumerated `lb:0,ub:2,madatory`
}

func (ie *MRDC_Parameters_ul_SharingEUTRA_NR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Encode MRDC_Parameters_ul_SharingEUTRA_NR", err)
	}
	return nil
}

func (ie *MRDC_Parameters_ul_SharingEUTRA_NR) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 2}, false); err != nil {
		return utils.WrapError("Decode MRDC_Parameters_ul_SharingEUTRA_NR", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
