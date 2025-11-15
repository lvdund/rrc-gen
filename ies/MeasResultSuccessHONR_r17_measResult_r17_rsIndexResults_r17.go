package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17 struct {
	resultsSSB_Indexes_r17    *ResultsPerSSB_IndexList    `optional`
	resultsCSI_RS_Indexes_r17 *ResultsPerCSI_RS_IndexList `optional`
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Indexes_r17 != nil, ie.resultsCSI_RS_Indexes_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Indexes_r17 != nil {
		if err = ie.resultsSSB_Indexes_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Indexes_r17", err)
		}
	}
	if ie.resultsCSI_RS_Indexes_r17 != nil {
		if err = ie.resultsCSI_RS_Indexes_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Indexes_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_rsIndexResults_r17) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_Indexes_r17Present bool
	if resultsSSB_Indexes_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_Indexes_r17Present bool
	if resultsCSI_RS_Indexes_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_Indexes_r17Present {
		ie.resultsSSB_Indexes_r17 = new(ResultsPerSSB_IndexList)
		if err = ie.resultsSSB_Indexes_r17.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Indexes_r17", err)
		}
	}
	if resultsCSI_RS_Indexes_r17Present {
		ie.resultsCSI_RS_Indexes_r17 = new(ResultsPerCSI_RS_IndexList)
		if err = ie.resultsCSI_RS_Indexes_r17.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Indexes_r17", err)
		}
	}
	return nil
}
