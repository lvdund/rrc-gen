package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ProcessingParameters struct {
	fallback            ProcessingParameters_fallback             `madatory`
	differentTB_PerSlot *ProcessingParameters_differentTB_PerSlot `optional`
}

func (ie *ProcessingParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.differentTB_PerSlot != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.fallback.Encode(w); err != nil {
		return utils.WrapError("Encode fallback", err)
	}
	if ie.differentTB_PerSlot != nil {
		if err = ie.differentTB_PerSlot.Encode(w); err != nil {
			return utils.WrapError("Encode differentTB_PerSlot", err)
		}
	}
	return nil
}

func (ie *ProcessingParameters) Decode(r *uper.UperReader) error {
	var err error
	var differentTB_PerSlotPresent bool
	if differentTB_PerSlotPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.fallback.Decode(r); err != nil {
		return utils.WrapError("Decode fallback", err)
	}
	if differentTB_PerSlotPresent {
		ie.differentTB_PerSlot = new(ProcessingParameters_differentTB_PerSlot)
		if err = ie.differentTB_PerSlot.Decode(r); err != nil {
			return utils.WrapError("Decode differentTB_PerSlot", err)
		}
	}
	return nil
}
