VAS\GetDetail
=============

Get detailed information

Usage
-----

Get detailed information of an item from an existing VAS \(Value Added Services\) account or subscription.

Availability
------------

All resellers have access to this command.

Contraints
----------

The query must meet the following requirements:
- The login ID and password must be valid.
- The VAS account or subscription must belong to this account.

Products
--------
- [Google Apps](#vas-get-detail-googleapps)
- [Whois Publicity Service](#vas-get-detail-whoispublicityservice)

Google Apps / Google Suite
--------------------------

#### Input Parameters

| Parameter  | Type | Status  | Description |
| ------------ | ------- | -------- | ------------------------------------------------------------------------------- |
| Command   | string | Required | VAS\_GetDetail. |
| UID | string | Required | Your Account ID.                                |
| PW      | string | Required | Your API Token. |
| SLD | string | Required | Second-level domain name \(e.g. "enom" in "enom.com"\).          |
| TLD     | string | Required | Top-level domain name \(extension\). |
| ProductType | string | Required | Type of VAS product. Permitted values: - GoogleApps              |
| IncludePrice | boolean | Optional | Return additional information about current, pro-rate and next billing items. |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.   |

#### Returned Parameters and Values

> ### Check the return parameter ErrCount. If greater than 0 the transaction failed. The parameter Err\(ErrCount\) can be presented to the client. Otherwise, the process returns parameters as defined below.

| Parameter               | Type | Description                                                                          |
| -------------------------------------- | ------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command                | string | Name of command executed.                                                                   |
| ErrorCount               | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                |
| ErrX                  | string | Error messages explaining the failure. These can be presented as is back to the client.                                   |
| Done                  | string | True value indicates this entire response has reached you successfully.                                            |
| ProductType              | string | Name of the product. Expected values: - GoogleApps                                                      |
| InternalDomain             | boolean | Is domain registered and located in the customer's account? This flag is useful to determine where to update domain name's host record for site verification. |
| CustomerCreated            | boolean | Is **Customer**created successfully during account provisioning?                                                |
| AdminUserCreated            | boolean | Is **Admin User** created successfully during account provisioning?                                              |
| SubscriptionCreated          | boolean | Is **Subscription** created successfully during account provisioning?                                             |
| *`<Customer>`* ID           | integer | Customer's **ID** number associated with the value added service.                                               |
| *`<Customer>`* DomainName       | string | Customer's **Domain Name** associated with the value added service.                                              |
| *`<Customer>`* DomainVerified     | string | Customer's **Domain Verification** status associated with the value added service. Expected values: - True - False                      |
| *`<Customer>`* AlternateEmail     | string | Customer's **Email Address** associated with the value added service.                                             |
| *`<Customer>`* Name          | string | Customer's **Full Name** associated with the value added service subscription.                                        |
| *`<Customer>`* FirstName        | string | Customer's **First Name** associated with the value added service subscription.                                        |
| *`<Customer>`* LastName        | string | Customer's **Last Name** associated with the value added service subscription.                                        |
| *`<Customer>`* OrganizationName    | string | Customer's **Organization Name** associated with the value added service subscription.                                    |
| *`<Customer>`* Address         | string | Customer's **Address** associated with the value added service subscription.                                         |
| *`<Customer>`* City          | string | Customer's **City** associated with the value added service subscription.                                           |
| *`<Customer>`* StateProvince      | string | Customer's **State** or **Province** associated with the value added service subscription.                                  |
| *`<Customer>`* PostalCode       | string | Customer's **Postal Code** associated with the value added service subscription.                                       |
| *`<Customer>`* CountryCode       | string | Customer's two-character **Country Codes** associated with the value added service subscription.                               |
| *`<Customer>`* PhoneNumber       | string | Customer's **Phone Number** associated with the value added service subscription.                                       |
| *`<Subscription>`* ID         | integer | Subscription **ID** number for the value added service subscription.                                             |
| *`<Subscription>`* Name        | string | Subscription **Name** associated with the subscription of the value added service.                                      |
| *`<Subscription>`* Type        | string | Subscription **Type** associated with the subscription of the value added service.                                      |
| *`<Subscription>`* IsInTrial      | string | Subscription confirmation if the value added service is within the **Trial Grace Period** or not. Expected values: - True - False               |
| *`<Subscription>`* Status       | string | Subscription current **Status** of the value added service. Expected values: - ACTIVE - SUSPENDED                               |
| *`<Subscription>`* SuspensionReasons  | object | Collection of subscription **Suspension Reason** of the value added service.                                         |
| *`<Subscription>`* QtyPurchased    | integer | Number of unit purchased                                                                    |
| *`<Subscription>`* QtyUsed       | integer | Number of unit used                                                                      |
| *`<Subscription>`* QtyAvailable    | integer | Number of unit available                                                                    |
| *`<Subscription>`* CreationTime    | string | Subscription Date and Time that the value added service was **created**. Format: *MM/DD/YYYY HH:MM:SS AM/PM*                         |
| *`<Subscription>`* TrialEndTime    | string | Subscription Date and Time that the value added service **exits** the Triage Grace Period. Format: *MM/DD/YYYY HH:MM:SS AM/PM*                |
| *`<Subscription>`* ResourceURL     | string | Subscription URL to access Google Apps' **Admin Console**.                                                  |
| *`<Subscription>`* VASType       | integer | Current Value Added Services type ID                                                              |
| *`<Subscription>`* VASDesc       | string | Current Value Added Services description                                                            |
| *`<Subscription>`* VASPrice      | decimal | Current Value Added Services price                                                               |
| *`<Subscription>`* VASProrateType   | integer | Pro-rated \(non-freetrial\) Value Added Services type ID                                                   |
| *`<Subscription>`* VASProrateDesc   | string | Pro-rated \(non-freetrial\) Value Added Services description                                                 |
| *`<Subscription>`* VASProratePrice   | decimal | Pro-rated \(non-freetrial\) Value Added Services price                                                    |
| *`<Subscription>`* VASNextBillingType | integer | Next billing period Value Added Services type ID                                                        |
| *`<Subscription>`* VASNextBillingDesc | string | Next billing period Value Added Services description                                                      |
| *`<Subscription>`* VASNextBillingPrice | decimal | Next billing period Value Added Services price                                                         |
| *`<Subscription>`* VASNextBillingDate | string | Next billing period Value Added Services date Format: *MM/DD/YYYY HH:MM:SS AM/PM*                                       |
| *`<User>`* ID             | string | User's **ID** number for the value added service subscription.                                                |
| *`<User>`* IsAdmin           | boolean | User's**Admin** status for the value added service subscription.                                               |
| *`<User>`* Suspended          | string | User's **Suspended** status for the value added service subscription.                                             |
| *`<User>`* AgreedToTerms        | boolean | User's **Agreed To Terms** status for the value added service subscription.                                          |
| *`<User>`* PrimaryEmail        | string | User's **Primary Email** for the value added service subscription.                                              |
| *`<User>`* FullName          | string | User's **Full Name** for the value added service subscription.                                                |
| *`<User>`* FirstName          | string | User's **First Name** for the value added service subscription.                                                |
| *`<User>`* LastName          | string | User's **Last Name** for the value added service subscription.                                                |
| *`<User>`* OrgUnitPath         | string | User's **Organization Unit Path** for the value added service subscription.                                          |
| *`<User>`* CreationTime        | string | User's **Creation Time** for the value added service subscription.                                              |
| *`<User>`* LastLoginTime        | string | User's **Last Login Time** for the value added service subscription.                                             |
| *`<User>`* ChangePasswordAtNextLogin  | boolean | User's **Change Password At Next Login** enforcement status for the value added service subscription.                             |
| *`<User>`* IncludeInGlobalAddressList | boolean | User's **Include In Global Address List** for the value added service subscription.                                      |
| *`<User>`* IPWhitelisted        | boolean | User's **IP Whitelisted** flag for the value added service subscription.                                           |
| *`<User>`* IsDelegatedAdmin      | boolean | User's **Delegated Admin** flag for the value added service subscription.                                           |
| *`<User>`* IsMailboxSetup       | boolean | User's **Mailbox Setup** flag for the value added service subscription.                                            |
| *`<User>`* - `<Email>`\*Address    | string | User's **Email Address** for the value added service subscription.                                              |
| *`<User>`* - `<Email>`\*Primary    | string | User's **Primary Email Address** flag for the value added service subscription.                                        |
| *`<User>`* LoginUrl          | string | User's URL to access Google Apps' **User Console**.                                                      |

#### Examples

```
https://resellertest.enom.com/interface.asp?command=VAS_GetDetail&uid=YourAccountID&pw=YourApiToken&SLD={Required}&TLD={Required}&ProductType=GoogleApps&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
<VAS_GetDetail>
  <ProductType>GoogleApps</ProductType>
  <InternalDomain>false</InternalDomain>
  <CustomerCreated>true</CustomerCreated>
  <AdminUserCreated>true</AdminUserCreated>
  <SubscriptionCreated>true</SubscriptionCreated>
  <Customer>
   <ID>C02pxtg2f</ID>
   <DomainName><![CDATA[buditest1030.net]]></DomainName>
   <DomainVerified>true</DomainVerified>
   <AlternateEmail><![CDATA[[email protected]]]></AlternateEmail>
   <Name><![CDATA[first last]]></Name>
   <FirstName><![CDATA[first]]></FirstName>
   <LastName><![CDATA[last]]></LastName>
   <OrganizationName><![CDATA[Daily Planet]]></OrganizationName>
   <Address />
   <City />
   <StateProvince />
   <PostalCode>98107</PostalCode>
   <CountryCode>US</CountryCode>
   <PhoneNumber>+1 425-555-5555</PhoneNumber>
   <Subscriptions>
     <Subscription>
      <ID>66991935</ID>
      <Name>Google-Apps-For-Business</Name>
      <Type>FLEXIBLE</Type>
      <IsInTrial>false</IsInTrial>
      <Status>SUSPENDED</Status>
      <SuspensionReasons>
        <Reason>RESELLER_INITIATED</Reason>
      </SuspensionReasons>
						<QtyPurchased>8</QtyPurchased>
						<QtyUsed>3</QtyUsed>
						<QtyAvailable>5</QtyAvailable>
      <CreationTime UTC="2016-07-11T17:01:03.000Z" Epoch="1468256463">
        7/11/2016 10:01:03 AM
      </CreationTime>
      <TrialEndTime UTC="" Epoch=""></TrialEndTime>
      <ResourceUrl><![CDATA[https://www.google.com/a/cpanel/
        goog-test.reseller.enom.com.buditest1030.net.gaprodtesting.rocks/
        AdminHome#DomainSettings/notab=1&subtab=subscriptions]]>
      <VASType>2410</VASType>
      <VASDesc>Google Apps for Work - Monthly</VASDesc>
      <VASPrice>4.50</VASPrice>
      <VASProrateType/>
      <VASProrateDesc/>
      <VASProratePrice/>
      <VASNextBillingType>2410</VASNextBillingType>
      <VASNextBillingDesc>Google Apps for Work - Monthly</VASNextBillingDesc>
      <VASNextBillingPrice>4.50</VASNextBillingPrice>
      </ResourceUrl>
     </Subscription>
   </Subscriptions>
   <Users>
     <Count>1</Count>
     <User>
      <ID>117072922639416417499</ID>
      <IsAdmin>true</IsAdmin>
      <Suspended>false</Suspended>
      <AgreedToTerms>true</AgreedToTerms>
      <PrimaryEmail><![CDATA[[email protected]]]></PrimaryEmail>
      <FullName><![CDATA[bp tester]]></FullName>
      <FirstName><![CDATA[bp]]></FirstName>
      <LastName><![CDATA[tester]]></LastName>
      <OrgUnitPath>/</OrgUnitPath>
      <CreationTime UTC="2016-07-22T17:08:59.000Z" Epoch="1469182139">
      	  7/22/2016 3:08:59 AM
      </CreationTime>
      <LastLoginTime UTC="2016-07-22T17:09:24.000Z" Epoch="1469182164">
        7/22/2016 3:09:24 AM
      </LastLoginTime>
      <ChangePasswordAtNextLogin>false</ChangePasswordAtNextLogin>
      <IncludeInGlobalAddressList>true</IncludeInGlobalAddressList>
      <IPWhitelisted>false</IPWhitelisted>
      <IsDelegatedAdmin>false</IsDelegatedAdmin>
      <IsMailboxSetup>false</IsMailboxSetup>
      <Emails>
        <Email>
         <Address><![CDATA[[email protected]]]></Address>
         <Primary>true</Primary>
        </Email>
      </Emails>
      <LoginUrl><![CDATA[https://accounts.google.com/AccountChooser
        [email protected]&service=CPanel&continue=
        https://admin.google.com/buditest1030.net/AcceptTermsOfService
        ?continue=https://mail.google.com/a/buditest1030.net]]>
      </LoginUrl>
     </User>
   </Users>
  </Customer>
</VAS_GetDetail>
<Command>VAS_GETDETAIL</Command>
<APIType>API.NET</APIType>
<Language>eng</Language>
<ErrCount>0</ErrCount>
<ResponseCount>0</ResponseCount>
<MinPeriod>1</MinPeriod>
<MaxPeriod>10</MaxPeriod>
<Server>krkdt198</Server>
<Site>eNom</Site>
<IsLockable>True</IsLockable>
<IsRealTimeTLD>True</IsRealTimeTLD>
<TimeDifference>+8.00</TimeDifference>
<ExecTime>2.805</ExecTime>
<Done>true</Done>
<TrackingKey>cb9c39cc-fe0e-4361-82a8-3d6c7a8ae23c</TrackingKey>
<RequestDateTime>8/24/2016 10:22:06 AM</RequestDateTime>
```
**[Back to Top](#top)**

Whois Publicity Service
-----------------------

**Product life cycle**

| ProdStatusID | ProdStatusDesc |
| ------------ | ---------------------- |
| 1      | Awaiting Configuration |
| 2 | Service Active     |
| 3      | Billing Failed |
| 4 | Pending Renewal    |
| 5      | Cancellation Pending |
| 6 | Service Canceled    |
| 7      | Pending Expiration |
| 8 | Service Expired    |
| 9      | Service Deleted |

**Private data visibility logic**

| ProdStatusID | ProdConsented | ProdEnabled | WHOIS display | Example scenarios |
| ------------ | ------------- | ----------- | ------------- | ------------------------------------------------------------------------------------------------------- |
| 2 | False     | False | Public    | Either the product is still in purchasing process or all flags are turned off by user. |
| 2 | False     | True | Public    | User just purchased the product. By default the switch is on, but has not consented. |
| 2 | True     | False | Public    | User switches off the product. |
| 2 | True     | True | **PRIVATE**  | Consented and enable flag is true. |
| Other than 2 | N/A      | N/A | Public    | Domain might have the product, it is either still in processing, deleted, expired or other conditions. |

**Contact data samples**

| Type |                                                              |
| ------- | -------------------------------------------------------------------------------------------------------------------------- |
| Public | `<FirstName>Data Protected</FirstName>\ <LastName>Data Protected</LastName>\ <Address1>123 Data Protected</Address1>` .. |
| Private | `<FirstName>John</FirstName>\ <LastName>Smith</LastName>\ <Address1>2710 Clark Avenue</Address1>` ... |

| Parameter | Type  | Status | Description                                                   |
| ------------ | ------- | -------- | --------------------------------------------------------------------------------------------------------------- |
| Command | string | Required | VAS\_GetDetail                                                 |
| UID     | string | Required | Your Account ID. |
| PW | string | Required | Your API Token.                                                |
| ProductType | string | Required | Type of VAS product. Permitted values: - WhoisPublicity |
| VASItemID | integer | Required | ID number of the VAS Item. Use the [VAS\_GetList](../docs/vas-get-list.md) command to retrieve the ID number |
|       | |     | |
| ResponseType | string | Optional | Format of response. Permitted values are: Text \(default\), HTML, or XML.                   |

| Output     | Type | Description                                                                                                                 |
| -------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Command    | string | Name of command executed.                                                                                                          |
| ErrorCount   | integer | The number of errors if any occurred. If greater than 0 check the Err\(1 to ErrCount\) values.                                                                       |
| ErrX      | string | Error messages explaining the failure. These can be presented as is back to the client.                                                                          |
| ProductType  | string | Name of the product. Expected values: - WhoisPublicity                                                                                           |
| VASItemID   | integer | ID number of the VAS Item.                                                                                                         |
| DomainNameID  | integer | Domain name ID.                                                                                                               |
| DomainName   | string | Domain name.                                                                                                                |
| AutoRenew   | boolean | Auto renewal flag. Expected values: - True - False                                                                                             |
| ProdType    | integer | Product ID.                                                                                                                 |
| ProdStatusID  | integer | Product Status ID. Expected values: - 1:Awaiting Configuration - 2:Service Active - 3:Billing Failed - 4: Pending Renewal - 5: Cancellation Pending - 6: Service Canceled - 7: Pending Expiration - 8: Service Expired - 9: Service Deleted |
| ProdStatusDesc | string | Product Status Description. See above.                                                                                                   |
| ProdEnabled  | boolean | User control for WHOIS \(WPS\) visibility. The registrant must consent to the GDPR Whois Publicity before this flag is honored by the system. Expected values: - True - False                               |
| ProdConsented | boolean | Flag or indicator to show if the registrant has consented to the GDPR Whois Publicity Expected values: - True - False                                                            |
| CreateDate   | datetime | Value added service **create**date time. Format: MM/DD/YYYY HH:mm:SS AM/PM \(Pacific Time\) UTC and Epoch times are also available in the output.                                             |
| UpdateDate   | datetime | Value added service **last update** date time. Format: MM/DD/YYYY HH:mm:SS AM/PM \(Pacific Time\) UTC and Epoch times are also available in the output.                                          |
| ExpDate    | datetime | Value added service **expiration**date time. Format: MM/DD/YYYY HH:mm:SS AM/PM \(Pacific Time\) UTC and Epoch times are also available in the output.                                           |

```
https://resellertest.enom.com/interface.asp?command=VAS_GetDetail&uid=YourAccountID&pw=YourApiToken&VASItemID={Required}&ProductType=WhoisPublicityService&responsetype=xml
```
```
<?xml version="1.0" encoding="utf-8"?>
<interface-response>
  <VAS_GetDetail>
    <ProductType>WhoisPublicity</ProductType>
    <VASItemID>1017648</VASItemID>
    <DomainName>postman-reseller-20181114t214351982z.com</DomainName>
    <DomainNameID>366376964</DomainNameID>
    <AutoRenew>false</AutoRenew>
    <ProdType>74</ProdType>
    <ProdStatusID>2</ProdStatusID>
    <ProdStatusDesc>Service Active</ProdStatusDesc>
    <ProdEnabled>True</ProdEnabled>
    <ProdConsented>True</ProdConsented>
    <CreateDate UTC="2018-11-14T21:43:55.000Z" Epoch="1542203035">11/14/2018 1:43:55 PM</CreateDate>
    <UpdateDate UTC="2018-11-14T21:43:55.000Z" Epoch="1542203035">11/14/2018 1:43:55 PM</UpdateDate>
    <ExpDate UTC="2019-11-15T05:43:00.000Z" Epoch="1573767780">11/14/2019 9:43:00 PM</ExpDate>
  </VAS_GetDetail>
  <Command>VAS_GETDETAIL</Command>
  <APIType>API.NET</APIType>
  <Language>eng</Language>
  <ErrCount>0</ErrCount>
  <ResponseCount>0</ResponseCount>
  <MinPeriod>1</MinPeriod>
  <MaxPeriod>10</MaxPeriod>
  <Server>krkdt198</Server>
  <Site>eNom</Site>
  <IsLockable/>
  <IsRealTimeTLD/>
  <TimeDifference>+0.00</TimeDifference>
  <ExecTime>0.691</ExecTime>
  <Done>true</Done>
  <TrackingKey>2b97ff4b-6641-401a-8b06-6dbaa4811c32</TrackingKey>
  <RequestDateTime>10/25/2018 1:37:29 PM</RequestDateTime>
```
Related Commands
----------------
- [VAS\_Verification](../docs/vas-verification.md)
- [VAS\_GetList](../docs/vas-get-list.md)
- [VAS\_GetDetail](../docs/vas-get-detail.md)
- [VAS\_Update](../docs/vas-update.md)
- [VAS\_Add](../docs/vas-add.md)
- [VAS\_Delete](../docs/vas-delete.md)
- [VAS\_Cancel](../docs/vas-cancel.md)
- [VAS\_GetPricing](../docs/vas-get-pricing.md)
- [Purchase](../docs/domains/registration/purchase.md)
- [AddToCart](../docs/add-to-cart.md)
- [PurchaseServices](../docs/purchaseservices.md)
- [GetHosts](../docs/domains/domain-management/host-records/gethosts.md)
- [SetHosts](../docs/sethosts.md)