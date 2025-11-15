package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDSCH_ConfigPTM_r17 struct {
	dataScramblingIdentityPDSCH_r17 *int64                                           `lb:0,ub:1023,optional`
	dmrs_ScramblingID0_r17          *int64                                           `lb:0,ub:65535,optional`
	pdsch_AggregationFactor_r17     *PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17 `optional`
}

func (ie *PDSCH_ConfigPTM_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.dataScramblingIdentityPDSCH_r17 != nil, ie.dmrs_ScramblingID0_r17 != nil, ie.pdsch_AggregationFactor_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.dataScramblingIdentityPDSCH_r17 != nil {
		if err = w.WriteInteger(*ie.dataScramblingIdentityPDSCH_r17, &uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Encode dataScramblingIdentityPDSCH_r17", err)
		}
	}
	if ie.dmrs_ScramblingID0_r17 != nil {
		if err = w.WriteInteger(*ie.dmrs_ScramblingID0_r17, &uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Encode dmrs_ScramblingID0_r17", err)
		}
	}
	if ie.pdsch_AggregationFactor_r17 != nil {
		if err = ie.pdsch_AggregationFactor_r17.Encode(w); err != nil {
			return utils.WrapError("Encode pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}

func (ie *PDSCH_ConfigPTM_r17) Decode(r *uper.UperReader) error {
	var err error
	var dataScramblingIdentityPDSCH_r17Present bool
	if dataScramblingIdentityPDSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var dmrs_ScramblingID0_r17Present bool
	if dmrs_ScramblingID0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var pdsch_AggregationFactor_r17Present bool
	if pdsch_AggregationFactor_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if dataScramblingIdentityPDSCH_r17Present {
		var tmp_int_dataScramblingIdentityPDSCH_r17 int64
		if tmp_int_dataScramblingIdentityPDSCH_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
			return utils.WrapError("Decode dataScramblingIdentityPDSCH_r17", err)
		}
		ie.dataScramblingIdentityPDSCH_r17 = &tmp_int_dataScramblingIdentityPDSCH_r17
	}
	if dmrs_ScramblingID0_r17Present {
		var tmp_int_dmrs_ScramblingID0_r17 int64
		if tmp_int_dmrs_ScramblingID0_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
			return utils.WrapError("Decode dmrs_ScramblingID0_r17", err)
		}
		ie.dmrs_ScramblingID0_r17 = &tmp_int_dmrs_ScramblingID0_r17
	}
	if pdsch_AggregationFactor_r17Present {
		ie.pdsch_AggregationFactor_r17 = new(PDSCH_ConfigPTM_r17_pdsch_AggregationFactor_r17)
		if err = ie.pdsch_AggregationFactor_r17.Decode(r); err != nil {
			return utils.WrapError("Decode pdsch_AggregationFactor_r17", err)
		}
	}
	return nil
}
