package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TDD_UL_DL_ConfigCommon struct {
	referenceSubcarrierSpacing SubcarrierSpacing  `madatory`
	pattern1                   TDD_UL_DL_Pattern  `madatory`
	pattern2                   *TDD_UL_DL_Pattern `optional`
}

func (ie *TDD_UL_DL_ConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.pattern2 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.referenceSubcarrierSpacing.Encode(w); err != nil {
		return utils.WrapError("Encode referenceSubcarrierSpacing", err)
	}
	if err = ie.pattern1.Encode(w); err != nil {
		return utils.WrapError("Encode pattern1", err)
	}
	if ie.pattern2 != nil {
		if err = ie.pattern2.Encode(w); err != nil {
			return utils.WrapError("Encode pattern2", err)
		}
	}
	return nil
}

func (ie *TDD_UL_DL_ConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var pattern2Present bool
	if pattern2Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.referenceSubcarrierSpacing.Decode(r); err != nil {
		return utils.WrapError("Decode referenceSubcarrierSpacing", err)
	}
	if err = ie.pattern1.Decode(r); err != nil {
		return utils.WrapError("Decode pattern1", err)
	}
	if pattern2Present {
		ie.pattern2 = new(TDD_UL_DL_Pattern)
		if err = ie.pattern2.Decode(r); err != nil {
			return utils.WrapError("Decode pattern2", err)
		}
	}
	return nil
}
