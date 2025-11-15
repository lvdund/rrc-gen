package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ProcessingParameters_differentTB_PerSlot struct {
	upto1 *NumberOfCarriers `optional`
	upto2 *NumberOfCarriers `optional`
	upto4 *NumberOfCarriers `optional`
	upto7 *NumberOfCarriers `optional`
}

func (ie *ProcessingParameters_differentTB_PerSlot) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.upto1 != nil, ie.upto2 != nil, ie.upto4 != nil, ie.upto7 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.upto1 != nil {
		if err = ie.upto1.Encode(w); err != nil {
			return utils.WrapError("Encode upto1", err)
		}
	}
	if ie.upto2 != nil {
		if err = ie.upto2.Encode(w); err != nil {
			return utils.WrapError("Encode upto2", err)
		}
	}
	if ie.upto4 != nil {
		if err = ie.upto4.Encode(w); err != nil {
			return utils.WrapError("Encode upto4", err)
		}
	}
	if ie.upto7 != nil {
		if err = ie.upto7.Encode(w); err != nil {
			return utils.WrapError("Encode upto7", err)
		}
	}
	return nil
}

func (ie *ProcessingParameters_differentTB_PerSlot) Decode(r *uper.UperReader) error {
	var err error
	var upto1Present bool
	if upto1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var upto2Present bool
	if upto2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var upto4Present bool
	if upto4Present, err = r.ReadBool(); err != nil {
		return err
	}
	var upto7Present bool
	if upto7Present, err = r.ReadBool(); err != nil {
		return err
	}
	if upto1Present {
		ie.upto1 = new(NumberOfCarriers)
		if err = ie.upto1.Decode(r); err != nil {
			return utils.WrapError("Decode upto1", err)
		}
	}
	if upto2Present {
		ie.upto2 = new(NumberOfCarriers)
		if err = ie.upto2.Decode(r); err != nil {
			return utils.WrapError("Decode upto2", err)
		}
	}
	if upto4Present {
		ie.upto4 = new(NumberOfCarriers)
		if err = ie.upto4.Decode(r); err != nil {
			return utils.WrapError("Decode upto4", err)
		}
	}
	if upto7Present {
		ie.upto7 = new(NumberOfCarriers)
		if err = ie.upto7.Decode(r); err != nil {
			return utils.WrapError("Decode upto7", err)
		}
	}
	return nil
}
