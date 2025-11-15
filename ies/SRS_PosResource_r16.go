package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_PosResource_r16 struct {
	srs_PosResourceId_r16      SRS_PosResourceId_r16                          `madatory`
	transmissionComb_r16       SRS_PosResource_r16_transmissionComb_r16       `lb:0,ub:1,madatory`
	resourceMapping_r16        SRS_PosResource_r16_resourceMapping_r16        `lb:0,ub:13,madatory,ext`
	freqDomainShift_r16        int64                                          `lb:0,ub:268,madatory,ext`
	freqHopping_r16            SRS_PosResource_r16_freqHopping_r16            `lb:0,ub:63,madatory,ext`
	groupOrSequenceHopping_r16 SRS_PosResource_r16_groupOrSequenceHopping_r16 `madatory,ext`
	resourceType_r16           Resourcetype_r16_SRS_PosResource_r16           `madatory,ext`
	sequenceId_r16             int64                                          `lb:0,ub:65535,madatory,ext`
	spatialRelationInfoPos_r16 *SRS_SpatialRelationInfoPos_r16                `optional,ext`
}

func (ie *SRS_PosResource_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.srs_PosResourceId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode srs_PosResourceId_r16", err)
	}
	if err = ie.transmissionComb_r16.Encode(w); err != nil {
		return utils.WrapError("Encode transmissionComb_r16", err)
	}
	return nil
}

func (ie *SRS_PosResource_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.srs_PosResourceId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode srs_PosResourceId_r16", err)
	}
	if err = ie.transmissionComb_r16.Decode(r); err != nil {
		return utils.WrapError("Decode transmissionComb_r16", err)
	}
	return nil
}
