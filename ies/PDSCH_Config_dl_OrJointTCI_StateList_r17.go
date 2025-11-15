package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_nothing uint64 = iota
	PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_explicitlist
	PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_unifiedTCI_StateRef_r17
)

type PDSCH_Config_dl_OrJointTCI_StateList_r17 struct {
	Choice                  uint64
	explicitlist            *PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist
	unifiedTCI_StateRef_r17 *ServingCellAndBWP_Id_r17
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_explicitlist:
		if err = ie.explicitlist.Encode(w); err != nil {
			err = utils.WrapError("Encode explicitlist", err)
		}
	case PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_unifiedTCI_StateRef_r17:
		if err = ie.unifiedTCI_StateRef_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode unifiedTCI_StateRef_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDSCH_Config_dl_OrJointTCI_StateList_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_explicitlist:
		ie.explicitlist = new(PDSCH_Config_dl_OrJointTCI_StateList_r17_explicitlist)
		if err = ie.explicitlist.Decode(r); err != nil {
			return utils.WrapError("Decode explicitlist", err)
		}
	case PDSCH_Config_dl_OrJointTCI_StateList_r17_Choice_unifiedTCI_StateRef_r17:
		ie.unifiedTCI_StateRef_r17 = new(ServingCellAndBWP_Id_r17)
		if err = ie.unifiedTCI_StateRef_r17.Decode(r); err != nil {
			return utils.WrapError("Decode unifiedTCI_StateRef_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
