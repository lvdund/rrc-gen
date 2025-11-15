package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1720 struct {
	smtc4list_r17 *SSB_MTC4List_r17 `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.smtc4list_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.smtc4list_r17 != nil {
		if err = ie.smtc4list_r17.Encode(w); err != nil {
			return utils.WrapError("Encode smtc4list_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1720) Decode(r *uper.UperReader) error {
	var err error
	var smtc4list_r17Present bool
	if smtc4list_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if smtc4list_r17Present {
		ie.smtc4list_r17 = new(SSB_MTC4List_r17)
		if err = ie.smtc4list_r17.Decode(r); err != nil {
			return utils.WrapError("Decode smtc4list_r17", err)
		}
	}
	return nil
}
