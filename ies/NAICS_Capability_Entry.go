package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NAICS_Capability_Entry struct {
	numberOfNAICS_CapableCC int64                                        `lb:1,ub:5,madatory`
	numberOfAggregatedPRB   NAICS_Capability_Entry_numberOfAggregatedPRB `madatory`
}

func (ie *NAICS_Capability_Entry) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.numberOfNAICS_CapableCC, &uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfNAICS_CapableCC", err)
	}
	if err = ie.numberOfAggregatedPRB.Encode(w); err != nil {
		return utils.WrapError("Encode numberOfAggregatedPRB", err)
	}
	return nil
}

func (ie *NAICS_Capability_Entry) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_numberOfNAICS_CapableCC int64
	if tmp_int_numberOfNAICS_CapableCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 5}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfNAICS_CapableCC", err)
	}
	ie.numberOfNAICS_CapableCC = tmp_int_numberOfNAICS_CapableCC
	if err = ie.numberOfAggregatedPRB.Decode(r); err != nil {
		return utils.WrapError("Decode numberOfAggregatedPRB", err)
	}
	return nil
}
