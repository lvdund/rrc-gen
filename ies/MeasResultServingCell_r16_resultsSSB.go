package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServingCell_r16_resultsSSB struct {
	best_ssb_Index   SSB_Index           `madatory`
	best_ssb_Results MeasQuantityResults `madatory`
	numberOfGoodSSB  int64               `lb:1,ub:maxNrofSSBs_r16,madatory`
}

func (ie *MeasResultServingCell_r16_resultsSSB) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.best_ssb_Index.Encode(w); err != nil {
		return utils.WrapError("Encode best_ssb_Index", err)
	}
	if err = ie.best_ssb_Results.Encode(w); err != nil {
		return utils.WrapError("Encode best_ssb_Results", err)
	}
	if err = w.WriteInteger(ie.numberOfGoodSSB, &uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
		return utils.WrapError("WriteInteger numberOfGoodSSB", err)
	}
	return nil
}

func (ie *MeasResultServingCell_r16_resultsSSB) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.best_ssb_Index.Decode(r); err != nil {
		return utils.WrapError("Decode best_ssb_Index", err)
	}
	if err = ie.best_ssb_Results.Decode(r); err != nil {
		return utils.WrapError("Decode best_ssb_Results", err)
	}
	var tmp_int_numberOfGoodSSB int64
	if tmp_int_numberOfGoodSSB, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
		return utils.WrapError("ReadInteger numberOfGoodSSB", err)
	}
	ie.numberOfGoodSSB = tmp_int_numberOfGoodSSB
	return nil
}
