package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CA_BandwidthClassEUTRA_Enum_a uper.Enumerated = 0
	CA_BandwidthClassEUTRA_Enum_b uper.Enumerated = 1
	CA_BandwidthClassEUTRA_Enum_c uper.Enumerated = 2
	CA_BandwidthClassEUTRA_Enum_d uper.Enumerated = 3
	CA_BandwidthClassEUTRA_Enum_e uper.Enumerated = 4
	CA_BandwidthClassEUTRA_Enum_f uper.Enumerated = 5
)

type CA_BandwidthClassEUTRA struct {
	Value uper.Enumerated `lb:0,ub:5,madatory`
}

func (ie *CA_BandwidthClassEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Encode CA_BandwidthClassEUTRA", err)
	}
	return nil
}

func (ie *CA_BandwidthClassEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 5}, false); err != nil {
		return utils.WrapError("Decode CA_BandwidthClassEUTRA", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
