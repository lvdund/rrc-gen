package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17 struct {
	resultsSSB_Cell_r17    *MeasQuantityResults `optional`
	resultsCSI_RS_Cell_r17 *MeasQuantityResults `optional`
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB_Cell_r17 != nil, ie.resultsCSI_RS_Cell_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.resultsSSB_Cell_r17 != nil {
		if err = ie.resultsSSB_Cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB_Cell_r17", err)
		}
	}
	if ie.resultsCSI_RS_Cell_r17 != nil {
		if err = ie.resultsCSI_RS_Cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode resultsCSI_RS_Cell_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultSuccessHONR_r17_measResult_r17_cellResults_r17) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSB_Cell_r17Present bool
	if resultsSSB_Cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var resultsCSI_RS_Cell_r17Present bool
	if resultsCSI_RS_Cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if resultsSSB_Cell_r17Present {
		ie.resultsSSB_Cell_r17 = new(MeasQuantityResults)
		if err = ie.resultsSSB_Cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB_Cell_r17", err)
		}
	}
	if resultsCSI_RS_Cell_r17Present {
		ie.resultsCSI_RS_Cell_r17 = new(MeasQuantityResults)
		if err = ie.resultsCSI_RS_Cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode resultsCSI_RS_Cell_r17", err)
		}
	}
	return nil
}
