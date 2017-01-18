package client

import "encoding/xml"

type HandlerGetAssets struct {
	Request  *RequestGetAssets
	Response *ResponseGetAssets
	client   *Client
}

type RequestGetAssets struct {
	XMLName xml.Name `xml:"wd:Get_Assets_Request"`
	XMLNsWd string   `xml:"xmlns:wd,attr"`

	AssetReference                  *RequestObjectList `xml:"wd:Request_References>wd:Asset_Reference"`
	CompanyReference                *RequestObjectList `xml:"wd:Request_Criteria>wd:Company_Reference"`
	IssuedToWorkerReference         *RequestObjectList `xml:"wd:Request_Criteria>wd:Issued_To_Worker_Reference"`
	SpendCategoryReference          *RequestObjectList `xml:"wd:Request_Criteria>wd:Spend_Category_Reference"`
	ItemReference                   *RequestObjectList `xml:"wd:Request_Criteria>wd:Item_Reference"`
	LocationReference               *RequestObjectList `xml:"wd:Request_Criteria>wd:Location_Reference"`
	AssetIdentifier                 *string            `xml:"wd:Request_Criteria>wd:Asset_Identifier,omitempty"`
	SerialNumber                    *string            `xml:"wd:Request_Criteria>wd:Serial_Number,omitempty"`
	WorktagReference                *RequestObjectList `xml:"wd:Request_Criteria>wd:Worktag_Reference"`
	AssetStatusReference            *RequestObjectList `xml:"wd:Request_Criteria>wd:Asset_Status_Reference"`
	AcquisitionMethodReference      *RequestObjectList `xml:"wd:Request_Criteria>wd:Acquisition_Method_Reference"`
	DisposalMethodReference         *RequestObjectList `xml:"wd:Request_Criteria>wd:Disposal_Method_Reference"`
	AcquiredFromDate                *Date              `xml:"wd:Request_Criteria>wd:Acquired_From_Date"`
	AcquiredToDate                  *Date              `xml:"wd:Request_Criteria>wd:Acquired_To_Date"`
	UpdatedFromDate                 *DateTime          `xml:"wd:Request_Criteria>wd:Updated_From_Date"`
	UpdatedToDate                   *DateTime          `xml:"wd:Request_Criteria>wd:Updated_To_Date"`
	AsOfEffectiveDate               *Date              `xml:"wd:Response_Filter>wd:As_Of_Effective_Date"`
	AsOfEntryDateTime               *DateTime          `xml:"wd:Response_Filter>wd:As_Of_Entry_DateTime"`
	Page                            int                `xml:"wd:Response_Filter>wd:Page"`
	Count                           int                `xml:"wd:Response_Filter>wd:Count"`
	IncludeReference                Boolean            `xml:"wd:Response_Group>wd:Include_Reference,omitempty"`
	IncludeCustodianData            Boolean            `xml:"wd:Response_Group>wd:Include_Custodian_Data,omitempty"`
	IncludeDepreciationData         Boolean            `xml:"wd:Response_Group>wd:Include_Depreciation_Data,omitempty"`
	IncludeDepreciationDetailData   Boolean            `xml:"wd:Response_Group>wd:Include_Depreciation_Detail_Data,omitempty"`
	IncludeDisposalData             Boolean            `xml:"wd:Response_Group>wd:Include_Disposal_Data,omitempty"`
	IncludeIntercompanyTransferData Boolean            `xml:"wd:Response_Group>wd:Include_Intercompany_Transfer_Data,omitempty"`
	IncludeImpairmentData           Boolean            `xml:"wd:Response_Group>wd:Include_Impairment_Data,omitempty"`
	IncludeReinstatementData        Boolean            `xml:"wd:Response_Group>wd:Include_Reinstatement_Data,omitempty"`
	IncludeInServiceScheduleData    Boolean            `xml:"wd:Response_Group>wd:Include_In_Service_Schedule_Data,omitempty"`
	IncludeCostAdjustmentData       Boolean            `xml:"wd:Response_Group>wd:Include_Cost_Adjustment_Data,omitempty"`
	IncludeAttachmentData           Boolean            `xml:"wd:Response_Group>wd:Include_Attachment_Data,omitempty"`
	IncludeReclassificationData     Boolean            `xml:"wd:Response_Group>wd:Include_Reclassification_Data,omitempty"`
	IncludeRemovalData              Boolean            `xml:"wd:Response_Group>wd:Include_Removal_Data,omitempty"`
}

