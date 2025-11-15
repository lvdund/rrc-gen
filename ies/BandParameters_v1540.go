package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540 struct {
	srs_CarrierSwitch *BandParameters_v1540_srs_CarrierSwitch `lb:1,ub:maxSimultaneousBands,optional`
	srs_TxSwitch      *BandParameters_v1540_srs_TxSwitch      `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1540) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_CarrierSwitch != nil, ie.srs_TxSwitch != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_CarrierSwitch != nil {
		if err = ie.srs_CarrierSwitch.Encode(w); err != nil {
			return utils.WrapError("Encode srs_CarrierSwitch", err)
		}
	}
	if ie.srs_TxSwitch != nil {
		if err = ie.srs_TxSwitch.Encode(w); err != nil {
			return utils.WrapError("Encode srs_TxSwitch", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1540) Decode(r *uper.UperReader) error {
	var err error
	var srs_CarrierSwitchPresent bool
	if srs_CarrierSwitchPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var srs_TxSwitchPresent bool
	if srs_TxSwitchPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_CarrierSwitchPresent {
		ie.srs_CarrierSwitch = new(BandParameters_v1540_srs_CarrierSwitch)
		if err = ie.srs_CarrierSwitch.Decode(r); err != nil {
			return utils.WrapError("Decode srs_CarrierSwitch", err)
		}
	}
	if srs_TxSwitchPresent {
		ie.srs_TxSwitch = new(BandParameters_v1540_srs_TxSwitch)
		if err = ie.srs_TxSwitch.Decode(r); err != nil {
			return utils.WrapError("Decode srs_TxSwitch", err)
		}
	}
	return nil
}
