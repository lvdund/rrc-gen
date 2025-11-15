package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DummyI struct {
	supportedSRS_TxPortSwitch DummyI_supportedSRS_TxPortSwitch `madatory`
	txSwitchImpactToRx        *DummyI_txSwitchImpactToRx       `optional`
}

func (ie *DummyI) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.txSwitchImpactToRx != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.supportedSRS_TxPortSwitch.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSRS_TxPortSwitch", err)
	}
	if ie.txSwitchImpactToRx != nil {
		if err = ie.txSwitchImpactToRx.Encode(w); err != nil {
			return utils.WrapError("Encode txSwitchImpactToRx", err)
		}
	}
	return nil
}

func (ie *DummyI) Decode(r *uper.UperReader) error {
	var err error
	var txSwitchImpactToRxPresent bool
	if txSwitchImpactToRxPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.supportedSRS_TxPortSwitch.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSRS_TxPortSwitch", err)
	}
	if txSwitchImpactToRxPresent {
		ie.txSwitchImpactToRx = new(DummyI_txSwitchImpactToRx)
		if err = ie.txSwitchImpactToRx.Decode(r); err != nil {
			return utils.WrapError("Decode txSwitchImpactToRx", err)
		}
	}
	return nil
}
