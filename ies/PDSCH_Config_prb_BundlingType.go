package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_prb_BundlingType_Choice_nothing uint64 = iota
	PDSCH_Config_prb_BundlingType_Choice_staticBundling
	PDSCH_Config_prb_BundlingType_Choice_dynamicBundling
)

type PDSCH_Config_prb_BundlingType struct {
	Choice          uint64
	staticBundling  *PDSCH_Config_prb_BundlingType_staticBundling
	dynamicBundling *PDSCH_Config_prb_BundlingType_dynamicBundling
}

func (ie *PDSCH_Config_prb_BundlingType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingType_Choice_staticBundling:
		if err = ie.staticBundling.Encode(w); err != nil {
			err = utils.WrapError("Encode staticBundling", err)
		}
	case PDSCH_Config_prb_BundlingType_Choice_dynamicBundling:
		if err = ie.dynamicBundling.Encode(w); err != nil {
			err = utils.WrapError("Encode dynamicBundling", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PDSCH_Config_prb_BundlingType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PDSCH_Config_prb_BundlingType_Choice_staticBundling:
		ie.staticBundling = new(PDSCH_Config_prb_BundlingType_staticBundling)
		if err = ie.staticBundling.Decode(r); err != nil {
			return utils.WrapError("Decode staticBundling", err)
		}
	case PDSCH_Config_prb_BundlingType_Choice_dynamicBundling:
		ie.dynamicBundling = new(PDSCH_Config_prb_BundlingType_dynamicBundling)
		if err = ie.dynamicBundling.Decode(r); err != nil {
			return utils.WrapError("Decode dynamicBundling", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
