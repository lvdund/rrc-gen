package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCLI_RSSI_r16 struct {
	rssi_ResourceId_r16 RSSI_ResourceId_r16 `madatory`
	cli_RSSI_Result_r16 CLI_RSSI_Range_r16  `madatory`
}

func (ie *MeasResultCLI_RSSI_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.rssi_ResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode rssi_ResourceId_r16", err)
	}
	if err = ie.cli_RSSI_Result_r16.Encode(w); err != nil {
		return utils.WrapError("Encode cli_RSSI_Result_r16", err)
	}
	return nil
}

func (ie *MeasResultCLI_RSSI_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.rssi_ResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode rssi_ResourceId_r16", err)
	}
	if err = ie.cli_RSSI_Result_r16.Decode(r); err != nil {
		return utils.WrapError("Decode cli_RSSI_Result_r16", err)
	}
	return nil
}
