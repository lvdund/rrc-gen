package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1610 struct {
	intraFR_NR_DC_PwrSharingMode1_r16   *CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode1_r16   `optional`
	intraFR_NR_DC_PwrSharingMode2_r16   *CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode2_r16   `optional`
	intraFR_NR_DC_DynamicPwrSharing_r16 *CA_ParametersNRDC_v1610_intraFR_NR_DC_DynamicPwrSharing_r16 `optional`
	asyncNRDC_r16                       *CA_ParametersNRDC_v1610_asyncNRDC_r16                       `optional`
}

func (ie *CA_ParametersNRDC_v1610) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.intraFR_NR_DC_PwrSharingMode1_r16 != nil, ie.intraFR_NR_DC_PwrSharingMode2_r16 != nil, ie.intraFR_NR_DC_DynamicPwrSharing_r16 != nil, ie.asyncNRDC_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.intraFR_NR_DC_PwrSharingMode1_r16 != nil {
		if err = ie.intraFR_NR_DC_PwrSharingMode1_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFR_NR_DC_PwrSharingMode1_r16", err)
		}
	}
	if ie.intraFR_NR_DC_PwrSharingMode2_r16 != nil {
		if err = ie.intraFR_NR_DC_PwrSharingMode2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFR_NR_DC_PwrSharingMode2_r16", err)
		}
	}
	if ie.intraFR_NR_DC_DynamicPwrSharing_r16 != nil {
		if err = ie.intraFR_NR_DC_DynamicPwrSharing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode intraFR_NR_DC_DynamicPwrSharing_r16", err)
		}
	}
	if ie.asyncNRDC_r16 != nil {
		if err = ie.asyncNRDC_r16.Encode(w); err != nil {
			return utils.WrapError("Encode asyncNRDC_r16", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1610) Decode(r *uper.UperReader) error {
	var err error
	var intraFR_NR_DC_PwrSharingMode1_r16Present bool
	if intraFR_NR_DC_PwrSharingMode1_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFR_NR_DC_PwrSharingMode2_r16Present bool
	if intraFR_NR_DC_PwrSharingMode2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var intraFR_NR_DC_DynamicPwrSharing_r16Present bool
	if intraFR_NR_DC_DynamicPwrSharing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var asyncNRDC_r16Present bool
	if asyncNRDC_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if intraFR_NR_DC_PwrSharingMode1_r16Present {
		ie.intraFR_NR_DC_PwrSharingMode1_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode1_r16)
		if err = ie.intraFR_NR_DC_PwrSharingMode1_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFR_NR_DC_PwrSharingMode1_r16", err)
		}
	}
	if intraFR_NR_DC_PwrSharingMode2_r16Present {
		ie.intraFR_NR_DC_PwrSharingMode2_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_PwrSharingMode2_r16)
		if err = ie.intraFR_NR_DC_PwrSharingMode2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFR_NR_DC_PwrSharingMode2_r16", err)
		}
	}
	if intraFR_NR_DC_DynamicPwrSharing_r16Present {
		ie.intraFR_NR_DC_DynamicPwrSharing_r16 = new(CA_ParametersNRDC_v1610_intraFR_NR_DC_DynamicPwrSharing_r16)
		if err = ie.intraFR_NR_DC_DynamicPwrSharing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode intraFR_NR_DC_DynamicPwrSharing_r16", err)
		}
	}
	if asyncNRDC_r16Present {
		ie.asyncNRDC_r16 = new(CA_ParametersNRDC_v1610_asyncNRDC_r16)
		if err = ie.asyncNRDC_r16.Decode(r); err != nil {
			return utils.WrapError("Decode asyncNRDC_r16", err)
		}
	}
	return nil
}
