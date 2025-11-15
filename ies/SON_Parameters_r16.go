package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SON_Parameters_r16 struct {
	rach_Report_r16        *SON_Parameters_r16_rach_Report_r16        `optional`
	rlfReportCHO_r17       *SON_Parameters_r16_rlfReportCHO_r17       `optional,ext-1`
	rlfReportDAPS_r17      *SON_Parameters_r16_rlfReportDAPS_r17      `optional,ext-1`
	success_HO_Report_r17  *SON_Parameters_r16_success_HO_Report_r17  `optional,ext-1`
	twoStepRACH_Report_r17 *SON_Parameters_r16_twoStepRACH_Report_r17 `optional,ext-1`
	pscell_MHI_Report_r17  *SON_Parameters_r16_pscell_MHI_Report_r17  `optional,ext-1`
	onDemandSI_Report_r17  *SON_Parameters_r16_onDemandSI_Report_r17  `optional,ext-1`
}

func (ie *SON_Parameters_r16) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.rlfReportCHO_r17 != nil || ie.rlfReportDAPS_r17 != nil || ie.success_HO_Report_r17 != nil || ie.twoStepRACH_Report_r17 != nil || ie.pscell_MHI_Report_r17 != nil || ie.onDemandSI_Report_r17 != nil
	preambleBits := []bool{hasExtensions, ie.rach_Report_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.rach_Report_r16 != nil {
		if err = ie.rach_Report_r16.Encode(w); err != nil {
			return utils.WrapError("Encode rach_Report_r16", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.rlfReportCHO_r17 != nil || ie.rlfReportDAPS_r17 != nil || ie.success_HO_Report_r17 != nil || ie.twoStepRACH_Report_r17 != nil || ie.pscell_MHI_Report_r17 != nil || ie.onDemandSI_Report_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SON_Parameters_r16", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.rlfReportCHO_r17 != nil, ie.rlfReportDAPS_r17 != nil, ie.success_HO_Report_r17 != nil, ie.twoStepRACH_Report_r17 != nil, ie.pscell_MHI_Report_r17 != nil, ie.onDemandSI_Report_r17 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode rlfReportCHO_r17 optional
			if ie.rlfReportCHO_r17 != nil {
				if err = ie.rlfReportCHO_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlfReportCHO_r17", err)
				}
			}
			// encode rlfReportDAPS_r17 optional
			if ie.rlfReportDAPS_r17 != nil {
				if err = ie.rlfReportDAPS_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode rlfReportDAPS_r17", err)
				}
			}
			// encode success_HO_Report_r17 optional
			if ie.success_HO_Report_r17 != nil {
				if err = ie.success_HO_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode success_HO_Report_r17", err)
				}
			}
			// encode twoStepRACH_Report_r17 optional
			if ie.twoStepRACH_Report_r17 != nil {
				if err = ie.twoStepRACH_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode twoStepRACH_Report_r17", err)
				}
			}
			// encode pscell_MHI_Report_r17 optional
			if ie.pscell_MHI_Report_r17 != nil {
				if err = ie.pscell_MHI_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode pscell_MHI_Report_r17", err)
				}
			}
			// encode onDemandSI_Report_r17 optional
			if ie.onDemandSI_Report_r17 != nil {
				if err = ie.onDemandSI_Report_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode onDemandSI_Report_r17", err)
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

func (ie *SON_Parameters_r16) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var rach_Report_r16Present bool
	if rach_Report_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if rach_Report_r16Present {
		ie.rach_Report_r16 = new(SON_Parameters_r16_rach_Report_r16)
		if err = ie.rach_Report_r16.Decode(r); err != nil {
			return utils.WrapError("Decode rach_Report_r16", err)
		}
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

			rlfReportCHO_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			rlfReportDAPS_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			success_HO_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			twoStepRACH_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			pscell_MHI_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			onDemandSI_Report_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode rlfReportCHO_r17 optional
			if rlfReportCHO_r17Present {
				ie.rlfReportCHO_r17 = new(SON_Parameters_r16_rlfReportCHO_r17)
				if err = ie.rlfReportCHO_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rlfReportCHO_r17", err)
				}
			}
			// decode rlfReportDAPS_r17 optional
			if rlfReportDAPS_r17Present {
				ie.rlfReportDAPS_r17 = new(SON_Parameters_r16_rlfReportDAPS_r17)
				if err = ie.rlfReportDAPS_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode rlfReportDAPS_r17", err)
				}
			}
			// decode success_HO_Report_r17 optional
			if success_HO_Report_r17Present {
				ie.success_HO_Report_r17 = new(SON_Parameters_r16_success_HO_Report_r17)
				if err = ie.success_HO_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode success_HO_Report_r17", err)
				}
			}
			// decode twoStepRACH_Report_r17 optional
			if twoStepRACH_Report_r17Present {
				ie.twoStepRACH_Report_r17 = new(SON_Parameters_r16_twoStepRACH_Report_r17)
				if err = ie.twoStepRACH_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode twoStepRACH_Report_r17", err)
				}
			}
			// decode pscell_MHI_Report_r17 optional
			if pscell_MHI_Report_r17Present {
				ie.pscell_MHI_Report_r17 = new(SON_Parameters_r16_pscell_MHI_Report_r17)
				if err = ie.pscell_MHI_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode pscell_MHI_Report_r17", err)
				}
			}
			// decode onDemandSI_Report_r17 optional
			if onDemandSI_Report_r17Present {
				ie.onDemandSI_Report_r17 = new(SON_Parameters_r16_onDemandSI_Report_r17)
				if err = ie.onDemandSI_Report_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode onDemandSI_Report_r17", err)
				}
			}
		}
	}
	return nil
}
