package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type UplinkDataCompression_r17_newSetup struct {
	bufferSize_r17 UplinkDataCompression_r17_newSetup_bufferSize_r17  `madatory`
	dictionary_r17 *UplinkDataCompression_r17_newSetup_dictionary_r17 `optional`
}

func (ie *UplinkDataCompression_r17_newSetup) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dictionary_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.bufferSize_r17.Encode(w); err != nil {
		return utils.WrapError("Encode bufferSize_r17", err)
	}
	if ie.dictionary_r17 != nil {
		if err = ie.dictionary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode dictionary_r17", err)
		}
	}
	return nil
}

func (ie *UplinkDataCompression_r17_newSetup) Decode(r *uper.UperReader) error {
	var err error
	var dictionary_r17Present bool
	if dictionary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.bufferSize_r17.Decode(r); err != nil {
		return utils.WrapError("Decode bufferSize_r17", err)
	}
	if dictionary_r17Present {
		ie.dictionary_r17 = new(UplinkDataCompression_r17_newSetup_dictionary_r17)
		if err = ie.dictionary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode dictionary_r17", err)
		}
	}
	return nil
}
