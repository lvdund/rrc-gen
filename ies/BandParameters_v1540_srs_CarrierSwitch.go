package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	BandParameters_v1540_srs_CarrierSwitch_Choice_nothing uint64 = iota
	BandParameters_v1540_srs_CarrierSwitch_Choice_nr
	BandParameters_v1540_srs_CarrierSwitch_Choice_eutra
)

type BandParameters_v1540_srs_CarrierSwitch struct {
	Choice uint64
	nr     *BandParameters_v1540_srs_CarrierSwitch_nr
	eutra  *BandParameters_v1540_srs_CarrierSwitch_eutra
}

func (ie *BandParameters_v1540_srs_CarrierSwitch) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_v1540_srs_CarrierSwitch_Choice_nr:
		if err = ie.nr.Encode(w); err != nil {
			err = utils.WrapError("Encode nr", err)
		}
	case BandParameters_v1540_srs_CarrierSwitch_Choice_eutra:
		if err = ie.eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *BandParameters_v1540_srs_CarrierSwitch) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case BandParameters_v1540_srs_CarrierSwitch_Choice_nr:
		ie.nr = new(BandParameters_v1540_srs_CarrierSwitch_nr)
		if err = ie.nr.Decode(r); err != nil {
			return utils.WrapError("Decode nr", err)
		}
	case BandParameters_v1540_srs_CarrierSwitch_Choice_eutra:
		ie.eutra = new(BandParameters_v1540_srs_CarrierSwitch_eutra)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
