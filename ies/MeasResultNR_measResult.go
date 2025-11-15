package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultNR_measResult struct {
	cellResults    *MeasResultNR_measResult_cellResults    `optional`
	rsIndexResults *MeasResultNR_measResult_rsIndexResults `optional`
}

func (ie *MeasResultNR_measResult) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellResults != nil, ie.rsIndexResults != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.cellResults != nil {
		if err = ie.cellResults.Encode(w); err != nil {
			return utils.WrapError("Encode cellResults", err)
		}
	}
	if ie.rsIndexResults != nil {
		if err = ie.rsIndexResults.Encode(w); err != nil {
			return utils.WrapError("Encode rsIndexResults", err)
		}
	}
	return nil
}

func (ie *MeasResultNR_measResult) Decode(r *uper.UperReader) error {
	var err error
	var cellResultsPresent bool
	if cellResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var rsIndexResultsPresent bool
	if rsIndexResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if cellResultsPresent {
		ie.cellResults = new(MeasResultNR_measResult_cellResults)
		if err = ie.cellResults.Decode(r); err != nil {
			return utils.WrapError("Decode cellResults", err)
		}
	}
	if rsIndexResultsPresent {
		ie.rsIndexResults = new(MeasResultNR_measResult_rsIndexResults)
		if err = ie.rsIndexResults.Decode(r); err != nil {
			return utils.WrapError("Decode rsIndexResults", err)
		}
	}
	return nil
}
