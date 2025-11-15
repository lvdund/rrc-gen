package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_nothing uint64 = iota
	RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_nrReconnectCellId_r16
	RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_eutraReconnectCellId_r16
)

type RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16 struct {
	Choice                   uint64
	nrReconnectCellId_r16    *CGI_Info_Logging_r16
	eutraReconnectCellId_r16 *CGI_InfoEUTRALogging
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_nrReconnectCellId_r16:
		if err = ie.nrReconnectCellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode nrReconnectCellId_r16", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_eutraReconnectCellId_r16:
		if err = ie.eutraReconnectCellId_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode eutraReconnectCellId_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(2, false); err != nil {
		return err
	}
	switch ie.Choice {
	case RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_nrReconnectCellId_r16:
		ie.nrReconnectCellId_r16 = new(CGI_Info_Logging_r16)
		if err = ie.nrReconnectCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode nrReconnectCellId_r16", err)
		}
	case RLF_Report_r16_nr_RLF_Report_r16_reconnectCellId_r16_Choice_eutraReconnectCellId_r16:
		ie.eutraReconnectCellId_r16 = new(CGI_InfoEUTRALogging)
		if err = ie.eutraReconnectCellId_r16.Decode(r); err != nil {
			return utils.WrapError("Decode eutraReconnectCellId_r16", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
