package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type QuantityConfigRS struct {
	ssb_FilterConfig    FilterConfig `madatory`
	csi_RS_FilterConfig FilterConfig `madatory`
}

func (ie *QuantityConfigRS) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ssb_FilterConfig.Encode(w); err != nil {
		return utils.WrapError("Encode ssb_FilterConfig", err)
	}
	if err = ie.csi_RS_FilterConfig.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_FilterConfig", err)
	}
	return nil
}

func (ie *QuantityConfigRS) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ssb_FilterConfig.Decode(r); err != nil {
		return utils.WrapError("Decode ssb_FilterConfig", err)
	}
	if err = ie.csi_RS_FilterConfig.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_FilterConfig", err)
	}
	return nil
}