type ResponseGetAssets struct {
	Version                         string              `xml:"version,attr,omitempty"`
	AssetReference                  *ResponseObjectList `xml:"Request_References>Asset_Reference"`
	CompanyReference                *ResponseObjectList `xml:"Request_Criteria>Company_Reference"`
	IssuedToWorkerReference         *ResponseObjectList `xml:"Request_Criteria>Issued_To_Worker_Reference"`
	SpendCategoryReference          *ResponseObjectList `xml:"Request_Criteria>Spend_Category_Reference"`
	ItemReference                   *ResponseObjectList `xml:"Request_Criteria>Item_Reference"`
	LocationReference               *ResponseObjectList `xml:"Request_Criteria>Location_Reference"`
	AssetIdentifier                 string              `xml:"Request_Criteria>Asset_Identifier,omitempty"`
	SerialNumber                    string              `xml:"Request_Criteria>Serial_Number,omitempty"`
	WorktagReference                *ResponseObjectList `xml:"Request_Criteria>Worktag_Reference"`
	AssetStatusReference            *ResponseObjectList `xml:"Request_Criteria>Asset_Status_Reference"`
	AcquisitionMethodReference      *ResponseObjectList `xml:"Request_Criteria>Acquisition_Method_Reference"`
	DisposalMethodReference         *ResponseObjectList `xml:"Request_Criteria>Disposal_Method_Reference"`
	AcquiredFromDate                *Date               `xml:"Request_Criteria>Acquired_From_Date"`
	AcquiredToDate                  *Date               `xml:"Request_Criteria>Acquired_To_Date"`
	UpdatedFromDate                 *DateTime           `xml:"Request_Criteria>Updated_From_Date"`
	UpdatedToDate                   *DateTime           `xml:"Request_Criteria>Updated_To_Date"`
	AsOfEffectiveDate               *Date               `xml:"Response_Filter>As_Of_Effective_Date"`
	AsOfEntryDateTime               *DateTime           `xml:"Response_Filter>As_Of_Entry_DateTime"`
	Page                            int                 `xml:"Response_Filter>Page,omitempty"`
	Count                           int                 `xml:"Response_Filter>Count,omitempty"`
	IncludeReference                Boolean             `xml:"Response_Group>Include_Reference,omitempty"`
	IncludeCustodianData            Boolean             `xml:"Response_Group>Include_Custodian_Data,omitempty"`
	IncludeDepreciationData         Boolean             `xml:"Response_Group>Include_Depreciation_Data,omitempty"`
	IncludeDepreciationDetailData   Boolean             `xml:"Response_Group>Include_Depreciation_Detail_Data,omitempty"`
	IncludeDisposalData             Boolean             `xml:"Response_Group>Include_Disposal_Data,omitempty"`
	IncludeIntercompanyTransferData Boolean             `xml:"Response_Group>Include_Intercompany_Transfer_Data,omitempty"`
	IncludeImpairmentData           Boolean             `xml:"Response_Group>Include_Impairment_Data,omitempty"`
	IncludeReinstatementData        Boolean             `xml:"Response_Group>Include_Reinstatement_Data,omitempty"`
	IncludeInServiceScheduleData    Boolean             `xml:"Response_Group>Include_In_Service_Schedule_Data,omitempty"`
	IncludeCostAdjustmentData       Boolean             `xml:"Response_Group>Include_Cost_Adjustment_Data,omitempty"`
	IncludeAttachmentData           Boolean             `xml:"Response_Group>Include_Attachment_Data,omitempty"`
	IncludeReclassificationData     Boolean             `xml:"Response_Group>Include_Reclassification_Data,omitempty"`
	IncludeRemovalData              Boolean             `xml:"Response_Group>Include_Removal_Data,omitempty"`
	TotalResults                    int                 `xml:"Response_Results>Total_Results"`
	TotalPages                      int                 `xml:"Response_Results>Total_Pages"`
	PageResults                     int                 `xml:"Response_Results>Page_Results"`
	CurrentPage                     int                 `xml:"Response_Results>Page"`
	Assets                          []*Asset            `xml:"Response_Data>Asset"`
}

