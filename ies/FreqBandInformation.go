package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	FreqBandInformation_Choice_nothing uint64 = iota
	FreqBandInformation_Choice_bandInformationEUTRA
	FreqBandInformation_Choice_bandInformationNR
)

type FreqBandInformation struct {
	Choice               uint64
	bandInformationEUTRA *FreqBandInformationEUTRA
	bandInformationNR    *FreqBandInformationNR
}

func (ie *FreqBandInformation) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FreqBandInformation_Choice_bandInformationEUTRA:
		if err = ie.bandInformationEUTRA.Encode(w); err != nil {
			err = utils.WrapError("Encode bandInformationEUTRA", err)
		}
	case FreqBandInformation_Choice_bandInformationNR:
		if err = ie.bandInformationNR.Encode(w); err != nil {
			err = utils.WrapError("Encode bandInformationNR", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *FreqBandInformation) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case FreqBandInformation_Choice_bandInformationEUTRA:
		ie.bandInformationEUTRA = new(FreqBandInformationEUTRA)
		if err = ie.bandInformationEUTRA.Decode(r); err != nil {
			return utils.WrapError("Decode bandInformationEUTRA", err)
		}
	case FreqBandInformation_Choice_bandInformationNR:
		ie.bandInformationNR = new(FreqBandInformationNR)
		if err = ie.bandInformationNR.Decode(r); err != nil {
			return utils.WrapError("Decode bandInformationNR", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
