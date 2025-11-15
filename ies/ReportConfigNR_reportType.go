package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigNR_reportType_Choice_nothing uint64 = iota
	ReportConfigNR_reportType_Choice_periodical
	ReportConfigNR_reportType_Choice_eventTriggered
	ReportConfigNR_reportType_Choice_reportCGI
	ReportConfigNR_reportType_Choice_reportSFTD
	ReportConfigNR_reportType_Choice_condTriggerConfig_r16
	ReportConfigNR_reportType_Choice_cli_Periodical_r16
	ReportConfigNR_reportType_Choice_cli_EventTriggered_r16
	ReportConfigNR_reportType_Choice_rxTxPeriodical_r17
)

type ReportConfigNR_reportType struct {
	Choice                 uint64
	periodical             *PeriodicalReportConfig
	eventTriggered         *EventTriggerConfig
	reportCGI              *ReportCGI
	reportSFTD             *ReportSFTD_NR
	condTriggerConfig_r16  *CondTriggerConfig_r16
	cli_Periodical_r16     *CLI_PeriodicalReportConfig_r16
	cli_EventTriggered_r16 *CLI_EventTriggerConfig_r16
	rxTxPeriodical_r17     *RxTxPeriodical_r17
}

func (ie *ReportConfigNR_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_reportType_Choice_periodical:
		if err = ie.periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode periodical", err)
		}
	case ReportConfigNR_reportType_Choice_eventTriggered:
		if err = ie.eventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode eventTriggered", err)
		}
	case ReportConfigNR_reportType_Choice_reportCGI:
		if err = ie.reportCGI.Encode(w); err != nil {
			err = utils.WrapError("Encode reportCGI", err)
		}
	case ReportConfigNR_reportType_Choice_reportSFTD:
		if err = ie.reportSFTD.Encode(w); err != nil {
			err = utils.WrapError("Encode reportSFTD", err)
		}
	case ReportConfigNR_reportType_Choice_condTriggerConfig_r16:
		if err = ie.condTriggerConfig_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode condTriggerConfig_r16", err)
		}
	case ReportConfigNR_reportType_Choice_cli_Periodical_r16:
		if err = ie.cli_Periodical_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode cli_Periodical_r16", err)
		}
	case ReportConfigNR_reportType_Choice_cli_EventTriggered_r16:
		if err = ie.cli_EventTriggered_r16.Encode(w); err != nil {
			err = utils.WrapError("Encode cli_EventTriggered_r16", err)
		}
	case ReportConfigNR_reportType_Choice_rxTxPeriodical_r17:
		if err = ie.rxTxPeriodical_r17.Encode(w); err != nil {
			err = utils.WrapError("Encode rxTxPeriodical_r17", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigNR_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(8, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigNR_reportType_Choice_periodical:
		ie.periodical = new(PeriodicalReportConfig)
		if err = ie.periodical.Decode(r); err != nil {
			return utils.WrapError("Decode periodical", err)
		}
	case ReportConfigNR_reportType_Choice_eventTriggered:
		ie.eventTriggered = new(EventTriggerConfig)
		if err = ie.eventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode eventTriggered", err)
		}
	case ReportConfigNR_reportType_Choice_reportCGI:
		ie.reportCGI = new(ReportCGI)
		if err = ie.reportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode reportCGI", err)
		}
	case ReportConfigNR_reportType_Choice_reportSFTD:
		ie.reportSFTD = new(ReportSFTD_NR)
		if err = ie.reportSFTD.Decode(r); err != nil {
			return utils.WrapError("Decode reportSFTD", err)
		}
	case ReportConfigNR_reportType_Choice_condTriggerConfig_r16:
		ie.condTriggerConfig_r16 = new(CondTriggerConfig_r16)
		if err = ie.condTriggerConfig_r16.Decode(r); err != nil {
			return utils.WrapError("Decode condTriggerConfig_r16", err)
		}
	case ReportConfigNR_reportType_Choice_cli_Periodical_r16:
		ie.cli_Periodical_r16 = new(CLI_PeriodicalReportConfig_r16)
		if err = ie.cli_Periodical_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cli_Periodical_r16", err)
		}
	case ReportConfigNR_reportType_Choice_cli_EventTriggered_r16:
		ie.cli_EventTriggered_r16 = new(CLI_EventTriggerConfig_r16)
		if err = ie.cli_EventTriggered_r16.Decode(r); err != nil {
			return utils.WrapError("Decode cli_EventTriggered_r16", err)
		}
	case ReportConfigNR_reportType_Choice_rxTxPeriodical_r17:
		ie.rxTxPeriodical_r17 = new(RxTxPeriodical_r17)
		if err = ie.rxTxPeriodical_r17.Decode(r); err != nil {
			return utils.WrapError("Decode rxTxPeriodical_r17", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
