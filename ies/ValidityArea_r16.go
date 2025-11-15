package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ValidityArea_r16 struct {
	carrierFreq_r16      ARFCN_ValueNR     `madatory`
	validityCellList_r16 *ValidityCellList `optional`
}

func (ie *ValidityArea_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.validityCellList_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.carrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode carrierFreq_r16", err)
	}
	if ie.validityCellList_r16 != nil {
		if err = ie.validityCellList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode validityCellList_r16", err)
		}
	}
	return nil
}

func (ie *ValidityArea_r16) Decode(r *uper.UperReader) error {
	var err error
	var validityCellList_r16Present bool
	if validityCellList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.carrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode carrierFreq_r16", err)
	}
	if validityCellList_r16Present {
		ie.validityCellList_r16 = new(ValidityCellList)
		if err = ie.validityCellList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode validityCellList_r16", err)
		}
	}
	return nil
}
