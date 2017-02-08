package certcenter

// Profile fetches basic informations about your profile
//
func Profile() (*ProfileResult, error) {
	req := new(apiRequest)
	req.result = new(ProfileResult)
	err := req.do("Profile")
	checkErr(err)
	return req.result.(*ProfileResult), err
}

// Limit inquires information about your current limit and used amount
//
func Limit() (*LimitResult, error) {
	req := new(apiRequest)
	req.result = new(LimitResult)
	err := req.do("Limit")
	checkErr(err)
	return req.result.(*LimitResult), err
}

// Products allows you to fetch a list of valid ProductCodes
//
func Products() (*ProductsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductsResult)
	err := req.do("Products")
	checkErr(err)
	return req.result.(*ProductsResult), err
}

// ProductDetails inquires detailed information on a particular ProductCode
//
func ProductDetails(ProductCode string) (*ProductDetailsResult, error) {
	req := new(apiRequest)
	req.result = new(ProductDetailsResult)
	req.request = &ProductDetailsRequest{
		ProductCode: ProductCode,
	}
	err := req.do("ProductDetails", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*ProductDetailsResult), err
}

// Quote allows you to generate an individual real-time quotation
//
func Quote(request *QuoteRequest) (*QuoteResult, error) {
	req := new(apiRequest)
	req.result = new(QuoteResult)
	req.request = request
	err := req.do("Quote", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*QuoteResult), err
}

// ValidateCSR allows you to parse and validate a PEM-encoded PKCS#10
//
func ValidateCSR(request *ValidateCSRRequest) (*ValidateCSRResult, error) {
	req := new(apiRequest)
	req.result = new(ValidateCSRResult)
	req.request = request
	err := req.do("ValidateCSR", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*ValidateCSRResult), err
}

// UserAgreement fetches the latest subscriber agreement from the CA
//
func UserAgreement(ProductCode string) (*UserAgreementResult, error) {
	req := new(apiRequest)
	req.result = new(UserAgreementResult)
	req.request = &UserAgreementRequest{
		ProductCode: ProductCode,
	}
	err := req.do("UserAgreement", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*UserAgreementResult), err
}

// ApproverList will fetch a list of valid email addresses
// for a particular CommonName and ProductCode
func ApproverList(request *ApproverListRequest) (*ApproverListResult, error) {
	req := new(apiRequest)
	req.result = new(ApproverListResult)
	req.request = request
	err := req.do("ApproverList", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*ApproverListResult), err
}

// Order allows you to submit orders for regular certificates
// as well as S/MIME and AlwaysOnSSL certificates
//
func Order(request *OrderRequest) (*OrderResult, error) {
	req := new(apiRequest)
	req.result = new(OrderResult)
	req.request = request
	err := req.do("Order", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*OrderResult), err
}

// PutApproverEmail allows you to reset the email address of the approver
//
func PutApproverEmail(request *PutApproverEmailRequest) (*PutApproverEmailResult, error) {
	req := new(apiRequest)
	req.result = new(PutApproverEmailResult)
	req.request = request
	err := req.do("ApproverEmail", CC_PARAM_TYPE_QS|CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*PutApproverEmailResult), err
}

// ResendApproverEmail allows you to resend the approver email to the approvers address
//
func ResendApproverEmail(request *ResendApproverEmailRequest) (*ResendApproverEmailResult, error) {
	req := new(apiRequest)
	req.result = new(ResendApproverEmailResult)
	req.request = request
	err := req.do("ApproverEmail", CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*ResendApproverEmailResult), err
}

// GetOrders gives you the capability to query and filter your orders
//
func GetOrders(request *GetOrdersRequest) (*GetOrdersResult, error) {
	req := new(apiRequest)
	req.result = new(GetOrdersResult)
	req.request = request
	err := req.do("Orders", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*GetOrdersResult), err
}

