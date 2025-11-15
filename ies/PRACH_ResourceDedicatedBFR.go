package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PRACH_ResourceDedicatedBFR_Choice_nothing uint64 = iota
	PRACH_ResourceDedicatedBFR_Choice_ssb
	PRACH_ResourceDedicatedBFR_Choice_csi_RS
)

type PRACH_ResourceDedicatedBFR struct {
	Choice uint64
	ssb    *BFR_SSB_Resource
	csi_RS *BFR_CSIRS_Resource
}

func (ie *PRACH_ResourceDedicatedBFR) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRACH_ResourceDedicatedBFR_Choice_ssb:
		if err = ie.ssb.Encode(w); err != nil {
			err = utils.WrapError("Encode ssb", err)
		}
	case PRACH_ResourceDedicatedBFR_Choice_csi_RS:
		if err = ie.csi_RS.Encode(w); err != nil {
			err = utils.WrapError("Encode csi_RS", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *PRACH_ResourceDedicatedBFR) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case PRACH_ResourceDedicatedBFR_Choice_ssb:
		ie.ssb = new(BFR_SSB_Resource)
		if err = ie.ssb.Decode(r); err != nil {
			return utils.WrapError("Decode ssb", err)
		}
	case PRACH_ResourceDedicatedBFR_Choice_csi_RS:
		ie.csi_RS = new(BFR_CSIRS_Resource)
		if err = ie.csi_RS.Decode(r); err != nil {
			return utils.WrapError("Decode csi_RS", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
