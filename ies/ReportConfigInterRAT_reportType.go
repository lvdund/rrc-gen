package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	ReportConfigInterRAT_reportType_Choice_nothing uint64 = iota
	ReportConfigInterRAT_reportType_Choice_periodical
	ReportConfigInterRAT_reportType_Choice_eventTriggered
	ReportConfigInterRAT_reportType_Choice_reportCGI
	ReportConfigInterRAT_reportType_Choice_reportSFTD
)

type ReportConfigInterRAT_reportType struct {
	Choice         uint64
	periodical     *PeriodicalReportConfigInterRAT
	eventTriggered *EventTriggerConfigInterRAT
	reportCGI      *ReportCGI_EUTRA
	reportSFTD     *ReportSFTD_EUTRA
}

func (ie *ReportConfigInterRAT_reportType) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigInterRAT_reportType_Choice_periodical:
		if err = ie.periodical.Encode(w); err != nil {
			err = utils.WrapError("Encode periodical", err)
		}
	case ReportConfigInterRAT_reportType_Choice_eventTriggered:
		if err = ie.eventTriggered.Encode(w); err != nil {
			err = utils.WrapError("Encode eventTriggered", err)
		}
	case ReportConfigInterRAT_reportType_Choice_reportCGI:
		if err = ie.reportCGI.Encode(w); err != nil {
			err = utils.WrapError("Encode reportCGI", err)
		}
	case ReportConfigInterRAT_reportType_Choice_reportSFTD:
		if err = ie.reportSFTD.Encode(w); err != nil {
			err = utils.WrapError("Encode reportSFTD", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *ReportConfigInterRAT_reportType) Decode(r *uper.UperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(4, false); err != nil {
		return err
	}
	switch ie.Choice {
	case ReportConfigInterRAT_reportType_Choice_periodical:
		ie.periodical = new(PeriodicalReportConfigInterRAT)
		if err = ie.periodical.Decode(r); err != nil {
			return utils.WrapError("Decode periodical", err)
		}
	case ReportConfigInterRAT_reportType_Choice_eventTriggered:
		ie.eventTriggered = new(EventTriggerConfigInterRAT)
		if err = ie.eventTriggered.Decode(r); err != nil {
			return utils.WrapError("Decode eventTriggered", err)
		}
	case ReportConfigInterRAT_reportType_Choice_reportCGI:
		ie.reportCGI = new(ReportCGI_EUTRA)
		if err = ie.reportCGI.Decode(r); err != nil {
			return utils.WrapError("Decode reportCGI", err)
		}
	case ReportConfigInterRAT_reportType_Choice_reportSFTD:
		ie.reportSFTD = new(ReportSFTD_EUTRA)
		if err = ie.reportSFTD.Decode(r); err != nil {
			return utils.WrapError("Decode reportSFTD", err)
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
