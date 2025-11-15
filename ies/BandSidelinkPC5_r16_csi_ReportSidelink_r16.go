package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandSidelinkPC5_r16_csi_ReportSidelink_r16 struct {
	csi_RS_PortsSidelink_r16 BandSidelinkPC5_r16_csi_ReportSidelink_r16_csi_RS_PortsSidelink_r16 `madatory`
}

func (ie *BandSidelinkPC5_r16_csi_ReportSidelink_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.csi_RS_PortsSidelink_r16.Encode(w); err != nil {
		return utils.WrapError("Encode csi_RS_PortsSidelink_r16", err)
	}
	return nil
}

func (ie *BandSidelinkPC5_r16_csi_ReportSidelink_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.csi_RS_PortsSidelink_r16.Decode(r); err != nil {
		return utils.WrapError("Decode csi_RS_PortsSidelink_r16", err)
	}
	return nil
}
