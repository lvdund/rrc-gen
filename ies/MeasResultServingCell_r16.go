package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultServingCell_r16 struct {
	resultsSSB_Cell MeasQuantityResults                   `madatory`
	resultsSSB      *MeasResultServingCell_r16_resultsSSB `lb:1,ub:maxNrofSSBs_r16,optional`
}

func (ie *MeasResultServingCell_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.resultsSSB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.resultsSSB_Cell.Encode(w); err != nil {
		return utils.WrapError("Encode resultsSSB_Cell", err)
	}
	if ie.resultsSSB != nil {
		if err = ie.resultsSSB.Encode(w); err != nil {
			return utils.WrapError("Encode resultsSSB", err)
		}
	}
	return nil
}

func (ie *MeasResultServingCell_r16) Decode(r *uper.UperReader) error {
	var err error
	var resultsSSBPresent bool
	if resultsSSBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.resultsSSB_Cell.Decode(r); err != nil {
		return utils.WrapError("Decode resultsSSB_Cell", err)
	}
	if resultsSSBPresent {
		ie.resultsSSB = new(MeasResultServingCell_r16_resultsSSB)
		if err = ie.resultsSSB.Decode(r); err != nil {
			return utils.WrapError("Decode resultsSSB", err)
		}
	}
	return nil
}
