package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FreqPriorityDedicatedSlicing_r17 struct {
	dl_ExplicitCarrierFreq_r17 ARFCN_ValueNR               `madatory`
	sliceInfoListDedicated_r17 *SliceInfoListDedicated_r17 `optional`
}

func (ie *FreqPriorityDedicatedSlicing_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.sliceInfoListDedicated_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.dl_ExplicitCarrierFreq_r17.Encode(w); err != nil {
		return utils.WrapError("Encode dl_ExplicitCarrierFreq_r17", err)
	}
	if ie.sliceInfoListDedicated_r17 != nil {
		if err = ie.sliceInfoListDedicated_r17.Encode(w); err != nil {
			return utils.WrapError("Encode sliceInfoListDedicated_r17", err)
		}
	}
	return nil
}

func (ie *FreqPriorityDedicatedSlicing_r17) Decode(r *uper.UperReader) error {
	var err error
	var sliceInfoListDedicated_r17Present bool
	if sliceInfoListDedicated_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.dl_ExplicitCarrierFreq_r17.Decode(r); err != nil {
		return utils.WrapError("Decode dl_ExplicitCarrierFreq_r17", err)
	}
	if sliceInfoListDedicated_r17Present {
		ie.sliceInfoListDedicated_r17 = new(SliceInfoListDedicated_r17)
		if err = ie.sliceInfoListDedicated_r17.Decode(r); err != nil {
			return utils.WrapError("Decode sliceInfoListDedicated_r17", err)
		}
	}
	return nil
}
