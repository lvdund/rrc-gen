package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_nothing uint64 = iota
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern0
	CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern1
)

type CSI_IM_Resource_csi_IM_ResourceElementPattern struct {
	Choice   uint64
	pattern0 *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0
	pattern1 *CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern0:
		if err = ie.pattern0.Encode(w); err != nil {
			err = utils.WrapError("Encode pattern0", err)
		}
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern1:
		if err = ie.pattern1.Encode(w); err != nil {
			err = utils.WrapError("Encode pattern1", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CSI_IM_Resource_csi_IM_ResourceElementPattern) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern0:
		ie.pattern0 = new(CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern0)
		if err = ie.pattern0.Decode(r); err != nil {
			return utils.WrapError("Decode pattern0", err)
		}
	case CSI_IM_Resource_csi_IM_ResourceElementPattern_Choice_pattern1:
		ie.pattern1 = new(CSI_IM_Resource_csi_IM_ResourceElementPattern_pattern1)
		if err = ie.pattern1.Decode(r); err != nil {
			return utils.WrapError("Decode pattern1", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
