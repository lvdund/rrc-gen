package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCP_ParametersMRDC struct {
	pdcp_DuplicationSplitSRB *PDCP_ParametersMRDC_pdcp_DuplicationSplitSRB `optional`
	pdcp_DuplicationSplitDRB *PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB `optional`
}

func (ie *PDCP_ParametersMRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pdcp_DuplicationSplitSRB != nil, ie.pdcp_DuplicationSplitDRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.pdcp_DuplicationSplitSRB != nil {
		if err = ie.pdcp_DuplicationSplitSRB.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_DuplicationSplitSRB", err)
		}
	}
	if ie.pdcp_DuplicationSplitDRB != nil {
		if err = ie.pdcp_DuplicationSplitDRB.Encode(w); err != nil {
			return utils.WrapError("Encode pdcp_DuplicationSplitDRB", err)
		}
	}
	return nil
}

func (ie *PDCP_ParametersMRDC) Decode(r *uper.UperReader) error {
	var err error
	var pdcp_DuplicationSplitSRBPresent bool
	if pdcp_DuplicationSplitSRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var pdcp_DuplicationSplitDRBPresent bool
	if pdcp_DuplicationSplitDRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if pdcp_DuplicationSplitSRBPresent {
		ie.pdcp_DuplicationSplitSRB = new(PDCP_ParametersMRDC_pdcp_DuplicationSplitSRB)
		if err = ie.pdcp_DuplicationSplitSRB.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_DuplicationSplitSRB", err)
		}
	}
	if pdcp_DuplicationSplitDRBPresent {
		ie.pdcp_DuplicationSplitDRB = new(PDCP_ParametersMRDC_pdcp_DuplicationSplitDRB)
		if err = ie.pdcp_DuplicationSplitDRB.Decode(r); err != nil {
			return utils.WrapError("Decode pdcp_DuplicationSplitDRB", err)
		}
	}
	return nil
}
