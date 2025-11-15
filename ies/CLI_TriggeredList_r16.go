package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	CLI_TriggeredList_r16_Choice_nothing uint64 = iota
	CLI_TriggeredList_r16_Choice_srs_RSRP_TriggeredList_r16
	CLI_TriggeredList_r16_Choice_cli_RSSI_TriggeredList_r16
)

type CLI_TriggeredList_r16 struct {
	Choice                     uint64
	srs_RSRP_TriggeredList_r16 *SRS_RSRP_TriggeredList_r16
	cli_RSSI_TriggeredList_r16 *CLI_RSSI_TriggeredList_r16
}

func (ie *CLI_TriggeredList_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CLI_TriggeredList_r16_Choice_srs_RSRP_TriggeredList_r16:
		if err = ie.srs_RSRP_TriggeredList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode srs_RSRP_TriggeredList_r16", err)
		}
	case CLI_TriggeredList_r16_Choice_cli_RSSI_TriggeredList_r16:
		if err = ie.cli_RSSI_TriggeredList_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode cli_RSSI_TriggeredList_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *CLI_TriggeredList_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case CLI_TriggeredList_r16_Choice_srs_RSRP_TriggeredList_r16:
		ie.srs_RSRP_TriggeredList_r16 = new(SRS_RSRP_TriggeredList_r16)
		if err = ie.srs_RSRP_TriggeredList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode srs_RSRP_TriggeredList_r16", err)
		}
	case CLI_TriggeredList_r16_Choice_cli_RSSI_TriggeredList_r16:
		ie.cli_RSSI_TriggeredList_r16 = new(CLI_RSSI_TriggeredList_r16)
		if err = ie.cli_RSSI_TriggeredList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cli_RSSI_TriggeredList_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
