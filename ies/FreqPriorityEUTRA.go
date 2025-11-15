package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityEUTRA struct {
	carrierFreq                ARFCN_ValueEUTRA            `madatory`
	cellReselectionPriority    CellReselectionPriority     `madatory`
	cellReselectionSubPriority *CellReselectionSubPriority `optional`
}

func (ie *FreqPriorityEUTRA) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.cellReselectionSubPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq", err)
	}
	if err = ie.cellReselectionPriority.Encode(w); err != nil {
		return utils.WrapError("Encode cellReselectionPriority", err)
	}
	if ie.cellReselectionSubPriority != nil {
		if err = ie.cellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode cellReselectionSubPriority", err)
		}
	}
	return nil
}

func (ie *FreqPriorityEUTRA) Decode(r *uper.UperReader) error {
	var err error
	var cellReselectionSubPriorityPresent bool
	if cellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq", err)
	}
	if err = ie.cellReselectionPriority.Decode(r); err != nil {
		return utils.WrapError("Decode cellReselectionPriority", err)
	}
	if cellReselectionSubPriorityPresent {
		ie.cellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.cellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode cellReselectionSubPriority", err)
		}
	}
	return nil
}
