package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1540_srs_TxSwitch struct {
	supportedSRS_TxPortSwitch BandParameters_v1540_srs_TxSwitch_supportedSRS_TxPortSwitch `madatory`
	txSwitchImpactToRx        *int64                                                      `lb:1,ub:32,optional`
	txSwitchWithAnotherBand   *int64                                                      `lb:1,ub:32,optional`
}

func (ie *BandParameters_v1540_srs_TxSwitch) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.txSwitchImpactToRx != nil, ie.txSwitchWithAnotherBand != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedSRS_TxPortSwitch.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSRS_TxPortSwitch", err)
	}
	if ie.txSwitchImpactToRx != nil {
		if err = w.WriteInteger(*ie.txSwitchImpactToRx, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode txSwitchImpactToRx", err)
		}
	}
	if ie.txSwitchWithAnotherBand != nil {
		if err = w.WriteInteger(*ie.txSwitchWithAnotherBand, &uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Encode txSwitchWithAnotherBand", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1540_srs_TxSwitch) Decode(r *uper.UperReader) error {
	var err error
	var txSwitchImpactToRxPresent bool
	if txSwitchImpactToRxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var txSwitchWithAnotherBandPresent bool
	if txSwitchWithAnotherBandPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedSRS_TxPortSwitch.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSRS_TxPortSwitch", err)
	}
	if txSwitchImpactToRxPresent {
		var tmp_int_txSwitchImpactToRx int64
		if tmp_int_txSwitchImpactToRx, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode txSwitchImpactToRx", err)
		}
		ie.txSwitchImpactToRx = &tmp_int_txSwitchImpactToRx
	}
	if txSwitchWithAnotherBandPresent {
		var tmp_int_txSwitchWithAnotherBand int64
		if tmp_int_txSwitchWithAnotherBand, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode txSwitchWithAnotherBand", err)
		}
		ie.txSwitchWithAnotherBand = &tmp_int_txSwitchWithAnotherBand
	}
	return nil
}
