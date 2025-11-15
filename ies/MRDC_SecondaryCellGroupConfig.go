package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_SecondaryCellGroupConfig struct {
	mrdc_ReleaseAndAdd      *MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd     `optional`
	mrdc_SecondaryCellGroup MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup `madatory`
}

func (ie *MRDC_SecondaryCellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.mrdc_ReleaseAndAdd != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.mrdc_ReleaseAndAdd != nil {
		if err = ie.mrdc_ReleaseAndAdd.Encode(w); err != nil {
			return utils.WrapError("Encode mrdc_ReleaseAndAdd", err)
		}
	}
	if err = ie.mrdc_SecondaryCellGroup.Encode(w); err != nil {
		return utils.WrapError("Encode mrdc_SecondaryCellGroup", err)
	}
	return nil
}

func (ie *MRDC_SecondaryCellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var mrdc_ReleaseAndAddPresent bool
	if mrdc_ReleaseAndAddPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if mrdc_ReleaseAndAddPresent {
		ie.mrdc_ReleaseAndAdd = new(MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd)
		if err = ie.mrdc_ReleaseAndAdd.Decode(r); err != nil {
			return utils.WrapError("Decode mrdc_ReleaseAndAdd", err)
		}
	}
	if err = ie.mrdc_SecondaryCellGroup.Decode(r); err != nil {
		return utils.WrapError("Decode mrdc_SecondaryCellGroup", err)
	}
	return nil
}
