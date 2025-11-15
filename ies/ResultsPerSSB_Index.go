package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_Index struct {
	ssb_Index   SSB_Index            `madatory`
	ssb_Results *MeasQuantityResults `optional`
}

func (ie *ResultsPerSSB_Index) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ssb_Results != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ssb_Index.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_Index", err)
	}
	if ie.ssb_Results != nil {
		if err = ie.ssb_Results.Encode(w); err != nil {
			return utils.WrapError("Encode ssb_Results", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_Index) Decode(r *uper.UperReader) error {
	var err error
	var ssb_ResultsPresent bool
	if ssb_ResultsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ssb_Index.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_Index", err)
	}
	if ssb_ResultsPresent {
		ie.ssb_Results = new(MeasQuantityResults)
		if err = ie.ssb_Results.Decode(r); err != nil {
			return utils.WrapError("Decode ssb_Results", err)
		}
	}
	return nil
}
