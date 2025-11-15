package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HARQ_ProcessPerSCS_r17 struct {
	scs_120kHz_r17 *HARQ_ProcessPerSCS_r17_scs_120kHz_r17 `optional`
	scs_480kHz_r17 *HARQ_ProcessPerSCS_r17_scs_480kHz_r17 `optional`
	scs_960kHz_r17 *HARQ_ProcessPerSCS_r17_scs_960kHz_r17 `optional`
}

func (ie *HARQ_ProcessPerSCS_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.scs_120kHz_r17 != nil, ie.scs_480kHz_r17 != nil, ie.scs_960kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.scs_120kHz_r17 != nil {
		if err = ie.scs_120kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_120kHz_r17", err)
		}
	}
	if ie.scs_480kHz_r17 != nil {
		if err = ie.scs_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_480kHz_r17", err)
		}
	}
	if ie.scs_960kHz_r17 != nil {
		if err = ie.scs_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode scs_960kHz_r17", err)
		}
	}
	return nil
}

func (ie *HARQ_ProcessPerSCS_r17) Decode(r *uper.UperReader) error {
	var err error
	var scs_120kHz_r17Present bool
	if scs_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_480kHz_r17Present bool
	if scs_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var scs_960kHz_r17Present bool
	if scs_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if scs_120kHz_r17Present {
		ie.scs_120kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_120kHz_r17)
		if err = ie.scs_120kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_120kHz_r17", err)
		}
	}
	if scs_480kHz_r17Present {
		ie.scs_480kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_480kHz_r17)
		if err = ie.scs_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_480kHz_r17", err)
		}
	}
	if scs_960kHz_r17Present {
		ie.scs_960kHz_r17 = new(HARQ_ProcessPerSCS_r17_scs_960kHz_r17)
		if err = ie.scs_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode scs_960kHz_r17", err)
		}
	}
	return nil
}
