package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	RA_Report_r16_raPurpose_r16_Enum_accessRelated             uper.Enumerated = 0
	RA_Report_r16_raPurpose_r16_Enum_beamFailureRecovery       uper.Enumerated = 1
	RA_Report_r16_raPurpose_r16_Enum_reconfigurationWithSync   uper.Enumerated = 2
	RA_Report_r16_raPurpose_r16_Enum_ulUnSynchronized          uper.Enumerated = 3
	RA_Report_r16_raPurpose_r16_Enum_schedulingRequestFailure  uper.Enumerated = 4
	RA_Report_r16_raPurpose_r16_Enum_noPUCCHResourceAvailable  uper.Enumerated = 5
	RA_Report_r16_raPurpose_r16_Enum_requestForOtherSI         uper.Enumerated = 6
	RA_Report_r16_raPurpose_r16_Enum_msg3RequestForOtherSI_r17 uper.Enumerated = 7
	RA_Report_r16_raPurpose_r16_Enum_spare8                    uper.Enumerated = 8
	RA_Report_r16_raPurpose_r16_Enum_spare7                    uper.Enumerated = 9
	RA_Report_r16_raPurpose_r16_Enum_spare6                    uper.Enumerated = 10
	RA_Report_r16_raPurpose_r16_Enum_spare5                    uper.Enumerated = 11
	RA_Report_r16_raPurpose_r16_Enum_spare4                    uper.Enumerated = 12
	RA_Report_r16_raPurpose_r16_Enum_spare3                    uper.Enumerated = 13
	RA_Report_r16_raPurpose_r16_Enum_spare2                    uper.Enumerated = 14
	RA_Report_r16_raPurpose_r16_Enum_spare1                    uper.Enumerated = 15
)

type RA_Report_r16_raPurpose_r16 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *RA_Report_r16_raPurpose_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode RA_Report_r16_raPurpose_r16", err)
	}
	return nil
}

func (ie *RA_Report_r16_raPurpose_r16) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode RA_Report_r16_raPurpose_r16", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
