package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_AssociatedReportConfigInfo struct {
	reportConfigId                      CSI_ReportConfigId                                       `madatory`
	resourcesForChannel                 *CSI_AssociatedReportConfigInfo_resourcesForChannel      `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional`
	csi_IM_ResourcesForInterference     *int64                                                   `lb:1,ub:maxNrofCSI_IM_ResourceSetsPerConfig,optional`
	nzp_CSI_RS_ResourcesForInterference *int64                                                   `lb:1,ub:maxNrofNZP_CSI_RS_ResourceSetsPerConfig,optional`
	resourcesForChannel2_r17            *CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17 `lb:1,ub:maxNrofAP_CSI_RS_ResourcesPerSet,optional,ext-1`
	csi_SSB_ResourceSetExt              *int64                                                   `lb:1,ub:maxNrofCSI_SSB_ResourceSetsPerConfigExt,optional,ext-1`
}

func (ie *CSI_AssociatedReportConfigInfo) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.resourcesForChannel2_r17 != nil || ie.csi_SSB_ResourceSetExt != nil
	preambleBits := []bool{hasExtensions, ie.resourcesForChannel != nil, ie.csi_IM_ResourcesForInterference != nil, ie.nzp_CSI_RS_ResourcesForInterference != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.reportConfigId.Encode(w); err != nil {
		return utils.WrapError("Encode reportConfigId", err)
	}
	if ie.resourcesForChannel != nil {
		if err = ie.resourcesForChannel.Encode(w); err != nil {
			return utils.WrapError("Encode resourcesForChannel", err)
		}
	}
	if ie.csi_IM_ResourcesForInterference != nil {
		if err = w.WriteInteger(*ie.csi_IM_ResourcesForInterference, &uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Encode csi_IM_ResourcesForInterference", err)
		}
	}
	if ie.nzp_CSI_RS_ResourcesForInterference != nil {
		if err = w.WriteInteger(*ie.nzp_CSI_RS_ResourcesForInterference, &uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Encode nzp_CSI_RS_ResourcesForInterference", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.resourcesForChannel2_r17 != nil || ie.csi_SSB_ResourceSetExt != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap CSI_AssociatedReportConfigInfo", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.resourcesForChannel2_r17 != nil, ie.csi_SSB_ResourceSetExt != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode resourcesForChannel2_r17 optional
			if ie.resourcesForChannel2_r17 != nil {
				if err = ie.resourcesForChannel2_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode resourcesForChannel2_r17", err)
				}
			}
			// encode csi_SSB_ResourceSetExt optional
			if ie.csi_SSB_ResourceSetExt != nil {
				if err = extWriter.WriteInteger(*ie.csi_SSB_ResourceSetExt, &uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
					return utils.WrapError("Encode csi_SSB_ResourceSetExt", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *CSI_AssociatedReportConfigInfo) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var resourcesForChannelPresent bool
	if resourcesForChannelPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var csi_IM_ResourcesForInterferencePresent bool
	if csi_IM_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var nzp_CSI_RS_ResourcesForInterferencePresent bool
	if nzp_CSI_RS_ResourcesForInterferencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.reportConfigId.Decode(r); err != nil {
		return utils.WrapError("Decode reportConfigId", err)
	}
	if resourcesForChannelPresent {
		ie.resourcesForChannel = new(CSI_AssociatedReportConfigInfo_resourcesForChannel)
		if err = ie.resourcesForChannel.Decode(r); err != nil {
			return utils.WrapError("Decode resourcesForChannel", err)
		}
	}
	if csi_IM_ResourcesForInterferencePresent {
		var tmp_int_csi_IM_ResourcesForInterference int64
		if tmp_int_csi_IM_ResourcesForInterference, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofCSI_IM_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Decode csi_IM_ResourcesForInterference", err)
		}
		ie.csi_IM_ResourcesForInterference = &tmp_int_csi_IM_ResourcesForInterference
	}
	if nzp_CSI_RS_ResourcesForInterferencePresent {
		var tmp_int_nzp_CSI_RS_ResourcesForInterference int64
		if tmp_int_nzp_CSI_RS_ResourcesForInterference, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourceSetsPerConfig}, false); err != nil {
			return utils.WrapError("Decode nzp_CSI_RS_ResourcesForInterference", err)
		}
		ie.nzp_CSI_RS_ResourcesForInterference = &tmp_int_nzp_CSI_RS_ResourcesForInterference
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			resourcesForChannel2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			csi_SSB_ResourceSetExtPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode resourcesForChannel2_r17 optional
			if resourcesForChannel2_r17Present {
				ie.resourcesForChannel2_r17 = new(CSI_AssociatedReportConfigInfo_resourcesForChannel2_r17)
				if err = ie.resourcesForChannel2_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode resourcesForChannel2_r17", err)
				}
			}
			// decode csi_SSB_ResourceSetExt optional
			if csi_SSB_ResourceSetExtPresent {
				var tmp_int_csi_SSB_ResourceSetExt int64
				if tmp_int_csi_SSB_ResourceSetExt, err = extReader.ReadInteger(&uper.Constraint{Lb: 1, Ub: maxNrofCSI_SSB_ResourceSetsPerConfigExt}, false); err != nil {
					return utils.WrapError("Decode csi_SSB_ResourceSetExt", err)
				}
				ie.csi_SSB_ResourceSetExt = &tmp_int_csi_SSB_ResourceSetExt
			}
		}
	}
	return nil
}
