package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB_Enum_supported uper.Enumerated = 0
)

type NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB", err)
	}
	return nil
}

func (ie *NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode NRDC_Parameters_v15c0_pdcp_DuplicationSplitSRB", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
