package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_Resource struct {
	pucch_ResourceId          PUCCH_ResourceId                          `madatory`
	startingPRB               PRB_Id                                    `madatory`
	intraSlotFrequencyHopping *PUCCH_Resource_intraSlotFrequencyHopping `optional`
	secondHopPRB              *PRB_Id                                   `optional`
	format                    PUCCH_Resource_format                     `madatory`
}

func (ie *PUCCH_Resource) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.intraSlotFrequencyHopping != nil, ie.secondHopPRB != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.pucch_ResourceId.Encode(w); err != nil {
		return utils.WrapError("Encode pucch_ResourceId", err)
	}
	if err = ie.startingPRB.Encode(w); err != nil {
		return utils.WrapError("Encode startingPRB", err)
	}
	if ie.intraSlotFrequencyHopping != nil {
		if err = ie.intraSlotFrequencyHopping.Encode(w); err != nil {
			return utils.WrapError("Encode intraSlotFrequencyHopping", err)
		}
	}
	if ie.secondHopPRB != nil {
		if err = ie.secondHopPRB.Encode(w); err != nil {
			return utils.WrapError("Encode secondHopPRB", err)
		}
	}
	if err = ie.format.Encode(w); err != nil {
		return utils.WrapError("Encode format", err)
	}
	return nil
}

func (ie *PUCCH_Resource) Decode(r *uper.UperReader) error {
	var err error
	var intraSlotFrequencyHoppingPresent bool
	if intraSlotFrequencyHoppingPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var secondHopPRBPresent bool
	if secondHopPRBPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.pucch_ResourceId.Decode(r); err != nil {
		return utils.WrapError("Decode pucch_ResourceId", err)
	}
	if err = ie.startingPRB.Decode(r); err != nil {
		return utils.WrapError("Decode startingPRB", err)
	}
	if intraSlotFrequencyHoppingPresent {
		ie.intraSlotFrequencyHopping = new(PUCCH_Resource_intraSlotFrequencyHopping)
		if err = ie.intraSlotFrequencyHopping.Decode(r); err != nil {
			return utils.WrapError("Decode intraSlotFrequencyHopping", err)
		}
	}
	if secondHopPRBPresent {
		ie.secondHopPRB = new(PRB_Id)
		if err = ie.secondHopPRB.Decode(r); err != nil {
			return utils.WrapError("Decode secondHopPRB", err)
		}
	}
	if err = ie.format.Decode(r); err != nil {
		return utils.WrapError("Decode format", err)
	}
	return nil
}
