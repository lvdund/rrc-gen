package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult_rsIndexResults struct {
	resultsSSB_Indexes    *ResultsPerSSB_IndexList    `optional`
	resultsCSI_RS_Indexes *ResultsPerCSI_RS_IndexList `optional`
}

func (ie *MeasResultNR_measResult_rsIndexResults) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Indexes != nil, ie.resultsCSI_RS_Indexes != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Indexes != nil {
		if err = ie.resultsSSB_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Indexes", err)
		}
	}
	if ie.resultsCSI_RS_Indexes != nil {
		if err = ie.resultsCSI_RS_Indexes.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Indexes", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult_rsIndexResults) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_IndexesPresent bool
	if resultsSSB_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_IndexesPresent bool
	if resultsCSI_RS_IndexesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_IndexesPresent {
		ie.resultsSSB_Indexes = new(ResultsPerSSB_IndexList)
		if err = ie.resultsSSB_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Indexes", err)
		}
	}
	if resultsCSI_RS_IndexesPresent {
		ie.resultsCSI_RS_Indexes = new(ResultsPerCSI_RS_IndexList)
		if err = ie.resultsCSI_RS_Indexes.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Indexes", err)
		}
	}
	return nil
}
