package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_ReportFramework struct {
	maxNumberPeriodicCSI_PerBWP_ForCSI_Report       int64                                                          `lb:1,ub:4,madatory`
	maxNumberAperiodicCSI_PerBWP_ForCSI_Report      int64                                                          `lb:1,ub:4,madatory`
	maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report int64                                                          `lb:0,ub:4,madatory`
	maxNumberPeriodicCSI_PerBWP_ForBeamReport       int64                                                          `lb:1,ub:4,madatory`
	maxNumberAperiodicCSI_PerBWP_ForBeamReport      int64                                                          `lb:1,ub:4,madatory`
	maxNumberAperiodicCSI_triggeringStatePerCC      CSI_ReportFramework_maxNumberAperiodicCSI_triggeringStatePerCC `madatory`
	maxNumberSemiPersistentCSI_PerBWP_ForBeamReport int64                                                          `lb:0,ub:4,madatory`
	simultaneousCSI_ReportsPerCC                    int64                                                          `lb:1,ub:8,madatory`
}

func (ie *CSI_ReportFramework) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.maxNumberPeriodicCSI_PerBWP_ForCSI_Report, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPeriodicCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.maxNumberAperiodicCSI_PerBWP_ForCSI_Report, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report", err)
	}
	if err = w.WriteInteger(ie.maxNumberPeriodicCSI_PerBWP_ForBeamReport, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberPeriodicCSI_PerBWP_ForBeamReport", err)
	}
	if err = w.WriteInteger(ie.maxNumberAperiodicCSI_PerBWP_ForBeamReport, &uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberAperiodicCSI_PerBWP_ForBeamReport", err)
	}
	if err = ie.maxNumberAperiodicCSI_triggeringStatePerCC.Encode(w); err != nil {
		return utils.WrapError("Encode maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	if err = w.WriteInteger(ie.maxNumberSemiPersistentCSI_PerBWP_ForBeamReport, &uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("WriteInteger maxNumberSemiPersistentCSI_PerBWP_ForBeamReport", err)
	}
	if err = w.WriteInteger(ie.simultaneousCSI_ReportsPerCC, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger simultaneousCSI_ReportsPerCC", err)
	}
	return nil
}

func (ie *CSI_ReportFramework) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_maxNumberPeriodicCSI_PerBWP_ForCSI_Report int64
	if tmp_int_maxNumberPeriodicCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPeriodicCSI_PerBWP_ForCSI_Report", err)
	}
	ie.maxNumberPeriodicCSI_PerBWP_ForCSI_Report = tmp_int_maxNumberPeriodicCSI_PerBWP_ForCSI_Report
	var tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_Report int64
	if tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicCSI_PerBWP_ForCSI_Report", err)
	}
	ie.maxNumberAperiodicCSI_PerBWP_ForCSI_Report = tmp_int_maxNumberAperiodicCSI_PerBWP_ForCSI_Report
	var tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report int64
	if tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report", err)
	}
	ie.maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report = tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForCSI_Report
	var tmp_int_maxNumberPeriodicCSI_PerBWP_ForBeamReport int64
	if tmp_int_maxNumberPeriodicCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberPeriodicCSI_PerBWP_ForBeamReport", err)
	}
	ie.maxNumberPeriodicCSI_PerBWP_ForBeamReport = tmp_int_maxNumberPeriodicCSI_PerBWP_ForBeamReport
	var tmp_int_maxNumberAperiodicCSI_PerBWP_ForBeamReport int64
	if tmp_int_maxNumberAperiodicCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberAperiodicCSI_PerBWP_ForBeamReport", err)
	}
	ie.maxNumberAperiodicCSI_PerBWP_ForBeamReport = tmp_int_maxNumberAperiodicCSI_PerBWP_ForBeamReport
	if err = ie.maxNumberAperiodicCSI_triggeringStatePerCC.Decode(r); err != nil {
		return utils.WrapError("Decode maxNumberAperiodicCSI_triggeringStatePerCC", err)
	}
	var tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForBeamReport int64
	if tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForBeamReport, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 4}, false); err != nil {
		return utils.WrapError("ReadInteger maxNumberSemiPersistentCSI_PerBWP_ForBeamReport", err)
	}
	ie.maxNumberSemiPersistentCSI_PerBWP_ForBeamReport = tmp_int_maxNumberSemiPersistentCSI_PerBWP_ForBeamReport
	var tmp_int_simultaneousCSI_ReportsPerCC int64
	if tmp_int_simultaneousCSI_ReportsPerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger simultaneousCSI_ReportsPerCC", err)
	}
	ie.simultaneousCSI_ReportsPerCC = tmp_int_simultaneousCSI_ReportsPerCC
	return nil
}
