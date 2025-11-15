package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RedirectedCarrierInfo_Choice_nothing uint64 = iota
	RedirectedCarrierInfo_Choice_nr
	RedirectedCarrierInfo_Choice_eutra
)

type RedirectedCarrierInfo struct {
	Choice uint64
	nr     *CarrierInfoNR
	eutra  *RedirectedCarrierInfo_EUTRA
}

func (ie *RedirectedCarrierInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RedirectedCarrierInfo_Choice_nr:
		if err = ie.nr.Encode(w); err != nil {
			err = utils.WrapError("Encode nr", err)
		}
	case RedirectedCarrierInfo_Choice_eutra:
		if err = ie.eutra.Encode(w); err != nil {
			err = utils.WrapError("Encode eutra", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RedirectedCarrierInfo) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RedirectedCarrierInfo_Choice_nr:
		ie.nr = new(CarrierInfoNR)
		if err = ie.nr.Decode(r); err != nil {
			return utils.WrapError("Decode nr", err)
		}
	case RedirectedCarrierInfo_Choice_eutra:
		ie.eutra = new(RedirectedCarrierInfo_EUTRA)
		if err = ie.eutra.Decode(r); err != nil {
			return utils.WrapError("Decode eutra", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
