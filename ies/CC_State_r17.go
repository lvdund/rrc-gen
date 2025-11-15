package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CC_State_r17 struct {
	dlCarrier_r17 *CarrierState_r17 `optional`
	ulCarrier_r17 *CarrierState_r17 `optional`
}

func (ie *CC_State_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dlCarrier_r17 != nil, ie.ulCarrier_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dlCarrier_r17 != nil {
		if err = ie.dlCarrier_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dlCarrier_r17", err)
		}
	}
	if ie.ulCarrier_r17 != nil {
		if err = ie.ulCarrier_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ulCarrier_r17", err)
		}
	}
	return nil
}

func (ie *CC_State_r17) Decode(r *uper.UperReader) error {
	var err error
	var dlCarrier_r17Present bool
	if dlCarrier_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ulCarrier_r17Present bool
	if ulCarrier_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dlCarrier_r17Present {
		ie.dlCarrier_r17 = new(CarrierState_r17)
		if err = ie.dlCarrier_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dlCarrier_r17", err)
		}
	}
	if ulCarrier_r17Present {
		ie.ulCarrier_r17 = new(CarrierState_r17)
		if err = ie.ulCarrier_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ulCarrier_r17", err)
		}
	}
	return nil
}
