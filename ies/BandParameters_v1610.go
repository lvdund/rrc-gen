package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandParameters_v1610 struct {
	srs_TxSwitch_v1610 *BandParameters_v1610_srs_TxSwitch_v1610 `optional`
}

func (ie *BandParameters_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.srs_TxSwitch_v1610 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.srs_TxSwitch_v1610 != nil {
		if err = ie.srs_TxSwitch_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode srs_TxSwitch_v1610", err)
		}
	}
	return nil
}

func (ie *BandParameters_v1610) Decode(r *uper.UperReader) error {
	var err error
	var srs_TxSwitch_v1610Present bool
	if srs_TxSwitch_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	if srs_TxSwitch_v1610Present {
		ie.srs_TxSwitch_v1610 = new(BandParameters_v1610_srs_TxSwitch_v1610)
		if err = ie.srs_TxSwitch_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode srs_TxSwitch_v1610", err)
		}
	}
	return nil
}
