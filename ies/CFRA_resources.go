package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CFRA_resources_Choice_nothing uint64 = iota
	CFRA_resources_Choice_ssb
	CFRA_resources_Choice_csirs
)

type CFRA_resources struct {
	Choice uint64
	ssb    *CFRA_resources_ssb
	csirs  *CFRA_resources_csirs
}

func (ie *CFRA_resources) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CFRA_resources_Choice_ssb:
		if err = ie.ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb", err)
		}
	case CFRA_resources_Choice_csirs:
		if err = ie.csirs.Encode(w); err != nil {
			err = utils.WrapError("Encode csirs", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CFRA_resources) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CFRA_resources_Choice_ssb:
		ie.ssb = new(CFRA_resources_ssb)
		if err = ie.ssb.Decode(r); err != nil {
			return utils.WrapError("Decode ssb", err)
		}
	case CFRA_resources_Choice_csirs:
		ie.csirs = new(CFRA_resources_csirs)
		if err = ie.csirs.Decode(r); err != nil {
			return utils.WrapError("Decode csirs", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
