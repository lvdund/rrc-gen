package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultRxTxTimeDiff_r17 struct {
	rxTxTimeDiff_ue_r17 *RxTxTimeDiff_r17 `optional`
}

func (ie *MeasResultRxTxTimeDiff_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.rxTxTimeDiff_ue_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rxTxTimeDiff_ue_r17 != nil {
		if err = ie.rxTxTimeDiff_ue_r17.Encode(w); err != nil {
			return utils.WrapError("Encode rxTxTimeDiff_ue_r17", err)
		}
	}
	return nil
}

func (ie *MeasResultRxTxTimeDiff_r17) Decode(r *uper.UperReader) error {
	var err error
	var rxTxTimeDiff_ue_r17Present bool
	if rxTxTimeDiff_ue_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rxTxTimeDiff_ue_r17Present {
		ie.rxTxTimeDiff_ue_r17 = new(RxTxTimeDiff_r17)
		if err = ie.rxTxTimeDiff_ue_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rxTxTimeDiff_ue_r17", err)
		}
	}
	return nil
}
