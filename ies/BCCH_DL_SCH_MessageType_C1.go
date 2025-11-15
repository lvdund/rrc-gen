package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BCCH_DL_SCH_MessageType_C1_Choice_nothing uint64 = iota
	BCCH_DL_SCH_MessageType_C1_Choice_systemInformation
	BCCH_DL_SCH_MessageType_C1_Choice_systemInformationBlockType1
)

type BCCH_DL_SCH_MessageType_C1 struct {
	Choice                      uint64
	systemInformation           *SystemInformation
	systemInformationBlockType1 *SIB1
}

func (ie *BCCH_DL_SCH_MessageType_C1) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_C1_Choice_systemInformation:
		if err = ie.systemInformation.Encode(w); err != nil {
			err = utils.WrapError("Encode systemInformation", err)
		}
	case BCCH_DL_SCH_MessageType_C1_Choice_systemInformationBlockType1:
		if err = ie.systemInformationBlockType1.Encode(w); err != nil {
			err = utils.WrapError("Encode systemInformationBlockType1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BCCH_DL_SCH_MessageType_C1) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BCCH_DL_SCH_MessageType_C1_Choice_systemInformation:
		ie.systemInformation = new(SystemInformation)
		if err = ie.systemInformation.Decode(r); err != nil {
			return utils.WrapError("Decode systemInformation", err)
		}
	case BCCH_DL_SCH_MessageType_C1_Choice_systemInformationBlockType1:
		ie.systemInformationBlockType1 = new(SIB1)
		if err = ie.systemInformationBlockType1.Decode(r); err != nil {
			return utils.WrapError("Decode systemInformationBlockType1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
