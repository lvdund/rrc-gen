package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult_cellResults struct {
	resultsSSB_Cell    *MeasQuantityResults `optional`
	resultsCSI_RS_Cell *MeasQuantityResults `optional`
}

func (ie *MeasResultNR_measResult_cellResults) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Cell != nil, ie.resultsCSI_RS_Cell != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Cell != nil {
		if err = ie.resultsSSB_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Cell", err)
		}
	}
	if ie.resultsCSI_RS_Cell != nil {
		if err = ie.resultsCSI_RS_Cell.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Cell", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult_cellResults) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_CellPresent bool
	if resultsSSB_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_CellPresent bool
	if resultsCSI_RS_CellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_CellPresent {
		ie.resultsSSB_Cell = new(MeasQuantityResults)
		if err = ie.resultsSSB_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Cell", err)
		}
	}
	if resultsCSI_RS_CellPresent {
		ie.resultsCSI_RS_Cell = new(MeasQuantityResults)
		if err = ie.resultsCSI_RS_Cell.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Cell", err)
		}
	}
	return nil
}
