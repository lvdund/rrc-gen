package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServCellInfoXCG_NR_r16 struct {
	dl_FreqInfo_NR_r16 *FrequencyConfig_NR_r16 `optional`
	ul_FreqInfo_NR_r16 *FrequencyConfig_NR_r16 `optional`
}

func (ie *ServCellInfoXCG_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dl_FreqInfo_NR_r16 != nil, ie.ul_FreqInfo_NR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dl_FreqInfo_NR_r16 != nil {
		if err = ie.dl_FreqInfo_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode dl_FreqInfo_NR_r16", err)
		}
	}
	if ie.ul_FreqInfo_NR_r16 != nil {
		if err = ie.ul_FreqInfo_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ul_FreqInfo_NR_r16", err)
		}
	}
	return nil
}

func (ie *ServCellInfoXCG_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	var dl_FreqInfo_NR_r16Present bool
	if dl_FreqInfo_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ul_FreqInfo_NR_r16Present bool
	if ul_FreqInfo_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dl_FreqInfo_NR_r16Present {
		ie.dl_FreqInfo_NR_r16 = new(FrequencyConfig_NR_r16)
		if err = ie.dl_FreqInfo_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode dl_FreqInfo_NR_r16", err)
		}
	}
	if ul_FreqInfo_NR_r16Present {
		ie.ul_FreqInfo_NR_r16 = new(FrequencyConfig_NR_r16)
		if err = ie.ul_FreqInfo_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ul_FreqInfo_NR_r16", err)
		}
	}
	return nil
}
