package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	SRS_Resource_srs_TCI_State_r17_Choice_nothing uint64 = iota
	SRS_Resource_srs_TCI_State_r17_Choice_srs_UL_TCI_State
	SRS_Resource_srs_TCI_State_r17_Choice_srs_DLorJointTCI_State
)

type SRS_Resource_srs_TCI_State_r17 struct {
	Choice                 uint64
	srs_UL_TCI_State       *TCI_UL_State_Id_r17
	srs_DLorJointTCI_State *TCI_StateId
}

func (ie *SRS_Resource_srs_TCI_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_srs_TCI_State_r17_Choice_srs_UL_TCI_State:
		if err = ie.srs_UL_TCI_State.Encode(w); err != nil {
			err = utils.WrapError("Encode srs_UL_TCI_State", err)
		}
	case SRS_Resource_srs_TCI_State_r17_Choice_srs_DLorJointTCI_State:
		if err = ie.srs_DLorJointTCI_State.Encode(w); err != nil {
			err = utils.WrapError("Encode srs_DLorJointTCI_State", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *SRS_Resource_srs_TCI_State_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case SRS_Resource_srs_TCI_State_r17_Choice_srs_UL_TCI_State:
		ie.srs_UL_TCI_State = new(TCI_UL_State_Id_r17)
		if err = ie.srs_UL_TCI_State.Decode(r); err != nil {
			return utils.WrapError("Decode srs_UL_TCI_State", err)
		}
	case SRS_Resource_srs_TCI_State_r17_Choice_srs_DLorJointTCI_State:
		ie.srs_DLorJointTCI_State = new(TCI_StateId)
		if err = ie.srs_DLorJointTCI_State.Decode(r); err != nil {
			return utils.WrapError("Decode srs_DLorJointTCI_State", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
