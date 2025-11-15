package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16 struct {
	rsrp_Result_r16        *RSRP_Range                  `optional`
	rsrq_Result_r16        *RSRQ_Range                  `optional`
	resultsSSB_Indexes_r16 *ResultsPerSSB_IndexList_r16 `optional`
}

func (ie *MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rsrp_Result_r16 != nil, ie.rsrq_Result_r16 != nil, ie.resultsSSB_Indexes_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rsrp_Result_r16 != nil {
		if err = ie.rsrp_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsrp_Result_r16", err)
		}
	}
	if ie.rsrq_Result_r16 != nil {
		if err = ie.rsrq_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rsrq_Result_r16", err)
		}
	}
	if ie.resultsSSB_Indexes_r16 != nil {
		if err = ie.resultsSSB_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Indexes_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var rsrp_Result_r16Present bool
	if rsrp_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var rsrq_Result_r16Present bool
	if rsrq_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsSSB_Indexes_r16Present bool
	if resultsSSB_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rsrp_Result_r16Present {
		ie.rsrp_Result_r16 = new(RSRP_Range)
		if err = ie.rsrp_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsrp_Result_r16", err)
		}
	}
	if rsrq_Result_r16Present {
		ie.rsrq_Result_r16 = new(RSRQ_Range)
		if err = ie.rsrq_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rsrq_Result_r16", err)
		}
	}
	if resultsSSB_Indexes_r16Present {
		ie.resultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList_r16)
		if err = ie.resultsSSB_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Indexes_r16", err)
		}
	}
	return nil
}
