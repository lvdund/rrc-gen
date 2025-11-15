package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_nothing uint64 = iota
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs15_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs30_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs60_r17
	NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs120_r17
)

type NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17 struct {
	Choice     uint64
	scs15_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs15_r17
	scs30_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17
	scs60_r17  *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs60_r17
	scs120_r17 *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs15_r17:
		if err = ie.scs15_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode scs15_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs30_r17:
		if err = ie.scs30_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode scs30_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs60_r17:
		if err = ie.scs60_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode scs60_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs120_r17:
		if err = ie.scs120_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode scs120_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs15_r17:
		ie.scs15_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs15_r17)
		if err = ie.scs15_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs15_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs30_r17:
		ie.scs30_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs30_r17)
		if err = ie.scs30_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs30_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs60_r17:
		ie.scs60_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs60_r17)
		if err = ie.scs60_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs60_r17", err)
		}
	case NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_Choice_scs120_r17:
		ie.scs120_r17 = new(NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17_scs120_r17)
		if err = ie.scs120_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs120_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
