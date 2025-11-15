package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRLFNR_r16_measResult_r16_cellResults_r16 struct {
	resultsSSB_Cell_r16    *MeasQuantityResults `optional`
	resultsCSI_RS_Cell_r16 *MeasQuantityResults `optional`
}

func (ie *MeasResultRLFNR_r16_measResult_r16_cellResults_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Cell_r16 != nil, ie.resultsCSI_RS_Cell_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Cell_r16 != nil {
		if err = ie.resultsSSB_Cell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Cell_r16", err)
		}
	}
	if ie.resultsCSI_RS_Cell_r16 != nil {
		if err = ie.resultsCSI_RS_Cell_r16.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Cell_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultRLFNR_r16_measResult_r16_cellResults_r16) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_Cell_r16Present bool
	if resultsSSB_Cell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_Cell_r16Present bool
	if resultsCSI_RS_Cell_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_Cell_r16Present {
		ie.resultsSSB_Cell_r16 = new(MeasQuantityResults)
		if err = ie.resultsSSB_Cell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Cell_r16", err)
		}
	}
	if resultsCSI_RS_Cell_r16Present {
		ie.resultsCSI_RS_Cell_r16 = new(MeasQuantityResults)
		if err = ie.resultsCSI_RS_Cell_r16.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Cell_r16", err)
		}
	}
	return nil
}
