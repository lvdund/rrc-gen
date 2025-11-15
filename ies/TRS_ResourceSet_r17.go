package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type TRS_ResourceSet_r17 struct {
	powerControlOffsetSS_r17        TRS_ResourceSet_r17_powerControlOffsetSS_r17 `madatory`
	scramblingID_Info_r17           TRS_ResourceSet_r17_scramblingID_Info_r17    `lb:2,ub:2,madatory`
	firstOFDMSymbolInTimeDomain_r17 int64                                        `lb:0,ub:9,madatory,ext`
	startingRB_r17                  int64                                        `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory,ext`
	nrofRBs_r17                     int64                                        `lb:24,ub:maxNrofPhysicalResourceBlocksPlus1,madatory,ext`
	ssb_Index_r17                   SSB_Index                                    `madatory,ext`
	periodicityAndOffset_r17        TRS_ResourceSet_r17_periodicityAndOffset_r17 `lb:0,ub:9,madatory,ext`
	frequencyDomainAllocation_r17   uper.BitString                               `lb:4,ub:4,madatory,ext`
	indBitID_r17                    int64                                        `lb:0,ub:5,madatory,ext`
	nrofResources_r17               TRS_ResourceSet_r17_nrofResources_r17        `madatory,ext`
}

func (ie *TRS_ResourceSet_r17) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.powerControlOffsetSS_r17.Encode(w); err != nil {
		return utils.WrapError("Encode powerControlOffsetSS_r17", err)
	}
	if err = ie.scramblingID_Info_r17.Encode(w); err != nil {
		return utils.WrapError("Encode scramblingID_Info_r17", err)
	}
	return nil
}

func (ie *TRS_ResourceSet_r17) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.powerControlOffsetSS_r17.Decode(r); err != nil {
		return utils.WrapError("Decode powerControlOffsetSS_r17", err)
	}
	if err = ie.scramblingID_Info_r17.Decode(r); err != nil {
		return utils.WrapError("Decode scramblingID_Info_r17", err)
	}
	return nil
}
