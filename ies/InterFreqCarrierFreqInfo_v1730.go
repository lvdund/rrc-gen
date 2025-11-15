package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type InterFreqCarrierFreqInfo_v1730 struct {
	channelAccessMode2_r17 *InterFreqCarrierFreqInfo_v1730_channelAccessMode2_r17 `optional`
}

func (ie *InterFreqCarrierFreqInfo_v1730) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.channelAccessMode2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.channelAccessMode2_r17 != nil {
		if err = ie.channelAccessMode2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode channelAccessMode2_r17", err)
		}
	}
	return nil
}

func (ie *InterFreqCarrierFreqInfo_v1730) Decode(r *uper.UperReader) error {
	var err error
	var channelAccessMode2_r17Present bool
	if channelAccessMode2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if channelAccessMode2_r17Present {
		ie.channelAccessMode2_r17 = new(InterFreqCarrierFreqInfo_v1730_channelAccessMode2_r17)
		if err = ie.channelAccessMode2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode channelAccessMode2_r17", err)
		}
	}
	return nil
}
