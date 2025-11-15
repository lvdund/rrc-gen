package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1610_srs_TxSwitch_v1610 struct {
	supportedSRS_TxPortSwitch_v1610 BandParameters_v1610_srs_TxSwitch_v1610_supportedSRS_TxPortSwitch_v1610 `madatory`
}

func (ie *BandParameters_v1610_srs_TxSwitch_v1610) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.supportedSRS_TxPortSwitch_v1610.Encode(w); err != nil {
		return utils.WrapError("Encode supportedSRS_TxPortSwitch_v1610", err)
	}
	return nil
}

func (ie *BandParameters_v1610_srs_TxSwitch_v1610) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.supportedSRS_TxPortSwitch_v1610.Decode(r); err != nil {
		return utils.WrapError("Decode supportedSRS_TxPortSwitch_v1610", err)
	}
	return nil
}
