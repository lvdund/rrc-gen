package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkConfigCommon struct {
	frequencyInfoUL  *FrequencyInfoUL   `optional`
	initialUplinkBWP *BWP_UplinkCommon  `optional`
	dummy            TimeAlignmentTimer `madatory`
}

func (ie *UplinkConfigCommon) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.frequencyInfoUL != nil, ie.initialUplinkBWP != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.frequencyInfoUL != nil {
		if err = ie.frequencyInfoUL.Encode(w); err != nil {
			return utils.WrapError("Encode frequencyInfoUL", err)
		}
	}
	if ie.initialUplinkBWP != nil {
		if err = ie.initialUplinkBWP.Encode(w); err != nil {
			return utils.WrapError("Encode initialUplinkBWP", err)
		}
	}
	if err = ie.dummy.Encode(w); err != nil {
		return utils.WrapError("Encode dummy", err)
	}
	return nil
}

func (ie *UplinkConfigCommon) Decode(r *uper.UperReader) error {
	var err error
	var frequencyInfoULPresent bool
	if frequencyInfoULPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var initialUplinkBWPPresent bool
	if initialUplinkBWPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if frequencyInfoULPresent {
		ie.frequencyInfoUL = new(FrequencyInfoUL)
		if err = ie.frequencyInfoUL.Decode(r); err != nil {
			return utils.WrapError("Decode frequencyInfoUL", err)
		}
	}
	if initialUplinkBWPPresent {
		ie.initialUplinkBWP = new(BWP_UplinkCommon)
		if err = ie.initialUplinkBWP.Decode(r); err != nil {
			return utils.WrapError("Decode initialUplinkBWP", err)
		}
	}
	if err = ie.dummy.Decode(r); err != nil {
		return utils.WrapError("Decode dummy", err)
	}
	return nil
}
