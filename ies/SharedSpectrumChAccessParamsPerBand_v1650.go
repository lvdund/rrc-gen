package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SharedSpectrumChAccessParamsPerBand_v1650 struct {
	extendedSearchSpaceSwitchWithDCI_r16 *SharedSpectrumChAccessParamsPerBand_v1650_extendedSearchSpaceSwitchWithDCI_r16 `optional`
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1650) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.extendedSearchSpaceSwitchWithDCI_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.extendedSearchSpaceSwitchWithDCI_r16 != nil {
		if err = ie.extendedSearchSpaceSwitchWithDCI_r16.Encode(w); err != nil {
			return utils.WrapError("Encode extendedSearchSpaceSwitchWithDCI_r16", err)
		}
	}
	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1650) Decode(r *uper.UperReader) error {
	var err error
	var extendedSearchSpaceSwitchWithDCI_r16Present bool
	if extendedSearchSpaceSwitchWithDCI_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if extendedSearchSpaceSwitchWithDCI_r16Present {
		ie.extendedSearchSpaceSwitchWithDCI_r16 = new(SharedSpectrumChAccessParamsPerBand_v1650_extendedSearchSpaceSwitchWithDCI_r16)
		if err = ie.extendedSearchSpaceSwitchWithDCI_r16.Decode(r); err != nil {
			return utils.WrapError("Decode extendedSearchSpaceSwitchWithDCI_r16", err)
		}
	}
	return nil
}
