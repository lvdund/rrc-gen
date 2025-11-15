package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Fr1_100mhz struct {
	scs_15kHz *Fr1_100mhz_scs_15kHz `optional`
	scs_30kHz *Fr1_100mhz_scs_30kHz `optional`
	scs_60kHz *Fr1_100mhz_scs_60kHz `optional`
}

func (ie *Fr1_100mhz) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_15kHz != nil, ie.scs_30kHz != nil, ie.scs_60kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_15kHz != nil {
		if err = ie.scs_15kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_15kHz", err)
		}
	}
	if ie.scs_30kHz != nil {
		if err = ie.scs_30kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_30kHz", err)
		}
	}
	if ie.scs_60kHz != nil {
		if err = ie.scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode scs_60kHz", err)
		}
	}
	return nil
}

func (ie *Fr1_100mhz) Decode(r *uper.UperReader) error {
	var err error
	var scs_15kHzPresent bool
	if scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_30kHzPresent bool
	if scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_60kHzPresent bool
	if scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_15kHzPresent {
		ie.scs_15kHz = new(Fr1_100mhz_scs_15kHz)
		if err = ie.scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_15kHz", err)
		}
	}
	if scs_30kHzPresent {
		ie.scs_30kHz = new(Fr1_100mhz_scs_30kHz)
		if err = ie.scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_30kHz", err)
		}
	}
	if scs_60kHzPresent {
		ie.scs_60kHz = new(Fr1_100mhz_scs_60kHz)
		if err = ie.scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode scs_60kHz", err)
		}
	}
	return nil
}
