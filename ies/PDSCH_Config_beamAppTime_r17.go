package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

const (
	PDSCH_Config_beamAppTime_r17_Enum_n1     uper.Enumerated = 0
	PDSCH_Config_beamAppTime_r17_Enum_n2     uper.Enumerated = 1
	PDSCH_Config_beamAppTime_r17_Enum_n4     uper.Enumerated = 2
	PDSCH_Config_beamAppTime_r17_Enum_n7     uper.Enumerated = 3
	PDSCH_Config_beamAppTime_r17_Enum_n14    uper.Enumerated = 4
	PDSCH_Config_beamAppTime_r17_Enum_n28    uper.Enumerated = 5
	PDSCH_Config_beamAppTime_r17_Enum_n42    uper.Enumerated = 6
	PDSCH_Config_beamAppTime_r17_Enum_n56    uper.Enumerated = 7
	PDSCH_Config_beamAppTime_r17_Enum_n70    uper.Enumerated = 8
	PDSCH_Config_beamAppTime_r17_Enum_n84    uper.Enumerated = 9
	PDSCH_Config_beamAppTime_r17_Enum_n98    uper.Enumerated = 10
	PDSCH_Config_beamAppTime_r17_Enum_n112   uper.Enumerated = 11
	PDSCH_Config_beamAppTime_r17_Enum_n224   uper.Enumerated = 12
	PDSCH_Config_beamAppTime_r17_Enum_n336   uper.Enumerated = 13
	PDSCH_Config_beamAppTime_r17_Enum_spare2 uper.Enumerated = 14
	PDSCH_Config_beamAppTime_r17_Enum_spare1 uper.Enumerated = 15
)

type PDSCH_Config_beamAppTime_r17 struct {
	Value uper.Enumerated `lb:0,ub:15,madatory`
}

func (ie *PDSCH_Config_beamAppTime_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteEnumerate(uint64(ie.Value), uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Encode PDSCH_Config_beamAppTime_r17", err)
	}
	return nil
}

func (ie *PDSCH_Config_beamAppTime_r17) Decode(r *uper.UperReader) error {
	var err error
	var v uint64
	if v, err = r.ReadEnumerate(uper.Constraint{Lb: 0, Ub: 15}, false); err != nil {
		return utils.WrapError("Decode PDSCH_Config_beamAppTime_r17", err)
	}
	ie.Value = uper.Enumerated(v)
	return nil
}