type Asset struct {
	AssetReference               *ResponseObject     `xml:"Asset_Reference"`
	AssetReferenceID             string              `xml:"Asset_Data>Asset_Reference_ID"`
	AssetID                      string              `xml:"Asset_Data>Asset_ID"`
	AssetName                    string              `xml:"Asset_Data>Asset_Name"`
	AssetDescription             string              `xml:"Asset_Data>Asset_Description"`
	TotalAssetCost               float32             `xml:"Asset_Data>Total_Asset_Cost"`
	Quantity                     int                 `xml:"Asset_Data>Quantity"`
	QuantityAvailable            int                 `xml:"Asset_Data>Quantity_Available"`
	ResidualValue                float32             `xml:"Asset_Data>Residual_Value"`
	DatePlacedInService          *DateTime           `xml:"Asset_Data>Date_Placed_In_Service"`
	AccountingTreatmentReference *ResponseObject     `xml:"Asset_Data>Accounting_Treatment_Reference"`
	AssetIdentifier              string              `xml:"Asset_Data>Asset_Identifier"`
	AssetStatusReference         *ResponseObject     `xml:"Asset_Data>Asset_Status_Reference"`
	SerialNumber                 string              `xml:"Asset_Data>Serial_Number"`
	Manufacturer                 string              `xml:"Asset_Data>Manufacturer"`
	BusinessUsePercentage        int                 `xml:"Asset_Data>Business_Use_Percentage"`
	EventInProgress              bool                `xml:"Asset_Data>Event_In_Progress"`
	AcquisitionCost              float32             `xml:"Asset_Data>Acquisition_Data>Acquisition_Cost"`
	AcquisitionDate              *Date               `xml:"Asset_Data>Acquisition_Data>Acquisition_Date"`
	PONumber                     string              `xml:"Asset_Data>Acquisition_Data>PO_Number"`
	ReceiptNumber                string              `xml:"Asset_Data>Acquisition_Data>Receipt_Number"`
	SupplierInvoiceNumber        string              `xml:"Asset_Data>Acquisition_Data>Supplier_Invoice_Number"`
	ResidualValueAtAcquisition   float32             `xml:"Asset_Data>Acquisition_Data>Residual_Value_At_Acquisition"`
	WorktagReference             *ResponseObjectList `xml:"Asset_Data>Acquisition_Data>Worktag_Reference"`
	ProjectNumber                string              `xml:"Asset_Data>Acquisition_Data>Project_Number"`
	ContractNumber               string              `xml:"Asset_Data>Acquisition_Data>Contract_Number"`
	ContractLineStartDate        *Date               `xml:"Asset_Data>Acquisition_Data>Contract_Line_Start_Date"`
	ContractLineEndDate          *Date               `xml:"Asset_Data>Acquisition_Data>Contract_Line_End_Date"`
}

func NewRequestGetAssets(clnt *Client) *HandlerGetAssets {
	return &HandlerGetAssets{
		Request: &RequestGetAssets{
			XMLNsWd: XMLNsWd,
			Count:   clnt.PageSize,
		},
		client: clnt,
	}
}

func (h *HandlerGetAssets) get() error {
	env, err := h.client.get(&Request{h.Request})
	if err != nil {
		return err
	}

	// Empty response
	if env.Body.GetAssets == nil {
		return nil
	}

	h.Response = env.Body.GetAssets
	return nil
}

func (h *HandlerGetAssets) GetRequestXML() ([]byte, error) {
	return h.client.getRequestXML(&Request{h.Request})
}

func (h *HandlerGetAssets) GetNextXML() ([]byte, error) {
	h.Request.Page++
	return h.client.getXML(&Request{h.Request})
}

func (h *HandlerGetAssets) GetNext() error {
	h.Request.Page++
	return h.get()
}

func (h *HandlerGetAssets) Last() bool {
	if h.Response == nil {
		return false
	}
	if h.Response.CurrentPage == h.Response.TotalPages {
		return true
	}
	return false
}
