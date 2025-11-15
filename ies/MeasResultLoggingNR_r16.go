package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultLoggingNR_r16 struct {
	physCellId_r16      PhysCellId          `madatory`
	resultsSSB_Cell_r16 MeasQuantityResults `madatory`
	numberOfGoodSSB_r16 *int64              `lb:1,ub:maxNrofSSBs_r16,optional`
}

func (ie *MeasResultLoggingNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.numberOfGoodSSB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.physCellId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode physCellId_r16", err)
	}
	if err = ie.resultsSSB_Cell_r16.Encode(w); err != nil {
		return utils.WrapError("Encode resultsSSB_Cell_r16", err)
	}
	if ie.numberOfGoodSSB_r16 != nil {
		if err = w.WriteInteger(*ie.numberOfGoodSSB_r16, &uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
			return utils.WrapError("Encode numberOfGoodSSB_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultLoggingNR_r16) Decode(r *uper.UperReader) error {
	var err error
	var numberOfGoodSSB_r16Present bool
	if numberOfGoodSSB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.physCellId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode physCellId_r16", err)
	}
	if err = ie.resultsSSB_Cell_r16.Decode(r); err != nil {
		return utils.WrapError("Decode resultsSSB_Cell_r16", err)
	}
	if numberOfGoodSSB_r16Present {
		var tmp_int_numberOfGoodSSB_r16 int64
		if tmp_int_numberOfGoodSSB_r16, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofSSBs_r16}, false); err != nil {
			return utils.WrapError("Decode numberOfGoodSSB_r16", err)
		}
		ie.numberOfGoodSSB_r16 = &tmp_int_numberOfGoodSSB_r16
	}
	return nil
}
