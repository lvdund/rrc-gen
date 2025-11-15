package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB_Enum_supported uper.Enumerated = 0
)

type PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB struct {
	Value uper.Enumerated `lb:0,ub:0,madatory`
}

func (ie *PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Encode PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB", err)
	}
	return nil
}

func (ie *PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("Decode PDCP_Parameters_pdcp_DuplicationMCG_OrSCG_DRB", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
