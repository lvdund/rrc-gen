package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig struct {
	transformPrecoderDisabled *PTRS_UplinkConfig_transformPrecoderDisabled `lb:2,ub:2,optional`
	transformPrecoderEnabled  *PTRS_UplinkConfig_transformPrecoderEnabled  `lb:5,ub:5,optional`
}

func (ie *PTRS_UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.transformPrecoderDisabled != nil, ie.transformPrecoderEnabled != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.transformPrecoderDisabled != nil {
		if err = ie.transformPrecoderDisabled.Encode(w); err != nil {
			return utils.WrapError("Encode transformPrecoderDisabled", err)
		}
	}
	if ie.transformPrecoderEnabled != nil {
		if err = ie.transformPrecoderEnabled.Encode(w); err != nil {
			return utils.WrapError("Encode transformPrecoderEnabled", err)
		}
	}
	return nil
}

func (ie *PTRS_UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var transformPrecoderDisabledPresent bool
	if transformPrecoderDisabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var transformPrecoderEnabledPresent bool
	if transformPrecoderEnabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if transformPrecoderDisabledPresent {
		ie.transformPrecoderDisabled = new(PTRS_UplinkConfig_transformPrecoderDisabled)
		if err = ie.transformPrecoderDisabled.Decode(r); err != nil {
			return utils.WrapError("Decode transformPrecoderDisabled", err)
		}
	}
	if transformPrecoderEnabledPresent {
		ie.transformPrecoderEnabled = new(PTRS_UplinkConfig_transformPrecoderEnabled)
		if err = ie.transformPrecoderEnabled.Decode(r); err != nil {
			return utils.WrapError("Decode transformPrecoderEnabled", err)
		}
	}
	return nil
}
