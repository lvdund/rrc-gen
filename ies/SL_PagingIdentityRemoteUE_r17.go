package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_PagingIdentityRemoteUE_r17 struct {
	ng_5G_S_TMSI_r17 NG_5G_S_TMSI  `madatory`
	fullI_RNTI_r17   *I_RNTI_Value `optional`
}

func (ie *SL_PagingIdentityRemoteUE_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.fullI_RNTI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.ng_5G_S_TMSI_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ng_5G_S_TMSI_r17", err)
	}
	if ie.fullI_RNTI_r17 != nil {
		if err = ie.fullI_RNTI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode fullI_RNTI_r17", err)
		}
	}
	return nil
}

func (ie *SL_PagingIdentityRemoteUE_r17) Decode(r *uper.UperReader) error {
	var err error
	var fullI_RNTI_r17Present bool
	if fullI_RNTI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.ng_5G_S_TMSI_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ng_5G_S_TMSI_r17", err)
	}
	if fullI_RNTI_r17Present {
		ie.fullI_RNTI_r17 = new(I_RNTI_Value)
		if err = ie.fullI_RNTI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode fullI_RNTI_r17", err)
		}
	}
	return nil
}
