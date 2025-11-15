package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeighbourCellInfo_r17 struct {
	epochTime_r17     EpochTime_r17     `madatory`
	ephemerisInfo_r17 EphemerisInfo_r17 `madatory`
}

func (ie *NeighbourCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.epochTime_r17.Encode(w); err != nil {
		return utils.WrapError("Encode epochTime_r17", err)
	}
	if err = ie.ephemerisInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ephemerisInfo_r17", err)
	}
	return nil
}

func (ie *NeighbourCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.epochTime_r17.Decode(r); err != nil {
		return utils.WrapError("Decode epochTime_r17", err)
	}
	if err = ie.ephemerisInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ephemerisInfo_r17", err)
	}
	return nil
}
