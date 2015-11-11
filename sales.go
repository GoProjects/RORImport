// sales
package main

import (
	"database/sql"
	"fmt"
	"strings"
	"regexp"
)

func salesUpdate(dealer_id int, toDB map[string][]string, mapRow map[string]string, db *sql.DB) ( bool, error) {
	var validMiles = regexp.MustCompile("^[[:digit:]]+")
	
	toDB["dealer_id"]	 	= []string{fmt.Sprint(dealer_id), "integer"}
	toDB["dealnumber"]	= []string{mapRow["AppointmentDate"] + " " + mapRow["AppointmentTime"], "timestamp_without_time_zone"}
	toDB["clientdealerid"]  = []string{mapRow["ClientDealerID"], "character_varying(255)"}
	toDB["filetype"]  = []string{mapRow["FileType"], "character_varying(255)"}
	toDB["acdealerid"]  = []string{mapRow["ACDealerID"], "character_varying(255)"}
	toDB["dmstype"]  = []string{mapRow["DMSType"], "character_varying(255)"}
	toDB["customernumber"]  = []string{mapRow["CustomerNumber"], "character_varying(255)"}
	toDB["customername"]  = []string{mapRow["CustomerName"], "character_varying(255)"}
	toDB["customerfirstname"]  = []string{mapRow["CustomerFirstName"], "character_varying(255)"}
	toDB["customerlastname"]  = []string{mapRow["CustomerLastName"], "character_varying(255)"}
	toDB["customeraddress"]  = []string{mapRow["CustomerAddress"], "character_varying(255)"}
	toDB["customercity"]  = []string{mapRow["CustomerCity"], "character_varying(255)"}
	toDB["customerstate"]  = []string{mapRow["CustomerState"], "character_varying(255)"}
	toDB["customerzip"]  = []string{mapRow["CustomerZip"], "character_varying(255)"}
	toDB["customercounty"]  = []string{mapRow["CustomerCounty"], "character_varying(255)"}
	toDB["customercellphone"]  = []string{mapRow["CustomerCellPhone"], "character_varying(255)"}
	toDB["customerworkphone"]  = []string{mapRow["CustomerWorkPhone"], "character_varying(255)"}
	toDB["customerhomephone"]  = []string{mapRow["CustomerHomePhone"], "character_varying(255)"}
	toDB["customerpagerphone"]  = []string{mapRow["CustomerPagerPhone"], "character_varying(255)"}
	toDB["customerbirthdate"]  = []string{mapRow["CustomerBirthDate"], "timestamp_without_time_zone"}
	toDB["customeremail"]  = []string{mapRow["CustomerEmail"], "character_varying(255)"}
	toDB["mailblock"]  = []string{mapRow["MailBlock"], "character_varying(255)"}
	toDB["cobuyername"]  = []string{mapRow["CoBuyerName"], "character_varying(255)"}
	toDB["cobuyerfirstname"]  = []string{mapRow["CoBuyerFirstName"], "character_varying(255)"}
	toDB["cobuyerlastname"]  = []string{mapRow["CoBuyerLastName"], "character_varying(255)"}
	toDB["cobuyeraddress"]  = []string{mapRow["CoBuyerAddress"], "character_varying(255)"}
	toDB["cobuyercity"]  = []string{mapRow["CoBuyerCity"], "character_varying(255)"}
	toDB["cobuyerstate"]  = []string{mapRow["CoBuyerState"], "character_varying(255)"}
	toDB["cobuyerzip"]  = []string{mapRow["CoBuyerZip"], "character_varying(255)"}
	toDB["cobuyercounty"]  = []string{mapRow["CoBuyerCounty"], "character_varying(255)"}
	toDB["cobuyerhomephone"]  = []string{mapRow["CoBuyerHomePhone"], "character_varying(255)"}
	toDB["cobuyerworkphone"]  = []string{mapRow["CoBuyerWorkPhone"], "character_varying(255)"}
	toDB["cobuyerbirthdate"]  = []string{mapRow["CoBuyerBirthDate"], "timestamp_without_time_zone"}
	toDB["salesman_1_number"]  = []string{mapRow["Salesman_1_Number"], "character_varying(255)"}
	toDB["salesman_1_name"]  = []string{mapRow["Salesman_1_Name"], "character_varying(255)"}
	toDB["salesman_2_name"]  = []string{mapRow["Salesman_2_Name"], "character_varying(255)"}
	toDB["salesman_2_number"]  = []string{mapRow["Salesman_2_Number"], "character_varying(255)"}
	toDB["salesman_3_name"]  = []string{mapRow["Salesman_3_Name"], "character_varying(255)"}
	toDB["salesman_3_number"]  = []string{mapRow["Salesman_3_Number"], "character_varying(255)"}
	toDB["salesman_4_name"]  = []string{mapRow["Salesman_4_Name"], "character_varying(255)"}
	toDB["salesman_4_number"]  = []string{mapRow["Salesman_4_Number"], "character_varying(255)"}
	toDB["salesman_5_name"]  = []string{mapRow["Salesman_5_Name"], "character_varying(255)"}
	toDB["salesman_5_number"]  = []string{mapRow["Salesman_5_Number"], "character_varying(255)"}
	toDB["salesman_6_name"]  = []string{mapRow["Salesman_6_Name"], "character_varying(255)"}
	toDB["salesman_6_number"]  = []string{mapRow["Salesman_6_Number"], "character_varying(255)"}
	toDB["closingmanagername"]  = []string{mapRow["ClosingManagerName"], "character_varying(255)"}
	toDB["closingmanagernumber"]  = []string{mapRow["ClosingManagerNumber"], "character_varying(255)"}
	toDB["f_and_i_managername"]  = []string{mapRow["F_AND_I_ManagerName"], "character_varying(255)"}
	toDB["f_and_i_managernumber"]  = []string{mapRow["F_AND_I_ManagerNumber"], "character_varying(255)"}
	toDB["salesmanagername"]  = []string{mapRow["SalesManagerName"], "character_varying(255)"}
	toDB["salesmanagernumber"]  = []string{mapRow["SalesManagerNumber"], "character_varying(255)"}
	toDB["entrydate"]  = []string{mapRow["EntryDate"], "timestamp_without_time_zone"}
	toDB["dealbookdate"]  = []string{mapRow["DealBookDate"], "timestamp_without_time_zone"}
	toDB["vehicleclassification"]  = []string{mapRow["VehicleClassification"], "character_varying(255)"}
	toDB["vehicleexteriorcolor"]  = []string{mapRow["VehicleExteriorColor"], "character_varying(255)"}
	toDB["vehicleinteriorcolor"]  = []string{mapRow["VehicleInteriorColor"], "character_varying(255)"}
	toDB["vehiclemake"]  = []string{mapRow["VehicleMake"], "character_varying(255)"}
	toDB["vehiclemileage"]  = []string{mapRow["VehicleMileage"], "character_varying(255)"}
	toDB["vehiclemodel"]  = []string{mapRow["VehicleModel"], "character_varying(255)"}
	toDB["vehiclestocknumber"]  = []string{mapRow["VehicleStockNumber"], "character_varying(255)"}
	toDB["vehicletype"]  = []string{mapRow["VehicleType"], "character_varying(255)"}
	toDB["vehiclevin"]  = []string{mapRow["VehicleVIN"], "character_varying(255)"}
	toDB["vehicleyear"]  = []string{mapRow["VehicleYear"], "character_varying(255)"}
	toDB["inservicedate"]  = []string{mapRow["InServiceDate"], "timestamp_without_time_zone"}
	toDB["holdbackamount"]  = []string{mapRow["HoldBackAmount"], "numeric"}
	toDB["dealtype"]  = []string{mapRow["DealType"], "character_varying(255)"}
	toDB["saletype"]  = []string{mapRow["SaleType"], "character_varying(255)"}
	toDB["bankaddress"]  = []string{mapRow["BankAddress"], "character_varying(255)"}
	toDB["bankcity"]  = []string{mapRow["BankCity"], "character_varying(255)"}
	toDB["bankcode"]  = []string{mapRow["BankCode"], "character_varying(255)"}
	toDB["bankname"]  = []string{mapRow["BankName"], "character_varying(255)"}
	toDB["bankstate"]  = []string{mapRow["BankState"], "character_varying(255)"}
	toDB["bankzip"]  = []string{mapRow["BankZip"], "character_varying(255)"}
	toDB["aftreserve"]  = []string{mapRow["AFTReserve"], "numeric"}
	toDB["ahpremium"]  = []string{mapRow["AHPremium"], "numeric"}
	toDB["ahres"]  = []string{mapRow["AHRes"], "numeric"}
	toDB["aprrate"]  = []string{mapRow["APRRate"], "numeric"}
	toDB["aprrate2"]  = []string{mapRow["APRRate2"], "numeric"}
	toDB["aprrate3"]  = []string{mapRow["APRRate3"], "numeric"}
	toDB["aprrate4"]  = []string{mapRow["APRRate4"], "numeric"}
	toDB["accountingdate"]  = []string{mapRow["AccountingDate"], "timestamp_without_time_zone"}
	toDB["amountfinanced"]  = []string{mapRow["AmountFinanced"], "numeric"}
	toDB["backgross"]  = []string{mapRow["BackGross"], "numeric"}
	toDB["balloonamount"]  = []string{mapRow["BalloonAmount"], "numeric"}
	toDB["basepayment"]  = []string{mapRow["BasePayment"], "numeric"}
	toDB["bodydescription"]  = []string{mapRow["BodyDescription"], "character_varying(255)"}
	toDB["bodydoorcount"]  = []string{mapRow["BodyDoorCount"], "character_varying(255)"}
	toDB["buyrate"]  = []string{mapRow["BuyRate"], "numeric"}
	toDB["cass_std_cart"]  = []string{mapRow["CASS_STD_CART"], "character_varying(255)"}
	toDB["cass_std_chkdgt"]  = []string{mapRow["CASS_STD_CHKDGT"], "character_varying(255)"}
	toDB["cass_std_city"]  = []string{mapRow["CASS_STD_CITY"], "character_varying(255)"}
	toDB["cass_std_dpbc"]  = []string{mapRow["CASS_STD_DPBC"], "character_varying(255)"}
	toDB["cass_std_error_cd"]  = []string{mapRow["CASS_STD_ERROR_CD"], "character_varying(255)"}
	toDB["cass_std_ews"]  = []string{mapRow["CASS_STD_EWS"], "character_varying(255)"}
	toDB["cass_std_fips"]  = []string{mapRow["CASS_STD_FIPS"], "character_varying(255)"}
	toDB["cass_std_lacs"]  = []string{mapRow["CASS_STD_LACS"], "character_varying(255)"}
	toDB["cass_std_lacsrt"]  = []string{mapRow["CASS_STD_LACSRT"], "character_varying(255)"}
	toDB["cass_std_line1"]  = []string{mapRow["CASS_STD_LINE1"], "character_varying(255)"}
	toDB["cass_std_line2"]  = []string{mapRow["CASS_STD_LINE2"], "character_varying(255)"}
	toDB["cass_std_lot"]  = []string{mapRow["CASS_STD_LOT"], "character_varying(255)"}
	toDB["cass_std_lotord"]  = []string{mapRow["CASS_STD_LOTORD"], "character_varying(255)"}
	toDB["cass_std_ndiapt"]  = []string{mapRow["CASS_STD_NDIAPT"], "character_varying(255)"}
	toDB["cass_std_ndirr"]  = []string{mapRow["CASS_STD_NDIRR"], "character_varying(255)"}
	toDB["cass_std_state"]  = []string{mapRow["CASS_STD_STATE"], "character_varying(255)"}
	toDB["cass_std_urb"]  = []string{mapRow["CASS_STD_URB"], "character_varying(255)"}
	toDB["cass_std_z4lom"]  = []string{mapRow["CASS_STD_Z4LOM"], "character_varying(255)"}
	toDB["cass_std_zip"]  = []string{mapRow["CASS_STD_ZIP"], "character_varying(255)"}
	toDB["cass_std_zip4"]  = []string{mapRow["CASS_STD_ZIP4"], "character_varying(255)"}
	toDB["cass_std_zipmov"]  = []string{mapRow["CASS_STD_ZIPMOV"], "character_varying(255)"}
	toDB["cashdeposit"]  = []string{mapRow["CashDeposit"], "numeric"}
	toDB["cashprice"]  = []string{mapRow["CashPrice"], "numeric"}
	toDB["cobuyeraddress2"]  = []string{mapRow["CoBuyerAddress2"], "character_varying(255)"}
	toDB["cobuyercell"]  = []string{mapRow["CoBuyerCell"], "character_varying(255)"}
	toDB["cobuyercustnum"]  = []string{mapRow["CoBuyerCustNum"], "character_varying(255)"}
	toDB["cobuyeremail"]  = []string{mapRow["CoBuyerEmail"], "character_varying(255)"}
	toDB["cobuyeremailblock"]  = []string{mapRow["CoBuyerEmailBlock"], "character_varying(255)"}
	toDB["cobuyeremployer"]  = []string{mapRow["CoBuyerEmployer"], "character_varying(255)"}
	toDB["cobuyermailblock"]  = []string{mapRow["CoBuyerMailBlock"], "character_varying(255)"}
	toDB["cobuyermiddlename"]  = []string{mapRow["CoBuyerMiddleName"], "character_varying(255)"}
	toDB["cobuyeroccupation"]  = []string{mapRow["CoBuyerOccupation"], "character_varying(255)"}
	toDB["cobuyeroptout"]  = []string{mapRow["CoBuyerOptOut"], "character_varying(255)"}
	toDB["cobuyerphoneblock"]  = []string{mapRow["CoBuyerPhoneBlock"], "character_varying(255)"}
	toDB["cobuyersalutation"]  = []string{mapRow["CoBuyerSalutation"], "character_varying(255)"}
	toDB["contractdate"]  = []string{mapRow["ContractDate"], "timestamp_without_time_zone"}
	toDB["cost"]  = []string{mapRow["Cost"], "numeric"}
	toDB["country"]  = []string{mapRow["Country"], "character_varying(255)"}
	toDB["creditlifecommision"]  = []string{mapRow["CreditLifeCommision"], "numeric"}
	toDB["creditlifeprem"]  = []string{mapRow["CreditLifePrem"], "numeric"}
	toDB["creditlifepremium"]  = []string{mapRow["CreditLifePremium"], "numeric"}
	toDB["creditliferes"]  = []string{mapRow["CreditLifeRes"], "numeric"}
	toDB["customeraddress2"]  = []string{mapRow["CustomerAddress2"], "character_varying(255)"}
	toDB["customermiddlename"]  = []string{mapRow["CustomerMiddleName"], "character_varying(255)"}
	toDB["customersalutation"]  = []string{mapRow["CustomerSalutation"], "character_varying(255)"}
	toDB["customersuffix"]  = []string{mapRow["CustomerSuffix"], "character_varying(255)"}
	toDB["dmvamount"]  = []string{mapRow["DMVAmount"], "numeric"}
	toDB["daytofirstpayment"]  = []string{mapRow["DayToFirstPayment"], "character_varying(255)"}
	toDB["dealdateoffset"]  = []string{mapRow["DealDateOffset"], "character_varying(255)"}
	toDB["dealnumber"]  = []string{mapRow["DealNumber"], "character_varying(255)"}
	toDB["dealstatus"]  = []string{mapRow["DealStatus"], "character_varying(255)"}
	toDB["dealerselect"]  = []string{mapRow["DealerSelect"], "character_varying(255)"}
	toDB["downpayment"]  = []string{mapRow["DownPayment"], "numeric"}
	toDB["downpayment2"]  = []string{mapRow["DownPayment2"], "numeric"}
	toDB["emailblock"]  = []string{mapRow["EmailBlock"], "character_varying(255)"}
	toDB["employer"]  = []string{mapRow["Employer"], "character_varying(255)"}
	toDB["enginedesc"]  = []string{mapRow["EngineDesc"], "character_varying(255)"}
	toDB["extendedwarrantydollar"]  = []string{mapRow["ExtendedWarrantyDollar"], "numeric"}
	toDB["extendedwarrantyflag"]  = []string{mapRow["ExtendedWarrantyFlag"], "character_varying(255)"}
	toDB["extendedwarrantylimitmiles"]  = []string{mapRow["ExtendedWarrantyLimitMiles"], "numeric"}
	toDB["extendedwarrantyname"]  = []string{mapRow["ExtendedWarrantyName"], "character_varying(255)"}
	toDB["extendedwarrantyprofit"]  = []string{mapRow["ExtendedWarrantyProfit"], "numeric"}
	toDB["extendedwarrantyterm"]  = []string{mapRow["ExtendedWarrantyTerm"], "character_varying(255)"}
	toDB["fee_10_commission"]  = []string{mapRow["Fee_10_Commission"], "numeric"}
	toDB["fee_10_fee"]  = []string{mapRow["Fee_10_Fee"], "numeric"}
	toDB["fee_10_name"]  = []string{mapRow["Fee_10_Name"], "character_varying(255)"}
	toDB["fee_1_commission"]  = []string{mapRow["Fee_1_Commission"], "numeric"}
	toDB["fee_1_fee"]  = []string{mapRow["Fee_1_Fee"], "numeric"}
	toDB["fee_1_name"]  = []string{mapRow["Fee_1_Name"], "character_varying(255)"}
	toDB["fee_2_commission"]  = []string{mapRow["Fee_2_Commission"], "numeric"}
	toDB["fee_2_fee"]  = []string{mapRow["Fee_2_Fee"], "numeric"}
	toDB["fee_2_name"]  = []string{mapRow["Fee_2_Name"], "character_varying(255)"}
	toDB["fee_3_commission"]  = []string{mapRow["Fee_3_Commission"], "numeric"}
	toDB["fee_3_fee"]  = []string{mapRow["Fee_3_Fee"], "numeric"}
	toDB["fee_3_name"]  = []string{mapRow["Fee_3_Name"], "character_varying(255)"}
	toDB["fee_4_commission"]  = []string{mapRow["Fee_4_Commission"], "numeric"}
	toDB["fee_4_fee"]  = []string{mapRow["Fee_4_Fee"], "numeric"}
	toDB["fee_4_name"]  = []string{mapRow["Fee_4_Name"], "character_varying(255)"}
	toDB["fee_5_commission"]  = []string{mapRow["Fee_5_Commission"], "numeric"}
	toDB["fee_5_fee"]  = []string{mapRow["Fee_5_Fee"], "numeric"}
	toDB["fee_5_name"]  = []string{mapRow["Fee_5_Name"], "character_varying(255)"}
	toDB["fee_6_commission"]  = []string{mapRow["Fee_6_Commission"], "numeric"}
	toDB["fee_6_fee"]  = []string{mapRow["Fee_6_Fee"], "numeric"}
	toDB["fee_6_name"]  = []string{mapRow["Fee_6_Name"], "character_varying(255)"}
	toDB["fee_7_commission"]  = []string{mapRow["Fee_7_Commission"], "numeric"}
	toDB["fee_7_fee"]  = []string{mapRow["Fee_7_Fee"], "numeric"}
	toDB["fee_7_name"]  = []string{mapRow["Fee_7_Name"], "character_varying(255)"}
	toDB["fee_8_commission"]  = []string{mapRow["Fee_8_Commission"], "numeric"}
	toDB["fee_8_fee"]  = []string{mapRow["Fee_8_Fee"], "numeric"}
	toDB["fee_8_name"]  = []string{mapRow["Fee_8_Name"], "character_varying(255)"}
	toDB["fee_9_commission"]  = []string{mapRow["Fee_9_Commission"], "numeric"}
	toDB["fee_9_fee"]  = []string{mapRow["Fee_9_Fee"], "numeric"}
	toDB["fee_9_name"]  = []string{mapRow["Fee_9_Name"], "character_varying(255)"}
	toDB["financecharge"]  = []string{mapRow["FinanceCharge"], "numeric"}
	toDB["financereserve"]  = []string{mapRow["FinanceReserve"], "numeric"}
	toDB["frontgross"]  = []string{mapRow["FrontGross"], "numeric"}
	toDB["gappremium"]  = []string{mapRow["GAPPremium"], "numeric"}
	toDB["globaloptout"]  = []string{mapRow["GlobalOptOut"], "character_varying(255)"}
	toDB["grossprofitsale"]  = []string{mapRow["GrossProfitSale"], "numeric"}
	toDB["incentives"]  = []string{mapRow["Incentives"], "character_varying(255)"}
	toDB["individualbusinessflag"]  = []string{mapRow["IndividualBusinessFlag"], "character_varying(255)"}
	toDB["insuranceaddress"]  = []string{mapRow["InsuranceAddress"], "character_varying(255)"}
	toDB["insuranceagentname"]  = []string{mapRow["InsuranceAgentName"], "character_varying(255)"}
	toDB["insurancecity"]  = []string{mapRow["InsuranceCity"], "character_varying(255)"}
	toDB["insurancecompensationdeduction"]  = []string{mapRow["InsuranceCompensationDeduction"], "character_varying(255)"}
	toDB["insuranceeffectivedate"]  = []string{mapRow["InsuranceEffectiveDate"], "timestamp_without_time_zone"}
	toDB["insuranceexpirationdate"]  = []string{mapRow["InsuranceExpirationDate"], "timestamp_without_time_zone"}
	toDB["insurancename"]  = []string{mapRow["InsuranceName"], "character_varying(255)"}
	toDB["insurancephone"]  = []string{mapRow["InsurancePhone"], "character_varying(255)"}
	toDB["insurancepolicynumber"]  = []string{mapRow["InsurancePolicyNumber"], "character_varying(255)"}
	toDB["insurancestate"]  = []string{mapRow["InsuranceState"], "character_varying(255)"}
	toDB["insurancezip"]  = []string{mapRow["InsuranceZip"], "character_varying(255)"}
	toDB["inventorydate"]  = []string{mapRow["InventoryDate"], "timestamp_without_time_zone"}
	toDB["invoiceamount"]  = []string{mapRow["InvoiceAmount"], "numeric"}
	toDB["journalsaleamount"]  = []string{mapRow["JournalSaleAmount"], "numeric"}
	toDB["language"]  = []string{mapRow["Language"], "character_varying(255)"}
	toDB["lastinstallmentdate"]  = []string{mapRow["LastInstallmentDate"], "timestamp_without_time_zone"}
	toDB["leaseannualmiles"]  = []string{mapRow["LeaseAnnualMiles"], "numeric"}
	toDB["leaseestimatedmiles"]  = []string{validMiles.FindString(mapRow["LeaseEstimatedMiles"]), "numeric"}
	toDB["leasefirstpaydate"]  = []string{mapRow["LeaseFirstPayDate"], "timestamp_without_time_zone"}
	toDB["leasenetcapcost"]  = []string{mapRow["LeaseNetCapCost"], "numeric"}
	toDB["leasepayment"]  = []string{mapRow["LeasePayment"], "numeric"}
	toDB["leaserate"]  = []string{mapRow["LeaseRate"], "numeric"}
	toDB["leaseterm"]  = []string{mapRow["LeaseTerm"], "numeric"}
	toDB["leasetotalcapreduction"]  = []string{mapRow["LeaseTotalCapReduction"], "numeric"}
	toDB["licensefee"]  = []string{mapRow["LicenseFee"], "numeric"}
	toDB["licenseplatenumber"]  = []string{mapRow["LicensePlateNumber"], "character_varying(255)"}
	toDB["listprice"]  = []string{mapRow["ListPrice"], "numeric"}
	toDB["msrp"]  = []string{mapRow["MSRP"], "numeric"}
	toDB["mileagerate"]  = []string{mapRow["MileageRate"], "numeric"}
	toDB["modeldescriptionofcarsold"]  = []string{mapRow["ModelDescriptionOfCarSold"], "character_varying(255)"}
	toDB["modelnumberofcarsold"]  = []string{mapRow["ModelNumberOfCarSold"], "character_varying(255)"}
	toDB["ncoa_ac_id"]  = []string{mapRow["NCOA_AC_ID"], "character_varying(255)"}
	toDB["nettradeamount"]  = []string{mapRow["NetTradeAmount"], "numeric"}
	toDB["occupation"]  = []string{mapRow["Occupation"], "character_varying(255)"}
	toDB["paymenttype"]  = []string{mapRow["PaymentType"], "character_varying(255)"}
	toDB["phoneblock"]  = []string{mapRow["PhoneBlock"], "character_varying(255)"}
	toDB["rosnumber"]  = []string{mapRow["ROSNumber"], "character_varying(255)"}
	toDB["realbookdate"]  = []string{mapRow["RealBookDate"], "timestamp_without_time_zone"}
	toDB["rebate"]  = []string{mapRow["Rebate"], "numeric"}
	toDB["registrationfee"]  = []string{mapRow["RegistrationFee"], "numeric"}
	toDB["residualamount"]  = []string{mapRow["ResidualAmount"], "numeric"}
	toDB["retailfirstpaydate"]  = []string{mapRow["RetailFirstPayDate"], "timestamp_without_time_zone"}
	toDB["retailpayment"]  = []string{mapRow["RetailPayment"], "numeric"}
	toDB["slct2"]  = []string{mapRow["SLCT2"], "character_varying(255)"}
	toDB["salesman2commission"]  = []string{mapRow["Salesman2Commission"], "numeric"}
	toDB["salesmancommission"]  = []string{mapRow["SalesmanCommission"], "numeric"}
	toDB["securitydeposit2"]  = []string{mapRow["SecurityDeposit2"], "numeric"}
	toDB["securitydesposit"]  = []string{mapRow["SecurityDesposit"], "numeric"}
	toDB["statedmvtotfee"]  = []string{mapRow["StateDMVTotFee"], "character_varying(255)"}
	toDB["statusdate"]  = []string{mapRow["StatusDate"], "timestamp_without_time_zone"}
	toDB["subtrimlevel"]  = []string{mapRow["SubTrimLevel"], "character_varying(255)"}
	toDB["term"]  = []string{mapRow["Term"], "numeric"}
	toDB["term2"]  = []string{mapRow["Term2"], "numeric"}
	toDB["totalaccessories"]  = []string{mapRow["TotalAccessories"], "numeric"}
	toDB["totaldriveoffamount"]  = []string{mapRow["TotalDriveOffAmount"], "numeric"}
	toDB["totalinsurancereserve"]  = []string{mapRow["TotalInsuranceReserve"], "numeric"}
	toDB["totalofpayments"]  = []string{mapRow["TotalOfPayments"], "numeric"}
	toDB["totalofpayments2"]  = []string{mapRow["TotalOfPayments2"], "numeric"}
	toDB["totalpickuppayment"]  = []string{mapRow["TotalPickupPayment"], "numeric"}
	toDB["totaltax"]  = []string{mapRow["TotalTax"], "numeric"}
	toDB["tradein_1_acv"]  = []string{mapRow["TradeIn_1_ACV"], "numeric"}
	toDB["tradein_1_exteriorcolor"]  = []string{mapRow["TradeIn_1_ExteriorColor"], "character_varying(255)"}
	toDB["tradein_1_gross"]  = []string{mapRow["TradeIn_1_Gross"], "numeric"}
	toDB["tradein_1_interiorcolor"]  = []string{mapRow["TradeIn_1_InteriorColor"], "character_varying(255)"}
	toDB["tradein_1_make"]  = []string{mapRow["TradeIn_1_Make"], "character_varying(255)"}
	toDB["tradein_1_mileage"]  = []string{mapRow["TradeIn_1_Mileage"], "character_varying(255)"}
	toDB["tradein_1_model"]  = []string{mapRow["TradeIn_1_Model"], "character_varying(255)"}
	toDB["tradein_1_payoff"]  = []string{mapRow["TradeIn_1_Payoff"], "numeric"}
	toDB["tradein_1_vin"]  = []string{mapRow["TradeIn_1_VIN"], "character_varying(255)"}
	toDB["tradein_1_year"]  = []string{mapRow["TradeIn_1_Year"], "character_varying(255)"}
	toDB["tradein_2_acv"]  = []string{mapRow["TradeIn_2_ACV"], "numeric"}
	toDB["tradein_2_exteriorcolor"]  = []string{mapRow["TradeIn_2_ExteriorColor"], "character_varying(255)"}
	toDB["tradein_2_gross"]  = []string{mapRow["TradeIn_2_Gross"], "numeric"}
	toDB["tradein_2_interiorcolor"]  = []string{mapRow["TradeIn_2_InteriorColor"], "character_varying(255)"}
	toDB["tradein_2_make"]  = []string{mapRow["TradeIn_2_Make"], "character_varying(255)"}
	toDB["tradein_2_mileage"]  = []string{mapRow["TradeIn_2_Mileage"], "character_varying(255)"}
	toDB["tradein_2_model"]  = []string{mapRow["TradeIn_2_Model"], "character_varying(255)"}
	toDB["tradein_2_payoff"]  = []string{mapRow["TradeIn_2_Payoff"], "numeric"}
	toDB["tradein_2_vin"]  = []string{mapRow["TradeIn_2_VIN"], "character_varying(255)"}
	toDB["tradein_2_year"]  = []string{mapRow["TradeIn_2_Year"], "character_varying(255)"}
	toDB["transmissiondesc"]  = []string{mapRow["TransmissionDesc"], "character_varying(255)"}
	toDB["trimlevel"]  = []string{mapRow["TrimLevel"], "character_varying(255)"}
	toDB["typecode"]  = []string{mapRow["TypeCode"], "character_varying(255)"}
	toDB["weight"]  = []string{mapRow["Weight"], "character_varying(255)"}
	toDB["cobuyercountry"]  = []string{mapRow["CoBuyerCountry"], "character_varying(255)"}
	
	delete(toDB, "last_name")
	delete(toDB, "first_name")
	
	cond := "dealer_id=" + fmt.Sprint(dealer_id)
	if strings.TrimSpace(toDB["dealnumber"][0]) != "" {
		cond += " and dealnumber=" + normalizeValue(toDB["dealnumber"][0], toDB["dealnumber"][1])
	}
	
	return procDB("sales", cond, toDB, db)
}