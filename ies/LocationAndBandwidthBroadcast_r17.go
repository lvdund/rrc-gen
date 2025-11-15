package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	LocationAndBandwidthBroadcast_r17_Choice_nothing uint64 = iota
	LocationAndBandwidthBroadcast_r17_Choice_sameAsSib1ConfiguredLocationAndBW
	LocationAndBandwidthBroadcast_r17_Choice_locationAndBandwidth
)

type LocationAndBandwidthBroadcast_r17 struct {
	Choice                            uint64
	sameAsSib1ConfiguredLocationAndBW uper.NULL `madatory`
	locationAndBandwidth              int64     `lb:0,ub:37949,madatory`
}

func (ie *LocationAndBandwidthBroadcast_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationAndBandwidthBroadcast_r17_Choice_sameAsSib1ConfiguredLocationAndBW:
		if err := w.WriteNull(); err != nil {
			err = utils.WrapError("Encode sameAsSib1ConfiguredLocationAndBW", err)
		}
	case LocationAndBandwidthBroadcast_r17_Choice_locationAndBandwidth:
		if err = w.WriteInteger(int64(ie.locationAndBandwidth), &uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			err = utils.WrapError("Encode locationAndBandwidth", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *LocationAndBandwidthBroadcast_r17) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case LocationAndBandwidthBroadcast_r17_Choice_sameAsSib1ConfiguredLocationAndBW:
		if err := r.ReadNull(); err != nil {
			return utils.WrapError("Decode sameAsSib1ConfiguredLocationAndBW", err)
		}
	case LocationAndBandwidthBroadcast_r17_Choice_locationAndBandwidth:
		var tmp_int_locationAndBandwidth int64
		if tmp_int_locationAndBandwidth, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 37949}, false); err != nil {
			return utils.WrapError("Decode locationAndBandwidth", err)
		}
		ie.locationAndBandwidth = tmp_int_locationAndBandwidth
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
