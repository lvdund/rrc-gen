package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReportSFTD_NR struct {
	reportSFTD_Meas           bool                                `madatory`
	reportRSRP                bool                                `madatory`
	reportSFTD_NeighMeas      *ReportSFTD_NR_reportSFTD_NeighMeas `optional,ext-1`
	drx_SFTD_NeighMeas        *ReportSFTD_NR_drx_SFTD_NeighMeas   `optional,ext-1`
	cellsForWhichToReportSFTD []PhysCellId                        `lb:1,ub:maxCellSFTD,optional,ext-1`
}

func (ie *ReportSFTD_NR) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.reportSFTD_NeighMeas != nil || ie.drx_SFTD_NeighMeas != nil || len(ie.cellsForWhichToReportSFTD) > 0
	preambleBits := []bool{hasExtensions}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBoolean(ie.reportSFTD_Meas); err != nil {
		return utils.WrapError("WriteBoolean reportSFTD_Meas", err)
	}
	if err = w.WriteBoolean(ie.reportRSRP); err != nil {
		return utils.WrapError("WriteBoolean reportRSRP", err)
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.reportSFTD_NeighMeas != nil || ie.drx_SFTD_NeighMeas != nil || len(ie.cellsForWhichToReportSFTD) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap ReportSFTD_NR", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.reportSFTD_NeighMeas != nil, ie.drx_SFTD_NeighMeas != nil, len(ie.cellsForWhichToReportSFTD) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode reportSFTD_NeighMeas optional
			if ie.reportSFTD_NeighMeas != nil {
				if err = ie.reportSFTD_NeighMeas.Encode(extWriter); err != nil {
					return utils.WrapError("Encode reportSFTD_NeighMeas", err)
				}
			}
			// encode drx_SFTD_NeighMeas optional
			if ie.drx_SFTD_NeighMeas != nil {
				if err = ie.drx_SFTD_NeighMeas.Encode(extWriter); err != nil {
					return utils.WrapError("Encode drx_SFTD_NeighMeas", err)
				}
			}
			// encode cellsForWhichToReportSFTD optional
			if len(ie.cellsForWhichToReportSFTD) > 0 {
				tmp_cellsForWhichToReportSFTD := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
				for _, i := range ie.cellsForWhichToReportSFTD {
					tmp_cellsForWhichToReportSFTD.Value = append(tmp_cellsForWhichToReportSFTD.Value, &i)
				}
				if err = tmp_cellsForWhichToReportSFTD.Encode(extWriter); err != nil {
					return utils.WrapError("Encode cellsForWhichToReportSFTD", err)
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

func (ie *ReportSFTD_NR) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bool_reportSFTD_Meas bool
	if tmp_bool_reportSFTD_Meas, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportSFTD_Meas", err)
	}
	ie.reportSFTD_Meas = tmp_bool_reportSFTD_Meas
	var tmp_bool_reportRSRP bool
	if tmp_bool_reportRSRP, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean reportRSRP", err)
	}
	ie.reportRSRP = tmp_bool_reportRSRP

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

			reportSFTD_NeighMeasPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			drx_SFTD_NeighMeasPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			cellsForWhichToReportSFTDPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode reportSFTD_NeighMeas optional
			if reportSFTD_NeighMeasPresent {
				ie.reportSFTD_NeighMeas = new(ReportSFTD_NR_reportSFTD_NeighMeas)
				if err = ie.reportSFTD_NeighMeas.Decode(extReader); err != nil {
					return utils.WrapError("Decode reportSFTD_NeighMeas", err)
				}
			}
			// decode drx_SFTD_NeighMeas optional
			if drx_SFTD_NeighMeasPresent {
				ie.drx_SFTD_NeighMeas = new(ReportSFTD_NR_drx_SFTD_NeighMeas)
				if err = ie.drx_SFTD_NeighMeas.Decode(extReader); err != nil {
					return utils.WrapError("Decode drx_SFTD_NeighMeas", err)
				}
			}
			// decode cellsForWhichToReportSFTD optional
			if cellsForWhichToReportSFTDPresent {
				tmp_cellsForWhichToReportSFTD := utils.NewSequence[*PhysCellId]([]*PhysCellId{}, uper.Constraint{Lb: 1, Ub: maxCellSFTD}, false)
				fn_cellsForWhichToReportSFTD := func() *PhysCellId {
					return new(PhysCellId)
				}
				if err = tmp_cellsForWhichToReportSFTD.Decode(extReader, fn_cellsForWhichToReportSFTD); err != nil {
					return utils.WrapError("Decode cellsForWhichToReportSFTD", err)
				}
				ie.cellsForWhichToReportSFTD = []PhysCellId{}
				for _, i := range tmp_cellsForWhichToReportSFTD.Value {
					ie.cellsForWhichToReportSFTD = append(ie.cellsForWhichToReportSFTD, *i)
				}
			}
		}
	}
	return nil
}