// GetModifiedOrders fetches modified orders. You can provide
// a timespan to specify which changes your're interested in
//
func GetModifiedOrders(request *GetModifiedOrdersRequest) (*GetModifiedOrdersResult, error) {
	req := new(apiRequest)
	req.result = new(GetModifiedOrdersResult)
	req.request = request
	err := req.do("GetModifiedOrders", CC_PARAM_TYPE_QS)
	checkErr(err)
	return req.result.(*GetModifiedOrdersResult), err
}

// GetOrder gives you the capability to query a particular order
//
func GetOrder(request *GetOrderRequest) (*GetOrderResult, error) {
	req := new(apiRequest)
	req.result = new(GetOrderResult)
	req.request = request
	err := req.do("Order", CC_PARAM_TYPE_QS|CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*GetOrderResult), err
}

// DeleteOrder gives you the capability to cancel a order
//
func DeleteOrder(request *DeleteOrderRequest) (*DeleteOrderResult, error) {
	req := new(apiRequest)
	req.result = new(DeleteOrderResult)
	req.request = request
	err := req.do("Order", CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*DeleteOrderResult), err
}

// Reissue allows you to replace an existent certificate in case
// of a key loss or algorithm/key-size upgrade
//
func Reissue(request *ReissueRequest) (*ReissueResult, error) {
	req := new(apiRequest)
	req.result = new(ReissueResult)
	req.request = request
	err := req.do("Reissue", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*ReissueResult), err
}

// Revoke allows you to mark a certificate as invalid.
//
func Revoke(request *RevokeRequest) (*RevokeResult, error) {
	req := new(apiRequest)
	req.result = new(RevokeResult)
	req.request = request
	err := req.do("Revoke", CC_PARAM_TYPE_BODY|CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*RevokeResult), err
}

// ValidateName checks a CommonName against the Symantec EE blacklist
// (AlwaysOnSSL/Symantec EE only)
//
func ValidateName(request *ValidateNameRequest) (*ValidateNameResult, error) {
	req := new(apiRequest)
	req.result = new(ValidateNameResult)
	req.request = request
	err := req.do("ValidateName", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*ValidateNameResult), err
}

// DNSData retrieve appropriate data for DNS based validation
// (AlwaysOnSSL/Symantec EE only)
//
func DNSData(request *DNSDataRequest) (*DNSDataResult, error) {
	req := new(apiRequest)
	req.result = new(DNSDataResult)
	req.request = request
	err := req.do("DNSData", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*DNSDataResult), err
}

// FileData retrieve appropriate data for FILE based validation
// (AlwaysOnSSL/Symantec EE only)
//
func FileData(request *FileDataRequest) (*FileDataResult, error) {
	req := new(apiRequest)
	req.result = new(FileDataResult)
	req.request = request
	err := req.do("FileData", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*FileDataResult), err
}

// VulnerabilityAssessment allows you to configure the
// Vulnerability Assessment (Symantec certificates, only!)
//
func VulnerabilityAssessment(request *VulnerabilityAssessmentRequest) (*VulnerabilityAssessmentResult, error) {
	req := new(apiRequest)
	req.result = new(VulnerabilityAssessmentResult)
	req.request = request
	err := req.do("VulnerabilityAssessment", CC_PARAM_TYPE_BODY)
	checkErr(err)
	return req.result.(*VulnerabilityAssessmentResult), err
}

// VulnerabilityAssessmentRescan let you initiate a re-scan for a certain order
//
func VulnerabilityAssessmentRescan(request *VulnerabilityAssessmentRescanRequest) (*VulnerabilityAssessmentRescanResult, error) {
	req := new(apiRequest)
	req.result = new(VulnerabilityAssessmentRescanResult)
	req.request = request
	err := req.do("VulnerabilityAssessment", CC_PARAM_TYPE_PATH)
	checkErr(err)
	return req.result.(*VulnerabilityAssessmentRescanResult), err
}
